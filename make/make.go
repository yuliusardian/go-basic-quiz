package main

import "fmt"

func main() {
	// declare type and assign into variable mm
	mm := make(map[string]string)
	// assign value Yulius Ardian Febrianto for array name in mm variable
	mm["name"] = "Yulius Ardian Febrianto"
	mm["pob"] = "Yogyakarta"
	mm["dob"] = "04 February 1996"
	fmt.Println(mm)
	fmt.Println(mm["name"])
	fmt.Println(mm["pob"])
	fmt.Println(mm["dob"])

	mss := make([]int, 3)
	mss[0] = 100
	mss[1] = 721
	fmt.Println(mss)
	fmt.Println(mss[0])
	fmt.Println(mss[1])
}
