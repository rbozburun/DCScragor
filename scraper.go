package main

import (
	"io/ioutil"
	"os"
	"strings"
	"fmt"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

type FeedItem struct {
	Title string
	URL   string
}

// Return the content of given file
func readFile(fname string) string {
	databyte, err := ioutil.ReadFile(fname)
	checkErr(err)
	return string(databyte)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}

func ParseRSS() {
	blogList := [2]string{"https://sibersaldirilar.com/feed/", "https://pwnlab.me/feed/"}
	fp := gofeed.NewParser()
	fp.Client = &http.Client{ Timeout: time.Second*10 } 

	feed_items := make([]FeedItem, 1)

	for true {
		for k := 0; k < len(blogList); k++ {
			feed, err := fp.ParseURL(blogList[k])
			
			//if err nil, everything is OK
			if err == nil {
				l.Printf("[INFO] RSS Parsing started for %s", blogList[k])
			items := feed.Items
			for i := 0; i < len(items); i++ {
				// Create FeedItemObj for each link Item in RSS
				if items[i].Title != "" && items[i].Link != "" {
					if strings.Contains(readFile("feed_item.list"), items[i].Link) {
						//l.Printf("[WARN] FeedItem already added to the feed_item.list. Title: %s Link: %s", items[i].Title, items[i].Link)
					} else {
						feedItem := FeedItem{Title: items[i].Title, URL: items[i].Link}
						feed_items = append(feed_items, feedItem)
						l.Printf("[INFO] FeedItem is created. Title: %s Link: %s", items[i].Title, items[i].Link)

						// Write link to the file
						file, err := os.OpenFile("feed_item.list", os.O_APPEND|os.O_WRONLY, 0644)
						if err != nil {
							l.Println(err)
						}
						defer file.Close()
						if _, err := file.WriteString(items[i].Link + "\n"); err != nil {
							l.Fatal(err)
						}

						// Send title and link to the DC channel
						msg := "New blog post published: **" + items[i].Title + "**\n" + items[i].Link
						Dg.ChannelMessageSend(botChID, msg)

					}
				} else {
					l.Printf("[WARN] FeedItem couldn't create since link or title is empty!. URL: %s", blogList[k])
				}
			}
			} else{
				fmt.Println(err)
			}

			
		}

		// Regenrate the items list
		feed_items = make([]FeedItem, 1)

		// Wait 8 hours
		time.Sleep(28800 * time.Second)
	}
}
