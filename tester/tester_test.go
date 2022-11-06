package tester_test

import (
	"github.com/huzhouv/gotools/tester"
	"testing"
)

func TestWrap(t *testing.T) {
	tr := tester.Wrap(t)
	tr.Case("wrapping testing.T")
	tr.Require(tr != nil, "wrapping should be success")
}
