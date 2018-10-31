// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"github.com/cyulei/agenda/datarw"
	"github.com/cyulei/agenda/entity"
	"github.com/spf13/cobra"
)

// deleteuserCmd represents the deleteuser command
var deleteuserCmd = &cobra.Command{
	Use:   "deleteuser",
	Short: "delete CurUser",
	Long: `deleteuser:delete CurUser and logout
	you must login before deleteuser
	For example:
	agenda deleteuser  			:delete CurUser and logout
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteuser called")
		deleteuser()
	},
}

func init() {
	rootCmd.AddCommand(deleteuserCmd)

}
func deleteuser() {
	curUser := datarw.GetCurUser()
	if curUser == nil { //是否已登陆
		fmt.Println("isn't login,please use command login")
		return
	}

	//获取所有用户
	users := datarw.GetUsers()

	for index, user := range users {
		if user.Name == curUser.Name {
			users = append(users[:index], users[index+1:]...)
			datarw.SaveUsers(users)
			datarw.SaveCurUser(nil) //登出
			fmt.Println("User:", curUser.Name, " has been deleted")

			/*会议相关*/

			cancleAllmeeting(*curUser) //当前用户取消所有其创建的会议
			exitAllmeeting(*curUser)   //当前用户退出所有会议

			return
		}
	}

	fmt.Println("error: unexpected to execute")

}

//取消user创建的会议
func cancleAllmeeting(user entity.User) {
	meetings := datarw.GetMeetings()
	var newMeetings []entity.Meeting

	for _, meeting := range meetings { //遍历会议
		if meeting.Sponsor != user.Name {
			newMeetings = append(newMeetings, meeting)
		}
	}

	datarw.SaveMeetings(newMeetings)
}

//user退出所有会议
func exitAllmeeting(user entity.User) {
	meetings := datarw.GetMeetings()

	for k, meeting := range meetings { //遍历会议
		for i, participator := range meeting.Participators { //遍历一个会议的成员

			if participator == user.Name {
				meetings[k].Participators = append(meetings[k].Participators[:i], meetings[k].Participators[i+1:]...)

				break //!!!
			}
		}

	}

	meetings = deleteEmptyMeeting(meetings)

	datarw.SaveMeetings(meetings)

}

func deleteEmptyMeeting(meetings []entity.Meeting) []entity.Meeting {
	var newMeetings []entity.Meeting

	for _, meeting := range meetings { //当会议成员不为空时，保留会议
		if len(meeting.Participators) != 0 {
			newMeetings = append(newMeetings, meeting)
		}
	}

	return newMeetings
}
