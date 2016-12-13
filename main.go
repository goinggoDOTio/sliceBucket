package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book Adventures of Sherlock Holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// create a slice of slice of string to hold slices of words
	buckets := make([][]string, 12)

	// loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	// print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, "-", len(buckets[i]))
	}
	// print the words in one of the buckets
	// fmt.Println(buckets[6])
	fmt.Println("len -", len(buckets))
	fmt.Println("cap -", cap(buckets))
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
