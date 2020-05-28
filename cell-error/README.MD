## 自定义error
 * 实现go原生的error接口，自定义一个解码函数MyErrorEncoder，在[]kithttp.ServerOption{}里面包裹进去，MyErrorEncoder函数将状态码直接写进http status code里面，这有利于日志错误分类的收集和统计
