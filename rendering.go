package main

import (
	"io"
	"os/exec"
	"path/filepath"
	"strings"
)

func rendering(output Output, content string) {
	cmdArgs := []string{
		"-o",
		Config.GetString("buildFolder") + "/" + output.Name + "." + output.Extention,
	}
	if output.Extention == "pdf" || output.Extention == "tex" {
		cmdArgs = append(cmdArgs, "--latex-engine="+Config.GetString("latexEngine"))
		filename := strings.TrimSuffix(output.Template, filepath.Ext(output.Template))
		tmpl := Config.GetString("buildFolder") + "/" + filename
		cmdArgs = append(cmdArgs, "--template="+tmpl)
	}
	cmd := exec.Command("pandoc", cmdArgs...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		Logger.Fatalf("[RENDERING] Can't stdin pipe : %s", err)
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, content)
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		Logger.Fatalf("[RENDERING] %s%s", out, err)
	}
}
