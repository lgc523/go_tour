package cmd

import (
	"log"

	"github.com/go-tour/tour/internal/sql2struct"
	"github.com/spf13/cobra"
)

var userName string
var passWd string
var host string
var charset string
var dbType string
var dbName string
var tableName string
var port string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql mapping",
	Long:  "sql 转换处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换",
	Long:  "sql 转换 struct",
	Run:   sql2StructHandler,
}

func sql2StructHandler(cmd *cobra.Command, args []string) {
	dbInfo := &sql2struct.DBInfo{
		DBType:   dbType,
		Host:     host,
		Port:     port,
		UserName: userName,
		Password: passWd,
		Charset:  charset,
	}
	dbModel := sql2struct.NewDBModel(dbInfo)
	err := dbModel.Conncet()
	if err != nil {
		log.Fatalf("[DbModel.connect err:%s]", err.Error())
	}
	tableColumns, err := dbModel.GetColumns(dbName, tableName)
	if err != nil {
		log.Fatalf("[DbModel.getColumns err:%s]", err.Error())
	}
	template := sql2struct.NewStructTemplate()
	structColumns := template.AssemblyColumns(tableColumns)
	err = template.Generate(tableName, structColumns)
	if err != nil {
		log.Fatalf("[template.generate err:%s]", err.Error())
	}
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&userName, "username", "u", "", "输入数据库账号")
	sql2structCmd.Flags().StringVarP(&passWd, "passwd", "p", "", "输入数据库账号密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "", "输入数据库实例地址")
	sql2structCmd.Flags().StringVarP(&port, "port", "P", "3306", "输入数据库实例端口号")
	sql2structCmd.Flags().StringVarP(&dbType, "dbtype", "", "", "输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "dbname", "d", "", "输入数据库名称")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "输入数据库编码")
	sql2structCmd.Flags().StringVarP(&tableName, "tablename", "t", "", "输入数据库表名")
}
