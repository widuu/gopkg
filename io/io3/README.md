#golang讲解（go语言）标准库分析之io(3)

####[1]type Reader

	type Reader interface {
	    Read(p []byte) (n int, err error)
	}

#####(1)func LimitReader(r Reader, n int64) Reader，我们之前就说了Reader这个结构，其实这就是对Reader的一次封装，限定了它读取字节数，其实他实现的就是io.LimitedReader{}这个结构

	import (
		"fmt"
		"io"
		"os"
		"reflect"
	)
	
	func main() {
		f, _ := os.Open("test.txt")
		defer f.Close()
		reader := io.LimitReader(f, 5)
		p := make([]byte, 5)
		fmt.Println(reflect.TypeOf(reader)) //*io.LimitedReader
		var total int
		for {
			n, err := reader.Read(p)
			if err == io.EOF {
				fmt.Println("read value", string(p[:total])) //read value hello
				fmt.Println(total)                           //5
				break
			}
			total = total + n
		}
	
	}

---

#####(2)func MultiReader(readers ...Reader) Reader这个函数一看就知道是封装了多个readers,跟上边的方法差不多，只是封装了多个而已，当然还去除了读取的限制，我们代码给大家测试一下

	import (
		"fmt"
		"io"
		"os"
		"reflect"
	)
	
	func main() {
		f1, _ := os.Open("test1.txt")
		f2, _ := os.Open("test.txt")
		defer f1.Close()
		defer f2.Close()
		reader := io.MultiReader(f1, f2) //*io.multiReader
		fmt.Println(reflect.TypeOf(reader))
		p := make([]byte, 10)
		var total int
		var data string
		for {
			n, err := reader.Read(p)
			if err == io.EOF {
				fmt.Println("read end", total) //read end 17
				break
			}
			total = total + n
			data = data + string(p[:n])
		}
		fmt.Println("read value", data)  //read value widuu2hello widuu
		fmt.Println("read count", total) //	read count 17
	}

---

#####(3)既然上边介绍读了，我这介绍个写吧`type Write`func MultiWriter(writers ...Writer) Writer一样的作用只不过是这次换成写了

	import (
		"fmt"
		"io"
		"io/ioutil"
		"os"
	)
	
	func main() {
		f1, _ := os.Create("1.txt")
		f2, _ := os.Create("2.txt")
		writer := io.MultiWriter(f1, f2)
		writer.Write([]byte("widuu"))
		//千万别这么逻辑来 ，我这是测试用的哈
		r1, _ := ioutil.ReadFile("1.txt")
		r2, _ := ioutil.ReadFile("2.txt")
		fmt.Println(string(r1)) //widuu
		fmt.Println(string(r2)) //widuu
	}

---

#####(4)func TeeReader(r Reader, w Writer) Reader这个方法有意思是从r中读取数据然后写入到w中，这个没有内部缓冲区，看下代码

	import (
		"fmt"
		"io"
		"os"
		"reflect"
	)
	
	func main() {
		r, _ := os.Open("test.txt")
		w, _ := os.Create("test2.txt")
		reader := io.TeeReader(r, w)
		fmt.Println(reflect.TypeOf(reader)) //*io.teeReader
		p := make([]byte, 10)
		n, _ := reader.Read(p)
		fmt.Println(string(p[:n])) //hello widu
	}

---

####[2]type SectionReader{}

	type SectionReader struct {
	    // contains filtered or unexported fields
	}

#####(1)func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader，你一看就知道了，其实就是通过这个方法获取到io.SectionReader,第一个参数读取器，第二个参数偏移量，第三个参数是读取多少

	import (
		"fmt"
		"io"
		"os"
		"reflect"
	)
	
	func main() {
		f, _ := os.Open("test.txt")
		sr := io.NewSectionReader(f, 2, 5)
		fmt.Println(reflect.TypeOf(sr)) //*io.SectionReader
	}

---

#####(2)func (s *SectionReader) Read(p []byte) (n int, err error)熟悉的read()其实就是读取数据用的，大家看函数就可以理解了，因为咱们经常遇到这个上两个都写这个了~~

	import (
		"fmt"
		"io"
		"os"
	)
	
	func main() {
		f, _ := os.Open("test.txt")
		defer f.Close()
		sr := io.NewSectionReader(f, 2, 5)
		p := make([]byte, 10)
		n, err := sr.Read(p)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(p[:n])) //llo w
	}

---

#####(3)func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)额这个跟之前的ReadAt是一样的，只不过只有一个偏移量，少了截取数，但是你要知道SectionReader做的是什么就把数据截取了，所以就不需要截取数了

	import (
		"fmt"
		"io"
		"os"
	)
	
	func main() {
		f, _ := os.Open("test.txt")
		defer f.Close()
		sr := io.NewSectionReader(f, 2, 5)
		p := make([]byte, 10)
		n, err := sr.ReadAt(p, 1)
		if err == io.EOF {
			fmt.Println(string(p[:n])) //	lo w
		}
	
	}

#####(4)func (s *SectionReader) Seek(offset int64, whence int) (int64, error)这个是设置文件指针的便宜量的，之前我们的os里边也是有个seek的，对SectionReader的读取起始点、当前读取点、结束点进行偏移，offset 偏移量，whence 设定选项 0:读取起始点，1:当前读取点，2:结束点(不好用)，其他：将抛出Seek: invalid whence异常
	
	import (
		"fmt"
		"io"
		"os"
	)
	
	func main() {
		f, _ := os.Open("test.txt")
		defer f.Close()
		sr := io.NewSectionReader(f, 2, 5)
		p := make([]byte, 10)
		sr.Seek(1, 0) 			  //相当于起始的地址偏移1
		n, err := sr.Read(p)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(p[:n])) //lo w 是不是达到了前边的ReadAt()
	}

---

#####(5)func (s *SectionReader) Size() int64返回的是可以读取的字节数，这个不受偏移指针的影响，也不受当前读取的影响，我们具体看下代码


	import (
		"fmt"
		"io"
		"os"
	)
	
	func main() {
		f, _ := os.Open("test.txt")
		defer f.Close()
		sr := io.NewSectionReader(f, 2, 5)
		fmt.Println(sr.Size()) //5
		p := make([]byte, 10)
		sr.Seek(1, 0) 			//相当于起始的地址偏移1
		n, err := sr.Read(p)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(p[:n])) //lo w
		fmt.Println(sr.Size())     //5
	}


今天我们就讲完了io的标准库了，明天讲哪个呢？还是关注明天吧哈！每天只讲一点golang的标准库，方便大家学习和使用，更多的时候去理解标准库，大家多动手，如果你喜欢请继续关注我们[微度网络-网络技术支持中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)