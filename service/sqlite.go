package service

import (
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name, Password, Email, Phone string
}

type Meeting struct {
  Title string
  Spon string
  Pr []string  
  Start string
  End string
}

func (u *User) Init(name, password, email, phone string) {
	u.Name= name
	u.Password= password
	u.Email= email
	u.Phone= phone
}


/*func main() {
	db, err := sql.Open("sqlite3", "./agenda.db")
    checkErr(err)
	sql_table := `
    CREATE TABLE IF NOT EXISTS userinfo (
	    username TEXT PRIMARY KEY NULL,
	    password TEXT NULL,
	    email TEXT NULL,
	    phone TEXT NULL
	);
	`
    db.Exec(sql_table)

    sql_table = `
    CREATE TABLE IF NOT EXISTS meetinginfo (
    	title TEXT PRIMARY KEY NULL,
	    sponser TEXT NULL,
	    participator TEXT NULL,
	    start TEXT NULL,
	    end TEXT NULL
	);
	`

	db.Exec(sql_table)

	db.Close()

	var u User
	u.Init("ricky", "123", "123", "123")
	createUser(u)
	createUser(u)
	listAllUsers()
	
}*/


func createUser(user User) {
	db, err := sql.Open("sqlite3", "./agenda.db")
    checkErr(err)
	stmt, err := db.Prepare("INSERT INTO userinfo(username, password, email, phone) values(?,?,?,?)")
	checkErr(err)
	_, err = stmt.Exec(user.Name, user.Password, user.Email, user.Phone)
	checkErr(err)
	db.Close()
}

func  listAllUsers() map[string]User{
	db, err := sql.Open("sqlite3", "./agenda.db")
    checkErr(err)
	rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    allUsers := map[string]User{}
    for rows.Next() {
    	var temp User
        err = rows.Scan(&temp.Name, &temp.Password, &temp.Email, &temp.Phone)
        checkErr(err)
    	allUsers[temp.Name] = temp
    }
    rows.Close()
    db.Close()
    return allUsers
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}


