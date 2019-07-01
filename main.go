package main

import "fmt"
import "github.com/triztian/my.special.library/sub-package"

func main() {
	fmt.Println(subpkg.Sum(1, 2))
}
