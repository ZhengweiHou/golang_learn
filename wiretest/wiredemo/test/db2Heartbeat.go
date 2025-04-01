package test

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"

	_ "github.com/ibmdb/go_ibm_db"
)

var (
	help   bool
	url    string
	port   string
	dbname string
	con    string
)

// stmtCache 缓存预处理语句
var stmtCache = struct {
	sync.RWMutex
	m map[string]*sql.Stmt
}{
	m: make(map[string]*sql.Stmt),
}

// getStmt 从缓存获取或创建预处理语句
func getStmt(db *sql.DB, query string) (*sql.Stmt, error) {
	stmtCache.RLock()
	stmt, ok := stmtCache.m[query]
	stmtCache.RUnlock()

	if ok {
		return stmt, nil
	}

	stmtCache.Lock()
	defer stmtCache.Unlock()
	
	// 再次检查，防止并发创建
	if stmt, ok := stmtCache.m[query]; ok {
		return stmt, nil
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	stmtCache.m[query] = stmt
	return stmt, nil
}

// clearStmtCache 清理缓存
func clearStmtCache() {
	stmtCache.Lock()
	defer stmtCache.Unlock()
	for _, stmt := range stmtCache.m {
		stmt.Close()
	}
	stmtCache.m = make(map[string]*sql.Stmt)
}

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.StringVar(&url, "url", "127.0.0.1", "set database `url`")
	flag.StringVar(&port, "port", "50000", "set `database port` port")
	flag.StringVar(&dbname, "dbname", "sample", "set `dbname`")

	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	flag.PrintDefaults()
	os.Exit(-1)
}

func logwriter() *bufio.Writer {
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(execPath)
	dirlog := filepath.Join(dir, "logs")
	// 验证目录存在性
	fmt.Println(dir)
	fmt.Println(dirlog)
	os.MkdirAll(dirlog, os.ModePerm)
	if info, err := os.Stat(dirlog); err == nil && info.IsDir() {
		fmt.Printf("目录 %s 已成功创建\n", dirlog)
	} else {
		panic("目录创建失败")
	}
	now := time.Now()
	logdate := now.Format("20060102")
	deldate := now.AddDate(0, 0, -7).Format("20060102")

	filename := filepath.Join(dirlog, dbname+"_db2heartbeat_"+logdate+".log")
	delfilename := filepath.Join(dirlog, dbname+"_db2heartbeat_"+deldate+".log")

	_, err = os.Stat(delfilename)
	if os.IsExist(err) {
		err = os.Remove(delfilename)
	}
	if err != nil {
		fmt.Printf("文件 %s 不存在，忽略\n", delfilename)
	}

	filelog, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(filename)
	defer filelog.Close()
	writer := bufio.NewWriterSize(filelog, 256)
	return writer
}

func getLineSeparator() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

func continue_with_err(db *sql.DB, writer *bufio.Writer) (*sql.DB, *bufio.Writer) {
	if err := db.Ping(); err != nil {
		log.Printf("数据库不可用: %v, 尝试重连中...\n", err)
		db.Close() // 关闭无效连接
		clearStmtCache() // 清理stmt缓存

		db, err = connectDB(con)
		if err != nil {
			panic(err)
		}
	}

	if err := writer.Flush(); err != nil {
		writer = logwriter()
	}
	if err := writer.Flush(); err != nil {
		panic(err)
	}
	return db, writer
}

func main() {
	var db *sql.DB
	var err error
	var startt, delta, runMilli int64
	var totalexec int64

	startMilli := time.Now().UnixMilli()
	loc, _ := time.LoadLocation("Asia/Shanghai")

	if time.Now().In(loc).Hour() < 8 {
		totalexec = time.Now().In(loc).Truncate(24*time.Hour).Add(8*time.Hour).UnixMilli() - time.Now().In(loc).UnixMilli()
	}
	if time.Now().In(loc).Hour() >= 8 {
		totalexec = time.Now().In(loc).Truncate(24*time.Hour).Add(32*time.Hour).UnixMilli() - time.Now().In(loc).UnixMilli()
	}

	line_separator := getLineSeparator()
	writer := logwriter()
	flag.Parse()
	if help {
		flag.Usage()
	}
	_, err = strconv.Atoi(port)
	if err != nil {
		fmt.Printf("%v\n", "database port is invalid")
		flag.Usage()
	}

	fmt.Println(url)
	fmt.Println(port)
	fmt.Println(dbname)

	con = "HOSTNAME=" + url + ";DATABASE=" + dbname + ";PORT=" + port + ";UID=dbinq;PWD=dbinq"

	db, err = connectDB(con)
	if err != nil {
		log.Fatalf("无法启动程序: %v", err)
	}

	dbrole, err := getRole(db)
	if err != nil {
		panic(err)
	}
	if dbrole != "STANDBY" {
		runMilli = time.Now().UnixMilli() - startMilli
		for runMilli < totalexec {
			startt = time.Now().UnixMilli()
			logtime := time.Now().In(loc).Format("2006-01-02-15.04.05.000")
			err = insert(db)
			if err != nil {
				db, writer = continue_with_err(db, writer)
			}
			err = display(db)
			if err != nil {
				db, writer = continue_with_err(db, writer)
			}
			err = delete(db)
			if err != nil {
				db, writer = continue_with_err(db, writer)
			}
			delta = time.Now().UnixMilli() - startt
			_, err = writer.WriteString(logtime + " dml     " + strconv.Itoa(int(delta)) + line_separator)
			if err != nil {
				db, writer = continue_with_err(db, writer)
			}

			time.Sleep(300 * time.Millisecond)
			runMilli = time.Now().UnixMilli() - startMilli
		}
	}
}

func connectDB(dataSourceName string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	loc, _ := time.LoadLocation("Asia/Shanghai")

	for attempts := 0; attempts < 100000; attempts++ {
		db, _ = sql.Open("go_ibm_db", dataSourceName)
		startt := time.Now().UnixMilli()
		logtime := time.Now().In(loc).Format("2006-01-02-15.04.05.000")
		if err = db.Ping(); err != nil {
			log.Printf("数据库不可用: %v, 尝试重连中...\n", err)
			db.Close()
			time.Sleep(5 * time.Second)
			continue
		}
		delta := time.Now().UnixMilli() - startt
		fmt.Println(logtime + " connect " + strconv.Itoa(int(delta)))
		log.Println("数据库连接成功!")
		return db, nil
	}

	return nil, fmt.Errorf("无法连接到数据库 %v", err)
}

func delete(db *sql.DB) error {
	st, err := getStmt(db, "delete from monitor.heartbeat where application_id=mon_get_application_id()")
	if err != nil {
		return err
	}
	_, err = st.Exec()
	return err
}

func insert(db *sql.DB) error {
	st, err := getStmt(db, "insert into monitor.heartbeat(application_id,insert_time) values(mon_get_application_id(),current timestamp)")
	if err != nil {
		return err
	}
	_, err = st.Exec()
	return err
}

func display(db *sql.DB) error {
	st, err := getStmt(db, "select application_id,insert_time from monitor.heartbeat")
	if err != nil {
		return err
	}
	return execquery(st)
}

func execquery(st *sql.Stmt) error {
	rows, err := st.Query()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var t string
		var x time.Time
		err = rows.Scan(&t, &x)
		if err != nil {
			return err
		}
	}
	return nil
}

func getRole(db *sql.DB) (string, error) {
	var t string
	st, err := getStmt(db, "select value from sysibmadm.dbcfg where name='hadr_db_role'")
	if err != nil {
		return "123", err
	}
	rows, err := st.Query()
	if err != nil {
		return "123", err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&t)
		if err != nil {
			return "123", err
		}
	}
	return t, err
}
