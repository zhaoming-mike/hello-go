package main
import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	s := "postgress://user:pass@host.com:4545/path?k=v&k2=v2#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	p := fmt.Println
	p("----------------")
	p("url.Scheme:", u.Scheme)
	p("url.User:", u.User)
	p("url.User.Username():", u.User.Username())

	pass, _ := u.User.Password()
	p("url.User.Password:", pass)

	p("url.Host:", u.Host)

	host := strings.Split(u.Host, ":")
	p("host[0]", host[0])
	p("host[1]", host[1])

	p("url.Path:", u.Path)
	p("url.Fragment:", u.Fragment)
	p("url.RawQuery:", u.RawQuery)

	kvmap, _ := url.ParseQuery(u.RawQuery)
	p("map:", kvmap)
	p("map[\"k\"][0]:", kvmap["k"][0])
	//p("map[\"k\"][1]:", kvmap["k"][1]) //panic: runtime error: index out of range

}
