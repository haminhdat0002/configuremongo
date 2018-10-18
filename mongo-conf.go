package configuremongo

import (
	"errors"
	"fmt"

	mgo "github.com/globalsign/mgo"
)

// NewMongo returns an instance of the Mongo checker. It reads its
// data from a collection confCollectionName inside the database that could be reach with the mongoDbConfString
// For more informations about this configuration string, please check the MGO driver : https://godoc.org/labix.org/v2/mgo#Dial
func NewMongo(mongoDbConfString string, confCollectionName string) *MONGO {
	return &MONGO{
		confString:         mongoDbConfString,
		confCollectionName: confCollectionName,
	}
}

// MONGO represents the MONGO Checker. It reads form the mongo collection and then pulls a value out of a map[string]interface{}.
type MONGO struct {
	values             map[string]interface{}
	confString         string
	confCollectionName string
}

type MongoConf struct {
	Name  string
	Value interface{}
}

//Setup initializes the Mongo Checker
func (m *MONGO) Setup() error {
	session, err := mgo.Dial(m.confString)
	if err != nil {
		return err
	}
	confCollection := session.DB("").C(m.confCollectionName)
	results := []MongoConf{}
	m.values = make(map[string]interface{})
	// results := []bson.D{}
	err = confCollection.Find(nil).All(&results)
	if err != nil {
		fmt.Println("could not retrieve configurations from mongo")
		return err
	} else {
		for _, conf := range results {
			m.values[conf.Name] = conf.Value
		}
		// fmt.Println(m.values)
		return nil
	}
}

func (m *MONGO) value(name string) (interface{}, error) {
	val, ok := m.values[name]
	if !ok {
		return nil, errors.New("that variable does not exist")
	}

	return val, nil
}

// Int returns an int if it exists in the mongo collection.
func (m *MONGO) Int(name string) (int, error) {
	v, err := m.value(name)
	if err != nil {
		return 0, err
	}

	f, ok := v.(float64)
	if !ok {
		return 0, errors.New(fmt.Sprintf("%T unable", v))
	}

	return int(f), nil
}

// Bool returns a bool if it exists in the mongo collection.
func (m *MONGO) Bool(name string) (bool, error) {
	v, err := m.value(name)
	if err != nil {
		return false, err
	}

	b, ok := v.(bool)
	if !ok {
		return false, errors.New("unable to cast")
	}

	return b, nil
}

// String returns a string if it exists in the mongo collection.
func (m *MONGO) String(name string) (string, error) {
	v, err := m.value(name)
	if err != nil {
		return "", err
	}

	s, ok := v.(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("unable to cast %T", v))
	}

	return s, nil
}
