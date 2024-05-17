package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

func main() {
	// コマンドラインから入力値を受け取る
	projectName := flag.String("project-name", "", "プロジェクト名")
	author := flag.String("author", "", "作者名")
	flag.Parse()

	// テンプレートファイルを読み込む
	tmpl, err := template.ParseFiles("main.go.tmpl")
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return
	}

	// 生成ファイルを開く
	file, err := os.Create("main.go")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	// テンプレートにデータを渡してファイルに書き込む
	err = tmpl.Execute(file, struct {
		ProjectName string
		Author      string
	}{
		ProjectName: *projectName,
		Author:      *author,
	})
	if err != nil {
		fmt.Println("Failed to execute template:", err)
		return
	}

	fmt.Println("File generated successfully.")
}
