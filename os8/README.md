#golang讲解（go语言）标准库分析之os(8)

>今天我们继续os包，不多说话了就说一句是golang的os包

####(1)type FileInfo里边就两个函数Stat和Lstat这个我就不多说了，使用方法相同前边咱们也介绍了

	import (
		"fmt"
		"os"
	)
	
	func main() {
		//这个和stat差不多
		f, _ := os.Lstat("1.go")
		fmt.Println(f) //	&{1.go {32 {2883792444 30345044} {2884358853 30345044} {3464863792 30345293} 0 43} {0 0} D:\test\1.go 0 0 0}
	}
	
	type FileMode
		func (m FileMode) IsDir() bool
	    func (m FileMode) IsRegular() bool
	    func (m FileMode) Perm() FileMode
	    func (m FileMode) String() string

---

####(2)这个FileMode主要是判断和输出权限用的，这个没有什么多讲的所以我就直接给大家测试一下代码

	import (
		"fmt"
		"os"
		"reflect"
	)
	
	func main() {
		f, _ := os.Stat("1.go")
		fi := f.Mode()
		fmt.Println(reflect.TypeOf(fi)) //os.FileMode
		fmt.Println(fi.IsDir())         //false 判断是否是目录
		fmt.Println(fi.IsRegular())     //true  判断是否是个常规文件
		fmt.Println(fi.Perm())          //-rw-rw-rw-	返回的unix浅显
		fmt.Println(fi.String())        //-rw-rw-rw- 	返回文件模式的字符串形式
	}

---

####(3)func (e *LinkError) Error() string 这个函数是获取LinkError的字符串形式，首先我们看下结构

	type LinkError struct {
	    Op  string
	    Old string
	    New string
	    Err error
	}
	
	import (
		"errors"
		"fmt"
		"os"
	)
	
	func main() {
		linkerror := &os.LinkError{
			Op:  "widuu",
			Old: "old",
			New: "new error",
			Err: errors.New("error test"),
		}
		fmt.Println(linkerror.Error())  //widuu old new error: error test
	}

---

####(4)func (e *PathError) Error() string 这个函数是获取PathError的字符串形式，然后我们看下它的结构
	
	type PathError struct {
	    Op   string
	    Path string
	    Err  error
	}
	import (
		"errors"
		"fmt"
		"os"
	)
	
	func main() {
		patherr := &os.PathError{
			Op:   "dir",
			Path: "widuu",
			Err:  errors.New("no path"),
		}
		fmt.Println(patherr.Error()) //dir widuu: no path
	}

---

####(5)type Process我们讲这个结构下的方法

	type Process struct {
	    Pid int
	}

>[1]FindProcess()函数原型func FindProcess(pid int) (p *Process, err error)输入进程的pid返回的是进程的指针结构和error信息

	import (
		"fmt"
		"os"
	)
	
	func main() {
		p, _ := os.FindProcess(6816) //现在我本机windows下的cmd
		fmt.Println(p)               //&{6816 236 0}
	}

>[2]os.StartProcess()函数原型func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)函数是新建一个进程，必须输入名称和参数还待指定环境，这是一个低级的接口，等我们讲os/exec的时候里边有更多的高级接口，返回的是进程的结构和PathError


	import (
		"fmt"
		"os"
	)
	
	func main() {
		attr := &os.ProcAttr{
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},  //这里注意哈
		}
		p, err := os.StartProcess(`c:\windows\system32\notepad.EXE`, []string{`c:\windows\system32\notepad.EXE`, `d:/1.txt`}, attr) //windows打开文件夹下1.txt
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(p)  //&{5488 240 0}
	}

---

>[3]func (p *Process) Kill() error 一看我们就知道是杀死进程，我们试验下

	import (
		"fmt"
		"os"
	)
	
	func main() {
		p, _ := os.FindProcess(7588) //现在我本机windows下的cmd
		fmt.Println(p)               //	&{7588 224 0}
		err := p.Kill()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("杀死进程")
	}


***很郁闷的事情是~~~我正用着这个进程godoc -http:=8080 还待重新来一次***
---


>[4]func (p *Process) Release() error这个是释放进程的资源的还拿上边的做实验

	import (
		"fmt"
		"os"
	)
	
	func main() {
		attr := &os.ProcAttr{
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		}
		p, err := os.StartProcess(`c:\windows\system32\notepad.EXE`, []string{`c:\windows\system32\notepad.EXE`, `d:/1.txt`}, attr) //windows打开文件夹下1.txt
		if err != nil {
			fmt.Println(err.Error()) //	&{7704 244 0}
		}
		fmt.Println(p)
		if err := p.Release(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

---

>[5]func (p *Process) Signal(sig Signal) error这个就是给一个进程发送信号用的

	import (
		"fmt"
		"os"
	)
	
	func main() {
		attr := &os.ProcAttr{
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		}
		p, _ := os.StartProcess(`c:\windows\system32\notepad.EXE`, []string{`c:\windows\system32\notepad.EXE`, `d:/1.txt`}, attr)
		fmt.Println(p)
		if err := p.Signal(os.Kill); err != nil { //向系统发送一个干掉他的信号 所以我们熟悉的记事本不出来了
			fmt.Println(err)
		}
	}

---

>[6]func (p *Process) Wait() (*ProcessState, error) 一看就知道了，这个函数是等待函数的进程完成，然后返回进程结构体指针的了

	import (
		"fmt"
		"os"
	)
	
	func main() {
		attr := &os.ProcAttr{
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		}
		p, _ := os.StartProcess(`c:\windows\system32\notepad.EXE`, []string{`c:\windows\system32\notepad.EXE`, `d:/1.txt`}, attr)
		pw, _ := p.Wait()
		fmt.Println(pw) //exit status 0
	}

---

####(6)我们讲述type ProcessState struct {}这个结构和里边的方法

>[1]func (p *ProcessState) Exited() bool查看进程是否退出

	import (
		"fmt"
		"os"
	)
	
	func main() {
		attr := &os.ProcAttr{
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		}
		p, _ := os.StartProcess(`c:\windows\system32\notepad.EXE`, []string{`c:\windows\system32\notepad.EXE`, `d:/1.txt`}, attr)
		fmt.Println(p)           //&{7764 252 0}
		pw, _ := p.Wait()
		fmt.Println(pw.Exited()) //当我关闭记事本的时候进程结束，所以就算true
	}

>[2]func (p *ProcessState) Pid() int                  很简单返回进程的pid,剩下的我就不一一说了直接说明了然后做实验

>[3]func (p *ProcessState) String() string            这个函数主要是获取进程状态的字符串

>[4]func (p *ProcessState) Success() bool             这个函数报告是否退出

>[5]func (p *ProcessState) Sys() interface{}          这个函数主要是获取进程的退出信息

>[6]func (p *ProcessState) SysUsage() interface{}     这个函数主要是获取进程资源使用情况

>[7]func (p *ProcessState) SystemTime() time.Duration 这个函数主要是获取的进程的系统cpu使用时间

>[8]func (p *ProcessState) UserTime() time.Duration   这个函数主要是获取的进程的用户cpu使用时间

	import (
		"fmt"
		"os"
	)
	
	func main() {
		attr := &os.ProcAttr{
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		}
		p, _ := os.StartProcess(`c:\windows\system32\notepad.EXE`, []string{`c:\windows\system32\notepad.EXE`, `d:/1.txt`}, attr)
		fmt.Println(p) 			    //&{44 252 0}
		pw, _ := p.Wait()
		fmt.Println(pw.Pid())       //44
		fmt.Println(pw.String())    //exit status 0
		fmt.Println(pw.Success())   //true
		fmt.Println(pw.Sys())       //{0}
		fmt.Println(pw.SysUsage())	//&{{1728533532 30345488} {1745163891 30345488} {1562500 0} {1562500 0}}
		fmt.Println(pw.SystemTime())//156.25ms
		fmt.Println(pw.UserTime())	//156.25ms
	}

---

####(7)func (e *SyscallError) Error() string 这个函数主要是获取SyscallError的字符串形式

	type SyscallError struct {
	    Syscall string
	    Err     error
	}
	import (
		"errors"
		"fmt"
		"os"
	)
	
	func main() {
		err := &os.SyscallError{
			Syscall: "widuu",
			Err:     errors.New("syscall error"),
		}
		fmt.Println(err.Error()) //widuu: syscall error
	}

***今天的os包就完了，明天我们继续讲os的子包，到时候多谢大家关注我的博客[微度网络|网络技术支持中心](http://www.widuu.com)http://www.widuu.com***