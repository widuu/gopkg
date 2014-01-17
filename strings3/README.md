#golang讲解（go语言）标准库分析之strings（2）

>今天我们继续哈，争取我们把strings包这个东西给弄完了，这已经也有两天没有好好更新了，所以这两天更新的比较多补充前两天的

#####(1)func Title(s string) string这个函数作用很简单，就是把输入的字符串首字母大写，参数是字符串返回的是字符串

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Title("hello word widuu write"))
	}

---

#####(2)func ToLower(s string) string这个函数是把字符串转换成小写，当然都是副本哦

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.ToLower("WIDUU")) //widuu
	}

---

#####(3)func ToTitleSpecial(_case unicode.SpecialCase, s string) string该函数把s字符串里面的每个单词转化为标题体，但是调用的是unicode.SpecialCase的ToTitle方法

	import (
	    "fmt"
	    "strings"
	    "unicode"
	)
	
	func main() {
	    var SC unicode.SpecialCase
	    fmt.Println(strings.ToTitleSpecial(SC, "Gopher"))
	    //GOPHER
	}

---

#####(4)func ToUpper(s string) string该函数就是把所有的字符串都大写

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.ToUpper("widuu")) //WIDUU
	}

---

#####(5)func ToUpperSpecial(_case unicode.SpecialCase, s string) string该函数把s字符串里面的每个单词转化为大写，但是调用的是unicode.SpecialCase的ToUpper方法

	import (
	    "fmt"
	    "strings"
	    "unicode"
	)
	
	func main() {
	    var SC unicode.SpecialCase
	    fmt.Println(strings.ToUpperSpecial(SC, "Gopher"))
	    //GOPHER
	}

---

#####(6)func Trim(s string, cutset string) string这个函数大家估计就看的多了，这个函数是去除两边的自定义字符的

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Trim("&&&&&nihao&&&&", "&")) //nihao
	}

---

#####(7)func TrimFunc(s string, f func(rune) bool) string该函数大家是不是看的多了，对就是自定义函数清楚，那我们自定义函数完成上边那个函数的效果吧

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.TrimFunc("&&&&nihao&&&&", func(r rune) bool {
			if r == '&' {
				return true
			}
			return false
		})) //nihao
	}

---

#####(8)func TrimLeft(s string, cutset string) string该函数是删除左边的定义的字符

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.TrimLeft("&widuu&", "&")) //widuu&
	}

#####(9)func TrimLeftFunc(s string, f func(rune) bool) string这个你说我还用再多说吗？上边遇到这种类型的是不是多的吓人，写的都要吐了，好吧根据自定义函数来删除左边

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.TrimLeftFunc("&widuu&", func(r rune) bool {
			if r == '&' {
				return true
			}
			return false
		})) //	widuu&是不是跟上边的那个函数效果一样了
	}

#####(10)func TrimPrefix(s, prefix string) string前边我们讲了HasPrefix判断前缀的，这个是去除前缀的，来我们see下

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		s := "widuu_xiaowei"
		if strings.HasPrefix(s, "widuu_") {
			fmt.Println(strings.TrimPrefix(s, "widuu_")) //xiaowei
			return
		}
	
		fmt.Println(s)
	}

#####(11)

 - func TrimRight(s string, cutset string) string
 - func TrimRightFunc(s string, f func(rune) bool) string
 跟上边的left一样，只是这个是右边而已

#####(12)func TrimSpace(s string) string,清楚文本里边的空白操作\r\n\t

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.TrimSpace("\r\t\n widuu \t")) //widuu
	}

#####(13)func TrimSuffix(s, suffix string) string这个函数去除后缀的，跟TrimPrfix这个正好相反，我们看实例

	import (
		"fmt"
		"strings"
	)
	
	func main() {
		s := "web_widuu"
		if strings.HasSuffix(s, "_widuu") {
			fmt.Println(strings.TrimSuffix(s, "_widuu")) //web
			return
		}
		fmt.Println(s)
	
	}


每天只讲一点golang的标准库，方便大家学习和使用，更多的时候去理解标准库，大家多动手，如果你喜欢请继续关注我们[微度网络-网络技术支持中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)