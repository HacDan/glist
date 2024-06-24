package main

import (
    "fmt"
    "strings"
    "os"

    "github.com/juliangruber/go-intersect/v2"
)

func main() {
    if len(os.Args) == 1 {
        // TODO: Write help message here
        fmt.Println("No command supplied")
        os.Exit(1)
    }

    command := os.Args[1]
     
    delimiter := findDelimter(os.Args[2])
     

    s1 := splitString(os.Args[2], delimiter)
    s2 := splitString(os.Args[3], delimiter)

    switch command {
    case "-u":
        fmt.Println(findUnion(s1, s2))
    case "-x":
        fmt.Println(findIntersect(s1, s2))
    case "-n":
        fmt.Println(findComplements(s1, s2))
    default:
        // TODO: Print help message here along with error
        fmt.Println("Incorrect command entered")
    }
}

func removeDuplicateStr(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func splitString(s string, d string) []string {
    return cleanSlice(strings.Split(s, d))    
}

func cleanSlice(s []string) []string {
    c := []string{}
    for _, item := range s {
        c = append(c, strings.TrimSpace(item))
    }
    return c
}

func findComplements(s1, s2 []string) []string {
    intersection := intersect.SimpleGeneric(s1, s2)
    s3 := append(s1, s2...)
    final := []string {}

    for _, s := range s3 {
        if contains(s, intersection) != true {
            final = append(final, s)      
        }
    }
    return final
}

func contains(s string, c []string) bool {
    for _, str := range c {
        if str == s {
            return true
        }
    }
    return false
}
func findUnion(s1, s2 []string) []string {
    s3 := append(s1, s2...)
    return removeDuplicateStr(s3)
}

func findIntersect(s1, s2 []string) []string {
    return intersect.SimpleGeneric(s1, s2)
}

func findDelimter(s string) string {
    // Common delimiters - add to in the future
    DELIMITERS := []string{",", " ", "\t", ";", "|", "\n"}
    for _, c := range DELIMITERS {
        if strings.Contains(s, c) {
            return c
        }
    }
    // TODO: Find way to return error here
    return ""
}

