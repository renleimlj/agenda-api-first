package main

import (
	"agenda-api-first/cli/entity"
	"agenda-api-first/cli/cmd"
	"agenda-api-first/cli/logger"
	"reflect"
	"testing"
)

var users = []entity.User{
	{Name: "u1", Password: "haha", Email: "1@qq.com", Phone: "123456"},
	{Name: "u2", Password: "xixi", Email: "2@qq.com", Phone: "223456"},
	{Name: "u3", Password: "meme", Email: "3@qq.com", Phone: "323456"},
}

var meetings = []entity.Meeting{
	{Title:"m1", Spon:"u1", Pr:["u2", "u3"], Start:"2017-10-01/11:11", End:"2017-10-02/11:11"},
	{Title:"m2", Spon:"u2", Pr:[ "u3"], Start:"2017-01-01/11:11", End:"2017-08-02/11:11"},
	{Title:"m3", Spon:"u3", Pr:["u1"], Start:"2017-11-01/11:11", End:"2017--12/11:11"},
}

func TestRegister(u *entity.User) {
	Logout()
	if res, err := Register(users[0].Name, users[0].Password, users[0].Email, users[0].Phone); err != nil {
		logger.Log("Register() Fail: u1")
	}
	if res, err := Register(users[1].Name, users[1].Password, users[1].Email, users[1].Phone); err != nil {
		logger.Log("Register() Fail: u2")
	}
	if res, err := Register(users[2].Name, users[2].Password, users[2].Email, users[2].Phone); err != nil {
		logger.Log("Register() Fail: u3")
	}
}

func TestLogout(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Logout", true},
	}
	for _, testi := range tests {
		t.Run(testi.name, func(t *testing.T) {
			if got := Logout(); got != testi.want {
				t.Errorf("Logout() = %v, want %v", got, testi.want)
			}
		})
	}
	Logout()
}

func TestLogin(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Login error", args{users[0].Name, users[1].Password}, false},
		{"Login succ", args{users[0].Name, users[0].Password}, true},
	}
	TestRegister(&users[0])
	for _, testi := range tests {
		t.Run(testi.name, func(t *testing.T) {
			if got := Login(testi.args.username, testi.args.password); got != testi.want {
				t.Errorf("Login() = %v, want %v", got, testi.want)
			}
		})
		Logout()
	}
	Logout()
}

func TestQuery(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Query", true},
	}
	for _, testi := range tests {
		t.Run(testi.name, func(t *testing.T) {
			if got := Query(); got != testi.want {
				t.Errorf("Query() = %v, want %v", got, testi.want)
			}
		})
	}
	Logout()
}

func TestCreateMeeting(t *testing.T) {
	type args struct {
		title        string
		username     string
		participator []string
		startDate    string
		endDate      string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"CreateMeeting new", args{meetings[0].Title, users[0].Name, nil, "0000-00-00/00:00", "0001-00-00/00:00"l}, true},
		{"CreateMeetingexit", args{ meetings[0].Title, users[0].Name, nil, "0000-00-00/00:00", "0001-00-00/00:00"}, false},
	}
	TestRegister(&users[0])
	for _, testi := range tests {
		t.Run(testi.name, func(t *testing.T) {
			if got := CreateMeeting(testi.args.title, testi.args.username, testi.args.participator, testi.args.startDate, testi.args.endDate); got != testi.want {
				t.Errorf("CreateMeeting() = %v, want %v", got, testi.want)
			}
		})
	}
	Logout()
}