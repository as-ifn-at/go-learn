package main

import (
	"fmt"
	"log/slog"
	"strings"
)

func main() {
	slog.Debug("askkd")

	// s := "Hello, 世界"
	// for i, rune := range s {
	// 	fmt.Printf("%d: %v\n", i, string(rune))
	// }

	// var sb strings.Builder

	// sb.WriteString("abcde")

	// fmt.Println(sb.Cap())

	// fmt.Println(sb.Len())

	// sb.WriteString("fghijkl")

	// fmt.Println(sb.Cap())

	// fmt.Println(sb.Len())

	// sb.WriteString("mnopqr")

	// fmt.Println(sb.Cap())

	// fmt.Printf("sb: %v\n", sb)

	// fmt.Printf("sb.String(): %v\n", sb.String())

	// t := "nijas"
	// r := []rune(t)
	// te := r[0]
	// r[0] = r[1]
	// r[1] = te

	// fmt.Printf("t: %v\n", t)

	// fmt.Printf("r: %v\n", r)
	// // strings.Compare()

	// fmt.Println("rune data here!")

	// for i := 0; i < len(r) - 1; i++ {
	// 	r[i] = r[i+1]
	// }
	// fmt.Printf("r: %v\n", r)

	// fmt.Printf("checkSameparity(): %v\n", checkSameparity(50,49))

	fmt.Printf("getSmallestString(): %v\n", getSmallestString("45320"))

}

func getSmallestString(s string) string {
    
    r := []rune(s)
    l := len(s)
    res := s

    for i:= 0; i<l-1; i++ {
        if checkSameparity(r[i], r[i+1]) {
            swap(i, i+1, r)
            if strings.Compare(string(r), res) < 0{
                res = string(r)
            } else {
                swap(i, i+1, r)
            }
        }
    }

    return res
    
}

func swap(i, j int, r []rune) {
    te := r[i]
    r[i] = r[j]
    r[j] = te
}

func checkSameparity(a, b rune) bool {
    if a % 2 == 0 && b % 2 == 0 {
        return true
    }
    if (b % 2 == 1 || b == 50) && (a % 2 == 1 || a == 50){
        return true
    }

    return false
}


