package controller

import (
	"github.com/daplho/cnat/cnat-operator/pkg/controller/at"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, at.Add)
}
