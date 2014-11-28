/*
[2014-11-24] Challenge #190 [Easy] Webscraping sentiments
http://www.reddit.com/r/dailyprogrammer/comments/2nauiv/20141124_challenge_190_easy_webscraping_sentiments/
*/


package main

import (
	"strings"
//	"fmt"
	"io/ioutil"
	"net/http"
	"flag"
)


func get_comments(s string) []string {
	fields := strings.Split(s, "<div class=\"Ct\">")
	for i:=1; i< len(fields); i++{
		comment := strings.Split(fields[i], "</div>")
		fields[i] = comment[0]
//		println(fields[i])
	}
	return fields[1:]
}

func string_contains(s string, a []string)bool {
	for i := 0; i < len(a); i++ {
		if strings.Contains(s, a[i]){
			return true
		}
	}
	return false
}



func main() {
	flag.Parse()
	args := flag.Args()

	// Get word list
	url := "https://www.youtube.com/watch?v=RFinNxS5KN4"
	if len(args) > 0 {
		url = args[0]
	}
	prefix := "https://plus.googleapis.com/u/0/_/widget/render/comments?first_party_property=YOUTUBE&href="


	new_url := prefix + url

	resp, err := http.Get(new_url)

	if err == nil{
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			content := string(body)
			comments := get_comments(content)
//			println(comments)
			happy:= []string{ "love", "loved", "like", "liked", "awesome", "amazing", "good", "great", "excellent" }
			sad := []string { "hate", "hated", "dislike", "disliked", "awful", "terrible", "bad", "painful", "worst" }

			happy_count :=0
			sad_count := 0
			for i:=0;i<len(comments);i++{
				if string_contains(comments[i] ,happy){
					happy_count++
				}
				if string_contains(comments[i], sad){
					sad_count++
				}
			}


			println(happy_count)
			println(sad_count)


			switch{
			case happy_count > sad_count:
				println("HAPPY")
			case happy_count < sad_count:
				println("SAD")
			default:
				println("NEUTRAL")
			}

			// search for <div class="Ct"> until </div>


		}
	}


}
