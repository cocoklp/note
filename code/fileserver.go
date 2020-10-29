package main

import (
	"fmt"
	"net/http"
	/*
		"os"
		"os/signal"*/
	"strings"
	"time"
)

func startServer() {
	http.Handle("/", TestHandle("."))
	s := &http.Server{
		Addr: ":8080",
	}
	s.ListenAndServe()
}

func main() {
	fmt.Println("before listen")
	go startServer()
	fmt.Println("after listen")
	/*
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		// Block until a signal is received.
		sig := <-c
		fmt.Printf("Trapped Signal; %v", sig)
	*/
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
	}
	for {
	}
}

type FileHandler struct {
	http.Handler
	count int8
}

func (f *FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	f.Handler.ServeHTTP(w, r)
	fmt.Println(r.URL, r.Method)
	url := r.URL
	fmt.Println("scheme:", url.Scheme, " opaque:", url.Opaque, " user:", url.User, " host:", url.Host, " path:", url.Path)
	fmt.Println(" rawpath:", url.RawPath, " forcequery:", url.ForceQuery, " rawquery:", url.RawQuery, " fragment:", url.Fragment)
	fullName := r.URL.Path
	nameSlice := strings.Split(fullName, "/")
	fmt.Println(nameSlice[len(nameSlice)-1])
}

func TestHandle(directory string) http.Handler {
	fh := new(FileHandler)
	fh.Handler = http.FileServer(http.Dir(directory))
	return fh
}
