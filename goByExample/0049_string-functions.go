package main

import (
	"fmt"
	. "strings"
)

var p = fmt.Println

func main() {
	p("Contains:	", Contains("test", "es"))
	p("Count:		", Count("test", "t"))
	p("HasPrefix:	", HasPrefix("test", "te"))
	p("HasSuffix:	", HasSuffix("test", "st"))
	p("Index:		", Index("test", "e"))
	p("Join:		", Join([]string{"a", "b"}, "-"))
	p("Repeat:		", Repeat("a", 5))
	p("Replace:		", Replace("foo", "o", "0", -1))
	p("Replace:		", Replace("foo", "o", "0", 1))
	p("Split:		", Split("a-b-c-d-e", "-"))
	p("ToLower:		", ToLower("TEST"))
	p("ToUpper:		", ToUpper("test"))
}
