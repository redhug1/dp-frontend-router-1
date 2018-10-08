// +build production
// Code generated by go-bindata.
// sources:
// templates/error.tmpl
// templates/main.tmpl
// templates/partials/footer.tmpl
// templates/partials/header.tmpl
// redirects/redirects.csv
// DO NOT EDIT!

package assets

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

var _templatesErrorTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\xc1\x8e\xd3\x30\x10\xbd\xe7\x2b\x46\xbe\x3b\xd5\x5e\x51\x1a\x21\x01\x12\x48\x70\x5b\xc4\x71\x35\xb1\x27\xb5\x55\xc7\x63\x79\x26\xe9\x56\x4b\xfe\x1d\xa5\xb0\x2c\xad\x7a\xe2\xf4\xc6\xcf\x6f\x34\xe3\xf7\xdc\xf9\xb8\x80\x4b\x28\xb2\x37\x05\x0f\x64\x03\xa1\xa7\x0a\x92\x50\xc9\xf4\xcd\xbf\xf7\xa7\x8a\xa5\x50\xdd\xd8\xf0\x70\xa7\xe9\xe9\x49\xa3\x26\x82\x31\xcd\x12\xac\x0d\x98\x46\x6b\x07\x56\xe5\xc9\xf4\x2f\x2f\xd0\x7e\xaa\x95\x6b\xfb\x78\x11\xad\x6b\xb7\x0b\x0f\x7d\xd3\xed\x7c\x5c\xfa\xe6\x15\xef\xcc\x83\x82\x99\x92\xb5\x13\xd6\x9b\x85\x0e\x35\x7a\xbb\xa9\xee\xf1\x8e\x13\x78\x92\xa3\x72\xb1\x17\x62\x9c\x53\xb2\xa7\xe8\x35\x6c\x72\xac\x1a\x5d\xa2\xeb\xc6\x81\x9f\x61\xe0\x67\x6b\x0b\x7a\x4f\xfe\x77\x5d\xc9\x5b\x8f\xf5\x78\x7d\xb2\x56\xa8\x60\x45\x25\x6f\x13\x8d\x0a\x42\x4e\x23\x67\xf2\x66\x7b\xcd\x7f\x6f\xa9\x27\xb6\x1a\x62\xf5\x62\xfa\xe6\xcd\xb4\x8f\x24\xae\xc6\xb2\x4d\x80\x9f\x20\x38\xd2\xe7\xc7\x6f\x5f\x61\x5d\x9b\xae\xf4\x5f\x46\x38\xf3\x0c\xa2\x31\x25\xa0\xec\x78\xce\xba\xf9\x56\x79\x48\x34\x09\x94\x44\x28\x04\x1d\x42\xa8\x34\xee\xcd\x84\x31\x29\xbf\x3b\xd1\xd0\x3a\x9e\x26\xca\x2a\xef\x39\x4b\x7b\xe0\xa5\x9d\x8f\x06\x2e\x31\xee\xcd\x07\xce\x8a\x4e\xe1\xbb\x98\xde\xfd\xa9\x67\xe9\x76\xd8\xb7\xf0\x83\x00\x0b\x27\x3e\x44\x21\x18\xb9\x02\xe6\x33\xc4\xec\x38\x2f\x94\x23\x65\x47\xa0\x21\x0a\x4c\x78\x86\x80\x0b\x81\xc3\x59\xc8\xb7\xdd\xae\xfc\xcd\xfc\x16\xde\x22\xb9\xfe\x13\xaf\xf8\x2b\x00\x00\xff\xff\x38\x7a\x05\x98\xac\x02\x00\x00")

func templatesErrorTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesErrorTmpl,
		"templates/error.tmpl",
	)
}

func templatesErrorTmpl() (*asset, error) {
	bytes, err := templatesErrorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/error.tmpl", size: 684, mode: os.FileMode(420), modTime: time.Unix(1513155567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x7f\x6f\xdb\x38\x12\xfd\x7b\xf3\x29\x66\xb5\xc0\xc9\x6e\x44\xc9\xda\x5c\x7c\x4e\x6c\x05\x48\xb7\xc1\x5e\x80\xb6\x09\xae\x29\xee\x0e\x41\x80\xd2\xd2\x48\x62\x4b\x91\x2a\x39\xb2\x6b\x38\xfa\xee\x07\xca\x3f\xe2\xa4\x69\x2e\x8b\xa0\xff\x88\xa2\x66\xe6\xbd\x37\xc3\x21\x4d\x4f\x7e\x7d\x73\xf1\xc7\xd5\x7f\x2f\xcf\xa0\xa4\x4a\x9e\xec\x4d\x36\x03\xf2\xec\x64\x0f\x00\x60\x42\x82\x24\x9e\x9c\x2b\x42\xa3\xb8\x04\x8b\x66\x86\x06\xd0\x18\x6d\x80\xc1\x45\x9e\x8b\x14\x21\xd7\x06\xde\x73\x12\xda\xb9\x7c\x20\x4e\xc2\x92\x48\xed\x24\x5a\x45\xaf\x90\x2a\x24\x0e\x25\x51\xcd\xf0\x6b\x23\x66\x89\xf7\x1f\xf6\xf1\x94\xfd\xa1\xab\x9a\x93\x98\x4a\xf4\x20\xd5\x8a\x50\x51\xe2\x9d\x9f\x25\x98\x15\x18\xa4\xa5\xd1\x15\x26\xb1\x07\xd1\x2e\x48\x5a\x72\x63\x91\x12\xaf\xa1\x9c\x8d\xbc\xfb\xb6\x0d\xc6\x5c\x64\x54\x26\x19\xce\x44\x8a\xac\x9b\x04\x42\x09\x12\x5c\x32\x9b\x72\x89\x49\x1c\x0e\x82\xc6\xa2\xe9\xa6\x7c\x2a\x3b\x1e\xc5\x2b\x4c\xbc\x99\xc0\x79\xad\x0d\x79\xbb\xc8\x2b\x53\xae\x4d\xc5\x89\x65\x48\x98\xba\x84\x77\x54\x13\x4a\xac\x4b\xad\x30\x51\xfa\x91\x48\x2a\xb1\x42\x96\x6a\xa9\xcd\x4e\xd0\x6f\x87\xa3\xc3\xa3\xc3\xd7\x8f\xf8\xf3\xba\x96\xc8\x2a\x3d\x15\x12\xd9\x1c\xa7\x8c\xd7\x35\xb3\xc4\xa9\xb1\x6c\xca\x0d\xb3\xb4\xb8\x57\xb4\x07\x48\x52\xa8\x2f\x60\x50\x26\x9e\x2d\xb5\xa1\xb4\x21\x10\xa9\xd3\x5b\x1a\xcc\x13\xcf\x2d\x84\x3d\x8e\xa2\x34\x53\xa1\x56\x36\x2c\xf4\x2c\x6c\xbe\x44\xdc\x5a\x24\x1b\x89\x8a\x17\x68\xa3\x9c\xcf\x5c\x4c\x28\x52\xed\x01\x2d\x6a\x4c\xbc\xce\x12\x7d\x63\x1d\xd6\xc9\xde\xde\x2f\xbf\x2c\x97\x9f\x26\xbf\x32\x76\x2d\x72\x90\x84\x70\x7e\x06\xa3\x9b\x93\x4f\x70\x0b\x96\xe7\xf8\xcf\xab\x77\x6f\xdb\xb6\x53\x74\x5f\x93\x13\x6f\x4b\x44\xda\x08\x5a\x2e\xc3\x4b\x4e\xae\xcd\xde\x8a\xa9\xe1\x66\x71\xda\x49\xb9\xe4\x54\xb6\x6d\x94\x5a\x1b\x69\x99\x31\x81\x61\x6a\xad\x77\xb2\xe1\xbd\x46\x95\x89\xfc\x86\xb1\xc7\x18\x77\x94\x9d\x9f\xc1\xd1\xcf\x51\x25\x90\x1d\xad\x35\x6d\x39\x7f\xa8\xea\x5e\xb5\x0a\x5a\xcb\x72\x1f\x7e\x8a\xb6\x8a\x0b\xf5\x50\x1b\x63\x4f\xe8\x9b\x44\xab\xad\x3f\x99\xea\x6c\x01\xa9\xe4\xd6\x3a\x1e\x08\xaf\x16\x35\x42\xdb\xba\x15\x5f\x21\x01\x61\x55\x4b\x4e\x08\x5e\xcd\x8d\xdb\x56\xb6\x8b\x45\xe3\x41\x08\x6d\xbb\xb7\xee\x67\x2e\x14\x88\x2c\xf1\xdc\x8b\x07\x46\x4b\xdc\xbc\x13\x9f\x0a\x95\xe1\xb7\xc4\x63\xf1\x5a\xe0\x0a\x7a\x21\x50\x66\xeb\x12\x4c\xba\x1c\x9e\x64\xcd\xb5\x26\xc7\xba\xe5\xb4\xa9\x11\x35\x6d\x10\x7b\x79\xa3\xba\x9d\xda\x13\x81\x0d\x74\x50\x04\x26\xe0\x41\xd5\x5f\x8a\x6b\xff\x4f\xad\x0b\x89\xa7\x8a\xcb\x85\x3b\xaf\x2e\xa6\x9f\x31\x25\xff\x26\x31\x63\x71\x6d\x6e\x12\xf7\xb8\xbd\xdd\xc6\xf7\x97\x1b\x48\x67\x08\xbf\x26\xab\xe1\xf6\xf6\xfa\xa6\x1f\xd6\x8d\x2d\x7b\xdc\x14\x4d\x85\x8a\x6c\xbf\x0d\x3a\xa3\x4c\xe2\x57\x0a\xe7\xf0\x86\x13\xf6\xfa\x63\x9e\xd8\x30\x35\xc8\x09\xcf\x24\x3a\xc7\x9e\xee\x07\x6b\xd0\x2a\xb1\x61\x81\xb4\x36\xd8\xd7\x8b\x2b\x5e\xbc\xe7\x15\xf6\x74\xff\x7a\x70\x33\xe6\x21\xb7\x0b\x95\x26\xf1\x98\x87\xd6\xa4\x49\x31\xae\xc2\x9a\x1b\x54\xf4\x5e\x67\x18\x0a\x65\xd1\xd0\x6b\xcc\xb5\xc1\x9e\x4b\x6f\x8d\xda\xf6\x7b\x73\xa1\x32\x3d\x0f\x32\x9d\x76\xda\x02\x7f\x55\x1f\x3f\xf0\xa3\x68\x3e\x9f\x87\x45\x57\x04\xc6\x37\x55\x08\x53\x5d\x45\x77\xb3\xcf\xd6\x0f\xfc\x82\xfb\xfd\xf1\x1a\xb2\xe0\x3d\x7f\x95\x84\x1f\x80\xff\xf1\x94\x1d\x0e\x47\x47\xbf\x0f\x0e\xfe\xc1\x0e\xfc\x00\x96\x3e\x97\x52\xcf\x4f\x55\x5a\x6a\xe3\x1f\x03\x99\x06\xdb\x7b\xb1\x16\x55\xe6\x22\x6b\x5e\xa0\x3b\x68\xbb\x20\x37\xf1\x8f\x41\xea\xb4\xfb\x19\x09\x6b\x4e\xa5\x3b\x06\x61\x1f\x0a\xa4\x0f\xc8\x4d\x5a\xf6\xfa\xb0\x7f\xe7\x51\x72\x5b\xde\x01\x6f\x16\x69\xd7\x7b\xb9\xed\x29\xd7\x3a\xd1\x2b\xb8\xba\x78\x73\x01\x0c\xfe\x5d\xa2\x02\xdb\x39\x81\xb0\x50\xe9\x19\x66\x40\x1a\x0c\xaa\x0c\x0d\x1a\x98\xa3\x2f\x25\x28\x5c\x7d\xe6\x59\xb6\xf1\x26\x34\x15\x08\x45\x1a\x9c\x5e\xf8\xf3\x14\x0c\xda\x5a\x2b\x8b\x3b\x54\x51\x04\x22\xef\x7d\x9f\x49\x92\x24\xe0\x47\x2b\x24\xff\x9e\xb8\x28\xea\x86\x19\x37\xa0\x9a\x6a\x8a\xe6\x22\xff\x17\xda\x46\x92\x85\x04\x96\xcb\xdf\x44\xee\x78\x1a\x49\xe1\x03\x73\xdb\xc2\x72\xf9\x84\x09\xa5\xc5\xb6\x85\x81\x4b\x5f\xe4\x6d\x3b\xfe\x9e\xd4\x20\x35\x46\xdd\x95\x75\x9d\xe9\x3e\xf8\x7f\x7b\x80\x98\xf8\xb0\xff\x50\xdf\x03\xc0\x16\x1c\x23\x3c\x92\xdb\xe3\x34\xdf\x85\xbf\x8a\xb6\x47\xdf\x5d\xd4\xa6\x95\xf7\xf6\xb6\x9e\xef\x90\xdb\xc6\x20\x90\xa8\x10\xb4\x5a\x2d\x07\x83\xd4\x60\x26\xc8\x2d\x9a\xfb\x79\x3b\x8e\xa2\x12\x65\x1d\x6e\x7b\xd9\xdd\x29\xba\xee\x5e\xb5\xfc\xf6\x7b\x54\x75\x70\x42\x15\xcc\x01\x32\xad\xd8\x54\x37\x2a\x45\xe6\x70\xa3\x87\x2d\xe6\x7c\x4c\x1c\xf7\xfa\xcb\xdd\x76\xc6\x19\x2a\x72\x2f\x57\xa2\xc2\x0b\x75\xe9\x1a\x3a\x00\x3f\xee\x1e\x31\x3b\x18\x80\xc5\x54\xab\xcc\xba\x7e\x07\x5f\x69\xd5\x5d\xaa\x78\x87\xe9\x1f\x43\x0c\x6d\x7f\xdc\x3e\xca\x75\xf0\x4c\xae\xdf\xdd\xe3\x20\x66\xc3\x17\x70\x0d\x9f\xc9\xe5\xf6\xba\x3f\x8c\x59\x3c\x7a\x01\x59\x3c\x7a\x26\xdb\xdf\xbb\x2a\x8e\x5c\x6a\x2f\xc9\x6d\xf0\x4c\xba\xc3\x2e\xb9\x41\x97\xdd\x8b\xd2\x7b\x2e\xe1\x70\x95\xdf\x20\xde\xff\x2b\x6c\xcf\x00\x1e\x74\x0f\x16\x3f\x2b\x8b\x35\xac\x45\x72\x20\xba\xa1\xde\xba\xd5\x83\x38\x1e\x0c\x06\x3f\x74\x38\x88\x83\x83\x27\x1d\x86\x71\x30\x7c\xd2\x21\x1e\xc5\x41\x3c\x7a\x1a\x63\x10\x07\xc3\xc1\xff\x41\x19\x38\x98\x5d\xa7\x49\xb4\xb9\x0d\x4c\x22\x77\xa7\x71\xe3\xea\x4f\xce\xff\x02\x00\x00\xff\xff\x2b\xf0\xc6\xc8\xfc\x0c\x00\x00")

func templatesMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMainTmpl,
		"templates/main.tmpl",
	)
}

func templatesMainTmpl() (*asset, error) {
	bytes, err := templatesMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/main.tmpl", size: 3324, mode: os.FileMode(420), modTime: time.Unix(1513155567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\xc1\x6e\xe3\x36\x10\xbd\xfb\x2b\x08\x9e\x2b\x73\xb1\x01\x7a\x28\x14\x01\x41\xd0\x6c\x17\x0d\xe2\x43\xb6\xe8\xd1\x18\x93\x23\x69\x6a\x6a\x28\x90\x23\x29\xee\xd7\x17\xb4\xe2\x38\xad\xbd\xe8\xc6\xd9\xf6\x62\xcb\x33\x7c\x6f\x9e\x39\x8f\x43\x95\x75\x08\x82\x51\x59\x0f\x29\x5d\xeb\x3e\x12\x4b\x51\xb4\xe4\x50\x57\x8b\xb2\xfd\x78\x48\x8c\x94\x06\xf0\x7e\xd7\x92\x73\xc8\xba\xba\x9b\x61\x9e\x78\x9b\x4a\xd3\x7e\xac\x16\xa5\xa3\xf1\xb0\x7a\x26\xd5\x7f\x0f\x4e\x11\xfa\x7e\x8e\x32\x8c\xe7\x00\x05\xc3\xa8\x6c\xf0\x45\x5e\xfa\x0f\xb4\x0d\x7e\x9f\x2a\x7c\x53\x04\xc6\x42\x5a\x8a\x6e\x8e\x74\xee\x18\xd9\xab\xbe\x3a\xa5\x5d\xaf\x5b\x04\x47\xdc\xe8\xea\x17\xf4\x7d\x69\xda\xab\x6a\x51\x0e\xfe\xdc\x4a\x4f\x49\x32\x8f\xa7\x73\x59\x12\xec\x72\x16\x54\x1b\xb1\xbe\xd6\xa6\x45\xdf\x1b\xb0\x16\x53\xa2\x0d\x79\x92\x9d\xae\x6e\x5e\xff\x2c\x0d\x54\x8b\xd2\x78\x7a\x23\xa7\x0d\x61\x4b\x98\x80\x5d\x1f\x69\x04\xbb\xd3\xd5\xed\x1c\x52\xc0\x4e\x3d\x07\x2f\x65\x17\x8c\x5d\xe6\xb6\x81\x1d\x09\x05\x4e\xba\xfa\x92\x63\x7b\xf2\x63\xf4\x35\xbf\x19\x7c\xfe\x74\x34\xfe\x57\xcd\xb9\xd9\x84\x41\xd4\xea\xe1\xf1\x7b\x76\x08\x32\xe9\x90\xcc\xd4\x82\x4c\xe8\x82\xae\x7e\x6f\x41\xd4\x84\xca\x85\x0b\xb6\xef\x40\x67\x21\x22\xc6\xa4\xab\xdb\xf9\xe1\x3d\x54\x81\x05\xac\x0c\x99\x6c\x7e\x54\xc3\x25\x7c\x8c\x53\xd2\xd5\x03\x4e\xef\x11\x23\x11\x38\xf5\x10\x91\xed\x0e\xd8\x35\x61\xc4\xc8\xc0\x16\x4d\x1d\x11\x5d\xe8\x42\x4d\x5c\x87\xd8\x41\x36\x48\x1d\x48\x57\x77\x73\x42\x85\x5a\x7d\x3e\xa6\xfe\x5f\xeb\xdc\x06\x66\xb4\xa2\x26\x92\x76\xbf\x7b\xdf\xcd\x40\xad\x48\x9f\x7e\x32\x46\x26\x12\xc1\xb8\xb4\xa1\x33\xab\x87\x47\x7d\xc0\x92\x0d\xfc\x3c\x34\x95\x40\x6c\x50\xae\xf5\x7a\xe3\x81\xb7\xba\xfa\x32\x63\xde\xde\x8e\x43\xd1\x69\x9a\x96\x35\x58\xdc\x84\xb0\x7d\x53\xe5\xbb\x67\xd0\xfb\x4a\xe7\x09\x8f\x8e\x78\x5f\xda\x86\xae\x07\xde\x99\x50\xd7\x64\xb1\xa8\x43\x26\xc8\x9d\x06\x5f\x24\x01\xa1\x24\x64\xd3\x37\xa9\xbb\xdf\xf3\x7e\xe6\xcb\xd5\xf5\xc3\xc6\x93\x5d\x36\x61\x74\xe8\x69\xc4\xb8\xdb\x6b\x04\x6b\xc3\xc0\x92\xcc\x6f\xbf\xae\x1e\x1e\x4d\x1a\x36\xc9\x46\xda\x60\x4c\xf9\x74\x7c\x93\xb6\x9f\x3b\x20\xaf\xc0\x63\x94\xaf\x8f\xbf\xc3\xd7\x7c\x97\x9d\x1a\xfb\xd5\x75\x77\x7a\xd1\x79\xb2\xc8\x69\x7f\xc7\x52\xd7\x9c\x4f\xae\xd7\xd4\x35\x5a\x81\x97\x6b\xbd\xfa\x74\xaf\xd5\x44\x4e\xda\x6b\xfd\xe3\x07\xad\x52\xb4\xc7\x8d\xb0\x8e\x97\x81\x53\xde\x89\xe5\xb0\x35\x90\x12\x4a\x32\xd4\x41\x83\xc9\x84\xc6\x2f\xfb\x7c\x3c\x16\x65\xff\xb5\x3a\x82\x4f\xa2\x3a\x88\x0d\x71\xe1\xb1\x96\x22\x75\x45\xf1\x41\x57\x8b\x1b\x9f\x0f\x27\x0b\xb2\x28\x4a\x0a\x46\x20\x0f\x1b\x8f\x6a\x60\x87\x51\x49\x8b\xaa\x84\x73\x3b\x7a\xec\xd4\xb3\x8d\x0e\x36\x81\x68\x5b\x1a\xf1\x45\xac\x0b\xd6\x84\x1e\xb9\x98\xa7\x4c\x87\x2c\xb3\x2e\x8b\x66\xc4\x98\x28\xb0\xb9\x32\xa7\x2d\x5a\xf5\xc8\xea\xd3\x0b\x46\xdd\xcf\x18\x35\x5e\x2d\x3f\xe4\x96\xa9\x32\xf5\xc0\xaf\xa5\xa9\xbd\x3e\x7c\x92\x3c\xcc\xf2\xbc\xa1\xa6\xcd\xff\x14\xbc\xd7\x55\x69\xf2\xf2\xea\x07\x85\x4f\x16\x7b\x51\x53\x8b\x11\x55\x90\x16\xe3\x44\x09\x55\x36\x37\xba\x45\x69\xfa\x53\x03\x1c\x3b\x4f\x2e\xbf\x28\xe1\xd4\x87\x98\x99\x5f\xbc\xf6\x47\x2a\x8e\x61\xfa\x13\x73\xbd\x8c\x3a\x05\x75\xee\x02\x90\x6f\xfe\x1d\x64\xe6\x9e\x57\x8b\xbf\x02\x00\x00\xff\xff\x00\x2c\x01\x11\xf1\x09\x00\x00")

func templatesPartialsFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsFooterTmpl,
		"templates/partials/footer.tmpl",
	)
}

func templatesPartialsFooterTmpl() (*asset, error) {
	bytes, err := templatesPartialsFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/footer.tmpl", size: 2545, mode: os.FileMode(420), modTime: time.Unix(1513155567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\x4d\x6f\xdb\x46\x13\xbe\xe7\x57\xcc\xbb\xef\x25\x3d\xac\x58\xf4\xd2\x26\xa0\x08\xb4\xae\x81\x18\x48\xd2\xc2\x6e\x0e\x45\x60\x38\x63\xee\x90\xdc\x6a\xb9\xcb\xee\xae\x64\x0b\x2a\xff\x7b\xb1\xe2\xa7\x25\x52\xb6\xd1\xa0\x4d\x74\x22\x87\xf3\xf1\xcc\xcc\xb3\xb3\xa3\xb8\x20\x14\x64\x93\x17\x00\x00\x31\x42\xaa\xd0\xb9\x25\x73\x2b\x59\x29\xa9\x57\x0c\x0a\x4b\xd9\x92\xfd\xbf\x44\xa9\x19\x78\xbc\x95\x5a\xd0\xfd\x92\x7d\xcb\x92\xab\x95\xac\xc0\x1b\x08\x9f\x20\x35\xda\x93\xf6\x71\x84\xad\x2b\x21\x37\x20\xc5\x92\x55\x98\xd3\xaf\xe8\x0b\xd6\xb9\x2e\xa4\x20\x96\xec\x76\x8b\x0f\x97\x17\x75\x1d\x47\x42\x6e\x92\x17\x83\x4d\xab\x75\x67\xb1\xaa\xc8\xb2\xc6\xdb\xe1\xd7\x06\x34\xa4\x46\xf1\xa0\x38\xd2\x3a\xd4\x4c\x8d\xda\xab\x71\x95\x73\xa3\x89\xfb\x42\x5a\xd1\x48\x4a\x31\x48\x0e\x3c\xb4\xc5\x08\xf8\x95\xc9\x0d\x1f\x97\x22\x9a\xd0\x0d\xbf\xdd\xee\x53\xfc\x3f\xce\x3f\xca\x0c\x94\x27\xb8\x38\x87\x1f\xae\x93\x4f\xf0\x17\x38\xcc\xe8\xcd\x6f\xef\xde\xd6\xf5\xa4\xdd\x3e\x96\x2c\xf3\x0e\x71\x08\xc8\xc0\xd9\x74\xc9\x0a\xef\x2b\xf7\x3a\x8a\x52\xa1\x17\x46\xbb\x45\x6e\x36\x8b\xf5\x2a\x42\xe7\xc8\xbb\x48\x96\x98\x93\x8b\x8c\x76\x3c\xd8\x44\x9b\xef\xfa\xe7\x45\xa5\x73\x06\xa8\xfc\x92\xfd\x92\x65\x32\x25\xc8\x8c\x85\xf7\xe8\xa5\xd1\xa8\xe0\xca\xa3\x97\xce\xcb\xd4\x9d\xca\xe5\x23\x69\x21\xb3\x6b\xce\x9f\x92\xc5\x28\xfb\xbc\xc9\xfe\xd5\x75\x12\x24\xff\x59\x09\xdc\xe6\x5f\x2c\x41\xcf\xfb\x41\xb0\xe7\xf5\x03\xd1\x0c\x2d\xfd\x9d\x69\x48\xe8\x7a\x5e\x8e\x44\xe1\xb8\x70\xee\x4a\xa8\xac\xd4\x9e\xf3\xe6\xf8\x1c\x03\x10\xaa\xa1\x2b\xea\x7c\x8d\x39\xf5\xc7\xad\x17\x4c\xa7\x19\x0b\x7f\xa8\x79\x73\xe3\xa5\x57\xc4\x92\xb3\x02\x75\x4e\xd0\xc9\x5f\xc7\x91\xf0\x73\x5e\xc4\xb1\x17\xe9\xa9\x9c\x89\xda\x14\x58\x66\x40\x7f\xc2\xe2\x6d\x6b\x01\x8c\x34\xab\xeb\x61\x08\x0d\xae\xc2\xf9\x03\x99\x1a\xdd\xe6\xdf\x1d\xc5\x28\xdd\x2e\x76\xbb\xc5\x95\xf4\xf4\xb3\x09\x73\xa8\xae\xbb\xc1\xc2\x92\xb3\x6d\x69\x91\x72\x78\x79\xf6\xfb\x37\xa1\x3f\xbb\x1d\x69\x71\x82\x81\x13\x80\xd2\xed\x73\x01\xcd\xa2\x39\xd7\xb9\x92\xae\x80\x97\xe7\xef\x1f\x45\x13\x47\x42\x4c\xb4\x38\x12\xea\x59\x24\x73\x94\x1a\x2d\xd0\x6e\xb9\xc6\x0d\x3c\x9d\x72\x8f\x30\x6d\xdd\x30\x4d\xe3\x86\xf7\x11\xd8\x64\xcc\x50\x27\xe7\xe1\x0f\x17\x5e\x78\xaa\x8c\xa6\x46\x34\x47\x46\x25\x67\xfc\x3c\xc2\xa5\xd1\xc5\x75\x18\x5f\xaf\xe0\x08\xc0\x68\x98\x5b\x52\x84\x8e\x52\x54\x14\xcc\x58\x72\xd9\x08\xa0\x93\x1c\x9d\xec\xa1\xf8\x4a\x7e\x31\x69\x94\xe4\x0b\x23\x8c\x32\xf9\x96\x25\xef\x86\x97\xaf\x03\xbd\xa6\x3b\x17\x60\x0b\x89\x5f\x07\x60\xbc\x35\x6b\xbf\x76\x2c\xf9\x31\x3c\x7c\x39\x98\x87\x89\x34\x89\xde\xa3\xcd\xc9\x2f\xd9\xcd\xad\xc2\xf0\x6e\x49\x2d\x99\x36\xa6\x22\x4d\x16\xb4\xb1\x94\x91\xb5\x64\xbb\x3c\xbb\xcb\xf7\x56\x99\x7c\x7c\xfb\xb2\xe4\x27\x65\xf2\x67\xa6\x1d\x47\xeb\xd3\xe3\x6b\xf4\x3a\x7e\x1c\x0d\xb4\xca\xca\xb2\x1b\x67\xd3\x43\x2a\xd6\x78\x38\x10\xd7\xaa\x33\x0f\xf5\xe0\x61\x51\xb5\x46\x4d\x5d\xfb\xa3\xa6\x3c\x50\x3d\xd9\x94\x18\xbb\x15\x39\xd8\xb4\x08\xd9\x7e\x42\x96\xa4\xd7\xdc\x9b\x3c\x57\xc4\x00\xad\xc4\xde\x63\x13\xa0\x57\x9e\x0e\x1a\xcc\x4f\x31\xc1\x55\xa8\x67\x4c\x3d\xdd\xfb\x70\xa0\xf4\x3a\x8e\x82\xda\x5c\x9f\x26\x1a\x38\xd3\xbc\x93\xa5\x81\x27\xd5\xc6\x11\xda\xb4\x68\x4a\xd3\x3c\x9f\x2a\x4e\xa7\x3d\x1d\xb5\xfd\xfa\x4f\xaa\x73\xb5\x77\xf1\x19\xea\x33\x41\xec\xe3\x3f\x32\x10\x30\x84\x9d\x60\x4f\x59\x41\x9a\xf5\x97\x68\x4f\x83\x7d\x15\xe8\xbe\x42\x2d\x48\x2c\x59\x86\xca\xcd\x5d\xc0\xc7\x07\xe2\xa9\x17\xeb\x03\x93\x7d\xf3\xfa\xbd\xe7\xc3\xe5\x05\xb0\x88\xd5\xf5\x91\x0e\xe7\x98\x7a\xb9\xa1\x76\x69\x69\x87\xcb\xb0\x9b\x1a\x25\xb8\x24\xce\x85\x74\x95\xc2\x2d\xbf\x55\x26\x5d\xcd\x76\x67\x98\x61\x07\xf8\xf5\x6a\xd8\x52\x4a\xc1\xbf\xef\xf7\x95\x57\xa3\x3f\x5f\x6f\x4c\x49\x9f\x7f\xf4\xf4\x73\xa3\xfd\x12\x47\xdd\x1f\xe3\xbf\x03\x00\x00\xff\xff\x92\xea\xe8\xf4\x22\x0f\x00\x00")

func templatesPartialsHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsHeaderTmpl,
		"templates/partials/header.tmpl",
	)
}

func templatesPartialsHeaderTmpl() (*asset, error) {
	bytes, err := templatesPartialsHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/header.tmpl", size: 3874, mode: os.FileMode(420), modTime: time.Unix(1513155567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _redirectsRedirectsCsv = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5c\xcb\x96\xdb\x36\x93\xde\xeb\x59\x9a\xa6\xed\x24\x33\xe7\xcc\xd6\x67\x16\xd9\x24\x0b\x3f\x40\x1f\x08\x28\x92\x65\x81\x28\x36\x2e\x6c\x6b\x9e\x7e\x0e\x00\xde\x45\x52\xa4\xc4\x76\xf2\x6f\xe2\x16\x2e\xf5\x7d\x05\xa0\x0a\x85\x02\x98\x14\x38\x29\x2a\xaf\x69\xae\xc9\x18\x41\x25\x18\x8b\xbc\xd2\x24\x1c\xb7\xb9\xa8\xd2\xb7\x12\xcd\x52\xe5\x5b\x89\x2f\x77\x05\x94\x60\x0b\x12\x24\x29\x47\x58\x93\x74\x4a\x2b\xa0\x4a\x42\x45\x95\x93\xcc\x22\x29\xa6\x04\xa7\xb2\x74\x0a\xed\x35\x2d\xc8\x19\x54\x79\xe4\xe3\x7f\x40\xa5\x91\x83\xb1\xcc\xa2\x97\x65\x32\xd2\xa6\x64\x52\x32\x0d\xcc\x04\x66\x9b\xe4\x8d\xe9\x6d\x11\x7c\xea\x54\x56\x51\xac\x64\x9c\x93\x53\xd6\xa4\xee\x62\x80\x5b\xd2\x5d\x41\x60\x8b\x4a\xc0\x4f\xca\x0c\xe8\xda\x0b\x1e\x8d\xda\x7d\x11\x63\x82\xb7\xb2\x4e\x29\x94\x95\xa4\x6b\x09\xca\x32\x25\x24\x3b\x93\xd3\x25\xd3\x17\xb0\xcd\x08\xa0\x7a\x27\x7d\x49\x63\x4d\x33\xe6\x58\xfb\x41\x08\xf4\x6e\xcb\x23\xc3\x87\xa5\x8e\x19\xcf\x8a\xef\xc7\x30\xa7\x1a\xb4\xf2\x30\x95\x3b\x4b\xe4\x8d\xf6\x4a\x58\xf6\x13\x4c\x3a\x2c\xcc\x50\x31\xc5\x21\xb2\x9e\xa9\x30\x4c\x89\x5e\x9c\x80\x0c\x39\x7a\xf2\x02\xce\xd6\x29\x01\xda\x16\x50\x32\x66\xac\x46\x5e\x58\xab\x81\xb5\xaa\x3e\xce\x65\xac\xeb\x71\xa4\xfa\x01\x42\x95\x75\xcb\x37\xac\x4c\x54\xc2\x4f\x7d\x33\x0c\x61\x5c\x41\x0f\x6b\x46\x3a\x2d\x76\x9f\x30\x9f\x97\x73\x4a\xcf\xde\x4c\xc0\x18\x54\xc2\x19\xab\xaf\x7e\x34\x34\x13\x90\x6a\xb0\x0c\x65\x5b\x1c\xd9\xc4\x32\xc3\x24\x98\xb0\x50\x03\x93\xad\x12\xc6\x84\xa6\xa2\xb4\xc1\x1d\xc3\xd2\xdb\x71\xe8\x5c\x54\xf8\xc8\xa0\xb8\xcb\xac\x9c\x47\x1d\x40\x6b\x03\xa4\x9e\x72\x01\xee\x32\x27\x6f\x85\xd4\x99\x49\xbf\x12\x29\xab\x58\xb0\xe7\x86\x95\xbb\x84\x69\x58\xe7\x72\xdb\x77\x4a\xa6\x15\x72\xc7\x7f\x8f\x8a\x4b\xcc\x75\xf8\x7b\x50\xec\x77\x84\x92\xd9\x76\x06\xbf\x7e\xfe\xf2\x85\x83\x32\xce\x8c\x3c\x31\xa8\x5c\x32\x25\x98\x12\xef\x7e\x69\x94\x4c\xf3\xc2\x37\xbd\xef\xf0\x77\x10\x18\xab\xb8\x9f\xc9\xbd\xad\x0c\x98\xb4\x05\x53\xc2\x10\x47\x26\x39\xd3\x90\x72\xe6\x0c\x18\xca\x04\x30\x5b\xc4\x11\x60\x92\x53\x41\x52\x83\x64\x16\x44\xa8\x30\xa8\x6c\x01\xee\xb2\x61\x7b\xbb\x8b\x31\x56\x72\x1d\x6c\xc5\x0d\xb4\x15\xdd\x1f\xed\xd6\xd4\x28\xa1\x94\x63\xb2\xab\x73\xba\x86\xeb\xba\x57\x58\x16\x38\x61\x3c\x2f\xf9\xe0\x65\x18\x51\x4a\x14\x57\x60\x7a\xa6\xd5\x07\xae\xbb\x0d\xd0\x9b\xf7\xfe\x51\xb3\xf8\x03\xc0\x5e\xab\xb1\x9a\x3d\x48\x1c\x4e\x56\x99\x3d\xa1\xc0\x0a\xc8\x9c\x66\xf3\x68\x07\xd9\x4e\x4d\x28\xd8\x59\x42\x49\xda\x32\x89\xf6\x8a\x6a\x62\xb0\x1f\x60\x43\x9b\x40\xef\x28\x78\x46\x6d\x0b\x13\x0d\xd0\x2f\x18\xa6\x35\xb2\x1c\x4c\x1a\x8b\xa2\x7a\xa1\x91\x37\x58\x6d\xad\x07\xf4\xfe\x08\x55\xc6\x7c\x80\xe1\x5b\xdd\xd7\xed\x1e\xcc\x58\xb3\x0d\x78\x5b\x5c\x04\x6b\x02\x40\x83\xff\x07\x7e\x25\x11\x8f\xb6\x10\x75\x6a\x5a\x09\x28\x29\xd7\xac\x2a\xb6\xfa\x89\x25\xa9\x13\x15\xe6\xc4\x0f\xc3\x89\x1a\x8c\x0d\x3b\x5c\x05\xca\xf8\x45\x19\x90\x9c\x69\xb7\xcb\x9e\x48\xdb\x72\x12\x52\xac\x08\x98\xa7\x32\x92\x74\x98\x29\xb7\xd2\x35\xe4\x68\x2c\xe8\xbe\x7d\x34\xb3\xb3\x86\x8f\xb1\xea\x8d\xc0\x8f\x18\x38\x29\x81\xb6\x19\x52\x81\x06\x98\x69\xb5\xe5\x3e\x3a\xd1\x11\x32\x7a\xd6\x7e\x93\x7e\xd0\xc2\x67\xc1\xc6\xaa\xde\x41\xfd\x08\x15\xfd\x18\x62\xcd\x64\x07\xe4\xf7\x3e\x29\xc1\xa2\xfa\x70\x45\xef\x60\xef\x8e\x87\x39\xab\xd0\x7a\x71\xc4\x2f\x1e\xbc\xf9\xcd\xfd\x4e\x50\x56\x4f\x07\xc8\x9b\xc4\x1f\xe1\x86\x79\x81\x52\x74\xce\x7e\xcf\xc2\xdb\xe7\x7e\x57\x70\x36\xbb\x8d\xe1\xd1\xb4\x02\x6d\x48\x29\x90\x8d\x1a\x58\xa3\x6c\xa3\xac\xb1\x12\x4f\x09\x9f\x28\xb1\x84\x72\x98\xe7\xe3\x92\x61\xc9\x94\x0d\x8b\xe2\x43\x5c\xdc\x14\xe1\xd1\x45\xc4\x49\x71\xa8\x9a\xf6\x19\x68\x8b\x7e\x62\x75\x1f\x76\xf6\x0d\x0e\x59\x55\x6b\x78\x13\x0d\xe7\x81\x57\x4c\xdc\x30\x0b\x52\xa2\x85\x89\x8d\x07\x6b\x03\x6d\x35\x28\x71\x27\xe7\x75\x2b\xe2\x86\xd4\x58\xd6\x11\xc6\xdb\x9e\x71\x6a\x92\x35\xaa\x9c\x4b\x32\x56\xa3\x40\x57\x0a\xcc\x32\xe4\x28\x61\x63\x9c\xb8\xcf\x98\x77\xe0\x7e\x80\x9a\xa5\x36\xec\x57\xa8\x35\xc4\x39\x4e\x8d\xe6\x78\x6a\x49\x68\x97\x57\x84\x86\x14\xaa\xfc\x81\xd0\xfe\x11\xdd\xb6\x83\x3f\xac\x30\xd6\xa4\xdb\x04\x67\xf3\xc3\x1c\xaa\x5d\x03\x30\x51\x6f\x05\x69\x25\xac\x0f\x69\x4c\x54\x16\xb4\xf2\x91\xc0\x30\x0f\x08\x9e\x0b\x68\x0e\xeb\x21\xfc\xbc\x84\x31\xb9\xa1\xa8\xde\x0b\x81\xaa\x51\x53\x48\xa9\x0e\xfc\x48\xc4\x9e\xab\x22\xc5\x50\x43\x89\x26\x04\xe6\x23\x6f\x34\x2f\x6a\x42\x62\x93\xcc\x67\xe8\x81\x02\x9d\x5f\x8f\x22\xd6\x49\x7b\x8e\xd2\xa0\x38\x64\xc0\x8f\xa3\x37\x23\xf9\x08\x2f\x01\x3f\x39\x18\xf3\x1e\x96\xd4\x33\x27\xff\x7d\xee\x61\x33\xea\x9e\x34\x1a\x2a\x45\xf5\xe0\x78\x9c\x91\x06\xcc\x95\x40\x0d\xdc\xf6\x27\xc7\x4c\xe0\xce\x74\xda\x40\xf0\x58\x8d\x55\x84\x63\xd2\x32\x79\x8c\xfc\x48\x99\x0a\x38\xfa\xcd\x2e\x24\x11\xba\x41\x3b\x3e\x27\x73\x0f\x71\x7e\x9a\xb6\xc6\x8d\x4c\xfb\x2d\xc0\x2f\x11\xff\x1b\x55\x5e\x90\xd3\xe3\xa4\x7f\xec\xcd\xc9\x58\x53\x81\xf6\xf5\xbb\x22\xd3\x05\x84\xd9\x4b\xc1\x59\xa8\xad\x37\x26\xdd\x9d\x02\xd6\xcc\x42\x73\x3b\xaa\x83\x85\xc6\xfb\xd0\x47\x2e\x50\xb6\x08\xdd\x73\xef\xeb\x4d\x2c\xf7\x61\xab\x08\x17\x33\x05\x49\x11\x73\x0b\xb8\x25\xe7\x35\x7f\xe5\xbb\x28\xf3\xe1\x64\x72\xf7\xd7\x3b\xda\xa2\x49\xa2\xf7\xfc\xb5\x62\xb2\x6b\xf1\x7c\x1e\x79\x16\xec\x56\xc3\x45\xd4\xc1\x0e\x11\xfe\x45\x4e\xce\x56\xce\x86\x89\x1d\x5c\xea\x0e\x7f\x94\xc0\x8c\xd3\xfd\xca\x09\xe2\x63\x4c\xcf\xa9\xac\x98\xf6\xb1\x91\xe9\x2f\xa8\xfa\x3b\xe5\xc7\xa1\x66\x55\xba\x8f\x79\x67\x06\x25\xa0\x17\xef\xbd\x25\x39\x8d\xa6\x9c\x51\xa9\x62\xc6\x80\xca\x63\x1a\x04\xae\x58\x6d\x98\xac\x5b\xb9\x2b\xfc\x27\x00\x77\xa2\xae\x61\xcf\x58\x74\x4b\x39\x94\xa3\x1a\xbd\x38\xd8\x23\x71\x85\xec\xad\xe8\x0d\x06\x1c\xac\x8a\x17\x4c\x33\x6e\x41\xc7\x53\x65\x5a\x50\x09\x6d\xf0\xd7\x39\xf4\x12\x04\x32\x67\x58\x3e\x52\x0b\x2c\xe3\x7e\x8f\xed\x2d\x34\x04\x8e\x02\x6b\x14\x8e\xc9\x8d\xb6\xff\x10\x8b\xb9\xa1\xd8\x40\x67\xf3\x1e\xe2\xff\x53\x49\xc6\x41\xa0\xa9\x9c\x85\x81\xab\xef\x13\x74\xc3\x87\x1a\x6d\xbb\xc1\x5d\xda\xa1\x50\x73\xaf\x37\x6e\x30\x0f\xcb\xd9\xc4\x9e\x59\x38\x7f\x04\xe1\x32\xfb\x98\xec\xf4\x02\xd0\xc1\x8a\xc4\x9e\x1f\xa8\x40\x07\xf0\x40\x40\x26\xd0\xb0\x73\xc8\xfc\x34\x94\x31\x03\xaa\x2a\xd2\xd6\xf7\xc1\xd1\x92\x7a\x46\xf8\x84\xf7\x22\xca\xbd\xdd\x35\x24\x12\x99\x0f\xce\x3a\x43\x6b\x1f\xd5\xa4\xa8\x38\x95\xde\xbf\xbe\x07\x2e\xad\x42\x21\xa3\xe2\x03\x20\xa6\x44\x46\x24\xb6\x6a\xb4\x0b\x6b\xaa\xdf\x22\xe8\xb3\x77\xd1\x95\xa6\x1f\xc0\x07\x1e\xa0\x64\x3a\x66\xb6\x99\x75\x66\xb6\xdd\xb3\x97\xd1\x43\xc8\xb1\x9e\xf7\xb1\xd7\x37\x2e\x5e\x30\xe5\xa3\x71\xea\x4e\x24\x25\x68\xbf\xe9\xf9\xe0\x9b\xbf\x39\x34\x43\x5f\x37\x5f\x57\xb2\xf5\xad\x6c\x33\xc6\x44\xb5\x65\xb0\xc7\x5f\x3e\x94\xa4\x6c\x21\xaf\x1f\xf1\xf4\x61\x49\xf4\x61\xce\xcc\x37\x0a\xae\xf2\x07\x9d\x3f\xc6\x1d\x4f\x11\x1e\x33\x95\x51\x64\xd2\x17\x07\x1d\x4c\xb8\xad\x06\x5d\xce\x37\xea\x62\xe0\x8c\xb4\x24\xce\x24\x73\xb6\x20\x1d\x3c\xd4\xc3\xc1\xf8\x02\x9f\xb1\xea\x3b\x88\xed\x4e\x2d\x7e\xec\x38\xed\xce\xa3\xfc\x83\xc3\xf4\xfc\xb5\xd8\xb0\xaa\x97\x76\xf4\xb5\xd8\x22\xca\x7f\x98\xe7\xd9\x1d\x30\x74\x65\x3e\x40\x80\x9f\x15\x70\xcb\x14\xc7\x56\x09\xaa\x50\x35\x77\xd3\xbe\x81\xb1\x57\x09\xcf\xc4\x27\xcb\x70\x63\x0d\x57\x71\x1f\x56\xd3\xff\x69\xae\xc6\x42\x73\xc0\x6c\x97\xf3\x19\x34\xb0\x1a\xfa\x47\x12\x35\x75\x07\xb6\xe7\xa1\xc6\x9a\xdd\xc7\x3c\x6c\xff\xa8\x19\x67\x8a\x5f\xf7\x1f\x51\x36\xef\x1f\x53\x84\xcd\x69\x8c\xc6\xdc\xe2\xea\xbe\x7d\x0c\x3f\xaa\x1f\x56\x77\xfe\x06\x84\x8b\x0f\x8c\xf6\x64\x34\x96\x51\xe7\xbd\xc1\x36\xf8\xad\xb9\xbd\x31\xf0\xf3\x8f\xc5\x17\xe4\x35\x7d\x7f\xe9\x6c\x58\xb2\x4c\x8e\xc5\xfc\x33\xf3\x32\x4f\xe4\xf0\x23\x8e\x2d\x00\xb2\x0c\xb8\x35\x94\x85\x9b\x13\xa6\xc4\x19\x14\x64\x68\x0d\xa9\x4e\x48\xec\xfb\xa1\xe7\x9e\x75\x26\xcd\xd3\xf9\x9e\xcb\x9d\x91\x00\xd9\x9e\x3a\xc2\x5f\xa4\x99\x1c\x3e\xaf\x6a\x5f\xac\x77\x95\x7b\x9e\x45\xdc\x93\x3d\x7d\xd1\x3e\x0f\xb2\xe7\xad\x60\xf9\xf6\x47\xdf\xe8\x7c\x45\x65\x9c\xf6\x63\x1a\x32\x95\x0a\xa1\xeb\x97\x39\x25\xfa\xce\x8f\x3e\x28\x7c\x1c\xae\x57\xca\x8f\xc8\xe8\x51\x46\xfc\x52\x0b\x4d\x45\x86\x9d\x25\x4c\xa6\xb3\xfd\xde\x23\x76\x5a\x6f\x3b\x52\x6b\x2f\xce\xf4\xab\x90\x8d\x80\x7b\x66\x8b\xde\x15\x68\x53\x60\x45\x99\xbb\xbc\x39\xb2\x20\x4c\xc1\xf4\x8d\x6f\xdc\x3c\x1f\x6b\x02\x9f\x7e\x27\xeb\x2e\x6d\xbb\x43\xdf\xc7\xf6\x62\x17\xda\x3f\x79\x7f\xe9\x2e\xa3\x4f\xff\x7c\x04\xa6\x04\x5a\xa7\x81\x94\x06\x03\x4c\xfb\xc3\xbb\x10\x50\x83\xa4\xaa\x8b\xb4\x0f\xb8\xd7\x7c\x08\xf9\x34\xfe\xae\xb1\x66\xd2\x01\x13\x02\x44\x5e\xb3\x99\xb5\xdf\xd7\xf7\x1f\xe5\xb0\xaa\xd2\xc4\x78\x71\xfb\x95\xe4\x58\xda\xca\x0a\xbf\x23\xf6\xb4\xed\xeb\xcd\x37\xc7\xb4\x05\x2d\xaf\xc3\xec\x46\x7c\xf3\xe9\xed\x47\x1a\xca\x9a\x27\x92\xcc\x18\xb0\xa3\xcc\xf2\xae\x8f\x3b\x1f\x06\x3a\xa5\xf1\x4f\x93\xe6\x60\x13\x63\xbd\x18\xf1\x52\x58\x5b\x99\xff\x49\xdb\xba\x4f\xa4\xcc\xa7\x9c\xea\x4f\xee\x72\x4a\x6b\x34\x8e\x49\x34\xf1\x6e\x37\xf5\x0b\x80\x93\xb2\xa0\x6c\x2a\xa0\xfe\x2d\xfd\xfb\xaf\xef\xaf\xdf\xff\xfe\xf6\x4a\x9c\xbb\x2a\x34\x7a\xe5\x24\x50\xe5\xaf\x96\x48\x7e\x2a\x6c\x29\x3b\xf9\xa4\x8c\xc0\xdc\xd3\xfa\x94\xa3\x2d\xdc\xf9\x13\x52\x2a\xaa\x84\x7b\x9a\x98\x61\x5c\xff\x89\xef\x68\x52\xe3\x83\x08\xa6\x45\xd2\x4b\x66\x72\xd2\x74\x0b\xfa\x36\x15\xfe\xfa\xfe\xfd\x7f\xbf\xbd\x0a\x34\x9c\x6a\xd0\xd7\x5f\x46\x7e\x11\xf7\xc9\x57\x1a\x13\xcb\x34\xa0\x73\x87\x02\x2c\xcd\xbc\x6e\x3d\xf4\x8d\xc6\x3a\xd4\x0f\x27\xaf\x5f\x3f\x7f\xf9\xef\xd5\x49\xf9\xfe\xe7\xb7\x38\xb1\x7f\x7e\x7b\x2d\x10\xb4\xf7\x1b\xd7\xd7\x1a\xe1\xfd\x88\x19\x69\x1c\x1c\x2e\x2c\xa6\x79\xcc\x53\xaa\x41\x36\xef\xb6\xe3\xc9\x3c\xc3\xdc\x69\x30\xa4\xda\x37\x6b\xcd\x5b\x7c\x71\xbe\x32\x0d\x8c\x32\xe7\xf5\xd3\x60\x50\x80\xe2\x30\xc9\x9c\x7c\xfd\xfc\xe5\x8f\xf8\x71\xce\xcb\xc1\x92\x7f\x38\x05\x5f\x3f\x7f\xf9\xaf\x4a\x53\x8d\x26\xac\xbb\x53\x1a\xbf\xaa\x4b\x0b\x90\x55\x6a\x41\x97\xe1\xb1\x76\x77\x13\xf5\xd2\xd6\xc7\x7f\xac\x66\xca\x64\xa4\xcb\x36\xed\x9c\x6b\x56\x96\x30\xd3\x71\x2c\x98\x13\x5d\x30\x44\xa7\xe1\x19\x02\xbf\x6e\x94\xdb\xb4\xf6\x8b\x24\x9c\x95\x47\x52\x37\xca\xf0\x6b\xaa\x69\x01\xc6\x9e\xba\xd3\x08\x27\x95\xbd\xa4\xfe\x14\x6c\x9d\x49\xdf\x0b\x66\xdf\x41\x50\x0a\x75\xf8\xc2\x72\xd8\x0a\xb4\x1f\xcc\x53\xaa\x4c\xf1\xfe\xd2\x79\x49\x54\x1d\x58\x46\x7a\xe9\x1a\xb2\xbf\xf4\x1c\x97\x37\x42\xda\x7c\x80\xb1\x4e\x5c\x29\xeb\x32\x0a\xef\x20\xe5\x19\x50\xe5\xe1\xd9\xb9\x06\x6f\x76\x57\x72\x2a\x8f\xe6\xf8\x68\xbf\xa8\x43\xa1\xff\x8d\x5a\x78\xf8\x9b\x8e\x63\x7a\x83\xbe\xf1\xc6\x07\x55\xc1\xa4\xdf\xd2\xba\x25\x7f\xf2\xa6\xef\xed\xa6\x7f\x94\xf1\xa5\xff\x33\xe9\xdd\x4d\xd2\x6d\x94\x89\x86\x8a\xb4\x4d\x33\x38\x6b\xc7\xf4\x35\xf1\x36\x92\x1a\x7b\x4e\x4a\xf3\xa6\x93\x51\x71\xf4\x32\x47\x26\x5c\xbb\x8f\x38\x7a\x96\x83\x23\x4f\xcb\x71\x42\xd1\x53\x39\x4c\xd3\xf0\x40\xe8\x5f\xa9\x98\x2c\x4d\x1a\xe6\x3c\x89\xbc\x12\x45\x36\x41\x95\x74\xf9\x97\x24\xe9\x93\x55\x09\xe9\xc4\x6a\x86\x0a\x55\x9e\x24\x0a\xc0\x9a\xe4\x76\x4e\x1b\x3d\xef\xa5\xc2\x14\xd9\x26\x1b\xe6\x54\xdf\x76\xa0\xd3\x60\x29\x86\xb6\x1d\xa5\xbe\x35\xe9\x96\x8e\x27\xb3\xa0\xa2\x62\xec\x6b\x6a\xbc\xb7\x11\x49\x9b\xc7\x48\x28\x4b\x42\xe0\xf8\xbb\xa7\xfd\x47\x58\x8a\x5f\x63\xbd\x2f\x4f\x9a\xf2\x56\x95\x3b\xe1\x61\xcf\x39\xc2\xb4\x28\x94\x05\x8c\x38\x11\xbf\x13\xb7\x96\x04\x70\x2f\xb8\x67\x77\x76\x26\x89\x67\xae\x2e\xc4\x4f\xfa\x33\x58\x4b\x30\x19\x6c\x28\x89\x06\xe3\xa4\x35\xa3\x55\xb5\x9d\xe2\xed\x67\x79\xb3\x0c\x07\x80\x0d\xde\x78\xd1\xc4\x19\x4d\xe2\x94\x0e\x4c\x61\x65\xe5\x3f\x9f\x1c\xed\xb5\x70\x97\x91\x80\xdb\x99\x67\x15\x36\x09\xaa\xf4\x1d\xce\xc3\x9f\x05\x95\xf0\x12\x37\xcd\x70\x23\xd6\x73\x8f\x3d\x87\x94\xb7\xa6\x1e\xdb\x4f\x37\x6e\x52\x84\xbb\x93\x8e\xf3\x92\x50\x71\xe9\x04\x18\x26\xe5\x6f\x4d\x3f\x5e\x61\xc1\x2b\x64\x4a\xe8\xd1\xff\xae\x01\x79\x6b\x13\x33\x5b\x6f\xb7\x67\xc7\x40\x21\xdc\x69\xf7\x3b\x71\x3f\x12\xe3\x9d\x3e\x05\x95\xb3\x1c\x55\xfe\x8e\xb6\xb8\x6d\x3c\xdd\xd0\x47\x62\x02\x93\x9e\x5b\x90\x14\x02\x8d\x5f\xc7\xee\xa6\x28\xf2\x1d\x90\x0a\xbf\xff\x05\x84\xba\xab\x98\xcd\x5c\x48\x49\x54\x30\x7b\x9b\xe3\x77\xfc\xd3\xb1\x51\xc8\xdc\xeb\xa2\x83\x03\x9d\x1b\x88\x83\x55\xb8\xf3\xe6\xf1\x60\x6d\xd6\xd0\x8e\x9e\x9b\x85\x07\x3a\x92\x67\x47\x4f\xd1\x02\xd2\xc1\x0a\xf5\xb5\x83\x2c\x4a\xc1\x8e\x9e\xa2\x59\x98\x83\x55\xc9\x58\x89\x3e\x10\x33\xe4\xfc\xba\x8e\xc5\x99\x3e\x5a\x95\x59\x98\x83\x55\x59\xbe\x38\xa6\x4a\x1d\xac\xcf\x32\xd6\xc1\x4a\x71\x8d\x65\x23\xb9\x7f\x67\x11\x0e\xf4\xf3\x55\x6d\x6d\xc6\xde\x8e\x9e\xc3\x65\x2a\xff\x1f\x00\x00\xff\xff\xc4\x66\xdd\x89\xc4\x4e\x00\x00")

func redirectsRedirectsCsvBytes() ([]byte, error) {
	return bindataRead(
		_redirectsRedirectsCsv,
		"redirects/redirects.csv",
	)
}

func redirectsRedirectsCsv() (*asset, error) {
	bytes, err := redirectsRedirectsCsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "redirects/redirects.csv", size: 20164, mode: os.FileMode(420), modTime: time.Unix(1539000048, 0)}
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
	"templates/error.tmpl": templatesErrorTmpl,
	"templates/main.tmpl": templatesMainTmpl,
	"templates/partials/footer.tmpl": templatesPartialsFooterTmpl,
	"templates/partials/header.tmpl": templatesPartialsHeaderTmpl,
	"redirects/redirects.csv": redirectsRedirectsCsv,
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
	"redirects": &bintree{nil, map[string]*bintree{
		"redirects.csv": &bintree{redirectsRedirectsCsv, map[string]*bintree{}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"error.tmpl": &bintree{templatesErrorTmpl, map[string]*bintree{}},
		"main.tmpl": &bintree{templatesMainTmpl, map[string]*bintree{}},
		"partials": &bintree{nil, map[string]*bintree{
			"footer.tmpl": &bintree{templatesPartialsFooterTmpl, map[string]*bintree{}},
			"header.tmpl": &bintree{templatesPartialsHeaderTmpl, map[string]*bintree{}},
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

