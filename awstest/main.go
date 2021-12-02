package main

import (
	A "awstest/aws"
	"fmt"
)

func main() {
	A.ShowBuckets()
	A.ShowKeys()
	//A.ShowObjects()
	fmt.Println("test")

}
