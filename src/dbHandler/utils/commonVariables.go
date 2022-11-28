package black_kiwi_db_utils

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

var ConnPool *pgxpool.Pool
var CatMap *map[black_kiwi_data_structs.Categories]int