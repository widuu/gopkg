#golang讲解（go语言）标准库分析之os/user os/signal (os包完结篇)

>今天我们讲golang os包的最后两个包user、signal

##os/user包
####(1)type UnknownUserError string其中这个里边有一个方法`func (e UnknownUserError) Error() string`返回木有用户的错误信息

	import (
		"fmt"
		"os/user"
	)
	
	func main() {
		var err user.UnknownUserError
		err = "这个我自己设定"
		fmt.Println(err.Error()) //user: unknown user 这个我自己设定
	}

---

####(2)type UnknownUserIdError int里边有一个方法`func (e UnknownUserIdError) Error() string`返回的依然是木有用户的错误信息

	import (
		"fmt"
		"os/user"
	)
	
	func main() {
		var err user.UnknownUserIdError
		err = 2
		fmt.Println(err.Error()) //user: unknown userid 2
	}

---

####(3)type User我们先看下结构是什么?当然User有个BUG我们会看到


	type User struct {
		    Uid      string // 用户id
		    Gid      string // 用户组id
		    Username string //用户名
		    Name     string //用户全名
		    HomeDir  string //用户家目录
		}
>[1]func Current() (*User, error)返回当前用户的信息

	import (
		"fmt"
		"os/user"
	)
	
	func main() {
		/*
			type User struct {
			    Uid      string // 用户id
			    Gid      string // 用户组id
			    Username string //用户名
			    Name     string //用户全名
			    HomeDir  string //用户家目录
			}
		*/
		user, _ := user.Current()
		//我是windows平台
		fmt.Println(user.Gid)      //S-1-5-21-955939588-3462822645-4196941772-513
		fmt.Println(user.HomeDir)  //C:\Users\Administrator
		fmt.Println(user.Uid)      //S-1-5-21-955939588-3462822645-4196941772-500
		fmt.Println(user.Name)     //nil
		fmt.Println(user.Username) //widuu\Administrator
	}

---

>[2]先看BUG Lookup and LookupId functions do not set Gid and HomeDir fields in the User struct returned on windows. 这个是windows平台上的错误，找不到HomeDir的错误和没有设置Gid我们看下，第一个func Lookup(username string) (*User, error)，我是windows平台写的看下代码和返回

	import (
		"fmt"
		"os/user"
	)
	
	func main() {
		user, err := user.Lookup("widuu\\Administrator")
		if err != nil {
			/*
			  我是windows平台 返回Unknown directory为什么呢？在标准库里我们可以看到有个BUG是这样写的
			  Lookup and LookupId functions do not set Gid and HomeDir fields in the User struct returned on windows.
			  这里我们就明白了，widnows唉唉唉~~~~
			*/
			fmt.Println(err)
		}
		fmt.Println(user.HomeDir)
	}

---

>[3]func LookupId(uid string) (*User, error)根据uid返回用户的信息，依然是windows平台错误，我放出代码自己测试吧

	import (
		"fmt"
		"os/user"
	)
	
	func main() {
		user, err := user.LookupId("S-1-5-21-955939588-3462822645-4196941772-500")
		if err != nil {
			fmt.Println(err) //Unknown directory 一样的windows bug
		}
		fmt.Println(user.HomeDir)
	}

---

##os/signal

####(1)这个包就有两个方法是传递信号第一个func Notify(c chan<- os.Signal, sig ...os.Signal)

	import (
		"fmt"
		"os"
		"os/signal"
	)
	
	func main() {
		//设置一个channel来发送信号
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
	
		// 一直运行一直到收到一个信号
	
		s := <-c
		
		fmt.Println("Got signal:", s)  //当我停止运行时 Got signal: interrupt
	
	}

---

>[2]func Stop(c chan<- os.Signal)这个函数是停止信号接收，给个代码让大家理解

	import (
		"fmt"
		"os"
		"os/signal"
	)
	
	func main() {
		//设置一个channel来发送信号
		c := make(chan os.Signal, 1)
		// 一直运行一直到收到一个信号
		signal.Notify(c, os.Interrupt, os.Kill)
	
		//终端信号的接收
		//signal.Stop(c)   //程序运行结过exit status 2
		s := <-c
	
		fmt.Println("Got signal:", s) //当我停止运行时 Got signal: interrupt
	
	}

*BUG  This package is not yet implemented on Plan 9.*

到这里我们的OS包就讲完了，然后我们继续讲其它的包，如果你喜欢请继续关注我们[微度网络-网络技术支持中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)



