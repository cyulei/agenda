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

//var cfgFile string
var password string

// deleteuserCmd represents the deleteuser command
var deleteuserCmd = &cobra.Command{
	Use:   "deleteuser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteuser called")
		deleteuser(password)
	},
}

func init() {
	rootCmd.AddCommand(deleteuserCmd)

	registerCmd.Flags().StringVarP(&password, "password", "p", "", "user's password")
}
func deleteuser(password string) {
	curUser := datarw.GetCurUser()
	if curUser == nil { //GetCurUser可以返回一个指针
		fmt.Println("isn't login,please use command login")
		return
	}
	if password == "" {
		fmt.Println("please input password")
		return
	}
	if password == curUser.Password {
		users := datarw.GetUsers()
		for index, user := range users {
			if user.Name == curUser.Name {
				users = append(users[:index], users[index+1:]...)
				datarw.SaveUsers(users)
				fmt.Println("User:", curUser.Name, " has been deleted")
				return
			}
		}
		fmt.Println("error: unexpected to execute")

	} else {
		fmt.Println("password incorrect")
	}
}
