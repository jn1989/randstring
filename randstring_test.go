package randstring

import (
	"regexp"
	"testing"
)

func TestRandString(t *testing.T) {
	reg := regexp.MustCompile(`[:alnum:]`) //正则匹配，生成的只有字母和数字
	s := RandString(20)
	if !reg.Match([]byte(s)) {
		t.Error("测试失败:", s)
	}
}
func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(20)
	}
}
