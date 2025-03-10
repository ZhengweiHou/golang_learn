module entdemo

go 1.24.0

// https://github.com/ZhengweiHou/ent/tree/v0.14.2_hzw
replace entgo.io/ent v0.14.2 => /home/houzw/document/git-rep/golang/ent

// https://github.com/ZhengweiHou/atlas/tree/v0.31.0_hzw
replace ariga.io/atlas v0.31.0 => /home/houzw/document/git-rep/golang/atlas

require (
	//	entgo.io/ent v0.14.2
	github.com/go-sql-driver/mysql v1.9.0
	github.com/ibmdb/go_ibm_db v0.5.2
)

require entgo.io/ent v0.14.2

require (
	ariga.io/atlas v0.31.0 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/ibmruntimes/go-recordio/v2 v2.0.0-20240416213906-ae0ad556db70 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/zclconf/go-cty v1.14.4 // indirect
	github.com/zclconf/go-cty-yaml v1.1.0 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
