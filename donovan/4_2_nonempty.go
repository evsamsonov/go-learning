package main

import "fmt"

func main() {
	s := []string{"test", "", "nonempty"}
	s = nonempty(s)
	fmt.Println(s)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s;
			i++
		}
	}

	return strings[:i]
}
