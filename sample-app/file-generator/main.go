package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

type fileContent struct {
	packageName    string
	interfaceName  string
	structName     string
	initMethodName string
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	for {
		items := append([]string{"..", "CREATE FILE"}, getDirectoryNames(currentDir)...)

		prompt := promptui.Select{
			Label: "Select a directory or action",
			Items: items,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . | bold }}",
				Selected: "{{ . | green | bold }}",
				Active:   "{{ . | cyan | bold }}",
				Inactive: "{{ . | yellow }}",
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Println("Prompt failed:", err)
			return
		}

		if result == ".." {
			currentDir = filepath.Dir(currentDir)
			continue
		} else if result == "CREATE FILE" {
			filePrompt := promptui.Prompt{
				Label: "Enter the name of the file to create",
			}

			fileName, err := filePrompt.Run()
			if err != nil {
				fmt.Println("Prompt failed:", err)
				return
			}

			filePath := filepath.Join(currentDir, fmt.Sprintf("%s.go", fileName))
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer file.Close()

			createContentOfFile(file, currentDir)

			fmt.Println("Created file:", filePath)
			return
		}

		selectedDir := filepath.Join(currentDir, result)

		if isDirectory(selectedDir) {
			fmt.Println("Navigating to:", selectedDir)
			currentDir = selectedDir
			continue
		} else {
			fmt.Println("Selected item is not a directory.")
		}
	}
}

func createContentOfFile(file *os.File, dir string) {
	structPrompt := promptui.Prompt{
		Label: "Enter the Struct Name",
	}
	structName, err := structPrompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}

	interfacePrompt := promptui.Prompt{
		Label: "Enter the Interface Name",
	}
	interfaceName, err := interfacePrompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}

	initMethodPrompt := promptui.Prompt{
		Label: "Enter the Init Method Name",
	}
	initMethod, err := initMethodPrompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}

	content := fileContent{
		packageName:    filepath.Base(dir),
		structName:     structName,
		interfaceName:  interfaceName,
		initMethodName: initMethod,
	}

	_, err = file.WriteString(fmt.Sprintf(
		"package %s\n\n",
		content.packageName,
	))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	if content.interfaceName != "" {
		_, err = file.WriteString(fmt.Sprintf(
			"type %s interface {\n}\n\n",
			content.interfaceName,
		))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	if content.structName != "" {
		_, err = file.WriteString(fmt.Sprintf(
			"type %s struct {\n}\n\n",
			content.structName,
		))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	if content.initMethodName != "" {
		_, err = file.WriteString(fmt.Sprintf(
			"func %s() %s {\n\treturn &%s{}\n}",
			content.initMethodName,
			content.interfaceName,
			content.structName,
		))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func getDirectoryNames(currentDir string) []string {
	files, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}
	return dirs
}

func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
