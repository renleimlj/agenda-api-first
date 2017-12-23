// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
  "agenda-api-first/cli/logger"
  "io/ioutil"
  "strings"
  "agenda-api-first/cli/entity"
)

// addPrCmd represents the addPr command
var addPrCmd = &cobra.Command{
  Use:   "addPr",
  Short: "--title=meeting --pr=participator",
  Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  Run: func(cmd *cobra.Command, args []string) {
    title, _ := cmd.Flags().GetString("title") 
    prs, _ := cmd.Flags().GetString("pr") 
    prs_arr := strings.Split(prs, ",") 
    meeting := entity.QueryMeetingByTitle(title)
    for _, each := range prs_arr { 
      meeting.AddParticipator(each)
      bytes,_ := ioutil.ReadFile("./CurUser")
      curuser := string(bytes)
      logger.Log("'" + curuser + "' called: addPr, title: " + title + ", participators: " + prs)
    }
  },
}

func init() {
  RootCmd.AddCommand(addPrCmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // addPrCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // addPrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
  addPrCmd.Flags().StringP("title", "t", "", "The title of the meeting the pr wanting to add in.")
  addPrCmd.Flags().StringP("pr", "p", "", "The pr wanting to add in.Please split by comma .") 
}
