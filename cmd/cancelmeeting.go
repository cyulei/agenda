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

	"os"

	"github.com/cyulei/agenda/entity"
	"github.com/spf13/cobra"
)

// cancelmeetingCmd represents the cancelmeeting command
var cancelmeetingCmd = &cobra.Command{
	Use:   "cancelmeeting",
	Short: "Cancellation of a meeting will delete the meeting.",
	Long: `cancelmeeting:If the founder of the conference can cancel the meeting through the title, the meeting will be deleted.
	For example:
	agenda cancelmeeting -t title1
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cancelRun()
		fmt.Println("cancelmeeting called")
	},
}
var cancel_title string

func init() {
	rootCmd.AddCommand(cancelmeetingCmd)

	cancelmeetingCmd.Flags().StringVarP(&cancel_title, "title", "t", "", "yes,even stupid gay knows what it means")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func cancelRun() {
	fileName := "datarw/Agenda.log"
	var logh *log.Logger
	logFile, errlog := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if errlog != nil {
		println("error with open log file ", fileName)
	} else {
		logh = log.New(logFile, "[Info]", log.Ldate|log.Ltime|log.Lshortfile)
		logh.Println("cancelmeeting called")
	}
	//load
	usr := datarw.GetCurUser()
	meetings := datarw.GetMeetings()
	res := make([]entity.Meeting, 0) //for writing

	//adjust the parameters
	if cancel_title == "" {
		fmt.Fprint(os.Stderr, "you must use a -t title to tell me which meeting you want to cancel")
		if errlog == nil {
			logh.SetPrefix("[Error]")
			logh.Println("lack of title parameter when cancel mmeting")
		}
		return
	}
	//check login state
	//if not log in exit
	var loginStatus = usr != nil
	//cancel
	var meetingExist = false
	var authoritySatisified = false

	if !loginStatus {
		println("please login first")
		if errlog == nil {
			logh.SetPrefix("[Error]")
			logh.Println("not log in when cancelmeeting")
		}
		return
	}
	for _, meeting := range meetings {
		if meeting.Title != cancel_title {
			res = append(res, meeting)
			continue
		}

		meetingExist = true
		if usr.Name != meeting.Sponsor {
			res = append(res, meeting)
			continue
		}

		authoritySatisified = true
		//all satisfied
		break
	}
	if !meetingExist {
		println("no such meeting :", cancel_title)
		if errlog == nil {
			logh.SetPrefix("[Warning]")
			logh.Println("user ", usr.Name, " delete a meeting that is not existed, the input title is ", cancel_title)
		}
		return
	}
	if !authoritySatisified {
		println("there is a meeting called :", cancel_title, "but you are not the sponsor of it,please try to relog in ?")
		if errlog == nil {
			logh.SetPrefix("[Warning]")
			logh.Println("not login when cancel meeting:", cancel_title)
		}
		return
	}
	datarw.SaveMeetings(res)
	println("successfully cancel the meeting :", cancel_title)
}
