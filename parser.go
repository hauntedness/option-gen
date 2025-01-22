package optiongen

import (
	"cmp"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"log/slog"
	"slices"

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
				fieldType := v.Origin().Type()
				fieldTypeString := types.TypeString(fieldType, func(p *types.Package) string {
					if p == nil || p.Path() == pkg.PkgPath {
						return ""
					}
					return p.Name()
				})
				field := Field{FieldName: v.Name(), FieldType: fieldTypeString}
				slog.Debug("Field", "detail", field, "pos", v.Pos())
				g.Fields = append(g.Fields, field)
			}
		}
		break
	}
	g.TypeName = typeName
	g.PackageName = pkg.Name
	return g, nil
}

func LoadDefinitions(packagePath string, typeList []string, conf *packages.Config) (gs []Gen, err error) {
	pkgs, err := packages.Load(conf, packagePath)
	if err != nil {
		return nil, err
	}
	if len(pkgs) != 1 {
		return nil, fmt.Errorf("package load: pkg not found")
	}
	pkg := pkgs[0]
	for ident, obj := range pkg.TypesInfo.Defs {
		if obj == nil || ident == nil {
			continue
		}
		if !slices.Contains(typeList, obj.Name()) {
			continue
		}
		g := Gen{}
		g.TypeName = obj.Name()
		g.PackageName = pkg.Name
		if named, ok := obj.Type().(*types.Named); ok {
			st, ok := named.Underlying().(*types.Struct)
			if !ok {
				return nil, fmt.Errorf("not a struct type: %v", named)
			}
			for i := 0; i < st.NumFields(); i++ {
				v := st.Field(i)
				fieldType := v.Origin().Type()
				fieldTypeString := types.TypeString(fieldType, func(p *types.Package) string {
					if p == nil || p.Path() == pkg.PkgPath {
						return ""
					}
					return p.Name()
				})
				field := Field{FieldName: v.Name(), FieldType: fieldTypeString}
				slog.Debug("Field", "detail", field, "pos", v.Pos())
				g.Fields = append(g.Fields, field)
			}
		}
		gs = append(gs, g)
	}

	if len(typeList) != len(gs) {
		var names []string
		for i := range gs {
			names = append(names, gs[i].TypeName)
		}
		return gs, fmt.Errorf("some types were missing in generation. given: %v, found: %v", typeList, names)
	}

	return gs, nil
}

// TODO: currently, the implementation doesn't understand package alias
// implement this function in future need
func Files(pkg *packages.Package) files {
	files := make([]file, 0, 20)
	for _, f := range pkg.Syntax {
		if f == nil {
			continue
		}
		var file file
		file.filename = pkg.Fset.File(f.Pos()).Name()
		file.pos = f.Pos()
		file.imports = f.Imports
		files = append(files, file)
	}
	slices.SortFunc(files, func(f1, f2 file) int {
		return cmp.Compare(f1.pos, f2.pos)
	})
	return files
}

type file struct {
	filename string
	imports  []*ast.ImportSpec
	pos      token.Pos
}

func (f file) Print() {
	slog.Info(f.filename, "pos", f.pos)
	for _, is := range f.imports {
		if is == nil {
			continue
		}
		alias := ""
		if is.Name != nil {
			alias = is.Name.Name
		}
		path := ""
		if is.Path != nil {
			path = is.Path.Value
		}
		slog.Info("syntax import", "alias", alias, "path", path)
	}
}

func (f file) ImportName(pkg string) (string, bool) {
	for _, is := range f.imports {
		if is == nil {
			continue
		}
		if is.Path != nil && is.Path.Value == pkg {
			if is.Name != nil {
				return is.Name.Name, true
			}
		}
	}
	return "", false
}

type files []file

func (fs files) Search(pos token.Pos) (file, bool) {
	p, _ := slices.BinarySearchFunc(fs, pos, func(f1 file, t token.Pos) int {
		return cmp.Compare(f1.pos, t)
	})
	if p < int(fs[p].pos) {
		return fs[p+1], true
	} else if p > int(fs[p].pos) {
		return fs[p], true
	}
	return file{}, false
}
