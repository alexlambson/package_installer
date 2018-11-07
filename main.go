package main

import (
	"fmt"
	"os"
	"package_installer/resolver"
)

func main() {
	/*
	Create the dependencies, this will be replaced with a parser
	 */
	a := &resolver.Node{Name: "a"}
	b := &resolver.Node{Name: "b"}
	c := &resolver.Node{Name: "c"}
	d := &resolver.Node{Name: "d"}
	e := &resolver.Node{Name: "e"}
	a.AddEdge(b)
	a.AddEdge(d)
	b.AddEdge(c)
	b.AddEdge(e)
	c.AddEdge(d)
	c.AddEdge(e)
	d.AddEdge(b)
	dependencyOrder := &[]string{}
	seenNodes := &[]string{}
	error := resolver.Resolve(a, dependencyOrder, seenNodes)
	if error != nil {
		fmt.Println(error)
		os.Exit(2)
	}
	for _, node := range *dependencyOrder {
		fmt.Print(node)
	}
}
