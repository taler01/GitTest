package main

//
//import (
//	"fmt"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//)
//
//type info struct {
//	ID   int
//	NAME string
//	AGE  int
//}
//
//func main() {
//	// 连接MySQL数据库；
//	db, err := gorm.Open("mysql", "root:1009@(127.0.0.1:3306)/cvd_datas?charset=utf8&parseTime=True&loc=Local")
//	if err != nil {
//		panic(err)
//		fmt.Printf("err:%v\n", err)
//	}
//	// 延迟关闭
//	defer db.Close()
//	// 创建表，自动迁移
//	//db.AutoMigrate(&info{})
//	//u1 := info{
//	//	ID:   0,
//	//	NAME: "test",
//	//	AGE:  20,
//	//}
//	//db.Create(u1)
//	// 进行数据库数据插入操作；
//	sql := "insert info (id, name, age) values(1, 龙, 34)"
//	ret, err := db.Exec(sql)
//	if err != nil {
//		fmt.Printf("insert failed,err:%v\n", err)
//		return
//	}
//	id, err := ret.LastInsertId()
//	if err != nil {
//		fmt.Printf("get id failed,err:%v\n", err)
//		return
//	}
//	fmt.Println("id", id)
//
//}
//func Database_Insert() {
//	//sql := "insert info (id, name, age) values(1, "龙", 34)"
//	//ret, err := db
//}
