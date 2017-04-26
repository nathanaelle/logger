package	logger	// import "github.com/nathanaelle/logger"

import	(
	"log"
	"fmt"
	"io"
)

type	(
	Definition	interface {
		LoggerFound(buffer []byte) (Logger,bool)
		DefaultLogger() Logger
		String() string
	}

	Logger	interface {
		io.Writer
		Print(v ...interface{})
		Println(v ...interface{})
		Printf(format string, v ...interface{})
		Fatal(v ...interface{})
		Fatalln(v ...interface{})
		Fatalf(format string, v ...interface{})
		Logger(prefix string, flag int)	*log.Logger
		String() string
	}

	Factory	struct {
		CustomFlagHelper	func([]string) string
		index			[]Definition
		deflt			Logger
		found			Logger
	}
)

var NoMatchingDef	error	= fmt.Errorf("No Matching Definition Found")
var logger	*Factory	= &Factory{}


func register(def Definition) Definition {
	logger.index = append(logger.index, def)
	return	def
}


func (l *Factory)Register(defs ...Definition) {
	l.index = append(l.index, defs...)
}

func (l *Factory)SetDefault(def Definition) {
	l.deflt = def.DefaultLogger()
}


func (l *Factory)String() string {
	if l.found != nil {
		return l.found.String()
	}
	if l.deflt != nil {
		return l.deflt.String()
	}
	return	""
}

func (l *Factory)Get() interface{} {
	return l.LoggerFound()
}


func (l *Factory)Set(data string) error {
	d := []byte(data)

	if l.index == nil || len(l.index) == 0 {
		l.Register(logger.index...)
	}

	for _,i := range l.index {
		if logger, ok := i.LoggerFound(d); ok {
			l.found = logger
			return nil
		}
	}

	if l.deflt != nil {
		l.found = l.deflt
		return nil
	}

	return	NoMatchingDef
}


func (l *Factory)FlagHelper() string {
	if l.index == nil || len(l.index) == 0 {
		l.Register(logger.index...)
	}

	a := make([]string,len(l.index))
	for i,d := range l.index {
		a[i] = d.String()
	}

	if l.CustomFlagHelper != nil {
		return l.CustomFlagHelper(a)
	}

	return fmt.Sprintf("possible values : %+v", a)
}

func (l *Factory)LoggerFound() Logger {
	if l.found != nil {
		return l.found
	}
	if l.deflt != nil {
		return l.deflt
	}

	log.Panicf("no logger.Logger found yet and no default logger.Definition")
	return nil
}

func (p *Factory)MarshalText() ([]byte, error) {
	return	[]byte(p.String()), nil
}

func (p *Factory)UnmarshalText(text []byte) error {
	return	p.Set(string(text))
}
