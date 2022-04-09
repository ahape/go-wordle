package main

func contains(arr []string, elem string) bool {
    for _, e := range arr {
        if e == elem {
            return true
        }
    }
    return false
}
