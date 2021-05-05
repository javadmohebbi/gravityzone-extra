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
	"github.com/javadmohebbi/gravityzone-extra/hardware"
)

/*
Author: M. Javad Mohebbi
This example will guide you how to get list of
detected hardware list on gravityzone
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

	iv := hardware.New(gzmongo.Client, context.Background())

	lst, err := iv.GetAll()

	if err != nil {
		panic(err)
	}

	var linesWithEndpoint, linesHwds []string
	linesWithEndpoint = append(linesWithEndpoint, fmt.Sprintf("%v,%v,%v,%v,%v",
		"EndpointName", "OS", "DeviceId", "DeviceName", "LastDetectionDate",
	))
	linesHwds = append(linesHwds, fmt.Sprintf("%v,%v,%v,%v",
		"DeviceId", "DeviceName", "LastDetectionDate", "# of Endpoints",
	))
	for _, l := range lst {
		for _, e := range l.Endpoints {
			ll := fmt.Sprintf("%v,%v,%v,%v,%v",
				e.Name, e.OperatingSystemVersion, l.DeviceID,
				l.DeviceName,
				l.LastDetectionDate,
			)
			linesWithEndpoint = append(linesWithEndpoint, ll)
		}

		ll := fmt.Sprintf("%v,%v,%v,%d",
			l.DeviceID,
			l.DeviceName,
			l.LastDetectionDate,
			len(l.Endpoints),
		)
		linesHwds = append(linesHwds, ll)

	}

	// for _, l := range linesApp {
	// 	fmt.Println(l)
	// }

	writeTo(*outPath, "hwds.csv", linesHwds)
	writeTo(*outPath, "hwds_details.csv", linesWithEndpoint)

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
