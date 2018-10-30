package datarw

//"github.com/cyulei/agenda/cmd"
import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cyulei/agenda/entity"
)

// GetMeetings get a []entity.Meeting from a file

func GetMeetings() []entity.Meeting {
	filePath := "datarw/Meetings.json"
	var Meetings []entity.Meeting
	if existFile(filePath) {
		josnStr, err := ioutil.ReadFile(filePath)
		checkError(err)
		//检查是否是空文件
		str := strings.Replace(string(josnStr), "\n", "", 1)
		if str == "" {
			//fmt.Println("Empty")
			return Meetings
		}
		err = json.Unmarshal(josnStr, &Meetings)
		checkError(err)
	}

	return Meetings

}

// SaveMeetings save a []entity.Meeting to a file
func SaveMeetings(MeetingsToSave []entity.Meeting) {
	filePath := "datarw/Meetings.json"
	//清空原文件
	if existFile(filePath) {
		os.Truncate(filePath, 0)
	}
	//转为json串
	josnStr, err := json.Marshal(MeetingsToSave)
	checkError(err)
	err = ioutil.WriteFile(filePath, josnStr, os.ModeAppend)
	checkError(err)
	//开放文件权限
	os.Chmod(filePath, 0777)
}

// TestMeeting is func to test
func TestMeeting() {
	Meetings := GetMeetings()

	participators := []string{"u2", "u1"}

	Meeting1 := entity.Meeting{"456", participators, entity.Date{1, 1, 1, 1, 1}, entity.Date{1, 1, 1, 1, 1}, "4588"}

	Meetings = append(Meetings, Meeting1)
	SaveMeetings(Meetings)
}
