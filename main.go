package main

import "log"

func main() {
	news, err := rbcParse()
	if err != nil {
		log.Fatal(err)
	}

    log.Printf("%+v", news)
}
