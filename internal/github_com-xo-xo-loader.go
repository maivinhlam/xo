// Code generated by 'yaegi extract github.com/KousukeUchiyama/xo/loader'. DO NOT EDIT.

package internal

import (
	"reflect"

	"github.com/KousukeUchiyama/xo/loader"
)

func init() {
	Symbols["github.com/KousukeUchiyama/xo/loader/loader"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"EnumValues":           reflect.ValueOf(loader.EnumValues),
		"Enums":                reflect.ValueOf(loader.Enums),
		"Flags":                reflect.ValueOf(loader.Flags),
		"IndexColumns":         reflect.ValueOf(loader.IndexColumns),
		"MysqlEnumValues":      reflect.ValueOf(loader.MysqlEnumValues),
		"MysqlGoType":          reflect.ValueOf(loader.MysqlGoType),
		"NthParam":             reflect.ValueOf(loader.NthParam),
		"OracleGoType":         reflect.ValueOf(loader.OracleGoType),
		"PQPostgresGoType":     reflect.ValueOf(loader.PQPostgresGoType),
		"PostgresFlags":        reflect.ValueOf(loader.PostgresFlags),
		"PostgresGoType":       reflect.ValueOf(loader.PostgresGoType),
		"PostgresIndexColumns": reflect.ValueOf(loader.PostgresIndexColumns),
		"PostgresTableColumns": reflect.ValueOf(loader.PostgresTableColumns),
		"PostgresViewStrip":    reflect.ValueOf(loader.PostgresViewStrip),
		"ProcParams":           reflect.ValueOf(loader.ProcParams),
		"Procs":                reflect.ValueOf(loader.Procs),
		"Register":             reflect.ValueOf(loader.Register),
		"Schema":               reflect.ValueOf(loader.Schema),
		"Sqlite3GoType":        reflect.ValueOf(loader.Sqlite3GoType),
		"SqlserverGoType":      reflect.ValueOf(loader.SqlserverGoType),
		"SqlserverViewStrip":   reflect.ValueOf(loader.SqlserverViewStrip),
		"StdlibPostgresGoType": reflect.ValueOf(loader.StdlibPostgresGoType),
		"TableColumns":         reflect.ValueOf(loader.TableColumns),
		"TableForeignKeys":     reflect.ValueOf(loader.TableForeignKeys),
		"TableIndexes":         reflect.ValueOf(loader.TableIndexes),
		"TableSequences":       reflect.ValueOf(loader.TableSequences),
		"Tables":               reflect.ValueOf(loader.Tables),
		"ViewCreate":           reflect.ValueOf(loader.ViewCreate),
		"ViewDrop":             reflect.ValueOf(loader.ViewDrop),
		"ViewSchema":           reflect.ValueOf(loader.ViewSchema),
		"ViewStrip":            reflect.ValueOf(loader.ViewStrip),
		"ViewTruncate":         reflect.ValueOf(loader.ViewTruncate),

		// type definitions
		"Loader": reflect.ValueOf((*loader.Loader)(nil)),
	}
}
