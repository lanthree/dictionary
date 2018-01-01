package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "net/http"
import "encoding/json"

type Word struct {
	Wordid int
	Value string
}

type WordsResult struct {
	Ret int
	Reason string
	Words []Word
}

type EmptyResult struct {
	Ret int
	Reason string
}

var (
	db *sql.DB
	err error
)

func OutputJson(w http.ResponseWriter, ret int, reason string) {

	out := &EmptyResult{ret, reason}
	b, err := json.Marshal(out)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)

	w.Write(b)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/words", http.StatusFound)
		return
	}

	OutputJson(w, 0, "Not Found")
	return
}

func WordsHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT * FROM words")
	if err != nil {
		panic(err.Error())
	}

	var words []Word
	var word Word

	for rows.Next() {
		err = rows.Scan(&word.Wordid, &word.Value)
		if err != nil {
			panic(err.Error())
		}
		words = append(words, word)
	}

	out := &WordsResult{0, "OK", words}
	b, err := json.Marshal(out)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(b)
}

func main() {

	db, err = sql.Open("mysql", "lanthree:@/dictionary")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//http.Handle("/", http.FileServer(http.Dir("/home/ubuntu/work/img/")))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("/home/ubuntu/work/img/"))))

	http.HandleFunc("/words", WordsHandler)
	http.HandleFunc("/", NotFoundHandler)

	http.ListenAndServe("0.0.0.0:80", nil)
}
