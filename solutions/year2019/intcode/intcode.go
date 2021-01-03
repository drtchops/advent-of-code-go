package intcode

import (
	"fmt"
	"strconv"

	"github.com/drtchops/aoc/utils"
)

type TermSig struct {
	ID     int64
	Output int64
	Err    error
}

type Optcode int64
type ParamMode int64

const (
	OptcodeAdd       Optcode = 1
	OptcodeMul       Optcode = 2
	OptcodeInput     Optcode = 3
	OptcodeOutput    Optcode = 4
	OptcodeJumpTrue  Optcode = 5
	OptcodeJumpFalse Optcode = 6
	OptcodeLessThan  Optcode = 7
	OptcodeEqual     Optcode = 8
	OptcodeRelative  Optcode = 9
	OptcodeTerm      Optcode = 99

	ParamPosition  ParamMode = 0
	ParamImmediate ParamMode = 1
	ParamRelative  ParamMode = 2
)

var paramCounts = map[Optcode]int64{
	OptcodeAdd:       3,
	OptcodeMul:       3,
	OptcodeInput:     1,
	OptcodeOutput:    1,
	OptcodeJumpTrue:  2,
	OptcodeJumpFalse: 2,
	OptcodeLessThan:  3,
	OptcodeEqual:     3,
	OptcodeRelative:  1,
	OptcodeTerm:      0,
}
var writeParams = map[Optcode]int64{
	OptcodeAdd:      2,
	OptcodeMul:      2,
	OptcodeInput:    0,
	OptcodeLessThan: 2,
	OptcodeEqual:    2,
}

func Parse(input string) []int64 {
	prog := utils.ParseInputInts(input, ",")
	prog = append(prog, make([]int64, len(prog)*10)...)
	return prog
}

func Run(prog []int64, id int64, input, output chan int64, term chan TermSig) {
	var ptr int64
	var relAddr int64
	var lastOutput int64
	var err error

	for {
		jump := false
		opt, pms := parseOptcode(prog[ptr])
		count, ok := paramCounts[opt]
		if !ok {
			err = fmt.Errorf("unknown opt, ptr=%d line=%v opt=%d pms=%v", ptr, prog[ptr:ptr+count+1], opt, pms)
			break
		}
		if int64(len(prog)) < ptr+count {
			err = fmt.Errorf("not enough values, ptr=%d line=%v opt=%d pms=%v", ptr, prog[ptr:ptr+count+1], opt, pms)
			break
		}

		if opt == OptcodeTerm {
			// fmt.Printf("ptr=%d line=%v opt=%d\n", ptr, prog[ptr:ptr+count+1], opt)
			break
		}

		params := make([]int64, count)
		var i int64
		for i = 0; i < count; i++ {
			var pm ParamMode
			if i < int64(len(pms)) {
				pm = pms[i]
			}

			writeIdx, ok := writeParams[opt]
			var val int64
			param := prog[ptr+i+1]

			if ok && i == writeIdx {
				if pm == ParamRelative {
					val = relAddr + param
				} else {
					val = param
				}
			} else {
				if pm == ParamPosition {
					val = prog[param]
				} else if pm == ParamImmediate {
					val = param
				} else {
					val = prog[relAddr+param]
				}
			}
			params[i] = val
		}

		// fmt.Printf("ptr=%d line=%v opt=%d params=%v pms=%v\n", ptr, prog[ptr:ptr+count+1], opt, params, pms)

		if opt == OptcodeAdd {
			writePtr := params[2]
			val := params[0] + params[1]
			if err := writeVal(prog, writePtr, val); err != nil {
				break
			}
		} else if opt == OptcodeMul {
			writePtr := params[2]
			val := params[0] * params[1]
			if err := writeVal(prog, writePtr, val); err != nil {
				break
			}
		} else if opt == OptcodeInput {
			writePtr := params[0]
			// fmt.Printf("%d input\n", id)
			val := <-input
			if err := writeVal(prog, writePtr, val); err != nil {
				break
			}
		} else if opt == OptcodeOutput {
			lastOutput = params[0]
			// fmt.Printf("%d output %d\n", id, lastOutput)
			output <- lastOutput
		} else if opt == OptcodeJumpTrue {
			if params[0] != 0 {
				ptr = params[1]
				jump = true
			}
		} else if opt == OptcodeJumpFalse {
			if params[0] == 0 {
				ptr = params[1]
				jump = true
			}
		} else if opt == OptcodeLessThan {
			var val int64
			if params[0] < params[1] {
				val = 1
			}
			writePtr := params[2]
			if err := writeVal(prog, writePtr, val); err != nil {
				break
			}
		} else if opt == OptcodeEqual {
			var val int64
			if params[0] == params[1] {
				val = 1
			}
			writePtr := params[2]
			if err := writeVal(prog, writePtr, val); err != nil {
				break
			}
		} else if opt == OptcodeRelative {
			relAddr += params[0]
		}

		if !jump {
			ptr += count + 1
		}
	}

	term <- TermSig{
		ID:     id,
		Output: lastOutput,
		Err:    err,
	}
}

func writeVal(prog []int64, ptr, val int64) error {
	if ptr < 0 || ptr >= int64(len(prog)) {
		return fmt.Errorf("address %d out of range", ptr)
	}
	prog[ptr] = val
	return nil
}

func parseOptcode(inst int64) (Optcode, []ParamMode) {
	is := strconv.FormatInt(inst, 10)

	opts := ""
	if len(is) == 1 {
		opts = is
	} else {
		opts = is[len(is)-2:]
	}
	opt, _ := strconv.ParseInt(opts, 10, 64)

	pms := make([]ParamMode, 0)
	for i := len(is) - 3; i >= 0; i-- {
		p, _ := strconv.ParseInt(string(is[i]), 10, 64)
		pms = append(pms, ParamMode(p))
	}

	return Optcode(opt), pms
}
