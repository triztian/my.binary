package cmd

import (
	"fmt"
	"github.com/triztian/my.special.library/really/deep/nested/sub-package"
)

func Run() error {
	fmt.Println(subpkg.Sum(21, 21))
	return nil
}
