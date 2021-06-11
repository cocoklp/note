# orm

初始化

```go
orm.RegisterDriver("postgres",orm.DRPostgres)  //注册驱动类型 驱动名称，数据库名称
url:=fmt.Sprintf("user=%s password=%s host = %s port =%s timezone = %s dnName = %s")
orm.RegisterDataBase("default","postgres"，url) // 获取orm连接对象
// 模型建表,首字母大写
type User struct {
    Id int
    Name string
    PassWord string
}
orm.RegisterModel(new(User))
orm.RunSyncdb("default", false,true) // 数据库别名 是否强制建表(即先drop再create) 是否显示sqwl语句
```

自动建表后，首字母变小写，驼峰变下划线。如果模型中变量是小写，则不会编译数据库总的表中字段。

```
err := dao.ReadData(&bean, "mep_name")
func ReadData(data interface{}, cols ...string) error {
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(data, cols...)
	return err
}
select from table where cols = ''
select ... from mep_meta where mep_name = bean.MepName set to bean
```



# controller

Beego [parser.go:147] Invalid @Param format. Needs at least 4 parameters panic: runtime error

注释中Param格式错误，缺少4个必须参数。