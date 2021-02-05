package resources

import (
  "errors"
  "gorm.io/gorm"

  "github.com/dee-ex/resource-saver/entities" 
)

func (repo *MySQL) GetLastID() (uint32, error) {
  var resource entities.Resource
  res := repo.db.Last(&resource)

  // in case table is empty
  if errors.Is(res.Error, gorm.ErrRecordNotFound) {
    return 0, nil
  }

  return resource.ID, res.Error
}

func (repo *MySQL) CreateResource(resource *entities.Resource) (error) {
  res := repo.db.Create(resource)

  return res.Error
}
