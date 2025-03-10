package main

import (
	"context"
	"entdemo/hzwent"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := hzwent.Open("mysql", "root:root@tcp(localhost:3306)/test?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	CreateStudent(context.Background(), client)

	stu, err := QueryStudent(context.Background(), client)
	fmt.Printf("stu:%v", stu)

}

func CreateStudentByFileds(ctx context.Context, client *hzwent.Client, name, stuno string, age int) (*hzwent.Student, error) {
	istu, err := client.Student.
		Create().
		SetName(name).
		SetStuNo(stuno).
		SetAge(age).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return istu, nil
}

func CreateStudentByObj(ctx context.Context, client *hzwent.Client, stu *hzwent.Student) (*hzwent.Student, error) {
	istu, err := client.Student.
		Create().
		SetName(stu.Name).
		SetStuNo(stu.StuNo).
		SetAge(stu.Age).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return istu, nil

}

func CreateStudent(ctx context.Context, client *hzwent.Client) (*hzwent.Student, error) {
	u, err := client.Student.
		Create().
		SetName("张三").
		SetStuNo("001").
		SetAge(18).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("student was created: ", u)
	return u, nil
}

func QueryStudent(ctx context.Context, client *hzwent.Client) ([]*hzwent.Student, error) {
	u, err := client.Student.
		Query().
		//		Where(student.Stuno("001")).
		All(ctx)
		// `Only` fails if no user found,
		// or more than 1 user returned.
		//Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("student returned: ", u)
	return u, nil
}

func UpdateStudent(ctx context.Context, client *hzwent.Client, stu *hzwent.Student) (*hzwent.Student, error) {
	u, err := client.Student.
		UpdateOne(stu).
		SetAge(20).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}
	log.Println("student was updated: ", u)
	return u, nil
}
