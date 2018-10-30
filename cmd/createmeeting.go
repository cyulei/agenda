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
	"strings"

	"github.com/cyulei/agenda/datarw"
	"github.com/cyulei/agenda/entity"
	"github.com/spf13/cobra"
)

var (
	sponsor_name         string
	create_meeting_title string
	create_start_date    string
	create_end_date      string
	create_participators string
)

// createmeetingCmd represents the createmeeting command
var createmeetingCmd = &cobra.Command{
	Use:   "createmeeting",
	Short: "Create a meeting",
	Long: `Current user can create a meeting. You should provide meeting title, start date and end date of this meeting
and all participators. For example:
createmeeting -t=new_meeting -s=2007-8-3-13-42 -e=2007-8-3-15-42 -p=xxx-xxx-xxx`,
	Run: func(cmd *cobra.Command, args []string) {
		//log
		fileName := "datarw/Agenda.log"
		logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
		defer logFile.Close()
		if err != nil {
			log.Fatalln("Open file error")
		}
		infoLog := log.New(logFile, "[Info]", log.Ldate|log.Ltime|log.Lshortfile)
		infoLog.Println("Cmd createmeeting called")

		current_user := datarw.GetCurUser()
		if current_user == nil {
			infoLog.SetPrefix("[Error]")
			infoLog.Println("Not log in yet")
			fmt.Println("Please log first")
			infoLog.Println("Cmd createmeeting failed")
			fmt.Println("createmeeting failed")
			return
		}
		sponsor_name = current_user.Name
		meetings := datarw.GetMeetings()
		users := datarw.GetUsers()
		change_participators := strings.Split(create_participators, "-")
		start_date_string := strings.Split(create_start_date, "-")
		end_date_string := strings.Split(create_end_date, "-")
		var s_date, e_date entity.Date
		if len(start_date_string) != 5 || len(end_date_string) != 5 {
			infoLog.SetPrefix("[Error]")
			infoLog.Println("Wrong date format")
			fmt.Println("Wrong date format. Should be Year-Month-Day-Hour-Minute")
			infoLog.Println("Cmd createmeeting failed")
			fmt.Println("createmeeting failed")
			return
		} else {
			s_date1, flag1 := entity.Convert(start_date_string)
			e_date1, flag2 := entity.Convert(end_date_string)
			s_date = s_date1
			e_date = e_date1
			if !flag1 || !flag2 {
				infoLog.SetPrefix("[Error]")
				infoLog.Println("Wrong date format")
				fmt.Println("Wrong date format. Should be Year-Month-Day-Hour-Minute")
				infoLog.Println("Cmd createmeeting failed")
				fmt.Println("createmeeting failed")
				return
			}
		}
		var temp_meeting entity.Meeting
		temp_meeting.Startdate = s_date
		temp_meeting.Enddate = e_date
		if !entity.IsParticipatorAvailable(current_user.Name, meetings, temp_meeting) {
			infoLog.SetPrefix("[Error]")
			infoLog.Println("Sponsor not free")
			fmt.Println("Sponsor is not free")
			infoLog.Println("Cmd createmeeting failed")
			fmt.Println("createmeeting failed")
			return
		}

		valid_participators, ok := entity.Check_participators(sponsor_name, change_participators, users, meetings, s_date, e_date)
		if entity.Check_title(create_meeting_title, meetings) && entity.Check_date(s_date, e_date) && ok {
			var new_meeting entity.Meeting
			new_meeting.Sponsor = current_user.Name
			new_meeting.Title = create_meeting_title
			new_meeting.Startdate = s_date
			new_meeting.Enddate = e_date
			new_meeting.Participators = valid_participators
			meetings = append(meetings, new_meeting)
			datarw.SaveMeetings(meetings)
			infoLog.SetPrefix("[Info]")
			infoLog.Println("Cmd createmeeting finished")
			fmt.Println("createmeeting finished")
		} else {
			if !entity.Check_title(create_meeting_title, meetings) {
				infoLog.SetPrefix("[Error]")
				infoLog.Println("Meeting exists")
				fmt.Println("Meeting exists, change meeting title")
			}
			if !entity.Check_date(s_date, e_date) {
				infoLog.SetPrefix("[Error]")
				infoLog.Println("Invalid start/end date")
				fmt.Println("Invalid start/end date, please check")
			}
			if !ok {
				infoLog.SetPrefix("[Error]")
				infoLog.Println("No valid participators")
				fmt.Println("No valid participators (busy or not exists)")
			}
			infoLog.SetPrefix("[Error]")
			infoLog.Println("Cmd createmeeting failed")
			fmt.Println("createmeeting failed")
		}

	},
}

func init() {
	rootCmd.AddCommand(createmeetingCmd)
	createmeetingCmd.Flags().StringVarP(&create_meeting_title, "title", "t", "", "meeting title")
	createmeetingCmd.MarkFlagRequired("title")
	createmeetingCmd.Flags().StringVarP(&create_start_date, "start", "s", "", "meeting start date")
	createmeetingCmd.MarkFlagRequired("tart")
	createmeetingCmd.Flags().StringVarP(&create_end_date, "end", "e", "", "meeting end date")
	createmeetingCmd.MarkFlagRequired("end")
	createmeetingCmd.Flags().StringVarP(&create_participators, "part", "p", "", "meeting participators")
	createmeetingCmd.MarkFlagRequired("part")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
