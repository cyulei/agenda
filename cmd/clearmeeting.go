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
	"github.com/spf13/cobra"
)

var info_show bool

// clearmeetingCmd represents the clearmeeting command
var clearmeetingCmd = &cobra.Command{
	Use:   "clearmeeting",
	Short: "Current user can clear meetings which he sponsors",
	Long: `Current user can clear all meetings that he sponsors, for example:\n
		clearmeeting -i clear all meetings and print titles of meeting being deleted`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			var delete_meetings []string
			//get current user
			var current_user entity.User
			current_user = get_current_user()
			//get all existed meetings
			meetings := get_all_meetings()
			//meetings after delete
			var final_meetings []entity.Meeting
			for i, j := range meetings {
				if j.Sponsor == current_user.Name {
					delete_meetings = append(delete_meetings, j.Title)
				} else {
					final_meetings = append(j)
				}
			}

			if info_show {
				for i, j := range delete_meetings {
					fmt.Println("deletemeeting" + strconv.Itoa(i) + ": " + j)
				}
			}
			fmt.Println("clearmeeting finished")
		*/
	},
}

func init() {

	rootCmd.AddCommand(clearmeetingCmd)
	clearmeetingCmd.Flags().BoolVarP(&info_show, "info", "i", false, "show meetings cleared")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
