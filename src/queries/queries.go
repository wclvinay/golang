package queries
import (
  "fmt"
  "net/http"
  "db"
  "cache"
  "encoding/json"
  "github.com/gorilla/mux"
  
)
type ResponseData struct{
  FLAG int `json:"flag"`
  Queries []QueriesData `json:"queries"`
}
type QueriesData struct {
	  Id     int  `json:"id"`
    ParentId   int `json:"parent_id"`
    Questions   string `json:"questions"`
    Answer   string `json:"answer"`
    Title    string    `json:"title"`
    Image    string    `json:"image"`
}
func init() {
}
func GetQueries(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query_type :=params["type"]
	fmt.Println("Query Type:",query_type)
	fmt.Println("Go MySQL Tutorial")
  queryValues := r.URL.Query()
  parentId := queryValues.Get("parent_id")
  if len(parentId) <= 0 {
    parentId = "0" 
  }
  var redisKey string
  redisKey = "queries_"+query_type+"_"+parentId
  fmt.Println(redisKey)
  redisData, err := cache.Cacheobj.Get(redisKey).Result()
  if len(redisData) > 0 {
     fmt.Println("RedisData:", redisData)
     raw := ResponseData{}
     json.Unmarshal([]byte(redisData),&raw)
      js,_ := json.Marshal(raw)
      w.Header().Set("Content-Type", "application/json")
      w.Write(js)
      return
  }
	results, err := db.Connection.Query("SELECT id,parent_id,questions,answer,title,image from queries where type=? and parent_id=?",query_type,parentId)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    datas := []QueriesData{}
    for results.Next() {
	    var id int
      var parent_id int
      var title string
      var questions string
      var answer string
      var image string
      err = results.Scan(&id, &parent_id, &questions,&answer,&title,&image)
      datas = append(datas, QueriesData{
  			Id : id,
  			ParentId: parent_id,
  			Questions : questions,
  			Answer : answer,
  			Title: title,
  			Image: image,		
	    })
    }
    finalData := ResponseData{1,datas}
    js,err := json.Marshal(finalData)
  	fmt.Println(string(js))
    cache.Cacheobj.Set(redisKey, js, 0).Err()
  	if err != nil {
     	   http.Error(w, err.Error(), http.StatusInternalServerError)
             return
      }
  	w.Header().Set("Content-Type", "application/json")
  	w.Write(js)
}
