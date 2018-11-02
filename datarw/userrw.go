package datarw

//"github.com/cyulei/agenda/cmd"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cyulei/agenda/entity"
)

var UsersfilePath string = "./datarw/Users.json"
var CurUserfilePath string = "./datarw/CurUser.json"

// GetUsers get a []entity.User from a file
func GetUsers() []entity.User {
	var users []entity.User
	if existFile(UsersfilePath) {
		josnStr, err := ioutil.ReadFile(UsersfilePath)
		checkError(err)
		//检查是否是空文件
		str := strings.Replace(string(josnStr), "\n", "", 1)
		if str == "" {
			//fmt.Println("Empty")
			return users
		}
		err = json.Unmarshal(josnStr, &users)
		checkError(err)
	}

	return users
}

// SaveUsers save a []entity.User to a file
func SaveUsers(usersToSave []entity.User) {
	//清空原文件
	os.Truncate(UsersfilePath, 0)

	//转为json串
	josnStr, err := json.Marshal(usersToSave)
	checkError(err)
	err = ioutil.WriteFile(UsersfilePath, josnStr, os.ModeAppend)
	checkError(err)
	//开放文件权限
	os.Chmod(UsersfilePath, 0777)
}

// GetCurUser get a *entity.User from a file
func GetCurUser() *entity.User {

	var user entity.User

	if existFile(CurUserfilePath) {
		//读取Json串
		josnStr, err := ioutil.ReadFile(CurUserfilePath)
		checkError(err)
		//检查是否是空文件
		str := strings.Replace(string(josnStr), "\n", "", 1)
		if str == "" {
			//fmt.Println("Empty")
			return nil
		}
		//解析Json串
		err = json.Unmarshal(josnStr, &user)
		checkError(err)

		return &user
	}

	return nil

}

// SaveCurUser save a entity.User to a file
func SaveCurUser(userToSave *entity.User) {
	//清空原文件
	os.Truncate(CurUserfilePath, 0)

	if userToSave != nil {
		//转为json串
		josnStr, err := json.Marshal(userToSave)
		checkError(err)
		err = ioutil.WriteFile(CurUserfilePath, josnStr, os.ModeAppend)
		checkError(err)
		//开放文件权限
		os.Chmod(CurUserfilePath, 0777)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
	}
}
func existFile(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// TestUser is func to test
func TestUser() {
	users := GetUsers()
	user1 := entity.User{"456", "456", "456", "4588"}
	users = append(users, user1)
	SaveUsers(users)
}
