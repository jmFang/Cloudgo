package entities

import (
	"fmt"
	"testing"
)

func TestService(t *testing.T) {
	//test saving
	user := NewUserInfo(UserInfo{UserName: "sysu", DepartName: "sdcs"})
	err := UserInfoService.Save(user)
	if err != nil {
		t.Error("=== Test 'Save' fails\n")
	} else {
		fmt.Printf("=== Test 'Save' Successfully!\n")
	}
	//tesing finding
	users := UserInfoService.FindAll()
	size := len(users)
	if size > 0 {
		if users[size-1].UserName == "sysu" && users[size-1].DepartName == "sdcs" {
			fmt.Printf("=== Test 'FindAll' Successfully!. Finding result is : %v \n", users[size-1])
		} else {
			t.Errorf("=== Test 'FindAll' fails. Username: %q, DepartName: %q not the wanted result, Username: %q, DepartName: %q\n", users[size-1].UserName, users[size-1].DepartName, user.UserName, user.DepartName)
		}
	}

	// tesing finding by id
	user1 := UserInfoService.FindByID(user.UID)
	if user1.UID != user.UID {
		t.Errorf("=== Test 'FindById' fails. Found id: %v not the wanted id: %v \n", user1.UID, user.UID)
	} else {
		fmt.Print("=== Test 'FindById' Successfully!.")
	}
}
