package main

import (

"fmt"
"bufio"
"os"
)



func main()  {


	var userinput = make(map[string]string)

	r := getuserinput(userinput,"v1")
	r = getuserinput(userinput,"v2")


	if r != nil{

		fmt.Print(r.Error())
	}


	for k, i:= range userinput{


		fmt.Print("[",k,",",i,"] ")

	}
}



func getuserinput(elements map[string]string,key string) (r error) {


	fmt.Println("enter name ")

	reader := bufio.NewReader(os.Stdin)

	input, r := reader.ReadString('\n')

	if r != nil{

		fmt.Println(r.Error())


		return r
	}else{

		elements[key] = input

		return nil
	}


}