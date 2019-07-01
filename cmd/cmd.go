package cmd

import "github.com/triztian/my.special.library/sub-package"
import "fmt"

func Run() error {
	fmt.Println(subpkg.Sum(21, 21))
	return nil
}
