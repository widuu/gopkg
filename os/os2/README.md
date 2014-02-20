>今天我们继续昨天的继续分析golang的os包但是今天由于我的时间比较紧所以可能分析的比较少，但是还是一样的味道！

####第一个我们分析的是os.Chtime()这个包，函数的原形是func Chtimes(name string, atime time.Time, mtime time.Time) error 输入string的文件的名称 访问时间 创建时间 返回的是error接口信息

	import (
		"fmt"
		"os"
		"time"
	)

	func main() {
		err := os.Chtimes("2.go", time.Now(), time.Now())  //改变时间
		if err != nil {
			fmt.Println(err)
		}
		fi, _ := os.Stat("2.go")
		fmt.Println(fi.ModTime())   //输出时间 2013-12-29 20:46:23.0005257 +0800 +0800
	}

---	
	
####第二个我们说的是os.Environ()获取系统的环境变量，函数原形是func Environ() []string返回是环境变量的[]string切片，说道这个就要和其他的一起说明了，那就是os.ClearEnv()清空环境变量

	func main() {
		data := os.Environ() //输出之前的环境变量 APPDATA=C:\Users\xiaolvge\AppData\Roaming CLASSPATH=.;D:\java\jdk1.6.0_38…………
		fmt.Println(data)
		os.Clearenv() //清空环境变量
		data = os.Environ()
		fmt.Println(data) //输出[]string类型的切片[]
	}

---
	
####第三个说的是os.Exit()就是中断程序返回自定义的int类型的代码，函数运行是func Exit(code int)输入一个int的值就可以了
	import (
		"fmt"
		"os"
	)

	func main() {
		func() {
			for {
				fmt.Println("这个是匿名函数")
				os.Exit(1)    //输出exit status 1中断操作
			}
		}()
	}
	
---

####第四个函数os.Expand()这个其实就是一个回调函数替换的方法，函数的原形是func Expand(s string, mapping func(string) string) string 输入的是一个string。对应的是func(string)string的替换字符串的方法，如果没有字符就替换为空

	import (
		"fmt"
		"os"
	)

	func main() {
		mapping := func(s string) string {
			m := map[string]string{"widuu": "www.widuu.com", "xiaowei": "widuu"}
			return m[s]
		}
		data := "hello $xiaowei blog address $widuu"
		fmt.Printf("%s", os.Expand(data, mapping)) //输出hello widuu blog address www.widuu.com
	}
	
---

####第五个函数跟第四个差不多但是很简单,os.ExpandEnv()把字符串的s替换成环境变量的内容，函数的原形是func ExpandEnv(s string) string，输入的当然是要替换的字符，输出的当然还是字符串了

	import (
		"fmt"
		"os"
	)

	func main() {
		data := "GOBIN PATH $GOBIN"
		fmt.Println(os.ExpandEnv(data)) //输出我本地的环境变量的GOBIN的地址GOBIN PATH C:\Go\bin
	}
	
---

####第六个是os.Hostname()这个函数看字面的思意就懂了，是返回主机的HostName(),函数的原形是func Hostname() (name string, err error)返回主机名字和一个error的接口信息
	import (
		"fmt"
		"os"
	)

	func main() {
		data, _ := os.Hostname()
		fmt.Println(data) //我是windows环境下返回我的win的主机名 WIDUU
	}
	
---

>今天我们就讲到这里如果您喜欢我们，请关注我们