package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 60, // 请求超时时间
}

type HttpError struct {
	Err error
}

func (he *HttpError) Error() string {
	return he.Err.Error()
}

// Handler 请求成功回调处理器
type Handler func(resp *http.Response)

// ErrHandler 请求失败回调处理器
type ErrHandler func(err *HttpError)

// ContentType 是请求头的 Content-Type 的封装类型，如 ContentTypeApplicationJson 等
type ContentType string

type builder struct {
	url         string
	params      []any
	method      string
	contentType ContentType
	callback    Handler
	errHandler  ErrHandler
	body        io.Reader
	err         error
	resp        *http.Response
}

func NewBuilder(url string, param ...any) *builder {
	return &builder{url: url, params: param}
}

func (b *builder) ContentType(contentType ContentType) *builder {
	b.contentType = contentType
	return b
}

func (b *builder) Body(body io.Reader) *builder {
	b.body = body
	return b
}

func (b *builder) BodyStr(body string) *builder {
	b.body = strings.NewReader(body)
	return b
}

func (b *builder) Get() *builder {
	b.method = http.MethodGet
	resp, err := client.Get(b.url)
	b.resp = resp
	if err == nil && resp != nil && resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
	}
	b.err = err
	return b
}

func (b *builder) Post() *builder {
	b.method = http.MethodPost
	resp, err := client.Post(b.url, string(b.contentType), b.body)
	b.resp = resp
	if err == nil && resp != nil && resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
	}
	b.err = err
	return b
}

func (b *builder) WhenSuccess(handler Handler) *builder {
	if b.err == nil && b.resp.StatusCode == http.StatusOK {
		handler(b.resp)
	}
	return b
}

func (b *builder) WhenFailed(handler ErrHandler) *builder {
	if b.err != nil {
		handler(&HttpError{b.err})
	}
	return b
}

func (b *builder) End() {
	if b.resp != nil && b.resp.Body != nil {
		b.resp.Body.Close()
		b.resp.Body = nil
	}
	b.resp = nil
	b = nil
}

// convenient GET methods

func MustGetString(url string) string {
	return GetString(url, func(err *HttpError) {
		panic(err)
	})
}

func GetString(url string, errHandler ErrHandler) string {
	var s string
	NewBuilder(url).Get().WhenSuccess(func(resp *http.Response) {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read response data error: %v", err))
		}
		s = string(bs)
	}).WhenFailed(errHandler).End()
	return s
}

func MustGetBytes(url string) []byte {
	return GetBytes(url, func(err *HttpError) {
		panic(fmt.Sprintf("request failed: %v", err))
	})
}

func GetBytes(url string, errHandler ErrHandler) []byte {
	var ret []byte
	NewBuilder(url).Get().WhenSuccess(func(resp *http.Response) {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read reponse data error: %v", err))
		}
		ret = bytes
	}).WhenFailed(errHandler).End()
	return ret
}

func MustGet(url string, handler Handler) {
	Get(url, handler, func(err *HttpError) {
		panic(err)
	})
}

func Get(url string, handler Handler, errHandler ErrHandler) {
	NewBuilder(url).Get().WhenSuccess(handler).WhenFailed(errHandler).End()
}

func MustGetJsonObject[T any](url string, t T) T {
	GetJsonObject(url, func(err *HttpError) {
		panic(fmt.Sprintf("request failed: %v", err))
	}, t)
	return t
}

func GetJsonObject[T any](url string, errHandler ErrHandler, t T) T {
	NewBuilder(url).Get().WhenSuccess(func(resp *http.Response) {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read response data error: %v", err))
		}
		err = json.Unmarshal(bs, t)
		if err != nil {
			panic(fmt.Sprintf("unmarshal error: %v", err))
		}
	}).WhenFailed(errHandler).End()
	return t
}

// convenient POST methods

func Post(url string, contentType string, body io.Reader) []byte {
	resp, err := client.Post(url, contentType, body)
	if err != nil {
		panic(fmt.Sprintf("request error: %v", err))
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("read response error: %v", err))
	}
	return b
}
