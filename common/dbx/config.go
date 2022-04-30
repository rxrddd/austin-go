package dbx

type DbConfig struct {
	Dsn          string
	MaxIdleConns int //空闲连接池中连接的最大数量
	MaxOpenConns int //打开数据库连接的最大数量
}
