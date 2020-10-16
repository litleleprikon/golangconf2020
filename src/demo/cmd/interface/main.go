package main

// IMPORTS START OMIT
import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"strings"
	"text/template"
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

	structs := findStructs(node)
	printProgram(node, fset, structs)
}

func printProgram(node ast.Node, fset *token.FileSet, structs map[string][]string) {
	//PRINTER START OMIT
	var buf strings.Builder
	printer.Fprint(&buf, fset, node)
	//PRINTER END OMIT

	// TMPL START OMIT
	tmplString := `
	{{ range $struct, $fields := . }}
		func (s {{ $struct }}) Hash() int {
			return {{ StringsJoin $fields ", " }}
		}

 	{{ end }}
	`
	tmpl := template.Must(
		template.New("print").
			Funcs(template.FuncMap{"StringsJoin": strings.Join}).
			Parse(tmplString))
	buf.WriteString("\n")
	err := tmpl.Execute(&buf, structs)
	// TMPL END OMIT

	if err != nil {
		log.Fatal(err)
	}
	// PRINT START OMIT
	formated, err := format.Source([]byte(buf.String()))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(formated))
	// PRINT END OMIT
}

func findStructs(node *ast.File) map[string][]string {
	result := make(map[string][]string)
	for _, f := range node.Decls {
		dec, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		if dec.Tok != token.TYPE {
			continue
		}
		for _, spec := range dec.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			fieldNames := []string{}
			for _, field := range structType.Fields.List {
				ident, ok := field.Type.(*ast.Ident)
				if !ok {
					continue
				}
				if ident.Name != "int" {
					continue
				}
				for _, name := range field.Names {
					fieldNames = append(fieldNames, name.Name)
				}
			}
			result[typeSpec.Name.Name] = fieldNames
		}
	}
	return result
}
