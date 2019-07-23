package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"time"
)

var (
	dbhostip   = "192.168.20.235"
	dbusername = "gotest"
	dbpassword = "Xiodi.cn123"
	dbname     = "gotest"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
		fmt.Println("err:", err)
	}
}

func OpenDB() (success bool, db *sql.DB) {
	var isOpen bool
	db, err := sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhostip+")/"+dbname+"?charset=utf8")
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	CheckErr(err)
	return isOpen, db
}

func Hello(response http.ResponseWriter, request *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(response, "Error: %s", err)
	} else {
		fmt.Fprintf(response, "Hostnames: %s", host)
	}
}

func Insert(response http.ResponseWriter, request *http.Request) {
	opened, db := OpenDB()
	if opened {
		fmt.Println("open success")
	} else {
		fmt.Println("open faile:")
	}

	info := insertToDB(db)
	println(info)

	fmt.Fprintf(response, info)
}

func insertToDB(db *sql.DB) (info string) {
	nowTimeStr := time.Now()
	stmt, err := db.Prepare("insert userinfo set username=?,departname=?,created=?,password=?")
	CheckErr(err)
	res, err := stmt.Exec("test", "it", nowTimeStr, "123456")
	CheckErr(err)
	id, err := res.LastInsertId()
	CheckErr(err)
	var insert_info string
	if err != nil {
		insert_info = "插入数据失败"
	} else {
		insert_info = "插入数据成功"
	}

	db.Close()
	println(id)
	return insert_info
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/insert", Insert)
	http.ListenAndServe(":8080", nil)
}
