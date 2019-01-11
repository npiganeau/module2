// Copyright 2019 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package module2

import (
	"fmt"
	"github.com/npiganeau/module1/subpackg"
)

func MyFunc(param subpackg.SPType) {
	fmt.Println("Do something", param.SomeField)
}
