package generator

import (
	"errors"
	"fmt"
	"go/ast"
)

func CastSliceBasicList(expr []ast.Expr) ([]*ast.BasicLit, error) {
	bls := make([]*ast.BasicLit, len(expr))
	for i, e := range expr {
		bl, ok := e.(*ast.BasicLit)
		if !ok {
			return nil, errors.New("cast to ast.BasicLit failed")
		}

		bls[i] = bl
	}

	return bls, nil
}

func ParseMapValues(cl *ast.CompositeLit) (map[string][]string, error) {
	results := make(map[string][]string, len(cl.Elts))

	for _, v := range cl.Elts {
		kvExpr, ok := v.(*ast.KeyValueExpr)
		if !ok {
			return nil, errors.New("cast to ast.KeyValueExpr failed")
		}

		key, ok := kvExpr.Key.(*ast.BasicLit)
		if !ok {
			return nil, errors.New("cast to ast.BasicList failed")
		}

		clValues, ok := kvExpr.Value.(*ast.CompositeLit)
		if !ok {
			return nil, errors.New("cast to ast.CompositeLit failed")
		}

		values, err := CastSliceBasicList(clValues.Elts)
		if err != nil {
			return nil, err
		}

		vData := make([]string, len(values))
		for i := 0; i < len(values); i++ {
			vData[i] = values[i].Value
		}

		results[key.Value] = vData
	}

	return results, nil
}

func parseMap(vs *ast.ValueSpec) (*Map, error) {
	for _, v := range vs.Values {
		lit, ok := v.(*ast.CompositeLit)
		if !ok {
			continue
		}

		mapValues, err := ParseMapValues(lit)
		if err != nil {
			return nil, fmt.Errorf("ParseMapValues: %w", err)
		}

		return NewMap(mapValues), nil
	}

	return nil, errors.New("map not found")
}

func ParseMap(specs []ast.Spec) (*Map, error) {

	for _, spec := range specs {
		vs, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}

		mapValues, err := parseMap(vs)
		if err != nil {
			return nil, err
		}

		return mapValues, nil
	}

	return nil, nil
}
