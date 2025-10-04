// Package globals provides global variables and constants for the application
package globals

import (
	ci "github.com/kubex-ecosystem/gdbase/internal/interfaces"
	svc "github.com/kubex-ecosystem/gdbase/internal/types"
)

func NewPropertyType[T any](name string, v *T, withMetrics bool, cb func(any) (bool, error)) ci.IProperty[T] {
	ci := svc.NewProperty(name, v, withMetrics, cb)

	//lint:ignore SA4023 explicação do motivo
	if ci == nil { //lint:ignore SA4023 explicação do motivo
		return nil
	}
	return ci.(*svc.Property[T])
}
