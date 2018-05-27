// cgo -godefs types_netbsd.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build arm,netbsd

package unix

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]uint8
	Ispeed int32
	Ospeed int32
}
