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
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "User log out",
	Long:  `User log out, input command mode like : logout`,
	Run: func(cmd *cobra.Command, args []string) {
		//log
		fileName := "datarw/Agenda.log"
		logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
		defer logFile.Close()
		if err != nil {
			log.Fatalln("Open file error")
		}
		infoLog := log.New(logFile, "[Info]", log.Ldate|log.Ltime|log.Lshortfile)
		infoLog.Println("Cmd logout called")

		//确定当前是登陆状态
		curUser := datarw.GetCurUser()
		if curUser == nil {
			infoLog.SetPrefix("[Error]")
			infoLog.Println("User is not logged in")
			fmt.Println("Please log in first!")
			return
		}
		datarw.SaveCurUser(nil)
		fmt.Println("Logout success!")
		//登出
		infoLog.SetPrefix("[Error]")
		infoLog.Println(curUser.Name + " logout success")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
