package main

import (
	"fmt"
	"log"
	"os"

	"github.com/na0chan-go/protobuf-lesson/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	// データを作成
	employee := &pb.Employee{
		Id:          1,
		Name:        "na0chan",
		Email:       "na0chan@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project: map[string]*pb.Company_Project{
			"ProjectX": &pb.Company_Project{},
		},
		Profile: &pb.Employee_Text{
			Text: "Hello, World!",
		},
		Birthday: &pb.Date{
			Year:  1996,
			Month: 2,
			Day:   18,
		},
	}
	// バイナリデータに変換
	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Failed to encode employee:", err)
	}

	// バイナリデータをファイルに出力
	err = os.WriteFile("employee.bin", binData, 0644)
	if err != nil {
		log.Fatalln("Failed to write employee:", err)
	}

	// バイナリデータをファイルから読み込む
	in, err := os.ReadFile("employee.bin")
	if err != nil {
		log.Fatalln("Failed to read employee:", err)
	}

	// 空の構造体を作成
	readEmployee := &pb.Employee{}
	// バイナリデータをデコード
	err = proto.Unmarshal(in, readEmployee)
	if err != nil {
		log.Fatalln("Failed to decode employee:", err)
	}

	// デシリアライズしたデータを出力
	fmt.Println(readEmployee)
}
