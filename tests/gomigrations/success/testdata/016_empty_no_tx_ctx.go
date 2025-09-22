package gomigrations

import (
	"github.com/fmarmol/goose/v3"
)

func init() {
	goose.AddMigrationNoTxContext(nil, nil)
}
