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
	"os"

	"github.com/cyulei/agenda/datarw"
	"github.com/cyulei/agenda/entity"
	"github.com/spf13/cobra"
)

// exitmeetingCmd represents the exitmeeting command
var exitmeetingCmd = &cobra.Command{
	Use:   "exitmeeting",
	Short: "Exit from the meeting as a member.",
	Long: `exitmeeting : you must login first , if you are the sponser of the meeting,it'll be canceled without assertain.
	For example:
	agenda exitmeeting -t=title1
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//flag.Parse()
		runExit()
		fmt.Println("exitmeeting called")
	},
}
var exit_title string

func init() {
	rootCmd.AddCommand(exitmeetingCmd)
	exitmeetingCmd.Flags().StringVarP(&exit_title, "title", "t", "empty title", "input the title of meeting you want to exit")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exitmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exitmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runExit() {

	//load data
	curUsr := datarw.GetCurUser()
	if curUsr == nil {
		println("please login first ")
		return
	}
	usr := curUsr.Name
	meetings := datarw.GetMeetings()
	var res = make([]entity.Meeting, 0)
	//check title
	if exit_title == "empty title" {
		fmt.Println("please input the title ")
		return
	}

	//find the meeting
	delete := false
	pos := -1
	inmeeting := false
	meetingExist := false
	meetingEmpty := false

	for i := 0; i < len(meetings); i++ {
		pos = i
		mt := &meetings[i]
		if mt.Title != exit_title {
			continue
		}
		meetingExist = true

		if mt.Sponsor == usr {
			delete = true

			break
		}

		pts := mt.Participators
		for j := 0; j < len(pts); j++ {
			pt := pts[j]
			if pt == usr { //usr is a participator of meeting mt,we need to remove usr from mt's participators
				inmeeting = true
				//remove
				parts := make([]string, 0)

				parts = append(parts, pts[0:j]...)
				parts = append(parts, pts[j+1:]...)
				//mt.participators = parts
				mt.Participators = parts

				if len(mt.Participators) == 0 {
					meetingEmpty = true
				}
				//break
				j = len(pts) + 100
				i = len(meetings) + 100
			}
		}

	}
	fileName := "datarw/Agenda.log"
	var logh *log.Logger
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		println("error with open log file ", fileName)
	} else {
		logh = log.New(logFile, "[Info]", log.Ldate|log.Ltime|log.Lshortfile)
		logh.Println("exitmeeting called")
	}

	if meetingExist == false {
		fmt.Println("no such meeting,", exit_title, "please check you title spelling")
		if err == nil {
			logh.SetPrefix("[Error]")
			logh.Println("exitmeeting a  meeting not exist")
		}

		//saveMeetings(meetings)
		//meeting_title
		return
	}
	if delete == true {
		fmt.Println("you are the sponsor of the meeting,yes you are sure to delete(cancel) the meeting")
		if err == nil {
			logh.SetPrefix("[Warning]")
			logh.Println("user ", usr, " exit  meeting ", exit_title, ",which is sponsed by himself/herself .meeting ", exit_title,
				" has been deleted now because of no sponser")
		}
		//delete
		res = append(meetings[0:pos], meetings[pos+1:]...)
		datarw.SaveMeetings(res)
		return
	}
	//res =meetings
	if inmeeting == false {
		fmt.Println("you are not in meeting of title :", exit_title, " please check your spelling ?")
		if err == nil {
			logh.SetPrefix("[Info]")
			logh.Println("user:", usr, " want to exitmeeting ", exit_title, " but he/she does not take part in it,and be surely refused")
		}
		return
	}
	if meetingEmpty == true {
		fmt.Println("meeting delete for no participators in")
		if err == nil {
			logh.SetPrefix("[Warning]")
			logh.Println("meeting ", exit_title, " has been deleted because of no participator")
		}
		res = append(meetings[0:pos], meetings[pos+1:]...)
		datarw.SaveMeetings(res)
		return
	}
	fmt.Println("successfully exit the meeting:", exit_title)
	datarw.SaveMeetings(meetings)
	return
}
