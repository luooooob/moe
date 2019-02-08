## web develope tooltik

此处需要一篇博客解释

## todo

ctx上挂一堆res和req里我常用到的简写方法

考虑一下要不要保留state,

把别人的巧克力融化,倒进自己模子里

##todo:

文档 ❓

c.Next() ✔️

Controller起个名字

handlerFunc是因为别人有handler这个概念，去掉了handler这一层，我们这个直接叫handlerFun有点突兀
用Handler, 有http.Handler的概念在前，会被弄混，所以也不合适
Controller像是当年一大波宣扬自己是MVC模式的框架搞出来的东西，叫这个有点过时
我自己瞎起的话就叫Poi好了（逃

配置文件

method拦截

测试

写的越多, bug越多, 能少写就少写

c.Assert()需要重新思考一下，err!=nil到底算true还是false, 类型匹配的完吗，要不要用反射？

c.Required(), 参数种类需要提前验证减少麻烦，创业那个表单的灵感
纸上谈兵闷头造轮子行不通的，哪怕试用范围再小也不能只靠脑补，实践出真知

JSONRPC了解一下？只能完全匹配的话干嘛不大胆一点，path只写两样信息，版本和action,直接往函数名上精确匹配


一开始肯定想搞一个支持普遍意义上RESTfulAPI的框架
因为不想用第三方，啥都得自己写
理想的功能列一下
1. 最基本的，完全路径匹配
2. 高级一点就是defaultServeMux像那样， 匹配最长的
3. 另一种思路是加一个*通配符来匹配
4. 匹配带参数的路径
5. defaultServeMux居然连request.method匹配都不支持，这个一定要有
想办法往类似jsonRPC那样改造，
