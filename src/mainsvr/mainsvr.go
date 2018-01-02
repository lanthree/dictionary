package main

import "log"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "net/http"
import "encoding/json"
import "time"

type Explanation struct {
	Explanationid int
	Wordid int
	Explanation string
	Tags string
	Sentence string
	BackgroundImg string

	ViewsCounter int
	ThumbupCounter int
	ThumbdownCounter int

	Author string
	Updatetime time.Time
}

type Word struct {
	Wordid int
	Value string
	Explanations []Explanation
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

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/words", http.StatusFound)
		return
	}

	out := &EmptyResult{-1, "Not Found"}
	b, err := json.Marshal(out)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write(b)
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
	rows.Close()

	var exp Explanation
	for idx := range words {
		rows, err := db.Query("SELECT * FROM explanations WHERE wordid=?", words[idx].Wordid)
		for rows.Next() {
			err = rows.Scan(&exp.Explanationid, &exp.Wordid, &exp.Explanation,
					&exp.Tags, &exp.Sentence, &exp.BackgroundImg,
					&exp.ViewsCounter, &exp.ThumbupCounter, &exp.ThumbdownCounter,
					&exp.Author, &exp.Updatetime)
			if err != nil {
				panic(err.Error())
			}
			words[idx].Explanations = append(words[idx].Explanations, exp)
		}
		rows.Close()
	}

	log.Println(words)

	out := &WordsResult{0, "OK", words}
	b, err := json.Marshal(out)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write(b)
}

func main() {

	db, err = sql.Open("mysql", "lanthree:@/dictionary?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//http.Handle("/", http.FileServer(http.Dir("/home/ubuntu/work/img/")))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("/home/ubuntu/work/img/"))))

	http.HandleFunc("/words", WordsHandler)
	http.HandleFunc("/", NotFoundHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
