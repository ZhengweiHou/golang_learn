package gormdemo

import (
	"context"
	"encoding/json"
	"fmt"
	"gormdemo/dao"
	"gormdemo/model"
	"log"
	"testing"
	"time"
)

func TestBasedao(t *testing.T) {
	ctx := context.Background()
	studao := dao.StudentDao{}

	stu1 := &model.Student{
		Name: "张三",
		Age:  12,
	}

	// insert one
	affect, _ := studao.InsertOne(&ctx, stu1)
	jstr, _ := json.Marshal(stu1)
	fmt.Printf("affact:%d stu: %s\n", affect, jstr)

	stukey := &model.Student{
		Id: stu1.Id,
	}

	// find by key
	stu2, err := studao.FindByPrimaryKey(&ctx, stukey)
	if err != nil {
		log.Fatal(err)
	}
	jstr, _ = json.Marshal(stu2)
	fmt.Printf("find: %s\n", jstr)

	// update by key
	stu2.Name = "张三三"
	affect, _ = studao.UpdateByPrimaryKey(&ctx, stu2)
	jstr, _ = json.Marshal(stu2)
	fmt.Printf("update by key affact:%d stu: %s\n", affect, jstr)
	// TODO 更新忽略空值

	// find by key after update
	stu3, err := studao.FindByPrimaryKey(&ctx, stukey)
	if err != nil {
		log.Fatal(err)
	}
	jstr, _ = json.Marshal(stu3)
	fmt.Printf("find after update: %s\n", jstr)
	// finds by name
	namearg := &model.FindByNameArg{
		Name: "hzw1",
	}
	stus, err := studao.FindByName(&ctx, *namearg)
	if err != nil {
		log.Fatal(err)
	}
	for _, stu := range stus {
		jstr, err := json.Marshal(stu)
		if err != nil {
			log.Fatalf("json marshal failed: %v", err)
		}
		fmt.Printf("finds by name stu: %s\n", jstr)
	}

	// delete by key
	//	delstu := &model.Student{
	//		Id: stu1.Id,
	//	}
	//	affect, _ = studao.DeleteByPrimaryKey(&ctx, delstu)
	//	fmt.Printf("delete by key affact:%d \n", affect)

	time.Sleep(time.Microsecond * 100)

}
