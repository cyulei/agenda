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
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login -u [username] -p [password]",
	Short: "User log in",
	Long:  `User log in, input command mode like : login -u username -p password`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		//fmt.Println(username, password)
		curUser := datarw.GetCurUser()
		if curUser != nil {
			fmt.Println("Already logged in!Please log out first!")
			return
		}
		//检测是否已经存在用户
		users := datarw.GetUsers()
		for i := 0; i < len(users); i++ {
			if users[i].Name == username {
				//得到用户密码是否正确
				if users[i].Password == password {
					//标示用户已经登陆
					fmt.Println("Login success!")
					datarw.SaveCurUser(&users[i])
					return
				}
				fmt.Println("Password erorr!")
				return
			}
		}
		fmt.Println("Username don't exist!Please register.")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("username", "u", "", "user name")
	loginCmd.Flags().StringP("password", "p", "", "user password")
}
