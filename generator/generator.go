package generator

import (
	"fmt"

	"github.com/releaseband/map-gen/generator/parser"
)

func Run(path string) error {
	fileDecl, err := parser.ParseFile(path)
	if err != nil {
		return err
	}

	for _, v := range fileDecl.Vars {
		fmt.Println(v.Name)
		fmt.Println(v.MapData)
	}
	return nil
}
