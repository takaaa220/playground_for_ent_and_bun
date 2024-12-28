package main

import (
	"fmt"

	"github.com/takaaa220/playgronund_for_ent_and_bun/blog/bun/migrate/migrations"
	"github.com/takaaa220/playgronund_for_ent_and_bun/bunutil"
	"github.com/uptrace/bun/migrate"
)

// see: https://github.com/uptrace/bun/tree/master/example/migrate

func main() {
	migrator := migrate.NewMigrator(bunutil.Connect(), migrations.Migrations)

	if err := bunutil.RunMigration(migrator); err != nil {
		panic(fmt.Errorf("failed to run migration: %w", err))
	}
}
