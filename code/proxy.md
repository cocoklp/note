```
package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const pulltimeout = 5000

func main() {
	http.HandleFunc("/", proxyServer)
	http.ListenAndServe(":8080", nil)
}

func proxyServer(rw http.ResponseWriter, req *http.Request) {
	realserver := "http://114.67.72.40:5555"
	fmt.Println(req.URL.String())
	rawPath := req.URL.String()
	oriH := req.Header

	//body := []byte("")
	body, _ := ioutil.ReadAll(req.Body)
	urlNew := fmt.Sprintf("%s%s", realserver, rawPath)

	fmt.Printf("urlRawPath %s\n", rawPath)
	fmt.Printf("url %s\n", urlNew)
	fmt.Printf("body %s\n", string(body))

	user := "admin"
	pass := "123456"
	encoded := base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))

	outReq, err := http.NewRequest(req.Method, urlNew, bytes.NewReader(body))
	outReq.Header.Add("User-Agent", "Mozilla/5.0")
	for k, hs := range oriH {
		for _, v := range hs {
			outReq.Header.Add(k, v)
		}
	}
	outReq.Header.Set("Authorization", "Basic "+encoded)
	outReq.Header.Set("Content-Type", "application/json")

	timeout := time.Duration(time.Duration(pulltimeout) * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	outReq = outReq.WithContext(ctx)
	err = httpDo(ctx, outReq, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		io.Copy(rw, resp.Body)
		return nil
	})
	fmt.Println(err)
}

/*
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
*/
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

```

