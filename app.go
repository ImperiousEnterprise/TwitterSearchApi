package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/labstack/gommon/log"
	"strconv"
	"encoding/json"
	"strings"
	"github.com/dghubble/go-twitter/twitter"
	"flag"
)

type App struct{
	Router *mux.Router
}

func (a *App) initialize() {
	a.Router = mux.NewRouter()
	a.activateRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) activateRoutes() {
	 var entry string
	flag.StringVar(&entry, "entry", "./index.html", "the entrypoint to serve.")

	//Setting API endpoints for VueJs to communicate with
	a.Router.HandleFunc("/search", SearchHandler).Methods("GET")
	//Restricting the search API endpoint to only allow GET Requests and all other request return 403
	a.Router.HandleFunc("/search", status(403, "GET")).Methods("POST", "PATCH", "DELETE")

	// Serve static assets directly.
	a.Router.PathPrefix("/scripts/vue-resource.js").HandlerFunc(IndexHandler("./scripts/vue-resource.js")).Methods("GET")

	// Set entry point to serve
	a.Router.PathPrefix("/").HandlerFunc(IndexHandler(entry))
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}

func SearchHandler(w http.ResponseWriter, r *http.Request){
	text := r.FormValue("q")
	count, _ := strconv.Atoi(r.FormValue("count"))

	//In the chance that q parameter is not included in url
	// dont carry out api call to twitter
	if text == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal("Sorry, you entered an invalid query. Please try your search again. ")
		w.Write(j)
	}

	// Search Tweets
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: text,
		Count: count ,
		TweetMode: "extended",
	})
	log.Info("Response Status from Search: " + resp.Status)

	if(err != nil){
		log.Error(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		j, _ := json.Marshal("Sorry, you entered an invalid query. Please try your search again. ")
		w.Write(j)
	}else{
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Status", resp.Status)
		j, _ := json.Marshal(search.Statuses)
		w.Write(j)
	}
}


// status is used to set a specific status code
func status(code int, allow ...string) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(code)
		if len(allow) > 0 {
			w.Write([]byte(`Allow: ` + strings.Join(allow, ", ")))
		}
	}
}