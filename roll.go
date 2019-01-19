package main

import (
    "fmt"
    "log"
    "math/rand"
    "os"
    "regexp"
    "strconv"
    "strings"
    "time"
)

const dieRollRegex string = `(\d+)?d\d+`
const constantRegex string = `[+|-]\d+\b`


func interpretRollStatement(rollStatement string) ([]string, []string) {
    dieRollCompiled := regexp.MustCompile(dieRollRegex)
    constantCompiled := regexp.MustCompile(constantRegex)
    matchedDieRolls := dieRollCompiled.FindAllString(rollStatement, -1)
    matchedConstants := constantCompiled.FindAllString(rollStatement, -1)
    return matchedDieRolls, matchedConstants
}


func rollDice(dieRolls []string) int {
    sum := 0
    random := rand.New(rand.NewSource(time.Now().UnixNano()))
    for _, dieRoll := range dieRolls {
        dieSplit := strings.Split(dieRoll, "d")

        amount, err := strconv.Atoi(dieSplit[0])
        if err != nil {
            amount = 1 // handle rolls like 'd6' instead of '1d6'
        }

        faces, err := strconv.Atoi(dieSplit[1])
        if err != nil {
            log.Fatal(err)
        }

        for n := 0; n < amount; n++ {
            sum += random.Intn(faces) + 1
        }
    }
    return sum
}


func addConstants(constants []string) int {
    sum := 0
    for _, constant := range constants {
        sign := constant[0]
        value, err := strconv.Atoi(constant[1:])
        if err != nil {
            log.Fatal(err)
        }
        if sign == '+' {
            sum += value
        } else {
            sum -= value
        }
    }
    return sum
}


func main() {
    rollStatement := os.Args[1]
    dieRolls, constants := interpretRollStatement(rollStatement)
    dieResult := rollDice(dieRolls)
    constantResult := addConstants(constants)
    fmt.Printf("Result: %v\n", dieResult + constantResult)
}
