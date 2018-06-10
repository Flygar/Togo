# Togo

#### Go[标准库](https://gowalker.org/search?q=gorepos)完整列表

#### 开启本地文档
`godoc -http=:6060`  
`godoc -http=:6060 -goroot="."`

#### Note
`go install`下载源码，编译并安装包
外部源码将被下载到`$GOPATH/src`目录下并被安装到`$GOPATH/PKG/"machine_arch"/`目录下


#### Method方法
任何类型都可以有方法(int、bool、string 或数组的*别名类型*)，除了接口，因为接口是一个抽象定义，但是方法却是具体实现;
类型和作用在它上面定义的方法必须在同一个包里定义;间接的方式：可以先定义该类型（比如：int 或 float）的别名类型，然后再为别名类型定义方法。或者将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。

#### 函数和方法的区别
>相同点: 在接收者是指针时，方法可以改变接收者的值（或状态），这点函数也可以做到（当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态）。
>方法是函数，所以同样的，不允许方法重载，即对于一个类型只能有一个给定名称的方法。但是如果基于接收者类型，是有重载的：具有同样名字的方法可以在 2 个或多个不同的接收者类型上存在.


```go
func (a *denseMatrix) Add(b Matrix) Matrix
func (a *sparseMatrix) Add(b Matrix) Matrix
func (_ receiver_type) methodName(parameter_list) (return_value_list) { ... }
```

- 指针方法和值方法都可以在指针或非指针上被调用:
    对于类型 T，如果在 *T 上存在方法 Meth()，并且 t 是这个类型的变量，那么 t.Meth() 会被自动转换为 (&t).Meth()
   

结构体：chapter_10

range:值拷贝

**指针方法和值方法都可以在指针或非指针上被调用**:[methodset1]()  
对于类型 T，并且 t 是这个类型的变量,如果在 *T 上存在方法 Meth()，那么 t.Meth() 会被自动转换为 (&t).Meth()