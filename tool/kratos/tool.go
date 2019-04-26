package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func toolAction(c *cli.Context) (err error) {
	if c.NArg() == 0 {
		sort.Slice(toolIndexs, func(i, j int) bool { return toolIndexs[i].BuildTime.After(toolIndexs[j].BuildTime) })
		for _, t := range toolIndexs {
			updateTime := t.BuildTime.Format("2006/01/02")
			fmt.Printf("%s%s: %s %s (%s) [%s]\n", color.HiMagentaString(t.Name), getNotice(t), color.HiCyanString(t.Summary), t.URL, t.Author, updateTime)
		}
		fmt.Println("\nExecute the install installer such as: kratos tool install demo")
		fmt.Println("Execution tool name Run the program such as: kratos tool demo")
		fmt.Println("\nInstall all tools: kratos tool install all")
		return
	}
	if c.Args().First() == "install" {
		name := c.Args().Get(1)
		if name == "all" {
			installAll()
		} else {
			install(name)
		}
		return
	}
	name := c.Args().First()
	for _, t := range toolIndexs {
		if name == t.Name {
			if !t.installed() || t.updated() {
				install(name)
			}
			pwd, _ := os.Getwd()
			var args []string
			if c.NArg() > 1 {
				args = []string(c.Args())[1:]
			}
			runTool(t.Name, pwd, t.toolPath(), args)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "%s\n not installed yet", name)
	return
}

func upgradeAction(c *cli.Context) error {
	install("kratos")
	return nil
}

func install(name string) {
	if name == "" {
		fmt.Fprintf(os.Stderr, color.HiRedString("Please fill in the name of the tool to be installed\n"))
		return
	}
	for _, t := range toolIndexs {
		if name == t.Name {
			t.install()
			return
		}
	}
	fmt.Fprintf(os.Stderr, color.HiRedString("Installation failed %s\n not found", name))
	return
}

func installAll() {
	for _, t := range toolIndexs {
		if t.Install != "" {
			t.install()
		}
	}
}

func getNotice(t *Tool) (notice string) {
	if !t.supportOS() || t.Install == "" {
		return
	}
	notice = color.HiGreenString("(Not Installed)")
	if f, err := os.Stat(t.toolPath()); err == nil {
		notice = color.HiBlueString("(Installed)")
		if t.BuildTime.After(f.ModTime()) {
			notice = color.RedString("(updated)")
		}
	}
	return
}

func runTool(name, dir, cmd string, args []string) (err error) {
	toolCmd := &exec.Cmd{
		Path:   cmd,
		Args:   append([]string{cmd}, args...),
		Dir:    dir,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Env:    os.Environ(),
	}
	if filepath.Base(cmd) == cmd {
		var lp string
		if lp, err = exec.LookPath(cmd); err == nil {
			toolCmd.Path = lp
		}
	}
	if err = toolCmd.Run(); err != nil {
		if e, ok := err.(*exec.ExitError); !ok || !e.Exited() {
			fmt.Fprintf(os.Stderr, "Error running %s: %v\n", name, err)
		}
	}
	return
}

// Tool .
type Tool struct {
	Name      string    `json:"name"`
	BuildTime time.Time `json:"build_time"`
	Install   string    `json:"install"`
	Summary   string    `json:"summary"`
	Platform  []string  `json:"platform"`
	Author    string    `json:"author"`
	URL       string    `json:"url"`
}

func (t Tool) supportOS() bool {
	for _, p := range t.Platform {
		if strings.ToLower(p) == runtime.GOOS {
			return true
		}
	}
	return false
}

func (t Tool) install() {
	if t.Install == "" {
		fmt.Fprintf(os.Stderr, color.RedString("%s: Automatic installation failed See the documentation %s\n for details", t.Name, t.URL))
		return
	}
	cmds := strings.Split(t.Install, " ")
	if len(cmds) > 0 {
		if err := runTool(t.Name, path.Dir(t.toolPath()), cmds[0], cmds[1:]); err == nil {
			color.Green("%s: Installation is successful!", t.Name)
		}
	}
}

func (t Tool) updated() bool {
	if !t.supportOS() || t.Install == "" {
		return false
	}
	if f, err := os.Stat(t.toolPath()); err == nil {
		if t.BuildTime.After(f.ModTime()) {
			return true
		}
	}
	return false
}

func (t Tool) toolPath() string {
	return filepath.Join(goPath(), "bin", t.Name)
}

func (t Tool) installed() bool {
	_, err := os.Stat(t.toolPath())
	return err == nil
}
