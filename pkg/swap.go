package pkg

// swap in mem functions
import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

var beforeSwap *MonitorSwap

func init() {
	beforeSwap, err = NewSwap()
}

type SwapUsage struct {
	In  float64
	Out float64
}

type MonitorSwap struct {
	swap_in  uint64
	swap_out uint64
}

func (this *MonitorSwap) Get() error {
	data, err := mem.SwapMemory()
	if err != nil {
		return err
	}
	this.swap_in = data.Sin
	this.swap_out = data.Sout

	return nil
}

func NewSwap() (*MonitorSwap, error) {
	data := &MonitorSwap{}
	err := data.Get()
	return data, err
}

func SwapIO() (string, error, SwapUsage) {
	var rs string
	var result SwapUsage
	after, err := NewSwap()
	if err != nil {
		return rs, err, result
	}
	si := after.swap_in - beforeSwap.swap_in
	so := after.swap_out - beforeSwap.swap_out

	// in := strings.Repeat(" ", 5-len(fmt.Sprintf("%d", si))) + fmt.Sprintf("%d", si)
	in := parseRepeatSpace(fmt.Sprintf("%d", si), 5)
	// out := strings.Repeat(" ", 5-len(fmt.Sprintf("%d", so))) + fmt.Sprintf("%d", so)
	out := parseRepeatSpace(fmt.Sprintf("%d", so), 5)
	result.In = float64(si)
	result.Out = float64(so)
	if si > 0 {
		rs += Colorize(in, "red", "", false, true)
	} else {
		rs += Colorize(in, "white", "", false, false)
	}

	if so > 0 {
		rs += Colorize(out, "red", "", false, true)
	} else {
		rs += Colorize(out, "white", "", false, false)
	}

	rs += Colorize("|", "dgreen", "", false, false)
	beforeSwap = after
	return rs, nil, result
}
