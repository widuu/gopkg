>今天我们讲os/exec包，这个我就不废话了

####(1)func LookPath(file string) (f string, err error)这个是搜索可执行的二进制的文件的路径，返回的是执行路径和error
	
	import (
		"fmt"
		"os/exec"
	)

	func main() {
		f, err := exec.LookPath("php")
		if err != nil {
			fmt.Println("not install php")
		}
		fmt.Println(f) // 输出我的phpD:\PHP\\php.exe
	}

---

####(2)type Cmd

	type Cmd struct {
	    Path string
	    Args []string
	    Env []string
	    Dir string
	    Stdin io.Reader
	    Stdout io.Writer
	    Stderr io.Writer
	    ExtraFiles []*os.File
	    SysProcAttr *syscall.SysProcAttr
	    Process *os.Process
	    ProcessState *os.ProcessState
	}

---
	
>[1]func Command(name string, arg ...string) *Cmd 输入文件的路径，参数字符串，返回的是*Cmd的结构上边的结构

	import (
		"fmt"
		"os/exec"
	)
	
	func main() {
		argv := []string{"-a"}
		c := exec.Command("ls", argv...)
		d, _ := c.Output()
		fmt.Println(string(d)) //因为装的git bash所以可以用ls -a
		/*
		 *	.
		 *	..
		 *	command.go
		 *	lookpath.go
		 */
	}

---

>[2]func (c *Cmd) CombinedOutput() ([]byte, error)这个函数返回标准输出和错误


	import (
		"fmt"
		"os/exec"
	)
	
	func main() {
		arv := []string{"-a"}
		c := exec.Command("ls", arv...)
		d, err := c.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(d)) //输出和command里边的output一样哈
	}

---

>[3]func (c *Cmd) Output() ([]byte, error)我们第一次就使用了,返回标准的输出

	import (
		"fmt"
		"os/exec"
	)
	
	func main() {
		c, err := exec.Command("date").Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(c)) //	Sat Jan  4 17:07:36     2014 这个是标准库里的例子
	}

---

>[4]func (c *Cmd) Run() error这个函数主要是执行*Cmd中的命令，并且会等待命令执行完成，如果命令执行不成功，则返回错误信息

	import (
		"fmt"
		"os/exec"
	)
	
	func main() {
		c, err := exec.Command("date")
		if err != nil {
			fmt.Println(err)
		}
		err = c.Run()
		if err != nil {
			fmt.Println(err)
		}
		d, _ := c.Output()
		fmt.Println(string(d)) //	Sat Jan  4 17:07:36    
	}

---

>[5]func (c *Cmd) Start() error这个函数主要是执行*Cmd中的命令，只是让命令开始执行，并不会等待命令执行完。

>[6]func (c *Cmd) Wait() error 这个函数是等待函数执行完成

	import (
		"fmt"
		"os/exec"
	)
	
	func main() {
		//官方的标准库的例子
		cmd := exec.Command("sleep", "5") //执行等待5秒
		err := cmd.Start()                //开始执行
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		fmt.Printf("Waiting for command to finish...\n")
		err = cmd.Wait()						//wait下边的函数等待执行完成
		fmt.Printf("Command finished with error: %v\n", err)
	}

---

>[7]func (c *Cmd) StderrPipe() (io.ReadCloser, error)这个函数主要是用于连接到命令启动时错误标准输出的管道，命令结束时，管道会自动关闭

	import (
		"fmt"
		"io/ioutil"
		"os/exec"
	)
	
	func main() {
		c := exec.Command("mv", "hello")
		i, err := c.StderrPipe()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		if err = c.Start(); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		b, _ := ioutil.ReadAll(i)
		if err := c.Wait(); err != nil {
			fmt.Printf("Error: %s\n", err) //Error: exit status 1 mv: missing file argument Try `mv --help' for more information.
		}
		fmt.Println(string(b))
	}

---

>[8]func (c *Cmd) StdinPipe() (io.WriteCloser, error)这个函数主要是用于连接到命令启动时标准输入的管道
	
	import (
		"bytes"
		"fmt"
		"os/exec"
	)
	
	func main() {
		var output bytes.Buffer
		cmd := exec.Command("cat")
		cmd.Stdout = &output
		stdin, _ := cmd.StdinPipe()
		cmd.Start() //执行开始
		stdin.Write([]byte("widuu test"))
		stdin.Close()
		cmd.Wait()                                        //等待执行完成
		fmt.Printf("The output is: %s\n", output.Bytes()) //The output is: widuu test!
	}

---

>[9]有输入就有输出 func (c *Cmd) StdoutPipe() (io.ReadCloser, error)

	import (
		"fmt"
		"io/ioutil"
		"os/exec"
	)
	
	func main() {
		cmd := exec.Command("ls", "-ll")
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()
		d, _ := ioutil.ReadAll(stdout) //从管道里读取数据
		cmd.Wait()
		fmt.Println(string(d))
		/*
			total 5
			-rw-r--r--    1 Administ Administ      268 Jan  4 17:07 combinedoutput.go
			-rw-r--r--    1 Administ Administ      277 Jan  4 17:01 command.go
			-rw-r--r--    1 Administ Administ      209 Jan  4 16:46 lookpath.go
			-rw-r--r--    1 Administ Administ      241 Jan  4 17:07 output.go
			-rw-r--r--    1 Administ Administ      271 Jan  4 17:13 run.go
			-rw-r--r--    1 Administ Administ      438 Jan  4 17:24 start.go
			-rw-r--r--    1 Administ Administ      497 Jan  4 17:32 stderrpipe.go
			-rw-r--r--    1 Administ Administ      552 Jan  6 14:10 stdinpipe.go
			-rw-r--r--    1 Administ Administ      235 Jan  6 14:14 stdoutpipe.go
		*/
	}

---

####(3)type Error结构是

	type Error struct {
	    Name string
	    Err  error
	}

>[1]func (e *Error) Error() string这个函数主要是输出命令执行失败的错误信息

	import (
		"errors"
		"fmt"
		"os/exec"
	)
	
	func main() {
		err := &exec.Error{
			Name: "widuu",
			Err:  errors.New("is not exists widuu"),
		}
	
		fmt.Println(err.Error()) //	exec: "widuu": is not exists widuu
	}

---

####(4)type ExitError 

	type ExitError struct {
	    *os.ProcessState
	}

>[1]func (e *ExitError) Error() string 这个函数主要是返回一个执行不成功命令的信息

	import (
	    "fmt"
	    "os/exec"
	)
	
	func main() {
	    cmd := exec.Command("mv", "Hello World!")
	    cmd.Run()
	    exitError := exec.ExitError{cmd.ProcessState}
	    fmt.Printf("The output is: %s\n", exitError.Error())   //The output is: exit status 1
	}

>这里我们exec包就讲完了，本文首发在我的博客，如果您喜欢可以随时关注我的博客**[微度网络-网络技术中心](http://www.widuu.com)[http://www.widuu.com](http://www.widuu.com)**