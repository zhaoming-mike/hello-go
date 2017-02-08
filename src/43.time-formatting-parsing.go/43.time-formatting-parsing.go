package main
import "fmt"
import "time"

func main() {
	p := fmt.Println

	t := time.Now()

	p(t.Format("2006-01-02T15:04:05Z07:00")) //RFC3339
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:01:02.999999-07:00"))


	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	withNanos := "2006-01-02T15:00:44.999444999-09:00"
	t1, e := time.Parse(withNanos, "2012-02-03T22:22:22.112233+00:00")
	p("t1:", t1)

	kitchen := "3:09PM"
	t2, e := time.Parse(kitchen, "9:03PM")
	p("t2:", t2)

	ansic := "Mon Jan _2 15:22:23 2006"
	_, e = time.Parse(ansic, "09:38PM")
	p("e:", e)	//e: parsing time "09:38PM" as "Mon Jan _2 15:22:23 2006": cannot parse "09:38PM" as "Mon"

	p("time.Kitchen:", t.Format(time.Kitchen))

}
