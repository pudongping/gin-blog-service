// Code generated for package configs by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.yaml
package configs

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

// Mode return file modify time
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

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\x5d\x4f\x1b\x47\x14\x7d\xdf\x5f\x71\xa5\x3c\x63\x16\x08\x24\x9a\xb7\x84\x10\x85\x0a\xb7\x56\xed\x28\x8f\xd5\x80\x87\x65\xab\xf5\xce\x32\x33\x4b\x36\x7d\x22\x95\x29\x49\x14\x30\x55\x70\x4c\x09\x69\x8b\x0a\x0a\xfd\xc0\x76\xab\x34\x45\xb6\x81\x3f\xe3\x99\x5d\x3f\xe5\x2f\x54\xb3\x6b\xbc\x76\x12\xe5\x6d\x67\xef\xbd\xe7\x9c\x7b\xe7\x9e\xc9\x13\xb6\x46\x18\x82\x6b\xa0\x0e\xb6\xe4\xb3\xc3\xde\xc6\x56\x78\x5e\x37\x00\xbe\xf6\xdd\x2c\x2d\x12\x04\x45\xb2\xe8\x5b\x00\xd7\x20\xaa\x5f\x84\xe7\x75\xb0\x6c\x17\xc2\xfd\x72\x74\xb9\x13\x1d\x3e\x57\x27\x87\xb2\x53\x31\x00\xee\x09\xe1\xe5\x28\x13\x08\x6e\x9a\xa6\xa9\xd3\x7b\xed\xbd\xa8\x7e\x14\xee\x97\xe1\x5e\xa1\x90\x83\xf0\xd5\x8f\x72\xe7\xaf\xf0\xcf\x86\xac\xfc\xa6\xe1\x09\x2e\x16\xec\x12\xa1\xbe\x40\x30\x13\x17\xc8\x8d\xc7\x51\xfd\x2c\x6a\xb4\x65\xe5\x65\xb8\x5f\x56\x07\xeb\xf2\xe8\x8d\x7a\xfe\x38\x6c\x9f\xaa\xda\xbb\x5e\xed\xad\x01\xf0\x80\xd9\x82\x7c\xb2\x4e\xfe\xf0\x93\xdc\x38\xfe\x74\xdd\x2d\xcf\xd3\x1d\xca\xd6\x6e\xb8\x7b\x32\xe8\xf0\x0e\x59\xc6\xbe\x23\x72\xd8\x22\x79\xfb\x3b\x82\x60\xc2\x4c\xdb\x4c\xd4\xab\x46\xa5\x77\xf8\xaf\xaa\x36\x7b\x9b\xba\xc9\x2c\x0e\x86\xb3\xe3\x74\xf5\x74\x3d\x11\x90\x52\x8f\x16\xf5\x69\x66\xa9\x2b\x48\x20\x86\xb4\x1b\x00\x0b\xd4\xca\xe3\x35\x92\xc3\x62\x05\x01\x17\x94\x61\x8b\x8c\x3b\xd4\xe2\x23\x03\x4c\x74\xab\xda\xb1\xbc\xac\xc9\xd3\x3d\xf9\xfd\x49\xf4\x5f\x43\x5e\x94\x13\x80\xbb\xb6\x43\xbe\xc4\x25\x82\x00\x7b\x5e\xfa\x6b\x2e\x10\x08\x32\x0e\xb5\x0c\x80\xfb\x9e\x43\x71\xf1\x63\x26\x3f\xfe\x1f\x93\x75\xcf\x9e\x75\x3b\xbf\xaa\x97\x9b\xdd\xf6\xbb\xa4\x93\xb0\xfd\xa4\x7b\xf9\x5a\x9e\xee\x85\xaf\xea\xf2\xbc\x9a\xc2\xc4\x1b\x73\x9f\x39\x08\x56\x84\xf0\xd0\xf8\xf8\xc4\xe4\x8d\x8c\x99\x31\x33\x13\x48\xdf\xfd\x38\x17\x58\xd8\x4b\x1f\x82\xca\x9d\xed\x70\xbf\x1c\xee\x9e\x74\x5b\xdb\xf2\xef\x6a\x78\xd4\xd2\x34\x71\x28\xd9\x3c\x79\xd0\x94\xaf\xd7\x07\x34\xf3\x25\x6c\x91\x2c\x0e\x92\x61\x4f\x7f\x08\xf7\xf1\xd8\xc3\xdf\x5b\xbd\xda\x5b\x79\xf4\x46\x36\x2b\xef\x3b\x4f\xb2\xb7\xdf\x77\x9e\x8e\xa2\xdd\x72\x1c\xfa\x70\x2e\x10\x1c\x7d\x16\x6d\x20\xb7\xa3\xd5\x00\x8c\x41\xe6\x5b\xcf\x4a\x3f\xc9\xe0\xdb\x73\x2d\x63\xae\x84\x6d\x07\x69\x07\x50\x2e\x10\xf0\x92\xf0\x32\xab\xab\x99\x25\x5a\x32\x00\x12\x47\x5c\x9f\x99\xd6\x42\x38\x61\xc9\x3d\x05\x41\xa0\x63\x98\xf3\x87\x94\x15\xaf\xce\xf3\x3c\x9f\x5f\x40\x20\x98\x4f\x0c\x80\xbb\x8c\x96\xae\x22\x05\x8a\xfa\x8c\xfa\xfc\xc5\x83\x82\x3e\xe6\xc9\x12\x23\x02\x01\x76\x48\x52\xcd\x7d\xed\x63\xcb\x76\xc7\x16\x1d\x6a\x8d\x71\xc2\xd6\xec\x25\x0d\x35\x17\x78\x36\x23\x08\x6e\x4c\x9a\xa6\x71\x07\x0b\xbc\x88\x39\x89\x1d\x5f\x6d\xaa\xad\xba\x6c\xbd\x48\x2d\x71\xbb\xf0\xc8\x23\x08\x4a\x8f\xf8\xaa\x33\xa2\x99\x51\x2a\x46\x44\x4f\x4c\x4e\x5d\x9f\x9e\x19\xf4\x9d\xee\xc0\xd4\x94\x39\x13\x43\x25\x85\x5a\xcc\x37\xa9\x98\x02\x5e\x74\x48\x8e\x91\x65\x3b\xe8\xc7\x0c\x80\xd9\x15\xcc\xb8\x6e\xc6\x17\xcb\x37\x63\x16\xc6\x63\x97\x23\x28\x24\xe3\xc8\xe2\x60\xbe\xe8\x90\x59\xea\xba\xfc\xca\xa8\x43\xd7\xfe\x4f\x74\xf9\xb3\xda\x3e\x56\xd5\x66\x92\xfb\x95\x47\xdc\x7e\xee\xd4\x50\xee\x70\x96\x0e\x67\x71\xb0\x60\x2f\x13\x11\x33\xc5\x2b\xa6\x1a\x95\xee\xd9\x1f\xbd\x17\x17\x6a\xfb\x38\x7e\xe4\x36\xd5\xc1\x2f\xfd\x57\xe4\xff\x00\x00\x00\xff\xff\x46\x3d\x3d\x0c\x2b\x05\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 1323, mode: os.FileMode(420), modTime: time.Unix(1642249081, 0)}
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
	"configs/config.yaml": configsConfigYaml,
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
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
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
