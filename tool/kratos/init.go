package main

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
)

const (
	_textModeFastInit    = "One-click initialization project"
	_textModeInteraction = "Custom project parameters"
	_textYes             = "Yes"
	_textNo              = "No"
)

func runInit(ctx *cli.Context) (err error) {
	if ctx.NumFlags() == 0 {
		if err = interact(); err != nil {
			return
		}
	}
	if !validate() {
		return nil
	}
	if err = create(); err != nil {
		fmt.Println("Project initialization failed: ", err.Error())
		return nil
	}
	fmt.Printf("The project [%s] was initialized successfully!\n", p.Path)
	return nil
}

func initPwd() (ok bool) {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	ps := strings.Split(pwd, string(os.PathSeparator))
	plen := len(ps)
	if plen < 1 {
		// At least one directory level: project name
		return
	}
	name := ps[plen-1]
	if name == "" {
		return
	}
	p.Name = name
	p.Path = pwd
	return true
}

func goPath() (gp string) {
	gopaths := strings.Split(os.Getenv("GOPATH"), ":")
	if len(gopaths) == 1 {
		return gopaths[0]
	}
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	abspwd, err := filepath.Abs(pwd)
	if err != nil {
		return
	}
	for _, gopath := range gopaths {
		absgp, err := filepath.Abs(gopath)
		if err != nil {
			return
		}
		if strings.HasPrefix(abspwd, absgp) {
			return absgp
		}
	}
	return build.Default.GOPATH
}

func interact() (err error) {
	qs1 := &survey.Select{
		Message: "How would you like to play?",
		Options: []string{_textModeFastInit, _textModeInteraction},
	}
	var ans1 string
	if err = survey.AskOne(qs1, &ans1, nil); err != nil {
		return
	}
	switch ans1 {
	case _textModeFastInit:
		if ok := initPwd(); !ok {
			fmt.Println("Fast initialization failed!")
		}
		return
	case _textModeInteraction:
		// go on
	default:
		return
	}
	qs := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Please enter a project name: ",
			},
			Validate: survey.Required,
		},
		{
			Name: "owner",
			Prompt: &survey.Input{
				Message: "Please enter the project leader: ",
			},
		},
		{
			Name: "useGRPC",
			Prompt: &survey.Select{
				Message: "Do you use gRPC? ",
				Options: []string{_textYes, _textNo},
				Default: _textNo,
			},
		},
		{
			Name: "here",
			Prompt: &survey.Select{
				Message: "Is the current directory? The default is GOPATH",
				Options: []string{_textYes, _textNo},
				Default: _textYes,
			},
		},
	}
	ans := struct {
		Name    string
		Owner   string
		UseGRPC string
		Here    string
	}{}
	if err = survey.Ask(qs, &ans); err != nil {
		return
	}
	p.Name = ans.Name
	p.Owner = ans.Owner
	if ans.UseGRPC == _textYes {
		p.WithGRPC = true
	}
	if ans.UseGRPC == _textYes {
		p.WithGRPC = true
	}
	if ans.Here == _textYes {
		p.Here = true
	}
	return
}
