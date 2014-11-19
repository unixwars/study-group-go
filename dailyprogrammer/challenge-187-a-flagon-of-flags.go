/*
(Easy): A Flagon of Flags
http://www.reddit.com/r/dailyprogrammer/comments/2l6dll/11032014_challenge_187_easy_a_flagon_of_flags/
*/

package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    var n int
    var line string

    fmt.Scanf("%d", &n)
    cmds := make(map [string]string)

    scanner := bufio.NewScanner(os.Stdin)
    for ; scanner.Scan() ; n-- {
        line = scanner.Text()
	if n==0 {
	   break
	}
	tokens := strings.Split(line, ":")
	short, long := tokens[0], tokens[1]
	cmds[short] = long
    }

    tokens := strings.Split(line, " ")
    for _, token := range tokens {
    	trimmed := strings.TrimLeft(token, "-")
    	if strings.HasPrefix(token, "--") {
	   fmt.Println("flag: ", trimmed)
	} else if strings.HasPrefix(token, "-") {
	   for _, p:= range trimmed {
	   	fmt.Println("flag: ", cmds[string(p)])
	   }
	} else {
	  fmt.Println("parameter: ", token)
	}
    }
}
