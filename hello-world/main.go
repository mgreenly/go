package main

import (
	"fmt"

	"github.com/mgreenly/go/hello-dep/hello"
	"github.com/mgreenly/go/hello-dep/world"
)

func main() {
	fmt.Println(hello.SayHello())
	fmt.Println(world.SayWorld())
}
