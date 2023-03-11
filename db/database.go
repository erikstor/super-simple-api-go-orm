package db

import (
	"fmt"
	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var dsn = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// var dsn = "root:@tcp(localhost:3306)/test-go?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println("Error en la conexión de la base de datos")
	//	// todo buscar un servicio para conectar con la base de datos
	//	//panic(err)
	//}
	fmt.Println("Conexión exitosa")
	return db
}()
