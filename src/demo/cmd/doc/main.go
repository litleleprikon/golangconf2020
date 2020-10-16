package main

// IMPORTS START OMIT
import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

// IMPORTS END OMIT

func main() {
	// PREPARE START OMIT
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "src/demo/example/main.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	// PREPARE END OMIT
	node = processAST(node)

	//PRINTER STARTS OMIT
	printer.Fprint(os.Stdout, fset, node)
	//PRINTER ENDS OMIT
}

func processAST(node *ast.File) *ast.File {

	//SEARCH START OMIT
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if fn.Name.IsExported() {
			if fn.Doc.Text() == "" {
				//SEARCH END OMIT
				//AST EDIT STARTS OMIT
				fn.Doc = &ast.CommentGroup{
					List: []*ast.Comment{
						&ast.Comment{
							Text:  fmt.Sprintf("// %s ... Place documentation here", fn.Name.Name),
							Slash: fn.Pos() - 1},
					},
				}
				//AST EDIT END OMIT
			}
		}
	}

	return node
}
