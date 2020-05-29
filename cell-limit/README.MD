## 自定义error
 * 限流这层按说是应该在网关那做的，这里为了学习go kit暂时放在这里
 * 使用gokit的endpointMiddleWare定义一个限流中间件，在main函数里面将要限流的接口用这个函数装饰一下，这个接口就能实现限流相关功能
 * 启动服务，访问`localhost:8000/user/2`,连续点击url的刷新按钮，开始会正常返回，过一会之后会返回我们在中间件里面自定义的错误"too many request，please waiting..."，http status code也会变成429
