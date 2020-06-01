## api监控
### 代码及环境搭建
- 我们的服务上线了，除了从日志可以查看一些运行情况外，其实现在服务接口对于我们来说是一个黑盒，我们不知道每个服务接口的延时，被调用次数等等指标的运行情况，如果想实时知道这些指标，就需要我们的api监控
- 这里在/cell/register/middleware/metrics.go里面实现一个中间件用来在main函数里面后面装饰IUserService接口，同时这里实现IUserService的几个方法
- 首先在/cells目录下，运行`docker-compose -f docker/docker-compose.yml up -d`用docker启动prometheus和grafana的相关环境，如果想停止使用命令`sudo docker-compose -f docker/docker-compose.yml stop`
- 然后在浏览器不停访问几次localhost:8000/user/1

### 配置grafana
- 登录Grafana：通过浏览器访问localhost:3000，使用用户名admin和yml配置的密码password登录；
- 创建数据源：点击create your first datasource，选择Prometheus，配置HTTP.URL（建议设置为http://[IP]:[port]，不要使用localhost），最后保存成功。
- 创建Dashboard：点击create your first dashboard，点击Title下拉菜单中的Edit，即可进入如下界面，可以选择Metric选项下拉发现会有很多可以供我们选择的指标监控，选择一个就可以展现在dashboard上面

