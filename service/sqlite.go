package service

import (
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
    "strings"
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

/*
func main() {
	db, err := sql.Open("sqlite3", "../agenda.db")
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

	sql_table = `
    CREATE TABLE IF NOT EXISTS keyinfo (
	    key TEXT PRIMARY KEY NULL,
	    name TEXT NULL
	);
	`
    db.Exec(sql_table)
    db.Close()

	var u User
	u.Init("ricky", "123", "123", "123")
	createUser(u)
	listAllUsers()

	var u User
	u.Init("kity", "123", "123", "123")
	createUser(u)

	var user User
	user.Init("land", "123", "123", "123")
	createUser(user)

	var m Meeting
	m.Title = "m1"
	m.Start = "2017-12-14/00:00"
	m.End = "2017-12-14/00:30"
	m.Spon = "ricky"
	m.Pr = []string{"kity", "land"}  
	createMeeting(m)
	listAllMeetings()
	
}
*/


func CreateUser(user User) error {
	db, err := sql.Open("sqlite3", "agenda.db")
    checkErr(err)
    if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO userinfo(username, password, email, phone) values(?,?,?,?)")
	checkErr(err)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Name, user.Password, user.Email, user.Phone)
	checkErr(err)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func  ListAllUsers() map[string]User{
	db, err := sql.Open("sqlite3", "agenda.db")
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

func CreateMeeting(meeting Meeting) error {
	db, err := sql.Open("sqlite3", "agenda.db")
	if err != nil {
		return err
	}
    checkErr(err)
    stmt, err := db.Prepare("INSERT INTO meetinginfo(title, sponser, participator, start, end) values(?,?,?,?,?)")
    if err != nil {
		return err
	}
	checkErr(err)
	temp := strings.Join(meeting.Pr, "/")
	_, err = stmt.Exec(meeting.Title, meeting.Spon, temp, meeting.Start, meeting.End)
	if err != nil {
		return err
	}
	checkErr(err)
	db.Close()
	return nil
}

func split(s rune) bool {
  if s == '/' {
    return true
  }
  return false
}

func ListAllMeetings() map[string]Meeting{
	db, err := sql.Open("sqlite3", "agenda.db")
    checkErr(err)
	rows, err := db.Query("SELECT * FROM meetinginfo")
    checkErr(err)
    allMeetings := map[string]Meeting{}
    for rows.Next() {
    	var temp Meeting
    	var par string
        err = rows.Scan(&temp.Title, &temp.Spon, &par, &temp.Start, &temp.End)
        checkErr(err)
        temp.Pr = strings.FieldsFunc(par, split)
    	allMeetings[temp.Title] = temp
    }
    rows.Close()
    db.Close()
    return allMeetings
}

func DeleteMeeting(title string) error{
	db, err := sql.Open("sqlite3", "agenda.db")
    checkErr(err)
    if err != nil {
		return err
	}
	stmt, err := db.Prepare("delete from meetinginfo where title=?")
	checkErr(err)
	if err != nil {
		return err
	}
	res, err1 := stmt.Exec(title)
	checkErr(err1)
	if err1 != nil {
		return err1
	}
	_, err1 = res.RowsAffected()
	checkErr(err1)
	if err1 != nil {
		return err1
	}
	db.Close()
	return nil
}

func CreateKey(key ,name string) error {
	db, err := sql.Open("sqlite3", "agenda.db")
    checkErr(err)
    if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO keyinfo(key, name) values(?,?)")
	checkErr(err)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(key, name)
	checkErr(err)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func DeleteKey(key string) error {
	db, err := sql.Open("sqlite3", "agenda.db")
    checkErr(err)
    if err != nil {
		return err
	}
	stmt, err := db.Prepare("delete from keyinfo where key=?")
	checkErr(err)
	if err != nil {
		return err
	}
	res, err1 := stmt.Exec(key)
	checkErr(err1)
	if err1 != nil {
		return err1
	}
	_, err1 = res.RowsAffected()
	checkErr(err1)
	if err1 != nil {
		return err1
	}
	db.Close()
	return nil
}

func FindUserbyKey(target_key string) string {
	db, err := sql.Open("sqlite3", "agenda.db")
    checkErr(err)
	rows, err := db.Query("SELECT * FROM keyinfo")
    checkErr(err)
    for rows.Next() {
    	var key string
    	var name string
        err = rows.Scan(&key, &name)
        checkErr(err)
        if key == target_key {
        	rows.Close()
    		db.Close()
         	return name
    	}
    }
    rows.Close()
    db.Close()
    return ""
}

func isLogin(target_name string) bool {
	db, err := sql.Open("sqlite3", "agenda.db")
    checkErr(err)
	rows, err := db.Query("SELECT * FROM keyinfo")
    checkErr(err)
    for rows.Next() {
    	var key string
    	var name string
        err = rows.Scan(&key, &name)
        checkErr(err)
        if name == target_name {
        	rows.Close()
    		db.Close()
         	return true
    	}
    }
    rows.Close()
    db.Close()
    return false
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}