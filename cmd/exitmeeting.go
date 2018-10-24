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

<<<<<<< HEAD
=======
	"github.com/cyulei/agenda/datarw"

>>>>>>> d823cadd36789c7ecde564f0e4d65da972ab2ebc
	"github.com/cyulei/agenda/entity"
	"github.com/spf13/cobra"
)

// exitmeetingCmd represents the exitmeeting command
var exitmeetingCmd = &cobra.Command{
	Use:   "exitmeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	var res []entity.Meeting
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

	for i := 0; i < len(meetings); i++ {
		mt := &meetings[i]
		if mt.Title != exit_title {
			continue
		}
		meetingExist = true

		if mt.Sponsor == usr {
			delete = true
			pos = i
			break
		}

		pts := mt.Participators
		for j := 0; j < len(pts); j++ {
			pt := pts[j]
			if pt == usr { //usr is a participator of meeting mt,we need to remove usr from mt's participators
				inmeeting = true
				//remove
				parts := make([]string, len(pts))

				parts = append(parts, pts[0:j]...)
				parts = append(parts, pts[j:]...)
				//mt.participators = parts
				mt.Participators = parts
				//break
				j = len(pts) + 100
				i = len(meetings) + 100
			}
		}

	}
	if meetingExist == false {
		fmt.Println("no such meeting,", exit_title, "please check you title spelling")
		//saveMeetings(meetings)
		//meeting_title
		return
	}
	if delete == true {
		fmt.Println("you are the sponsor of the meeting,yes you are sure to delete(cancel) the meeting")
		//delete
		res = append(meetings[0:pos], meetings[pos:]...)
		datarw.SaveMeetings(res)
		return
	}
	//res =meetings
	if inmeeting == false {
		fmt.Println("you are not in meeting of title :", exit_title, " please check your spelling ?")
		return
	}
	fmt.Println("successfully exit the meeting:", exit_title)
	datarw.SaveMeetings(meetings)
	return
}
