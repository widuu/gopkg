package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	volume := filepath.VolumeName("D:/test/path/filepath")
	fmt.Println(volume) //	D:
}
