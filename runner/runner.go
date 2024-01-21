package runner

import (
	"fineC/input"
	"fineC/output"
	"fineC/util"
)

type Runner struct {
	Input  input.Input
	Output output.Output
}

func (r *Runner) Run() {
	ch := make(chan []byte, util.APICounter)
	r.Input.Crawling(ch)
	r.Output.Convert(ch)
	r.Output.Send()
}
