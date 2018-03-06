This is a configuration checker for the nice go configuration tool : https://github.com/paked/configure

It will enable you to use a mongo db collection as a source of your settings.
This project require the MGO mongo driver : https://github.com/go-mgo/mgo/tree/v2
Make sure they are available in your project (if you want to run this project separatly or run the test, just run 
```
dep ensure
```
and it will made available.


Create a MONGO checker and use it :

```go
	conf.Use(configure-mongo.NewMongo())
```
