#golang讲解（go语言）标准库分析之io(1)

>今天我们开新的标准库io也是继续上一次我们没有讲完的标准库，上一次我们只说了io/ioutil这次我们继续开讲整个io库

#####(1)`func Copy(dst Writer, src Reader) (written int64, err error)`这个函数是从一个文件读取拷贝到另外一个文件，一直拷贝到读取文件的EOF，所以不会返回io.EOF错误，参数是写入目标器和读取目标器，返回int64的拷贝字节数和err信息

	import (
		"fmt"
		"io"
		"os"
	)
	
	func main() {
		r, _ := os.Open("test.txt")
		w, _ := os.Create("write.txt")
		num, err := io.Copy(w, w)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(num) //返回int64的11 打开我的write.txt正是test.txt里边的hello widuu
	}	

---

#####(2)`func CopyN(dst Writer, src Reader, n int64) (written int64, err error)`看函数就知道了跟上述的是一样的，只是多加了一个读取数的限制，然后我们看下代码

	import (
		"fmt"
		"io"
		"io/ioutil"
		"os"
	)
	
	func main() {
		r, _ := os.Open("test.txt")
		w, _ := os.Create("write1.txt")
		num, err := io.CopyN(w, r, 5)
		if err != nil {
			fmt.Println(err)
		}
		defer r.Close()
		b, _ := ioutil.ReadFile("write1.txt")
		fmt.Println(string(b)) //输出 hello
		fmt.Println(num)       //5
	}

---

#####(3)`func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)`这个函数就是从读取器中读取数据放到我们的buf中，限定了最小的读取字节数，如果我们读取的数据小于最小读取器，譬如你设定min的值是8，但是你读取的数据字节数是5就会返回一个`io.ErrUnexpectedEOF`,如果大于就会返回`io.ErrShortBuffer`,读取完毕会有`io.EOF`~~，多讲一些哈，这个Reader只要我们满足这个interface就可以用这个
	type Reader interface {
	    Read(p []byte) (n int, err error)
	}
>我们讲os包的时候就发现了，其中*File就支持`func (f *File) Read(b []byte) (n int, err error)`，不懂的大家可以回头看看

	import (
		"fmt"
		"io"
		"os"
	)
	
	func main() {
		r, _ := os.Open("write1.txt")
		b := make([]byte, 20)
		defer r.Close()
		var total int
		for {
			n, err := io.ReadAtLeast(r, b, 8)
			if err == nil {
				fmt.Println("Read enough value:", string(b)) //	Read enough value: hello widuu
			}
			if err == io.ErrUnexpectedEOF { //读取了的数据小于我们限定的最小读取数据8
				fmt.Println("Read fewer value:", string(b[0:n]))
			}
			
			if err == io.ErrShortBuffer{   //这个是我们设定的buf也就是b小于我们限定8
				fmt.Println("buf too Short")
				os.Exit(1)
			}
			if err == io.EOF { //读完了 输出
				fmt.Println("Read end total", total) //Read end total 11
				break
			}
			total = total + n
		}
	}

---

#####(4)`func ReadFull(r Reader, buf []byte) (n int, err error)`这个函数和上边的函数是相似，只不过是读取len(buf)个，放在buf中

	import (
		"fmt"
		"io"
		"os"
	)
	
	func main() {
		r, _ := os.Open("write.txt")
		b := make([]byte, 20)
		num, err := io.ReadFull(r, b)
		defer r.Close()
		if err == io.EOF {
			fmt.Println("Read end total", num)
		}
		if err == io.ErrUnexpectedEOF {
			fmt.Println("Read fewer value:", string(b[:num])) //Read fewer value: hello widuu，依然是buf长度大于读取的长度
			return
		}
	
		fmt.Println("Read  value:", string(b)) //如果b是5 就出现这里
	}

---

#####(5)`func WriteString(w Writer, s string) (n int, err error)`弄完读了，当然带要写了，这个函数主要是向写入目标中写入字符创，返回是写入的字节数还有error错误，主要是权限的错误，其中写入呀！都是writer这个结构就可以写入

	type Writer interface {
	    Write(p []byte) (n int, err error)
	}

>跟read一样我们的*File是有`func (f *File) Write(b []byte) (n int, err error)`，当然其实我们的*File中就已经有WirteString了`func (f *File) WriteString(s string) (ret int, err error)`

	import (
		"fmt"
		"io"
		"io/ioutil"
		"os"
	)
	
	func main() {
		w, _ := os.OpenFile("write1.txt", os.O_RDWR, os.ModePerm)
		n, err := io.WriteString(w, "ni hao ma")
		if err != nil {
			fmt.Println(err) //当我用os.open()的时候木有权限  悲催的~~输出write write1.txt: Access is denied.
		}
		defer w.Close()
		b, _ := ioutil.ReadFile("write1.txt")
		fmt.Println("write total", n) //write total 9
		fmt.Println(string(b))        // ni hao ma
	}

---

#####(6)今天最后一个我们讲一个结构`type LimitedReader`
	
	type LimitedReader struct {
	    R Reader // 读取器了
	    N int64  // 最大字节限制
	}

>[1]只实现了一个方法func (l *LimitedReader) Read(p []byte) (n int, err error)其实我们不难发现这个跟我们的ReadAtLast()就是亲兄弟的节奏

	import (
		"fmt"
		"io"
		"os"
	)
	
	func main() {
		reader, _ := os.Open("test.txt")
		limitedreader := io.LimitedReader{
			R: reader,
			N: 20,
		}
		p := make([]byte, 10)
		var total int
		for {
			n, err := limitedreader.Read(p)
			if err == io.EOF {
				fmt.Println("read total", total)     //read total 11
				fmt.Println("read value", string(p)) //read value hello widuu
				break
			}
			total = total + n
	
		}
	
	}

每天只讲一点golang的标准库，方便大家学习和使用，更多的时候去理解标准库，大家多动手，如果你喜欢请继续关注我们[微度网络-网络技术支持中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)