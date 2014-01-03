> 今天我们继续讲golang标准库的os包，依然是没有废话直接来

###(1)(f *File).Name()这个函数是返回文件的名称，函数原型func (f *File) Name() string要文件的指针操作，返回字符串，感觉比较鸡助的方法底层实现
    
    func (f *File) Name() string { return f.name }
     import (
    	"fmt"
    	"os"
    )
    
    func main() {
    	f, _ := os.Open("1.go")
    	fmt.Println(f.Name()) //输出1.go
    }

###(2)(f *File).Read()这个是函数的指针来操作的,属于*FIlE的method,函数原型func (f *File) Read(b []byte) (n int, err error)输入读取的字节数，返回字节的长度和error信息

    import (
    	"fmt"
    	"os"
    )
    
    func main() {
    	b := make([]byte, 100) //设置读取的字节数
    	f, _ := os.Open("11.go")
    	n, _ := f.Read(b)
    	fmt.Println(n)	
    	fmt.Println(string(b[:n])) //输出内容 为什么是n而不直接输入100呢？底层这样实现的
    	/*
    		n, e := f.read(b)
       		if n < 0 {
    			n = 0
    	}
    		if n == 0 && len(b) > 0 && e == nil {
      			return 0, io.EOF
      		}
      	*/
      	//所以字节不足100就读取n
    }

###(3)(f *File).ReadAt()这个函数的原型是func (f *File) ReadAt(b []byte, off int64) (n int, err error)加入了下标，可以自定义读取多少
	import (
		"fmt"
		"os"
	)
	
	func main() {
		f, _ := os.Open("11.go")
		b := make([]byte, 20)
		n, _ := f.ReadAt(b, 15)
		fmt.Println(n)
		fmt.Println(string(b[:n]))
	}
###(4)(f *File).Readdir()函数原型func (f *File) Readdir(n int) (fi []FileInfo, err error)，我们要打开一个文件夹，然后设置读取文件夹文件的个数，返回的是文件的fileinfo信息
	import (
		"fmt"
		"os"
	)
	
	func main() {
		f, err := os.Open("src")	//打开一个目录
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		ff, _ := f.Readdir(10)    //设置读取的数量 <=0是读取所有的文件 返回的[]fileinfo
		for i, fi := range ff {
			fmt.Printf("filename %d: %+v\n", i, fi.Name())  //我们输出文件的名称
		}
	}
---
###(5)(f *File).Readdirnames这个函数的作用是读取目录内的文件名，其实上一个函数我们已经实现了这个函数的功能，函数的原型func (f *File) Readdirnames(n int) (names []string, err error)，跟上边一下只不过返回的是文件名 []string的slice
	import (
		"fmt"
		"os"
	)
	
	func main() {
		f, _ := os.Open("bin")
		names, err := f.Readdirnames(0)
		if err != nil {
			fmt.Println(err)
		}
		for i, name := range names {
			fmt.Printf("filename %d: %s\n", i, name)
		}
	}
---

###(6)(f *File).Seek()这个函数大家一看就懂了，就是偏移指针的地址，函数的原型是func (f *File) Seek(offset int64, whence int) (ret int64, err error) 其中offset是文件指针的位置 whence为0时代表相对文件开始的位置，1代表相对当前位置，2代表相对文件结尾的位置 ret返回的是现在指针的位置
	import (
		"fmt"
		"os"
	)
	
	func main() {
		b := make([]byte, 10)
		f, _ := os.Open("1.go")
		defer f.Close()
		f.Seek(1, 0)				//相当于开始位置偏移1
		n, _ := f.Read(b)
		fmt.Println(string(b[:n]))  //原字符package 输出ackage
	}
###(7)(f *File).Stat()其中跟前边的os.Stat()一样都是返回Fileinfo所以不多讲了
###(8)(f *File).Truncate()这个函数跟前边的os.Truncate()函数是一样的我就不多讲了 大家看os(5)
###(9)(f *File) Write像文件中写入内容，函数原型func (f *File) Write(b []byte) (n int, err error)返回的是n写入的字节数
	import (
		"fmt"
		"os"
	)
	
	func main() {
		f, _ := os.OpenFile("1.go", os.O_RDWR|os.O_APPEND, 0755) //以追加和读写的方式去打开文件
		n, _ := f.Write([]byte("helloword"))                     //我们写入hellword
		fmt.Println(n)                                           //打印写入的字节数
		b := make([]byte, 20)
		f.Seek(0, 0) 											//指针返回到0
		data, _ := f.Read(b)
		fmt.Println(string(b[:data])) 							//输出了packagehelloword
	}
---
###(10)(f *File) WriteAt()在偏移位置多少的地方写入，函数原型是func (f *File) WriteAt(b []byte, off int64) (n int, err error)返回值是一样的
	import (
		"fmt"
		"os"
	)
	
	func main() {
		f, _ := os.OpenFile("1.go", os.O_RDWR, os.ModePerm)
		f.WriteAt([]byte("widuu"), 10) //在偏移10的地方写入
		b := make([]byte, 20)
		d, _ := f.ReadAt(b, 10)    //偏移10的地方开始读取
		fmt.Println(string(b[:d])) //widuudhellowordhello
	}
---
###(11)(f *File).WriteString()这个很简单了，写入字符串函数原型func (f *File) WriteString(s string) (ret int, err error)返回值一样的了
	import (
		"fmt"
		"os"
	)
	
	func main() {
		f, _ := os.OpenFile("2.go", os.O_RDWR, os.ModePerm)
		n, _ := f.WriteString("hello word widuu")	//写入字符串
		fmt.Println(n)
		b := make([]byte, n)
		f.Seek(0, 0)				//一定要把偏移地址归0否则就一直在写入的结尾处
		c, _ := f.Read(b)
		fmt.Println(string(b[:c])) //返回hello word widuu
	}

>好了今天我们就讲完这些，明天我们继续讲os包



