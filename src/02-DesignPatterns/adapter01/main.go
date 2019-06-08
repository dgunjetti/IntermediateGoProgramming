package main

import "fmt"

//START1 OMIT
type LegacyPrinter interface {
	Print(s string) string
}
type LegacyPrinterImp struct{}

func (l *LegacyPrinterImp) Print(s string) string {
	return fmt.Sprintf("Legacy: %s", s)
}

type ModernPrinter interface {
	PrintStored() string
}
type PrintAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

//END1 OMIT

//START2 OMIT
func (p *PrintAdapter) PrintStored() string {
	if p.OldPrinter != nil {
		newMsg := fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
		return newMsg
	} else {
		return p.Msg
	}
}
func main() {
	adapter := PrintAdapter{
		OldPrinter: nil,
		Msg:        "hello",
	}
	fmt.Println(adapter.PrintStored())

	adapter = PrintAdapter{
		OldPrinter: &LegacyPrinterImp{},
		Msg:        "hello",
	}
	fmt.Println(adapter.PrintStored())
}

//END2 OMIT
