#golang讲解（go语言）标准库分析之reflect开山篇

>(1)`func Copy(dst, src Value) int`这个函数是将src中的数据副本拷贝到dst中，一直到dst装满或者src中没有数据了，src和dst必须是相同类型的数组或者slice,并且大家注意 src是Value类型，返回的拷贝的个数，

	type Value struct {
	    // contains filtered or unexported fields
	}

>所以 src和dst是结构中的数据

	import (
		"fmt"
		"reflect"
	)
	
	type Test struct {
		A []string
		B []string
	}
	
	func main() {
		var a Test
		a.A = append(a.A, []string{"a", "b", "c", "d"}...)
		a.B = append(a.B, "c", "d", "f", "g")
		n := reflect.Copy(reflect.ValueOf(a.B), reflect.ValueOf(a.A))
		fmt.Println(n)
		fmt.Println(a.B, a.A)
	}

---

>(2)`func DeepEqual(a1, a2 interface{}) bool` 这个函数对比两个类型是否完全相等，如果值相等并且底层类型也相等，那么这两个就相等返回true,否则就返回false

	import (
		"fmt"
		"reflect"
	)
	
	type Tz int
	
	func main() {
		var a int
		var b Tz
		a = 1
		b = 1
		fmt.Println(reflect.DeepEqual(a, b)) //false
		var c int
		c = 1
		fmt.Println(reflect.DeepEqual(a, c)) //true
		c = 2
		fmt.Println(reflect.DeepEqual(a, c)) //false
	
	}
