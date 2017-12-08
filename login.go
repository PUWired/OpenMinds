package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

var username int64

func login(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    username, _ = strconv.ParseInt(r.Form["username"][0], 10, 64)
    password := r.Form["password"][0]
    role := r.Form["role"][0]
    username = uHash(username)
    password = pHash(password)
    fmt.Println(username, password, role)
    http.Redirect(w, r, "presicoun/messenger.html", http.StatusMovedPermanently)
}

func pHash(word string) string {
	l := len(word)
	var hash string
	for l != 0 {
		l--
		hash += string(pow(13, int64(word[l]))%251)
	}
	return hash
}

func uHash(input int64) int64 {
	var ans int64 = 0
	for input != 0 {
		ans = 16 * ans + pow(91, input%16) % 51
		input /= 16
	}
	return ans
}

func pow(base int64, exp int64) int64 {
	var power int64 = 1
	for exp > 0 {
		power *= base
		exp--
	}
	return power
}

func main() {
    go http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
