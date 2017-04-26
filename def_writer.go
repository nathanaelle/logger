package	logger	// import "github.com/nathanaelle/logger"

import	(
	"io"
	"log"
	"fmt"
	"os"
)

type	(
	d_writer struct {
		name	string
		w	io.Writer
	}

	l_writer struct {
		name	string
		w	io.Writer
	}
)


func GenericWriter(name string, w io.Writer) Definition {
	return &d_writer { name, w }
}

func (d d_writer)LoggerFound(buffer []byte) (Logger,bool) {
	if string(buffer) == d.name {
		return &l_writer{ d.name, d.w }, true
	}
	return nil, false
}

func (d d_writer)DefaultLogger() Logger  {
	return	nil
}

func (d d_writer)String() string {
	return	d.name
}

func (w l_writer)Print(v ...interface{}) {
	fmt.Fprint(w.w, v...)
}

func (w l_writer)Println(v ...interface{}) {
	fmt.Fprintln(w.w, v...)
}

func (w l_writer)Printf(format string, v ...interface{}) {
	fmt.Fprintf(w.w, format, v...)
}

func (w l_writer)Fatal(v ...interface{}) {
	s := fmt.Sprint(v...)
	w.Print(s)
	os.Exit(1)
}

func (w l_writer)Fatalln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	w.Print(s)
	os.Exit(1)
}

func (w l_writer)Fatalf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	w.Print(s)
	os.Exit(1)
}

func (w l_writer)Logger(prefix string, flag int)	*log.Logger {
	return log.New(w.w, prefix, flag)
}

func (w l_writer)Write(d []byte) (int,error) {
	return w.w.Write(d)
}


func (w l_writer)String() string {
	return w.name
}
