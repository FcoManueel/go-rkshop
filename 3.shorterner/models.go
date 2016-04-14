package main

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
