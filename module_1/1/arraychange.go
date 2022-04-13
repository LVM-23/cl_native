package main

import "fmt"

func main(){
	  strArray  := [5]string{"I","am" ,"stupid","and","weak"}

	for  index,_ := range strArray{
		if index ==2 {
			strArray[index] = "smart"
		}

		 if index ==4 {
			 strArray[index] = "strong"
		 }
	}
	fmt.Println(strArray)
}


