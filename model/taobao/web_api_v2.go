package taobao

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	DEFAULT_CHARSET = "UTF-8"
)

func DoPost(url string, timeout time.Duration, params, header map[string]string) (*HttpResponse, error) {
	query := buildQuery(params)
	return _doPost(url, "application/x-www-form-urlencoded;charset=UTF-8", query, timeout, header)
}

func _doPost(url, ctype, content string, timeout time.Duration, headerMap map[string]string) (*HttpResponse, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetBodyString(content)
	req.SetRequestURI(url)
	req.Header.SetMethod(METHOD_POST)

	return _doTimeout(req, resp, headerMap, ctype, timeout)
}

func _doTimeout(req *fasthttp.Request, resp *fasthttp.Response, headerMap map[string]string, ctype string, timeout time.Duration) (*HttpResponse, error) {
	if len(headerMap) == 0 || headerMap[TOP_HTTP_DNS_HOST] == "" {
		req.Header.Set("Host", string(req.URI().Host()))
	} else {
		req.Header.Set("Host", headerMap[TOP_HTTP_DNS_HOST])
	}

	req.Header.Set("Accept", "text/xml,text/javascript")
	req.Header.Set("User-Agent", "top-sdk-java")
	req.Header.Set("Content-Type", ctype)
	for k, v := range headerMap {
		if !(TOP_HTTP_DNS_HOST == k) {
			req.Header.Set(k, v)
		}
	}

	if err := fasthttp.DoTimeout(req, resp, timeout); err != nil {
		return nil, err
	}

	head := make(map[string]string)
	resp.Header.VisitAll(func(key, value []byte) {
		head[string(key)] = string(value)
	})

	hResp := &HttpResponse{
		Code:   resp.StatusCode(),
		Body:   resp.Body(),
		Header: head,
	}

	debugPrint(func() (head []byte, body []byte) {
		return []byte(req.Header.String()), req.Body()
	}, false)

	debugPrint(func() (head []byte, body []byte) {
		return []byte(resp.Header.String()), hResp.Body
	}, true)

	return hResp, nil
}

func buildQuery(params map[string]string) string {
	var pkv []string
	for k, v := range params {
		if k != "" && v != "" {
			pkv = append(pkv, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
		}
	}
	return strings.Join(pkv, "&")
}

func DoPostWithJson(url string, params map[string]interface{}, timeout time.Duration) (*HttpResponse, error) {
	ctype := "application/json;charset=UTF-8"
	bs, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return _doPost(url, ctype, string(bs), timeout, nil)
}

func DoGet(url string, params map[string]string, timeout time.Duration) (*HttpResponse, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if len(params) > 0 {
		if query := buildQuery(params); strings.Contains(url, "?") {
			url += query
		} else {
			url = url + "?" + query
		}
	}
	req.SetRequestURI(url)
	req.Header.SetMethod(METHOD_GET)

	return _doTimeout(req, resp, nil, "application/x-www-form-urlencoded", timeout)
}

var Debug = false
var maxDebugBody = 3 * 1024 * 1024

func debugPrint(f func() (head, body []byte), isResp bool) {
	if !Debug {
		return
	}
	h, b := f()

	pre := ">"
	if isResp {
		pre = "<"
	}

	sc := bufio.NewScanner(bytes.NewReader(h))
	for sc.Scan() {
		println(pre, sc.Text())
	}

	if len(b) > maxDebugBody {
		b = b[:maxDebugBody]
		b = append(b, []byte("...")...)
	}
	println(string(b))
}
