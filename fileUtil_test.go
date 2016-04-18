package fileUtil

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetFileSize(t *testing.T) {
	path := "http://dl.360safe.com/setup.exe"
	size, err := GetFileSize(path)
	if err != nil {
		t.Error("url size error")
	} else {
		t.Log(fmt.Sprintf("size = %d", size))
		t.Log("pass")
	}
	path = "http:adasd/adas/d/asd"
	size, err = GetFileSize(path)
	if err != nil {
		t.Log(err)
		t.Log(fmt.Sprintf("size = %d", size))
		t.Log("pass")
	} else {
		t.Error("net work error")
	}
}

func TestFileExists(t *testing.T) {
	filePath := "/tmp/foo/bar"
	rs := FileExists(filePath)
	os := runtime.GOOS

	if !rs {
		t.Log("pass")
	} else {
		t.Error("net work error")
	}

	filePath = "/etc/bashrc"
	successFlag := true
	if os == "windows" {
		successFlag = false
	}
	rs = FileExists(filePath)
	if rs == successFlag {
		t.Log("pass")
	} else {
		t.Errorf("file error in %s", filePath)
	}
}

func TestGetFileName(t *testing.T) {
	if "asd.zip" == GetFileName("/sdf/sd/fsd/fds/f/asd.zip") {
		t.Log("pass")
	} else {
		t.Error("error in GetFileName")
	}
}

func TestGetExt(t *testing.T) {
	filename := `.asdasdasd`
	if len(GetExt(filename)) == 0 {
		t.Log(filename + " pass")
	} else {
		t.Errorf("ext error in %s", filename)
	}

	filename = ``
	if len(GetExt(filename)) == 0 {
		t.Log(filename + " pass")
	} else {
		t.Errorf("ext error in %s", filename)
	}

	filename = `foobar`
	if len(GetExt(filename)) == 0 {
		t.Log(filename + " pass")
	} else {
		t.Errorf("ext error in %s", filename)
	}

	filename = `.asdasdasd.`
	if len(GetExt(filename)) == 0 {
		t.Log(filename + " pass")
	} else {
		t.Errorf("ext error in %s", filename)
	}

	filename = `http://dl.360safe.com/setup.exe`
	if GetExt(filename) == `.exe` {
		t.Log(filename + " pass")
	} else {
		t.Errorf("ext error in %s", filename)
	}

	filename = `/foo/.bar/foo.bar`
	if GetExt(filename) == `.bar` {
		t.Log(filename + " pass")
	} else {
		t.Errorf("ext error in %s", filename)
	}
}

func TestReadAll(t *testing.T) {
	filePath := "/tmp/foobar.barfoo"
	rs := ReadAll(filePath)

	if len(rs) == 0 {
		t.Log("pass")
	} else {
		t.Error(filePath + " err")
	}

	filePath = "/etc/bashrc"
	rs = ReadAll(filePath)
	os := runtime.GOOS
	if os == "windows" {
		if len(rs) == 0 {
			t.Log("pass")
		} else {
			t.Error(filePath + " err")
		}
	} else {
		if len(rs) > 0 {
			t.Log("pass")
		} else {
			t.Error(filePath + " err")
		}
	}
}

func TestReadAllOk(t *testing.T) {
	os := runtime.GOOS
	filePath := "/tmp/foobar.barfoo"
	rs, err := ReadAllOk(filePath)

	if len(rs) == 0 && err != nil {
		t.Log("pass")
	} else {
		t.Error(filePath + " err")
	}

	filePath = "/etc/"
	rs, err = ReadAllOk(filePath)

	if os == "windows" {
		if len(rs) == 0 && err != nil {
			t.Log("pass")
		} else {
			t.Error(filePath + " err")
		}
	} else {
		if len(rs) == 0 && err != nil {
			t.Log("pass")
		} else {
			t.Error(filePath + " err")
		}
	}

	filePath = "/etc/bashrc"
	rs, err = ReadAllOk(filePath)
	if os == "windows" {
		if len(rs) == 0 && err != nil {
			t.Log("pass")
		} else {
			t.Error(filePath + " err")
		}
	} else {
		if len(rs) > 0 && err == nil {
			t.Log("pass")
		} else {
			t.Error(filePath + " err")
		}
	}
}

func TestTruncate(t *testing.T) {
	filePath := "/tmp/foo"
	WriteString(filePath, "123123", 0666)
	Truncate(filePath)
	if len(ReadAll(filePath)) == 0 {
		t.Log("pass")
	} else {
		t.Error(filePath + " err")
	}
}

func TestMd5File(t *testing.T) {
	filePath := "/tmp/foo.bar"
	WriteString(filePath, "foobar", 0666)

	if rs, err := Md5File(filePath); rs == `3858f62230ac3c915f300c664312c63f` && err == nil {
		t.Log("pass")
	} else {
		t.Log(rs)
		t.Error(filePath + " err")
	}
}
