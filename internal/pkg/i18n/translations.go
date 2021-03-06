// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package i18n generated by go-bindata.// sources:
// resources/i18n/en.toml
// resources/i18n/pt_BR.toml
package i18n

import (
	"bytes"
	"compress/gzip"
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
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _resourcesI18nEnToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x51\x8b\x14\x49\x12\x7e\xf7\x57\x84\x03\xc7\xdd\x41\x5b\xfd\x3e\x70\xc8\x70\x7a\x77\x03\x3a\xca\x38\x22\x9e\x25\x12\x5d\x19\xdd\x95\x37\x99\x19\x75\x99\x91\xdd\x96\xaf\x8b\xb0\xc8\x82\x82\xb0\xb0\x8b\x8b\xfe\x01\x9f\xf6\x65\x7f\xcf\xfc\x81\xf5\x27\x2c\x99\x55\xdd\x53\xd5\x53\xe3\xca\x3c\x0c\xd9\x19\xdf\x17\x91\x11\xf1\x45\xc5\x33\xed\xb4\x14\x95\x55\x85\xa2\x50\x79\xdd\x88\x66\xf7\xfc\x06\x4b\x4d\x1e\xfe\x01\x07\xc7\x4e\x8b\x46\xa3\x5f\x11\x78\x2d\x50\xb1\x5b\xea\x55\xf4\x98\xcc\x0e\x6e\xdc\xe8\xe0\x1b\x32\x15\x5b\x1a\xc0\x4e\xb5\x54\xb5\x26\xd0\x01\x10\x1a\x83\xb2\x64\x6f\x41\x6a\x14\xa8\xc9\x34\x01\x5a\x8e\x80\x4e\xa5\xff\x1e\x84\xd0\x82\x30\x04\x5c\x13\x88\xb6\x04\x8b\xb6\x74\x2b\xbd\xd6\x6e\x95\x2d\xa5\x26\x68\x78\x43\x3e\x59\x55\x9e\x50\xfa\xf3\x32\x1a\x10\xb2\xc9\x03\x85\x74\x49\x2f\xa9\x8a\x42\xa0\x6d\xc3\x5e\xd0\x49\xe9\x04\xc3\x79\x00\xac\x3c\x87\x30\xf0\x97\xbc\xb3\x5f\xa1\xd3\xaf\xf2\x6b\x60\xa3\xa5\x06\xab\x9d\xb6\xd1\x76\x51\x24\x93\xfc\x6b\x10\x74\x0a\xbd\x0a\xb3\xd2\x29\x32\x7a\x4d\x3e\x85\x86\x51\xd8\xb1\x6d\x93\x63\x45\x6b\x32\xdc\x90\x0f\x3d\x84\xaa\xe8\xb5\xb4\x45\xe9\x4a\xf7\x94\x23\x54\xe8\x60\xad\x69\x03\x29\x82\x87\x5e\xaf\xb1\x6a\xe1\x21\x1b\x5d\xb5\xf0\xb7\x5a\xa4\x39\x9c\xcf\xb5\x0b\x7a\x55\x4b\x28\x5e\xc5\xa6\xa8\xd8\x16\x0b\x3f\x6f\xd8\x68\xd1\x15\xde\x6a\x32\x46\x2b\x54\xf4\xf7\xe4\x71\x41\x22\xe4\x21\x3a\x45\x3e\x07\x98\x99\x2b\xb6\x56\x8b\x25\x27\xd9\xf3\xae\x46\x21\x56\x15\x85\xb0\x8c\x66\x50\xa6\xd2\x5d\x7c\x78\x0d\xbb\x1a\x77\x79\xb8\xb4\xbc\x39\x89\x7f\xb1\xd1\x4e\xf1\x26\x8c\x78\xbe\x8d\x64\x83\xde\x69\xb7\x1a\x21\x4b\x49\x7f\x17\x3f\x7f\xfa\xfd\xb7\xb7\x00\x4f\x8e\x4e\x4f\x8e\x4f\xfe\x0d\xdd\x79\x1f\x38\xe9\xba\x23\xd8\x22\x77\x90\x94\x08\x76\xa1\xa8\x09\x15\xf9\x01\xe2\xe2\xf3\x3b\xf8\x67\x77\x09\x9e\x1a\x0e\x5a\xd8\xb7\x70\xf1\xf9\xdd\x20\xd2\x31\x7a\xc2\xef\x04\xc3\x04\x3c\xdd\x6e\x53\x37\x0c\xe1\xc3\xeb\xa9\x10\x50\x29\x52\x83\xd4\x99\xf6\xe6\x9f\x70\x7e\x5b\x60\x5f\xe7\x45\xa5\xc6\xdc\x49\xa1\xa3\x0c\x9f\x46\x07\xe5\x41\xd2\x3f\x2a\x95\x89\xcb\x83\xd4\x81\xe9\x84\xe0\x68\x33\x74\x66\xd1\x45\x34\x26\x35\xfe\xf5\x2e\xb4\x5b\xf2\xc8\xc5\x56\x21\xe7\x44\x4d\x96\xfb\x68\xd2\x64\x49\x71\xcc\xfe\x93\xf0\x3a\x03\x6b\xa3\xd3\xd2\x0e\x9c\xcf\x4a\xb7\x88\x92\x47\xc6\x46\x1b\x03\x8e\x48\xa5\x40\x1b\xcf\x6b\xad\x08\x10\x56\x5a\xb2\x7d\x27\xd2\xc4\x93\x26\x53\x34\x18\x06\x73\x24\x69\x29\xbd\x4d\x6a\xb2\xd9\xb0\x74\x57\x9e\x9f\xdc\xa3\x53\x33\x70\x68\xbb\x90\xf4\x28\xe7\xbc\x30\x7a\x85\xc2\x5e\x9b\x16\x30\x40\x79\xd0\x27\xa0\x3c\xc8\xc2\x7c\x44\x04\x35\x6f\xf2\xe8\xe0\x0e\xcd\x2e\x07\x44\x2f\xd1\x36\x86\x0e\x4b\xf7\x2c\x8d\x85\x70\x38\x9f\xaf\xb4\xd4\x71\x91\x52\x38\xff\x6f\x6c\x8e\xcf\xe6\xbe\x9b\xaf\xb7\xb6\xc1\xcf\x17\x86\x17\x73\x8b\x41\xc8\xcf\x77\x0f\x99\x77\xa3\xf2\x45\x6f\x35\x3f\xbd\x7b\x74\xe7\xfe\xdd\xc2\xaa\xe7\x5f\xab\xcd\xff\x23\x85\xbd\xaf\xc0\x13\x8e\x26\xcf\x6a\x30\xfa\x9c\xb6\xa5\xbf\xae\x0a\xb7\xaf\xe7\x26\xef\x79\xa8\xc5\x63\x81\x0d\x06\x70\x2c\xd0\x70\x08\x7a\x61\xae\xb0\xef\xf7\xb2\x74\xc9\x4a\x03\x7a\x06\x8d\x21\x0c\x04\x92\x2e\x56\xa8\x1d\xa4\x77\xfb\x62\x18\x80\x76\xab\x51\x0c\x03\xef\x47\xe3\x6e\x1a\x7b\x2a\x8a\x4b\x1a\x89\xa9\x90\x68\xae\x0e\x93\xd2\x7d\xf9\xf8\xfe\x47\x38\x7b\x7c\xf6\xe0\xf4\xf8\xe8\x1e\xa4\xd3\x20\xb9\x7b\xc0\xe9\xd1\xb9\x05\x4f\xe1\x44\x8b\x19\x7e\x58\xff\xd3\xb5\x4c\xff\x0d\x4c\xd2\xdb\xb6\xc0\xe1\x55\xf0\x82\x55\x3b\x9e\xf6\xdf\xff\x04\x97\x62\xee\x49\x7a\x82\xb2\xbf\x7f\xd0\x50\xd7\x87\x8d\xe7\xff\x51\x25\x9d\x52\xf2\x77\x73\x89\x6b\xf6\x5a\x08\x84\x5e\x0a\x90\x4a\x79\x2a\xa6\xa2\xee\xb3\x39\xf0\x7d\x86\xe7\x49\x7e\x86\xf9\xbc\x2b\xe1\x40\x78\x6d\xaf\x7d\x1f\x5d\x96\x9e\x50\x90\xbc\x0e\x10\xc1\x26\xad\x0c\x5b\x83\x1c\xc9\x76\xb5\x88\x21\xd5\x2e\xbf\x24\x85\xbe\x0b\xc2\x92\x78\x5d\x5d\xad\xd4\x97\x8f\xef\xdf\xc0\xfd\x7c\x19\x52\x99\xde\x5c\x87\x99\x28\x52\x0f\xdb\x53\x4d\x8f\xda\x9b\x65\x67\x9c\x57\x1c\x88\x21\x2d\x21\x9e\xd7\xdd\x1e\xd1\xaf\x0d\x60\xd9\x13\xac\xd1\xc4\xdc\xe8\x23\x09\xcd\x4a\xb7\xa1\x6e\x70\x55\x6c\x4c\xca\x3d\x3a\x76\xad\xe5\x18\x40\xa1\x20\xe0\x22\x0d\xc1\xc6\xb3\x8a\xf9\x52\xc1\x92\x50\xa2\xa7\xd2\xc5\x40\x69\x4d\x11\x1d\x24\x3d\x30\xdd\x55\x1e\x43\x9d\x3b\xda\x4b\x28\xa6\x42\x9f\x90\xfa\x1d\xee\x76\xb3\x95\x27\xba\xbd\x9f\xa0\x25\xb3\xec\xb5\xff\x76\x70\xa3\xd9\x60\x1b\xc0\xb2\xd2\xcb\xb6\x6b\x97\xaa\x66\x5d\x6d\x0b\x95\x5e\xda\xb5\x9d\xed\x93\xb9\x9b\xa1\x5d\x0b\xf5\x9e\x7c\x74\x85\xb4\x0d\x4d\x15\xf0\xed\x77\xf0\xaf\x07\xa7\xf7\x1f\xdf\x3b\x82\xd3\xc7\x27\x70\xf6\xf4\xe1\x5d\x48\xbf\x0e\xea\xb2\x07\x9f\xa8\xe5\x3e\xc3\x14\xd8\x70\x85\x66\x72\x5b\x39\x76\xc0\x5e\x75\x8b\x68\xea\xd7\x5d\x13\x67\x88\x69\x67\x39\x7b\x36\x06\x81\x3a\x6f\xb3\x97\x8d\x0e\x06\xdd\x2a\xe2\x8a\x40\xbb\x20\x68\x0c\xa9\x34\xf2\x73\xaa\x2c\x56\xb5\x76\x34\x2b\x9d\x5e\x66\x06\xc5\xee\xaf\x69\x36\xba\xac\x84\x1e\x90\x32\xca\x81\xb6\xbe\x47\x22\x4a\xab\xa3\xea\xfc\x29\xae\xce\x69\x24\xcc\xdd\xc3\x26\x0a\xfe\x88\xba\x4e\x03\x45\x4b\x8c\x46\x76\xe1\x66\x17\x6d\x43\x97\x83\x45\xbb\x35\x1a\xad\x2e\xe9\xf6\xc7\x79\x6f\x70\x85\x62\x96\xe2\x0a\xb4\x3b\x07\x40\x4f\x40\x0e\x17\x29\x09\xcf\xfe\xb2\x7e\xde\xf9\x68\xa2\x14\x58\x55\xd4\x48\xc1\x8d\xec\x2d\x4c\x4f\x29\x4c\x59\x4d\x94\x78\x64\xa9\xa8\x32\xda\xd1\x3e\xe1\x2f\x3f\xc0\x09\x4f\x5a\x4d\x10\x0e\x2d\xc7\x4d\x32\xee\xcf\x4f\x5d\x1b\x5c\x67\x3c\xc1\x7c\x9d\x7d\x57\xc3\x11\xfb\xbb\x5f\xfb\xca\x5e\x6b\x3e\xc1\xdf\x23\xfe\x08\x00\x00\xff\xff\xa7\x67\x78\x39\xe7\x0d\x00\x00")

func resourcesI18nEnTomlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesI18nEnToml,
		"resources/i18n/en.toml",
	)
}

func resourcesI18nEnToml() (*asset, error) {
	bytes, err := resourcesI18nEnTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/i18n/en.toml", size: 3559, mode: os.FileMode(420), modTime: time.Unix(1611341256, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesI18nPt_brToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xce\xcc\xcb\x2c\xd1\x4b\xce\x4d\xd1\x4b\x49\x2d\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x8b\xe5\xca\x2f\xc9\x48\x2d\x52\xb0\x55\x50\xf2\xcc\xcb\x4c\xce\x4c\xcc\xc9\xac\x4a\x2c\x52\x48\x2c\x56\x48\xce\xcf\x4b\xcb\x4c\x2f\x2d\x4a\x3c\xbc\xfc\xf0\xd6\xd4\x62\x85\x94\x7c\x85\xa2\xcc\x12\x25\x40\x00\x00\x00\xff\xff\xb7\x34\x31\xce\x46\x00\x00\x00")

func resourcesI18nPt_brTomlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesI18nPt_brToml,
		"resources/i18n/pt_BR.toml",
	)
}

func resourcesI18nPt_brToml() (*asset, error) {
	bytes, err := resourcesI18nPt_brTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/i18n/pt_BR.toml", size: 70, mode: os.FileMode(420), modTime: time.Unix(1611239673, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"resources/i18n/en.toml":    resourcesI18nEnToml,
	"resources/i18n/pt_BR.toml": resourcesI18nPt_brToml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"resources": &bintree{nil, map[string]*bintree{
		"i18n": &bintree{nil, map[string]*bintree{
			"en.toml":    &bintree{resourcesI18nEnToml, map[string]*bintree{}},
			"pt_BR.toml": &bintree{resourcesI18nPt_brToml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
