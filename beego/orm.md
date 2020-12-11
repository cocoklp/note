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

