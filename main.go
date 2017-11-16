package main
import (
	"github.com/dghubble/go-twitter/twitter"
	 "github.com/dghubble/oauth1"
)

var config = oauth1.NewConfig("", "")
var token = oauth1.NewToken("", "")
var httpClient = config.Client(oauth1.NoContext, token)

// Twitter client
var client = twitter.NewClient(httpClient)

func main() {
	a := App{}
	a.initialize()
	a.Run(":8080")

}

