package main

import (
	"fmt"
	"sort"
)

// ç»å®ä¸äºè®¡ç®æºè¯¾ç¨ï¼æ¯ä¸ªè¯¾ç¨é½æåç½®è¯¾ç¨ï¼åªæå®æäºåç½®è¯¾ç¨æå¯ä»¥å¼å§å½åè¯¾ç¨çå­¦ä¹ ï¼æä»¬çç®æ æ¯éæ©åºä¸ç»è¯¾ç¨ï¼è¿ç»è¯¾ç¨å¿é¡»ç¡®ä¿æé¡ºåºå­¦ä¹ æ¶ï¼è½å¨é¨è¢«å®æãæ¯ä¸ªè¯¾ç¨çåç½®è¯¾ç¨å¦ä¸
// prereqsè®°å½äºæ¯ä¸ªè¯¾ç¨çåç½®è¯¾ç¨
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// æææåºï¼åç½®æ¡ä»¶æææåå¾
// é¡¶ç¹è¡¨ç¤ºè¯¾ç¨ï¼è¾¹è¡¨ç¤ºä¾èµå³ç³»
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	seen := make(map[string]bool)
	var order []string
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}

	}
	var keys []string

	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	visitAll(keys)
	return order
}
