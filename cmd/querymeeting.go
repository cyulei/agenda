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

	"github.com/cyulei/agenda/datarw"
	"github.com/cyulei/agenda/entity"
	"github.com/spf13/cobra"
)

// querymeetingCmd represents the querymeeting command
var querymeetingCmd = &cobra.Command{
	Use:   "querymeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("querymeeting called")
		runQuery()
	},
}
var query_title string
var query_sDate string
var query_eDate string
var query_all bool

func init() {
	rootCmd.AddCommand(querymeetingCmd)
	querymeetingCmd.Flags().StringVarP(&query_title, "title", "t", "", "the title you want to query")
	querymeetingCmd.Flags().StringVarP(&query_sDate, "start time", "s", "", "format yyyy-mm-dd-hh:mm")
	querymeetingCmd.Flags().StringVarP(&query_eDate, "end time", "e", "", "format yyyy-mm-dd-hh:mm")
	querymeetingCmd.Flags().BoolVarP(&query_all, "all user or current user", "a", false, "query meetings "+
		"all user has been appeared,if you want query for current user,please don't use it")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// querymeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// querymeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func runQuery() {
	//	load
	//datarw.SaveMeetings
	var meetings = datarw.GetMeetings()
	var usr = datarw.GetCurUser()

	var res = make([]entity.Meeting, 0) //to display

	//parse the arguments
	var title_limited = query_title != ""
	//var time_limited = false
	var start_limited = false
	var end_limited = false
	var usr_limited = query_all
	var usr_logged = usr != nil

	var sdate = entity.Date{}
	var edate = entity.Date{}

	if usr_limited && !usr_logged { //想要查询当前登录的用户的会议但是又没有登录
		println("you can not query for the current logged user,please login first! ")
		return
	}
	if query_sDate != "" {
		start_limited = true
	} else if query_eDate != "" {
		end_limited = true
	}
	//time_limited = start_limited || end_limited

	if len(query_sDate) != 16 && start_limited || len(query_eDate) != 16 && end_limited {
		println("date format error,yyyy-mm-dd-hh:mm")
		return
	}
	//func date
	if start_limited {
		var err, err1, err2, err3, err4 error
		sdate.Year, err = strconv.Atoi((string)(query_sDate[0:4]))
		sdate.Month, err1 = strconv.Atoi(query_sDate[5:7])
		sdate.Day, err2 = strconv.Atoi(query_sDate[8:10])
		sdate.Hour, err3 = strconv.Atoi(query_sDate[11:13])
		sdate.Minute, err4 = strconv.Atoi(query_sDate[14:16])

		if err != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			println("date format error,yyyy-mm-dd-hh:mm")
			return
		}
		if !entity.IsValid(sdate) {
			println(sdate.Year, sdate.Month, sdate.Day, sdate.Hour, sdate.Minute)
			println("check your date number,pay attention to max day of the month and care for the leap year")
			return
		}
	}
	if end_limited {
		var err, err1, err2, err3, err4 error
		edate.Year, err = strconv.Atoi((string)(query_eDate[0:4]))
		edate.Month, err1 = strconv.Atoi(query_eDate[5:7])
		edate.Day, err2 = strconv.Atoi(query_eDate[8:10])
		edate.Hour, err3 = strconv.Atoi(query_eDate[11:13])
		edate.Minute, err4 = strconv.Atoi(query_eDate[14:16])

		if err != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			println("date format error,yyyy-mm-dd-hh:mm")
			return
		}
		if entity.IsValid(edate) {
			println("check your date number,pay attention to max day of the month and care for the leap year")
			return
		}
	}

	//end
	//start querying

	for _, meeting := range meetings {
		if title_limited && meeting.Title != query_title { //has limitation on title but not satisfied
			continue
		}
		//-----------------------not satisfied date compare-------------------------------------------------------------------
		//no title limitation or has title limitation and already satisfied
		if start_limited && entity.Compare(sdate, meeting.Startdate) > 0 { //has limitation on start date but not satisfied
			continue
		}
		if end_limited && entity.Compare(edate, meeting.Enddate) < 0 { //after the given edate ,which is not  supposed
			continue
		}
		if usr_limited {

			if usr.Name == meeting.Sponsor {
				//we add this meeting to result
			} else {
				var f = false
				for _, parts := range meeting.Participators {
					if parts == usr.Name { //satisfied we can display it
						f = true
						break
					}
				}
				if f == false { // not satisfied we cannot display this meeting
					continue
				}
			}
		}
		//all request satisfied
		res = append(res, meeting)
	}
	DisplayMeeting(res)
}

func DisplayMeeting(mt []entity.Meeting) {

	standardMeetingLength := 12
	standardNameLength := 8
	//standardTimeLength := 16
	println("-----------------Display Meeting---------------------------")
	println("Title\t\t\tSponsor\t\t\tStart Time\t\tEnd Time\t\tParticipators")
	for _, meeting := range mt {
		print(meeting.Title)
		for j := 4; j <= standardMeetingLength; j += 4 {
			if len(meeting.Title) < j {
				for k := j - 4; k < standardMeetingLength; k += 4 {
					print("\t")
				}
			}
		}
		//print("\t\t")
		print(meeting.Sponsor)
		for j := 4; j <= standardNameLength; j += 4 {
			if len(meeting.Sponsor) < j {
				for k := j - 4; k < standardNameLength; k += 4 {
					print("\t")
				}
			}
		}
		//print("\n")
		sd := meeting.Startdate
		ed := meeting.Enddate
		year := sd.Year
		month := sd.Month
		day := sd.Day
		hour := sd.Hour
		minute := sd.Minute
		//var info []byte
		info := fmt.Sprintf("%04d-%02d-%02d-%02d:%02d", year, month, day, hour, minute)
		//fmt.sprintf(info, "%04d-%02d-%02d-%02d:%02d", year, month, day, hour, minute)
		print(info)
		print("\t")
		year = ed.Year
		month = ed.Month
		day = ed.Day
		hour = ed.Hour
		minute = ed.Minute

		info = fmt.Sprintf("%04d-%02d-%02d-%02d:%02d", year, month, day, hour, minute)
		print(info)
		print("\t")
		for _, p := range meeting.Participators {
			print(p)
			for j := 4; j <= standardNameLength; j += 4 {
				if len(p) < j {
					for k := j - 4; k < standardNameLength; k += 4 {
						print("\t")
					}
				}
			}
		}

		println()
	}
	println("-----------------------------------------------------------")
}
