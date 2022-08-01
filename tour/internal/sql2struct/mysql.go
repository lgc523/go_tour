package sql2struct

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Port     string
	Charset  string
}

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
	ColumnDefault interface{}
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Conncet() error {
	var err error
	s := "%s:%s@tcp(%s:%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Port,
		m.DBInfo.Charset,
	)
	log.Printf("Connect MySQL Protocol is:\t{%s}", dsn)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		log.Fatalf("[connect db fail,cause:%s]", err.Error())
		return err
	}
	return nil
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	querySQL :=
		"SELECT COLUMN_NAME,DATA_TYPE,COLUMN_KEY,IS_NULLABLE,COLUMN_TYPE,COLUMN_COMMENT,COLUMN_DEFAULT " +
			" FROM COLUMNS " +
			" WHERE TABLE_SCHEMA = ?  AND TABLE_NAME = ? "
	r, err := m.DBEngine.Query(querySQL, dbName, tableName)
	if err != nil {
		log.Printf("[execute sql %s,failed cause: %s\n,param:{%s}-{%s}]", querySQL, err, dbName, tableName)
		return nil, err
	}
	if r == nil {
		return nil, fmt.Errorf("[Table :%s:%s has no data]", dbName, tableName)
	}
	defer r.Close()

	var columns []*TableColumn

	for r.Next() {
		var column TableColumn
		err := r.Scan(&column.ColumnName,
			&column.DataType,
			&column.ColumnKey,
			&column.IsNullable,
			&column.ColumnType,
			&column.ColumnComment,
			&column.ColumnDefault,
		)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}

var DBTypeToStructType = map[string]string{
	"int":       "int",
	"tinyint":   "int8",
	"bool":      "bool",
	"varchar":   "string",
	"timestamp": "time.Time",
	"datetime":  "time.Time",
	"decimal":   "decimal.Decimal",
}
