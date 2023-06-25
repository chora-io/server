module github.com/choraio/server/api

go 1.20

require (
	github.com/choraio/server/db v0.0.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/piprate/json-gold v0.5.0
	github.com/rs/zerolog v1.29.1
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.15.0
)

require (
	github.com/cosmos/btcutil v1.0.5 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/pressly/goose/v3 v3.10.0 // indirect
	github.com/simukti/sqldb-logger v0.0.0-20220521163925-faf2f2be0eb6 // indirect
	github.com/simukti/sqldb-logger/logadapter/zerologadapter v0.0.0-20220521163925-faf2f2be0eb6 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/choraio/server/db => ../db
)
