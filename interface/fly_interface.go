package _interface

type Fly interface {
	Fly()
}

type ReadWriter interface {
	Read(p []byte) (n int , err error)
	Write(p []byte) (n int , err error)
}