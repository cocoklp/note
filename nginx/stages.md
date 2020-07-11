location /test {
set $a 32;
echo $a;

set $a 56;
echo $a;
}
输出两个56，set阶段始终在content阶段之前执行，与代码的先后顺序无关

1st：post-read 阶段
	读取并解析完请求头之后立即开始运行


bytes_sent
	nginx返回给客户端的字节数，包括响应头和响应体。1字节占 8 bite

body_bytes_sent
	nginx 返回给客户端的响应体的字节数，不包含响应头

content_length
	等同于 http_content_length ，即 nginx 从客户端收到的请求头中 content-length 字段的值

request_length：
	请求的字节数，包括请求行、请求头、请求体。是请求解析过程中不断累加的。

upstream_response_length
	从upstream服务器获得的响应体的字节数，不包括响应头。

upstream_bytes_received
	从upstream服务器获得的响应体和响应头的字节数