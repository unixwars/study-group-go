/*
[2014-11-24] Challenge #190 [Easy] Webscraping sentiments
http://www.reddit.com/r/dailyprogrammer/comments/2nauiv/20141124_challenge_190_easy_webscraping_sentiments/
*/

package main

import (
  "fmt"
  "regexp"
  "net/http"
  "log"
  "io/ioutil"
  "strings"
  "flag"
)

const (
  url_prefix    = "https://plus.googleapis.com/u/0/_/widget/render/comments?first_party_property=YOUTUBE&href="
  default_video = "https://www.youtube.com/watch?v=qpgTC9MDx1o"
  comment_re    = "<div class=\"[C,c][T,t]\">(.*?)</div>"
)


/* 
 * Sentiment stuf 
 */
type Sentiment int
type Sentiments []Sentiment
type Frequencies map[Sentiment]int

const (
  NONE Sentiment = iota
  SAD 
  HAPPY
)

var sentiments = []string{"Undetermined", "Sad", "Happy"}
func (s Sentiment) String() string {
  return sentiments[int(s)]
}

func (freqs Frequencies) MostFrequent() Sentiment {
  max := NONE
  for k,v := range freqs {
    if v > freqs[max] {
      max = k
    }
  }
  return max
}


/* 
 * Corpus stuff
 */
type Keywords []string
type Corpus []Keywords

var corpus = Corpus{
  Keywords{},
  Keywords{"hate","hated","dislike","disliked","awful","terrible","bad","painful","worst","idiot"},
  Keywords{"love","loved","like","liked","awesome","amazing","good","great","excellent","nice","amazing"},
}

func (slice Keywords) Contains(value string) bool {
  for _, v := range slice {
    if (strings.Contains(value, v)) {
      return true
    }
  }
  return false
}

func (c Corpus) Evaluate(text string) Sentiment {
  words := strings.Split(text, " ")

  freqs := make(Frequencies)
  for _, word := range words {
    for i, words := range c {
      if words.Contains(word) {
        freqs[Sentiment(i)] += 1
      }
    }
  }
  return freqs.MostFrequent()
}


/*
 * Other ...
 */
func Download(url string) string {
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer	res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    log.Fatal(err)
  }
  return string(body)
}

func ParseComments(text string) []string {
  text = strings.ToLower(text)
  text = strings.Replace(text, ",", " ", -1)
  text = strings.Replace(text, ".", " ", -1)
  text = strings.Replace(text, ";", " ", -1)
  text = strings.Replace(text, ":", " ", -1)

  re := regexp.MustCompile(comment_re)
  submatches := re.FindAllStringSubmatch(text, -1)

  comments := make([]string, len(submatches))
  for i, submatch := range submatches {
    comments[i] = submatch[1]
  }

  return comments
}

func main() {
  flag.Parse()
  args := flag.Args()
  video_url := default_video
  if len(args) > 0 {
    video_url = args[0]
  }

  body := Download(url_prefix + video_url)
  comments := ParseComments(body)

  counter := make(Frequencies)
  total := 0

  for _, cmt := range comments {
    total++
    sentiment := corpus.Evaluate(cmt)
    if sentiment != NONE {
        counter[sentiment] += 1
    }

    fmt.Println("Sentence: ", cmt)
    fmt.Println("Sentiment:", sentiment)
  }

  fmt.Println("\nSample size:", total, "persons")
  fmt.Println("General sentiment:", counter.MostFrequent(), counter)
}
