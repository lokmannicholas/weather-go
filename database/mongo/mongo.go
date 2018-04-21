package mongo

import (
	"github.com/lokmannicholas/weather-go/config"
	"gopkg.in/mgo.v2"
)

type MongoDB struct {
	Session *mgo.Session
}

var (
	mdb *MongoDB
	s   *mgo.Session
)

func GetMongoDB() *MongoDB {
	if mdb == nil {
		session, err := mgo.Dial(config.Get().MongoDB.Host)
		if err != nil {
			panic(err)
		}
		session.SetMode(mgo.Monotonic, true)
		s = session
		mdb = &MongoDB{
			Session: s,
		}
	}

	return mdb
}

func (mdb *MongoDB) GetCollection(collec string) *mgo.Collection {
	return mdb.Session.DB(config.Get().MongoDB.DBName).C(collec)
}
