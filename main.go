package main

import (
	"fmt"

	"github.com/rafialariq/go-cicd/pkg"
)

func main() {
	result := pkg.Sum(8)
	fmt.Println(result)
}
