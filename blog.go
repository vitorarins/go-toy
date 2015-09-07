package main

import (
	"fmt"
	"os"

	"github.com/vitorarins/go-blog/Godeps/_workspace/src/github.com/go-martini/martini"
	"github.com/vitorarins/go-blog/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"github.com/vitorarins/go-blog/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
)

type document struct {
	Id      bson.ObjectId `bson:"_id"`
	Title   string        `bson:"title"`
	Content string        `bson:"content"`
}

func main() {
	uri := os.Getenv("MONGO_URL")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	martiniClass := martini.Classic()
	martiniClass.Get("/", func() string {
		return "Hello world!"
	})

	martiniClass.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	martiniClass.Get("/document", func() string {

		sess, err := mgo.Dial(uri)
		if err != nil {
			fmt.Printf("Can't connect to mongo, go error %v\n", err)
			os.Exit(1)
		}
		defer sess.Close()

		var documentFound document
		err = sess.DB("test").C("documents").Find(bson.M{}).One(&documentFound)
		if err != nil {
			fmt.Printf("got an error finding a doc %v\n")
			os.Exit(1)
		}

		return fmt.Sprintf("Found document: %+v\n", documentFound)
	})

	martiniClass.Run()
}
