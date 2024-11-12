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
	conn         *sqlx.DB
	transactions map[string]*sql.Tx
}

var pgOnce = sync.Once{}
var connector = PgConnector{conn: nil}

func (PgConnector) Getconnection() *PgConnector {
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

		connector.conn = conn
	})

	return &connector
}

func (pg PgConnector) Get(dest *interface{}, query string, params map[string]interface{}) error {
	return pg.conn.Get(dest, query, params)
}
func (pg PgConnector) Select(dest *[]interface{}, query string, params map[string]interface{}) error {
	return pg.conn.Select(dest, query, params)
}

func (pg PgConnector) Exec(cmd string, params map[string]interface{}) error {
	_, err := pg.conn.Exec(cmd, params)
	return err
}

// Implement the transaction semantics, converts it to a more flexible handle-based manipuliation
func (pg PgConnector) BeginTx() string {
	id := uuid.New().String()
	tx, err := pg.conn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	pg.transactions[id] = tx
	return id
}

func (pg PgConnector) CommitTx(txId string) {

}
