package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"os"
	"time"
)

//#######################
/*
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

应该wrap这个error，抛给上层
调用了github的上其他库，err直接返还给业务层
另外通过wrap errors 可以给err添加足够的上下文，比如报错的sql语句，
还能保留调用堆栈和root error
在业务层 可以 使用 errors.Cause 获取 root error，再进行和 sentinel error 判定。

*/


// my code
/*
CREATE TABLE `user` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT,
	`name` varchar(45) DEFAULT '',
	`age` int(11) NOT NULL DEFAULT '0',
	`sex` tinyint(3) NOT NULL DEFAULT '0',
	`phone` varchar(45) NOT NULL DEFAULT '',
	PRIMARY KEY (`id`))ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

MariaDB [golang_learn]> select * from user;
+----+-------+-----+-----+-------+
| id | name  | age | sex | phone |
+----+-------+-----+-----+-------+
|  3 | gao   |  10 |   0 |       |
|  4 | wang  |  11 |   0 |       |
|  5 | zhang |  12 |   0 |       |
+----+-------+-----+-----+-------+
3 rows in set (0.001 sec)
*/

const (
	USERNAME = "gaodong"
	PASSWORD = "1q2w3e4r"
	NETWORK  = "tcp"
	SERVER   = "192.168.27.132"
	PORT     = 3306
	DATABASE = "golang_learn"
)

type userInfo struct {
	id int
	name string
	age int
	sex int
	phone string
}

func init_db() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	//fmt.Println(dsn)
	db,err := sql.Open("mysql",dsn)

	if err != nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return nil,err
	}
	db.SetConnMaxLifetime(100*time.Second)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	if err := db.Ping(); err != nil {

		//fmt.Println("open database fail")
		fmt.Printf("Ping database fail,err:%v\n",err)
		return nil,err
	}
	//fmt.Println("connnect success")
	return db,nil
}
func query_name_by_id(db *sql.DB, userId int, name *string) (error) {
	err := db.QueryRow("select name from user where id=?", userId).Scan(name)
	if err != nil {
		//fmt.Printf("db.QueryRow err:%v\n",err)
		//if err == sql.ErrNoRows {
		//	fmt.Println("ok")
		//}
		return errors.Wrapf(err,"select name from user where id=%d", userId)
	}

	return nil
}

func main() {
	db, err := init_db()
	if err != nil {
		fmt.Printf("init_db err:%v\n", err)
		os.Exit(1)
	}
	var user userInfo
	user.id = 1
	err = query_name_by_id(db, user.id, &user.name)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			fmt.Printf("sql.ErrNoRows\n")
		}
		fmt.Printf("err:%v\n", err)
		os.Exit(2)
	}
	fmt.Println(user)
}
//#####my result####
/*
err:select name from user where id=1: sql: no rows in result set
Process finished with exit code 2
*/