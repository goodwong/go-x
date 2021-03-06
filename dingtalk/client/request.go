package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
	"time"
)

const defaultHTTPClientTimeout = time.Duration(2 * time.Second)

// New new
// Usage 1，自定义获取token的函数：
//	client.New(&client.Config{
//		TokenFunc: getTokenFunc,
//	})
// Usage 2, 使用默认的token函数：
//	client := client.New(&client.Config{
//		TokenAPI: tokenAPI, // tokenAPI自己替换里面的APPKEY&APPSECRET等内容
//	})
func New(cfg *Config) *APIClient {
	client := &APIClient{
		HTTPClientTimeout: defaultHTTPClientTimeout,
		mu:                &sync.RWMutex{},
	}

	if cfg.TokenFunc != nil {
		client.tokenFunc = cfg.TokenFunc
	} else if cfg.TokenAPI != "" {
		client.tokenAPI = cfg.TokenAPI
		client.tokenFunc = client.defaultAccessTokenFunc
	} else {
		panic("client.New() miss cfg.TokenFunc or cfg.TokenAPI")
	}

	return client
}

// APIClient 类型
type APIClient struct {
	HTTPClientTimeout time.Duration
	mu                *sync.RWMutex
	accessToken       *accessTokenStruct
	tokenAPI          string
	tokenFunc         func() (accessToken *accessTokenStruct, err error)
}

// Config config
type Config struct {
	TokenAPI  string
	TokenFunc func() (accessToken *accessTokenStruct, err error)
}

// 封装的request
// 做两件事：
// 1. 自动替换ACCESS_TOKEN
// 2. ACCESS_TOKEN过期自动重试
func (client *APIClient) request(method, url, contentType string, body io.Reader, result interface{}) error {
	retryFn := func() (result interface{}, retry bool, err error) {
		if strings.Contains(url, "ACCESS_TOKEN") {
			// 替换access_token
			token, err := client.GetAccessToken()
			if err != nil {
				// GetAccessToken 本身会遇忙重试
				return nil, false, fmt.Errorf("获取AccessToken失败：%s", err)
			}
			r := strings.NewReplacer("ACCESS_TOKEN", token)
			url = r.Replace(url)
		}

		// request
		httpClient := newHTTPClient()
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			return nil, true, err
		}
		req.Header.Set("Content-Type", contentType)
		// debug(httputil.DumpRequestOut(req, true)) // 打印请求、响应的原文格式，这功能老牛逼了
		var resp *http.Response
		resp, err = httpClient.Do(req)
		if err != nil {
			return nil, true, fmt.Errorf("请求失败：%s", err)
		}
		// debug(httputil.DumpResponse(resp, true)) // 打印请求、响应的原文格式，这功能老牛逼了

		// response
		respBytes, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, true, fmt.Errorf("读取失败：%s", err)
		}

		// 验证
		contentType := resp.Header.Get("Content-type")
		status := resp.StatusCode
		var errInfo responseErrorInfo
		errInfo, err = checkResult(status, contentType, respBytes)
		// 需要重试的情况
		switch errInfo.Errcode {
		case -1:
			return nil, true, errors.New(errInfo.Errmsg) // 服务器忙
		case 40001, 40014, 41001, 42001:
			log.Println("access_token失效！", err)
			if _, err = client.RefreshAccessToken(); err != nil { // accessToken 失效，重新获取，并重试
				return nil, false, fmt.Errorf("刷新access_token失败：%s", err)
			}
			return nil, true, fmt.Errorf("access_token失效：%s", err)
		}
		// 其它错误，返回
		if err != nil {
			return nil, false, fmt.Errorf("API返回错误：%s", err)
		}

		// 正常：返回
		return respBytes, false, nil
	}

	resp, err := retry(retryFn, 3) // 重试3次
	if err != nil {
		return err
	}
	respBytes, _ := resp.([]byte)

	// 解析
	switch v := result.(type) {
	case nil:
		// do nothing
	case *[]byte:
		*v = respBytes
	default:
		err = json.Unmarshal(respBytes, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Get 请求
func (client *APIClient) Get(url string, result interface{}) error {
	return client.request("GET", url, "application/json", nil, result)
}

// PostJSON 请求
func (client *APIClient) PostJSON(url string, payload interface{}, result interface{}) error {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false) // 禁用转码
	_ = encoder.Encode(payload)

	return client.request("POST", url, "application/json", buffer, result)
}

// PostBlob 请求
func (client *APIClient) PostBlob(url, field string, blob []byte, result interface{}) error {
	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)
	part, err := writer.CreateFormFile(field, "media.png")
	if err != nil {
		return err
	}
	part.Write(blob)
	err = writer.Close()
	if err != nil {
		return err
	}

	return client.request("POST", url, writer.FormDataContentType(), buffer, result)
}

type retryFn func() (result interface{}, retry bool, err error)

func retry(fn retryFn, attempts int) (result interface{}, err error) {
	for {
		result, tryAgain, err := fn()
		attempts--

		// 成功
		if err == nil {
			return result, nil
		}

		// 失败，重试
		if attempts > 0 && tryAgain {
			continue
		}

		// 重试机会用尽，或无需重试
		return nil, err
	}
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Timeout: defaultHTTPClientTimeout,
	}
}

type responseErrorInfo struct {
	Errcode int
	Errmsg  string
}

func checkResult(status int, contentType string, readBytes []byte) (responseErrorInfo, error) {
	var err error
	var result responseErrorInfo

	switch {
	case strings.HasPrefix(contentType, "image/"):
		// 文件媒体类型
		// 不解析内容
	case strings.Contains(contentType, "application/json"):
		fallthrough
	case strings.Contains(contentType, "text/plain"):
		// json解码
		err = json.Unmarshal(readBytes, &result)
		if err != nil {
			return result, err
		}
		// 业务检查
		if result.Errcode != 0 {
			return result, errors.New(result.Errmsg)
		}
	}

	// 检查status code
	if (status < 200 || status >= 300) && err == nil {
		err = fmt.Errorf("http状态码错误：%d", status)
	}
	return result, err
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}
