package main

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
)

type user struct {
	name     string `bson:"name"`
	lastname string `bson:"lastname"`
	age      int    `bson:"age"`
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Panicln(err)
	}
	defer session.Close()

	result := make([]user, 2)

	if err := session.DB("mydb").C("trans").Find(nil).All(&result); err != nil {
		log.Panicln(err)
	}
	for _, res := range result {
		fmt.Println(res.age)
		fmt.Println(res.name)
		fmt.Println(res.lastname)
	}

}
