package main

import "fmt"
import "time"

func main() {
	p := fmt.Println
	now := time.Now()
	p("now:", now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("then:", then)

	p("year:", then.Year())
	p("month:", then.Month())
	p("day:", then.Day())
	p("hour:", then.Hour())
	p("minute:", then.Minute())
	p("second:", then.Second())
	p("nanosecond:", then.Nanosecond())
	p("location:", then.Location())
	p("weekday:", then.Weekday())
	p("before now:", then.Before(now))
	p("after now:", then.After(now))
	p("equal now:", then.Equal(now))

	diff := now.Sub(then)
	p("diff now:", diff)

	p("diff.Hours():", diff.Hours())
	p("diff.Minutes():", diff.Minutes())
	p("diff.Seconds():", diff.Seconds())
	p("diff.Nanoseconds():", diff.Nanoseconds())
	p("then add + diff:", then.Add(diff))
	p("then add - diff:", then.Add(-diff))
}
