



https://www.zhihu.com/question/19786827

https://www.cnblogs.com/lonelydreamer/p/6169469.html



https协议是无状态的协议，所以服务端需要记录用户的状态时，需要某种机制来标识具体的用户，该机制就是session。session保存在服务端，有一个唯一标识，保存在内存、数据库、文件等均可，对于集群需要考虑session的转移，在大型网站一般都有专门的session集群，用来保存用户会话。

首次创建session时，服务端给客户端一个sessionId，客户端在每次http请求的时候发送cookie到服务器，cookie中记录sessionId，服务器通过sessionId来识别用户身份。

# cookie

客户端技术，把每个用户是数据以cookie的形式写给用户各自的浏览器。用户使用浏览器访问服务其中的web资源时，就带着各自的数据去访问。

浏览器允许网页服务器在浏览器里存一小段数据，什么数据都行，你自己的格式自己去解析就好了．具体的做法是，浏览器第一次访问服务器时，服务器应答中就会包含需要浏览器请求的数据，浏览器收到服务器的应答，并把数据保存起来．当浏览器再次访问服务器时，浏览器就在请求里包含这段数据．由于这段数据不是我们的主要业务，只不过是我们进行主业务时的一点小插曲，故称这段数据为cookie，是甜甜圈，小点心，以区别于正餐．

# session

服务端技术。是保存在服务端的一个数据结构，用来跟踪用户的状态。

为每一个用户的浏览器创建一个其独享的session对象。所以可以把数据放到各自的session中，用户再去访问时其他web资源再从用户各自的session中取出数据为用户服务。