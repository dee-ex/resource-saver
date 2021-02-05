package resources

import (
  "gorm.io/gorm"

  "github.com/dee-ex/resource-saver/entities"
)

type (
  Repository interface {
    GetLastID() (uint32, error)
    CreateResource(resource *entities.Resource) (error)
  }

  MySQL struct {
    db *gorm.DB
  }

  Service struct {
    repo Repository
  }

  ResourceCreationInput struct {
    Resource string
  }
)
