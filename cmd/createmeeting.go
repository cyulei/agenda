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
	"strconv"
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
	Long: `Current user can create a meeting. You should provide meeting title, start date and end date of this meeting\n
	and all participators. For example:\n
	createmeeting -t=new_meeting -s=2007-8-3-13-42 -d=2007-8-3-15-42 -p=xxx-xxx-xxx`,
	Run: func(cmd *cobra.Command, args []string) {
		current_user := datarw.GetCurUser()
		if current_user == nil {
			fmt.Println("Please log first")
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
			fmt.Println("Wrong date format. Should be Year-Month-Day-Hour-Minute")
			return
		} else {
			s_date1, flag1 := convert(start_date_string)
			e_date1, flag2 := convert(end_date_string)
			s_date = s_date1
			e_date = e_date1
			if !flag1 || !flag2 {
				fmt.Println("Wrong date format. Should be Year-Month-Day-Hour-Minute")
				return
			}
		}
		var temp_meeting entity.Meeting
		temp_meeting.Startdate = s_date
		temp_meeting.Enddate = e_date
		if !isParticipatorAvailable(current_user.Name, meetings, temp_meeting) {
			fmt.Println("Sponsor is not free")
			return
		}

		valid_participators, ok := check_participators(change_participators, users, meetings, s_date, e_date)
		if check_title(create_meeting_title, meetings) && check_date(s_date, e_date) && ok {
			var new_meeting entity.Meeting
			new_meeting.Sponsor = current_user.Name
			new_meeting.Title = create_meeting_title
			new_meeting.Startdate = s_date
			new_meeting.Enddate = e_date
			new_meeting.Participators = valid_participators
			meetings = append(meetings, new_meeting)
			datarw.SaveMeetings(meetings)
			fmt.Println("createmeeting finished")
		} else {
			if !check_title(create_meeting_title, meetings) {
				fmt.Println("Meeting exists, change meeting title")
			}
			if !check_date(s_date, e_date) {
				fmt.Println("Invalid start/end date, please check")
			}
			if !ok {
				fmt.Println("No valid participators (busy or not exists)")
			}
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

func convert(date_string []string) (entity.Date, bool) {
	year, flag1 := strconv.Atoi(date_string[0])
	month, flag2 := strconv.Atoi(date_string[1])
	day, flag3 := strconv.Atoi(date_string[2])
	hour, flag4 := strconv.Atoi(date_string[3])
	minute, flag5 := strconv.Atoi(date_string[4])

	var date entity.Date
	if flag1 == nil && flag2 == nil && flag3 == nil && flag4 == nil && flag5 == nil {
		date.Year = year
		date.Month = month
		date.Day = day
		date.Hour = hour
		date.Minute = minute
		return date, true
	} else {
		return date, false
	}

}

func check_participators(participators []string, all_users []entity.User, all_meetings []entity.Meeting, s_date entity.Date, e_date entity.Date) ([]string, bool) {
	var valid_participators []string
	var temp_meeting entity.Meeting
	temp_meeting.Startdate = s_date
	temp_meeting.Enddate = e_date

	for _, j := range participators {
		if isParticipatorExist(j, all_users) && j != sponsor_name {
			if isParticipatorAvailable(j, all_meetings, temp_meeting) {
				valid_participators = append(valid_participators, j)
			}
		}
	}
	if len(valid_participators) == 0 {
		return valid_participators, false
	} else {
		return valid_participators, true
	}
}

func check_title(meeting_title string, all_meetings []entity.Meeting) bool {
	for _, j := range all_meetings {
		if j.Title == meeting_title {
			return false
		}
	}
	return true
}

func check_date(date1 entity.Date, date2 entity.Date) bool {
	if entity.IsValid(date1) && entity.IsValid(date2) {
		if entity.Compare(date2, date1) >= 0 {
			return true
		}
	}
	return false
}
