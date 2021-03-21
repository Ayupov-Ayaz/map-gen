package parser

type Variant struct {
	Name    string
	MapData map[string][]string
}

func NewVariant(name string, mapData map[string][]string) *Variant {
	return &Variant{Name: name, MapData: mapData}
}

type FileDeclaration struct {
	Vars []Variant
}

func (d *FileDeclaration) AddVariants(variants []Variant) {
	if len(d.Vars) == 0 {
		d.Vars = variants
	} else {
		d.Vars = append(d.Vars, variants...)
	}
}
