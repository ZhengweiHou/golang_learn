package gormdemo

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/ibmdb"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var conhzw = "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"

func getHzwDb() *gorm.DB {
	dialector := ibmdb.Open(concrud)

	conf := &gorm.Config{
		PrepareStmt: true,
		//		DryRun:      true, // 生成SQL而不执行
		Logger: logger.New(log.New(log.Writer(), "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Info}),
	}

	db, err := gorm.Open(dialector, conf)
	if err != nil {
		log.Fatalf("open err: %v", err)
	}

	return db
}

// 新增
func TestHzwCreate(t *testing.T) {
	db := getHzwDb()

	// ========= Create =========
	hzw := Hzw{
		Name: "hzw",
		Age:  20,
	}
	// ---- DryRun show Create sql ----
	//stmt := db.Session(&gorm.Session{DryRun: true}).Create(&hzw).Statement
	//log.Printf("Create SQL: %v\n", stmt.SQL.String())
	//log.Printf("Create SQL Vars: %v\n", stmt.Vars)

	// ---- Create ----
	// TODO CreateAt什么时候赋值的
	result := db.Create(&hzw)
	if result.Error != nil {
		log.Fatalf("Create failed: %v", result.Error)
	}
	jstr, err := json.Marshal(hzw)
	if err != nil {
		log.Fatalf("json marshal failed: %v", err)
	}
	fmt.Printf("Affected: %v Created hzw: %s\n", result.RowsAffected, jstr)

}

// 新增 带主键
func TestHzwCreateWithIdA(t *testing.T) {
	db := getHzwDb()

	// ========= Create =========
	hzw := Hzw{
		ID:   111,
		Name: "hzw",
		Age:  20,
	}
	result := db.Create(&hzw)
	if result.Error != nil {
		log.Fatalf("Create failed: %v", result.Error)
	}
	jstr, err := json.Marshal(hzw)
	if err != nil {
		log.Fatalf("json marshal failed: %v", err)
	}
	fmt.Printf("Affected: %v Created hzw: %s\n", result.RowsAffected, jstr)

}

// 批量新增
func TestHzwCreateMulti(t *testing.T) {
	db := getHzwDb()

	// 批次插入
	hzws := []*Hzw{
		{Name: "hzw1", Age: 21},
		{Name: "hzw2", Age: 22},
	}
	result := db.Create(hzws)
	if result.Error != nil {
		log.Fatalf("Create failed: %v", result.Error)
	}
	log.Printf("Affected: %v\n", result.RowsAffected)
	for _, hzw := range hzws {
		jstr, err := json.Marshal(hzw)
		if err != nil {
			log.Fatalf("json marshal failed: %v", err)
		}
		fmt.Printf("Created hzw: %s\n", jstr)
	}

	// 批次插入 指定批次数量
	hzws2 := []*Hzw{
		{Name: "hzw1", Age: 21}, {Name: "hzw2", Age: 22}, {Name: "hzw3", Age: 21}, {Name: "hzw4", Age: 22},
		{Name: "hzw1", Age: 21}, {Name: "hzw2", Age: 22}, {Name: "hzw3", Age: 21}, {Name: "hzw4", Age: 22},
		{Name: "hzw1", Age: 21}, {Name: "hzw2", Age: 22}, {Name: "hzw3", Age: 21}, {Name: "hzw4", Age: 22},
		{Name: "hzw1", Age: 21}, {Name: "hzw2", Age: 22}, {Name: "hzw3", Age: 21}, {Name: "hzw4", Age: 22},
	}
	db.CreateInBatches(hzws2, 4)
	for _, hzw := range hzws2 {
		jstr, err := json.Marshal(hzw)
		if err != nil {
			log.Fatalf("json marshal failed: %v", err)
		}
		fmt.Printf("Created hzw: %s\n", jstr)
	}
}

// 查询
func TestHzwFind(t *testing.T) {
	db := getHzwDb() // DB2

	//	== find
	hzw := &Hzw{}
	db.Find(hzw) // 注意这种查询会全表查询，只返回一条
	//-- SELECT * FROM hzw
	jstr, _ := json.Marshal(hzw)
	fmt.Printf("find hzw: %s\n", jstr)

	// == find with id
	hzw = &Hzw{}
	hzw.ID = 1
	db.Find(hzw)
	// -- SELECT * FROM hzw WHERE hzw.id = 1
	jstr, _ = json.Marshal(hzw)
	fmt.Printf("find with id: %s\n", jstr)

	// == find where
	hzw = &Hzw{}
	db.Where("ID = ?", 1).Find(hzw)
	// -- SELECT * FROM hzw WHERE ID = 1
	jstr, _ = json.Marshal(hzw)
	fmt.Printf("find by id hzw: %s\n", jstr)

	// == find limit
	hzws := []Hzw{}
	db.Limit(4).Find(&hzws)
	// -- SELECT * FROM hzw FETCH FIRST 4 ROWS ONLY
	jstr, _ = json.Marshal(hzws)
	fmt.Printf("finds limit hzws: %s\n", jstr)

	// == find with ids
	hzws = []Hzw{}
	db.Find(&hzws, []int{1, 2, 3})
	// -- SELECT * FROM hzw WHERE hzw.id IN (1,2,3)
	jstr, _ = json.Marshal(hzws)
	fmt.Printf("finds with ids: %s\n", jstr)

	// == first 默认主键升序 若model没定义主键，则按model第一个字段排序
	hzw = &Hzw{}
	db.First(hzw)
	// -- SELECT * FROM hzw ORDER BY hzw.id FETCH FIRST 1 ROWS ONLY
	fmt.Printf("first:%v", hzw)

	// == first 根据主键检索
	hzw = &Hzw{}
	db.First(hzw, 2)
	// -- SELECT * FROM hzw WHERE hzw.id = 2 ORDER BY hzw.id FETCH FIRST 1 ROWS ONLY
	fmt.Printf("first by id:%v", hzw)

	// == last 主键DESC 第一个 若model没定义主键，则按model第一个字段排序
	hzw = &Hzw{} // 有主键会按主键查询
	db.Last(hzw)
	// -- SELECT * FROM hzw ORDER BY hzw.id DESC FETCH FIRST 1 ROWS ONLY
	fmt.Printf("last:%v", hzw)

	// == take 没有指定排序
	hzw = &Hzw{}
	db.Take(hzw)
	// -- SELECT * FROM hzw FETCH FIRST 1 ROWS ONLY
	fmt.Printf("take:%v", hzw)

	// == take、first、last 若有主键则会将主键作为查询条件
	hzw = &Hzw{}
	hzw.ID = 3
	db.Take(hzw)
	// -- SELECT * FROM hzw WHERE hzw.id = 1 FETCH FIRST 1 ROWS ONLY
	fmt.Printf("take:%v", hzw)

	// == count
	count := int64(0)
	db.Model(&Hzw{}).Count(&count)
	fmt.Printf("count: %d \n", count)

	// == distinct
	hzws = []Hzw{}
	tx := db.Distinct("NAME").Find(&hzws)
	fmt.Printf("distinct txerr:%v \n", tx.Error)
	for _, hzw := range hzws {
		fmt.Printf("%s\n", hzw.Name)
	}
}

// 删除
func TestHzwDelete(t *testing.T) {
	db := getHzwDb()
	hzw := Hzw{Name: "deltest"}
	db.Create(&hzw)
	fmt.Printf("insert id:%d hzw:%v\n", hzw.ID, hzw)

	selhzw := &Hzw{}
	res := db.Find(selhzw, hzw.ID)
	fmt.Printf("find affect:%d\n", res.RowsAffected)
	fmt.Printf("find by id:%d, finded:%v\n", hzw.ID, selhzw)

	// 软删除
	db.Delete(&hzw)
	selhzw = &Hzw{}
	res = db.Find(selhzw, hzw.ID)
	fmt.Printf("find affect:%d\n", res.RowsAffected)

	tx := db.Unscoped()
	// 硬查询
	selhzw = &Hzw{}
	res = tx.Find(selhzw, hzw.ID)
	fmt.Printf("unscoped find affect:%d\n hzw:%v\n", res.RowsAffected, selhzw)
	// 硬删除
	tx.Delete(&hzw)
}

// 更新
func TestHzwUpdate(t *testing.T) {
	db := getHzwDb()
	//db := getMysqlDb()

	// == Save() 没主键时为Create操作 ==
	hzw := &Hzw{Name: "upserttest", Age: 11}
	db.Save(hzw) // Save没指定主键时为 insert
	jstr, _ := json.Marshal(hzw)
	fmt.Printf("save hzw: %s\n", jstr)

	// == Model + updates 只更新非空字段 ==
	hzw2 := &Hzw{
		ID:  hzw.ID,
		Age: 22,
	}
	db.Model(hzw2).Updates(hzw2) // 只更新非空字段
	jstr, _ = json.Marshal(hzw2)
	fmt.Printf("save hzw: %s\n", jstr) // 其他字段不会自动从库中赋值到 实体 中

	// == Save() 有主键时为Update操作 但包含所有空值 ==
	hzw3 := &Hzw{
		ID:  hzw.ID,
		Age: 33,
	}
	db.Save(hzw3) // 指定的主键为update，默认更新所有字段包含所有空值
	jstr, _ = json.Marshal(hzw3)
	fmt.Printf("save hzw: %s\n", jstr) // Name被更新为空串了

}
