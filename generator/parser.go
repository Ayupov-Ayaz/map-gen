package generator

import (
	"errors"
	"go/ast"
	"strings"
)

const mapGenCommentTag = "//map_gen"

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

func parseMapGenComments(comments []*ast.Comment) bool {
	for _, c := range comments {
		if strings.HasPrefix(c.Text, mapGenCommentTag) {

		}
	}

	return false
}

func getCommentFromSingleVar(decl *ast.GenDecl) []*ast.Comment {
	if decl.Doc != nil && len(decl.Doc.List) > 0 {
		return decl.Doc.List
	}

	return nil
}

func getCommentFromMultiVar(decl *ast.GenDecl) []*ast.Comment {
	if len(decl.Specs) == 0 {
		return nil
	}

	var comments []*ast.Comment

	for _, s := range decl.Specs {
		vs, ok := s.(*ast.ValueSpec)
		if !ok || vs.Doc == nil || vs.Doc.List == nil {
			continue
		}

		comments = append(comments, vs.Doc.List...)
	}

	return comments
}

func getComments(decl *ast.GenDecl) []*ast.Comment {
	comments := getCommentFromSingleVar(decl)
	if comments != nil {
		return comments
	}

	return getCommentFromMultiVar(decl)
}

func parseVars(decl *ast.GenDecl) ([]Variant, error) {
	comments := getComments(decl)
	parseMapGenComments(comments)

	return nil, nil
}
