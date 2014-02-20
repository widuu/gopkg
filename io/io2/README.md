#golang讲解（go语言）标准库分析之io(2)

>今天我们继续讲golang的io的标准库

####[1]type PipeReader

	type PipeReader struct {
	    // contains filtered or unexported fields
	}

#####(1)func Pipe() (*PipeReader, *PipeWriter)创建一个管道，并返回它的读取器和写入器，这个会在内存中进行管道同步，它的开启会io.Reader然后等待io.Writer的输入，没有内部缓冲，它是安全的调用Read和Write彼此和并行调用写

	import (
		"fmt"
		"io"
		"reflect"
	)
	
	func main() {
		r, w := io.Pipe()
		fmt.Println(reflect.TypeOf(r))	//*io.PipeReader
		fmt.Println(reflect.TypeOf(w))	//*io.PipeWriter
	}

---

#####(2)func (r *PipeReader) Close() error管道关闭后，正在进行或后续的写入Write操作返回ErrClosedPipe

	import (
		"fmt"
		"io"
	)
	
	func main() {
		r, w := io.Pipe()
		r.Close()
		_, err := w.Write([]byte("hello widuu")) //具体可以看下边的func (w *PipeWriter) Write(data []byte) (n int, err error)，当然也就这样用用的了，它是阻塞的后边具体介绍
	
		if err == io.ErrClosedPipe {
			fmt.Println("管道已经关闭无法写入") //管道已经关闭无法写入
		}
	}

---

#####(3)func (r *PipeReader) CloseWithError(err error) error这个就是上边的r.Close关闭的时候，写入器会返回错误的信息

	import (
		"errors"
		"fmt"
		"io"
	)
	
	func main() {
		r, w := io.Pipe()
		r.Close()
		err := errors.New("管道符关闭了") //errors这个包我们前边已经说过了，就一个方法New不会的可以看看前边的
		r.CloseWithError(err)
		_, err = w.Write([]byte("test"))
		if err != nil {
			fmt.Println(err) //管道符关闭了
		}
	}

---

#####(4)func (r *PipeReader) Read(data []byte) (n int, err error)标准的阅读接口，它从管道中读取数据、阻塞一直到一个写入接口关闭，如果写入端发生错误，它就会返回错误，否则返回的EOF

	import (
		"fmt"
		"io"
	)
	
	func main() {
		r, w := io.Pipe()
		go w.Write([]byte("hello widuu"))
		d := make([]byte, 11)
		n, _ := r.Read(d) //从管道里读取数据
		fmt.Println(string(d))
		fmt.Println(n)
	}

---

####[2]type PipeWriter

	type PipeWriter struct {
	    // contains filtered or unexported fields
	}

#####(1)func (w *PipeWriter) Close() error关闭管道，关闭时正在进行的Read操作将返回EOF，若管道内仍有未读取的数据，后续仍可正常读取

	import (
		"fmt"
		"io"
	)
	
	func main() {
		r, w := io.Pipe()
		go w.Write([]byte("hello word"))
	
		data := make([]byte, 10)
		n, err := r.Read(data)
		w.Close()
		if err == io.EOF {
			fmt.Println("executing read return EOF")
			fmt.Println("executing read reads number", n)
		}
		n, _ = r.Read(data)
		fmt.Println(string(data))          //hello word
		fmt.Println("next read number", n) //next read number 0
	}

---

#####(2)func (w *PipeWriter) CloseWithError(err error) error这个函数和read里边的CloseWithError是大同小异的，关闭管道，关闭时正在进行的Read操作将返回参数传入的异常，若管道内仍有未读取的数据，后续仍可正常读取

	import (
		"errors"
		"fmt"
		"io"
	)
	
	func main() {
		r, w := io.Pipe()
		go w.Write([]byte("hello widuu"))
		newerr := errors.New("your daye 突然关闭了")
		w.CloseWithError(newerr)
		data := make([]byte, 10)
		_, err := r.Read(data)
		if err != nil {
			fmt.Println(err) //your daye 突然关闭了
		}
	}

---

#####(3)func (w *PipeWriter) Write(data []byte) (n int, err error)终于来打write了，这个是把字节切片写入管道，返回的是写入字节数和error,前边用到的太多了，随便哪一个吧

	import (
		"fmt"
		"io"
	)
	
	func main() {
		r, w := io.Pipe()
		go w.Write([]byte("hello widuu")) //写入的是[]byte,注意官方文档写的是，写入管道阻塞，一直到所有数据的读取结束
		data := make([]byte, 11)
		n, _ := r.Read(data)
		fmt.Println(string(data))     //hello widuu
		fmt.Println("read number", n) //read number 10
	}

每天只讲一点golang的标准库，方便大家学习和使用，更多的时候去理解标准库，大家多动手，如果你喜欢请继续关注我们[微度网络-网络技术支持中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)