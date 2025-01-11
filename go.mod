module github.com/XDwanj/go-gsgm

go 1.22.0

toolchain go1.23.1

// cli
require github.com/spf13/cobra v1.8.1

require github.com/mitchellh/go-homedir v1.1.0

// util
require github.com/bwmarrin/snowflake v0.3.0

require github.com/duke-git/lancet/v2 v2.3.4

require golang.org/x/image v0.23.0

require gopkg.in/yaml.v3 v3.0.1

require github.com/jedib0t/go-pretty/v6 v6.6.5

// db
require github.com/jmoiron/sqlx v1.4.0

require github.com/mattn/go-sqlite3 v1.14.24

require (
	github.com/qustavo/sqlhooks/v2 v2.1.0
	gorm.io/driver/sqlite v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
)

// indirect
require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/exp v0.0.0-20250106191152-7588d65b2ba8 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
