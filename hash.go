package main 

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "strconv"
    "encoding/json"
    "strings"
)

var kvalue1 map[int] string
var kvalue2 map[int] string
var kvalue3 map[int] string

type Response struct{
    Key  int        `json:"key"`
    Value string    `json:"value"`

}

func main(){
  kvalue1 = make(map[int] string)
  kvalue2 = make(map[int] string)
  kvalue3 = make(map[int] string)

    go func(){
    mux1 := httprouter.New()    
    mux1.PUT("/keys/:id/:value",put)
    mux1.GET("/keys/:id",get)
    mux1.GET("/keys",getall)
    ser := http.Server{
            Addr:        "0.0.0.0:3000",
            Handler: mux1,
    }
    ser.ListenAndServe()
    }()
    

    go func(){
    mux2 := httprouter.New()
    mux2.PUT("/keys/:id/:value",put)
    mux2.GET("/keys/:id",get)
    mux2.GET("/keys",getall)
    ser2 := http.Server{
            Addr:        "0.0.0.0:3001",
            Handler: mux2,
    }
    ser2.ListenAndServe()
    }()

    mux3 := httprouter.New()
    mux3.PUT("/keys/:id/:value",put)
    mux3.GET("/keys/:id",get)
    mux3.GET("/keys",getall)
    ser3 := http.Server{
            Addr:        "0.0.0.0:3002",
            Handler: mux3,
    }
    ser3.ListenAndServe()


}


func put(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
    key := p.ByName("id")
    value := p.ByName("value")
  var port []string
    index, _ := strconv.Atoi(key)
  port = strings.Split(req.Host,":")
  if(port[1]=="3000"){
      kvalue1[index] = value    

  } else if (port[1]=="3001"){
      kvalue2[index] = value 

  } else{
      kvalue3[index] = value  

    }
  
}

func get(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
log.Println("Inside GET!!")
  key := p.ByName("id")
  index, _ := strconv.Atoi(key)
  var port []string
  var response Response
  port = strings.Split(req.Host,":")
  if(port[1]=="3000"){
    response.Key = index
    response.Value = kvalue1[index]   

  } else if (port[1]=="3001"){
    response.Key = index
    response.Value = kvalue2[index] 

  } else{
    response.Key = index
    response.Value = kvalue3[index] 

    }
    payload, err := json.Marshal(response)  
    if err != nil {
         http.Error(rw,"Bad Request" , http.StatusInternalServerError)
        return
    }
    rw.Header().Set("Content-Type", "application/json")
    rw.Write(payload)
}


func getall(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
    var response []Response
    var pair Response
  var port []string
  port = strings.Split(req.Host,":")
  if(port[1]=="3000"){
      for key, value := range kvalue1 {
      pair.Key = key
      pair.Value = value
       response = append(response, pair)
      }    

  } else if (port[1]=="3001"){
      for key, value := range kvalue2 {
      pair.Key = key
      pair.Value = value
       response = append(response, pair)
      } 

  } else{
      for key, value := range kvalue3 {
      pair.Key = key
      pair.Value = value
       response = append(response, pair)
      } 

    }
    
    
    resData, err := json.Marshal(response)  
    if err != nil {
         http.Error(rw,"Bad Request" , http.StatusInternalServerError)
        return
    }
    rw.Header().Set("Content-Type", "application/json")
    rw.Write(resData)
}