package jsoner

import (
	"bytes"
	"encoding/json"
)

type MarshalOption func(bs []byte) []byte
type UnmarshalOption func(d *json.Decoder)

func Indent(bs []byte) []byte {
	var buf bytes.Buffer
	err := json.Indent(&buf, bs, "", "  ")
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func UseNumber(d *json.Decoder) {
	d.UseNumber()
}

func DisallowUnknownFields(d *json.Decoder) {
	d.DisallowUnknownFields()
}

func ToJson(t any, options ...MarshalOption) string {
	bs, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	for _, option := range options {
		bs = option(bs)
	}
	return string(bs)
}

func Parse(str string, t any, options ...UnmarshalOption) any {
	d := json.NewDecoder(bytes.NewReader([]byte(str)))
	for _, option := range options {
		option(d)
	}
	err := d.Decode(t)
	if err != nil {
		panic(err)
	}
	return t
}
