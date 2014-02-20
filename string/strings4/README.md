#golang讲解（go语言）标准库分析之strings结束篇

今天我们完结了这个strings的包，其实我们就剩下了type Reader和type Replacer这个我们之间讲过io的包，这样大家理解起来就比较省劲了！

#####(1)func NewReader(s string) *Reader通过读取一个字符串之后返回Reader对象,然后实现io.Reader, io.ReaderAt, io.Seeker, io.ByteScanner, 和io.RuneScanner 接口

	import (
		"fmt"
		"reflect"
		"strings"
	)
	
	func main() {
		reader := strings.NewReader("widuu web")
		fmt.Println(reflect.TypeOf(reader)) //*strings.Reader
		fmt.Println(reader.Len())           //9
	}

---

#####(2)func (r *Reader) Len() int这个函数很简单就是实现上边的接口的读取的字符串的数，上边已经写了，哈哈这里就省了哈

---

#####(3)func (r *Reader) Read(b []byte) (n int, err error),读取数据到b中，返回读取的实际大小n，如果出错返回err，例如EOF或者b的长度为0

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		reader := strings.NewReader("hello widuu web")
		n, err := reader.Read(make([]byte, 10))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n) //10
	}

---

#####(4)func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)这个跟我们前边写的是一样的都是根据偏移量来读取数据的

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		reader := strings.NewReader("hello widuu web")
		b := make([]byte, 10)
		n, err := reader.ReadAt(b, 2)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b[:n])) //llo widuu
	}

---

#####(5)func (r *Reader) ReadByte() (b byte, err error),这个估计我都不用说大家都知道就是读取一个byte数据，然后返回的直接就是byte

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		reader := strings.NewReader("hello widuu web")
		data, err := reader.ReadByte()
		if err == nil {
			fmt.Println(string(data))
			return
		}
		fmt.Println(err) //返回一个字节h
	}

---

#####(6)func (r *Reader) ReadRune() (ch rune, size int, err error)这个同上只不过返回的是一个rune类型的

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		reader := strings.NewReader("hello widuu web")
		r, n, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}
	
		fmt.Println(string(r), n) //h,1
	}

---

#####(7)func (r *Reader) Seek(offset int64, whence int) (int64, error)根据whence来移动offset,改变指针的

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		reader := strings.NewReader("hello widuu")
		b := make([]byte, 8)
		n, _ := reader.Read(b)
		fmt.Println(string(b[:n])) //hello wi 这个时候没有读到要读d
		reader.Seek(2, 1)          //这个时候我们在读取位置 偏移2
		n1, _ := reader.Read(b)
		fmt.Println(string(b[:n1])) //u
	}

#####(8)func (r *Reader) UnreadByte() error当前读取的位置向前移一个byte

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		reader := strings.NewReader("hello widuu")
		b := make([]byte, 8)
		n, _ := reader.Read(b)
		fmt.Println(string(b[:n])) //hello wi 这个时候没有读到要读d
		reader.Seek(2, 1)          //这个时候我们在读取位置 偏移2
		reader.UnreadByte()		   //然后向前偏移1字节
		n1, _ := reader.Read(b)
		fmt.Println(string(b[:n1])) //uu
	}

---

#####(9)func (r *Reader) UnreadRune() error这个一样只不过类型变了而已，参照上边的

---

#####(10)func (r *Reader) WriteTo(w io.Writer) (n int64, err error)这个借口主要继承了io.Write的功能所以我们可以这样做

	import (
		"fmt"
		"os"
		"strings"
	)
	
	func main() {
		reader := strings.NewReader("hello widuu")
		w, _ := os.Create("test.txt")
		defer w.Close()
		n, err := reader.WriteTo(w)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n) //11
	}

---

#####(11)func NewReplacer(oldnew ...string) *Replacer NewReplacer从列表中返回一个新的替换后的字符串,oldnew是slice

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
		fmt.Println(r.Replace("This is <b>HTML</b>!")) //This is &lt;b&gt;HTML&lt;/b&gt;!
	}

---

#####(2)func (r *Replacer) Replace(s string) string这个上边已经写出函数了，就是对字符串进行替换，大家可以参考上边的程序

---

#####(3)func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)函数讲字符串替换后，然后写入w中

	import (
		"fmt"
		"io/ioutil"
		"os"
		"strings"
	)
	
	func main() {
		r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
		w, _ := os.OpenFile("test1.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		defer w.Close()
		n, err := r.WriteString(w, "This is <b>HTML</b>!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
		d, _ := ioutil.ReadFile("test1.txt") //32
		fmt.Println(string(d))               //This is &lt;b&gt;HTML&lt;/b&gt;!
	}

每天只讲一点golang的标准库，方便大家学习和使用，更多的时候去理解标准库，大家多动手，如果你喜欢请继续关注我们[微度网络-网络技术支持中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)