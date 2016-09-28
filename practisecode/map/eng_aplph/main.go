package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www-personal.umich.edu/~jlawler/wordlist")
	if err != nil {
		log.Fatalln(err)
	}
	// bs; byteslice
	bs, _ := ioutil.ReadAll(res.Body) // allows to see body of the website
	str := string(bs)
	fmt.Println(str)
}
