package generator

import (
	"errors"
	"go/ast"
)

func parseImports(d *ast.GenDecl) ([]string, error) {
	imports := make([]string, len(d.Specs))
	for i, spec := range d.Specs {
		imp, ok := spec.(*ast.ImportSpec)
		if !ok {
			return nil, errors.New("cast import failed")
		}

		imports[i] = imp.Path.Value
	}

	return imports, nil
}
