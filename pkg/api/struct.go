package api

import (
	"fmt"

	diinotools "github.com/ksrkelvin/diinoTools"
)

// DiinoAPI - DiinoAPI struct
type DiinoAPI struct {
	Diino *diinotools.Diino
}

// New - New DiinoAPI
func New() (api *DiinoAPI, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	diino, err := diinotools.New()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	newDiinoAPI := &DiinoAPI{
		Diino: diino,
	}

	return newDiinoAPI, err
}
