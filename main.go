package main

import (
	"os"
	"path/filepath"
)

func main() {
	up, _ := os.UserHomeDir()
	psprofpath := filepath.Join(up, "Documents", "WindowsPowerShell", "Microsoft.PowerShell_profile.ps1")
	os.MkdirAll(filepath.Dir(psprofpath), os.ModePerm)
	file, _ := os.OpenFile(psprofpath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer file.Close()
	file.WriteString(`Start-Process "C:\Windows\System32\notepad.exe"` + "\n")
}
