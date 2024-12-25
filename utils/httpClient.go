package utils

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	client  *resty.Client
	reqLog  func(...interface{})
	respLog func(...interface{})
}

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

type Option func(hc *HttpClient)
type ReqOption func(req *resty.Request)

func NewHttpClient(opts ...Option) *HttpClient {
	instance := &HttpClient{}
	// default client
	instance.client = resty.NewWithClient(
		&http.Client{Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: false},
			DisableKeepAlives: true},
			Timeout: 10 * time.Second})
	instance.reqLog = defaultLog
	instance.respLog = defaultLog
	instance.RegisterMiddleware()

	for _, opt := range opts {
		opt(instance)
	}

	return instance
}

// defaultLog 默认日志输出
func defaultLog(v ...interface{}) {
	log.Println(v...)
}

// SetRestyClient 设置resty client
func SetRestyClient(client *resty.Client) Option {
	return func(hc *HttpClient) {
		hc.client = client
	}
}

// SetReqLog 设置请求日志
func SetReqLog(l func(...interface{})) Option {
	return func(hc *HttpClient) {
		hc.reqLog = l
	}
}

// SetRespLog 设置相应日志
func SetRespLog(l func(...interface{})) Option {
	return func(hc *HttpClient) {
		hc.respLog = l
	}
}

// RegisterMiddleware 注册resty中间件
func (hc *HttpClient) RegisterMiddleware() {
	// Registering Request Middleware
	hc.client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36")
		if hc.client != nil {
			hc.reqLog(fmt.Sprintf("info:request\turl:%s\tquery_param:%+v\tform_data:%+v\tbody:%+v\theader=%+v", req.URL, req.QueryParam, req.FormData, req.Body, req.Header))
		}

		return nil // if its success otherwise return error
	})

	// Registering Response Middleware
	hc.client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		if hc.respLog != nil {
			hc.respLog(fmt.Sprintf("info:response\turl:%s\tquery_param:%+v\tform_data:%+v\tbody:%+v\tresponse:%s\theader=%+v", resp.Request.URL, resp.Request.QueryParam, resp.Request.FormData, resp.Request.Body, string(resp.Body()), resp.Request.Header))
		}

		return nil
	})
}

// Do 发送请求
// endpoint: host地址，如：http(s)://aaa.bbb.com
// uri: path，如：/user
func (hc *HttpClient) Do(endpoint, method, uri string, opts ...ReqOption) (*resty.Response, error) {
	req := hc.client.R()

	for _, opt := range opts {
		opt(req)
	}

	u := fmt.Sprintf("%s%s", endpoint, uri)
	switch method {
	case MethodGet:
		return req.Get(u)
	case MethodPost:
		return req.Post(u)
	}
	return nil, errors.New("method error")
}

// SetReqHeader 设置请求头
func SetReqHeader(header, value string) ReqOption {
	return func(req *resty.Request) {
		req.SetHeader(header, value)
	}
}

// SetReqQueryString 设置URL参数
func SetReqQueryString(queryString string) ReqOption {
	return func(req *resty.Request) {
		req.SetQueryString(queryString)
	}
}

// SetReqQueryParam 设置URL参数
func SetReqQueryParam(param, value string) ReqOption {
	return func(req *resty.Request) {
		req.SetQueryParam(param, value)
	}
}

// SetReqQueryParams 设置URL参数
func SetReqQueryParamsMap(data map[string]string) ReqOption {
	return func(req *resty.Request) {
		for k, v := range data {
			req.SetQueryParam(k, v)
		}
	}
}

// SetReqQueryParams 设置URL参数
func SetReqQueryParams(params map[string]string) ReqOption {
	return func(req *resty.Request) {
		req.SetQueryParams(params)
	}
}

// SetReqBody 设置Body
func SetReqBody(body interface{}) ReqOption {
	return func(req *resty.Request) {
		req.SetBody(body)
	}
}

// SetReqFormData 设置Form data
func SetReqFormData(data map[string]string) ReqOption {
	return func(req *resty.Request) {
		req.SetFormData(data)
	}
}
