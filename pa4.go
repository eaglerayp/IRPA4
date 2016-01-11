package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func showCluster(doc_cluster [1095]int) {
	ClusterDocsMap := make(map[int][]int)
	for index := 0; index < 1095; index++ {
		// if no this cluster, create slice
		if _, val := ClusterDocsMap[doc_cluster[index]]; val {
			list := []int{index} //array initialize
			ClusterDocsMap[doc_cluster[index]] = list
		} else {
			// else append
			ClusterDocsMap[doc_cluster[index]] = append(ClusterDocsMap[doc_cluster[index]], index)
		}
	}

	// display
	for _, list := range ClusterDocsMap {
		for doc := range list {
			fmt.Println(doc + 1)
		}
		fmt.Println()
	}
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
	// below doc always = index, means 1.txt = doc 0
	for doc1 := 0; doc1 < docnum; doc1++ {
		for doc2 := doc1 + 1; doc2 < docnum; doc2++ {
			sim := CosineSimilarity(doclist[doc1], doclist[doc2])
			SimArray = append(SimArray, sim)
			simIndex++
			array := []int{doc1, doc2}
			SimMap[sim] = array
		}
	}
	// initial doc-cluster map, doc 0 belongs to cluster 0
	var doc_cluster [1095]int
	for index := 0; index < 1095; index++ {
		doc_cluster[index] = index
	}
	sort.Float64s(SimArray)
	// small to big
	// single link,
	cluster_count := 1095
	for index := len(SimArray) - 1; cluster_count >= 8; index-- {
		// // max cosine~
		doc1 := SimMap[SimArray[index]][0]
		doc2 := SimMap[SimArray[index]][1]
		// //merge
		// if doc1 != doc2 {
		// 	cluster_count--
		// 	doc_cluster[doc2] = doc_cluster[doc1]
		// }
		fmt.Println(SimArray[index], doc1, doc2)
	}
	// showCluster(doc_cluster)
}
