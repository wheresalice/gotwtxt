/*
Copyright © 2021 Wheresalice

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// followingCmd represents the following command
var followingCmd = &cobra.Command{
	Use:   "following",
	Short: "List who you are following",
	Long:  `List all of the users that you are following`,
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
		followingFile := path.Join(home, ".twtxt_following.ini")
		following, err := ini.Load(followingFile)
		if err != nil {
			log.Fatalln("Couldn't parse following file " + followingFile)
		}
		b, _ := json.MarshalIndent(following.Section("").KeysHash(), "", " ")
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(followingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// followingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// followingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
