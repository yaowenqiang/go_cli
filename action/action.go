package main
import (
	"flag"
	"log"
	"plugin"
)
func main () {
	path := flag.String("plugin", "", "Plugin to execute")

	flag.Parse()

	if *path == "" {
		log.Fatal("Path to plugin must be provided")
	}
	
	p , err := plugin.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	f, err := p.Lookup("ThingToDo")
	if err != nil {
		log.Fatal(err)
	}
	
	thingsToDo, ok := f.(func())

	if !ok {
		log.Fatal("Could not find function 'ThingToDo' in plugin")
	}
	thingsToDo()

	log.Println("Did the thing")
	

}
// go build --buildmode=plugin -o=plugin.so sousrc/plugin/plugin.go
