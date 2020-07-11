# 软删除
deleted标签

type tb1 struct {
  DeleteAt      time.Time `json:"deleteAt" xorm:"delete_at deleted"`
}

CREATE TABLE `dns_change_record` (
  `delete_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='test';

xorm.Find()/Get()/Count()
  只返回未被删除的数据
xorm.Unscoped().Find()/Get()/Count()
  返回所有数据
xorm.Delete()
  软删除，会给delete_at字段赋值，表示已删除
xorm.Unscoped().Delete
  彻底删除

xorm.Update 不会对deleted字段进行处理

对delete_at的处理为
	插入时，没有 delete_at字段，所以表里是默认值
	查询时，增加条件 `delete_at` IS NULL OR `delete_at` = '0001-01-01 00:00:00'
	所以如果 default是0001-01-01 00:00:00，那么命中的 = 条件
	如果default 是 null命中null条件，如果default是 0000-00-00 00:00:00，这个值是null，仍然会命中null
