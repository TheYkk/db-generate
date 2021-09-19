package main

import (
	"bufio"
	"github.com/bits-and-blooms/bloom/v3"
	"log"
	"os"
)

func main() {
	passFile := os.Args[1]
	log.SetFlags(log.Lshortfile)

	// Error rate is 1 in 1 Million
	filter := bloom.NewWithEstimates(6000000, 0.000001)

	pass, err := os.Open(passFile)
	if err != nil {
		log.Fatal(err)
	}
	defer pass.Close()

	scanner := bufio.NewScanner(pass)
	for scanner.Scan() {
		// First 40 chars is sha1 sum
		filter.AddString(scanner.Text()[:40])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fl, err := os.Create("db.db")
	if err != nil {
		log.Fatal(err)
	}
	// Export bitset to file
	_, err = filter.WriteTo(fl)
	if err != nil {
		log.Fatal(err)
	}
}
