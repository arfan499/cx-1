package assert

import (
	"fmt"
	"github.com/skycoin/cx/cx/ast"
	"github.com/skycoin/cx/cx/constants"
)

func ArgOfType(arg *ast.CXArgument, t int) {
	if arg.Type != t {
		panic(fmt.Sprintf("Argument %s, expected type %s, got %s",
			arg.Name, constants.TypeNames[t], constants.TypeNames[arg.Type]))
	}
}

func ArgAtomic(arg *ast.CXArgument) {
	if arg.IsStruct || arg.IsPointer || arg.IsArray || arg.IsSlice {
		panic(fmt.Sprintf("Argument %s of type %s is not atomic", arg.Name, constants.TypeNames[arg.Type]))
	}
}

func ArgNotAtomic(arg *ast.CXArgument) {
	if !arg.IsStruct && !arg.IsPointer && !arg.IsArray && !arg.IsSlice {
		panic(fmt.Sprintf("Argument %s of type %s is atomic", arg.Name, constants.TypeNames[arg.Type]))
	}
}

func ArgPointer(arg *ast.CXArgument) {
	if !arg.IsPointer {
		panic(fmt.Sprintf("Argument %s is not a pointer", arg.Name))
	}
}

func ArgNotPointer(arg *ast.CXArgument) {
	if arg.IsPointer {
		panic(fmt.Sprintf("Argument %s is a pointer", arg.Name))
	}
}

func ArgSlice(arg *ast.CXArgument) {
	if !arg.IsSlice {
		panic(fmt.Sprintf("Argument %s is not a slice", arg.Name))
	}
}

func ArgNotSlice(arg *ast.CXArgument) {
	if arg.IsSlice {
		panic(fmt.Sprintf("Argument %s is a slice", arg.Name))
	}
}

func ArgStruct(arg *ast.CXArgument) {
	if !arg.IsStruct {
		panic(fmt.Sprintf("Argument %s is not a struct", arg.Name))
	}
}

func ArgNotStruct(arg *ast.CXArgument) {
	if arg.IsStruct {
		panic(fmt.Sprintf("Argument %s is a struct", arg.Name))
	}
}

func ArgArray(arg *ast.CXArgument) {
	if !arg.IsArray {
		panic(fmt.Sprintf("Argument %s is not a array", arg.Name))
	}
}

func ArgNotArray(arg *ast.CXArgument) {
	if arg.IsArray {
		panic(fmt.Sprintf("Argument %s is a array", arg.Name))
	}
}
