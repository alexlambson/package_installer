package resolver

import "fmt"

type CircularDependencyError struct {
	s string
}

func (e CircularDependencyError) Error() string {
	return e.s
}

func nameInSlice(name string, resolved []string) bool {
	for _, resolvedName := range resolved {
		if name == resolvedName {
			return true
		}
	}
	return false
}

func Resolve(node *Node, resolved *[]string, seen *[]string) (e error) {
	e = nil
	*seen = append(*seen, node.Name)
	for _, edge := range node.Edges{
		if !nameInSlice(edge.Name, *resolved) {
			if nameInSlice(edge.Name, *seen){
				errorString := fmt.Sprintf("Circular Dependency found: %s -> %s", node.Name, edge.Name)
				e = CircularDependencyError{s:errorString}
				return
			}
			Resolve(edge, resolved, seen)
		}
	}
	*resolved = append(*resolved, node.Name)
	return
}