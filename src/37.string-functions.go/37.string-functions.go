package main
import "fmt"
import s "strings"

var p = fmt.Println

func main() {
	p("Contains:	", s.Contains("test", "es"))
	p("Count:		", s.Count("test", "t"))
	p("HasPrefix:	", s.HasPrefix("test", "te"))
	p("HasSufffix:	", s.HasSuffix("test", "st"))
	p("Index:		", s.Index("test", "e"))
	p("Join:		", s.Join([]string { "a", "b", "c"}, "--"))
	p("Repeat:		", s.Repeat("a", 10))
	p("Replace:	", s.Replace("foof00foof00foo", "o", "0", -1))
	p("Replace:	", s.Replace("foof00foof00foo", "o", "0", 3))
	p("Split:		", s.Split("a/b/f/f/d/sd/fs/fsfs/sdfsdf", "/"))
	p("转成大写:	", s.ToUpper("mike is smart boy!"))
	p("转成小写:	", s.ToLower("MICHAEL JACKSON"))
	p()
}
