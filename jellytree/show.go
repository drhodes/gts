package main
import "fmt"

func show(n *node, depth int) string {
	padding := ""
	for i:=0; i<depth*2; i++ {
		padding += " "
	}

	if n == nil {
		return padding + "nil"
	}

	xs := fmt.Sprint(padding, n.Group.Show()) + "\n"
	l := show(n.left, depth+1) + "\n"
	r := show(n.right, depth+1) + "\n"
	return xs+l+r
}

func (j Jtree) Show() string {
	return show(j.root, 0)
}
