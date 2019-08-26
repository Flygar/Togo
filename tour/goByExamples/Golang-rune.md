---
title: 'Golang rune and byte'
---

>The difference between golang's string and byte and rune

<!-- more -->
>在Go当中 string底层是用byte数组存的，并且是不可以改变的。

### 源码中的定义
```go
// string is the set of all strings of 8-bit bytes, conventionally but not
// necessarily representing UTF-8-encoded text. A string may be empty, but
// not nil. Values of string type are immutable.
type string string

// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.
type byte = uint8

// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
type rune = int32
```
**byte**
*unit8的别名，用来区分变量指的是byte还是unit8*
>8bit, 1字节, 无符号, 2的8次方=256;总计可表示256个字符, 表示Unicode中的英文, 标点符号绰绰有余;math.MaxUint8为255(从0计数);

```go
var a byte
a = 'A' // a指的就是Unicode中十进制65所代表的符号‘A’;其实也可以当成unit8整数来用,a的unit8的值是65
//a= “严格” // Unicode，每个中文由3个byte组成
//a的类型: uint8, a的值: 65, a的Unicode符号:A
fmt.Printf("a的类型: %T, a的值: %v, a的Unicode符号:%v", a, a, string(a))

var b uint8
b = 66 // b指的就是unit8类型整数66;也可以当作byte类型来用,代表了Unicode中十进制66所代表的符号‘B’;
//b='B'
fmt.Printf("\nb的类型: %T, b的值: %v, b的Unicode符号:%v", b, b, string(b))

fmt.Println(math.MaxUint8)
```

**rune**
*int32的别名，用来区分变量指的是rune还是int32*
>32bit, 4字节, 有符号, 2的31次方=2147483648; 理论上总计可表示2的32次方个字符(负数也算), 英文, 中文, 日文, 韩文都可以表示;math.MaxInt32=2147483647(从0计数);

```go
//rune 对应unicode中的码点
var a rune  
a = '严' //a指的就是Unicode中的符号‘严’;其实也可以当成int32类型整数来用,a的int32的值是20005;
//a的类型: int32, a的值: 20005, a的ASCII符号:A
fmt.Printf("a的类型: %T, a的值: %v, a的Unicode符号:%v", a, a, string(a))

var b int32
b = 20006 // b指的就是int32类型整数20006;也可以当作rune类型来用,代表了Unicode中十进制20006所代表的符号‘並’;
//b='並'
fmt.Printf("\nb的类型: %T, b的值: %v, b的Unicode符号:%v", b, b, string(b))

fmt.Println(math.MaxInt32)
```

**string**
*字符串*
```go
//------------
b := []byte{'A', 'B', 'a', 'b'}
fmt.Printf("%v\n", b)
fmt.Printf("%v\n", string(b))
//通过占位符的方式格式化输出
fmt.Printf("%s\n", b)
//------------
a := "AB够浪"                    // Unicode，每个中文由3个byte组成
fmt.Println("总字节长度: ", len(a)) // 长度为 2+3*2=8
fmt.Println(a)                 //AB够浪

a0 := []byte(a) //将字符串转化为slice
a0 = append(a0, byte('A'))
fmt.Println("总字节长度: ", len(a0)) // 长度为 2+3*2+1=9
fmt.Println(a0)                 //[65 66 229 164 159 230 181 170 65]
fmt.Println(string(a0))         //AB够浪A

a1 := []rune(a)                 // 将字符串转化为rune,UTF-8
fmt.Println("总字节长度: ", len(a1)) // 长度为4
fmt.Println(a1)                 //[65 66 22815 28010]
fmt.Println(string(a1))         //AB够浪

fmt.Println("++++++++++++++++++++++++++++++++++")

for k, v := range []rune(a) {
    fmt.Printf("字符串中第%d个字符: %s; UTF-8 的值: %d\n", k, string(v), v)
    //fmt.Println(k," - ",[]byte(a)[k])
    fmt.Println("	|")
    fmt.Printf("字符: %s 由字节 %v 组成\n", string(v), []byte(string([]rune(a)[k])))
    for i, j := range []byte(string(v)) {
        fmt.Printf("字符: %s 第 %d 个字节 %v \n", string(v), i+1, j)
    }
    fmt.Println("----------------------------------")
}
```

**string and int**
```go
//str:="hello"
strInt := "12345"

//test:=int(strInt) //Cannot convert expression of type string to type int
str2int, err := strconv.Atoi(strInt)
if err != nil {
    fmt.Println("err:", err)
}
fmt.Println(str2int)

//打印十进制数65所代表的二进制数
fmt.Printf("%0b\n", 65) // 01000001

//十进制65代表了二进制数01000001，在Utf-8中字符表示为'A'
fmt.Println(string(65)) // A
//十进制20005代表了二进制01001110 00100101(unicode)，在Utf-8中变成(三字节:1110xxxx 10xxxxxx 10xxxxxx)字符表示为'严'
fmt.Println(string(20005)) // '严'
```