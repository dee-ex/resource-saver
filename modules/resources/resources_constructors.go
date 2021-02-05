package resources

import (
  "github.com/dee-ex/resource-saver/infrastructures"
)

func NewMySQLRepository() (*MySQL, error) {
  db, err := infrastructures.NewMySQLSession()
  if err != nil {
    return nil, err
  }
  return &MySQL{db: db}, nil
}

func NewService(repo Repository) (*Service) {
  return &Service{repo: repo}
}
