>今天我们继续说os包，OK我们直接来
 
####(1)os.Mkdir创建单个目录函数原形func Mkdir(name string, perm FileMode) error输入一个目录的名称和目录的权限，我们可以用默认的os.ModePerm然后返回的是一个error的信息，我们看下,也一块复习前边的一点知识
	
	import (
		"fmt"
		"os"
	)

	func main() {
		var path string
		if os.IsPathSeparator('\\') {  //前边的判断是否是系统的分隔符
			path = "\\"
		} else {
			path = "/"
		}
		fmt.Println(path)
		dir, _ := os.Getwd()  //当前的目录
		err := os.Mkdir(dir+path+"md", os.ModePerm)  //在当前目录下生成md目录
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("创建目录" + dir + path + "md成功")
	}
	
---
	
####(2)os.MkdirAll()函数原形是func MkdirAll(path string, perm FileMode) error输入的是多级目录结构和权限返回的是error的信息

	import (
		"fmt"
		"os"
	)

	func main() {
		dir, _ := os.Getwd()
		err := os.MkdirAll(dir+"/a/b/c", os.ModePerm)  //生成多级目录
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("创建文件夹" + dir + "/a/b/c成功")
	}

---	
	
####(3)第三个我们要说的os.NewSyscallError()函数的原形是func NewSyscallError(syscall string, err error) error输入的第一个是系统调用名称，第二个是错误信息的对象，返回是SyscallError的信息

	import (
		"errors"
		"fmt"
		"os"
	)

	func main() {
		fmt.Printf("%s\n", os.NewSyscallError("custom", errors.New("something wrong"))) //后边我们会继续讲Syscall包
	}
	
---

####(4)第四个是os.ReadLink()函数原形是func Readlink(name string) (string, error)输入的是链接的名称返回的是目标文件和err的错误返回信息

	import (
		"fmt"
		"os"
	)

	func main() {
		data, _ := os.Readlink("src") //我的linux返回的是/home/widuu/go/src
		fmt.Println(data)
	}

---	
	
####(5)第五个函数os.Remove()函数的原形是func Remove(name string) error输入的是目录的名称返回的是error信息

	import (
		"fmt"
		"os"
	)

	func main() {
		err := os.Remove("md") //删除我们刚创建的md目录
		if err != nil {
			fmt.Println(err)
		}
	}

---	
	
####(6)删除os.RemoveAll()这个函数的原形是func RemoveAll(path string) error和创建一样只不过是删除

	import (
		"fmt"
		"os"
	)

	func main() {
		path, _ := os.Getwd()
		err := os.RemoveAll(path + "/a")  //删除/a/b/c目录
		if err != nil {
			fmt.Println(err)
		}
	}

>好了今天我们就先到这里，然后我们明天继续~~~首发在我的博客[微度网络-网络技术中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)