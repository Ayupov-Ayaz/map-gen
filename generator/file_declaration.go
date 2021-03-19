package generator

import "go/ast"

type FileDeclaration struct {
	Imports []string
	Vars    []ast.Decl
}

func (d *FileDeclaration) AddImports(imports []string) {
	if len(d.Imports) == 0 {
		d.Imports = imports
	} else {
		d.Imports = append(d.Imports, imports...)
	}
}
