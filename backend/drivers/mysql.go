// Implement the MySQL connector
package drivers

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"wolf/config"
)

type _MySQLConnector struct {
	Conn *sql.DB
}

var once = sync.Once{}

func GetMySQLConnector() *_MySQLConnector {
	connector := &_MySQLConnector{Conn: nil}
	once.Do(func() {
		deployConfig := config.GetDeployConfig()
		dataSourceName := deployConfig.MySQL.Username + ":" + deployConfig.MySQL.Password + "@tcp(" + deployConfig.MySQL.Host + ":" + deployConfig.MySQL.Port + ")/" + deployConfig.MySQL.Database
		conn, err := sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Fatal(err)
		}
		connector.Conn = conn
	})
	return connector
}

func (c *_MySQLConnector) Exec(query string, args ...interface{}) (*sql.Rows, error) {
	return c.Conn.Query(query, args...)
}
