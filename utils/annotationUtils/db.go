package annotationUtils

import "github.com/dave/dst"

const (
	tableNameAnno = "@tableName"
	databaseAnno  = "@database"
	joinAnno      = "@join"
)

// region tableName

func ExistTableName(annotations dst.Decorations) bool {
	return Exist(annotations, tableNameAnno)
}

func GetTableName(annotations dst.Decorations) (val string, exist bool) {
	return FindStr(annotations, tableNameAnno)
}

// endregion

// region database

func ExistDatabaseName(annotations dst.Decorations) bool {
	return Exist(annotations, databaseAnno)
}

func GetDatabaseName(annotations dst.Decorations) (val string, exist bool) {
	return FindStr(annotations, databaseAnno)
}

// endregion

// region join

func ExistJoin(annotations dst.Decorations) bool {
	return Exist(annotations, joinAnno)
}

// endregion
