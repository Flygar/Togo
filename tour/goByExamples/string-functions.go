/*
标准库的 strings 包提供了很多有用的字符串相关的函数
你可以在 [`strings`](http://golang.org/pkg/strings/)包文档中找到更多的函数
*/
package main

import "fmt"

// strings包别名s
import s "strings"

// 我们给 `fmt.Println` 一个短名字的别名，我们随后将会经常用到。
var p = fmt.Println

func main() {
	p("Contains:  ", s.Contains("test", "es"))        //true
	p("Count:     ", s.Count("teeeestttt", "ee"))     //2,统计的是飞重叠的次数
	p("Count:     ", s.Count("test", "")-1)           //返回字符串长度4
	p("Fields:    ", s.Fields(" t e  s  t  \t"))      //[t e s t];将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。
	p("HasPrefix: ", s.HasPrefix("test", "te"))       //true
	p("HasSuffix: ", s.HasSuffix("test", "st"))       //true
	p("Index:     ", s.Index("test", "e"))            //1
	p("LastIndex: ", s.LastIndex("test", "t"))        //3
	p("IndexRune: ", s.IndexRune("我好渣", '渣'))         //6，前面2个汉子各占3个字节
	p("Join:      ", s.Join([]string{"a", "b"}, "-")) //a-b
	p("Repeat:    ", s.Repeat("a", 5))                //aaaaa
	p("Replace:   ", s.Replace("foo", "o", "0", -1))  //f00
	p("Replace:   ", s.Replace("foo", "o", "0", 1))   //f0o
	p("Split:     ", s.Split("a-b-c-d-e", "-"))       // [a b c d e]
	p("ToLower:   ", s.ToLower("TEST"))               // test
	p("ToUpper:   ", s.ToUpper("test"))               // TEST
	p("TrimSpace: ", s.TrimSpace(" test "))           //来剔除字符串开头和结尾的空白符号;
	p("Trim:      ", s.Trim("test", "t"))             //来将开头和结尾的 t 去除掉。

	p()

}
