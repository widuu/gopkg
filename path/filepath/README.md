#golang讲解（go语言）标准库分析之filepath

####(1)func Abs(path string) (string, error)返回文件的绝对路径，如果参数不是绝对路径，文件+文件工作目录组成绝对路径，返回绝对路径和错误信息

	package main
	
	import (
		"fmt"
		"path/filepath"
	)
	
	func main() {
		dir, err := filepath.Abs("../clean.go")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(dir)	//D:\test\path\clean.go
	}


---

####(2)func Base(path string) string  见：path.Base是同样的用法

####(3)func Clean(path string) string 见path.Clean

####(4)func Dir(path string) string   见path.Dir

####(5)func EvalSymlinks(path string) (string, error) 返回一个链接文件的实际的路径。例如在/home下创建一个link.log的文件，然后cd xuchdong; ln -s /home/link.log link， 使用该函数可以找到原始的文件即/home/link.log

	package main
	
	import (
	    "fmt"
	    "path/filepath"
	)
	
	func main(){
	    // 环境准备：首先在/home目录下创建一个link.log的文件,
	    // 然后在当前目录下使用ln -s /home/link.log link_other
	
	    path, _ := filepath.EvalSymlinks("link_other")
	    fmt.Printf("%s\n", path) // /home/link.log
	}

---

####(6)func Ext(path string) string 见path.Ext

####(7)func FromSlash(path string) string 将所有的"/"使用路径分隔符替换。分隔符见os.PathSeparator, linux默认的分隔符是"/"。

	package main
	
	import (
		"fmt"
		"path/filepath"
	)
	
	func main() {
		a := filepath.FromSlash("D:/test/path/filepath/abs.go")
		fmt.Println(a) //windows D:\test\path\filepath\abs.go
	}

---

####(8)func Glob(pattern string) (matches []string, err error)返回目录下所有匹配的文件

	package main
	
	import (
		"fmt"
		"path/filepath"
	)
	
	func main() {
		m, _ := filepath.Glob("*.go")
		for i, v := range m {
			fmt.Println(i, v) //0 abs.go 1 base.go 2 fromslash.go 3 glob.go
		}
	}

---

####(9)func HasPrefix(p, prefix string) bool 存在兼容性问题不建议使用和string.Hasprefix用法一致

	package main
	
	import (
		"fmt"
		"path/filepath"
	)
	
	func main() {
		b := filepath.HasPrefix("D:/test/golang/", "D:/")
		fmt.Println(b) //true
	}

####(10)func IsAbs(path string) (b bool) 见path.IsAbs()

####(11)func Join(elem ...string) string 见path.Join()

####(12)func Match(pattern, name string) (matched bool, err error) 见path.Match

####(13)func Rel(basepath, targpath string) (string, error) 返回以basepath为基准的相对路径。Join(basepath, Rel(basepath, targpath)) 等于targpath。

	package main
	
	import (
		"fmt"
		"path/filepath"
	)
	
	func main() {
		dir, _ := filepath.Rel("D:/test/", "D:/test/path/filepath/rel.go")
		fmt.Println(dir)  //path\filepath\rel.go
	}

---

####(14)func Split(path string) (dir, file string) 见path.Split

####(15)func SplitList(path string) []string 将路径字符串使用路径列表分隔符分开。路径列表分隔符见os.PathListSeparator, linux的 路径列表分隔符是":", windows的路径列表分隔符是";"，主要用于PATH或是GOPATH等环境变 量。

	package main
	
	import (
		"fmt"
		"path/filepath"
	)
	
	func main() {
		dir := filepath.SplitList("D:/test/path;D:/test/path/filepath")
		for index, v := range dir {
			fmt.Println(index, v) //0 D:/test/path 1 D:/test/path/filepath
		
		}
	}

---

####(15)func ToSlash(path string) string 将所有的路径分隔符使用"/"替换。分隔符见os.PathSeparator, linux默认的分隔符是"/"。

	package main
	
	import (
		"fmt"
		"path/filepath"
	)
	
	func main() {
		a := filepath.ToSlash("D:\\test\\path\\pathfile")
		fmt.Println(a) //D:/test/path/pathfile
	}
	
---

####(16)func VolumeName(path string) (v string) 返回存储卷的卷标

	package main
	
	import (
		"fmt"
		"path/filepath"
	)
	
	func main() {
		volume := filepath.VolumeName("D:/test/path/filepath")
		fmt.Println(volume) //	D:
	}

---

####func Walk(root string, walkFn WalkFunc) error 根据walkFunc遍历目录下的所有文件，其中我们使用的type WalkFunc是type WalkFunc func(path string, info os.FileInfo, err error) error

	package main
	
	import (
		"fmt"
		"os"
		"path/filepath"
	)
	
	func walkfunc(path string, info os.FileInfo, err error) error {
		fmt.Printf("%s\n", path)
		return nil
	}
	
	func main() {
		filepath.Walk("D:/test/path/filepath", walkfunc)
		// D:/test/path/filepath
		// D:\test\path\filepath\abs.go
		// D:\test\path\filepath\base.go
		// D:\test\path\filepath\fromslash.go
		// D:\test\path\filepath\glob.go
		// D:\test\path\filepath\hasprefix.go
		// D:\test\path\filepath\rel.go
		// D:\test\path\filepath\splitlist.go
		// D:\test\path\filepath\toslash.go
		// D:\test\path\filepath\volumename.go
		// D:\test\path\filepath\walk.go
	
	}

每天只讲一点golang的标准库，方便大家学习和使用，更多的时候去理解标准库，大家多动手，如果你喜欢请继续关注我们[微度网络-网络技术支持中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)