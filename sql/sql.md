# 表2作为表1的条件
select * from tb1 where tb1_col1 in (select tb2_col1 from tb2 as s where tb2_col2 ='value');

# 两个表分页
## union
连接两个以上的select语句的结果组合到一个结果集合中，多个select语句会删除重复数据
select user_pin from instance_info where valid =1 union select user_pin from instance_usage_info where valid =1 limit 1,3;
## union all
包含重复数据
