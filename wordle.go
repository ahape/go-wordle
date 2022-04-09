package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "time"
)

type Wordle struct {
    config Config
    startDate time.Time
    answer string
    shareText string
    wordleNum int
    wordLen int
    maxTries int
    tries int
    solved bool
}

func NewWordle() *Wordle {
    w := &Wordle{
        maxTries: 6,
        wordLen: 5,
        config: getJson(),
        startDate: time.Date(2021, time.June, 19, 0, 0, 0, 0, time.Local),
    }
    w.loadAnswer()
    return w
}

func (w *Wordle) loadAnswer() {
    year, month, day := time.Now().Date()
    todayDate := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
    w.wordleNum = int(todayDate.Sub(w.startDate) / (24 * time.Hour))
    w.answer = w.config["answers"][w.wordleNum]
}

func (w *Wordle) colorizedResponse(word string) {
    var normalText, normalEmoji string = "\033[0m", "\u2B1C"
    var greenText, greenEmoji string = "\033[32m", "\U0001f7e9"
    var yellowText, yellowEmoji string = "\033[33m", "\U0001f7e8"
    var text string
    for i, char := range strings.Split(word, "") {
        if char == string(w.answer[i]) {
            w.shareText += greenEmoji
            text += greenText + strings.ToUpper(char) + normalText
        } else if contains(strings.Split(w.answer, ""), char) {
            w.shareText += yellowEmoji
            text += yellowText + strings.ToUpper(char) + normalText
        } else {
            w.shareText += normalEmoji
            text += strings.ToUpper(char)
        }
    }
    w.shareText += "\n"
    fmt.Println("==> " + text)
}

func (w Wordle) printShareText() {
    tries := w.tries + 1
    var triesText interface{} = "X"
    if (tries <= w.maxTries) {
        triesText = tries
    }
    fmt.Println("\n=== COPY BELOW TO SHARE WITH FRIENDS ===\n")
    fmt.Printf("Wordle %d %v/%d\n", w.wordleNum, triesText, w.maxTries)
    fmt.Println(w.shareText)
}

func (w *Wordle) Start() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("This is Wordle. Enter the correct %d letter word. You have %d tries\n", w.wordLen, w.maxTries)
    for w.tries < w.maxTries && !w.solved {
        word, _ := reader.ReadString('\n')
        word = strings.ToLower(strings.TrimSpace(word))
        if word == w.answer {
            w.colorizedResponse(word)
            w.solved = true
        } else if contains(w.config["allWords"], word) {
            w.colorizedResponse(word)
            w.tries++
            fmt.Printf("%d tries left. Guess again\n", w.maxTries - w.tries)
        } else if len(word) != w.wordLen {
            fmt.Printf("Your word has to be %d letters. Guess again\n", w.wordLen)
        } else {
            fmt.Println("Not a valid word. Guess again")
        }
    }
    if w.solved {
        fmt.Println("You solved it!")
    } else {
        fmt.Printf("Ouch! You failed. The word was '%s'\n", w.answer)
    }
    w.printShareText()
}
