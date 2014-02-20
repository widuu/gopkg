#golang讲解（go语言）标准库分析之string开篇

>今天我们继续讲golang标准库的string这个包，我选定这个是看了下别的包真的需要这个功能所以我们来这个包

>string包实现了简单的函数方法来操作字符串。其中的type Reader很简单，因为会返回一个*Reader这就用到我们前边讲的io了，废话不多说进入正题

---

>(1)func Contains(s, substr string) bool这个函数是查找某个字符是否在这个字符串中存在，存在返回true

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Contains("widuu", "wi")) //true
		fmt.Println(strings.Contains("wi", "widuu")) //false
	}

---

>(2)func ContainsAny(s, chars string) bool这个是查询字符串中是否包含多个字符

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.ContainsAny("widuu", "w&d")) //true
	}

---

>(3)func ContainsRune(s string, r rune) bool,这里边当然是字符串中是否包含rune类型，其中rune类型是utf8.RUneCountString可以完整表示全部Unicode字符的类型

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.ContainsRune("widuu", rune('w'))) //true
		fmt.Println(strings.ContainsRune("widuu", 20))        //fasle
	}

---

>(4)func Count(s, sep string) int这个的作用就是输出，在一段字符串中有多少匹配到的字符

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Count("widuu", "uu")) //1
		fmt.Println(strings.Count("widuu", "u"))  //2
	}

---

>(5)func EqualFold(s, t string) bool这个是判断s,t两个字符串在完全小写的情况下是否相等uff8编码

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.EqualFold("WIDUU", "widuu")) //true
	}

---

>(6)func Fields(s string) []string，这个函数的作用是按照1：n个空格来分割字符串最后返回的是[]string的切片

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Fields("hello widuu golang")) //out 	[hello widuu golang]
	}

---

>(7)func FieldsFunc(s string, f func(rune) bool) []string一看就了解了，这就是根据自定义函数分割了

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.FieldsFunc("widuunhellonword", split)) //	[widuu hello word]根据n字符分割
	}
	
	func split(s rune) bool {
		if s == 'n' {
			return true
		}
		return false
	}

>(8)func HasPrefix(s, prefix string) bool，判断是否以什么字符串开头的

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.HasPrefix("widuu_web", "widuu")) //true
	}

---

>(9)func HasSuffix(s, suffix string) bool,判断是否有什么后缀，返回bool值

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.HasSuffix("nihaowiduu", "widuu")) //true
	}


每天只讲一点golang的标准库，方便大家学习和使用，更多的时候去理解标准库，大家多动手，如果你喜欢请继续关注我们[微度网络-网络技术支持中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)