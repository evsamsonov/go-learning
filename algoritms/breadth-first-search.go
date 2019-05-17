package main

import "fmt"

func main() {
	search("you")
}

func search(name string) bool {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	var searchQueue []string
	searchQueue = append(searchQueue, graph[name]...)
	searched := make(map[string]int)

	for len(searchQueue) > 0 {
		var person = searchQueue[0]
		searchQueue = searchQueue[1:]

		_, isSearched := searched[person]
		if isSearched == false {
			if isDragDialerPerson(person) {
				fmt.Println(person + " is the drag dialer!")
				return true
			}

			searchQueue = append(searchQueue, graph[person]...)
			searched[person] = 1
		}
	}

	return false
}

func isDragDialerPerson(name string) bool {
	return string(name[len(name) - 1]) == "m"
}
