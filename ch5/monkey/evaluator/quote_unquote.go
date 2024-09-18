// evaluator/quote_unquote.go

package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
