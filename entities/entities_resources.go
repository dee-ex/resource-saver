package entities

import (
  "time"
)

type (
  Resource struct {
    ID uint32              // id primary key
    Origin string          // origin resource
    Code string            // code mapping to retrieve resource
    DateCreated *time.Time  // date created
    LastAccess *time.Time   // last time access this resource
    Active bool            // base on last time access
    Deleted bool           // for cleaning purposes
  }
)
