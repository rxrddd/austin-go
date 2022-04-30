package idgen

import (
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
)

var idWorker *snowflake.Snowflake

func init() {
	idWorker, _ = snowflake.NewSnowflake(int64(0), int64(0))
}

func NextID() int64 {
	return idWorker.NextVal()
}
