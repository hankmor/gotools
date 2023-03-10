package testool_test

import (
	"github.com/hankmor/gotools/testool"
	"testing"
)

func TestWrap(t *testing.T) {
	tr := testool.Wrap(t)
	tr.Case("wrapping testing.T")
	tr.Require(tr != nil, "wrapping should be success")
}
