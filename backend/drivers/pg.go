/**
* @description Implement a customized postgres connection driver
* @author Yuting Xie
* @date Nov 11, 2024
 */

package drivers

import (
	"database/sql"
	"log"
	"sync"
	"time"
	"wolf/config"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PgConnector struct {
	Conn         *sqlx.DB
	transactions map[string]*sql.Tx
}

var pgOnce = sync.Once{}
var connector = PgConnector{Conn: nil}

func GetPgConnector() *PgConnector {
	pgOnce.Do(func() {
		deployConfig := config.GetDeployConfig()
		conn := sqlx.MustConnect("postgres", deployConfig.Pg.Uri)

		err := conn.Ping()
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

// Implement the transaction semantics, converts it to a more flexible handle-based manipuliation
func (pg PgConnector) BeginTx() string {
	id := uuid.New().String()
	tx, err := pg.Conn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	pg.transactions[id] = tx
	return id
}

func (pg PgConnector) CommitTx(txId string) {

}
