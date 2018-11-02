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
	"strconv"

	"github.com/cyulei/agenda/datarw"
	"github.com/cyulei/agenda/entity"
	"github.com/spf13/cobra"
)

var info_show bool

// clearmeetingCmd represents the clearmeeting command
var clearmeetingCmd = &cobra.Command{
	Use:   "clearmeeting",
	Short: "Current user can clear meetings which he sponsors",
	Long: `clearmeeting:Current user can clear all meetings that he sponsors and see the details of the cleanup.
		For example:
		agenda clearmeeting -i 
		clear all meetings and print titles of meeting being deleted`,
	Run: func(cmd *cobra.Command, args []string) {
		//log
		fileName := "datarw/Agenda.log"
		logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
		defer logFile.Close()
		if err != nil {
			log.Fatalln("Open file error")
		}
		infoLog := log.New(logFile, "[Info]", log.Ldate|log.Ltime|log.Lshortfile)

		var delete_meetings []string
		//get current user
		current_user := datarw.GetCurUser()
		if current_user == nil {
			infoLog.SetPrefix("[Error]")
			infoLog.Println("Not log in yet")
			fmt.Println("Please log first")
			infoLog.Println("Cmd clearmeeting failed")
			fmt.Println("Cmd clearmeeting failed")
			return
		}
		infoLog.Println("Current User: " + current_user.Name + ", Cmd clearmeeting called")
		//get all existed meetings
		meetings := datarw.GetMeetings()
		//meetings after delete
		var final_meetings []entity.Meeting
		for _, j := range meetings {
			if j.Sponsor == current_user.Name {
				delete_meetings = append(delete_meetings, j.Title)
			} else {
				final_meetings = append(final_meetings, j)
			}
		}
		datarw.SaveMeetings(final_meetings)
		if info_show {
			for i, j := range delete_meetings {
				fmt.Println("deletemeeting" + strconv.Itoa(i) + ": " + j)
			}
		}
		infoLog.SetPrefix("[Info]")
		infoLog.Println("Current User: " + current_user.Name + ", Cmd clearmeeting finished")
		fmt.Println("clearmeeting finished")

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
