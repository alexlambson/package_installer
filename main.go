package main

import (
	"fmt"
	"package_installer/logic"
)


func main() {
	/*
	Create the dependencies, this will be replaced with a parser
	 */
	a := logic.Node{Name:"a"}
	b := logic.Node{Name:"b"}
	c := logic.Node{Name:"c"}
	d := logic.Node{Name:"d"}
	e := logic.Node{Name:"e"}
	a.AddEdge(b)
	a.AddEdge(d)
	b.AddEdge(c)
	b.AddEdge(e)
	c.AddEdge(d)
	c.AddEdge(e)

	fmt.Println(a)
}
