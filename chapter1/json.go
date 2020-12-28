package main

import (
	"fmt";
	"encoding/json";	
)

type Foo struct {
	Bar string
	Baz string
}

func main() {

	f := Foo{"Saurabh Shinde", "Jay Shree Ram"}
	b, _ := json.Marshal(f);
	fmt.Println(string(b)); 

	json.Unmarshal(b, &f);
}