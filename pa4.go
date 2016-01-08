package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readDoc(docid int) map[int]float64 {
	file, err := os.Open("docs/" + strconv.Itoa(docid) + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	m := make(map[int]float64)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		terms := strings.Split(scanner.Text(), "\t\t")
		id, _ := strconv.Atoi(terms[0])
		tfidf, _ := strconv.ParseFloat(terms[1], 64)
		m[id] = tfidf
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return m
}

func CosineSimilarity(doc1 map[int]float64, doc2 map[int]float64) float64 {
	Similarity := 0.0
	for index, _ := range doc1 {
		Similarity += doc1[index] * doc2[index]
	}
	return Similarity
}

func main() {
	docnum := 1095
	var doclist [1095]map[int]float64
	for i := 1; i <= docnum; i++ {
		doclist[i-1] = readDoc(i)
	}
	var SimArray []float64
	simIndex := 0
	SimMap := make(map[float64][]int)
	for doc1 := 0; doc1 < docnum; doc1++ {
		for doc2 := doc1 + 1; doc2 < docnum; doc2++ {
			sim := CosineSimilarity(doclist[doc1], doclist[doc2])
			SimArray = append(SimArray, sim)
			simIndex++
			array := []int{doc1, doc2}
			SimMap[sim] = array
		}
	}
	fmt.Println(simIndex)
}
