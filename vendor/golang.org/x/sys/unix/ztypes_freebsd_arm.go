// cgo -godefs -- -fsigned-char types_freebsd.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build arm,freebsd

package unix

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]uint8
	Ispeed uint32
	Ospeed uint32
}
