package sqlite

import (
	"fmt"
	D "github.com/fbaube/dsmnd"
	// FU "github.com/fbaube/fileutils"
	RU "github.com/fbaube/repoutils"
	S "strings"
	// "time"
)

/* REF
https://www.sqlite.org/lang_createtable.html
Create Table [ If Not Exists ] [schemaname.]tablename
1) as select-statement
2) ( columndef,+ tableconstraint,* )
[ table-options ] ;
*/

/* REF
TABL Fundatype = "tabl" // This datum describes a table !
INTG = "1234" // "INT"   // SQLITE_INTEGER 1
FLOT = "1.0f" // "FLOAT" // SQLITE_FLOAT   2
TEXT = "AaZz" // "TEXT"  // SQLITE_TEXT    3
BLOB = "blob" // "BLOB"  // SQLITE_BLOB    4
NULL = "null" // "NULL"  // SQLITE_NULL    5
PKEY = "pkey" // PRIMARY KEY (SQLite "INTEGER")
FKEY = "fkey" // FOREIGN KEY (SQLite "INTEGER")
LIST = "list" // table type, one per enumeration ??
OTHR = "othr"
*/

/*
	FKEY's

idx_inbatch integer not null references inbatch,
foreign key(idx_inbatch) references inbatch(idx_inbatch)
*/
func (pSR *SqliteRepo) BuildCreateTableStmt(pTD *RU.TableDescriptor) (string, error) {
	var sb S.Builder
	sb.WriteString(fmt.Sprintf("CREATE TABLE %s(\n", pTD.Name))
	fmt.Printf("GenCreTblStmt: ")
	for _, pCS := range pTD.ColumnSpecs {
		cnm := pCS.StorName // column name
		fdt := pCS.Fundatype
		fmt.Printf("%s:%s, ", cnm, fdt)
	}
	fmt.Printf("\n")

	// PRIMARY KEY IS ASSUMED
	// idx_tbl integer not null primary key autoincrement,
	sb.WriteString(pTD.IDName +
		" integer not null primary key autoincrement, " +
		"-- NOTE: integer, not int. \n")

	for _, pCS := range pTD.ColumnSpecs {
		cnm := pCS.StorName // column name
		fmt.Println("Creating column: %s \n", pCS.String())
		switch pCS.Fundatype {
		case D.PKEY:
			panic("DUPE PRIMARY KEY")
		case D.FKEY:
			idxnam := pCS.StorName
			tblnam := pCS.DispName
			sb.WriteString(cnm + fmt.Sprintf(
				" foreign key(%s) references %s(%s),\n",
				idxnam, tblnam, idxnam))
		case D.TEXT:
			sb.WriteString(cnm + " text not null,\n")
		case D.INTG:
			// filect int not null check (filect >= 0) default 0
			sb.WriteString(cnm + " int not null,\n")
		/* Unimplem'd:
		case D.FLOT:
		case D.BLOB:
		case D.NULL:
		case D.LIST:
		case D.OTHR:
		*/
		default:
			panic(pCS.Fundatype)
		}
	}
	// trim off final ",\n"
	ss := sb.String()
	sb2 := ss[0:len(ss)-2] + "\n) STRICT;"
	return sb2, nil
}

/*
CREATE TABLE contentity( -- STRICT!
idx_inbatch integer not null references inbatch,
relfp text not null check (typeof(relfp) == 'text'),
absfp text not null check (typeof(absfp) == 'text'),
t_cre text not null check (typeof(t_cre) == 'text'),
t_imp text not null check (typeof(t_imp) == 'text'),
t_edt text not null check (typeof(t_edt) == 'text'),
descr text not null check (typeof(descr) == 'text'),
mimetype text not null check (typeof(mimetype) == 'text'),
mtype     text not null check (typeof(mtype)     == 'text'),
xmlcontype text not null check (typeof(xmlcontype) == 'text'),
ditaflavor text not null check (typeof(ditaflavor) == 'text'),
ditacontype text not null check (typeof(ditacontype) == 'text'),
foreign key(idx_inbatch) references inbatch(idx_inbatch)
);
*/
