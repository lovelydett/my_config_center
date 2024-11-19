/**
* @description Implement a customized postgres connection driver
* @author Yuting Xie
* @date Nov 11, 2024
 */

package drivers

import (
	"context"
	"time"
	"wolf/config"

	"github.com/jmoiron/sqlx"
)

var Pg *sqlx.DB

func init() {
	deployConfig := config.GetDeployConfig()
	Pg = sqlx.MustConnect("postgres", deployConfig.Pg.Uri)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := Pg.PingContext(ctx)
	if err != nil {
		if err == context.Canceled || err == context.DeadlineExceeded {
			panic("Ping Postgres time out")
		}
		panic("Failed to ping postgres")
	}

	Pg.SetMaxOpenConns(25)
	Pg.SetMaxIdleConns(25)
	Pg.SetConnMaxLifetime(5 * time.Minute) // Refresh connections periodically
}
