package main

import (
	"fmt"
	R "github.com/fbaube/repo"
	RS "github.com/fbaube/reposqlite"
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
	rs, e := RS.NewRepoAtPath("mmmc.db")
	if e != nil {
		panic(e)
	}
	i1, _ = rs.(R.DBImplementation)
	fmt.Printf("R.DBImplementation: %T \n", i1)
	i2, _ = R.DBImplementation(rs)
	fmt.Printf("R.DBImplementation: %T \n", i2)
	_, _ = rs.(R.DBEntity)
	_, _ = rs.(R.DBBackups)
	_, _ = rs.(R.SessionLifecycle)
	_, _ = rs.(R.StatementBuilder)
	_, _ = rs.(R.QueryRunner)
}
