package main
import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "bufio"
    "os"
    "strings"
    "time"
)
type Config map[string][]string

func getJson() Config {
    file, _ := ioutil.ReadFile("./words.json")
    var parsed Config
    _ = json.Unmarshal(file, &parsed)
    return parsed
}

func getAnswer(config Config) (string, int) {
    wordleStartDate := time.Date(2021, time.June, 19, 0, 0, 0, 0, time.Local)
    year, month, day := time.Now().Date()
    todayDate := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
    answerIndex := int(todayDate.Sub(wordleStartDate) / (24 * time.Hour))
    return config["answers"][answerIndex], answerIndex
}

func contains(arr []string, elem string) bool {
    for _, e := range arr {
        if e == elem {
            return true
        }
    }
    return false
}

func colorizedResponse(word string, answer string) string {
    var normalText, normalEmoji string = "\033[0m", "\u2B1C"
    var greenText, greenEmoji string = "\033[32m", "\U0001f7e9"
    var yellowText, yellowEmoji string = "\033[33m", "\U0001f7e8"
    var emojis, text string = "", ""
    for i, char := range strings.Split(word, "") {
        if char == string(answer[i]) {
            emojis += greenEmoji
            text += greenText + strings.ToUpper(char) + normalText
        } else if contains(strings.Split(answer, ""), char) {
            emojis += yellowEmoji
            text += yellowText + strings.ToUpper(char) + normalText
        } else {
            emojis += normalEmoji
            text += strings.ToUpper(char)
        }
    }
    fmt.Println("==> " + text)
    return emojis + "\n"
}

func printShareText(shareText string, wordleNum int, tries int, maxTries int) {
    var triesText interface{} = "X"
    if (tries < maxTries) {
        triesText = tries
    }
    fmt.Println("\n=== COPY BELOW TO SHARE WITH FRIENDS ===\n")
    fmt.Printf("Wordle %d %v/%d\n", wordleNum, triesText, maxTries)
    fmt.Println(shareText)
}

func main() {
    const wordLen, maxTries int = 5, 6
    config := getJson()
    reader := bufio.NewReader(os.Stdin)
    answer, wordleNum := getAnswer(config)
    allWords := append(config["answers"], config["others"]...)
    var shareText string
    var solved bool
    var tries int
    fmt.Printf("This is Wordle. Enter the correct %d letter word. You have %d tries\n", wordLen, maxTries)
    for tries < maxTries && !solved {
        word, _ := reader.ReadString('\n')
        word = strings.ToLower(strings.TrimSpace(word))
        if word == answer {
            shareText += colorizedResponse(word, answer)
            solved = true
        } else if contains(allWords, word) {
            shareText += colorizedResponse(word, answer)
            tries++
            fmt.Printf("%d tries left. Guess again\n", maxTries - tries)
        } else if len(word) != wordLen {
            fmt.Printf("Your word has to be %d letters. Guess again\n", wordLen)
        } else {
            fmt.Println("Not a valid word. Guess again")
        }
    }
    if solved {
        fmt.Println("You solved it!")
    } else {
        fmt.Printf("Ouch! You failed. The word was '%s'\n", answer)
    }
    printShareText(shareText, wordleNum, tries + 1, maxTries)
}
