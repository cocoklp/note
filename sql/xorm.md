# 软删除
deleted标签

type tb1 struct {
  DeleteAt      time.Time `json:"deleteAt" xorm:"delete_at deleted"`
}

CREATE TABLE `dns_change_record` (
  `delete_at` datetime NOT NULL COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='test';

xorm.Find()/Get()/Count()
  只返回未被删除的数据
xorm.Unscoped().Find()/Get()/Count()
  返回所有数据
xorm.Delete()
  软删除，会给delete_at字段赋值，表示已删除
xorm.Unscoped().Delete
  彻底删除

