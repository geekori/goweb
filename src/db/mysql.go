package main

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*

MySQL数据库

https://github.com/go-sql-driver/mysql

安全MySQL驱动

go get -u github.com/go-sql-driver/mysql
-u：强制升级

 */
func main() {
	db, err :=  sql.Open("mysql", "root:12345678@/meituan?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return
	}
	db.Exec("drop table t_products")
	sqlStmt := `create table t_products(id integer not null primary key, name text,price float)`
	//  创建表
	_,err = db.Exec(sqlStmt)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	//  开始事务
	tx, err := db.Begin()


	if err != nil {
		log.Fatal(err)
		return
	}

	stmt, err := tx.Prepare("insert into t_products(id,name,price) values(?,?,?)")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer stmt.Close()
	for i:= 0; i < 10;i++ {
		_,err = stmt.Exec(i+1,fmt.Sprintf("产品%d", i + 1),float64(i+1) * 54.8)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	//  提交事务
	tx.Commit()

	//  查询
	rows,err := db.Query("select id,name,price from t_products")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var price float64
		err = rows.Scan(&id, &name, &price)
		if err != nil {
			log.Fatal(err)
			return
		}
		priceStr := fmt.Sprintf("%.2f",price)
		fmt.Println(id,name,priceStr)
	}

	stmt,_ = db.Prepare("select name,price from t_products where id=?")
	defer stmt.Close()
	var name string
	var price float64

	stmt.QueryRow("6").Scan(&name, &price)
	priceStr := fmt.Sprintf("%.2f",price)
	fmt.Println("-----------------")
	fmt.Println(name,priceStr)

	//  删除记录
	stmt,_ = db.Prepare("delete from t_products where id=?")

	stmt.Exec(9)



}


