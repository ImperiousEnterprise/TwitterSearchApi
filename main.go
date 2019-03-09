package main
import (
	"github.com/dghubble/go-twitter/twitter"
	 "github.com/dghubble/oauth1"
)

var config = oauth1.NewConfig("HDfyHzFk7qblM5MRXzJ9fS9QE", "flwRR7UN5Xn6tvrEmUCbMWzxTzGvofLavPx1B0cPNTzgurmA1O")
var token = oauth1.NewToken("2373250494-kJ4kb0DYqB8Xpzc65hzUgFGIEFPkfeERP5j6jtT", "RZIgffDBzOQKQtqyCAhwWWYkaAiDCRtU5OvdQsiBQ4QW1")
var httpClient = config.Client(oauth1.NoContext, token)

// Twitter client
var client = twitter.NewClient(httpClient)

func main() {
	a := App{}
	a.initialize()
	a.Run(":8080")

}

