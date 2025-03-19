module gormdemo

go 1.24.0

replace gorm.io/gorm v1.25.12 => /home/houzw/document/git-rep/golang/gorm

replace gorm.io/driver/ibmdb v1.0.0 => /home/houzw/document/git-rep/HOUZW/ibmdb

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
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/ibmruntimes/go-recordio/v2 v2.0.0-20240416213906-ae0ad556db70 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/microsoft/go-mssqldb v1.7.2 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
