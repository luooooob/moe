# Moe

一个超轻量无依赖的API开发框架

## 起步

## 用法

## 运行测试

## 版权

[MIT](LICENSE)

## todo

## todo

1. 文档, 即[readme.md](readme.md)和[readme_zh.md](readme_zh.md)

2. 路由自己写, 只匹配完全相同的, 不做模糊匹配, 应该不是很难

3. c.Required()功能, 参数种类需要提前验证减少麻烦

2. middleware, c.Next()等关键功能. 初步完成过一次,可能需要重构(固定几个middleware不留自定义空间?给request打几个标签几个固定字段不不留context的get,set?)

5. 从json配置文件启动

6. 优雅的错误处理和log功能

7. 测试, 写的越多, bug越多, 能少写就少写

其他
1. 开发缘由,需要一篇博客解释,从node的koa到gorilla/mux, gin.~~并不是简单的把别人的巧克力融化,倒进自己模子里~~

2. 给controller重新起个名字. handlerFunc是因为别人有handler这个概念，去掉了handler这一层，我们这个直接叫handlerFun有点突兀, 用Handler, 有http.Handler的概念在前，会被弄混，所以也不合适
Controller像是当年一大波宣扬自己是MVC模式的框架搞出来的东西，叫这个有点过时
我自己瞎起的话就叫Poi好了（逃. 待定

3. 思考,一个轻量化的框架用的时候不能直接写函数要包成接口的实现是不是太麻烦了, 要不要搞控制反转这一套

