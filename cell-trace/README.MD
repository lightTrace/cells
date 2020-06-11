## 链路追踪
- 通过docker安装jaeger环境：`sudo docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest`,能正常访问 http://127.0.0.1:16686/ 则安装成功
- 然后代码和metric的套路是一样的，这里在/cell-trace/register/middleware/trace.go里面实现一个中间件Metric用来在main函数里面后面装饰IUserService接口，同时这里实现IUserService的几个方法
- main.go的trace相关代码`tracer, _, err := middleware.NewJaegerTracer("user_agent_server")
               	if err != nil{
               		fmt.Println(err)
               	}
               	svc = middleware.NewTracerMiddleware(tracer)(svc)`
- 然后这里需要在方法里面都加上context，最后重启register服务
- 然后在浏览器访问一次localhost:8000/user/1，在jaeger的后台里面可以通过Find Traces按钮来查询刚刚的链路请求，甚至可以做到函数级别的查看
