// go:generate go run github.com/rafa-mori/gdbase/internalmodels/user GenUser
package main

import (
	"fmt"

	gl "github.com/rafa-mori/gdbase/logger"
)

func main() {
	if err := RegX().Execute(); err != nil {
		gl.Log("error", fmt.Sprintf("Error: %v", err))
	}
}
