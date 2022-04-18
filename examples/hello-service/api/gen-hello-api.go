package main

import (
	"os"
	"path"

	log "github.com/ipfs/go-log"
	"github.com/ipld/edelweiss/compile"
	"github.com/ipld/edelweiss/defs"
)

var proto = defs.Defs{

	// hello service definition
	defs.Named{
		Name: "Hello",
		Type: defs.Service{
			Methods: defs.Methods{
				defs.Method{
					Name: "Hello",
					Type: defs.Fn{
						Arg:    defs.Ref{Name: "HelloRequest"},
						Return: defs.Ref{Name: "HelloResponse"},
					},
				},
			},
		},
	},

	// request type
	defs.Named{
		Name: "HelloRequest",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{
					Name: "Name",
					Type: defs.String{},
				},
				defs.Field{
					Name: "Address",
					Type: defs.Inductive{
						Cases: defs.Cases{
							defs.Case{Name: "US", GoName: "US", Type: defs.Ref{Name: "USAddress"}},
							defs.Case{Name: "SouthKorea", GoName: "SK", Type: defs.Ref{Name: "SKAddress"}},
						},
						Default: defs.DefaultCase{GoKeyName: "Other", GoValueName: "Address", Type: defs.List{Element: defs.String{}}},
					},
				},
			},
		},
	},

	defs.Named{
		Name: "USAddress",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{Name: "street", GoName: "Street", Type: defs.String{}},
				defs.Field{Name: "city", GoName: "City", Type: defs.String{}},
				defs.Field{Name: "state", GoName: "State", Type: defs.Ref{Name: "State"}},
				defs.Field{Name: "zip", GoName: "ZIP", Type: defs.Int{}},
			},
		},
	},

	defs.Named{
		Name: "USAddress",
		Type: defs.Union{
			Cases: defs.Cases{
				defs.Case{Name: "ca", GoName: "CA", Type: defs.SingletonString{"CA"}},
				defs.Case{Name: "ny", GoName: "NY", Type: defs.SingletonString{"NY"}},
				defs.Case{Name: "other", GoName: "Other", Type: defs.String{}},
			},
		},
	},

	defs.Named{
		Name: "SKAddress",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{Name: "street", GoName: "Street", Type: defs.String{}},
				defs.Field{Name: "city", GoName: "City", Type: defs.String{}},
				defs.Field{Name: "province", GoName: "Province", Type: defs.String{}},
				defs.Field{Name: "postal_code", GoName: "PostalCode", Type: defs.Int{}},
			},
		},
	},

	// response type
	defs.Named{
		Name: "HelloResponse",
		Type: defs.Union{
			Cases: defs.Cases{
				defs.Case{Name: "", GoName: "", Type: XXX},
			},
		},
	},
}

var logger = log.Logger("api generator")

func main() {
	wd, err := os.Getwd()
	if err != nil {
		logger.Errorf("working dir (%v)\n", err)
		os.Exit(-1)
	}
	dir := path.Join(wd, "proto")
	x := &compile.GoPkgCodegen{
		GoPkgDirPath: dir,
		GoPkgPath:    "github.com/ipld/edelweiss/examples/hello-service/api/proto",
		Defs:         proto,
	}
	goFile, err := x.Compile()
	if err != nil {
		logger.Errorf("compilation (%v)\n", err)
		os.Exit(-1)
	}
	if err = os.Mkdir(dir, 0755); err != nil {
		logger.Errorf("making pkg dir (%v)\n", err)
		os.Exit(-1)
	}
	if err = goFile.Build(); err != nil {
		logger.Errorf("build (%v)\n", err)
		os.Exit(-1)
	}
}
