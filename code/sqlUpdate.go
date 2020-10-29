package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"time"
)

var DbPool *xorm.Engine
var DataupdateTimer *DataUpdater

type DataUpdater struct {
	endKeeper chan bool
	ticker    *time.Ticker
	mt        *sync.Mutex
	wg        *sync.WaitGroup
}

type DbConf struct {
	User     string `json:"user"`
	PassWord string `json:"passWord"`
	Host     string `json:"host"`
	Name     string `json:"name"`
}

const (
	updatePeriod = 10 //5 * 60 // 5min
)

type UserTest struct {
	Id       int64
	Name     string
	Desc     string `xorm:"desc"`
	DeleteAt int64  `xorm:"deleted"`
	UpdateAt int64  `xorm:"updated"`
}

func main() {
	dbConf := &DbConf{
		User:     "root",
		PassWord: "123456",
		Host:     "172.28.17.130:3306",
		Name:     "klp_test",
	}
	var err error
	DbPool, err = newDb(dbConf)
	if err != nil {
		fmt.Printf("newdb error: %s \n", err.Error())
		return
	}
	user := make([]UserTest)
	fmt.Println(DbPool.Find(&user))
	fmt.Println(user)
	DataupdateTimer = NewDataUpdateTimer()
	go DataupdateTimer.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until a signal is received.
	sig := <-c
	fmt.Printf("Trapped Signal; %v \n", sig)

	// Stop the service gracefully.
	DataupdateTimer.Stop()
	DbPool.Close()
	fmt.Printf("Gracefully Quit! \n")
	os.Exit(0)
}

func (s *DataUpdater) UpdateData() error {
	fmt.Printf("Update \n")
	return nil
}
func (s *DataUpdater) Start() {
	s.wg.Add(1)
	fmt.Printf("start update data \n")
	defer s.wg.Done()
	for {
		select {
		case <-s.ticker.C:
			err := s.UpdateData()
			if err != nil {
				fmt.Printf("ERROR!!! update data fail, %v \n", err)
			}
		case <-s.endKeeper:
			fmt.Printf("Close update data \n")
			return
		}
	}
}

func (s *DataUpdater) Stop() {
	s.ticker.Stop()
	close(s.endKeeper)
	s.wg.Wait()
}

func NewDataUpdateTimer() *DataUpdater {
	s := &DataUpdater{}
	s.endKeeper = make(chan bool, 1)
	s.ticker = time.NewTicker(time.Second * time.Duration(updatePeriod))
	s.mt = &sync.Mutex{}
	s.wg = &sync.WaitGroup{}
	return s
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
