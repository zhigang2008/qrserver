// mongotest project main.go
package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Position struct {
	Longitude int
	Latitude  int
}

type LData struct {
	Sid   string
	Name  string
	Value int
	Pos   Position
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&LData{"001", "设备1号", 200, Position{120, 30}},
		&LData{"002", "设备2号", 134, Position{20, 65}})
	if err != nil {
		panic(err)
	}

	result := LData{}
	err = c.Find(bson.M{"name": "设备1号"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("数值:", result.Value)
}
