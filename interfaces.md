### 接口
>go没有类和继承的概念 
> 
>通过**接口**来实现面向对象的特性
>
>接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。

定义接口：
```go
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}
```