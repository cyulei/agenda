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

	"github.com/cyulei/agenda/entity"
	"github.com/spf13/cobra"
)

var (
	meeting_title     string
	add_flag          bool
	delete_flag       bool
	participator_name string
)

// changeparticipatorCmd represents the changeparticipator command
var changeparticipatorCmd = &cobra.Command{
	Use:   "changeparticipator",
	Short: "Current user can change participators of a meeting",
	Long: `Current user can change participators of a meeting he sponsors. The adding process\n
		need date checks, that is to say participators need to have free time for this meeting.\n
		If a meeting has no participators after this cmd, this meeting will be deleted.For exanple:\n
		changeparticipator xxx(meeting-title) -d xxx|xxx|xxx`,
	Run: func(cmd *cobra.Command, args []string) {
		//get current user
		var current_user entity.User
		current_user = get_current_user()
		change_participators := strings.Split(participator_name, "|")
		if delete_flag {
			finished := false
			meetings := get_all_meetings()
			for i, j := range meetings {
				if j.Sponsor == current_user.Name && j.Title == meeting_title {
					for m, k := range j.Participators {
						if k.Name == participator_name {
							j.Participators = append(j.Participators[:m], j.Participators[m+1:]...)
							finished = true
							if len(j.Participators) == 0 {

							}
							break
						}
					}
					if !finished {
						break
					} else {

					}
				}
			}
			fmt.Println("changeparticipator called")
		}
	},
}

func init() {
	rootCmd.AddCommand(changeparticipatorCmd)
	changeparticipatorCmd.Flags().StringVarP(&meeting_title, "title", "mt", "", "meeting title")
	changeparticipatorCmd.MarkFlagRequired("title")
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
