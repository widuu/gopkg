>今天我们讲golang标准库的os包type File struct{},还是废话少说直接上代码

####(1)os.Create()这个函数是创见一个文件，函数的原型是func Create(name string) (file *File, err error)输入的是名称字符串类型，返回的是一个File的指针和一个error
	import (
		"fmt"
		"os"
		"reflect"
	)
	
	func main() {
		f, _ := os.Create("widuu_2.go")
		defer f.Close()
		fmt.Println(reflect.ValueOf(f).Type()) //*os.File
	}



> 这个函数的原理其实是这样的OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666) O_RDWR也就是说用读写的权限，O_CREATE然后文件存在忽略，不存在创建它，O_TRUNC文件存在截取长度为0，这就解释了为什么我们明明有这个文件，我擦，创建之后哭了~啥都没有了~~用的时候需谨慎，先判断文件是否存在~

---
####(2)os.OpenFile函数的原型是func OpenFile(name string, flag int, perm FileMode) (file *File, err error)要指定文件权限和打开的方式,就是我们上边所用到的
	import (
		"fmt"
		"os"
	)
	
	func main() {
		f, _ := os.OpenFile("10.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		defer f.Close()
		fmt.Println(f.Stat())
	}

>这个就是上边的Create()只不过权限是0777以及下边的操作等大部分用到OpenFile()

####(3)os.Open()这个函数是打开文件使用的，函数原型是func Open(name string) (file *File, err error)，返回值就不说了一样的，它的其实原理是这样的OpenFile(name, O_RDONLY, 0)以读文件的模式打开
	import (
		"fmt"
		"os"
		"reflect"
	)
	
	func main() {
		f, _ := os.Open("1.go")
		defer f.Close()
	}

---

####(4)os.Stat()这个是获取fileinfo的结构描述func Stat(name string) (fi FileInfo, err error)返回了Fileinfo这个结构，我们再前边也详细讲了
，其实它是怎么实现的呢？因为我们没讲syscall所以我们就讲一个，譬如FileInfo底层获取fs := &fileStat{name: basename(name)}然后后边逻辑大家可以看源代码
	import (
		"fmt"
		"os"
	)
	
	func main() {
		f, _ := os.Stat("1.go")
		fmt.Println(f.Size())
	}

---

####(5)os.Fd()返回文件的句柄，函数原型是func (file *File) Fd() uintptr函数是这样的uintptr(file.fd) 返回的是文件的句柄，句柄是什么？句柄，是整个windows编程的基础。一个句柄是指使用的一个唯一的整数
	import (
		"fmt"
		"os"
	)
	
	func main() {
		f, _ := os.Open("1.go")
		fmt.Println(f.Fd()) //我的平台句柄是228
	}

---

####(6)os.Pipe()这个函数获取的函数的读写指针，函数原型func Pipe() (r *File, w *File, err error)
	import (
		"fmt"
		"os"
	)
	
	func main() {
		r, w, _ := os.Pipe()
		fmt.Println(r, w) //&{0xc08402e120} &{0xc08402e180}
	}

---

####(7)os.NewFile()函数原型是func NewFile(fd uintptr, name string) *File 第一个传入的是句柄，然后是文件名称，这个函数并不是真的创建了一个文件，是新建一个文件不保存，然后返回文件的指针
	import (
		"fmt"
		"os"
	)
	
	func main() {
		f, _ := os.Open("ini.go")
		defer f.Close()
		f1 := os.NewFile(f.Fd(), "ceshi.go") //输如ini.go的句柄
		defer f1.Close()
		fd, _ := f1.Stat()
		fmt.Println(fd.ModTime())	//返回的是ini.go的创建时间2013-11-27 09:11:50.2793737 +0800 CST
	
	}

---

####(8)(f *File).Chdir()修改工作目录，函数原型func (f *File) Chdir() error，这个时候f必须是目录了,但是吧这个不支持windows
	import (
		"fmt"
		"os"
	)
	
	func main() {
		dir, _ := os.Getwd()
		fmt.Println(dir)
		f, _ := os.Open("views")
		err := f.Chdir()
		if err != nil {
			fmt.Println(err)
		}
		dir1, _ := os.Getwd()
		fmt.Println(dir1)
	}

---

####(9)Chmod(),Chown()都是一样的，当然大家可以直接用os.Chmod(),os.Chdir()函数来改变

>今天也就到这里，明天我们继续哈~~~
