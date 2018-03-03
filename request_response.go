package main

func main() {
	handleResponses()
}

var driverName = "mysql"
var dataSourceName = "test:test@/reqres"

func check(e error) {
	if e != nil {
		panic(e)
	}
}