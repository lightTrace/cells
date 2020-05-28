## 使用go-kit构建一个http服务
### 创建service
 * gokit的设计理念，是将Service将作为核心业务逻辑实现部分。本文使用UserService实现IUserService接口获取用户姓名和更新用户姓名的服务，使用map模拟数据库存储更新数据
### 创建endpoint
 * 使用endpoints.go创建请求、响应模型，即GetUserNameRequest，GetUserNameResponse...，gokit采用装饰着模式，把Service应该执行的逻辑封装到Endpoint方法中执行。Endpoint的作用是：调用IUserService中相应的方法处理请求对象（即GetUserNameRequest），返回响应对象（即GetUserNameResponse）
### 创建transport
 * Transport层用于接收用户网络请求并将其转为Endpoint可以处理的对象，然后交由Endpoint执行，最后将处理结果转为响应对象向用户响应
### 运行main.go即启动http服务了
 * GET方式访问`localhost:8000/user/1` 返回jack
 * GET方式访问`localhost:8000/user/2` 会返回用户不存在的错误，因为初始化map的时候只放进id为1，name为jack的服务
 * POST方式访问 `localhost:8000/user/1/rose` 修改了id为1的用户名为rose，可再次GET方式访问`localhost:8000/user/1验证
