package main

import (
	"backend/ent/migrate"
	"backend/utils"
	"context"
	"fmt"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataSourceName := fmt.Sprintf("mysql://%s:%s@%s:%s/%s?charset=utf8&parseTime=True",
		utils.GetEnv("DATABASE_USER", "mysqluser"),
		utils.GetEnv("DATABASE_PASSWORD", "mysqlpass"),
		utils.GetEnv("DATABASE_HOST", "database"),
		utils.GetEnv("DATABASE_PORT", "3306"),
		utils.GetEnv("DATABASE_NAME", "entdemo"),
	)
    fmt.Println(dataSourceName)

	ctx := context.Background()

	dir, err := atlas.NewLocalDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeInspect),
		schema.WithDialect(dialect.MySQL),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	err = migrate.NamedDiff(ctx, dataSourceName, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
