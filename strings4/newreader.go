package main

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
