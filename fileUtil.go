package fileUtil

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var (
	empty = ``
)

// FileExists returns whether a file exists
func FileExists(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

// DirExists returns whether a directory exists
func DirExists(dirPath string) bool {
	f, err := os.Stat(dirPath)
	if err != nil {
		return false
	}

	return f.IsDir()
}

// GetFileSize returns file size , support http(s)
func GetFileSize(uri string) (fileSize int64, err error) {
	urlStruct, err := url.Parse(uri)
	if err != nil {
		return
	}
	// local file
	if len(urlStruct.Scheme) == 0 {
		f, Ferr := os.Stat(uri)
		if Ferr != nil {
			return 0, Ferr
		}
		return f.Size(), nil
	}
	// http(s)
	resp, err := http.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("Connection Error")
	}

	if resp.ContentLength <= 0 {
		fileSize, _ = strconv.ParseInt(resp.Header["Accept-Length"][0], 10, 64)
	} else {
		fileSize = resp.ContentLength
	}

	return
}

// GetFileName returns file name
func GetFileName(filePath string) string {
	return path.Base(filePath)
}

// GetExt returns extension name
// Will return `` when giving a string `.foo` or `.foo.bar.` etc
func GetExt(filePath string) string {
	if len(filePath) == 0 {
		return empty
	}
	if -1 == strings.Index(filePath, `.`) {
		return empty
	}
	if ok, _ := regexp.MatchString(`^\.[^\.]*$`, filePath); ok {
		return empty
	}
	if string(filePath[len(filePath)-1]) == `.` {
		return empty
	}
	return path.Ext(filePath)
}

// ReadAll returns file content,will return `` if err
func ReadAll(filePath string) string {
	f, err := os.Stat(filePath)
	if err != nil {
		return empty
	}
	if f.IsDir() {
		return empty
	}
	fo, err := os.Open(filePath)
	if err != nil {
		return empty
	}
	defer fo.Close()
	fd, err := ioutil.ReadAll(fo)
	if err != nil {
		return empty
	}
	return string(fd)
}

// ReadAllOk returns file content with err
func ReadAllOk(filePath string) (content string, err error) {
	f, err := os.Stat(filePath)
	if err != nil {
		return
	}
	if f.IsDir() {
		return empty, errors.New("not a file")
	}
	fo, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer fo.Close()
	fd, err := ioutil.ReadAll(fo)
	if err != nil {
		return
	}
	return string(fd), nil
}

// Truncate warps os.Truncate
func Truncate(path string) (err error) {
	err = os.Truncate(path, 0)
	if err != nil {
		return
	}
	return
}

// MkdirAll warps os.MkdirAll
func MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// WriteString warps os.file.WriteString
func WriteString(path, s string, perm os.FileMode) (n int, err error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	defer f.Close()
	if err != nil {
		return
	}
	n, err = f.WriteString(s)
	if err != nil {
		return
	}
	return
}

// CopyFile forked from https://github.com/koding/file/blob/master/file.go#L90
func CopyFile(src, dst string) (err error) {
	sf, err := os.Open(src)
	if err != nil {
		return
	}
	defer sf.Close()

	fi, err := sf.Stat()
	if err != nil {
		return
	}

	if fi.IsDir() {
		return errors.New("src is a directory, please provide a file")
	}

	df, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fi.Mode())
	if err != nil {
		return err
	}
	defer df.Close()

	if _, err := io.Copy(df, sf); err != nil {
		return err
	}

	return nil
}

//Md5File returns md5 of a local file
func Md5File(filePath string) (md5String string, err error) {
	if FileExists(filePath) {
		f, err := os.Open(filePath)
		if err != nil {
			return ``, err
		}
		md5Object := md5.New()
		io.Copy(md5Object, f)
		defer f.Close()
		md5String = fmt.Sprintf("%x", md5Object.Sum(nil))
		return md5String, nil
	}
	return ``, errors.New("file does not exist")
}
