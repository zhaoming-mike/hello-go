package main
import "fmt"
import "time"

func main() {
	p := fmt.Println
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	p("now:", now)

	millis := nanos / 1000000
	p("secs:", secs)
	p("millis:", millis)
	p("nanos:", nanos)

	p("date:", time.Unix(secs, 0))
	p("date:", time.Unix(0, nanos))
}
