package main

import "fmt"
import "runtime"

func main(){
	runtime.GOMAXPROCS(2)
	
	go tampilkan_pesan()
	tampilkan_pesan2()
	
	var input string
	fmt.Scanln(&input)
}

func tampilkan_pesan(){
	for i:=0;i<100;i++{
		fmt.Println("Hai Jack")
	}
}

func tampilkan_pesan2(){
	for i:=0;i<100;i++{
		fmt.Println("Hai Night")
	}
}