package main

import (
	"github.com/rafa-mori/gdbase/internal/module"
	gl "github.com/rafa-mori/gdbase/internal/module/logger"
)

// main initializes the logger and creates a new GDBase instance.
func main() {
	if err := module.RegX().Command().Execute(); err != nil {
		gl.Log("fatal", err.Error())
	}
}
