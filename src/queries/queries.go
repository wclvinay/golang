package queries
import (
  "fmt"
  "net/http"
  "db"
  "encoding/json"
  "github.com/gorilla/mux"
  "database/sql"
)
type ResponseData struct {
	Id     int  `json:"id"`
    ParentId   int `json:"parent_id"`
    //Type   string `json:"type"`
    Questions   string `json:"questions"`
    Answer   string `json:"answer"`
    Title    string    `json:"title"`
    Image    string    `json:"image"`
}
var con *sql.DB
func init() {
    con = db.CreateCon()
}
func GetQueries(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query_type :=params["type"]
	fmt.Println(query_type)
	fmt.Println("Go MySQL Tutorial")
	results, err := con.Query("SELECT id,parent_id,questions,answer,title,image from queries where type=?",query_type)
	fmt.Println(results)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    datas := []ResponseData{}
    for results.Next() {
		var id int
        var parent_id int
        //var type  string
        var title string
        var questions string
        var answer string
        var image string
        err = results.Scan(&id, &parent_id, &questions,&answer,&title,&image)
        datas = append(datas, ResponseData{
			Id : id,
			ParentId: parent_id,
			Questions : questions,
			Answer : answer,
			Title: title,
			Image: image,		
		})
		
    }
    js,err := json.Marshal(datas)
	fmt.Println(string(js))
	if err != nil {
   	   http.Error(w, err.Error(), http.StatusInternalServerError)
           return
    }
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
