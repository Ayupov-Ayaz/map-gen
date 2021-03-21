package tests

import (
	"testing"

	"github.com/releaseband/map-gen/generator"
)

func TestGenerate(t *testing.T) {
	const path = "./example_src.go"

	err := generator.Run(path)
	if err != nil {
		t.Fatal(err)
	}
}
