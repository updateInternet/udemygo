package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-personal.umich.edu/~jlawler/wordlist")
	if err != nil {
		log.Fatalln(err)
	}

	//put all words in a map; it makes bit faster
	words := make(map[string]string)

	// sc; scanner
	// NewScanner calls Reader interface,
	sc := bufio.NewScanner(res.Body)
	//bufio points to a *Scanner (pointer scanner) which allows to call Split
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k, _ := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}
}
