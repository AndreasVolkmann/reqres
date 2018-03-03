package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type request struct {
	id   int
	user string
	note string
	date string
}



func main() {
	path := "D:\\Dev\\go\\src\\github.com\\AndreasVolkmann\\reqres\\request.csv"
	lines, err := buffered(path)
	check(err)
	fmt.Println("Requests:")
	requests := transformLines(lines)
	for _, v := range requests {
		fmt.Println(v)
	}

	insertRequests(requests)
}

func insertRequests(requests []request) {
	db, err := sql.Open("mysql", "test:test@/reqres")
	check(err)
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO reqres.requests VALUES (?, ?, ?, ?)")
	check(err)
	defer stmtIns.Close()
	for _, req := range requests {
		_, err = stmtIns.Exec(req.id, req.user, req.note, req.date)
		check(err)
	}
}

func transformLines(lines []string) []request {
	return Map(lines, func(line string) request {
		return transformLineToRequest(line)
	})
}

func transformLineToRequest(line string) request {
	split := strings.Split(line, ",")
	id, err := strconv.Atoi(split[0])
	check(err)
	return request{id, split[1], split[2], split[3]}
}

func buffered(path string) ([]string, error) {
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	lines = append(lines[:0], lines[1:]...) // drop header
	return lines, sc.Err()
}

func Map(vs []string, f func(string) request) []request {
	vsm := make([]request, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}