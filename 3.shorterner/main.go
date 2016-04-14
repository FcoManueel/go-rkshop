//IMPORTANT

package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

var URLsRepository = []URLPair{
	URLPair{
		ID:      "g",
		LongUrl: "http://www.google.com",
	},
	URLPair{
		ID:      "m",
		LongUrl: "http://www.github.com/FcoManueel",
	},
}

type URLPair struct {
	ID      string `json:"id"`
	LongUrl string `json:"longUrl"`
}

func getID(url string) string {
	i := strings.LastIndex(url, "/")
	return url[i+1:]
}

func expand(shortURL string) string {
	for _, pair := range URLsRepository {
		if pair.ID == getID(shortURL) {
			return pair.LongUrl
		}
	}
	return ""
}

func Redirect(w http.ResponseWriter, req *http.Request) {
	longURL := expand(req.RequestURI)
	if longURL == "" {
		log.Println("No se encontro nada. ", "Req: ", req.RequestURI)
		return
	}
	http.Redirect(w, req, longURL, http.StatusSeeOther)
}

func main() {
	log.Println("[go-short] starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "root!")
	})
	http.HandleFunc("/r/", Redirect)
	err := http.ListenAndServe(":8080", nil)

	log.Fatal(err)
}
