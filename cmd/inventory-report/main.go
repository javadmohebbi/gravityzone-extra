package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/javadmohebbi/gravityzone-extra/gzmongo"
	"github.com/javadmohebbi/gravityzone-extra/inventory"
)

/*
Author: M. Javad Mohebbi
This example will guide you how to get list of
installed applications on gravityzone
MongoDB address & password needed for getting the information


Just fork this https://github.com/javadmohebbi/gravityzone-extra repo & add more functionalities
*/

func main() {
	mongoAddress := flag.String("mongo-address", "localhost:27017", "Mongo db address & port")
	mongoPass := flag.String("mongo-pass", "", "Mongo db password")

	outPath := flag.String("out", "/home/bdadmin", "Path for output")

	flag.Parse()

	if *mongoAddress == "" {
		log.Fatalln("-mongo-address argument must be provided")
	}

	if *mongoPass == "" {
		log.Fatalln("-mongo-pass argument must be provided")
	}

	// create new instance of mongo db
	// to fetch info from DB
	gzmongo := gzmongo.New(*mongoAddress, *mongoPass)

	// disconnect this app from mongo db
	defer gzmongo.Disconnect()

	iv := inventory.New(gzmongo.Client, context.Background())

	lst, err := iv.GetAll()

	if err != nil {
		panic(err)
	}

	var linesWithEndpoint, linesApp []string
	for _, l := range lst {
		for _, e := range l.Endpoints {
			ll := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v",
				e.EndpointName, e.EndpointOS, l.Name,
				l.PathString,
				l.Specifics.Hash, l.Specifics.Version, l.Specifics.DiscoveredOn,
			)
			linesWithEndpoint = append(linesWithEndpoint, ll)
		}

		ll := fmt.Sprintf("%v,%v,%v,%v,%v,%d",
			l.Name,
			l.PathString,
			l.Specifics.Hash, l.Specifics.Version, l.Specifics.DiscoveredOn,
			len(l.Endpoints),
		)
		linesApp = append(linesApp, ll)

	}

	// for _, l := range linesApp {
	// 	fmt.Println(l)
	// }

	writeTo(*outPath, "apps.csv", linesApp)
	writeTo(*outPath, "app_details.csv", linesWithEndpoint)

	return

}

func writeTo(path, filename string, lines []string) {
	f := path + "/" + filename
	file, err := os.Create(f)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range lines {
		v := strings.Split(value, ",")
		err := writer.Write(v)
		checkError("Cannot write to file", err)
	}

	log.Println("file '", f, "' created")
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
