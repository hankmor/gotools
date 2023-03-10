package str_test

import (
	"github.com/hankmor/gotools/assert"
	"github.com/hankmor/gotools/str"
	"github.com/hankmor/gotools/testool"
	"testing"
)

func TestBlurEmail(t *testing.T) {
	type Case struct {
		email  string
		expect string
	}
	cases := []Case{
		{"1313831783@qq.com", "13****3@qq.com"},
		{"belonk@126.com", "be****k@126.com"},
	}
	for _, c := range cases {
		dst := str.BlurEmail(c.email)
		if dst != c.expect {
			t.Errorf("test failed, expect: %v, but found: %v", c.expect, dst)
		}
	}
}

func TestEndsWith(t *testing.T) {
	assert.True(str.EndsWith("", ""))
	assert.True(str.EndsWith("a", ""))
	assert.True(!str.EndsWith("", "a"))

	s := "aaabb123b"
	assert.True(str.EndsWith(s, "b"))
	assert.True(str.EndsWith(s, "3b"))
	assert.True(str.EndsWith(s, "23b"))
	assert.True(str.EndsWith(s, "123b"))
	assert.True(!str.EndsWith(s, "a"))

	assert.True(str.StartsWith("", ""))
	assert.True(str.StartsWith("a", ""))
	assert.True(!str.StartsWith("", "a"))

	assert.True(str.StartsWith(s, "a"))
	assert.True(str.StartsWith(s, "aa"))
	assert.True(str.StartsWith(s, "aaa"))
	assert.True(str.StartsWith(s, "aaab"))
	assert.True(!str.StartsWith(s, "aaab1"))
	assert.True(!str.StartsWith(s, "1aaab1"))
}

func TestCamelCaseToUnderscore(t *testing.T) {
	cs := [][]string{
		{"HelloWorld", "hello_world"},
		{"helloWorld", "hello_world"},
		{"Helloworld", "helloworld"},
		{"AbcDEFGh", "abc_def_gh"},
		{"AbcDefGh", "abc_def_gh"},
		{"abcDefGh", "abc_def_gh"},
		{"abcDefGh😄", "abc_def_gh😄"},
	}

	tl := testool.Wrap(t)
	tl.Case("camelcase to underscore")

	for _, c := range cs {
		r := str.CamelCaseToUnderscore(c[0])
		tl.Require(r == c[1], "expect result is: %v, but is: %v", c[1], r)
	}
}

func TestUnderscoreToCamelCase(t *testing.T) {
	cs := [][]string{
		{"HelloWorld", "hello_world"},
		{"HelloWorld", "hello_world"},
		{"Helloworld", "helloworld"},
		{"AbcDefGh", "abc_def_gh"},
		{"AbcDefGh中文", "abc_def_gh中文"},
	}

	tl := testool.Wrap(t)
	tl.Case("camelcase to underscore")

	for _, c := range cs {
		r := str.UnderscoreToCamelCase(c[1])
		tl.Require(r == c[0], "expect result is: %v, but is: %v", c[0], r)
	}
}
