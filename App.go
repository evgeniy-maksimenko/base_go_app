package main


import (
	"fmt"
	s "strings"
)

var log = fmt.Println

func main() {

	log("Contains:  ", s.Contains("test","t"))
	log("Count:     ", s.Count("test", "t"))
	log("HasPrefix: ", s.HasPrefix("test","te"))
	log("HasSuffix: ", s.HasSuffix("test","st"))
	log("Index:     ", s.Index("test", "e"))
	log("Join:      ", s.Join([]string{"a", "b"}, "-"))
	log("Repeat:    ", s.Repeat("a", 5))
	log("Replace:   ", s.Replace("foo", "o", "0", -1))
	log("Replace:   ", s.Replace("foo", "o", "0", 1))
	log("Split:     ", s.Split("a-b-c-d-e","-"))
	log("ToLower:   ", s.ToLower("TEST"))
	log("ToUpper:   ", s.ToUpper("test"))

}
