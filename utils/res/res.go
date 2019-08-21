// Code generated by go-bindata. DO NOT EDIT.
// sources:
// cert/yckit_demo.p12
package res

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

var _certYckit_demoP12 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xd5\x77\x38\xd5\x0f\xdf\x07\xf0\xb3\xec\x93\xbd\xc7\xb1\x92\xed\x1c\x11\x59\xd1\xb0\x8f\x15\xca\x0a\x65\xc4\xe1\x58\x21\x19\xe1\xeb\x38\xc7\xcc\x08\x49\x64\x44\xa5\x24\x7b\x73\x32\x4e\x8e\x4e\xf8\x25\x59\x47\x36\x71\x6c\x59\x19\x79\xae\xfb\xba\x9e\xe7\xb9\xef\xfb\xcf\xcf\xf5\x79\xbf\xff\x7d\xbf\x50\xc0\x39\x29\x08\x18\x8a\x02\xe0\x6b\xf4\x4c\x72\x78\x23\xfc\xc1\x39\x30\x03\xb8\x08\x80\xff\x84\x01\xf0\x71\x14\x00\xff\x81\x02\x18\xee\xff\xfb\x45\x5f\x04\x30\xd8\xa1\x00\x06\x1b\x08\x18\x84\x02\x18\xcc\xfe\xa3\x84\x12\xa6\x67\xfe\xdf\x03\x0e\xa6\x47\xb1\xc2\x18\x05\x33\xf9\x97\x56\x33\x08\x13\x10\x08\x23\x28\x06\x60\x10\xdb\x74\xb0\x94\x6c\x31\x71\x22\xf9\xbe\x2a\x93\xb7\x7b\x69\xf1\x38\xcb\xd9\xe3\xd6\x4e\xd4\xbc\xd5\xd8\xef\xfe\x62\xbb\xd9\x5a\xcf\xb4\xba\x45\x75\xbf\xfd\x56\x9f\xc3\xc3\x46\x4a\xd2\xb7\x16\x63\xb1\xf6\x79\xe5\xd1\x1f\xe7\xc4\x84\xf4\xbc\x0c\x35\xf4\xaa\xe6\x2e\x41\x05\xb7\x8e\x89\xfc\xf3\x51\x7d\xf5\x2c\x51\x86\xd6\xf6\xe8\x2e\x39\xfe\x01\x7a\x74\x29\xa3\x8c\xa6\xb3\xde\x81\x4b\xb4\x52\xf2\xb2\x20\x61\x7e\x5f\x32\xbf\x33\xa1\xd8\x08\xd2\x90\x04\x8e\x94\x18\xda\x92\xb9\xaa\x07\x9a\x4b\x2d\x3d\x72\xe8\xb5\x8f\x8b\x75\xd2\xf8\x4a\x49\x9f\x7a\x64\xf6\x9e\xc3\x82\x74\xb8\x34\xb4\xe6\x96\x28\x48\x75\xe9\x66\x36\x11\x0a\xf3\x78\x18\xca\xa9\x8c\x6b\x9c\xa2\xa0\x3b\xb5\x25\x88\xbb\xc8\x92\xd3\xe7\xb9\x40\xdf\x51\xd9\xb3\x48\xca\x2a\x92\x27\x7e\x4e\x8b\x11\x00\x02\x2c\x64\x56\x44\x1f\xb7\xf7\x9d\xf6\xbe\x2c\x4e\xb5\xde\x5a\xd9\x13\x97\x38\x74\x91\x78\xed\xb3\x59\x3c\x11\xaa\x11\xc0\xb2\xc9\x7f\xe0\xba\x6a\x64\xd0\x96\xbd\x56\xf9\x06\x51\x36\x6c\xb0\xdc\xdd\xdf\x5e\xa1\x8a\xa6\xe2\x7b\x17\x15\xc5\x23\x1c\x0d\x97\xa4\x2a\x50\xab\xe4\xa0\xdc\x82\x80\xed\x10\xee\xe1\x2b\x23\xaf\x32\x3a\x16\x5a\xbe\x2e\x70\x70\x68\xcc\xb1\x9c\xad\xe9\x55\x1b\x1b\xef\x8b\xf5\xbd\xa9\x99\x3e\xa3\x22\x26\xdb\xa7\x13\x2e\x0c\xbb\x48\x74\x4d\x23\xad\x65\x3c\x12\x3e\xa4\x91\x2d\x85\x1a\xd9\xed\x45\x4e\xe1\x55\xeb\xd7\x9a\x08\x44\x05\xbe\x61\x01\xcd\x0f\x21\xdf\x6d\x6e\x1e\xba\x0d\x8b\x64\xd4\x46\x09\x31\x40\x1f\x55\xa1\x0f\x76\xc1\xf7\x6e\xc9\x0e\x58\x21\x39\xfd\x1a\x6a\xcb\x14\x3c\x71\xa1\xb8\x15\xb6\x00\xb5\x6b\xe2\xa4\xd4\xa3\xc0\x1e\x62\x6f\x62\x9f\x61\xe1\xac\x31\xf5\xdc\x74\xb5\xd1\xee\xa7\x36\x4c\x43\x8e\xfd\x1b\x79\x63\xbf\x95\x25\x1e\xf7\x4c\x49\xbf\x64\x49\xbc\xd8\x81\xea\x7c\xd9\x71\x32\xe7\xb6\xa3\x88\x3a\x3a\x6e\xad\xcf\x7c\x61\x3d\xcb\x4a\x49\xae\x07\xe1\x62\xf2\xd4\x34\x3b\xc4\x16\xeb\x14\xcd\xc7\xab\x5a\x37\x10\x72\xe9\x71\x3f\xde\xd8\x7f\xfc\xf1\x19\xfc\x5d\x4d\x4e\x4e\xd3\xd2\xa9\x68\x23\x6e\x75\xbd\x7b\xe6\xf2\x64\xf6\x26\xb7\x57\x0c\xd8\xc8\x97\x48\xd3\x69\x68\x49\xfc\x24\x45\xaf\x46\x54\xd3\x19\x6f\x03\x1f\x57\xfc\x6e\x21\x54\x36\x38\xd9\xda\x0b\xf4\xa9\xd7\x22\xca\xeb\x82\x95\x82\xc4\x5f\xb7\x67\xfb\x17\xf9\x25\xe0\x13\x8b\xa0\xf6\x19\x69\x4d\x6b\xf6\x1d\x8b\x7c\x89\xa2\xdf\x16\xdc\x61\xb5\x49\x56\x7c\xce\x93\xdb\xbe\x86\x9e\xe9\x53\x71\xda\xe5\x04\x6e\x25\x81\x0a\x36\xc5\x22\xc7\x57\xaf\xf4\x45\x17\x1e\x3c\x69\x9b\xd5\x2a\x88\xcd\x8a\xfc\xfa\x12\xd4\x24\x95\x3f\x1c\xae\x62\x61\xf1\xf1\x46\xe9\x95\x02\x79\xbb\x6a\x29\xfd\x47\xd7\x7e\xc3\x7c\xdf\x4b\xd6\xfd\x6c\x9c\x38\xbc\x95\x93\xb2\x31\x57\xc6\x83\xfc\x30\x22\xf4\x0b\xa5\xd7\xdb\x3b\x19\x12\x4e\x56\xc9\x26\xcc\x8b\xab\xa5\x35\x8f\x16\x94\x6c\x75\x47\x15\x4c\x9e\xc3\x5f\x6f\x74\xe2\x16\xcc\x50\x66\x0c\xc1\x69\xbf\x7d\xcb\x5c\x58\x77\xf6\xd6\x81\x1f\x83\xe4\x54\x26\x3e\x8b\x2e\xf3\xe7\xb2\xa4\xcc\x27\xae\x31\x5f\x6c\x7e\xad\x9f\x11\x4c\xe2\xd1\x8e\xfb\x3b\x86\x8f\x60\xd7\xbe\xae\xc3\xee\xa8\x09\x27\x59\x72\x2c\x3d\xec\x38\xc9\x0c\xe5\xff\xf2\x3b\x88\x5f\x02\xe8\x4f\x33\x50\x30\x66\xcb\x9f\x79\x0b\x74\x9a\xa9\x9c\x8e\x04\xcd\xa4\x04\x36\x2b\x6b\x95\xd4\x74\xe5\x2e\x2a\x50\x4e\xab\xba\x26\x96\x6f\x18\x6d\x0a\x9c\xe1\xf0\x13\xe6\xff\xb8\x0a\xee\xdd\x50\xd2\x8a\x93\x79\xbb\xd2\x98\xa0\x35\x10\x77\xfb\x86\xe3\x13\x85\xc9\x7c\xd6\xa3\xcf\x39\x66\xf0\x20\xa1\x5d\x27\x7b\xe6\x2a\xc2\x81\xdf\x79\xcf\x98\xd2\xac\xeb\xf5\x16\x29\x37\x45\x33\x46\x79\x5a\xc6\xac\x82\xb0\x71\x63\x74\xc2\x16\xfc\x67\x88\x23\x5a\x2c\xee\x4d\xc0\x28\x98\x2e\xf7\x44\x36\x5c\x4f\xcb\x3d\xa4\x49\x8e\xe4\x84\x11\x5d\x60\x52\x3c\x59\x6e\xb1\x55\x2a\xb4\x90\x7e\x4a\x9a\x96\x25\xc5\xd8\xae\x6e\x9a\x6d\x4f\xe9\xcc\x08\x7d\xe4\x39\x8e\x43\x43\x3a\x93\x7b\x57\x17\xc8\xc8\x33\xbe\x07\xc4\xa9\xba\x2a\xb0\x86\x5d\x07\x23\x93\x2c\x51\x45\x77\xff\xa6\xb6\x98\x61\x2b\x17\x74\x23\xc8\xd8\xef\xe8\x62\xea\x63\x6d\x17\xf9\x7d\x3f\x66\x8e\x67\xe1\x25\x71\x35\xf5\x55\x7b\x27\xaa\xa3\x17\xf2\x5e\x2f\xeb\x08\x13\xb1\x68\xc4\x0e\xbb\x78\xc6\xe3\xf9\x89\xd9\x6c\x45\x6f\x9e\x73\xd2\xf7\x77\x4b\x46\xb2\xa2\x89\x78\x1b\x86\x0d\x3d\xb7\xce\xb3\xe3\x63\xc3\x3f\x82\xc7\x33\xdc\xaf\xdf\x26\x94\x43\x0f\x13\x0f\xac\xa5\xd2\xc9\x5f\x0d\xa0\x0c\xa4\x7e\x40\x5f\xda\x96\x87\x7a\x69\x2d\x24\xc6\xc3\xd1\xe4\x97\x45\xd4\x46\x94\x7f\xaa\xb5\x04\xc7\x69\xf4\x3b\x0a\x6b\xe3\xdd\xf9\x42\xc9\xf4\x3a\xc8\x8e\xd8\x1d\x2b\x75\xaa\xd1\x7c\x79\x69\x64\xd0\xe9\x8c\xd5\x7a\x53\xde\x17\xec\xfd\x62\x69\xab\x6a\xd9\x2f\x54\x35\xd6\x05\x77\xd7\xed\xb9\xc1\xce\x6f\x2f\x82\xf5\x29\xe8\xe8\x19\x46\x02\xca\x8a\xa7\x55\x20\xc4\x34\x9b\x19\x63\x4f\x5e\xef\x2b\xd9\x8a\x9b\x8f\x6f\x0a\x75\xe7\x34\xbf\xc7\xda\x21\xfb\xa1\xfc\xc0\x54\xa5\xe0\x55\xb7\xdb\xcb\xaf\xfa\x5c\xcf\xaf\xb0\x10\xdb\x0e\xb7\x11\x2f\x1e\xa2\xf3\x18\xad\x1f\x45\xc7\xfd\xb9\xf5\xb0\xea\xed\x1c\xd6\xfa\x83\x28\x5b\xb1\xc7\x60\x9a\x8b\xa0\xce\xa4\x37\xa6\xa5\x71\x65\x5d\x35\xcc\x6f\xaf\x0c\x4a\x50\xb0\xea\x3e\xc7\x25\x3d\x8a\x97\xeb\x39\xcc\xe0\xd8\x17\x98\xca\x6f\x20\x90\x72\xe3\xfa\xab\xdb\xcb\x6c\x0f\x26\xc0\xe2\x02\x02\xcf\x30\xe2\x59\xb8\xdd\x48\x52\x78\x60\x6c\xc3\xec\xd6\x53\xce\x1f\x12\x90\x1e\x88\x76\xe3\x70\x86\x23\xca\x56\x33\x78\x61\xea\x52\xe6\xe3\x81\x8c\xd3\x19\xf6\x86\xe8\x83\xbb\x0e\xe1\x27\x2c\xf9\x0e\xd2\x2c\xa6\xd6\xcc\x8f\x00\x07\x26\xb7\xe8\xec\xf8\xc4\x81\xa9\x5d\x46\xfa\x69\x59\x96\x36\x65\xae\x4c\x79\x18\x86\xf5\xde\xde\xc3\x48\xfb\x48\x73\xac\x39\xdd\x05\x8c\x5e\xca\x13\x43\xe6\xbc\x54\xcb\x1a\xea\xa9\xb3\xc3\x6f\xc3\x39\xe2\xd7\x68\x91\x93\x08\x06\xad\xa5\xef\x99\x08\xbf\x60\xab\x3c\x47\xc3\x9b\xfa\x03\xcf\xa3\xf4\xcf\x29\xce\xa8\x34\x4f\xfd\x81\x97\xc4\x34\x58\x7b\x3e\x08\x9a\xad\x31\x31\xfa\x5b\x21\x0d\x70\x3f\x54\x56\x95\x37\xad\x44\xff\xb5\xad\x96\x74\x0e\xc5\xe2\xfd\x38\x79\x5d\x57\xc7\xc2\x9b\x44\xa4\xa5\xad\xc7\x09\x51\x1e\xfb\x1d\x7e\xd5\x41\xf1\x81\xb9\x80\xad\xdc\xd4\xb7\xcb\x24\x23\x02\xa3\x55\x6a\xb3\x0a\x2a\x0b\xbf\xb9\x51\xe3\xa6\xef\x2e\x53\xd9\x00\xcb\x0d\xf8\xc4\x55\xda\xc5\x7f\x13\xf1\x39\x38\x39\x2c\x23\x2d\xbb\xee\x5a\x5c\x8a\xcc\x63\xad\x85\x33\x37\xde\xf6\xfd\x56\x31\x76\x51\xd3\x68\xae\x7c\x4b\x4d\xf3\x9e\x05\xd9\x6f\xf6\xcd\x43\xeb\x9f\x35\x15\x21\xd4\xa0\x67\xb4\xbd\x03\xd7\x37\xfe\xda\xd4\x14\x33\xb6\xeb\xf9\xd1\x9e\x8c\xc6\x35\xf7\x83\x8c\x48\x7c\x0e\x0b\x6b\x95\x8d\x4c\x12\x20\x33\x17\x99\xa4\xc5\xf1\x55\x44\x71\xbe\xf0\xb5\x9a\xa6\x22\x23\xbf\x4b\xe2\x57\xdb\x75\xff\x11\xf4\x71\x8d\x1a\x33\x1f\xf8\x35\x84\x5f\xb1\x3b\xae\xb8\x21\x64\x54\xaf\xd6\xc1\x7f\x25\x98\xfa\x04\x1a\xce\x55\xa4\xfd\x27\xaf\xff\x8b\x4f\xea\x0b\xb3\x07\x22\x79\x93\x6a\xa8\x48\x82\x9a\x37\xbb\x4d\xd2\xe6\x9d\x90\xeb\x8a\xdd\x1e\x58\x1c\x4b\x9a\x36\xca\xe2\xd5\x71\x60\x6b\xe6\xe5\x70\x41\xd3\xc8\xa7\x81\x76\x06\x25\x94\xc2\xdc\xd7\x81\xba\x0e\x03\xca\x28\x98\x08\x42\xdb\x74\x75\x9a\x6f\x81\x58\xbb\xb1\xf5\xde\xb8\x5b\x0a\x2d\x7c\xf5\x72\x03\x35\x7a\xc0\x73\xd9\x16\xfd\x85\x39\x45\x85\x34\xae\x6b\x3f\x98\xd6\xe7\x44\x7c\xea\x08\x96\xe2\xda\xba\xca\xcb\x21\x77\x96\x70\xde\x4b\x0c\xf8\xf3\xc1\xc5\x89\x6c\x51\x9b\x77\xba\x69\xf4\x2a\x02\xf6\xa6\xdd\xc2\xb3\x66\x6d\xe2\xc3\xe4\x92\x8b\xad\x35\x1d\x49\xf8\x4f\xef\x8b\x3b\x6c\xd2\x3d\x47\xd5\x0e\x8c\x3c\xfc\x30\xca\xee\x08\x18\xed\x75\x92\xc3\x7a\xcb\x6d\xf2\xa1\x7d\x80\xe7\x98\x80\xe2\xc5\x7b\x5a\x14\xdc\x6a\xa6\x84\x56\x6d\xbc\xdf\xe7\x74\x29\x86\x66\xa1\xf6\x1b\xcb\x77\x93\xe2\x29\x57\x9e\x60\xaa\x8b\xdf\x34\x2c\x27\x23\x30\xcd\x8b\x06\x4a\x32\xf7\x24\x87\x3b\x3d\xe4\x6c\xfa\xf8\xd5\x51\x00\x9d\xe7\x7f\x29\x4c\x77\x1b\x06\xd0\x59\xa3\x00\x3a\x0b\x14\x40\x67\x4a\xcf\xf2\x7f\xbc\x32\x83\x21\x45\x00\x6c\x03\x05\xc0\x68\xff\xa5\x2e\xf4\x5f\xea\x22\xa9\xc9\x9f\xe6\x7b\x2a\xdd\xff\xa5\x2e\x0c\x80\x91\x89\x66\x0b\xf3\xdd\xc3\xe3\xe7\xf5\xe3\x9d\x94\xd4\xfc\x16\x9f\x56\x98\x89\xb5\x0e\x0a\xa4\xbf\xec\x68\x50\xce\xc6\x44\x61\xfd\x2c\xbc\x44\xc2\x35\x58\x1f\xf4\x8d\xcd\xc5\x56\x06\x06\x33\x09\x6f\x1c\xd5\x7f\xba\x25\x98\xe6\x88\xbc\x65\x2e\x16\x46\x9d\xce\x31\x77\x20\x7f\x7f\x16\x3a\x57\x5a\x9c\x51\x7d\xcf\xf4\x33\x77\x63\xc9\xcb\x95\x21\x56\x12\x24\x00\x87\x82\x35\xa1\x86\x74\xf4\xa6\xdd\x2e\x0e\x22\xcd\xec\x3d\xfe\xa6\x73\xdd\xfc\xaa\xe9\x74\x6d\xb8\x4b\x72\xd6\x59\xa3\x72\xc8\x59\x9e\xec\x38\x96\x63\x4e\x36\x0c\xad\x5e\xd3\x57\x8f\x53\x60\x89\xbf\x9d\x84\x05\x4f\x52\x6e\x96\x6f\x1a\x33\x6b\xac\x74\x54\xfc\x2c\x0e\xf4\xa6\x99\x1a\xee\x66\x16\xb5\x2a\xab\x20\xdf\x74\xb8\x38\xd6\x86\xda\x37\x0f\x34\x05\xbc\xf7\x9e\x44\x08\x42\x1e\x8c\xeb\x61\x22\x62\x47\xe2\x03\xdc\x7d\xcd\x6f\x5c\xb9\x48\xe3\x78\xb3\xd3\x19\xd3\x76\xff\x24\xfd\x12\x52\x18\x95\x69\x6a\xfb\x11\xa2\x55\xad\x70\x14\x48\x06\xf7\x04\x6c\xaf\xa4\x7c\x5e\xf4\x60\x1b\x34\x5a\x31\xf6\x2d\x12\x28\xd4\xdf\x8d\xe3\x66\xf5\x43\x9e\x2c\xaf\x3a\x0b\x5b\xa9\xce\x7a\xdf\xcc\x81\x1c\x47\xcf\x3e\x5c\xbc\x9d\x68\x6d\xa9\xe9\xf8\x44\x7d\xbf\x7c\xd7\xfa\xfc\x67\xff\x53\xc5\x02\x82\xff\xf7\x73\xda\xb1\xda\xe1\xae\x38\x1f\xc5\x91\x42\xce\xb0\xcb\xaf\x3f\x8c\xe4\x36\x3e\x69\x49\xc6\x95\xbd\x05\x27\x23\xe4\x7c\xe6\x0b\x95\xd8\xef\xd9\xfd\xf0\x5b\x8f\xa8\x3a\xf0\xf6\xc5\x4a\x60\x52\x4c\xba\xc3\x3f\xca\x76\x5c\x0d\x7b\xec\x0b\x18\xe8\x52\xc3\x17\x9f\xf4\x28\xab\x1b\xce\xfd\x85\xe7\xe2\x3c\xf4\x7e\x2b\xae\x10\x4f\x90\xbc\x89\xa3\xfc\x6e\x8f\x12\xab\x25\x19\xa9\x1e\xe2\x7d\xe7\x5b\x68\x6b\xbd\xb2\x63\x32\x2c\x4a\xa9\x33\xee\x9c\xa5\x72\xd6\x5e\x3f\x93\xba\x6c\x85\xdb\xe5\x5b\xe1\x9f\x4c\x57\x26\xef\x54\x5c\xfa\x48\xbb\x61\xb5\x94\xcb\xfa\x7e\x30\x3e\xd8\xe4\xfc\x27\xd7\x94\x10\x1d\xe7\x97\xcd\xcd\x4c\x3e\xe3\x7c\x7b\x39\x1a\x3d\x82\x98\x3d\x34\x7f\xed\x05\x20\x68\xde\x6a\x6b\xdf\x65\x3d\x8b\x6c\x17\x30\x1d\x06\x04\x14\xb7\x3d\x93\xe0\x67\xde\x61\x08\x3e\xa1\x13\xdc\xe5\xd7\xa3\xfd\xf3\x53\x1e\x1f\x9a\x10\x4f\xa7\xda\x4d\xad\xbc\x08\xf4\x6d\x0d\x44\xf6\xd2\x17\x38\xa1\x09\x36\x49\x66\x39\x9f\xd0\xb1\x02\xea\xfa\x3f\x18\x3c\x2e\x88\x8e\xf4\xe3\x23\xa1\x60\x83\x30\x9b\x63\x26\xca\xa0\x39\x38\x86\x39\x88\x5c\xf2\x4f\x01\x24\x6a\x3b\x3f\x71\xb1\x83\x29\x73\x34\x10\x9d\x7a\xd5\xb2\x46\xe5\x5a\xb1\xf5\xb2\xc8\x97\x06\xd0\x5a\xf9\x58\xb5\x09\x02\xec\xf6\x93\xfa\xf2\xbd\x2a\xa7\xc0\x44\x36\x94\x7d\x88\xbc\xfc\x7d\xea\x3e\x9c\xce\x64\xd9\x64\x68\x3f\x95\x4d\x28\x49\xf2\xfb\x69\x43\x10\xa5\x2d\xa3\x08\x06\x04\x8c\xeb\xc2\x99\xb1\x3e\x81\xb4\x00\x87\x14\x5e\x5e\xa1\x36\x9c\x62\x4c\xce\xd6\x8c\xc3\xc0\x97\xbb\xc1\xa9\x92\xbe\x85\x7f\x67\xe3\x48\xa3\x05\x14\x9c\x0e\xe0\x71\x35\x55\x71\xed\x6e\x21\xdf\x18\xa6\x33\x60\x24\x41\xb5\xab\xa9\xbe\x8e\x7b\xe0\x35\x3b\x14\x69\xce\x7a\x75\x4c\x57\xaf\xcb\x59\xc4\xfe\xc1\xd6\x3b\xb9\x46\xaa\xe7\x7b\x89\x5d\xc2\x6c\x5f\xc5\x4f\xab\x18\x1a\x66\xe4\xf9\xed\x72\x01\xf6\xae\xd0\x59\xf8\x5d\xaf\x14\xcb\xee\xd1\x3f\x52\xec\xf4\x4d\x4f\xf4\xcc\x31\xaa\xb4\x02\x0b\x9b\x53\xb2\xba\xe5\x10\xf2\xd7\xae\xaa\x4a\x82\x97\x45\x79\x4f\x3f\xf4\x4c\xe1\x45\x45\xfd\xc0\x41\x44\x80\x4a\x3f\x83\xd4\x8a\x1c\x26\x6c\xef\x33\xf7\xea\xf5\x7a\xe9\x1c\x97\xca\xbc\xb9\x93\xeb\xcd\x22\x7f\x7e\x63\x73\x91\xcd\x80\x63\x7a\xb9\xd7\x53\x80\x9e\xaf\xff\x8a\x7f\x60\xc6\xb1\x71\x3b\x5b\x09\xb5\xf3\xb2\x60\x58\x07\x62\x9b\x65\xf9\xce\x33\x4b\xaf\xb7\xb1\x7b\xa5\xe7\x53\x73\x47\x5e\xda\xe8\xed\xe0\xb4\x69\x45\xec\x58\x8e\x7c\x79\x8e\xf7\x73\x48\x9a\x0b\x0f\x5c\xa6\xf5\xce\x0e\x17\xc2\xb0\xb4\x53\xf3\x13\xe6\x5d\xc7\xcc\xa3\xef\x04\x15\x38\x4c\xe4\x16\xda\x14\x6f\x05\xb5\x03\x0d\x77\xfa\xa4\x4c\x6b\x66\x1d\xc6\x6b\x2d\x09\x38\x27\xee\xc7\x80\x17\x5d\x0f\xcc\x03\x8f\x1c\x4c\x02\x3b\xde\x25\x6f\xde\xe4\xfc\x29\xe2\xef\xe9\xc8\x52\x75\x99\xf8\x0b\x83\x46\x05\x0d\xdf\xfd\x9c\xd8\x61\xa3\xcc\xc7\xec\x1a\xad\x88\xe5\x2f\xb3\x85\x43\xab\xb9\x41\xde\xfd\xc7\xe4\xf6\xbc\x8d\x2b\x11\x95\x62\xc5\xf7\x93\xe8\x5d\xea\xeb\x2e\x4a\xd7\xbc\x12\x0e\xbf\x5e\x83\xad\xa9\xa0\x12\xae\xf1\xb0\x74\x47\x29\x29\x52\x4a\x91\x3b\xd9\xb9\xbc\x55\xaf\xc5\x86\x27\x07\x1a\x74\xc3\x63\xb7\x81\xe1\x8a\xe9\x8d\x34\x8e\x17\x8e\x21\x61\xf6\x36\x42\x52\xde\x6c\x2f\x18\x34\x22\xbe\xd8\x4e\x17\x5c\x94\x63\x61\x81\x2f\xda\x1c\x58\x2a\x54\xbd\x3b\x18\x4a\xe9\xa6\xd5\x56\x39\xeb\x76\x17\x31\xde\xaa\xb9\x9c\x3c\xdd\x64\x64\xb7\x42\xaf\x85\x07\xf9\x9a\xdc\x48\x3f\xa5\x41\x4a\xa5\x23\x6f\x4f\x3e\x80\xa3\x9c\xa0\x5e\x47\x82\x3e\xf1\x1f\x6d\x83\x99\x03\x52\x3d\xb3\xcc\x45\xe5\x42\x69\xac\x85\x4f\xa1\xc1\x7d\x3a\x51\xc1\xd4\xad\x9e\x2a\x5c\xf1\xce\x53\xc2\x45\x6c\x46\x73\x74\x5a\x93\x8e\x21\xb5\x37\x6a\xa7\x92\xdb\xfd\xc5\x97\xc3\x56\x0a\x2d\x02\xaa\xaa\xc4\x2d\x92\x5f\x14\x57\x4f\x50\x19\x34\x3a\x30\x91\x43\xc4\xa2\xe4\xba\x50\xaa\xb3\xeb\xc4\x50\xac\x8e\xa7\x01\x52\x56\x29\x47\xc2\x66\x44\x6d\xf7\x82\x5a\x4a\xa6\x89\xc4\xef\x0b\x5d\xcf\xf5\xa5\x15\x5e\x7c\xc4\x6c\x28\x48\x6c\x90\xa9\x8f\x16\x9f\x0d\x4e\x7c\x6d\x55\x76\x53\x36\x41\x49\xfe\xff\x3a\x32\x71\x2b\xf3\x22\xb8\x41\x8f\x40\x6e\x20\x1f\x90\x37\x28\x18\xe4\x02\x72\x07\x79\x80\xb0\x20\xff\xff\x0c\xf1\x28\xf3\xc2\xb8\x03\xb3\x92\x73\xaf\x45\xae\x0e\x29\xf6\x98\xad\x01\xcb\x14\x81\xcb\xa5\x8a\x8a\x28\x14\x4a\x1c\xc5\x44\x4f\x27\xcf\x0a\x85\x08\xd2\x81\x60\xdc\x52\x56\x03\x71\xc9\x22\xcd\xa4\x0e\x0c\x39\xe6\xfb\x55\x83\xd2\xc4\x6e\x97\x36\x18\xa3\x57\x9b\x6e\x05\xad\x9b\xd3\x11\x02\x06\xff\x4f\x00\x00\x00\xff\xff\xca\x48\x2d\x87\x29\x0d\x00\x00")

func certYckit_demoP12Bytes() ([]byte, error) {
	return bindataRead(
		_certYckit_demoP12,
		"cert/yckit_demo.p12",
	)
}

func certYckit_demoP12() (*asset, error) {
	bytes, err := certYckit_demoP12Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cert/yckit_demo.p12", size: 3369, mode: os.FileMode(420), modTime: time.Unix(1565602417, 0)}
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
	"cert/yckit_demo.p12": certYckit_demoP12,
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
	"cert": &bintree{nil, map[string]*bintree{
		"yckit_demo.p12": &bintree{certYckit_demoP12, map[string]*bintree{}},
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
