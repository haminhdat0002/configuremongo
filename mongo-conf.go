package configuremongo

import (
	"errors"
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// NewJSONFromFile returns an instance of the JSON checker. It reads its
// data from a file which its location has been specified through the path
// parameter
func NewMongo(mongoDbConfString string) *MONGO {
	return &MONGO{
		confString: mongoDbConfString,
	}
}

// JSON represents the JSON Checker. It reads an io.Reader and then pulls a value out of a map[string]interface{}.
type MONGO struct {
	values     map[string]interface{}
	confString string
}

type MongoConf struct {
	Name  string
	Value interface{}
}

//Setup initializes the JSON Checker
func (m *MONGO) Setup() error {
	session, err := mgo.Dial(m.confString)
	if err != nil {
		return err
	}
	confCollection := session.DB("").C("confs")
	results := []MongoConf{}
	// results := []bson.D{}
	err = confCollection.Find(nil).All(&results)
	if err != nil {
		fmt.Println("could not retrieve configurations")
	} else {
		fmt.Println(results)
	}
	m.values = make(map[string]interface{})

	for _, conf := range results {
		m.values[conf.Name] = conf.Value
	}
	fmt.Println(m.values)
	return nil
}

func (m *MONGO) value(name string) (interface{}, error) {
	val, ok := m.values[name]
	if !ok {
		return nil, errors.New("that variable does not exist")
	}

	return val, nil
}

// Int returns an int if it exists within the marshalled JSON io.Reader.
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

// Bool returns a bool if it exists within the marshalled JSON io.Reader.
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

// String returns a string if it exists within the marshalled JSON io.Reader.
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
