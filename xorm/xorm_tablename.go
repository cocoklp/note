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

	dbConf := &DbConf{
		User:     "test",
		PassWord: "test",
		Host:     "test",
		Name:     "test",
	}
	var err error
	DbPool, err = newDb(dbConf)
	if err != nil {
		return
	}

	{
		fmt.Println("test table name start")
		employee := new(Employee)
		dnRn := new(DnRegion)
		DbPool.Table(employee).Get(employee)
		DbPool.Table(dnRn).Get(dnRn)
		DbPool.Table(dnRn.TableName()).Get(dnRn)
		fmt.Println("test table name end")
	}
	{
		dnRn := &Employee{
			Name: "Max",
		}
		{
			salaries := make([]string, 0, 10)
			err := DbPool.Table(dnRn.TableName()).Where("name = ?", dnRn.Name).Cols("salary").Find(&salaries, dnRn)
			fmt.Println(err, salaries, dnRn)
		}
		salaries := make([]Employee, 0, 10)
		err := DbPool.Table(dnRn.TableName()).Find(&salaries, dnRn)
		fmt.Println(err, salaries, dnRn)
	}
	return
}

type Employee struct {
}

func (e *Employee) TableName() string {
	return "employee"
}

type DnRegion struct {
}

func (t *DnRegion) TableName() string {
	return "dn_region"
}

func getdata(user *UserTest) (bool, error) {
	return DbPool.Get(user)
}

type UserTest struct {
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
