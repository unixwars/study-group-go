/*
[2014-11-17] Challenge #189 [Easy] Hangman!
http://www.reddit.com/r/dailyprogrammer/comments/2mlfxp/20141117_challenge_189_easy_hangman/
*/


package main

import (
	"strings"
	"fmt"
	"math/rand"
	"time"
	"io/ioutil"
	"flag"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

// word list - http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt
func read_word_file(filename string)string{
	data,_ := ioutil.ReadFile(filename)
	return string(data)
}

func print_info(current []string,
				guesses_left int,
				guessed_letters string){

	draw_hangman(guesses_left)
	word := strings.Join(current ," " )
	fmt.Printf("%s    Guesses left:%2d       Guessed Letters:%s\n",word,guesses_left,guessed_letters)
}


func draw_hangman(guesses_left int){
	a:= []string {
		"  +---+ ",
		"  |   | ",
		"  0   | ",
		" /|\\  | ",
		" / \\  | ",
		"      | ",
		"======= "}

	b:= []string {
		"  77778 ",
		"  6   8 ",
		"  5   8 ",
		" 342  8 ",
		" 1 0  8 ",
		"      8 ",
		"99999999"}

	for x:=0; x<len(a); x++{
		arow := a[x]
		brow := b[x]
		var new_row [8]byte
		for y:=0; y<len(arow); y++ {
			t := brow[y]
			s := arow[y]
			if t != 32{
				i := t-48
				if int(i) < guesses_left {
					s = 32
				}
			}
			new_row[y] = s
		}
		fmt.Println(string(new_row[:]))
	}
}



func main(){
	flag.Parse()
	args := flag.Args()

	// Get word list
	word_list := "twas brillig and the slythy toves did gyre and gimble in the wabe jazz quiz"
	if len(args) > 0 {
		word_list = read_word_file(args[0])
	}

	words := strings.Fields(word_list)
	fmt.Println("Word count =", len(words))

	// select a word
	myword := strings.ToLower( words[ random(0, len(words)-1) ] )
	word_len := len(myword)
	characters := strings.Split(myword, "")

	// Debug stuff
	fmt.Println("My word =", myword, characters)
	fmt.Println("Word length =", word_len)

	guessed := strings.Split(strings.Repeat("_", word_len),"")
	guessed_chars := ""
	guesses_left := 10

	var guess string
	letters_left := 1
	for ; guesses_left>0 && letters_left > 0 ; {
		print_info(guessed, guesses_left, guessed_chars)
		fmt.Scanf("%s", &guess)
		guessed_chars = guessed_chars+guess
		fail := 1
		letters_left = 0
		for i := 0; i < word_len; i++ {
			if characters[i] == guess {
				guessed[i] = guess
				fail = 0
			}
			if guessed[i] == "_" {
				letters_left++
			}
		}
		guesses_left = guesses_left - fail
	}

	print_info(guessed, guesses_left, guessed_chars)

	if letters_left > 0{
		fmt.Printf("FAILED !!   The word was %s", myword)
	}else{
		fmt.Println("SUCCESS !!")
	}
}
