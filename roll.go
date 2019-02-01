package main

import (
    "fmt"
    "log"
    "math"
    "math/rand"
    "os"
    "regexp"
    "strconv"
    "strings"
    "time"
)

const dieRollRegex = `[b|w]?(\d+)?d\d+`
const constantRegex = `[+|-]\d+\b`
var fullRollStatementRegex = fmt.Sprintf(`^(%s[+|-]?)*(\d+)*$`, dieRollRegex)
var verbose = false

func max(integers []int) int {
    max := 0
    for _, n := range integers {
        if n > max {
            max = n
        }
    }
    return max
}


func min(integers []int) int {
    min := math.MaxInt64
    for _, n := range integers {
        if n < min {
            min = n
        }
    }
    return min
}


func sum(integers []int) int {
    sum := 0
    for _, n := range integers {
        sum += n
    }
    return sum
}


func interpretRollStatement(rollStatement string) ([]string, []string) {
    dieRollCompiled := regexp.MustCompile(dieRollRegex)
    constantCompiled := regexp.MustCompile(constantRegex)
    matchedDieRolls := dieRollCompiled.FindAllString(rollStatement, -1)
    matchedConstants := constantCompiled.FindAllString(rollStatement, -1)
    return matchedDieRolls, matchedConstants
}


func rollDice(dieRolls []string) int {
    sum := 0
    //random := rand.New(rand.NewSource(time.Now().UnixNano()))
    rand.Seed(time.Now().UnixNano())
    for _, dieRoll := range dieRolls {
        sum += rollDie(dieRoll)
    }
    return sum
}

func rollDie(dieRoll string) int {
    mode := 'n' // for 'normal'
    if dieRoll[0] == 'b' || dieRoll[0] == 'w' {
        mode = rune(dieRoll[0])
        dieRoll = dieRoll[1:]
    }

    dieSplit := strings.Split(dieRoll, "d")
    dieCount, err := strconv.Atoi(dieSplit[0])
    if err != nil {
        dieCount = 1 // handle rolls like 'd6' instead of '1d6'
    }

    faces, err := strconv.Atoi(dieSplit[1])
    if err != nil {
        log.Fatal(err)
    }

    var resultList []int
    for n := 0; n < dieCount; n++ {
        resultList = append(resultList, rand.Intn(faces) + 1)
    }
    if verbose {
        fmt.Printf("Rolls: %v\n", resultList)
    }
    if mode == 'b' {
        return max(resultList)
    } else if mode == 'w' {
        return min(resultList)
    }
    return sum(resultList)
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


func parseArguments(args []string) ([]string, []string, []string) {
    var rollStatements []string
    var options []string
    var malformed []string
    fullRollStatementCompiled := regexp.MustCompile(fullRollStatementRegex)

    for _, arg := range args {
        if fullRollStatementCompiled.FindString(arg) != "" {
            rollStatements = append(rollStatements, arg)
        } else if arg == "-v" {
            options = append(options, arg)
        } else {
            malformed = append(malformed, arg)
        }
    }
    return rollStatements, options, malformed
}


func main() {
    rollStatements, options, malformed := parseArguments(os.Args[1:])
    if malformed != nil {
        fmt.Printf("ERROR: malformed argument(s): %v\n", malformed)
        os.Exit(1)
    }

    for _, opt := range options {
        if opt == "-v" {
            verbose = true
        }
    }

    for _, rollStatement := range rollStatements {
        dieRolls, constants := interpretRollStatement(rollStatement)
        dieResult := rollDice(dieRolls)
        constantResult := addConstants(constants)
        fmt.Printf("Result: %v\n", dieResult + constantResult)
    }
}
