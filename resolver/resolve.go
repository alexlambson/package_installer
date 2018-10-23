package resolver

import "fmt"

func Resolve(node *Node, resolved *[]string){
	fmt.Println(node.Name)
	for _, edge := range node.Edges{
		Resolve(edge, resolved)
	}
	*resolved = append(*resolved, node.Name)
}