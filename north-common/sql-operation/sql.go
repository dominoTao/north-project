package sql_operation

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"north-project/north-user-baseinfo/pkg/view"
)

const (
	USERNAME = "root"
	//PASSWORD = "liuyaozong123"
	//HOST = "39.106.3.240"
	HOST = "127.0.0.1"
	PASSWORD = "anchor484NB"
	PORT = 3306
	DATABASE = "project"
	//数据库连接的超时时间。
	CONNECT_TIMEOUT = "102s"
)

var DB *sql.DB

func init() {
	_, _ = InitDB()
}

func InitDB() (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s&interpolateParams=true&charset=utf8", USERNAME, PASSWORD, HOST, PORT, DATABASE, CONNECT_TIMEOUT)
	DB, _ = sql.Open("mysql", url)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err:= DB.Ping();err!=nil {
		return nil, fmt.Errorf("open database fail")
	}
	fmt.Println("connect database success")
	return DB, nil
}

func InsertUser(DB *sql.DB, pe view.UserLogin) bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`name`, `pass`) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	// 将参数传递到sql语句中并且执行
	res, err := stmt.Exec(pe.Name, pe.Pass)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	// 提交事务
	tx.Commit()
	// 获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

func DeleteUser(DB *sql.DB, pe view.UserLogin) bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	// 准备sql语句
	stmt, err := tx.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	// 将参数传递到sql语句中并且执行
	res, err := stmt.Exec(pe.Id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	// 提交事务
	tx.Commit()
	// 获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

func UpdateUser(DB *sql.DB, pe view.UserLogin) bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	// 准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET name = ?, pass = ? WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	// 将参数传递到sql语句中并且执行
	res, err := stmt.Exec(pe.Name, pe.Pass, pe.Id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	// 提交事务
	tx.Commit()
	// 获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

func SelectUserById(DB *sql.DB, id int) view.UserLogin {
	var pe view.UserLogin
	err := DB.QueryRow("SELECT * FROM user WHERE id = ?", id).Scan(&pe.Id, &pe.Name, &pe.Pass)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return pe
}

func SelectUserByName(DB *sql.DB, name, pass string) bool {
	var nameS string
	err := DB.QueryRow("SELECT name FROM user WHERE name =  ? AND pass = ?", name, pass).Scan(&nameS)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return len(nameS) > 0
}

func SelectUsers(DB *sql.DB) []view.UserLogin {
	pes := make([]view.UserLogin, 0)
	query, err := DB.Query("SELECT id,name,pass FROM user")
	if err != nil {
		fmt.Println("查询出错了")
	}
	for query.Next() {

		var id int
		var name, pass string
		err := query.Scan(&id, &name, &pass)
		if err != nil {
			fmt.Println("查询失败")
			return pes
		}
		pes = append(pes, view.UserLogin{
			Id: id,
			Name: name,
			Pass: pass,
		})
	}
	return pes
}