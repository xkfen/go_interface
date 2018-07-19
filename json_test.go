package main

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"fmt"
	"regexp"
)

type Movie struct {
	Title string 
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

func TestJsonMarshal(t *testing.T){
	movies := []Movie{
		{Title:"test",Year:2018, Color:false, Actors:[]string{"as", "erf"}},
		{Title:"test1", Year:2018, Color:true, Actors:[]string{"as", "erf"}},
	}
	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "  ")
	assert.NoError(t, err)
	fmt.Printf("%s\n", data)
}


// 正则匹配是否是字母
func checkIsLetter(str string) bool {
	match, _ := regexp.MatchString("^[A-Za-z]+$", str)
	return match
}

func TestCheckIsLetter(t *testing.T){
	flag := checkIsLetter("A写")
	assert.Equal(t, false, flag)
}

// 字符串原样输出(并且包括转义字符)
func TestString(t *testing.T){
	a := "testssd\n"
	fmt.Printf("%s", a)
}
