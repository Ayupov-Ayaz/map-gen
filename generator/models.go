package generator

type Map struct {
	values map[string][]string
}

func NewMap(values map[string][]string) *Map {
	return &Map{
		values: values,
	}
}

type Variant struct {
	Name string
	Maps *Map
}

func NewVariant(name string, mapData *Map) *Variant {
	return &Variant{Name: name, Maps: mapData}
}

type FileDeclaration struct {
	Imports []string
	Vars    []Variant
}

func (d *FileDeclaration) AddImports(imports []string) {
	if len(d.Imports) == 0 {
		d.Imports = imports
	} else {
		d.Imports = append(d.Imports, imports...)
	}
}

func (d *FileDeclaration) AddVariants(variants []Variant) {
	if len(d.Vars) == 0 {
		d.Vars = variants
	} else {
		d.Vars = append(d.Vars, variants...)
	}
}
