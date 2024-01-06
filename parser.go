package optiongen

import (
	"fmt"
	"go/types"
	"log/slog"

	"golang.org/x/tools/go/packages"
)

func LoadDefinition(packagePath string, typeName string, conf *packages.Config) (g Gen, err error) {
	pkgs, err := packages.Load(conf, packagePath)
	if err != nil {
		return g, err
	}
	if len(pkgs) != 1 {
		return g, fmt.Errorf("package load: pkg not found")
	}
	pkg := pkgs[0]
	for ident, obj := range pkg.TypesInfo.Defs {
		if obj == nil || ident == nil {
			continue
		}
		if obj.Name() != typeName {
			continue
		}
		if named, ok := obj.Type().(*types.Named); ok {
			st, ok := named.Underlying().(*types.Struct)
			if !ok {
				return g, fmt.Errorf("not a struct type: %v", named)
			}
			for i := 0; i < st.NumFields(); i++ {
				v := st.Field(i)
				field := Field{FieldName: v.Name(), FieldType: v.Origin().Type().String()}
				slog.Debug("Field", "detail", field, "pos", v.Pos())
				g.Fields = append(g.Fields, field)
			}
		}
	}
	g.TypeName = typeName
	return g, nil
}

// TODO: currently, the implementation doesn't understand package alias
// implement this function in future need
func TodoImports(pkg *packages.Package) {
	for k, p := range pkg.Imports {
		slog.Info("import", "k", k, "p", p)
	}
	for _, f := range pkg.Syntax {

		if f != nil && f.Name != nil {
			slog.Info("file", "name", f.Name.Name, "pos", f.Decls[0])
		}
		for _, is := range f.Imports {
			if is == nil {
				continue
			}
			name := ""
			if is.Name != nil {
				name = is.Name.Name
			}
			path := ""
			if is.Path != nil {
				path = is.Path.Value
			}
			slog.Info("syntax import", "name", name, "path", path)
		}
	}
}
