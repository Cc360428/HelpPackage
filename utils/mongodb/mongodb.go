// mongodb
package mongodb

import (
	"gopkg.in/mgo.v2"
)

func NewMongdodb(address string) (session *mgo.Session, err error) {
	session, err = mgo.Dial(address)
	return session, nil
}
