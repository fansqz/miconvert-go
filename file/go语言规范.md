# go语言规范

- 说明：这个文件用于记录我第一次用go语言写完整项目所遵循的规范
- 很多是直接借鉴网上别人的博客，有价值直接cv下来了。

## 命名规范



1. 包命名规范

   ~~~go
   //小写，简短，不用下划线，尽量不与标准库冲突
   package demo
   ~~~

2. 文件命名

   ~~~go
   //小写+下划线
   my_demo.go
   ~~~

3. 结构体命名

   ~~~go
   //驼峰命名，私有的首字母小写，公开的首字母大写
   type User struct {
       Username string
   }
   ~~~

4. 接口命名

   ~~~go
   //同结构体命名
   //单个函数结构名以“er”为后缀
   type Reader interface {
           Read(p []byte) (n int, err error)
   }
   ~~~

5. 变量命名

   ~~~go
   /*
   遵循驼峰
   如果变量为私有，且特有名词为首个单词，则使用小写，如 apiClient
   其它情况都应当使用该名词原有的写法，如 APIClient、repoID、UserID
   错误示例：UrlArray，应该写成 urlArray 或者 URLArray
   若变量类型为 bool 类型，则名称应以 Has, Is, Can 或 Allow 开头
   */
   var isExist bool
   var hasConflict bool
   var canManage bool
   var allowGitHook bool
   ~~~

6. 常量命名

   ~~~go
   //常量均需使用全部大写字母组成，并使用下划线分词
   const APP_VER = "1.0"
   //如果是枚举类型的常量，需要先创建相应类型：
   type Scheme string
   const (
       HTTP  Scheme = "http"
       HTTPS Scheme = "https"
   )
   ~~~

## 注释

- go 语言自带的 godoc 工具可以根据注释生成文档，生成可以自动生成对应的网站（ [http://golang.org](https://link.zhihu.com/?target=http%3A//golang.org) 就是使用 godoc 工具直接生成的），注释的质量决定了生成的文档的质量。每个包都应该有一个包注释

- 这个我的想法是，不如直接用goland的插件goanno



1. 结构体注释

   ~~~go
   //这里说下结构体的注释，结构用途申明一下，每个属性在后面说明用处即可
   // User ， 用户对象，定义了用户的基础信息
   type User struct{
       Username  string // 用户名
       Email     string // 邮箱
   }
   ~~~



## 代码风格



1. 关于缩进和折行

   我们使用Goland开发工具，可以直接使用快捷键：ctrl+alt+L，即可。

2. import规范

   ~~~go
   //不管一个包还是多个，统一
   import (
       "fmt"
   )
   ~~~

   

## 异常处理

- 详细看这篇https://mp.weixin.qq.com/s/RFF2gSikqXiWXIaOxQZsxQ

1. go的try_catch

   ~~~go
   func SomeProcess() (err error) { // <-- 注意，err 变量必须在这里有定义
       defer func() {
           if err == nil {
               return
           }
   
           // 这下面的逻辑，就当作 catch 作用了
           if errors.Is(err, somepkg.ErrRecordNotExist) {
               err = nil       // 这里是举一个例子，有可能捕获到某些错误，对于该函数而言不算错误，因此 err = nil
           } else if errors.Like(err, somepkg.ErrConnectionClosed) {
               // ...          // 或者是说遇到连接断开的操作时，可能需要做一些重连操作之类的；甚至乎还可以在这里重连成功之后，重新拉起一次请求
           } else {
               // ...
           }
       }()
   
       // ...
   
       if err = DoSomething(); err != nil {
           return
       }
   
       // ...
   }
   ~~~

   - 这里会有个问题，就是异常处理前置了，可读性没那么友好

2. 异常的返回

   - 有三种流派，我比较偏向于断言式流派

   ~~~go
   //异常定义
   type ErrRecordNotExist errImpl
   
   type ErrPermissionDenined errImpl
   
   type ErrOperationTimeout errImpl
   
   type errImpl struct {
       msg string
   }
   
   func (e *errImpl) Error() string {
       return e.msg
   }
   ~~~

   ~~~go
   //异常处理
   if err == nil {
       // OK
   } else if _, ok := err.(*ErrRecordNotExist); ok {
       // 处理记录不存在的错误
   } else if _, ok := err.(*ErrPermissionDenined); ok {
       // 处理权限错误
   } else {
       // 处理其他类型的错误
   }
   ~~~

3. 服务/系统错误信息返回

   - 对于需要向用户透露错误的信息，就用传统的code-message形式

   [go语言web开发系列之二:gin框架接口站统一返回restful格式的数据_老刘你真牛的博客-CSDN博客_gin 统一返回格式](https://blog.csdn.net/weixin_43881017/article/details/111178255)

   [封装ResultVO实现统一返回结果 - 掘金 (juejin.cn)](https://juejin.cn/post/6995932258662088717)

    两者做个结合吧

   - 对于未定义错误，或者不能给用户提供错误信息的，使用hash



## 关于循环依赖问题

- go似乎很难解决循环依赖问题。go一旦检验到a包依赖于b，b包依赖于a，就会报错
- 所以只能用合理的设计去解决这个问题

## 传递context



