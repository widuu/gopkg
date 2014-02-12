#golang讲解（go语言）标准库分析之path


>新的一年我们继续开始golang的标准库的讲解，先找回点状态，所以今天我们来讲解一下path，因为这个标准库灰常简单，东西很少，所以我们就讲解这个~~，然后大家看下边吧

#####(1)`func Base(path string) string`这个函数主要是用来返回最后一个元素的路径,如果路径为空返回.如果路径由斜线组成,返回/

	package main
	
	import (
		"fmt"
		"path"
	)
	
	func main() {
		path := path.Base("D:/gopkg/gopkg")
		fmt.Println(path) //gopkg
	}

---

#####(2)`func Clean(path string) string`这个函数主要是返回等价的最短路径 清理规则：
 1. 用一个斜线替换多个斜线
 1. 清除当前路径.
 1. 清除内部的..和他前面的元素，如a/b/.. 得到结果a
 1. 以/..开头的，变成/ 

---

	package main
	
	import (
		"fmt"
		"path"
	)
	
	func main() {
		fmt.Println(path.Clean("a/c"))               // a/c
		fmt.Println(path.Clean("a//c"))              //以一个/代替// , a/c
		fmt.Println(path.Clean("a/c/."))             //清除. , a/c
		fmt.Println(path.Clean("a/c/b/.."))          // 清除内部..以及前面的元素b; a/c
		fmt.Println(path.Clean("111/../a/c"))        // 清除内部..以及前面的元素111; a/c
		fmt.Println(path.Clean("/../a/b/../././/c")) // 清除/..开头,..以及前面的元素b; /a/c
	}

---

#####(3)`func Dir(path string) string`函数返回路径的最后一个元素，通常返回最后一个元素的路径目录，如果目录为空则返回.

	package main
	
	import (
		"fmt"
		"path"
	)
	
	func main() {
		fmt.Println(path.Dir("/a/b/c")) // /a/b
		fmt.Println(path.Dir(""))       // .
	}

---

#####(4)`func Ext(path string) string`函数返回路径文件的扩展名

	package main
	
	import (
		"fmt"
		"path"
	)
	
	func main() {
		fmt.Println(path.Ext("D:/gopkg/gopkg/ext.go")) //.go
	}

---

#####(5)`func IsAbs(path string) bool`返回的是否是绝对路径

	package main
	
	import (
		"fmt"
		"path"
	)
	
	func main() {
		fmt.Println(path.IsAbs("D:/gopkg/gopkg")) //false
		fmt.Println(path.IsAbs("/dev/null"))      //true
	}

---

#####(6)`func Join(elem ...string) string`这个函数主要是连接路径，返回的结果是已经Clean的，如果是空路径就忽略

	package main
	
	import (
		"fmt"
		"path"
	)
	
	func main() {
		fmt.Println(path.Join("a", "b", "c"))          // a/b/c
		fmt.Println(path.Join("a", "", "c"))           // a/c
		fmt.Println(path.Join("a", "../bb/../c", "c")) // c/c
	}

---

#####(7)`func Match(pattern, name string) (matched bool, err error)`匹配文件名完全匹配返回true,如果是shell，语法格式如下：

	pattern:
		{ term }
	term:
		'*'         matches any sequence of non-/ characters
		'?'         matches any single non-/ character
		'[' [ '^' ] { character-range } ']'
		            character class (must be non-empty)
		c           matches character c (c != '*', '?', '\\', '[')
		'\\' c      matches character c
	
	character-range:
		c           matches character c (c != '\\', '-', ']')
		'\\' c      matches character c
		lo '-' hi   matches character c for lo <= c <= hi

	package main
	
	import (
	    "fmt"
	    "path"
	)
	
	func main() {
	    fmt.Println(path.Match("*","alll")) //true nil
	    fmt.Println(path.Match("*","a/lll")) //false nil
	    fmt.Println(path.Match("?","alll")) //false nil
	    fmt.Println(path.Match("?","a")) //true nil
	    fmt.Println(path.Match("a","a")) //true nil
	    fmt.Println(path.Match("\\a","a")) //true nil
	}

---

#####(8)`func Split(path string) (dir, file string)`这个函数是分割目录和文件的

	package main
	
	import (
		"fmt"
		"path"
	)
	
	func main() {
		fmt.Println(path.Split("widuu/widuu.go")) //	widuu/ widuu.go
	}


每天只讲一点golang的标准库，方便大家学习和使用，更多的时候去理解标准库，大家多动手，如果你喜欢请继续关注我们[微度网络-网络技术支持中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)