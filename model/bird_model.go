package model

import "fmt"

type Bird struct {

}


func (bird *Bird) Fly(){
	fmt.Println("----fly")
}