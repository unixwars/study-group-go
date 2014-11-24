/*
http://www.reddit.com/r/dailyprogrammer/comments/2ledaj/11052014_challenge_187_intermediate_finding_time/
[11/05/2014] Challenge #187 [Intermediate] Finding Time to Reddit

Find longest slot for daily Reddit, and weekly analysis on minutes spent on tasks with percentage of time.
*/

package main

import (
  "os"
  "log"
  "bufio"
  "strings"
  "time"
  "sort"
  "fmt"
)

const (
  FILE_NAME string = "challenge-187-finding-time-to-reddit.input.txt"
  TIME_FMT string = "03:04 PM"
  BEGIN_FMT string = " 1-2-2006: " + TIME_FMT
  DATE_FMT string = "2006-01-02 Monday"
  REDDIT string = "reddit"
)

type Event struct {
  Begin time.Time
  End time.Time
  Duration time.Duration
  Title string
}

func (e Event) String() string {
  return fmt.Sprintf("%s - %s (%7s) %s", e.Begin.Format(TIME_FMT), e.End.Format(TIME_FMT), e.Duration, e.Title)
}


// Types and sort.Interface
type Events []*Event
type EventsByDate     struct { Events }
type EventsByDuration struct { Events }

func (e Events)           Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e Events)           Len() int           { return len(e) }
func (e EventsByDate)     Less(i,j int) bool  { return e.Events[i].Begin.Before(e.Events[j].Begin) }
func (e EventsByDuration) Less(i, j int) bool { return e.Events[i].Duration < e.Events[j].Duration }

func main() {
  events := ParseFile(FILE_NAME)
  sort.Sort(EventsByDate{events})

  day := Events{}

  // Find Reddit time
  for slicer, next := DaySlicer(events), true ; next ; {
    day, next = slicer()
    r := MakeRedditEvent(day)
    events = append(events, r)
  }

  // Show days
  sort.Sort(EventsByDate{events})
  for slicer, next := DaySlicer(events), true ; next ; {
    day, next = slicer()
    fmt.Println("\nEvents for", day[0].Begin.Format(DATE_FMT))
    ShowEvents(day)
  }

  // Stats
  ShowStats(events)
}

func ParseFile(filename string) Events {
  var events Events

  file, err := os.Open(filename)
  if err != nil {
     log.Fatal(err)
  }

  defer file.Close()
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    event,ok := ParseEvent(line)
    if ok {
      events = append(events, event)
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return events
}

func ParseEvent(line string) (*Event, bool) {
    event := &Event{}

    tokens := strings.Split(line, " to ")
    fromStr, restStr := tokens[0], tokens[1]
    begin, err := time.Parse(BEGIN_FMT, fromStr)
    if err != nil {
      return event, false
    }
    
    tokens = strings.Split(restStr, " -- ")
    deltaStr, title := tokens[0], tokens[1]
    delta, err := time.Parse(TIME_FMT, deltaStr)
    if err != nil {
      return event, false
    }

    event.Begin = begin
    event.End = delta.AddDate(begin.Year(), int(begin.Month())-1, begin.Day()-1)
    event.Duration = event.End.Sub(begin)
    event.Title = title

    return event, true
}

/* 
Generate event slices for same consecutive day
*/
func DaySlicer(events Events) func() (Events, bool) {
    start := 0
    stop := 0
    i := 0

    isSameDay := func(begin, end time.Time) bool {
      return begin.Year() == end.Year() && begin.YearDay() == end.YearDay()
    }

    return func() (Events, bool) {
      for isSameDay(events[start].Begin, events[stop].Begin) && stop < len(events)-1 {
        stop++
      }

      i, start = start, stop
      next := true
      if stop == len(events)-1 {
      	 next = false
      }
      return events[i:stop], next
    }
}

func MakeRedditEvent(day Events) *Event {
  slots := Events{}

  // Calculate free slots
  for i,j := 0,1 ; j < len(day) ; i,j = i+1, j+1 {
    duration := day[j].Begin.Sub(day[i].End)
    event := &Event{
      Begin: day[i].End,
      End: day[i].End.Add(duration),
      Duration: duration,
      Title: REDDIT,
    }
    slots = append(slots, event)
  }

  // Return longest
  sort.Sort(EventsByDuration{slots})
  return slots[len(slots)-1]
}

func ShowEvents(events Events) {
  for i,e := range events {
    fmt.Printf("%2d %s\n", i, *e)
  }
}

func ShowStats(events Events) {
  activities := make(map[string]time.Duration)
  total := time.Duration(0)

  for _,e := range events {
    activities[e.Title] = activities[e.Title] + e.Duration
    total += e.Duration
  }

  fmt.Printf("\nStats for total time of %s\n", total)
  for k,v := range activities {
    fmt.Printf("%8s (%5.2f%%)\t\t%s\n", v, float64(v)*100/float64(total), k) 
  }
}
