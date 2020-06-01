## 服务注册与发现
- 其实我一直想问为什么要服务注册和发现，根据路由直接调用api不简单省事吗？看了很多资料，就我自己觉得，就是在用最合理的方式解决服务调用之间的强依赖，也有点依赖倒置的味道，对于上层服务来说，反正我需要这样的一个服务，不仅需要，而且调用起来还要得心应手，并且保证高可用性，下层服务你自己看着办吧？下层服务表示很委屈，于是就委托注册中心来做这件事了。
- 所以本章里面代码很多，但是就是围绕register，discover以及注册中心来做的
- 首先在/cells，运行`sudo docker-compose -f docker/docker-compose.yml up`用docker启动一个consul注册中心，启动成功后访问localhost:8500,看到控制界面
- 然后在/cell-register/register目录下运行`go run main.go -consulHost localhost -consulPort 8500 -serviceHost 192.168.192.145 -servicePort 8000`,就将我们的服务注册到上一步启动的consul里面了，可以通过浏览器页面看到
- 然后在/cell-register/discover目录下运行`go run main.go -consulHost localhost -consulPort 8500`,这个服务就成功调用了register注册的user服务的getUserName接口了，并且通过9001端口暴露出去成为一个新的接口，可以访问localhost:9001/user/1获取结果
- 可以通过不同的port启动两个register服务，调用的时候我们采用随机负载均衡的方式调用两个端口提供的register的服务，这里如果采用的不是不同的端口而是不同的几台机器就实现了这个服务的分布式注册且高可用
