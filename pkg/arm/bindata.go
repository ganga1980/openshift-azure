// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// data/azuredeploy.json
// data/master-startup.sh
// data/node-startup.sh
package arm

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

var _azuredeployJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x5b\xeb\x6f\xdb\x38\x12\xff\xde\xbf\x82\xd0\x15\x48\x5b\xc4\xaf\x3c\xba\x87\x00\xf7\x21\x8f\xb6\x31\x9a\x34\x46\x94\xee\x7e\x58\x04\x05\x4d\x8e\x6d\x6e\x24\x52\x47\x52\x4e\x5d\x9f\xff\xf7\x03\xf5\xb2\x64\x49\xa6\xe4\x78\x77\x7b\x38\xf6\x43\x12\x89\x33\x24\x67\x7e\x33\x9c\x87\xba\x7c\x85\x10\x42\xce\x6b\x45\x66\xe0\x63\xe7\x0c\x39\x33\xad\x03\x75\xd6\xeb\xc5\x4f\xba\x3e\xe6\x78\x0a\x3e\x70\xdd\xc5\x3f\x42\x09\x5d\x22\xfc\xe4\x9d\xea\x1d\xf5\x07\xa7\x9d\xfe\xa0\xd3\x1f\xf4\x28\x04\x9e\x58\x98\x79\x0f\xe0\x07\x1e\xd6\xd0\xfd\x43\x09\xfe\x0f\xe7\x30\x5e\x81\x08\xae\x81\xeb\x5f\x41\x2a\x26\xb8\x59\x68\xd0\xed\x9b\x7f\xe9\x84\x00\x4b\xec\x83\x06\xa9\x9c\x33\xb4\x5c\x25\x4f\xe7\x58\x32\x3c\xf6\xa0\xf0\x50\x82\x12\xa1\x24\xd1\xc3\xdf\xa3\x47\x66\x2c\xb3\xdf\xa2\x49\x7a\x11\x80\x59\xe6\x96\x11\x29\x94\x98\xe8\xee\x5d\x00\x12\x6b\x26\x38\xf6\x86\x5c\xb1\xe9\x4c\xab\xde\xb3\x90\x4f\x2a\xc0\x86\xd7\x61\x91\x9e\x63\x3f\xa2\x5f\x2e\x51\xf7\x52\xf0\x09\x9b\x76\x6f\xc4\x74\xca\xf8\xf4\xb7\x94\x06\xfd\x07\xfd\xa1\xd0\x6a\xb5\x49\x8a\x03\x96\x3b\xe7\x51\x7f\xf0\x4b\xa7\x7f\xdc\x19\x9c\x76\x02\x09\x73\x06\xcf\x9b\xf3\x3d\x41\xa2\x7d\xe5\x96\xd3\x98\x71\x90\x2e\xc8\x39\x23\xd0\xbd\x49\x26\xd4\x2d\x18\x48\x11\x80\xd4\x2c\x16\x53\xe1\x5d\xf4\x5e\x3d\x85\x95\x2f\xa2\x97\x5f\x92\x83\x8e\x40\x7e\xba\x38\xea\x0f\xfe\xe9\x94\xe6\xad\x0e\xcb\x3c\x27\x80\x75\x28\x6b\x56\x8c\x57\x05\x2c\xc9\x6c\x2d\x89\x41\x99\xef\xab\xea\xbf\x56\x87\xaf\x96\xcb\x0e\x62\x13\xc4\x85\x46\x6f\x18\xa7\xf0\xbd\x42\x2c\xa3\xec\xd8\xdd\xf3\x29\x70\x3d\x12\xc2\x1b\x49\x31\x61\x1e\x28\xd4\x7f\xdb\xfd\x95\x83\x76\xc3\x31\x07\x3d\xbc\x42\xab\x55\x73\xa4\x7c\x01\x6d\x80\xd1\x9b\x33\xa9\x43\xec\x25\x7f\x96\x20\x52\xd2\xf3\x7b\xa3\xe7\xe3\xfe\xde\xf5\x9b\x62\x71\xce\x41\xb7\xd5\x3d\xa6\x54\x82\x52\xae\x01\x6c\xbd\xae\x92\x59\x23\x09\x13\xf6\x7d\xc3\xb0\x4a\x93\x07\xfd\xd8\x74\x7b\x15\x50\x31\xe3\xb1\x11\x80\x54\xa4\x99\xfa\xa5\xaa\x77\x8a\xf2\xf2\xa0\x30\xc1\xa1\xb7\x29\x92\xc2\x54\x8b\x78\x0a\x73\x0b\x42\x88\x5c\x54\x7a\xd0\xa3\x93\xea\x93\xa2\x12\x88\xeb\x9f\x3e\x6e\x07\x3b\x70\x6a\x30\xba\x5c\xf6\xde\xa1\x87\xbb\xab\xbb\x33\xf4\x0c\xc8\xc7\x0b\x34\x06\x64\x3c\x20\xd2\x02\xe1\xb9\x60\x14\xe9\x19\xa0\xe1\x08\x61\x4e\xd1\xcd\x05\x9a\x81\x04\x63\x29\xcf\x80\x94\x16\x01\x0a\x15\xe3\x53\xf4\xae\x97\xf2\xba\xc2\xe0\x0b\xee\x82\x56\x68\x22\x24\x72\xaf\xbe\x74\x11\x7a\x98\x01\x47\xb7\x58\x69\x90\x88\x0b\x0a\x0a\xa9\x99\x08\x3d\x8a\x88\xf0\x01\x85\x01\xc2\x0a\xdd\x03\xa6\x8b\x1c\x23\x1c\x6a\xe1\x63\xcd\x08\xf6\xbc\xc5\x61\xb4\xfc\x0c\x38\x01\xb3\x41\xe0\x1a\x24\x50\xc4\xb8\x16\x88\x08\xae\x18\x4d\x7c\x6d\xb4\x28\xce\xb1\xb9\x11\x98\x5e\x60\x0f\x73\x02\x12\x25\x06\x80\x26\x52\x70\x6d\xf6\x6d\xce\x76\x3e\x1a\x22\x05\x72\x0e\x32\x26\x4b\x05\xd5\xd8\x6c\x83\x70\xec\x31\x32\x1c\x9d\xc7\xfa\x2c\xfb\xf6\xbf\xdc\x70\x59\xd0\xc1\x01\x8b\x0f\xd5\xd6\x80\xb3\xd3\x78\xe9\x76\x6e\x41\xcf\x04\x35\x7c\xaf\x16\x1c\xfb\x8c\x54\x18\x80\xc3\xa8\x07\x0f\xcc\x07\x11\xea\x21\xbf\x65\x3c\xd4\xd1\x02\x83\xd3\x8a\xb9\x94\x2b\x17\xb4\x51\xc0\x16\x77\x4e\x85\x8f\x19\x37\xb7\xc5\x0d\x1e\x83\xb7\x71\x37\xc6\x60\xba\xb9\xb8\x34\x13\x62\x0b\xca\x84\x62\xf3\xfb\x1b\x02\xa9\xbb\xad\x32\x69\x5e\x60\xc5\x88\x53\x6f\x4e\xe9\xaf\x8d\x01\xe3\xe5\x30\xb9\x33\x58\x28\x04\xc0\xa9\xba\xe3\x95\x1e\xcd\xf9\x3d\x8d\x5a\x86\xf4\xcd\x41\x03\xcc\x1e\x1c\xa2\x83\x3c\x6c\x0e\xde\x3e\x16\x8f\xfc\xf8\x67\x81\xd5\x1b\xef\x0e\xd6\x31\x26\x4f\xc0\x69\x72\x0a\x73\x21\xbf\xc8\xc1\x27\xec\xaa\x9d\x6f\x85\x8b\xad\x8a\x53\x64\x14\x79\xd2\xe1\x28\x46\x6a\x18\xfb\xa5\x17\x6d\x2b\xe5\xb9\xaf\x8b\x27\x90\x6c\x8e\x35\xb4\xb5\xf1\x22\x8f\x22\x80\xac\x8b\xa2\xd8\x49\x98\x15\xf6\x0e\xcd\xcd\x51\x7d\x45\xd6\xbf\x69\xa6\x59\xc6\xc7\x22\xe4\xf4\x0b\xd6\x19\xce\xb6\x4f\xbb\x0f\xe3\x44\xa2\x72\xda\xda\x07\x30\x3e\xcd\x66\xee\x8a\x90\x40\x48\xdd\x39\x39\x39\xde\x17\x42\xca\x76\xd5\x4a\xc1\x44\x70\x82\xf5\x9b\xed\x7a\x2e\x78\x41\xa3\xe3\xbc\x23\x38\x78\x7b\x88\x0e\x7a\x15\xe6\x9d\x3e\xb3\x83\xc0\x02\xe0\x84\xcf\x48\x48\xed\x9c\xa1\x93\x93\x63\xcb\x7c\xe0\x26\x2c\xfa\xe8\x09\x6c\x2e\xae\xe1\xc8\x39\x43\x13\xec\x29\xb0\x90\xd5\xf8\x83\xbf\x45\x9c\x75\xbe\x29\x7b\xf1\x62\xa1\xa6\x8c\x1a\x4b\xb5\x45\xd8\x50\xa0\x33\xa7\xbd\x62\x4a\x4b\x36\x0e\xd3\x5b\xe8\xca\x1a\x9e\xa3\xc4\x0e\xc6\xf5\xa9\xc9\xc6\xee\xf6\x2a\xff\x68\x65\xd5\x4b\x8d\xf5\xc5\xd2\x0e\xa4\xd0\x82\x44\xb6\xe9\x3c\x90\x60\x0f\x39\x43\x85\xa3\x12\xa1\x6e\xe4\xd0\xe2\xc3\xfd\x4c\x4e\x8c\x99\x2c\x61\x8e\xbd\x21\x77\x81\x08\x4e\x0d\x89\x0d\x57\x3c\xf4\xc7\x20\xef\x26\xa3\xf4\x34\x47\x36\x1d\x34\x45\xfa\xfe\x95\xf5\x93\x44\xb5\xae\x16\x12\x4f\xa1\xa7\xe2\x9f\xe7\x84\x88\x90\x6b\x7b\x5c\x7b\xda\xe9\xbf\xef\x0c\x4e\xff\xb4\x24\x28\x97\x2d\xdc\xc3\xd4\xf8\x8a\x85\x5b\xd8\xe2\xae\xe5\x2d\x1c\x93\x3f\x24\x92\x70\x35\xe6\x14\x4b\xfa\xed\xe6\xde\xfd\x3f\x91\x67\xfc\xe3\xaf\x96\xe6\x06\x63\x8d\x6b\xd2\xc7\x4c\xa6\x24\xda\x66\xad\x4e\xa2\x3a\x88\xc4\x7c\x0a\xe8\x35\x0e\x02\x74\xf6\xaf\xb6\x45\xbf\xd5\x2e\xd9\x1f\x8f\x7f\xba\x40\x42\xc9\xf4\xe2\x93\x14\x61\xb0\xaf\x92\xc1\xeb\xdd\xd5\xcb\xd5\xb4\x63\x38\xe0\x20\xe8\x9a\xa4\x7a\x97\xba\x6f\x72\xa4\x17\x87\xb3\xd8\xf3\xc4\xf3\x37\xa5\x66\x7b\x2b\xb5\x11\x12\x27\x29\x8e\x49\x79\x36\x2b\xe2\xa5\xe9\x14\x14\x91\x2c\x48\x05\x1b\xd1\x20\xd7\xbd\x46\x5a\xe2\xc9\xc4\x9e\x1c\x51\x50\x9a\xf1\x48\xec\xe7\x9b\x45\xbe\x77\x2d\x88\x4d\x18\x75\x6f\x00\x1a\x81\xe0\xa8\x73\x74\x64\x25\x66\x12\x48\xba\xef\x61\x9c\x8a\xd8\x03\x22\x26\x8c\xda\x4c\xdc\xd5\x1f\xb4\xbc\xc6\x2c\xd3\xe3\x88\xa9\xbd\x10\x62\xba\xc2\xf9\xdf\xb5\xbf\x32\x93\xaa\x3e\xfc\x3b\x86\xf5\xbd\xf0\x00\x39\x7e\x54\x3a\x72\x0a\xc6\x9b\x1f\x8d\xf1\x19\xb5\xae\x7e\x26\x84\x5e\x3f\x3c\x8c\xdc\xbf\x15\xa3\x27\x27\xc7\x96\x08\x0e\xed\x05\xa5\xd6\xa8\xec\x7f\x0d\xa5\x49\x39\x7e\xf3\xe5\x96\x22\x7e\xfa\xab\xf5\xde\xb9\x14\x7e\x10\x6a\x48\xbb\x4b\xb7\x98\xcc\x18\x07\x97\x60\x0f\x5c\x68\x10\x57\xfc\xd2\x19\x1c\x75\xfa\x83\xfd\xdf\x3c\xc5\x0a\x66\xae\x03\x17\x19\x6b\x5d\x37\x2d\x23\xb7\x94\x91\x36\x9a\x69\x26\x23\x9b\x73\xd0\x26\xe9\xda\xe8\x80\x6c\xf5\x11\xad\xd7\xb5\xe4\x81\x1b\xab\xb7\xe5\x5e\x19\x3c\x98\x55\x4a\xd7\x77\x22\x74\x6b\x1d\x37\xf5\x69\x4a\x55\xd2\x6f\x2a\xad\x36\xaf\xd0\x0c\x64\x3e\x80\xab\x6a\x12\x10\x1c\x60\x12\xdb\x70\xba\xd6\x65\x3e\x76\xac\x20\xc9\x45\x9f\x31\x2c\x6e\x5d\xf6\x03\xaa\x6b\xfd\xeb\x3e\xae\xd1\x66\x1a\xac\x0e\x7d\x3c\x85\xfb\x44\xaa\xd1\xd1\x9c\x92\x62\x9d\xc0\xc3\xd5\x35\x99\xc2\x06\x0a\x2c\xdd\xcf\x5f\x6b\x84\x84\xb2\x02\xa9\x9a\xc5\x42\x29\x11\x8f\xd2\xb7\x5b\x59\x48\x41\x43\xa2\x2b\x19\xdc\x4d\x26\x39\xe2\x2a\x31\x54\x20\xcc\x1a\xc1\x89\x39\xc8\x40\x8a\x39\x4b\xec\xbf\xa6\xc2\xe5\x84\xc1\x54\x62\x0a\x23\xe1\x31\xb2\xa8\xef\xe1\xf8\x82\xc6\xce\x08\xf3\x10\x7b\xcd\x1a\xfd\x45\x37\x95\x84\xda\xf5\x4b\x08\x65\x9b\x82\xe2\x5e\xab\xcf\xf8\x57\x05\x32\x55\x27\xf1\x44\x48\x3b\xa1\x2a\xf5\x1c\x0a\x64\x24\xf6\x9d\x72\xdd\x65\xca\x63\x31\x6f\x27\x9d\x6d\x6c\x3c\xc6\xc3\xef\xed\xea\x7f\x0e\x65\x0a\x8f\x3d\x18\x61\xa5\x9e\x85\xa4\xe7\xa1\x9e\x01\xd7\x2c\x73\xba\x5a\x86\xb6\xda\xa3\x09\xa0\x1b\xd5\xb9\xe2\x92\xfb\x67\x58\x6c\xef\xc4\xe7\x87\x9d\x6b\xc6\xfd\x09\x16\x57\x58\xe3\x44\x72\xae\x7b\x3d\x4a\x97\x3b\x57\xae\x96\x8c\x4f\xd7\xd8\x76\xdd\xeb\xcf\xb0\xe8\x66\x33\xb6\x98\x47\xfd\x69\xb0\x36\xe7\x76\x7a\x33\xe1\x43\x6f\xad\xe8\x5e\x57\xa9\x59\x0f\x87\x7a\x26\x24\xfb\x01\xf4\xdb\x93\x39\x70\x23\xbe\xf5\x1d\x85\x74\x94\x3f\x42\x68\x46\x5f\x13\x14\x54\x1f\xd7\x49\x6a\x02\x8d\x20\xcf\x62\xbf\x37\x01\x09\x3c\xf9\x16\x63\x47\xef\x58\x62\x2d\x8c\xf3\x69\xe0\x96\x9a\xf4\x90\x5e\xe4\x22\x8b\xc2\x89\xae\xa6\xb6\x8e\xba\xc0\x62\xbe\x8e\x7c\x4a\x6c\x92\xa8\x68\xed\x73\x23\x27\xeb\x29\xb0\x8a\xab\xa2\xed\x55\xe2\x9e\xaa\x20\xba\xd1\xb3\xcb\xfb\x10\x1d\x94\x03\xb9\x48\xb1\xd1\xa5\x5f\xcb\xa5\x14\x01\x6c\x09\x39\xd2\xb1\xa5\xda\xec\x08\x75\xc5\xd4\x93\xdd\x69\x91\xc8\x6b\x4f\xcd\x71\xef\x01\xd3\xdf\x24\xd3\x60\x93\x39\x91\x80\x35\xdc\x65\xb9\xcc\x47\x29\xfc\xe8\x30\x36\xc2\xf8\x5b\x41\xda\x68\x67\x28\x67\x3d\xe7\xc5\x8a\xd3\x48\x82\xcf\x42\xbf\x5c\x70\xda\x1c\xdb\x8c\x78\xa7\x4c\x33\xda\x14\xc5\x1a\x9b\x23\xd8\x5d\x6f\x83\x13\x52\xa6\x9e\x4c\x7c\xf4\xe9\xc2\x39\x43\xc7\x96\x2c\x09\x55\x49\xff\x83\x1f\xe8\x45\x03\x6f\xeb\x78\xa1\x99\xdf\xdf\x51\x62\x8f\x36\x44\xd6\x79\xc0\x24\x00\x6e\xe4\x01\x93\xb9\x43\xae\x41\x4e\x30\x81\x86\x1d\xf9\x74\x34\x90\x77\x56\x3c\xb3\xe6\xdb\xa8\x65\x41\x20\x47\xc3\x7c\x2c\x17\x8d\x6e\xfc\x8c\x28\x6e\x56\x0e\x47\x1f\x85\x7c\xc6\x92\xc6\x26\xd9\x82\xbe\x2a\xcd\x68\xbc\x65\xd4\xb8\xd5\xdf\x3c\x9b\xa9\x48\x64\xea\x86\xa5\x6b\xb6\xde\x61\xd0\x12\x0f\xf9\xd1\x5c\x12\x08\x15\x3e\xcc\x4a\x6a\xd2\xcd\xc3\x19\xb4\x23\x70\x36\xe8\xdb\x83\xa8\xc0\x20\xfe\x7a\x72\xa7\xc5\xd1\x1a\x0e\xcb\x25\x9b\x94\xb3\xfb\xd5\x6a\xb9\x2c\xa7\xfc\xc9\xed\xb5\x5c\x9a\x2b\x76\xb5\x6a\xd6\x84\xad\x4d\xfa\x0f\xd1\x41\x2f\xf9\x02\xb4\x97\x7c\xc6\x79\xf0\xf6\x71\xb9\x04\x4e\xab\xbe\x18\xb3\x8d\x86\x18\x2b\x88\x20\xf9\xae\x25\x48\xca\x4c\xed\x3f\x07\xa8\xe4\xba\x46\x56\x4b\x4c\xad\xf7\xf5\x32\x6c\x65\x7c\x6a\x5b\xf9\x3b\xb1\xb4\x47\xdc\x65\x8a\x9d\x2f\xe1\xba\xe1\xe4\x8b\x39\x17\x2d\xbe\x35\xb3\x8d\x17\x89\xf9\xa7\xf9\xc4\xa6\x6e\xb4\xd7\x9d\x35\x1c\x78\xf9\x52\xcd\x66\x6e\x4f\xe3\xec\x7c\xb6\x45\x3c\x95\xf3\x6b\xc2\x1c\xf8\xae\x81\x9b\x74\xa3\x51\xa0\x93\xcd\xde\x6b\x50\x43\x94\x2d\x04\x47\xbb\x06\x35\x38\xd4\xe2\x6b\x5c\x43\xba\x65\x5c\xc8\x75\xc5\xb9\x45\x90\x12\x48\xa1\x81\x68\xa0\xd6\x8f\x89\x2b\xc9\xe3\x06\x4a\x92\xea\x5d\x60\x05\xef\x4f\x3e\x70\x22\x28\xa0\x37\xae\xc6\x52\x87\x41\xce\x8d\xbc\xad\xff\xba\xb8\x6a\x34\x8d\x41\x0a\x19\xf0\xda\x7e\xcf\xa3\xff\xfa\xf4\x61\xad\xd5\x86\xec\x54\x4e\x10\x4d\xb7\x90\x36\x0c\x2e\x43\xa5\x85\xef\xc6\x42\x69\x41\x7b\x8d\x39\xf5\x40\xe6\x7b\x06\xdd\xbe\x5d\x4a\x7b\x36\xa3\x72\x69\xb1\xae\x73\xb2\xe9\x69\x92\x7a\xb8\x23\x42\x1d\x84\x3a\x16\xdd\xab\xd5\xab\xff\x06\x00\x00\xff\xff\x20\xd8\x4d\xf4\xaa\x36\x00\x00")

func azuredeployJsonBytes() ([]byte, error) {
	return bindataRead(
		_azuredeployJson,
		"azuredeploy.json",
	)
}

func azuredeployJson() (*asset, error) {
	bytes, err := azuredeployJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "azuredeploy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\xfd\x73\x1a\x39\x96\x3f\x6f\xff\x15\x6f\xf1\xec\x65\x37\x95\x06\xec\x38\x99\x84\x89\xb7\x0a\x63\x9c\x50\x21\x36\x0b\x24\x7b\x73\x73\x53\x2e\xd1\xfd\x00\x8d\x1b\xa9\x23\xa9\x49\x98\x24\xff\xfb\x95\x3e\xfa\x0b\x1a\x8c\x89\x67\xb7\xee\xea\x3c\x55\x19\x90\xde\x97\xde\x97\x9e\x9e\xba\x39\xfa\x73\x63\x42\x59\x63\x42\xe4\x1c\x7c\xfc\xec\x79\x47\x30\xbe\xbe\xb8\x6e\x41\x03\x55\xd0\x08\x99\x5c\x10\xf9\xb1\x1e\x36\xb8\xa0\x33\xca\xfc\x24\x96\x4a\x20\x59\xf8\x21\x93\xf5\x80\xb3\x29\x50\x09\x41\x22\x04\x32\x15\xad\x60\x4e\x44\x18\xf0\x10\xc3\x9f\x80\x2a\xef\x08\x62\xc1\x27\x64\x12\xad\x40\xce\x79\x12\x85\xec\x91\x82\x09\x7a\xde\xa8\x3b\xfc\xd0\xeb\x74\x6f\xc6\x3f\x0f\xba\x67\x96\xb2\x47\xa7\xf0\x0b\xf8\x53\xa8\x19\xc6\x72\x25\x35\x75\x3a\x6b\x10\xc5\x17\x34\xf0\x79\x8c\x4c\xce\xe9\x54\xf9\x8c\x87\x58\x83\x5f\x7f\x02\x35\x47\xe6\x01\x00\x94\xc8\xad\xc3\x7b\x53\xea\x69\xe2\x7f\x86\x99\xc0\x18\x1a\x4b\x22\x1a\x11\x9d\x34\x42\x1e\xdc\xa2\xb0\xcb\x9c\x4a\x45\x26\x19\xc1\xc5\xed\x54\xd6\x3f\x4f\xa5\x96\xa6\x11\xe2\xb2\x11\x52\x79\xdb\x20\xbf\x27\x02\x1b\x02\x25\x4f\x44\x80\x7e\x4c\x84\x3a\xf6\x00\x30\x98\x73\x78\xb4\x1b\x0c\x36\xb8\x82\x26\x0f\x33\x11\x7f\x4c\xb8\x22\x00\x4d\x68\x3e\x82\xbf\xff\x3d\x17\xc6\x03\x90\x2b\xa9\x70\x11\xa8\x08\xa4\xe2\x31\x58\xcc\xba\x44\xb1\xa4\x01\x6a\x31\x79\xc2\xd4\x3a\x65\x0f\x40\xa0\x54\x5c\x60\xc0\x19\xf8\xc3\x8a\xf9\x80\x28\xb0\x9c\xec\x50\x23\x24\xb8\xe0\xac\xfe\x9b\xe4\x0c\x5e\xbd\x7a\xd4\xbd\xbe\x7c\xe4\x7d\xf1\x00\x6a\x11\x9f\xf9\xa1\xa0\x4b\x14\xb5\x16\xd4\x7e\xe3\x89\x60\x24\x0a\x6b\xde\x37\xaf\x7b\x7d\xb9\x26\x21\x11\x6a\x5d\x44\xad\xf8\xd4\x9b\x02\xce\x24\x0d\x51\xc0\x94\x04\x0a\xd4\x9c\xa8\x0d\xd5\xca\x40\xd2\xe3\x46\x94\xb0\xe6\x36\x9f\xd2\x0e\x45\x84\xa2\x8a\x72\xb6\x03\xfd\xa7\x02\xb7\x44\x25\x02\x41\x2a\x41\x14\xce\x56\x30\xe5\x42\xeb\x87\xfe\x8e\x12\xe8\xd4\x3b\x02\x86\x18\x62\x58\xe5\x1f\xa8\x82\x70\xa7\x77\xec\x90\xff\xeb\x57\x50\x22\xc1\xad\xee\x51\x00\x5d\x63\x68\x1d\x23\xc4\x29\x49\x22\x25\xab\x1d\x63\xcd\xf0\x1a\x6f\xbb\xd9\xcd\xac\xb6\x84\x91\xa4\x76\x7e\x7d\x3d\x1e\x8d\x87\xed\xc1\x4d\xe7\xfa\xea\xb2\xf7\xfa\xe6\xaa\xfd\xae\x7b\xa6\x43\xca\xb7\xf1\xe6\x2f\x88\x54\x28\x6a\x29\xd3\x3c\x10\x7f\xf8\x52\x8c\xb3\x6f\x26\x0e\x3d\x4f\x62\x08\x3e\x05\x1f\xa1\x26\x8f\x2e\xba\xe7\xef\x5f\xdf\xf4\xaf\x5f\xf7\xbb\x1f\xba\xfd\xb3\x93\xf5\x81\xd3\xa3\x1a\xec\x45\x55\xdb\x29\x94\x0a\x28\x03\x15\xc4\x4f\x4e\x9e\xbe\x68\xfe\x04\x21\xf7\x8e\x36\x26\x7e\x7c\x99\x41\x98\x0f\x2f\x4e\x4f\x9f\xa6\x1f\x4e\xed\x87\xe6\xb3\xa7\x90\x84\xee\x83\x1e\x79\xd9\x7c\x69\xc9\xfd\x29\x16\x5c\xf1\xb3\x1f\xbe\x84\x52\xfd\xe5\x2f\x4f\x1e\x7f\xf3\xfe\x14\x73\xa1\xec\xc0\xd1\xd1\xe3\x27\xdf\xbc\x3f\xd1\x58\x91\x49\x84\x12\xfc\x36\x5c\x8f\x6e\x2e\x7b\xc3\xee\x3f\xdb\xfd\xfe\x4d\xbb\xdf\xbf\xfe\x27\xf8\x31\xfc\x60\x88\x80\xbf\xd0\x51\xa0\x10\x7c\xdf\xfe\xff\xaa\xfb\x4f\x3d\x98\x4e\xfb\xa1\x26\x0d\x3f\x98\x7f\xfd\xdf\xa0\xdd\xe9\x74\x07\x63\x2f\xe4\x0c\x3d\x2f\x65\xe2\x4b\xb2\x44\x58\xd7\x7c\x3a\xeb\x79\x62\x01\xbe\x98\x5a\x1d\x6a\xcb\x36\x1e\xdb\xcf\x36\x87\x36\xac\xed\x1a\x8f\x3d\x6f\x42\x24\x3e\x3f\x05\x3f\x84\x57\xaf\x5e\xc1\x97\x2f\xd0\x41\xa1\xda\xf2\x7c\xa5\x50\x42\xbd\x63\xe8\xd6\xf5\x18\x9d\xd2\x80\x28\x94\xf5\xae\x0a\xc2\x0e\x31\x63\xf0\x15\xce\x0d\x7e\x97\xe9\xc0\x83\x6f\xdf\x9c\x48\x86\x65\x40\xea\x81\x50\x07\x72\x18\xa1\x58\xa2\xd8\x83\x8b\xb4\x80\x95\x9c\x06\x82\x2e\x89\xc2\xb7\xb8\xda\x97\xdf\x5b\x5c\xed\xc5\xee\x16\x57\x07\x2e\x6c\x80\x7b\x2d\x2b\xc6\x07\x58\x94\xe1\x75\xe7\x92\x0c\x2b\xbd\xa0\x4d\x5e\x3f\x93\x45\xf4\x8e\x08\x39\x27\x51\xc6\xa5\x1d\x2e\x28\x7b\x9b\x4c\xd0\x3a\xdd\x56\xda\xce\xd5\x74\x9c\x9a\x7f\xea\xb7\x19\xce\x01\xaa\xbb\xcb\xe3\x8a\xdc\x9c\xe3\x79\x8b\xdb\x90\x0a\x1d\x79\x15\xae\xcf\xc8\x02\xc3\x3f\xc8\xfb\xcb\x9c\xec\xff\xea\x5a\xd5\xfe\xc1\x21\x71\x3f\x96\xdb\xd8\xec\xe9\x3a\x1d\xb2\xd3\x69\x36\x78\x1d\x16\x0c\xd7\x69\xe1\xd5\xe1\x4c\xf2\x08\xef\xb3\x40\x63\xbd\x46\xe0\x10\xbf\x67\xad\x1b\x52\xec\xbf\xf2\xb2\x10\x87\x29\xe1\x52\x70\xa6\x06\x82\x7f\x5e\xdd\xcf\xc2\x53\x8d\xe7\xc7\x1a\xf1\x70\xa7\x1a\xd9\xfa\x6b\x44\x67\x8c\xb2\xd9\xfd\x04\x70\xb5\x9b\x2f\xe9\x8c\x7d\x67\xa6\xda\x10\x63\x7f\x13\xac\x49\x71\x78\x56\xee\x44\x14\x99\x3a\x38\xac\x2d\xf6\xf7\xa6\x6b\x27\xc4\xfe\xcb\xaf\x90\xe1\x30\x15\xb4\x67\x33\x81\x33\xa2\xb8\xc8\x1d\xf2\x3e\xca\x20\x19\xbe\x5f\xf0\xcc\xef\x52\x48\xa5\x48\xfb\xab\x66\x8b\x44\x87\xa9\xe7\x9d\xa1\xa9\xf7\xbc\x08\xd5\xc1\xae\x72\x6b\xf1\x1f\xc2\x5b\xaa\x04\xba\xb7\xdb\xac\xc9\xf3\x3d\xaa\xb1\x19\xec\x50\xc5\xb8\x34\xf6\x50\x6a\x29\x0a\x73\x6f\xa5\x94\x64\xf9\x1e\x95\xec\x53\xc0\x56\x4a\xf0\x00\x05\x6d\x49\x82\x7b\xab\x60\x57\x89\x3b\x48\x26\x11\x0d\x2a\xf8\xbb\x24\xde\x0e\x02\x7d\xe8\x7c\x8b\xab\x7a\x06\x7a\xbf\x5c\x4e\x2c\x05\x59\x8f\x0d\xfe\x16\x31\xb6\xea\x61\x43\x8e\x43\xb9\x5b\x0e\xd5\xec\x53\x66\x6f\xd4\x80\x48\xf9\x29\xdc\x93\xc7\x5c\xc5\x06\x7c\xbf\x3a\x3b\x8f\xf1\xfd\x0a\x6d\xc7\x24\xef\x7f\x15\xc2\xdc\x95\xdd\x5e\xde\xd2\x31\x95\xbf\xfe\xc7\x36\xe7\x5e\xbd\xea\x5e\x5f\x7a\xdd\x71\xe7\xe2\xa6\x7d\xf1\xa1\x3b\x1c\xf7\x46\xdd\x9b\x4e\xbf\xd7\xbd\x1a\xdf\xbc\x1f\xf6\x47\x67\x73\xa5\x62\xd9\x6a\x34\x7e\xf8\xeb\x9c\x4b\xa5\x4b\x9f\xbf\xb5\xf4\xa1\xda\xe2\x74\xba\xc3\xf1\xcd\x65\xaf\xdf\x3d\xab\x3c\x98\x59\x18\x4b\xcd\x80\xb6\xdf\x8f\xdf\x9c\x99\xce\x87\x99\xba\x68\x8f\xdb\x37\x17\xbd\xe1\x59\xb9\x1b\x61\xe6\xba\xfd\x6e\x67\xdc\xbb\xbe\xba\x19\xf7\xde\x75\xaf\xdf\x8f\xcf\x4e\x9e\x35\x9b\x76\xea\x4d\xb7\x3d\x1c\x9f\x77\xdb\xe3\x9b\xde\xd5\xb8\x3b\xfc\xd0\xee\x9f\x65\x73\xbd\xab\xde\xb8\xd7\xee\x17\x56\x33\xe8\x76\x87\xbb\xd6\xf2\x62\x0d\xb3\xd3\x7f\x3f\x1a\x77\x87\x67\x56\x8d\x7e\xd3\xfc\x65\xb8\xa5\x51\x83\xfd\xa4\x38\x74\x5c\x09\x78\xbc\x09\x78\x52\x09\x78\x52\x90\xe7\x6d\xf7\xe7\x2d\xaa\xd5\xbe\x69\x40\xfa\xbd\xd1\xb8\x7b\x55\x69\xaf\x66\xdd\xfc\x57\xb0\x95\x03\xde\x54\x47\x0e\x9a\xb2\x36\x3d\x9f\x82\x96\xec\xa8\xc1\xac\xb2\x78\x76\x66\x2d\x80\x6d\x37\xba\x99\xaf\x58\x5c\x76\x1c\xcd\xa1\xc6\x43\x6d\x8a\x8b\x9b\x4e\x7b\x1d\xd8\xd5\xbe\x06\xf4\x1f\xef\xaf\xc7\xed\x9b\xf3\x76\xe7\x6d\xf7\xea\xe2\xe6\xfc\xe7\x71\x77\x74\x76\x7a\xf2\xf2\xf4\xe5\xf3\x1f\x4f\x5e\x3e\xb7\x30\x77\x53\xba\xbe\xf4\xbc\xa0\x7c\x64\x2c\x1c\x2a\x2b\xc6\xcd\x56\x91\x16\xe1\x77\x61\xc6\xb7\xb4\x11\x10\x5f\x89\x44\xaa\x86\x6d\xfd\x36\x08\x0b\xe6\x5c\xc8\x42\xe4\x3a\x62\x49\x1c\x12\x85\x7e\x0a\x5f\x0c\xdf\xaa\xc4\xed\x1a\x73\xf5\x15\x59\x44\x2e\xa0\x49\xb8\xa0\x52\x52\xce\x6c\x4e\x69\x79\x00\x71\x94\xcc\x68\xe1\x3b\x40\x3b\xfa\x44\x56\x72\x90\x44\x51\x6f\x41\x66\x28\xed\x28\x80\x25\x97\x08\xa2\x28\x67\xe9\x20\xc0\x2d\x65\x61\x0b\x2e\x6c\xe7\xb1\x5d\x66\x90\x01\x91\x98\x7e\x40\xa1\x27\x5a\xb0\x3c\xce\x86\x43\x2a\xc9\x24\xc2\x16\x4c\x49\x24\xd1\x0c\x9f\x27\x34\x0a\x1d\xb5\xbb\x58\x6f\xa1\x6a\x25\x2a\x11\x2a\x88\x63\xc6\xaf\x97\x28\x04\x0d\xef\x5c\xdc\xdd\x1c\x32\x4a\x05\x16\x03\x1e\x0e\x04\x4a\x54\xdf\x43\x7d\x87\x46\x33\xcf\xa8\x53\xde\x30\x46\x1a\xf0\x88\x06\xab\xc3\xd8\xe1\x67\x0c\x12\x0d\x39\x4c\xa2\x5c\x21\x00\x3e\x2c\x88\x0a\xe6\x86\x7e\x9b\x31\xae\x0c\xb9\x02\x80\x06\xb9\xc5\x55\x0b\xa8\xf1\x93\x7a\x49\xac\x10\xd9\xca\xcf\x48\x17\x70\x00\x96\x24\x4a\xb0\x05\x35\x1d\xfb\xb5\xc2\x8c\xce\x29\xad\x5c\x1c\x3f\x44\x46\x31\x2c\x00\x70\x36\x74\xf7\x23\x6b\x52\xa4\xd7\x26\x2d\x88\x79\x28\xb7\x4c\x4d\xb4\xb9\x8a\x93\x02\x7f\xc3\x40\xb5\xd2\x96\x7b\xfa\x27\x6f\x69\x7c\x6d\x38\x45\x46\x8e\x4b\x42\xa3\x44\xe0\x1a\x9c\x35\x52\x41\xf9\xce\x3e\x24\x09\xa9\xca\xc3\x09\x99\x76\xf0\xd0\x21\xe7\xa7\x90\x42\x00\xe6\xe5\x69\x8f\x4d\xb9\x5d\x58\x80\x42\x5d\x52\x1d\x19\x3b\x8e\x52\x46\x0c\x5c\xed\x84\xd3\x89\x93\xc4\xb4\x8f\x4b\x8c\x64\xcb\xf3\xb5\xe1\xd7\xfc\x80\x24\x6a\x9e\x8b\x23\xf0\x63\x82\x52\xbd\x41\x12\xa2\x70\xc2\x18\xe1\x3a\xed\x16\x54\xb4\x19\x0a\x00\x7c\xb1\xe0\xec\x8a\x2c\x52\xeb\xf8\x5b\x84\xf2\xac\xd7\x29\x41\x2c\x97\x81\xc0\x29\xfd\x9c\x63\xfd\xa7\x3f\xc4\x05\x57\xe8\x77\x35\x8c\x6f\x46\x67\x82\x27\xb1\x05\xdf\x84\x7b\xad\x27\xcd\x60\x22\x51\x68\x37\xda\x06\xf9\x5e\xa2\xf0\x02\xce\x94\xe0\x51\x84\x05\x2b\x60\x84\x41\x1e\x2d\x11\x0f\x6e\xaf\x8c\x37\xae\x97\x4f\x7e\x8e\xac\x5d\xc9\x95\x89\xa6\xca\x64\x33\x5d\x74\x5b\x02\xb6\x01\x91\xc5\x63\x66\xcd\x8a\x36\x89\x73\xa7\xd4\x8e\x15\x2d\x8c\x02\xcb\x16\xd4\x1e\xd7\xbc\x80\x0b\xd9\x8e\x22\xfe\x09\xc3\x6b\x93\xf8\x65\xcb\x0b\x99\xcc\x57\x33\xa1\x2c\x6c\x87\xa1\x40\x29\x5b\x90\xee\xe3\x2f\x9a\xcf\x9e\xba\xb9\x2b\x54\x9f\xb8\xb8\x6d\x81\x0a\xe2\x53\x0f\xb3\x5e\x43\xea\x80\x01\x69\x41\x45\x9f\xb2\xb8\x92\x2d\xfd\x8e\xc2\x4a\xb6\x74\x23\x00\x12\x11\x19\xcb\xf8\xb0\xb5\x8a\xd4\x48\x23\xc5\x05\x99\x61\xbe\x2a\x5d\xb9\x0a\x86\x0a\xa5\x9b\xb2\x8e\xd3\x2a\x4c\xd4\x29\xaf\x02\x2c\xa7\x3d\x6d\xd3\x91\xb6\xe9\x1a\x99\x62\xfe\xaa\x00\x2b\x12\x31\x19\x2f\x97\x6c\xca\xc5\x82\xa8\x56\xf1\x2c\xd0\xcb\x21\x2e\xcd\x2c\x7c\x05\x94\x01\x89\x75\xa5\x6e\xf1\x8b\x79\x43\x53\xa1\x4c\x69\xef\x8d\x86\x38\xa3\x52\x89\xd5\x1b\xa7\x93\x96\xbb\x20\xf5\x85\x9b\xa8\xbb\x7b\xbe\xba\x5c\x06\xad\x67\xcd\x66\xd3\xb3\xd9\xc8\x1e\x11\x5c\x22\xba\x2d\x76\x03\x8a\x86\xdd\x6e\xcc\x8a\x8e\xc4\xa6\x3d\x2b\xda\x04\x00\x31\x17\xaa\x05\xc7\xcd\x93\x67\x4d\x2f\xd7\x7e\x51\x1e\xcd\x9d\xc4\xd4\x1e\x42\xdb\x62\x96\x2c\x90\xa5\x1b\x7c\x10\xf1\x24\x74\xe5\x4a\x1a\xb2\xc5\xb2\xc6\xcc\xc7\x82\x2f\x69\x88\xc2\x5e\x85\x9a\x63\x4a\x01\x39\x9d\xcd\x32\x8f\x06\x32\x9f\x45\xc2\x14\x5d\xe0\x1a\x79\x89\x4a\x51\x36\x93\xf5\xdb\x17\xda\x69\x1a\xcb\x63\x12\xc5\x73\x72\x7c\x96\x25\x79\x69\xad\xee\x4f\x48\x70\x8b\x2c\x4c\x11\xb5\x67\x3e\x2d\x01\x2c\x30\xa4\xc4\x57\xab\x18\x33\xe6\x71\x1c\xe9\xa3\x37\xe5\xac\xb1\x64\x61\xbd\xe0\x9f\xe6\x52\x6f\x92\x68\xd1\xf3\xb0\xfe\x57\xaa\x23\x88\x12\x93\xc7\xa4\x6d\xaf\xfa\xda\x09\xfc\xa9\x36\x70\x05\xa7\xf2\xfd\x41\x15\xfa\x2d\xae\xf6\xc0\xb6\x4e\x62\xbf\xf7\x06\x2d\x38\x3e\xf9\xd1\x24\xa5\xe3\xbb\xf7\xbf\x6d\x3d\xa1\x52\xd2\xdc\xd6\xac\x01\x90\xc1\x1c\xc3\x24\x4b\xf5\x16\xbc\xea\xbc\x9f\xc2\x99\xc7\x19\xf2\xcc\x2e\x47\xc9\x84\xa1\xf6\xed\x1f\x4f\xea\x4f\x4d\x22\x6d\x1c\x3f\xf7\x2c\x96\x95\xda\x58\x2d\xcb\x1d\x7d\xce\x63\xed\x32\x1d\xb7\x27\x32\x66\x37\x96\xb5\x8a\x93\x04\x01\xc6\x7a\x5a\x21\x53\xe3\x55\x8c\xb2\xb5\x8f\xdb\x3c\x29\xc2\x38\x49\x01\x26\x89\x90\xaa\x05\xcf\x9b\x4d\xcf\x95\x7f\x29\xd5\xbd\x88\x1a\xa4\x8f\xb1\x6c\xc1\x53\x43\x61\x63\x2d\x6f\x93\x49\x9a\xec\x36\x36\xc4\x62\x3f\xc1\x8e\xd8\x46\xcf\xfb\x61\xdf\xe4\xc3\x58\x50\xa6\xa0\x96\x66\xfa\x9a\x49\x90\x8a\x50\x66\x1b\x52\x34\xc0\xfa\x40\xf0\x18\x85\xa2\x28\xeb\xd7\x22\x98\xa3\x79\xea\x82\x8b\x81\xe0\xda\xb3\xcc\x95\xcc\xc8\x5d\xc9\xe8\xcc\x6a\xe9\xa7\xc9\x11\xbe\xc2\xc7\x84\x2b\x93\x55\x99\xdd\xd6\xf2\x84\xe3\xbc\xd5\x6d\x77\x6e\xd3\x09\x68\x28\x74\xae\xaa\x1f\x9f\xbc\xb0\xf6\x3c\x35\x1a\xd0\x5b\x90\xb5\x76\x1f\xd9\x4c\xcd\x5b\xf0\xd2\x33\x75\x8a\x49\xca\xbd\x81\xa3\xd2\xe9\x5d\x0c\x1d\x25\xb7\xb3\x36\xb4\xd2\x1c\xef\x81\x39\x4e\xd9\xda\x41\x60\x38\x27\xaa\x70\x8e\xe3\x4b\xe9\x4b\xc3\x21\x77\xb0\x02\xd5\x75\x27\xe3\xe5\xca\x8c\x48\x89\xea\x5f\xac\xdc\x5a\x7a\x9f\xd5\xa8\x15\x15\xad\xeb\x31\xc2\x54\xf1\xd4\xb8\x40\x35\xe7\x61\x0b\x48\xa2\xf4\xe6\x49\x43\x64\x8a\xaa\xd5\xc0\x25\x22\xa7\xb1\x88\xcf\x28\x2b\x54\xd2\x0b\x12\xc7\x94\xcd\xde\x39\xe4\x20\x22\x74\xe1\xe5\x67\x81\xb6\x4e\x5b\xd0\xbe\x30\x43\xe5\x9c\xb6\xe5\x38\x63\x28\x14\x4e\x07\xb8\x20\x34\x2a\x9e\x69\xcc\x40\xf6\x9d\x86\xc5\x39\x99\x4c\xbc\xd2\x61\xa4\x30\xa7\xbf\x67\x5f\x63\x81\x53\x14\x02\xc3\xf7\xae\xde\x2c\x42\x26\x8c\x7e\x4c\xf0\xa6\x80\x60\x33\x52\xef\xc2\xd8\xec\xaf\x94\x85\xf8\x79\xb7\xa5\xda\x89\x9a\xa7\x16\xea\xad\x6b\x12\x9a\x7f\xab\xa7\x5f\xea\x2e\x71\x5e\x94\xad\x93\x33\x1d\x61\x20\x50\xfd\x01\x8c\x2d\xe1\x4d\xb6\xb6\x12\xd1\x6e\xd5\xbb\x58\xa7\xe0\x40\xd2\xfa\xcf\x99\x31\x51\x73\x2e\xe8\xef\x58\xe5\xd1\xc6\x5f\xea\x0b\x1a\x08\x2e\xf9\x54\x71\x16\x51\xa6\x37\xbb\x85\xf3\x75\xed\xb2\x63\x64\xc4\xa8\xa0\xd6\x30\x11\x73\xd2\xc8\x48\xd6\x36\xe5\x03\x50\xfc\x16\xd9\xc3\x31\x33\xe4\xd6\x18\xf9\x10\xcc\x49\x14\x21\x9b\x15\x8f\x8d\xf7\x74\xfe\x3e\x0f\x48\x04\xa6\x45\xcc\x45\xb8\x7f\x08\x4c\xb7\xed\x71\x59\xbf\xb9\x68\xa9\x37\x63\xdb\xb3\x1e\x38\x3e\x15\x36\x73\x9b\x5d\xbb\x50\x3a\xfe\xbb\x33\x7d\x2a\xc1\x61\xbc\x2f\xff\x71\x71\x55\xa6\x26\x71\xad\x2d\x96\x0d\xbd\x23\x9f\xdb\x33\x1c\xe9\x2d\x2e\xd4\x3b\x64\xba\xc9\xba\x69\x9b\xe9\xa5\x64\xc5\x41\x1b\x1b\x72\x7b\xb1\x61\xc1\x7c\x69\xe1\x4c\x87\xce\x73\x8e\x59\x14\x41\x17\x09\x52\x8e\xf5\xf0\x9a\x18\x2f\x9e\x9f\x3a\x39\x32\x5f\xaf\x02\x7b\xd6\x6c\x7a\xb1\xe0\xbf\x61\x50\xc8\xd5\xee\xf0\x70\xc5\x43\x1c\x99\x63\x2f\x17\x2d\x30\x4f\xf2\x09\xf3\xe0\x42\xf1\x30\xd5\x08\xf8\x22\x4e\x14\xa6\x65\xb1\xc4\x20\x11\x54\xad\xf4\xe9\x33\xd0\xa6\x73\xc9\x3f\x90\xd9\xc8\x90\x18\xaf\x97\xcd\x56\xe3\x24\x9d\xec\x93\x09\x46\x72\x60\xee\xe2\x6c\xcf\xe5\x99\x3d\xae\xd3\x70\x1d\xef\xb8\x99\xfe\xf9\xc7\x2f\xd3\xbf\x86\x19\xf5\x04\x4f\x74\xd1\x9e\x2f\x45\x26\x93\x90\x2f\x08\x65\x7b\x27\xb8\x7d\xdc\x6e\xc8\x13\x7b\x6d\xa8\x67\x6d\xd6\x33\x8e\x38\x4a\xb9\x15\x7d\x47\x96\x2e\x98\x72\xd1\x16\x84\x91\x19\x86\x59\x3b\xc4\x4f\xf5\x6e\x3e\x9b\x5e\x94\x09\x2e\x3d\x1e\x47\x7c\xb5\x2d\xd2\xe2\xec\x6a\xab\xd4\x1f\xa8\xbc\x98\x02\x88\xd3\x4b\x36\x0d\xec\xf8\xee\xb8\x48\x93\xb6\x6f\x91\xd6\xdd\x95\x6d\x83\xd3\xd3\xea\xae\x41\x45\x91\x5e\xb8\xe2\x29\x76\x8d\xb2\xb5\xac\xd7\xeb\x85\x7b\x0b\xbd\xf6\xcf\x43\xdb\x7b\x92\x3d\x76\x19\xd1\xd9\x5c\x59\x07\xce\x7a\x52\x63\xba\x40\x9e\xa8\xf5\x58\x34\xcf\xdd\x14\xef\x3b\x8d\x4f\xfa\x05\xf1\xf6\x7a\x62\xa8\xdc\x86\xd9\xeb\xf9\x9e\xac\x50\xc8\xb6\x33\x3f\x6d\x01\xfc\x41\x79\x6f\xc9\xa3\x64\x51\x68\x3c\x84\x2b\x46\x16\x34\x30\xc9\x5a\xa7\x14\xca\x66\xdd\x52\xc7\xd1\xde\x5f\x6c\xb9\x27\xa8\xca\x43\xd9\x93\xdc\x6b\x7b\x8b\xdd\x2d\x46\xa5\x04\xe7\x39\xc4\x96\xe7\x9b\x3c\xa4\x77\x0e\x7b\xd4\x28\xb5\x42\x4a\x48\xba\xc2\x58\xbf\xba\x2c\xa5\x62\x64\x81\x58\xc5\xbb\x89\x74\x59\xb0\x83\xc6\xee\x35\x97\x0e\x7a\xe5\xe7\xd6\xf3\x25\xd7\x5a\x50\x5b\x1e\xd7\x9e\xe8\x51\xbd\x72\xfd\xdd\xb6\x6b\xec\x58\x2c\x30\xb4\xce\x56\x6b\xc1\x2f\xc6\xf8\x5f\x9c\x0b\xd4\xb4\xcd\x34\xfc\x15\xff\x60\xac\xf5\x5f\x9c\x19\x8b\x45\x34\x50\xb6\xdd\xfd\xed\x49\x35\xc6\x3b\xf2\xb9\x7b\x3e\xfa\xe0\x6c\x9c\xb0\xbb\xc1\x5f\x77\xba\x83\x8b\xfb\x20\x98\x8a\xfa\x82\xca\xdb\x7b\x20\xa9\x60\xde\x63\x3a\x23\xf2\xb0\x3d\x9d\x52\x46\xd5\x6a\x37\xca\x15\xd7\x1c\xf6\x5b\xf3\x6b\x64\x28\x48\x34\xc8\x15\xba\x13\x7c\xc0\xc3\x31\x8f\x50\x68\x48\xbd\x83\x8d\x09\x65\xea\x0e\x9c\xce\x1c\x83\x5b\x0d\xfc\x0e\x17\x5c\xac\x06\x3a\xbf\x25\x02\xf7\x44\xd2\x4b\xb9\x07\x8a\xd5\xeb\x39\x65\x21\x65\xb3\x6a\x78\xe2\xda\x3d\xb5\x56\x36\x06\x50\x4b\x37\x92\x54\xc3\xc5\x49\x80\x5a\x64\xb6\xd0\xcc\xdb\xb2\x71\x81\x33\xed\xaf\x85\xc1\x5f\xb3\xcf\x69\xcd\xeb\x04\x28\x88\x3a\x2c\x60\x69\xa8\x5f\x9d\x57\x53\x2e\xa8\x4e\x53\xdb\xbd\x3a\x2d\x19\x46\xb1\x40\x12\x0e\x2c\x86\x8d\x0a\x03\xf7\x09\x75\xe6\xae\xb5\xe0\x78\xa7\xae\xd6\x1d\xea\x60\x42\x7d\x24\x52\xb9\x8d\x03\x0f\x97\xe7\x9c\x44\x84\x05\x18\xa6\x17\x4c\xae\x2c\xd1\x4a\xba\x2f\x29\xed\x34\x03\x73\x40\x6c\x2f\x39\x0d\x07\x3c\x94\xbb\xc4\x32\x95\xcd\x5d\xf4\xbe\x5b\x4d\x26\x4e\x5c\xe4\x50\xce\xee\x4d\x67\xb7\xcf\x32\x45\x77\xfa\xad\x96\xe0\x77\xce\x0a\xb7\x7d\xdb\x5d\x53\xa7\xcb\x0a\xb1\x4e\x72\x5f\x75\xef\x11\xd9\x37\x53\x1e\x99\xfd\xd7\x14\x12\x70\xfc\xfc\x45\xfd\xf9\xd3\xfa\xf1\xc9\xcb\xfa\xf1\xf3\x47\x15\x4f\x82\x0b\x94\x3c\x5a\xda\x66\x69\xe5\xd3\xe0\xa5\xd6\x6a\xc5\x26\xb2\xad\xf5\x9a\xef\x23\xf9\x8e\xd5\xd1\xb0\xe9\x31\x4a\x8f\xc1\x57\x18\x29\x41\xd9\x2c\xdb\xa4\xd2\xb7\x9c\x28\x5b\xa2\x54\x74\x46\x14\x82\x9a\x23\xf8\xfe\x82\x30\x3a\x45\xa9\xfc\x44\x44\xe0\x1e\xec\x83\x98\x08\xb2\x40\x85\x02\x08\x0b\x41\x22\x02\x9d\x02\x55\xb0\xd0\x3a\xf2\x8e\x60\x8e\x51\x0c\x89\x04\xa2\x80\x44\xd1\xa6\xf4\x46\x07\x31\x0f\xa5\x7d\xd2\xa7\xf8\x60\x40\xd5\x66\x3f\xe0\xa1\xb7\x40\x45\x42\xa2\x48\xcb\xcb\xee\x67\xed\x6b\x43\x46\xed\x31\x09\xd0\x5e\xbf\xf8\xf6\x9d\x2e\x4f\xc6\x18\xb4\x5c\x73\xdb\x94\x40\xae\x08\x25\x62\x96\xdd\xc6\x7d\x75\xc6\x95\xa8\xc0\x27\xee\x4b\x1d\x2a\x9e\x43\x72\x73\xf8\x19\x83\x94\xaf\x26\xbd\x58\x90\xbc\x21\x6f\x5e\x43\x94\x73\xf7\xcd\x0f\xcc\x07\x73\xbf\x52\xaa\x1f\xec\x5d\x44\x57\x05\xa1\xb9\x99\x59\x6f\x0c\xd8\x0b\x99\x24\x8a\xdc\x4d\xba\x7b\x0a\xc2\x9e\xdb\xe9\x12\x19\x4a\x39\x10\x7c\x92\x75\x7a\xb4\x4c\x79\x13\xa3\x24\x13\x64\x17\x05\x81\x8a\x0a\x23\xbe\x1f\x10\xd3\x2c\x2f\x8c\x6d\x3c\x70\x52\x02\x4f\x7b\xf3\x95\x08\xd9\xe3\x35\x45\x94\xb4\x1f\xbf\x1d\x23\x2b\x5c\x1d\x06\xb2\x30\xe6\x7a\x13\x2d\x8c\x6e\xbd\x97\xcb\x41\xd2\x8b\x80\x39\x92\x48\xcd\xdd\x84\xce\x00\x94\x44\x17\x18\x91\x55\x56\xa6\x9f\x3e\x2b\x34\x35\x32\x2b\xa6\x27\x49\xd3\x02\xff\x9c\x3d\x29\xa1\x8f\x33\x34\xc2\x59\x56\xc0\xea\x41\x5b\xf5\xbe\x33\xc7\x97\xd4\xec\xe6\x65\xb6\x01\x51\xf3\x56\xbe\x40\x47\x23\xe7\xe4\x2e\x51\xdc\xb8\xde\xb0\xae\x59\xb4\x2a\x50\x2e\xd3\x59\x7b\x33\x6e\x8d\x96\x0e\x03\x33\xaa\x0f\x41\x94\xcd\x2e\xa8\xd8\xc4\xd1\xfa\xca\xcf\x49\x96\x8d\x15\x3f\xbd\xf2\xe4\xd2\xb2\x4b\x17\x5c\xb1\x84\xaa\x05\x6c\xc5\xdc\x10\x7a\x5d\xe4\xea\x72\x38\xcf\x05\x24\xa6\x7b\x54\xfe\xeb\xc9\xc0\xd6\x23\x56\x96\xd2\x83\x1f\x01\x5f\xc4\x9c\x21\x53\x2d\x20\x31\xcd\xd2\x46\xfe\xf9\xc0\xac\x61\x5e\x15\x4d\x8d\x66\x82\xd9\x7d\xb1\x94\x5d\xc4\x18\x75\x9d\xed\xf5\x4c\x54\x86\x14\xf1\x59\x84\x4b\x8c\xce\x4e\xab\xf2\x4b\xfe\x42\x70\x75\x62\xe9\xd8\x6b\xbc\x41\x44\x18\x3e\x4c\x6a\xd1\xd1\xf7\x3a\x7f\x78\x28\xb5\xb4\x8d\xb4\xdf\xf3\x51\x73\xe7\x6a\x8f\xe7\x2e\xa1\x06\x73\xd4\xba\x7e\x33\x1e\x0f\x46\x7b\x84\x24\x80\x5a\x3b\x52\x1f\x37\x0b\x2e\x94\x2a\x56\xc7\x0d\xbd\xa7\x94\x0d\x8d\xb4\x7a\x08\x59\x9d\x48\x5b\x65\x7d\xe8\x44\x52\x72\x9b\x52\x16\x28\xb9\xd0\xbe\x39\x65\x6b\x71\x51\x45\xb9\x74\x77\xbb\x8d\xc3\xf7\xa4\x98\xcd\xc5\x55\x2f\x6d\x1f\x22\x9b\xcb\xd9\xb5\x98\xbb\xb2\x50\xe1\x21\x97\x03\xb2\x91\xe5\x5c\x7e\x36\xe7\x8f\xc8\x35\x65\x0e\xdf\x93\x73\xa8\x54\xc8\x36\x9e\xb5\x3d\x3d\x3d\xfd\xbf\x92\x96\x4e\x77\x84\x7a\x95\xb9\xfe\x3f\x92\xff\x37\x45\x32\x9d\xc2\x2f\xbf\x40\xad\x50\x20\xd6\xe0\xec\x0c\x6a\xa5\xe7\xe0\x6b\xf0\x6b\xfe\x73\x1b\x3b\x62\x5f\xae\x58\x70\x70\xd0\x6b\xe4\xfb\x47\xfb\x66\xc0\x8c\x56\x2c\xb8\x77\xa0\x94\x44\x78\x58\x0f\xbe\xd1\xcb\x6c\xdc\xf0\x44\xfd\x5b\x1d\xec\xfe\xbe\x31\xa5\xc5\x57\xad\xd3\xf2\x34\xe2\xb3\x89\xa0\xe1\x0c\x77\x6d\x02\x19\xd0\xc1\xde\x90\xb3\x79\x00\x97\xe8\xf3\xd9\xb9\x21\x76\xa0\x5f\x14\x85\x79\x58\xe7\x30\xbf\x14\x51\xf2\x8b\x7c\xa4\x0c\xf9\x40\xa9\xaa\x22\x19\xae\x9f\x90\xee\x75\xca\xe2\xb3\x12\xfa\x92\x08\x3f\x1f\x7b\x30\x3f\xde\x74\xbe\x4d\x85\x3d\x7c\x10\xec\xa6\xe9\xad\x2b\x6d\xb7\xf4\x4e\x29\x65\x35\x99\x14\x9c\xfe\x82\x8a\x30\x3f\xa1\xf2\x57\x78\x6c\xbb\x60\x2d\xf8\x5b\xfd\xf1\xd1\x7f\x1f\x57\x9c\xe1\xd3\x5f\x51\xd9\xff\x0d\xb1\xef\xe6\x51\xf9\x73\x0f\x47\xf0\xa6\xdd\x79\xab\x0b\x80\x78\x05\x6b\x93\xa0\x38\x4c\x38\x57\x52\x09\x12\x17\xc7\x25\xb7\x3f\xff\x93\x89\xeb\x7e\x3c\x48\xe3\x43\xa2\x05\xa5\xcc\x34\xcf\x4c\x36\x3e\x02\xfb\xeb\x44\x12\x15\x7c\xa2\x51\x04\x8c\x2b\x98\x12\x1a\xb9\xd6\x99\x32\xa0\x76\xc1\x96\x84\x3d\xd0\x42\xc0\x85\xc0\x40\x45\xab\x7a\xe5\xcb\x3b\xeb\xd2\x6e\x00\x54\xc9\xee\x79\x47\x9a\x3f\xb6\xa0\xea\x57\x6b\x20\x10\x44\xce\x21\xe2\x3c\x96\x90\x30\x45\xa3\x54\x2e\x2a\x21\x89\xbd\xfc\xd7\x92\xec\xcb\x04\x95\x44\xb2\x1f\x4f\x5a\xff\x6d\xa5\x5d\xc0\xf0\x1f\xc5\x4c\x2d\x38\x57\x0d\x23\xf5\xfa\xca\xef\x76\x94\x22\x76\xc3\x2d\xfa\x7f\x02\x00\x00\xff\xff\xb7\xc0\x9a\xfd\x22\x4c\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x7f\x6f\xdb\x36\x10\xfd\x9f\x9f\xe2\xea\x0c\xcb\x56\x80\x52\xd2\x15\xc1\xa6\xa6\x05\x92\xd4\x2d\x8a\xa6\xb1\x91\x66\x05\x8a\x61\x08\x68\xf2\x6c\xb3\xa6\x48\xf5\x48\xb9\xd1\x5c\x7f\xf7\x81\xa2\xe3\x9f\x69\x92\x61\xff\xd9\xc7\x77\xef\x1e\x1f\x8f\x47\xed\x3d\xc9\x07\xda\xe6\x03\xe1\xc7\xc0\xf1\x86\xb1\x3d\xb8\xea\xbd\xee\x15\x90\x63\x90\xb9\xb2\xbe\x14\xfe\x6b\xa6\x72\x47\x7a\xa4\x2d\xaf\x2b\x1f\x08\x45\xc9\x95\xf5\x99\x74\x76\x08\xda\x83\xac\x89\xd0\x06\xd3\xc0\x58\x90\x92\x4e\xa1\x7a\x01\x3a\xb0\x3d\xa8\xc8\x0d\xc4\xc0\x34\xe0\xc7\xae\x36\xca\xee\x07\x18\x20\x63\x1f\xbb\x97\x9f\xde\x9d\x75\xaf\xaf\x3e\xf7\xbb\x2f\x13\x33\xd3\x43\xf8\x0b\xf8\x10\x3a\x6d\x61\xdf\xf8\xc8\xae\x47\xb9\x08\xae\xd4\x92\xbb\x0a\xad\x1f\xeb\x61\xe0\xd6\x29\xec\xc0\xdf\x2f\x20\x8c\xd1\x32\x00\x80\x0d\xba\x6d\x3c\x1b\x6a\x16\xc9\x9f\xc0\x88\xb0\x82\x7c\x2a\x28\x37\x7a\x90\x2b\x27\x27\x48\x69\x9b\x43\x1f\xc4\x60\x49\x58\x4e\x86\x3e\xbb\x19\xfa\xa8\x26\x57\x38\xcd\x95\xf6\x93\x5c\xfc\x53\x13\xe6\x84\xde\xd5\x24\x91\x57\x82\xc2\x21\x03\x40\x39\x76\xb0\x7f\x3f\x0c\x76\xaa\x42\xa4\x87\x11\x55\x5f\x6b\x17\x04\xc0\x01\x1c\xec\xc3\xab\x57\x2b\x31\x0c\xc0\x37\x3e\x60\x29\x83\x01\x1f\x5c\x05\x29\x33\xf3\x48\x53\x2d\x31\xca\x74\xb5\x0d\xdb\xcc\x0c\x80\xd0\x07\x47\x28\x9d\x05\x7e\xb9\xb3\x3e\x9b\x71\xd0\x43\xc0\xaf\x90\x75\x6f\x02\x89\xec\xd2\x19\x84\x8e\xb6\x43\x12\x1d\x98\xcf\x19\x80\x14\x01\x92\x94\x94\x93\x2b\x81\xa5\xb3\xd9\x17\xef\x2c\x1c\x1f\xef\x77\x7b\x6f\xf6\xd9\x8c\x01\x74\x8c\x1b\x71\x45\x7a\x8a\xd4\x29\xa0\xf3\xc5\xd5\x64\x85\x51\x1d\x36\x67\xdd\xde\x9b\xb6\x14\x5a\x95\x48\xd7\x77\x23\x28\x6c\x6f\x27\x1e\x52\xb2\xf2\xb4\xd7\xbb\xfa\x78\x75\x79\xd2\xbf\x3e\xeb\x5d\xbc\x79\xf7\xf6\xfa\xe2\xe4\x43\xf7\x65\x3c\x74\x9e\x3a\x82\xcf\x66\x1b\xda\xe7\xf3\xa5\x75\xab\xae\xf9\x69\xb6\xde\x14\xf3\xb6\x69\x18\xf3\xa8\x80\x6b\xe0\x08\x1d\xbf\xf7\xba\x7b\xfa\xe7\xdb\xeb\xf3\xde\xdb\xf3\xee\xa7\xee\xf9\xcb\x67\xdb\x81\xe7\x7b\x1d\x78\x14\x2b\x95\xc0\x69\x98\xb0\x18\xa4\xca\x9f\xa6\xdf\xa9\xb1\xf3\x52\xf8\x80\x94\x3f\x65\x6c\x20\x3c\x1e\x3d\x07\xae\xe0\xf8\xf8\x18\x66\x33\x38\x6d\x03\x5d\x1b\xaf\x0c\xfc\xf2\x59\x94\xe6\x83\x20\x3f\x16\x06\xb2\xb3\xb6\x62\x76\xe1\x14\x9e\x3a\x17\x7c\x20\x51\xbd\xaf\x07\x98\x94\xfc\x0a\xf3\xf9\xe2\x8c\x16\x55\xa2\x94\x7c\x70\x8b\xcc\x26\x4b\xe8\x43\x55\xcf\x90\xc2\x89\x3f\x6d\x02\xfa\x65\xd5\x18\xd3\x43\x2d\x45\x40\xbf\x29\xa1\x5d\xfa\x41\xf5\xf6\x8c\x96\x12\x2a\xa4\x4c\x52\x78\xa8\x7c\x9f\xf4\x54\x04\x7c\x8f\xcd\x7f\x10\xf1\x1e\x9b\x47\x6b\x98\x60\xc3\xe4\xb8\x74\x0a\x0e\x8e\x0e\x0e\xe0\x71\x19\xbb\xb0\x3b\xad\xfd\xdf\xde\x9e\x89\xfb\x0c\x95\xa2\x75\x50\x56\xbb\x72\xd2\x52\x8a\x57\x13\x9d\x4b\xc1\x03\xd5\x3e\xe4\x69\xee\xe4\xc2\xca\xb1\x23\x9f\xaf\xc6\xe6\x82\xac\xae\x94\x08\xc8\x6f\xf1\xb7\xb7\xce\x8a\x12\xe3\x5d\x44\x82\xc3\xa3\xdf\xb3\xa3\xdf\xb2\xc3\x67\x7f\x64\x87\x47\xfb\x77\xc8\x8a\xc3\xcd\x4c\xdb\xe9\xcf\xca\x89\xd2\x04\x7c\x53\xa1\x34\xae\x56\x15\xb9\xa9\x56\x48\x8c\xad\xe6\xc9\x5d\xeb\x69\x62\xa6\xb7\x64\x39\x5c\x66\x2b\xbb\x22\xb6\xbf\xc0\xc6\x18\x7c\x87\x8f\x81\xb4\x1d\xc5\xb9\x12\xc7\xcc\x9a\x86\xdb\x59\x67\xdc\x68\x40\x5a\x8d\x70\xb7\x76\xbb\x83\xca\x29\xbf\x02\x65\x8d\x28\xcd\xb2\xb6\xa8\xf4\x27\x24\xaf\x9d\x2d\x60\x7a\xc8\x26\xda\xaa\x02\xfa\x4e\xb1\x12\x83\x50\x22\x88\x82\x01\x44\xbb\x0a\x58\x95\x49\x11\x5f\x09\x89\x05\xc4\x06\xe1\x69\xde\x31\x5f\xa1\x8c\x09\xd2\xd9\x20\xb4\x45\xf2\xf1\x1f\x07\x5d\x8a\x11\x16\xb0\xb6\xd1\x73\x37\x3a\x6d\xc9\xde\xc5\x25\xf8\x0e\xf1\x69\xc0\x34\x3c\x21\xe1\xfb\xb5\x31\x7d\x67\xb4\x6c\x0a\x38\x31\xdf\x44\xe3\xdb\xb5\x5d\x31\x00\x1e\x65\x4d\x3a\x34\x67\xce\x06\xbc\x09\x45\x1b\x04\xa8\x48\x4f\xb5\xc1\x11\xaa\x02\x02\xd5\x09\x3b\x75\xa6\x2e\xf1\x43\x7c\x4e\x7c\x02\xf2\xf4\xb8\xf4\x45\x18\x17\x90\xfb\x20\x02\x2e\x08\x52\xad\x55\x64\x13\xb9\x79\xf0\xeb\x19\x69\x08\xf2\x16\xc0\xb7\x10\x84\x42\xf5\xac\x69\xd6\x24\x6d\xd2\x62\x90\x1b\x64\xab\xff\x0f\xa6\xb6\x1d\xe1\x46\x1b\xe9\x53\x41\x7c\x15\xdb\xa1\x18\x3b\x1f\x2e\x30\x7c\x73\x34\x59\xc6\x92\x45\x8b\xb3\x8b\x80\x96\xfe\xd6\xd4\xb5\x52\x1b\xcd\xb7\x6b\xd8\x0f\x93\x7f\x78\x7b\x1e\xb6\xf0\x3e\x4e\xb6\x6d\xda\xfd\xea\x17\xa6\x6c\xda\xd4\xde\xb1\x3d\xb0\x2e\x60\x01\x77\xbd\x80\x20\x29\x7e\x3d\x1a\xe7\x2a\x0f\xb5\x0d\xda\x2c\xb4\xc6\x6f\xc3\xba\x62\xab\xa7\x1f\xad\x18\x18\xbc\x93\x64\xf9\x25\xb0\xfd\xa1\x70\x1f\x18\x7e\x66\xff\x06\x00\x00\xff\xff\xae\x0f\xee\xfa\xc0\x0a\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"azuredeploy.json":  azuredeployJson,
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
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
	"azuredeploy.json":  {azuredeployJson, map[string]*bintree{}},
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
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
