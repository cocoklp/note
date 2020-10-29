package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	//"os"
	"time"
)

func main() {
	dbConf := &DbConf{}
	var err error
	DbPool, err = newDb(dbConf)
	if err != nil {
		return
	}
	users := make([]DnWafRuleCondition, 0)

	rule := new(DnWafRules)
	cond := new(DnWafCondition)
	err = DbPool.Join("LEFT", cond.TableName(), cond.TableName()+".domain = "+rule.TableName()+".domain").
		In(rule.TableName()+".id", []int{81, 82, 83, 94, 95}).Find(&users)
	fmt.Println(users, err)

}

type Conf struct {
	DbConf `json:"dbConf"`
}

type Rulecond struct {
	DnWafRules         `xorm:"extends"`
	DnWafRuleCondition `xorm:"extends"`
	DnWafCondition     `xorm:"extends"`
}

func (d *Rulecond) TableName() string {
	return "dn_waf_rules"
}

type DnWafRules struct {
}

func (d *DnWafRules) TableName() string {
	return "dn_waf_rules"
}

type DnWafCondition struct {
}

func (c *DnWafCondition) TableName() string {
	return "dn_waf_condition"
}

type DnWafRuleCondition struct {
	DnWafRules     `xorm:"extends"`
	DnWafCondition `xorm:"extends"`
}

func (d *DnWafRuleCondition) TableName() string {
	return "dn_waf_rules"
}

type ListWafRuleData struct {
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
	dB.ShowSQL(true)
	err = dB.Ping()
	if err != nil {
		fmt.Println("db ping err: %a", err.Error())
		return nil, err
	}
	return dB, nil
}
