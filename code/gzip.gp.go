package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
)

func CompressStrWithGzip(s *string) (string, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	defer gz.Close()
	_, err := gz.Write([]byte(*s))
	if err != nil {
		return "", err
	}

	err = gz.Close()
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func UncompressStrWithGzip(s *string) (string, error) {
	b := new(bytes.Buffer)
	b.WriteString(*s)
	gr, err := gzip.NewReader(b)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	_, err = result.ReadFrom(gr)
	if err != nil {
		return "", err
	}

	if err = gr.Close(); err != nil {
		return "", err
	}
	return result.String(), nil
}

func main() {

	str := os.Args[1]
	str = `joijoih
	joihoihoj
	joihpoj[hhhhhhhhhhpoupo
	joiujo
	hoiy89tyuihp
	hoi898fihpo
	oioi`
	fmt.Println(str, len(str))
	compStr, err := CompressStrWithGzip(&str)
	fmt.Println(compStr, len(compStr), err)
	uncompStr, err := UncompressStrWithGzip(&compStr)
	fmt.Println(uncompStr, len(uncompStr), err)
}
