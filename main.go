package main

import (
	"fmt"
	R "github.com/fbaube/repo"
	"github.com/fbaube/repo/sqlite"
	_ "github.com/fbaube/sqlite3"
)

/* REF
type SimpleRepo interface {
	DBImplementation // "sqlite"
	DBEntity         // path, open, isOpen, etc.
	DBBackups        // cp, mv, restoreFrom
	SessionLifecycle // open, close, etc.
	StatementBuilder // using struct RU.QuerySpec
	QueryRunner      // return 0,1,N rows
}
*/

func main() {
	panic("oops")
	sr, e := sqlite.NewRepoAtPath("mmmc.db")
	if e != nil {
		panic(e)
	}
	i1, _ = sr.(R.DBImplementation)
	fmt.Printf("R.DBImplementation: %T \n", i1)
	i2, _ = R.DBImplementation(sr)
	fmt.Printf("R.DBImplementation: %T \n", i2)
	_, _ = sr.(R.DBEntity)
	_, _ = sr.(R.DBBackups)
	_, _ = sr.(R.SessionLifecycle)
	_, _ = sr.(R.StatementBuilder)
	_, _ = sr.(R.QueryRunner)
}
