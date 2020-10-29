package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const pulltimeout = 5000
const pulltimes = 2

func main() {
	url := os.Args[1]
	fmt.Println(time.Now())
	//rbody, e, _, hs := tryGetRespBodyWithHeaders("", url, pulltimeout, pulltimes, nil)
	code, _, hs, err := DoRequest("GET", url, "", nil, nil, pulltimeout)
	fmt.Println(time.Now())
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(strings.HasSuffix(err.Error(), "connection timed out"))
	}

	headers := http.Header(hs)
	fmt.Println(code, err, hs, headers)
}

func tryGetRespBodyWithHeaders(host string, queryAddr string, timeout int64, tryTimes int,
	inHeaders map[string]string) ([]byte, error, int, map[string][]string) {
	var try int
	var rbody []byte
	var e error
	var headers map[string][]string
	for try < tryTimes {
		try++
		statusCode, body, hs, e := DoRequest("GET", queryAddr, host, inHeaders, nil, timeout)
		if e == nil && statusCode >= http.StatusOK && statusCode < http.StatusBadRequest {
			rbody = body
			headers = hs
			break
		}
	}
	fmt.Printf("pulled webcache page[%s], length of body(bytes)[%d], error %v", queryAddr, len(rbody), e)
	return rbody, e, try, headers
}

// reqType is one of HTTP request strings (GET, POST, PUT, DELETE, etc.)
func DoRequest(reqType string, url string, host string, headers map[string]string, data []byte, timeoutMs int64) (int, []byte, map[string][]string, error) {
	var reader io.Reader
	if data != nil && len(data) > 0 {
		reader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(reqType, url, reader)
	if err != nil {
		return 0, nil, nil, err
	}

	if host != "" {
		req.Host = host
	}
	// I strongly advise setting user agent as some servers ignore request without it
	req.Header.Set("User-Agent", "Mozilla/5.0")
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	var (
		statusCode int
		body       []byte
		timeout    time.Duration
		ctx        context.Context
		cancel     context.CancelFunc
		header     map[string][]string
	)
	timeout = time.Duration(time.Duration(timeoutMs) * time.Millisecond)
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req = req.WithContext(ctx)
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		body, _ = ioutil.ReadAll(resp.Body)
		statusCode = resp.StatusCode
		header = resp.Header

		return nil
	})

	return statusCode, body, header, err
}

// httpDo issues the HTTP request and calls f with the response. If ctx.Done is
// closed while the request or f is running, httpDo cancels the request, waits
// for f to exit, and returns ctx.Err. Otherwise, httpDo returns f's error.
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass thehe response to f.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	/*
		resp, err := client.Do(req)
				fmt.Println(resp, err)
						return err
	*/

	c := make(chan error, 1)
	go func() { c <- f(client.Do(req)) }()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}

}
