This is a configuration checker for the nice go configuration tool : https://github.com/paked/configure

It will enable you to use a mongo db collection as a source of your settings.
This project require the MGO mongo driver : https://github.com/go-mgo/mgo/tree/v2
Make sure they are available in your project (if you want to run this project separatly or run the test, just run
```
dep ensure
```
and it will made available.


Create a MONGO checker and use it (checking in the "confs" collection from configuration values) :

```go
conf.Use(configuremongo.NewMongo("mongodb://127.0.0.1:27017/test_conf_db"))
```

The mongo db collection must contain documents with a field named "name" and a field named "value" containing its typed value.
