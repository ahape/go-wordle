package main

import (
    "io/ioutil"
    "encoding/json"
)

func getJson() Config {
    file, _ := ioutil.ReadFile("../words.json")
    var parsed Config
    _ = json.Unmarshal(file, &parsed)
    parsed["allWords"] = append(parsed["answers"], parsed["others"]...)
    return parsed
}
