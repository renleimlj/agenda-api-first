package service

import (
    "testing"
)

func Test_UserOperations(t *testing.T) {
	t.Log("Test User")
	var u User
	u.Init("hero", "123", "13909209949@163.com", "13909209949")
	err := CreateUser(u)
	users := ListAllUsers()
	if err != nil {
		t.Error(err.Error())
		t.Error("User测试未通过")
	} else {
		_ , ok := users[u.Name]
		if !ok {
			t.Error("User测试未通过")
		} else {
			t.Log("User测试通过")
		}
	}
}

func Test_MeetingOperations(t *testing.T) {
	t.Log("Test Meeting")
	var m Meeting
	m.Title = "testmeeting"
	m.Start = "2017-12-15/00:00"
	m.End = "2017-12-15/00:30"
	m.Spon = "ricky"
	m.Pr = []string{"kity", "land"}  
	err := CreateMeeting(m)
	meetings := ListAllMeetings()
	if err != nil {
		t.Error(err.Error())
		t.Error("Meeting测试未通过")
	} else {
		_, ok := meetings[m.Title]
		if !ok {
			t.Error("Meeting测试未通过")
		} else {
			t.Log("MeetingCreate测试通过")
		}
	}
	err = DeleteMeeting(m.Title) 
	if err != nil {
		t.Error(err.Error())
		t.Error("Meeting测试未通过")
	} else {
		_, ok := meetings[m.Title]
		if !ok {
			t.Error("MeetingDelete测试通过")
		} else {
			t.Log("MeetingDelete测试未通过")
		}
	}
}

func Test_KeyOperations(t *testing.T) {
	t.Log("Test Key")
	CreateKey("123456789", "hero")
	temp := FindUserbyKey("123456789")
	if temp != "hero" {
		t.Error("Key测试未通过")
	}
	flag := isLogin("hero")
	if !flag {
		t.Error("Key测试未通过")
	}
	err := DeleteKey("123456789")
	if err != nil {
		t.Error("Key测试未通过")
	}
	temp = FindUserbyKey("123456789")
	if temp != "" {
		t.Error("Key测试未通过")
	}
	t.Log("Key测试完成")
}