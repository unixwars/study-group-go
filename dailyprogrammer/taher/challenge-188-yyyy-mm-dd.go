/*
iso 8601 standard for dates tells us the proper way to do an extended day is yyyy-mm-dd
yyyy = year
mm = month
dd = day

A company's database has become polluted with mixed date formats. They could be one of 6 different formats
 yyyy-mm-dd
 mm/dd/yy
 mm#yy#dd
 dd*mm*yyyy
 (month word) dd, yy
 (month word) dd, yyyy

 (month word) can be: Jan Feb Mar Apr May Jun Jul Aug Sep Oct Nov Dec
Note if is yyyy it is a full 4 digit year. If it is yy then it is only the last 2 digits of the year. Years only go between 1950-2049.

Input:
You will be given 1000 dates to correct.

Output:
You must output the dates to the proper iso 8601 standard of yyyy-mm-dd

Challenge Input:
 https://gist.github.com/coderd00d/a88d4d2da014203898af
*/

package main

import (
    "fmt"
    "time"
    "errors"
    "os"
    "bufio"
)

// Reference date: Mon Jan 2 15:04:05 MST 2006
var dateFormats = []string{"2006-01-02", "01/02/06", "01#06#02", "02*01*2006", "Jan 02, 06", "Jan 02, 2006"}
var isoFormat = dateFormats[0]
var inputFile = "challenge-188-yyyy-mm-dd.input.txt"

func Parse(dateString string) (string, error) {
    for _, df := range dateFormats {
    	dt, err := time.Parse(df, dateString)
	if err == nil {
	   if dt.Year() > 2049 {
	      dt = dt.AddDate(-100,0,0)
	   }
	   return dt.Format(isoFormat), nil
	}
    }
    return "", errors.New("Unrecognised format")
}


func main() {
    file, err := os.Open(inputFile)
    defer file.Close()
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	line := scanner.Text()
    	iso, err := Parse(line)
	if err == nil {
	   fmt.Println(iso)
	} else {
	   fmt.Println("**********", line)
	}
    }

    if err := scanner.Err(); err != nil {
     	panic(err)
    }
}

