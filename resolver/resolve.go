package resolver

import "fmt"

func Resolve(node *Node){
	fmt.Println(node.Name)
	for _, edge := range node.Edges{
		Resolve(edge)
	}
}