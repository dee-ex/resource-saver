package infrastructures

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

func OpenMySQLSession(username, password, host, port, dbname string) (*gorm.DB, error) {
  dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?parseTime=true"

  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }
  mysql_db, err := db.DB()
  if err != nil {
    return nil, err
  }
  mysql_db.SetMaxOpenConns(10)
  mysql_db.SetConnMaxLifetime(15)

  return db, nil
}

func NewMySQLSession() (*gorm.DB, error) {
  return OpenMySQLSession("root", "", "localhost", "3306", "resource_saver")
}
