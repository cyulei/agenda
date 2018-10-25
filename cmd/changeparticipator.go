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
	"strings"

	"github.com/cyulei/agenda/datarw"
	"github.com/cyulei/agenda/entity"
	"github.com/spf13/cobra"
)

var (
	change_meeting_title string
	add_flag             bool
	delete_flag          bool
	participator_name    string
)

// changeparticipatorCmd represents the changeparticipator command
var changeparticipatorCmd = &cobra.Command{
	Use:   "changeparticipator",
	Short: "Current user can change participators of a meeting",
	Long: `Current user can change participators of a meeting he sponsors. The adding process\n
		need date checks, that is to say participators need to have free time for this meeting.\n
		If a meeting has no participators after this cmd, this meeting will be deleted. For exanple:\n
		changeparticipator xxx(meeting-title) -d/-a xxx-xxx-xxx`,
	Run: func(cmd *cobra.Command, args []string) {
		current_user := datarw.GetCurUser() //current user
		if current_user == nil {
			fmt.Println("Please log first")
			return
		}

		change_participators := strings.Split(participator_name, "-") //change name list
		meetings := datarw.GetMeetings()                              //all meetings
		meeting_exist := false
		var final_participators []string
		if delete_flag {
			//delete participators from a meeting

			var delete_participators []string
			for i, j := range meetings {
				if j.Sponsor == current_user.Name && j.Title == change_meeting_title {
					meeting_exist = true

					for _, k := range j.Participators {
						if isParticipatorinList(k, change_participators) {
							delete_participators = append(delete_participators, k)
						} else {
							final_participators = append(final_participators, k)
						}
					}

					if len(final_participators) == 0 {
						meetings = append(meetings[:i], meetings[i+1:]...)
					} else {
						j.Participators = final_participators
					}
					datarw.SaveMeetings(meetings)
					break
				}
			}
			if !meeting_exist {
				fmt.Println("No such meeting, check meeting title")
			} else if len(delete_participators) != len(change_participators) {
				fmt.Println("Some users don't exist in this meeting. Already delete: ")
				for _, j := range delete_participators {
					fmt.Println(j)
				}
			}
		} else {
			//add participators to a meeting
			var valid_participators []string
			var all_users []entity.User
			all_users = datarw.GetUsers()
			for _, j := range change_participators {
				if !isParticipatorExist(j, all_users) {
					fmt.Println(j + " is not a valid user")
				} else {
					valid_participators = append(valid_participators, j)
				}
			}
			if len(valid_participators) != 0 {
				for i, j := range meetings {
					if j.Sponsor == current_user.Name && j.Title == change_meeting_title {
						final_participators = j.Participators
						meeting_exist = true
						for _, k := range valid_participators {
							if isParticipatorExistinMeeting(k, j) {
								fmt.Println(k + " is already in this meeting")
							} else {
								if isParticipatorAvailable(k, meetings, j) {
									meetings[i].Participators = append(meetings[i].Participators, k)
								} else {
									fmt.Println(k + " is not free")
								}
							}
						}
						datarw.SaveMeetings(meetings)
						break
					}
				}
				if !meeting_exist {
					fmt.Println("No such meeting, check meeting title")
				}
			}
		}
		fmt.Println("changeparticipator finished")
	},
}

func init() {
	rootCmd.AddCommand(changeparticipatorCmd)
	changeparticipatorCmd.Flags().StringVarP(&change_meeting_title, "chptitle", "t", "", "meeting title")
	changeparticipatorCmd.MarkFlagRequired("chptitle")
	changeparticipatorCmd.Flags().BoolVarP(&add_flag, "add", "a", true, "add participator")
	changeparticipatorCmd.Flags().BoolVarP(&delete_flag, "delete", "d", false, "delete participator")
	changeparticipatorCmd.Flags().StringVarP(&participator_name, "name", "n", "", "participator's name")
	changeparticipatorCmd.MarkFlagRequired("name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changeparticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// changeparticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func isParticipatorinList(name string, participators []string) bool {
	for _, j := range participators {
		if name == j {
			return true
		}
	}
	return false
}

func isParticipatorExist(name string, participators []entity.User) bool {
	for _, j := range participators {
		if name == j.Name {
			return true
		}
	}
	return false
}

func isParticipatorExistinMeeting(name string, meeting entity.Meeting) bool {
	if name == meeting.Sponsor {
		return true
	}
	for _, j := range meeting.Participators {
		if name == j {
			return true
		}
	}
	return false
}

func isParticipatorAvailable(name string, all_meetings []entity.Meeting, current_meeting entity.Meeting) bool {
	start_date := current_meeting.Startdate
	end_date := current_meeting.Enddate
	for _, j := range all_meetings {
		if isParticipatorExistinMeeting(name, j) {
			if entity.Compare(j.Startdate, end_date) >= 0 || entity.Compare(start_date, j.Enddate) >= 0 {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
