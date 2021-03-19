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

func isMultiVariant(decl *ast.GenDecl) bool {
	return !isSingleVariable(decl) && len(decl.Specs) > 0
}

func getSingleVarComment(decl *ast.GenDecl) (Variant, bool) {
	if decl.Doc != nil {
		for _, c := range decl.Doc.List {
			name, ok := getName(c.Text)
			if ok {
				return NewVariant(name), true
			}
		}
	}

	return Variant{}, false
}

func getMultiVarComment(decl *ast.GenDecl) []Variant {
	var variants []Variant

	for _, s := range decl.Specs {
		vs, ok := s.(*ast.ValueSpec)

		list := vs.Doc.List
		count := len(list)

		if !ok || vs.Doc == nil || count == 0 {
			continue
		}

		if variants == nil {
			variants = make([]Variant, 0, count)
		}

		for _, c := range list {
			name, ok := getName(c.Text)
			if ok {
				variants = append(variants, NewVariant(name))
			}
		}
	}

	return variants
}

func parseVar(decl *ast.GenDecl) ([]Variant, error) {
	var variants []Variant

	if isSingleVariable(decl) {
		variant, ok := getSingleVarComment(decl)
		if ok {
			variants = append(variants, variant)
		}

	} else if isMultiVariant(decl) {
		mVariants := getMultiVarComment(decl)
		if len(mVariants) > 0 {
			variants = append(variants, mVariants...)
		}
	}

	for _, c := range variants {
		fmt.Println(c.Name)
	}

	return nil, nil
}
