package main

import (
    "fmt"
    "os"
    "regexp"
)

const dieRollRegex string = `(\d+)?d\d+`
const constantRegex string = `[+|-]\d+\b`


func interpretRollStatement(rollStatement string) ([]string, []string) {
    dieRollCompiled := regexp.MustCompile(dieRollRegex)
    constantCompiled := regexp.MustCompile(constantRegex)
    matchedDieRolls := dieRollCompiled.FindAllString(rollStatement, -1)
    matchedConstants := constantCompiled.FindAllString(rollStatement, -1)
    fmt.Printf("Dice: %v\n", matchedDieRolls)
    fmt.Printf("Constants: %v\n", matchedConstants)
    return matchedDieRolls, matchedConstants
}


func main() {
    rollStatement := os.Args[1]
    interpretRollStatement(rollStatement)
}
