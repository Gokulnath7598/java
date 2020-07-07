package main

import ("fmt"
)

func main(){

	var Name string = "Ramya"

	var Age int16 = 21

	Age1 := int64(Age)

	weight := 48.2

	weight1 := int(weight)

	com := 1+10i

	var torf = (1==1)

	fmt.Println(torf)

	fmt.Printf("The Name of the user is %v is of type %T\n",Name,Name)
	fmt.Printf("The Age Of %v is %v is of Type %T\n",Name,Age1,Age1)
	fmt.Printf("Weight of %v is %v which is of type %T\n",Name,weight1,weight1)
	fmt.Printf("the type of %v is %T",com,com)

}
