package hclq

import (
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/token"
)

// Set traverses the document and calls either listAction or valueAction depending
// on whether or not the value is a list or a literal value. These functions will
// be invoked for ALL matching nodes in the query.
func (doc *HclDocument) Set(queryString string, listAction func(*ast.ListType) error, valueAction func(*token.Token) error) error {
	resultPairs, err := doc.Query(queryString)
	if err != nil {
		return err
	}

	for _, pair := range resultPairs {
		list, ok := pair.Node.(*ast.ListType)
		if ok {
			err := listAction(list)
			if err != nil {
				return err
			}
			continue
		}
		literal, ok := pair.Node.(*ast.LiteralType)
		if ok {
			err := valueAction(&literal.Token)
			if err != nil {
				return err
			}
			continue
		}
	}
	return nil
}
