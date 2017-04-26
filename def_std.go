package	logger	// import "github.com/nathanaelle/logger"

import	(
	"os"
)

type	(
	d_stdout struct {}
	d_stderr struct {}
)

var	(
	Stderr		Definition = register( new(d_stderr) )
	Stdout		Definition = register( new(d_stdout) )
)


func (_ *d_stdout)LoggerFound(buffer []byte) (Logger,bool) {
	if string(buffer) == "stdout" {
		return &l_writer{ "stdout", os.Stdout }, true
	}
	return nil, false
}

func (_ *d_stdout)DefaultLogger() Logger {
	return	&l_writer{ "stdout", os.Stdout }
}

func (_ *d_stdout)String() string {
	return	"stdout"
}

func (_ *d_stderr)LoggerFound(buffer []byte) (Logger,bool) {
	if string(buffer) == "stderr" {
		return &l_writer{ "stderr", os.Stderr }, true
	}
	return nil, false
}

func (_ *d_stderr)DefaultLogger() Logger {
	return	&l_writer{ "stderr", os.Stderr }
}

func (_ *d_stderr)String() string {
	return	"stderr"
}
