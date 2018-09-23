# Togo

#### Go[标准库](https://gowalker.org/search?q=gorepos)完整列表

#### 开启本地文档
`godoc -http=:6060`  
`godoc -http=:6060 -goroot="."`

#### Note
外部源码将被下载到`$GOPATH/src`目录下并被安装到`$GOPATH/PKG/"machine_arch"/`目录下  
值类型: 变量的值存储在栈中  
引用类型: slices, maps, interface, channel; 被引用的变量会存储在堆中，以便进行垃圾回收，且比栈拥有更大的内存空间。  
字符串: 值类型, 且值不可改变  
range: 值拷贝  
数组: 值拷贝  
切片: 本质就是指向数组的指针,引用传递,所以不需要对切片使用指针  
在函数调用时: 像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）。  
