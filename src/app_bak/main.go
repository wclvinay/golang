package main

import (
  "encoding/json"
  "net/http"
)

type Profile struct {
  Name    string
  Hobbies []string
}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  //profile := Profile{"Alex", []string{"snowboarding", "programming"}}
   profile :=[]Profile{}
   profile1 :=Profile{
        Name:"Alex",
        Hobbies : []string{"abc","bcd"},
   }
   profile= append(profile,profile1);
   profile2 :=Profile{
        Name:"Job",
        Hobbies : []string{"xabc","vbcd"},
   }
   profile= append(profile,profile2);
  js, err := json.Marshal(profile)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

