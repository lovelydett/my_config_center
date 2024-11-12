/**
* @description Implement a customized postgres connection driver
* @author Yuting Xie
* @date Nov 11, 2024
 */

package drivers

import (
	"log"
	"sync"
	"time"
	"wolf/config"

	"github.com/jmoiron/sqlx"
)

type PgConnector struct {
	Conn *sqlx.DB
}

var pgOnce = sync.Once{}
var connector = PgConnector{Conn: nil}

func (PgConnector) GetConnection() *PgConnector {
	pgOnce.Do(func() {
		deployConfig := config.GetDeployConfig()
		conn, err := sqlx.Open("postgres", deployConfig.Pg.Uri)
		if err != nil {
			log.Fatal(err)
		}

		conn.SetMaxOpenConns(25)
		conn.SetMaxIdleConns(25)
		conn.SetConnMaxLifetime(5 * time.Minute) // Refresh connections periodically

		connector.Conn = conn
	})

	return &connector
}
