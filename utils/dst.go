package utils

import (
	"fmt"
	"github.com/dave/dst"
)

func GetTypeName(expr dst.Expr) string {
	switch expr.(type) {
	case *dst.Ident:
		return expr.(*dst.Ident).Name
	case *dst.SelectorExpr:
		return expr.(*dst.SelectorExpr).X.(*dst.Ident).Name + "." + expr.(*dst.SelectorExpr).Sel.Name
	case *dst.ArrayType:
		return "[]" + GetTypeName(expr.(*dst.ArrayType).Elt)
	case *dst.InterfaceType:
		return "interface{}"
	case *dst.StarExpr:
		return "*" + GetTypeName(expr.(*dst.StarExpr).X)
	case *dst.IndexExpr:
		return fmt.Sprintf("%v[%v]", GetTypeName(expr.(*dst.IndexExpr).X), GetTypeName(expr.(*dst.IndexExpr).Index))
	case *dst.MapType:
		return fmt.Sprintf("map[%v]%v", GetTypeName(expr.(*dst.MapType).Key), GetTypeName(expr.(*dst.MapType).Value))
	default:
		return ""
	}
}
