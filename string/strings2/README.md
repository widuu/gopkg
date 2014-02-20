#golang讲解（go语言）标准库分析之strings（2）

>其实我也不明白是什么让我坚持下来一点一点的分析pkg包，但是分析的时候我感觉还是很高兴的！今天我们继续strings包

#####(1)func Index(s, sep string) int 这个函数是查找字符串，然后返回当前的位置，输入的都是string类型，然后int的位置信息

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Index("widuu", "i")) //1
		fmt.Println(strings.Index("widuu", "u")) //3
	}

---

#####(2)func IndexAny(s, chars string) int 这个函数是一样的查找，字符串第一次出现的位置，如果不存在就返回-1

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.IndexAny("widuu", "u")) //3
	}

---

#####(3)func IndexByte(s string, c byte) int,这个函数功能还是查找第一次粗线的位置，只不过这次C是byte类型的，查找到返回位置，找不到返回-1

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.IndexByte("hello xiaowei", 'x')) //6
	}

---

#####(4)func IndexRune(s string, r rune) int，还是查找位置，只不过这次是rune类型的

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.IndexRune("widuu", rune('w'))) //0
	}

---

#####(5)func IndexFunc(s string, f func(rune) bool) int这个函数大家一看就知道了，是通过类型的转换来用函数查找位置，我们来代码看下哈

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.IndexFunc("nihaoma", split)) //3
	}
	
	func split(r rune) bool {
		if r == 'a' {
			return true
		}
		return false
	}

---

#####(6)func Join(a []string, sep string) string,这个跟php中的implode差不多，这个函数是将一个[]string的切片通过分隔符，分割成一个字符串

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		s := []string{"hello", "word", "xiaowei"}
		fmt.Println(strings.Join(s, "-")) //	hello-word-xiaowei
	}

---

#####(7)func LastIndex(s, sep string) int 看到这个大家可能也明白了查找的是最后出现的位置，正好跟index相反

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.LastIndex("widuu", "u")) //	4
	}

---

#####(8)func LastIndexAny(s, chars string) int这个跟indexAny正好相反，也是查找最后一个

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.LastIndexAny("widuu", "u")) //	4
	}

---

#####(9)func LastIndexFunc(s string, f func(rune) bool) int,这个和IndexFunc正好相反，通过函数查找最后一个位置

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.LastIndexFunc("nihaoma", split)) //6
	}
	
	func split(r rune) bool {
		if r == 'a' {
			return true
		}
		return false
	}

---

#####(10)func Map(mapping func(rune) rune, s string) string,该函数以此读取s中的字符，传入mapping函数，然后返回的字符链接起来，说白了就是字符串的每一个字符通过mapping函数的处理，最后返回处理好的字符串，如果处理不正确，那么就抛弃该字符

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Map(func(r rune) rune {
			if r == 'a' {
				return 'i'
			}
			return r
		}, "hello waduu")) //hello widuu
	}

---

#####(11)func Repeat(s string, count int) string,这个函数就是输出几个重复的字符串，count代表重复的次数

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Repeat("widuu", 3)) //widuuwiduuwiduu
	}

---

#####(12)func Replace(s, old, new string, n int) string,这个函数是把一个字符串中的字符替换成你定义的字符，n是替换的数值，如果n<0就是替换全部，返回的知识一个副本

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		s := "widuuweb"
		fmt.Println(strings.Replace(s, "u", "a", 1))  //widauweb
		fmt.Println(strings.Replace(s, "u", "a", -1)) //widaaweb
		fmt.Println(s)                                //widuuweb
	}

---

#####(13)func Split(s, sep string) []string,有join就有Split这个就是把字符串按照指定的分隔符切割成slice

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Split("a,b,c,d,e", ",")) //[a b c d e]
	}

---

#####(15)func SplitAfter(s, sep string) []string,这个函数是在前边的切割完成之后再后边在加上sep分割符

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.SplitAfter("a,b,c,d", ",")) //[a, b, c, d]
	}

---

#####(16)func SplitAfterN(s, sep string, n int) []string该函数s根据sep分割，返回分割之后子字符串的slice,和split一样，只是返回的子字符串保留sep，如果sep为空，那么每一个字符都分割

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.SplitAfterN("a,b,c,d,r", ",", 4)) //["a," "b," "c," "d,r"]
		fmt.Println(strings.SplitAfterN("a,b,c,d,r", ",", 5)) //["a," "b," "c," "d," "r"]
	}

---

#####(17)func SplitN(s, sep string, n int) []string,这个是切割字符串的时候自己定义长度，如果sep为空，那么每一个字符都分割

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.SplitN("a,b,c", ",", 2)) //[a b,c]
	}

每天只讲一点golang的标准库，方便大家学习和使用，更多的时候去理解标准库，大家多动手，如果你喜欢请继续关注我们[微度网络-网络技术支持中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)