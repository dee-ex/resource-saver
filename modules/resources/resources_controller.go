package resources

import (
  "net/http"
  "encoding/json"

  "github.com/dee-ex/resource-saver/utils"
)

func ResourceCreationHandler(w http.ResponseWriter, r *http.Request) {
  // take resource from input
  var input ResourceCreationInput
  err := json.NewDecoder(r.Body).Decode(&input)
  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  // validate input
  if len(input.Resource) == 0 {
    http.Error(w, "No resource found. Make sure to pass your resource properly.", 400)
    return
  }

  // init repository and service
  repo, err := NewMySQLRepository()
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  serv := NewService(repo)

  // create resource from data given
  res, err := serv.CreateResource(input)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }

  // fetch the result
  utils.JSONResponse(w, 200, res)
}

func ResourceRetrievalHandler(w http.ResponseWriter, r *http.Request) {
  // some code here 
}
