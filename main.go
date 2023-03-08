package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	targetUrl := "https://api.openai.com/" // 目标域名和端口
	target, err := url.Parse(targetUrl)
	if err != nil {
		log.Fatal(err)
	}

	// 创建反向代理
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 修改请求头，将Host设置为目标域名
	proxy.Director = func(req *http.Request) {
		req.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}

	// 打印HTTP请求和响应的日志
	proxy.ModifyResponse = func(resp *http.Response) error {
		// 打印HTTP请求的日志
		requestDump, err := httputil.DumpRequest(resp.Request, true)
		if err != nil {
			log.Printf("Failed to dump request: \n%v\n", err)
		} else {
			log.Printf("%s Request: %s\n", time.Now().Format("2006-01-02 15:04:05"), string(requestDump))
		}

		// 打印HTTP响应的日志
		responseDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Printf("Failed to dump response: %v\n", err)
		} else {
			log.Printf("%s Response: \n%s\n", time.Now().Format("2006-01-02 15:04:05"), string(responseDump))
		}

		return nil
	}

	// 设置日志前缀和输出位置
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 启动HTTP服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Printf("Starting server on port 9000...\n")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}
