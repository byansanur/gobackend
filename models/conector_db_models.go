package models

import (
	"../config_db"
)

var (
	db  = config_db.DBInit()
	idb = config_db.InDB{DB: db}
	// idb = inDB

)
