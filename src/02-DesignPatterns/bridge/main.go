package main

import (
	"fmt"
	"io"
)

//START1 OMIT
type PrinterAPI interface {
	PrintMsg(string) error
}
type PrinterAbstraction interface {
	Print() error
}

type Impl1 struct{}
type Impl2 struct {
	Writer io.Writer
}

//END1 OMIT
//START2 OMIT
func (i *Impl1) PrintMsg(s string) error {
	fmt.Println(s)
	return nil
}
func (i *Impl2) PrintMsg(s string) error {
	if i.Writer != nil {
		fmt.Fprintf(i.Writer, "%s", s)
		return nil
	}
	return fmt.Errorf("Writer not set")
}

//END2 OMIT
type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(b []byte) (int, error) {
	n := len(b)
	t.Msg = string(b)
	return n, nil
}

//START3 OMIT
type User1 struct {
	Msg     string
	Printer PrinterAPI
}

func (u *User1) Print() error {
	u.Printer.PrintMsg(u.Msg)
	return nil
}

type User2 struct {
	Msg     string
	Printer PrinterAPI
}

func (u *User2) Print() error {
	u.Printer.PrintMsg(fmt.Sprintf("Msg from user2: %s", u.Msg))
	return nil
}

//END3 OMIT

func main() {
	//START4 OMIT
	u1 := User1{
		Msg: "hello",
		Printer: &Impl2{
			Writer: &TestWriter{},
		},
	}
	u1.Print()

	u2 := User2{
		Msg:     "hello",
		Printer: &Impl1{},
	}
	u2.Print()
	//END4 OMIT
}
