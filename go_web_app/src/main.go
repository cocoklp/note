package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type MyHandler struct{}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	//paths := string(r.URL.Path)
	for _, v := range r.URL.Path {
		log.Println(string(v))
	}

	log.Println(path)
	data, err := ioutil.ReadFile(string(path))
	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 My Friend -" + http.StatusText(404)))
	}
}

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":5500", nil)
}

/*
func main() {

	//	http.HandleFunc("/", someFunc)
	//	http.HandleFunc("/about", aboutFunc)
	//	http.ListenAndServe(":5500", nil) // nil 表示使用默认mux

	// 自己创建路由服务
	mMux := http.NewServeMux()
	mMux.HandleFunc("/", someFunc)
	http.ListenAndServe(":5501", mMux)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello go web some app"))
}

func aboutFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello about page"))
}
*/
