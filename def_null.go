package	logger	// import "github.com/nathanaelle/logger"

import	(
	"log"
	"os"
)

type	(
	d_null	struct {}
	l_null	struct {}
	w_null	struct {}
)

var NoLog	Definition	= register( new(d_null) )
var lognull	Logger		= new(l_null)


func (_ w_null)Write(d []byte) (int,error) {
	return len(d),nil
}


func (_ d_null)LoggerFound(buffer []byte) (Logger,bool) {
	if string(buffer) == "null" {
		return lognull, true
	}
	return nil, false
}

func (_ d_null)DefaultLogger() Logger {
	return lognull
}

func (_ d_null)String() string {
	return	"null"
}


func (_ l_null)Write(d []byte) (int,error) {
	return len(d),nil
}

func (_ l_null)Print(_ ...interface{}) {
}

func (_ l_null)Println(_ ...interface{}) {
}

func (_ l_null)Printf(_ string, _ ...interface{}) {
}

func (_ l_null)Fatal(_ ...interface{}) {
	os.Exit(1)
}

func (_ l_null)Fatalln(_ ...interface{}) {
	os.Exit(1)
}

func (_ l_null)Fatalf(_ string, _ ...interface{}) {
	os.Exit(1)
}

func (l l_null)Logger(prefix string, flag int)	*log.Logger {
	return log.New(l, prefix, flag)
}


func (_ l_null)String() string {
	return "null"
}
