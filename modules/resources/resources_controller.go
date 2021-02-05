package resources

import (
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
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
  // take out code value
  vars := mux.Vars(r)
  code := vars["code"]

  // init repository and service
  repo, err := NewMySQLRepository()
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  serv := NewService(repo)

  // get resource by code
  res, err := serv.GetResourceByCode(code)

  // offset id in database is 1
  // 0 id means there is no row matches input code
  if res.ID == 0 {
    http.Error(w, "Code not found. Please check your code again.", 404)
    return
  }

  // fetch the result
  utils.JSONResponse(w, 200, res)
}
