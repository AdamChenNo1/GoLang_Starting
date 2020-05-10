package main

import(
	"fmt"
)

func EditingExample(){
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n",s2) 
}

func main(){
	EditingExample()
}