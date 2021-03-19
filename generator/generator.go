package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func parseFile(path string) (*FileDeclaration, error) {
	set, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile: %w", err)
	}

	fileDecl := &FileDeclaration{}

	for _, d := range set.Decls {
		decl, ok := d.(*ast.GenDecl)
		if ok {
			switch decl.Tok {
			case token.IMPORT:
				imports, err := parseImports(decl)
				if err != nil {
					return nil, err
				}

				fileDecl.AddImports(imports)
				// todo: Var
			}
		}
	}

	return fileDecl, nil
}

func Run(path string) error {
	fileDeclaration, err := parseFile(path)
	if err != nil {
		return err
	}

	fmt.Println(fileDeclaration)

	return nil
}
