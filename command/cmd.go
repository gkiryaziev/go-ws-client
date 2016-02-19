package command

import (
	"log"
	"os/exec"
	"strings"
)

type Command struct {
}

type args struct {
	cmd       string
	cmdArgs   []string
	cleanArgs []string
}

// Constructor
func NewCommand() *Command {
	return &Command{}
}

// Command
func (this *Command) Command(c string) *args {
	return &args{c, nil, nil}
}

// Args for command
func (this *args) Arg(a string) *args {
	this.cmdArgs = append(this.cmdArgs, a)
	return this
}

// String for clean
func (this *args) Clean(a string) *args {
	this.cleanArgs = append(this.cleanArgs, a)
	return this
}

// Run the command
func (this *args) Run() string {
	cmdRes := cmdExec(this.cmd, this.cmdArgs...)
	if cmdRes == "" {
		return cmdRes
	}
	return clean(cmdRes, this.cleanArgs...)
}

func cmdExec(name string, args ...string) string {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(out)
}

func clean(str string, args ...string) string {
	s := strings.TrimSpace(str)
	for _, arg := range args {
		s = strings.Replace(s, arg, "", -1)
	}
	return s
}
