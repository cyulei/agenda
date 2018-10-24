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

	"github.com/cyulei/agenda/datarw"

	"os"

	"github.com/cyulei/agenda/entity"
	"github.com/spf13/cobra"
)

// cancelmeetingCmd represents the cancelmeeting command
var cancelmeetingCmd = &cobra.Command{
	Use:   "cancelmeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	//load
	usr := datarw.GetCurUser()
	meetings := datarw.GetMeetings()
	res := make([]entity.Meeting, 1) //for writing

	//adjust the parameters
	if cancel_title == "" {
		fmt.Fprint(os.Stderr, "you must use a -t title to tell me which meeting you want to cancel")
		return
	}
	//check login state
	//if not log in exit

	//cancel
	var meetingExist = false
	var authoritySatisified = false

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
		return
	}
	if !authoritySatisified {
		println("there is a meeting called :", cancel_title, "but you are not the sponsor of it,please try to relog in ?")
		return
	}
	datarw.SaveMeetings(res)
	println("successfully cancel the meeting :", cancel_title)
}
