location /test {
set $a 32;
echo $a;

set $a 56;
echo $a;
}
输出两个56，set阶段始终在content阶段之前执行，与代码的先后顺序无关

1st：post-read 阶段
	读取并解析完请求头之后立即开始运行