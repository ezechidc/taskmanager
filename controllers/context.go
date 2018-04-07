package controllers

import (
	"github.com/ezechidc/taskmanager/common"
	mgo "gopkg.in/mgo.v2"
)

//Context Struct used for maintaining HTTP Request Context
type Context struct {
	MongoSession *mgo.Session
	User         string
}

//DbCollection Returns mgo.collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

// Close mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

// NewContext creates a new Context object for each HTTP request
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}
