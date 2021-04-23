# token



## JWT

1. 客户端调用登录接口，传入用户名和密码，获取token

2. 服务端确认用户名密码正确，创建JWT并返回给客户端

3. 客户端拿到JWT，进行存储(缓存、db、cookie)，后续请求，在HTTP请求头种加上JWT

4. 服务端校验JWT，校验通过后返回相关资源和数据

   JWT还可以存用户或者角色信息。

jwt由三部分组成

1. header

   jwt基本信息

2. payload

   标准中注册的声明

   公共的声明

   ​	可以添加任何信息

   私有的声明

   ​	提供者和消费者共同定义的声明，不建议存放敏感信息。

3. signature

   创建签名需要使用 Base64 编码后的 header 和 payload 以及一个秘钥。将 base64 加密后的 header 和 base64 加密后的 payload 使用. 连接组成的字符串，通过 header 中声明的加密方式进行加盐 secret 组合加密，然后就构成了 jwt 的第三部分。比如：HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload), secret)

   