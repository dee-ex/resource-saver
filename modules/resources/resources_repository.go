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

func (repo *MySQL) GetList(offset, limit int) ([]entities.Resource, error) {
	var resources []entities.Resource
	res := repo.db.Offset(offset).Limit(limit).Find(&resources)

	return resources, res.Error
}

func (repo *MySQL) GetResourceByID(id uint32) (*entities.Resource, error) {
	var resource entities.Resource
	res := repo.db.Where("id = ?", id).Find(&resource)

	return &resource, res.Error
}

func (repo *MySQL) GetResourceByCode(code string) (*entities.Resource, error) {
	var resource entities.Resource
	res := repo.db.Where("code = ?", code).Find(&resource)

	return &resource, res.Error
}

func (repo *MySQL) CreateResource(resource *entities.Resource) (error) {
  res := repo.db.Create(resource)

  return res.Error
}

func (repo *MySQL) UpdateOriginByID(id uint32, origin string) (error) {
	res := repo.db.Model(&entities.Resource{}).Where("id = ?", id).Update("origin", origin)

	return res.Error
}

func (repo *MySQL) UpdateOriginByCode(code, origin string) (error) {
	res := repo.db.Model(&entities.Resource{}).Where("code = ?", code).Update("origin", origin)

	return res.Error
}

func (repo *MySQL) UpdateCodeByID(id uint32, code string) (error) {
	res := repo.db.Model(&entities.Resource{}).Where("id = ?", id).Update("code", code)

	return res.Error
}

func (repo *MySQL) DeleteResourceByID(id uint32) (error) {
	res := repo.db.Model(&entities.Resource{}).Where("id = ?", id).Update("deleted", true)

	return res.Error
}

func (repo *MySQL) DeleteResourceByCode(code string) (error) {
	res := repo.db.Model(&entities.Resource{}).Where("code = ?", code).Update("deleted", true)

	return res.Error
}
