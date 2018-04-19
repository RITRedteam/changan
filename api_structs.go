package main

import "gopkg.in/mgo.v2/bson"

type APIReport struct {
	DeviceID bson.ObjectId `json:"device_id"`
	Report   string        `json:"report"`
	Title    string        `json:"title"`
	IP       string        `json:"ip"`
}
