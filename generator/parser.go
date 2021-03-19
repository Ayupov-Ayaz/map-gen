package generator

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
)

const mapGenCommentTag = "//map_gen:"

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

func getName(commentText string) (string, bool) {
	if strings.HasPrefix(commentText, mapGenCommentTag) {
		elements := strings.Split(strings.Replace(commentText, mapGenCommentTag, "", 1), ";")

		for _, e := range elements {
			if strings.HasPrefix(e, "name=") {
				return strings.Replace(e, "name=", "", 1), true
			}
		}
	}

	return "", false
}

func isSingleVariable(decl *ast.GenDecl) bool {
	return decl.Doc != nil && len(decl.Doc.List) > 0
}

func parseSingleVariantData(decl *ast.GenDecl) (*Variant, error) {
	var (
		name string
		ok   bool
	)

	if decl.Doc != nil {
		for _, c := range decl.Doc.List {
			name, ok = getName(c.Text)
			if ok {
				break
			}
		}
	}

	if !ok {
		return nil, nil
	}

	mapData, err := ParseMap(decl.Specs)
	if err != nil {
		return nil, fmt.Errorf("ParseMap: %w", err)
	}

	return NewVariant(name, mapData), nil
}

func parseVar(decl *ast.GenDecl) ([]Variant, error) {
	var variants []Variant

	if isSingleVariable(decl) {
		variant, err := parseSingleVariantData(decl)
		if err != nil {
			return nil, fmt.Errorf("parseSingleVariantData: %w", err)
		}

		if variant != nil {
			variants = append(variants, *variant)
		}
	}

	return variants, nil
}
