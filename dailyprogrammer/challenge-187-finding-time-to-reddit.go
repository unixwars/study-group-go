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
)

type Event struct {
  Begin time.Time
  End time.Time
  Duration time.Duration
  Title string
}

type Schedule []*Event

// Schedule implements sort.Interface for []Event based on the Begin field.
func (a Schedule) Len() int           { return len(a) }
func (a Schedule) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Schedule) Less(i, j int) bool { return a[i].Begin.Before(a[j].Begin) }

// ByDuration implements sort.Interface for []Event based on the Duration field.
type ByDuration Schedule
func (a ByDuration) Less(i, j int) bool { return a[i].Duration < a[j].Duration }

func ParseFile(filename string) *Schedule{
  var sched Schedule

  file, err := os.Open(filename)
  if err != nil {
     log.Fatal(err)
  }

  defer file.Close()
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    event,ok := ParseLine(line)
    if ok {
      sched = append(sched, event)
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  sort.Sort(sched)
  return &sched
}

func ParseLine(line string) (*Event, bool) {
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
Returns a closure that can be used to iterate over Event slices containing Events of the same day. 
The input is assumed sorted, so unsorted arrays might return different slices with elements of same dates.
*/
func DayIterator(events []*Event) func() ([]*Event, bool) {
    start := 0
    stop := 0
    i := 0

    return func() ([]*Event, bool) {
      for IsSameDay(events[start].Begin, events[stop].Begin) && stop < len(events) {
        stop++
      }

      i, start = start, stop
      next := true
      if stop == len(events) {
      	 next = false
      }
      fmt.Println("++++ ", i, start, stop, len(events), next)
      return events[i:stop], next
    }
}

func IsSameDay(begin, end time.Time) bool {
  return begin.Year() == end.Year() && begin.YearDay() == end.YearDay()
}

func main() {
  events := ParseFile(FILE_NAME)
  printArray(events)
  iterator := DayIterator(*events)

  for {
    day, next := iterator()
    fmt.Println("\n** ", next)
    if !next {
      break
    }
    printArray(day)
  }
}

func printArray(events []*Event) {
     for i,e := range events {
     	 fmt.Println(i," ", e)
     }
     fmt.Println("--- Length: ", len(events))
}

