package main


import "fmt"

func main(){
	fmt.Print("enter the number")
	var n int
	fmt.Scanln(&n)

	if(n%2==0){
		fmt.Println(n,"is even number")
	}else{
		fmt.Println(n, "is odd number")
	}

}