package ini

import (
	"os"
	"reflect"
	"testing"
)

const (
	NON_ERROR = "open /usr/local/var/go/src/github.com/goindow/ini/non_exists.ini: no such file or directory"
)

var (
	pwd, _        = os.Getwd()
	file          = pwd + "/test.ini"
	nonExistsFile = pwd + "/non_exists.ini"
	sample        = Conf{
		"default": Section{
			"a": "1",
			"b": "2",
			"c": "1 + 1 = 2",
		},
		"user": Section{
			"name":   "hyb",
			"age":    "99",
			"gender": "male",
		},
		"profile": Section{
			"github": "github.com/goindow",
			"email":  "76788424@qq.com",
		},
	}
)

func Test_Read_Err(t *testing.T) {
	if _, err := Read(nonExistsFile); err == nil {
		fail(t, "ini.Read() should return an error("+NON_ERROR+")")
	}
}

func Test_Read_Nil(t *testing.T) {
	content, err := Read(file)
	if err != nil {
		fail(t, "ini.Read() should not return an error")
	}
	if reflect.TypeOf(sample) != reflect.TypeOf(content) {
		fail(t, "ini.Read() should return an ini.Conf")
	}
	if len(sample) != len(content) {
		fail(t, "ini.Read() return an wrong result with different length")
	}
	for ck, cv := range content {
		sv, ok := sample[ck]
		// content 中的 ck 在 sample 中不存在
		if !ok {
			fail(t, "ini.Read() return an wrong result with different sectionName")
			return
		}
		// 比较 section（sv、cv）
		if reflect.TypeOf(sv) != reflect.TypeOf(cv) {
			fail(t, "ini.Read() should return an ini.Conf with ini.Section")
			return
		}
		if len(sv) != len(cv) {
			fail(t, "ini.Read() return an wrong result with different length section")
		}
		// 比较 section 内容
		for cvk, cvv := range cv {
			svv, is := sv[cvk]
			// cv 中的 cvk 在 sv 中不存在
			if !is {
				fail(t, "ini.Read() return an wrong result with different section")
				return
			}
			// 比较 section 的 value
			if cvv != svv {
				fail(t, "ini.Read() return an wrong result")
				return
			}
		}
	}
}

func Benchmark_Read(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Read(file)
	}
}

func fail(t *testing.T, s string) {
	echo(t, s, 1)
}

func ok(t *testing.T, s string) {
	echo(t, s, 2)
}

func echo(t *testing.T, s string, level uint) {
	switch level {
	case 1:
		t.Error("[fail] " + s)
	case 2:
		t.Log("[ok] " + s)
	}
}
