join
https://www.cnblogs.com/fudashi/p/7491039.html
笛卡尔积，将A表的每一条记录和B表的每一条记录强行拼在一起。A有n条，B有m条，则笛卡尔积有m*n条

select *From user_test, user;
inner join:
	// 从结果中挑出 on条件成立的部分
	select *From user_test join user on user_test.id=user.id;

left join:
	// 两表交集外加左表剩余部分
	select *From user_test left join user on user_test.id=user.id;
right join:
	// 两表交集外加右表剩余部分
	select *From user_test right join user on user_test.id=user.id;
outer join:
	// 两表并集
	select *From user_test right join user on user_test.id=user.id
	union
	select *From user_test left join user on user_test.id=user.id;
using:
	MySQL中连接SQL语句中，ON子句的语法格式为：table1.column_name = table2.column_name。当模式设计对联接表的列采用了相同的命名样式时，就可以使用 USING 语法来简化 ON 语法，格式为：USING(column_name)。
所以，USING的功能相当于ON，区别在于USING指定一个属性名用于连接两个表，而ON指定一个条件。另外，SELECT *时，USING会去除USING指定的列.


create table Student(sid varchar(10),sname varchar(10),sage datetime,ssex nvarchar(10));
insert into Student values('01' , '赵雷' , '1990-01-01' , '男');
insert into Student values('02' , '钱电' , '1990-12-21' , '男');
insert into Student values('03' , '孙风' , '1990-05-20' , '男');
insert into Student values('04' , '李云' , '1990-08-06' , '男');
insert into Student values('05' , '周梅' , '1991-12-01' , '女');
insert into Student values('06' , '吴兰' , '1992-03-01' , '女');
insert into Student values('07' , '郑竹' , '1989-07-01' , '女');
insert into Student values('08' , '王菊' , '1990-01-20' , '女');
create table Course(cid varchar(10),cname varchar(10),tid varchar(10));
insert into Course values('01' , '语文' , '02');
insert into Course values('02' , '数学' , '01');
insert into Course values('03' , '英语' , '03');
create table Teacher(tid varchar(10),tname varchar(10));
insert into Teacher values('01' , '张三');
insert into Teacher values('02' , '李四');
insert into Teacher values('03' , '王五');
create table SC(sid varchar(10),cid varchar(10),score decimal(18,1));
insert into SC values('01' , '01' , 80);
insert into SC values('01' , '02' , 90);
insert into SC values('01' , '03' , 99);
insert into SC values('02' , '01' , 70);
insert into SC values('02' , '02' , 60);
insert into SC values('02' , '03' , 80);
insert into SC values('03' , '01' , 80);
insert into SC values('03' , '02' , 80);
insert into SC values('03' , '03' , 80);
insert into SC values('04' , '01' , 50);
insert into SC values('04' , '02' , 30);
insert into SC values('04' , '03' , 20);
insert into SC values('05' , '01' , 76);
insert into SC values('05' , '02' , 87);
insert into SC values('06' , '01' , 31);
insert into SC values('06' , '03' , 34);
insert into SC values('07' , '02' , 89);
insert into SC values('07' , '03' , 98);

1.查询"01"课程比"02"课程成绩高的学生的信息及课程分数
select *From Student right join (select t1.sid, class1, class2  from (select sid, score as class1 from SC  where SC.cid = '01') as t1, (select sid, score as class2 from SC where SC.cid = '02') as t2 where t1.sid = t2.sid and t1.class1>t2.class2) r on Student.sid=r.sid;


2.查询存在" 01 "课程但可能不存在" 02 "课程的情况(不存在时显示为 null )
select *from (select *from SC where cid = '01') t1 left join (select *from SC where cid = '02') t2 on t1.sid = t2.sid;

3. 查询01和02都存在的情况
 select *from (select *from SC where cid = '01') t1 join (select *from SC where cid = '02') t2 on t1.sid = t2.sid;

4. 选择了02但是没有选择01的情况
select *from SC where sid not in (select sid from SC where cid = '01') and cid = '02';

6. 平均成绩大于等于60分的同学的学生编号和姓名和平均成绩
select SC.sid,avg(score),Student.sname from SC right join Student on Student.sid=SC.sid group by SC.sid having avg(score) >= 60;
```
// having VS where
where在group by 前面
select sum(score) from student where gender='boy' group by name having sum(score)>210;
where的作用是对查询结果进行分组前,把不符合where条件的行去掉,在分组前过滤数据,条件中不能包含聚合函数,where条件限制特定的行.
having 的作用是筛选满足条件的组,在分组之后过滤数据,条件中可以包含聚合函数

having和where都可以用的场景:
	select price, name from goods where price > 100;
	select price, name from goods havind price > 100;
只可以用where不可以用having的场景:
	select name from goods where price > 100;
	select name from goods having price > 100; // 错误,因为没有price字段,having不能用,where是对表进行检索.
只可以用having不可以用where
	select id, avg(price) as agprice from goods group by id having agprice > 100;
	select id, avg(price) as agprice from goods where agprice > 100 group by id;  // 报错,因为没有agvs字段
```
7. 查询在 SC 表存在成绩的学生信息
select *from Student where sid in (select sid from SC);

8.查询所有同学的学生编号、学生姓名、选课总数、所有课程的成绩总和
select *from (select sid,count(1),sum(score) from SC group by sid) as t1 join  (select sname,sid from Student) as t2 on t1.sid = t2.sid;

9.查询学过「张三」老师授课的同学的信息
select * from Student where sid in (select sid from SC as t1 where cid in (select cid from Course  where tid = (select tid from Teacher where tname = '张三')));

10. 查询没有学全所有课程的同学的信息
select *from Student where sid not in (select sid from SC  group by SC.sid having count(cid) >= (select count(*) from Course));

11. 查询至少有一门课与学号为" 01 "的同学所学相同的同学的信息
select * from SC inner join Student on SC.sid=Student.sid where SC.cid in (select cid from SC where sid = '01') and SC.sid != '01' group by Student.sid;




