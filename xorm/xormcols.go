package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"time"
)

type DbConf struct {
	User     string `json:"user"`
	PassWord string `json:"passWord"`
	Host     string `json:"host"`
	Name     string `json:"name"`
}

type DataUpdater struct {
	endKeeper chan bool
	ticker    *time.Ticker
	mt        *sync.Mutex
	wg        *sync.WaitGroup
}

const (
	updatePeriod = 5 * 60 // 5min
)

func NewDataUpdateTimer() *DataUpdater {
	s := &DataUpdater{}
	s.endKeeper = make(chan bool, 1)
	s.ticker = time.NewTicker(time.Second * time.Duration(updatePeriod))
	s.mt = &sync.Mutex{}
	s.wg = &sync.WaitGroup{}
	return s
}

func main() {

	dbConf := &DbConf{}
	var err error
	DbPool, err = newDb(dbConf)
	if err != nil {
		return
	}
	{
		dnR := new(DnReGion)

		//datas := make([]string, 0, 10)
		//err := DbPool.Table(dnR.TableName()).Cols("country","city").Find(&datas)  // 报错
		datas, err := (DbPool.Table(dnR.TableName()).Cols("country", "city").Distinct("country").Query())
		if err != nil {
			fmt.Println(err)
		}
		for _, data := range datas {
			fmt.Println(string(data["country"]))
			fmt.Println(string(data["city"]))
		}
	}
	return
}

func get() (*DnReGion, bool, error) {
	dnR := new(DnReGion)
	fmt.Printf("%p\n", dnR)
	has, err := DbPool.Where("id = ?", 1234).Get(dnR)
	fmt.Printf("%p\n", dnR)
	return dnR, has, err
}

type Employee struct {
	Id        int64  `xorm:"id autoincr"`
	Name      string `xorm:"name"`
	Salary    string `xorm:"salary"`
	ManagerId int64  `xorm:"ManagerId"`
}

func getdata(user *UserTest) (bool, error) {
	return DbPool.Get(user)
}

type UserTest struct {
	Id       int64
	Name     string
	Desc     string `xorm:"desc"`
	DeleteAt int64  `xorm:"deleted"`
	UpdateAt int64  `xorm:"updated"`
}

type Conf struct {
	DbConf `json:"dbConf"`
}

var DbPool *xorm.Engine

func loadConfig(path string) (*Conf, error) {
	opts := new(Conf)
	confBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(confBytes))
	err = json.Unmarshal(confBytes, opts)
	if err != nil {
		return nil, err
	}
	return opts, nil
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
	dB, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		fmt.Println("xorm new engin err: %s ", err.Error())
		return nil, err
	}
	//dB.SetTableMapper(core.SameMapper{})

	dB.ShowSQL(true)
	err = dB.Ping()
	if err != nil {
		fmt.Println("db ping err: %a", err.Error())
		return nil, err
	}
	return dB, nil
}
