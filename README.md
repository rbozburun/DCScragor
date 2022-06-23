
# DC Scragor

RSS feed scraper for provided URLS.

## Running from the Server

```javascript
go run .
```



## How it works?

```javascript
Available commands:
----------------------------
!dcscragor help                 Prints this information 
!dcscragor add_blog <URL>       Adds your RSS feed URL to the scrape wish list.
```

### Details
This bot scrapes provided RSS feeds as hardcoded 3 times in a day. If a new post has been published, it sends the title and URL of the post to the specifed channel with **BOT-RSS-SHARE-CHANNEL-ID**. 

You can change **BOT-TOKEN** variable and **BOT-RSS-SHARE-CHANNEL-ID** then use it.

