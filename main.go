package main

import (
	"go_interface/model"
	"go_interface/interface"
)

func main(){
	var fly _interface.Fly = new(model.Bird)
	fly.Fly()
}