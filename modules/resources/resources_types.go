package resources

import (
  "gorm.io/gorm"

  "github.com/dee-ex/resource-saver/entities"
)

type (
  Repository interface {
    GetLastID() (uint32, error)
    GetList(offset, limit int) ([]entities.Resource, error)
    GetResourceByID(id uint32) (*entities.Resource, error)
    GetResourceByCode(code string) (*entities.Resource, error)
    CreateResource(resource *entities.Resource) (error)
    UpdateOriginByID(id uint32, origin string) (error)
    UpdateOriginByCode(code, origin string) (error)
    UpdateCodeByID(id uint32, code string) (error)
    DeleteResourceByID(id uint32) (error)
    DeleteResourceByCode(code string) (error)
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
