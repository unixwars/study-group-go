package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (int, error) {
    n, err := rot.r.Read(p)
    for i, v := range p {
	if  v >= 'A' && v <= 'M' || v >= 'a' && v <= 'm' {
    	     p[i] += 13
	} else if v >= 'N' && v <= 'Z' || v >= 'n' && v <= 'z' {
	     p[i] -= 13
	}
    }
    return n, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
