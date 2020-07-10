package controller

import (
	"kube-hive/pkg/controller/drone"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, drone.Add)
}
