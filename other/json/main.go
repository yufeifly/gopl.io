package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(` [
        { "Name" : "Platypus" , "Order" : "Monotremata" } ,
        { "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" }
    ] `)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, val := range animals {
		fmt.Printf("%+v\n", val.Name)
	}

}
