package resources

import (
  "time"

  "github.com/dee-ex/resource-saver/entities"
)

func (serv *Service) CreateResource(data ResourceCreationInput) (*entities.Resource, error) {
  // get last/previous id to generate code
  last_id, err := serv.repo.GetLastID()
  if err != nil {
    return nil, err
  }

  // previous id is the key to generate current code
  code := CodeGenerator(last_id)

  // prepare addtional fields
  current_time := time.Now()

  // make an instance
  resource := entities.Resource{ID: 0,
                                Code: code,
                                Active: true,
                                Deleted: false,
                                Origin: data.Resource,
                                LastAccess: &current_time,
                                DateCreated: &current_time,
                                }

  // save resource to database
  err = serv.repo.CreateResource(&resource)
  if err != nil {
    return nil, err
  }

  // if no error
  return &resource, nil
}
