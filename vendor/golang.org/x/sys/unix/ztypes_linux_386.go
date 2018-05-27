// cgo -godefs -- -Wall -Werror -static -I/tmp/include -m32 linux/types.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build 386,linux

package unix

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Line   uint8
	Cc     [19]uint8
	Ispeed uint32
	Ospeed uint32
}
