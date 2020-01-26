package accountclient

// ClientLogger  The interface the logger should support
type ClientLogger interface {
	Fatalln(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}
