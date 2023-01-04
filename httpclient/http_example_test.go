package httpclient_test

import (
	"encoding/json"
	"fmt"
	"gotools/httpclient"
	"io/ioutil"
	"log"
	"net/http"
)

// 运行测试时会执行 example 代码，示例方法必须使用 fmt.println 输出，最后约定用 Output 输出一致的结果
// 打印结果，如后边的 Output 输出结果必须一致，否则测试时会失败

func ExampleBuilderGet() {
	// 启动示例 http 服务
	startServer()

	var ret string
	// 构建 get 请求
	httpclient.NewBuilder("http://localhost:1234/html").Get().WhenSuccess(func(resp *http.Response) { // 请求成功处理
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("response body should be readable but not: %v", err)
		}
		ret = string(body)
	}).WhenFailed(func(err *httpclient.HttpError) { // 请求失败处理
		panic(err)
	}).End() // 请求完成，清理资源

	fmt.Println(ret)
	// Output:
	// <html><head>test</head><body><h1>test page!</h1></body></html>
}

func ExampleGetString() {
	startServer()

	s := httpclient.GetString("http://localhost:1234/json", func(err *httpclient.HttpError) {
		if err != nil {
			log.Fatalf("should has no error but found: %v", err)
		}
	})
	fmt.Println(s)
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleMustGetString() {
	startServer()

	s := httpclient.MustGetString("http://localhost:1234/json")
	fmt.Println(s)
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleMustGet() {
	startServer()

	var s string
	httpclient.MustGet("http://localhost:1234/json", func(resp *http.Response) {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		s = string(bs)
	})
	fmt.Println(s)
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleGet() {
	startServer()

	var s string
	httpclient.Get("http://localhost:1234/json", func(resp *http.Response) {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		s = string(bs)
	}, func(err *httpclient.HttpError) {
		panic(err)
	})

	fmt.Println(s)
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleGetBytes() {
	startServer()

	bs := httpclient.GetBytes("http://localhost:1234/json", func(err *httpclient.HttpError) {
		panic(err)
	})
	fmt.Println(string(bs))
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleMustGetBytes() {
	startServer()

	bs := httpclient.MustGetBytes("http://localhost:1234/json")
	fmt.Println(string(bs))
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

type user struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float32 `json:"height"`
}

func ExampleGetJsonObject() {
	startServer()

	u := httpclient.GetJsonObject("http://localhost:1234/json", func(err *httpclient.HttpError) {
		panic(err)
	}, &user{})

	bs, _ := json.Marshal(u)
	fmt.Println(string(bs))
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleMustGetJsonObject() {
	startServer()

	u := httpclient.MustGetJsonObject("http://localhost:1234/json", &user{})

	bs, _ := json.Marshal(u)
	fmt.Println(string(bs))
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}
