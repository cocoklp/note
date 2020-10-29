package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func main() {
	dbConf := &DbConf{
		User:     "test",
		PassWord: "test",
		Host:     "ip:port",
		Name:     "db",
	}
	DbPool, err := newDb(dbConf)
	if err != nil {
		return
	}
	file, err := os.Open("test.png") // For read access.
	if err != nil {
		fmt.Println(err)
	}

	bytes := make([]byte, 1024*1024*4)
	_, err = file.Read(bytes)
	if err != nil {
		fmt.Println(err)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
	data := string(bytes)
	dataA, err := CompressStrWithGzip(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(data), len(dataA), len(data)/len(dataA))
	dn := DnRegion{
		Country: "test",
		City:    "test",
		Content: dataA,
	}
	dataW, err := UncompressStrWithGzip(&dnN.Content)
	if err != nil {
		fmt.Println(err)
	}
	err = WriteFileAtomic("./test2.png", []byte(dataW), 0755)
	if err != nil {
		fmt.Println(err)
	}

	{
		err = WriteFileAtomic("./test.cache", []byte("bbbbbaaaa"), 0755)
		fmt.Println(err)
	}
}
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

func WriteFileAtomic(filename string, data []byte, perm os.FileMode) error {
	dir, name := path.Split(filename)
	f, err := ioutil.TempFile(dir, name)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err == nil {
		err = f.Sync()
	}
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	if permErr := os.Chmod(f.Name(), perm); err == nil {
		err = permErr
	}
	if err == nil {
		err = os.Rename(f.Name(), filename)
	}
	// Any err should result in full cleanup.
	if err != nil {
		os.Remove(f.Name())
	}
	return err
}

type DbConf struct {
	User     string `json:"user"`
	PassWord string `json:"passWord"`
	Host     string `json:"host"`
	Name     string `json:"name"`
}

func newDb(dbConf *DbConf) (*xorm.Engine, error) {
	driverName := "mysql"
	dsnConfig := &mysql.Config{
		User:                 dbConf.User,
		Passwd:               dbConf.PassWord,
		Addr:                 dbConf.Host,
		Net:                  "tcp",
		DBName:               dbConf.Name,
		AllowNativePasswords: true,
		ReadTimeout:          time.Duration(60) * time.Second, //  I/O read timeout
		WriteTimeout:         time.Duration(60) * time.Second, //  I/O write timeout
		Timeout:              time.Duration(60) * time.Second, // Dial timeout
	}
	dataSourceName := dsnConfig.FormatDSN()
	fmt.Printf("dataSourceName %s\n", dataSourceName)
	//dataSourceName := "root:123456@tcp(172.28.17.130:3306)/cloud_wafapi_db?allowNativePasswords=true&readTimeout=1s&timeout=1s&writeTimeout=1s&maxAllowedPacket=0"
	dB, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		fmt.Println("xorm new engin err: %s ", err.Error())
		return nil, err
	}
	dB.ShowSQL(false)
	err = dB.Ping()
	if err != nil {
		fmt.Println("db ping err: %a", err.Error())
		return nil, err
	}
	return dB, nil
}
