// +build !some.constraint.with-dash

package cmd

import (
	"fmt"
	"github.com/triztian/my.special.library/really/deep/nested/sub-package"
)

func run() error {
	fmt.Println(subpkg.Sum(21, 21))
	return nil
}
