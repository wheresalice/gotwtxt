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
	"fmt"
	"github.com/gregjones/httpcache/diskcache"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/bcampbell/fuzzytime"
	"github.com/fatih/color"
	"github.com/gregjones/httpcache"
	"github.com/spf13/cobra"
)

type tweet struct {
	username  string
	timestamp fuzzytime.DateTime
	text      string
}

type tweetSlice []tweet

func (s tweetSlice) Less(i, j int) bool {
	return s[i].timestamp.ISOFormat() < s[j].timestamp.ISOFormat()
}
func (s tweetSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s tweetSlice) Len() int      { return len(s) }

// timelineCmd represents the timeline command
var timelineCmd = &cobra.Command{
	Use:   "timeline",
	Short: "Print the timeline of people you are following",
	Long: `Fetch all of the posts of the accounts you are following and print them in chronological order.  Will use caching if the server sends the right headers`,
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
		cache := diskcache.New(path.Join(home, ".twtxt_cache"))
		transport := httpcache.NewTransport(cache)
		client := &http.Client{Transport: transport}

		followingFile := path.Join(home, ".twtxt_following.ini")
		following, err := ini.Load(followingFile)
		if err != nil {
			log.Fatalln("Couldn't parse following file " + followingFile)
		}

		var tweets tweetSlice
		for username, url := range following.Section("").KeysHash() {
			resp, err := client.Get(url)
			if err != nil {
				log.Printf("Failed to fetch %s: %v\n", url, err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Failed to read from %s: %v\n", url, err)
			}
			sb := string(body)
			lines := strings.Split(sb, "\n")
			for i, line := range lines {
				// skip empty lines and comments
				if len(line) == 0 || string(line[0]) == "#" {
					continue
				}
					r := regexp.MustCompile(` {4,}|\t`) // whilst this is supposed to be a tab, it is sometimes multiple spaces
				lineFields := r.Split(line, 2)
				if len(lineFields) != 2 {
					log.Printf("Failed to parse line %v of %s", i+1, url)
				} else {
					dt, _, err := fuzzytime.Extract(lineFields[0])
					if err != nil {
						log.Println(err)
					}
					tweets = append(tweets, tweet{
						username:  username,
						timestamp: dt,
						text:      lineFields[1],
					})
				}
				//println(username + line)
			}
		}
		sort.Sort(tweets)
		for _, t := range tweets {
			blue := color.New(color.FgBlue).SprintFunc()
			white := color.New(color.FgHiWhite).SprintFunc()
			fmt.Println(blue(t.timestamp.ISOFormat()) + " <" + t.username + "> " + white(t.text))
		}
	},
}

func init() {
	rootCmd.AddCommand(timelineCmd)
}
