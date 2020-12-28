package main

import (
	"fmt"
)

func strlen(s string, c chan int) {
	c <-len(s); //placing the  data in the bucket [channel]
}


func main() {
	c := make(chan int)

	go strlen("Salutations", c); 

	go strlen("World", c);

	x, y := <-c, <-c; // extracting data out of bucket

	fmt.Println(x, y, x + y);
}