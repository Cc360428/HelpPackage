// Package mongodb
package mongodb

import (
	"gopkg.in/mgo.v2"
)

func NewMongodb(address string) (session *mgo.Session, err error) {
	return mgo.Dial(address)
}
