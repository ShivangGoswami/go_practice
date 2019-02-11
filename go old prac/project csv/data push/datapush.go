package main

import (
	"bufio"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	_ "github.com/go-goracle/goracle"
)

func isValidDelimiter(d *string) error {
	if len(*d) == 0 {
		return errors.New("Please enter a delimiter")
	} else if len(*d) != 1 {
		return errors.New("delimiter must be of single character please use tr command and then pipe the data")
	} else if unicode.IsLetter(([]rune(*d))[0]) || unicode.IsDigit(([]rune(*d))[0]) {
		return errors.New("keeping delimiter as a character or a digit may lead to inconsistency")
	}
	return nil
}

func schemaDatahandler(s []string) {
	for _, txt := range s {
		if txt == "" || unicode.IsDigit(rune(txt[0])) {
			log.Fatal("Inconsistent Format:Column name badly defined")
		}
	}
}
func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func stringbuilder(dat *os.File, schema []string, sam []string) string {
	sam = rowSplitCorrecter(sam)
	reservedset := []string{"ACCESS", "ADD", "ALL", "ALTER", "AND", "ANY",
		"AS", "ASC", "AUDIT", "BETWEEN", "BY", "CHAR", "CHECK", "CLUSTER", "COLUMN",
		"COMMENT", "COMPRESS", "CONNECT", "CREATE", "CURRENT", "DATE", "DECIMAL",
		"DEFAULT", "DELETE", "DESC", "DISTINCT", "DROP", "ELSE", "EXCLUSIVE", "EXISTS",
		"FILE", "FLOAT", "FOR", "FROM", "GRANT", "GROUP", "HAVING", "IDENTIFIED",
		"IMMEDIATE", "IN", "INCREMENT", "INDEX", "INITIAL", "INSERT", "INTEGER",
		"INTERSECT", "INTO", "IS", "LEVEL", "LIKE", "LOCK", "LONG", "MAXEXTENTS",
		"MINUS", "MLSLABEL", "MODE", "MODIFY", "NOAUDIT", "NOCOMPRESS", "NOT",
		"NOWAIT", "NULL", "NUMBER", "OF", "OFFLINE", "ON", "ONLINE", "OPTION",
		"OR", "ORDER", "PCTFREE", "PRIOR", "PRIVILEGES", "PUBLIC", "RAW",
		"RENAME", "RESOURCE", "REVOKE", "ROW", "ROWID", "ROWNUM",
		"ROWS", "SELECT", "SESSION", "SET", "SHARE", "SIZE", "SMALLINT",
		"START", "SUCCESSFUL", "SYNONYM", "SYSDATE", "TABLE", "THEN", "TO",
		"TRIGGER", "UID", "UNION", "UNIQUE", "UPDATE", "USER", "VALIDATE", "VALUES",
		"VARCHAR", "VARCHAR2", "VIEW", "WHENEVER", "WHERE", "WITH"}
	reservedWordSet := strings.Join(reservedset, ",")
	name := dat.Name()
	str := "Create table " + name[0:len(name)-len(filepath.Ext(name))] + "("
	if len(schema) != len(sam) {
		log.Fatal("Inconsistent Format:Headers mismatch with data")
	}
	for i, datatype := range schema {
		if strings.Contains(reservedWordSet, strings.ToUpper(datatype)) {
			datatype += "_appender"
		}
		datatype = strings.Replace(datatype, " ", "_", -1)
		if _, err := strconv.ParseInt(sam[i], 10, 64); err == nil { //64 due to our system is 64 bit
			str += datatype + " " + "int" + ","
		} else if _, err := strconv.ParseFloat(sam[i], 10); err == nil {
			str += datatype + " " + "number" + ","
		} else {
			str += datatype + " " + "varchar2(500)" + ","
		}
	}
	str = strings.TrimRight(str, ",")
	str += " )"
	return str
}
func schemaErrorHandler(schema []string, delimit *string) error {
	delimiterlist := []string{",", ";", "'", "|", "/", "\\"}
	for _, datatype := range schema {
		datatype = strings.Replace(datatype, " ", "_", -1)
		for _, del := range delimiterlist {
			if strings.Contains(datatype, del) {
				err := errors.New("Inconsistent format:Other Delimiters present")
				return err
			}
		}
	}
	return nil
}
func queryBuilder(dat *os.File, row []string) string {
	name := dat.Name()
	str := "Insert into " + name[0:len(name)-len(filepath.Ext(name))] + " values("
	for _, val := range row {
		val = strings.Replace(val, "NaN", "NULL", -1)
		val = strings.Replace(val, "'", "''", -1)
		val = strings.Replace(val, "$", "", -1)
		if _, err := strconv.ParseInt(val, 10, 64); err == nil { //64 due to our system is 64 bit
			str += val + ","
		} else if _, err := strconv.ParseFloat(val, 10); err == nil {
			str += val + ","
		} else if strings.Contains(val, "NULL") {
			str += val + ","
		} else {
			val = strings.Replace(val, "&", "and", -1)
			str += "'" + val + "',"
		}
	}
	str = strings.TrimRight(str, ",")
	str += ")"
	return str
}
func rowSplitCorrecter(row []string) []string {
	newrow := []string{}
	i := 0
	temp := ""
	for i < len(row) {
		temp = ""
		if strings.Contains(row[i], "\"\"") {
			row[i] = strings.Replace(row[i], "\"\"", "", -1)
			if strings.Count(row[i], "\"") == 2 {
				row[i] = strings.Replace(row[i], "\"", "", -1)
				newrow = append(newrow, row[i])
				i++
				continue
			}
		}
		if !strings.Contains(row[i], "\"") {
			newrow = append(newrow, row[i])
		} else {
			if strings.Count(row[i], "\"") == 2 {
				row[i] = strings.Replace(row[i], "\"", "", -1)
				newrow = append(newrow, row[i])
				i++
				continue
			}
			temp += row[i] + ","
			i++
			for !strings.Contains(row[i], "\"") {
				temp += row[i] + ","
				i++
			}
			temp += row[i] + ","
			i++
			temp = strings.TrimRight(temp, ",")
			temp = strings.Replace(temp, "\"", "", -1)
			newrow = append(newrow, temp)
			i--
		}
		i++
	}
	return newrow
}
func insertRow(dat *os.File, db *sql.DB, row []string) {
	row = rowSplitCorrecter(row)
	if len(row) != 13 {
		log.Println("Ignoring this row(Value mismatch):", row)
		return
	}
	tx, err := db.Begin()
	checkerr(err)
	defer tx.Rollback()
	fmt.Println(queryBuilder(dat, row))
	stmt, err := tx.Prepare(queryBuilder(dat, row))
	checkerr(err)
	_, err = stmt.Exec()
	checkerr(err)
	stmt.Close()
	err = tx.Commit()
	checkerr(err)
}
func counter(rowCount *int) {
	log.Println(*rowCount, "Rows Processed")
	*rowCount++
}
func rowCountFetcher(dat *os.File, db *sql.DB) int {
	var count int
	tname := dat.Name()
	tname = tname[0 : len(tname)-len(filepath.Ext(tname))]
	db.QueryRow("select count(*) from " + tname).Scan(&count)
	return count
}
func main() {
	delimit := flag.String("delimiter", " ", " a delimiter to distinguish data in the flat file")
	flag.Parse()
	err := isValidDelimiter(delimit)
	checkerr(err)
	dat, err := os.Open("googleplaystore.csv")
	checkerr(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	scanner.Scan()
	schema := strings.Split(scanner.Text(), *delimit)
	err = schemaErrorHandler(schema, delimit)
	checkerr(err)
	schemaDatahandler(schema)
	scanner.Scan()
	db, err := sql.Open("goracle", "system/1001289@localhost:1521/xe")
	defer db.Close()
	checkerr(err)
	_, err = db.Exec(stringbuilder(dat, schema, strings.Split(scanner.Text(), *delimit)))
	fmt.Println("Schema Builder Query:", stringbuilder(dat, schema, strings.Split(scanner.Text(), *delimit)))
	checkerr(err)
	log.Println("Schema Successfully created!!!!!!!")
	insertRow(dat, db, strings.Split(scanner.Text(), *delimit))
	rowCount := 1
	counter(&rowCount)
	for scanner.Scan() {
		insertRow(dat, db, strings.Split(scanner.Text(), *delimit))
		counter(&rowCount)
	}
	rowCountsucc := rowCountFetcher(dat, db)
	log.Println("Sucessfully inserted:", rowCountsucc, "Ignored:", rowCount-rowCountsucc-1)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
