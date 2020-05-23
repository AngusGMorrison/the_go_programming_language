// Using the token-based decoder API, write a program that will read an arbitrary XML document
// and construct a tree of generic nodes that represent it. Nods are of two kinds: CharData nodes
// represent text strings, and Element nodes represent named elements and theri attributes. Each
// element node has a slice of child nodes.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	decoder := xml.NewDecoder(os.Stdin)
	root, err := parse(decoder)
	if err != io.EOF && err != nil {
		log.Fatalf("xmlparser: %v", err)
	}
	printXML(root)
}

func parse(d *xml.Decoder) (Node, error) {
	var stack []*Element

	for {
		token, err := d.Token()
		if err != nil {
			return nil, err
		}

		switch token := token.(type) {
		case xml.StartElement:
			el := Element{token.Name, token.Attr, []Node{}}
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, &el)
			}
			stack = append(stack, &el)
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("unexpected end element: %+v", token)
			} else if len(stack) == 1 { // end of tree
				return stack[0], nil
			}
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, CharData(token))
			}
		}
	}
}

var indent int = -2

func printXML(n Node) {
	if n == nil {
		return
	}
	indent += 2
	switch n := n.(type) {
	case *Element:
		fmt.Printf("%*c<%s%s>\n", indent, ' ', n.Type.Local, formatAttrs(n.Attr))
		for _, child := range n.Children {
			printXML(child)
		}
		fmt.Printf("%*c</%s>\n", indent, ' ', n.Type.Local)
	case CharData:
		fmt.Printf("%*c%s\n", indent, ' ', string(n))
	}
	indent -= 2
}

func formatAttrs(attrs []xml.Attr) string {
	fmtdAttrs := make([]string, len(attrs))
	for i, attr := range attrs {
		fmtdAttrs[i] = fmt.Sprintf(" %s=%q", attr.Name, attr.Value)
	}
	return strings.Join(fmtdAttrs, "")
}
