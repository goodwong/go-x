// Code generated by go-bindata. DO NOT EDIT.
// sources:
// password_login.html (7.649kB)

package templates

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _password_loginHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4d\x8f\xdb\x36\x13\xbe\xeb\x57\xcc\xab\x04\xb1\x0d\x78\xa5\xfd\x78\x17\x08\x14\xc9\x45\x90\x2c\x50\x14\x41\x13\x20\xdb\x43\x91\x2c\x1a\x5a\x1a\x5b\x6c\x68\x52\x25\x29\x7b\x8d\x85\x81\x9e\x7a\x6d\x0f\xbd\xe4\x5e\x14\xbd\xf4\xd4\x5b\xf2\x77\x8a\x04\xf9\x17\x85\x44\x7f\xd1\x92\xbc\x1f\x69\x0b\x24\xb5\x0e\x0b\x9b\x33\xf3\xcc\x0c\xe7\x99\x21\xbd\x0a\xff\xf7\xf0\xf1\x83\xd3\xaf\x9f\x9c\x40\xaa\x47\xac\xe7\x38\x61\x8a\x24\xe9\x39\x00\x00\xa1\xd2\x53\x86\xe6\x73\xf1\x0c\x84\x1c\xf5\x12\x3a\x86\x8b\xe5\x52\xf1\x64\x24\x49\x28\x1f\x06\x70\x9c\x9d\xc3\xfe\xbd\xa5\x6c\xe6\x58\x96\xc0\x48\x1f\xd9\x86\x6d\x42\x55\xc6\xc8\x34\x00\xca\x19\xe5\xb8\xd7\x67\x22\x7e\x79\xcf\x52\x99\xd0\x44\xa7\x01\x1c\xe3\xc8\x5e\xd7\x78\xae\xf7\x08\xa3\x43\x1e\x80\xa4\xc3\x54\xdb\xe2\x31\x4a\x4d\x63\xc2\x16\x2a\x5a\x64\xf7\xea\xc3\xde\xcf\xce\xe1\xa0\xf8\xb3\xbf\x3d\x7e\xca\xb3\x5c\x77\xed\xb5\x22\x08\x22\x91\x34\x6d\xc9\xdd\x39\x76\x6d\x4a\x87\x95\x9c\xfa\x42\x26\x28\x03\x38\xc8\xce\x41\x09\x46\x13\xb8\x45\x08\x69\x0e\xa9\x9f\x6b\x2d\xf8\xf5\x9c\xf7\x49\xfc\x72\x28\x45\xce\x93\x00\x6e\x0d\x06\x03\x5b\x3a\xa2\x7c\x6f\x1e\xde\xdd\xeb\x47\x57\xfc\x0d\xfd\x39\x6d\x42\xdf\x50\xc9\x09\xfb\x22\x99\xce\x29\x95\x1e\xf5\xde\xbd\x7a\xf3\xfe\xd5\x0f\x41\xe8\xa7\x47\xf3\xc5\x32\x97\x98\x11\xa5\x22\x97\x89\x21\xe5\xee\x8a\x74\x61\x42\xc7\x3d\x2b\x8c\xb0\x24\x52\x2f\x97\x2c\x08\x7d\xf3\xd9\x96\x2f\x8b\xc2\xc9\x08\x23\x37\x97\xcc\xed\x85\xfe\x62\x75\x0d\xda\xb7\xb0\x1b\x3d\x8d\x50\xa7\x22\x69\x70\x56\xb2\x62\xee\xc9\x28\xba\x05\xab\x49\x9f\x61\x02\x63\xc2\x72\x8c\xdc\x27\x8f\x9f\x9e\xba\xe0\x5f\xdf\x73\x42\x34\xb9\x5a\x92\x85\xa6\x0b\x52\x4c\x54\x74\xdc\x73\x0c\x21\xdc\x5c\xa1\x2c\xc4\x6e\x00\x6e\x82\x23\x51\x7c\x77\x0d\x83\xdd\x8c\x28\x35\x11\x32\x59\xc8\xbe\x52\x28\x0f\x0e\x8f\xba\xae\x33\xfb\x90\xbd\xaa\x0f\x76\xce\x53\x3d\xcd\x30\x72\x55\xde\x1f\x51\xed\xf6\xde\xfe\xf8\xd3\x9f\xaf\x7f\x09\x7d\x23\xac\xb8\xba\xdc\x97\x44\x95\x09\xae\x10\x94\x26\x3a\x57\x57\x28\x90\x51\xac\x16\xe8\x46\xc5\x59\xb8\xbf\x5a\x81\x16\xda\xf3\x22\xfd\x7f\x19\xc3\x65\xd4\x0c\xfd\x72\xee\x3a\xab\xfe\x79\xf3\xfb\xbb\xd7\xbf\x36\xf5\x8f\x44\x8e\x93\x4f\xad\x7f\xee\x9f\x3e\xf8\xfc\xdf\x6b\xa0\xa3\x65\x03\x49\x1c\x48\x54\xe9\x37\x5a\xbc\x44\x5e\x74\x8a\xeb\xcc\x6c\x94\x5d\xab\xfc\xe3\xad\x72\x7c\xf3\x56\x79\xfb\xc7\x6f\xef\x7f\xfe\x7e\xcb\x51\x23\x72\xfd\x89\xf5\xca\xc3\x93\x47\x27\xa7\x27\x37\xaa\xd2\x8e\x90\x0d\x55\xfd\x9b\x08\xa9\x62\x49\x33\xbd\x52\x64\xa8\x21\xc1\x31\x8d\x11\x22\x68\x4d\xb0\xdf\xb2\x44\xb9\x64\x0a\x22\x48\x44\x9c\x8f\x90\x6b\xef\xbb\x1c\xe5\xf4\x29\x32\x8c\xb5\x90\xf7\x19\x6b\xb7\x16\xbe\x9f\x95\x61\xe6\x92\x9d\xb5\x3a\x4b\x88\xc2\xfc\xd9\xfe\x99\x57\x6e\x13\x44\xf0\xe2\xf6\x05\x13\x31\xd1\x54\x70\x2f\x95\x38\x98\x7d\x96\x49\x31\xa6\x09\xca\x68\x71\x17\xb8\x23\x71\x84\xa3\x3e\xca\xe8\xe0\x8e\x09\x2c\xba\x7d\x61\x3e\xcc\x5e\xd8\xc0\x07\x2b\x60\x0b\xd5\xd6\x3a\xdc\xe6\xbe\xea\xc1\xba\xdc\x5e\x92\x7b\xa1\xb2\x96\xed\x98\x48\x28\xaf\x8e\x10\x19\xeb\x67\xfb\x67\x96\xb0\x3c\x17\x97\xc2\x83\xb3\x4d\x4b\x91\xeb\xa5\xf4\x70\x25\x2d\x31\x3d\x92\x24\x27\x63\xe4\xfa\x11\x55\x1a\x39\xca\xf6\xa2\x11\xba\x30\xc8\x79\x5c\x24\x05\x6d\xec\x6c\x5c\xc4\xd1\xcb\x24\x16\x66\x0f\x71\x40\x72\xa6\xdb\x1d\xc7\x92\xcf\x8b\x0c\x11\xe8\x94\x2a\x3b\xc7\xfa\xe2\x9a\xdd\xac\x80\x98\x59\xd0\x80\x53\xf6\x90\x01\x31\x7a\xcd\x38\xc5\x19\x08\x11\x7c\xf1\xf4\xf1\x97\x5e\x46\xa4\xc2\xf6\xe5\x81\x15\x36\x4b\xc4\x4e\x05\xd2\x74\xed\xe5\xa1\x19\xbd\x75\xfe\x2e\x10\x96\x93\xe2\x2a\xdb\xb4\x50\xae\x07\x2a\x0f\x73\x88\x0c\x17\xae\x94\x96\x5d\x30\x13\xe4\x92\xd2\x2d\xcf\xf3\x5a\x96\xc2\xc2\xfd\x86\x8a\xa5\x33\x40\x1d\xa7\xed\x5c\xb2\xee\x06\x5d\x8a\xc7\x54\xa8\x5b\x59\x8f\x25\x26\xc8\x35\x25\x4c\x05\xe0\x52\x1e\xb3\x3c\x41\xb7\xaa\x57\xfc\xe2\x42\xa9\x82\x1a\xe8\xe2\x69\xdd\x8f\x63\xcc\x74\x2b\x80\x16\xc9\x32\x46\x4d\x3f\xfa\xdf\x2a\xc1\x5b\x55\xb0\xd2\xe2\x81\xe0\x1a\xb9\xde\x3b\x9d\x66\x58\x6b\x57\x31\x9b\x55\x91\x8a\x1f\x80\x81\xe1\x95\xd2\x92\xf2\x21\x1d\x4c\xdb\xc5\x16\x77\x6c\xdd\x59\xa7\x62\xea\xe9\x14\x79\x5b\x42\xd4\x6b\xc8\x29\x16\x5c\x09\x86\x1e\x13\xc3\xb6\xac\xda\x43\xb5\x6e\x2f\x6e\x5f\x48\xcf\xac\xcd\x02\x58\x7d\x39\xc5\x73\xbd\x36\xe7\xd6\x1f\x89\x3a\x97\x1c\x64\x35\xd9\xad\x11\x4b\xaf\x60\x55\xbb\x73\x83\xb4\xb4\x9c\x36\x48\xca\x78\xec\x36\x6d\xc8\x1b\xea\x28\xb9\x51\x05\xd9\x05\x9e\x33\xd6\x85\xc3\x6d\x18\x0d\xd9\x97\x3b\x00\x31\xd1\x71\x5a\x33\xff\xb6\x86\x51\x0f\x56\x66\x9e\x4a\x31\x01\xac\x77\x76\xcd\x02\xd4\x07\x34\x9f\x04\x2b\x42\x5c\x3c\xe7\x75\xb7\xfd\x82\x1a\xd6\xe2\xcc\x7d\xce\x6b\x18\x32\x5b\x9b\x13\xb3\x2e\x0c\x08\x53\xb8\xb6\x64\xc6\xcd\xee\x0c\xf9\x08\xcf\x90\xdd\xf8\xdf\x8d\xff\xdd\xf8\xff\x38\xc6\xff\xfa\x14\x36\x17\xfa\xff\xca\xcc\xdd\x0d\xc8\x8f\x6a\x40\xee\x66\xde\x6e\xe6\x6d\x64\xfe\x01\x33\x0f\xcc\x3b\xb7\xf9\xff\x96\x9c\xd0\x37\x6f\xdb\x9c\xd0\x2f\xdf\xe8\xfe\x15\x00\x00\xff\xff\x88\xcb\x64\xd6\xe1\x1d\x00\x00")

func password_loginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_password_loginHtml,
		"password_login.html",
	)
}

func password_loginHtml() (*asset, error) {
	bytes, err := password_loginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "password_login.html", size: 7649, mode: os.FileMode(0644), modTime: time.Unix(1563956262, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8f, 0xb6, 0x8c, 0x83, 0xf, 0x51, 0xeb, 0xd1, 0xce, 0xc3, 0xf0, 0x21, 0x57, 0x77, 0xc0, 0xbe, 0x1, 0x48, 0xe9, 0x86, 0x89, 0xc4, 0x6d, 0x7c, 0x1a, 0xa3, 0xb2, 0x0, 0xd9, 0xf4, 0xab, 0x8f}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"password_login.html": password_loginHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"password_login.html": &bintree{password_loginHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
