module gormdemo

//go 1.24.0
go 1.23.6

replace gorm.io/gorm v1.25.12 => /home/houzw/document/git-rep/golang/gorm

replace gorm.io/driver/ibmdb v1.0.0 => /home/houzw/document/git-rep/HOUZW/golang/ibmdb

//replace gorm.io/driver/ibmdb v1.0.0 => github.com/ZhengweiHou/ibmdb v0.0.0-20250312083856-50baf4c8d628

//require gorm.io/gorm v1.25.12

require (
	github.com/ibmdb/go_ibm_db v0.5.2
	gorm.io/driver/ibmdb v1.0.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/driver/postgres v1.5.11
	gorm.io/driver/sqlite v1.5.7
	gorm.io/driver/sqlserver v1.5.4
	gorm.io/gorm v1.25.12
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.0 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/ibmruntimes/go-recordio/v2 v2.0.0-20241213170836-956c90c77e2f // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.2 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	github.com/microsoft/go-mssqldb v1.8.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)
