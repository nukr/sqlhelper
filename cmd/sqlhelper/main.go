package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	sqlhelper "github.com/nukr/sqlhelper/pkg"
)

func main() {
	pkg := os.Args[1]
	fmt.Printf("package %s\n\n", pkg)
	dir, err := ioutil.ReadDir("files")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range dir {
		file, err := os.Open("files/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		scanner := sqlhelper.NewScanner(bufio.NewReader(file))
		parser := sqlhelper.NewParser(scanner)
		parser.Parse(os.Stdout)
	}
}
