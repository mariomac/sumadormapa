package main

import (
	"fmt"
	 "github.com/mariomac/sumadormapa/sumador"
)

func main() {
	cs := sumador.Claves{}
	cs.Incrementa("a")
	cs.Incrementa("b")
	cs.Incrementa("c")
	cs.Incrementa("a")
	cs.Incrementa("b")
	cs.Incrementa("a")

	for k, v := range cs {
		fmt.Println(k, " -> ", v)
	}
}
