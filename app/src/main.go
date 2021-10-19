package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	_, err := gorm.Open(mysql.Open("go_test:p@55w0rd@tcp(mysql_container:3306)/go_database?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database!")
		panic(err.Error())
	}

	fmt.Println(`
。
+　。☆・　+
・゜☆・　Environment　｡☆ °
+　☆。☆　building　☆
+　☆。☆　completed!　☆
ﾟ　|。・/☆　。+
r⌒|ヽﾟ/。／・☆
|丶|　/'／。☆
|=='-/∠_〉。+
|======／
|====／∧∞∧
|==／（｡･ω･｡）
|／
	`)

}
