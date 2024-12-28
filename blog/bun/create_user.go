package blob_bun

import (
	"fmt"

	"github.com/takaaa220/playgronund_for_ent_and_bun/bunutil"
)

func CreateUser() error {
	db := bunutil.Connect()
	fmt.Println(db)

	return nil
}
