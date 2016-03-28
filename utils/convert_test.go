package utils

 import (
 	"fmt"
 	"testing"
 )

 type person struct {
 	Name string
 	Age  int
 }

 func TestMapToStruct(t *testing.T) {
 	m := make(map[string]interface{})
 	m["Name"] = "luis"
 	m["Age"] = 27
 	fmt.Println(m)
 	p := &person{}
 	err := MapToStruct(m, p)
 	if err != nil {
 		fmt.Println(err)
 	}
 	fmt.Println(p)
 }
