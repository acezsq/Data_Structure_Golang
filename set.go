package main

import "fmt"

type set map[string]struct{}

func (s set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s set) Add(key string) {
	s[key] = struct{}{}
}

func (s set) Delete(key string) {
	delete(s, key)
}

func main() {
	s := make(set)
	s.Add("zsq")
	s.Add("ace")
	fmt.Println(s.Has("zsq")) // true
	fmt.Println(s.Has("ace")) // true
	s.Delete("zsq")
	fmt.Println(s.Has("zsq")) // false
}
