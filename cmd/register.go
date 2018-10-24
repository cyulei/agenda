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
	"regexp"

	"github.com/cyulei/agenda/datarw"
	"github.com/cyulei/agenda/entity"

	"github.com/spf13/cobra"
)

//var cfgFile string
var name, registerPassword, email, phone string

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		register(name, registerPassword, email, phone)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	registerCmd.Flags().StringVarP(&name, "name", "n", "", "user's name")
	registerCmd.Flags().StringVarP(&registerPassword, "password", "p", "", "user's password")
	registerCmd.Flags().StringVarP(&email, "email", "e", "", "user's email")
	registerCmd.Flags().StringVarP(&phone, "phone", "t", "", "user's phone")
}

func register(name string, password string, email string, phone string) {
	if isValidName(name) && isValidPassword(password) && isValidEmail(email) && isValidPhone(phone) {
		users := datarw.GetUsers()
		if !hasName(name, users) {
			newuser := entity.User{Name: name, Password: password, Email: email, Phone: phone}
			users = append(users, newuser)
			datarw.SaveUsers(users)
			fmt.Println("Registration complete")
		}

	}

}

//Judge username exists
func hasName(name string, users []entity.User) bool {
	for _, user := range users {
		if user.Name == name {
			fmt.Println("The Username has been registered")
			return false
		}
	}
	return true
}
func isValidName(n string) bool {
	b := []byte(n)

	val, _ := regexp.Match(".+", b)

	return val
}
func isValidPassword(p string) bool {
	b := []byte(p)

	val, _ := regexp.Match(".+", b)
	if len(p) < 9 {
		fmt.Println("The password must be longer than 8 digits")
		val = false
	}

	return val
}
func isValidEmail(e string) bool {
	b := []byte(e)

	val, _ := regexp.Match("\\w*@\\w*\\.w*", b)

	if !val {
		fmt.Println("The Email is invaild")
	}
	return val
}
func isValidPhone(p string) bool {
	b := []byte(p)

	val, _ := regexp.Match("[0-9]+", b)

	return val
}
