package	logger	// import "github.com/nathanaelle/logger"

import	(
	"os"
	"flag"
	"strings"
)


func ExampleFactory_Set() {
	fs	:= flag.NewFlagSet("", flag.ExitOnError)

	fact	:= &Factory {
		CustomFlagHelper: func(d []string) string {
			return "choose a flag in : "+strings.Join(d, ", ")
		},
	}

	fact.Register(Stdout, NoLog)
	fact.SetDefault(NoLog)

	fs.SetOutput(os.Stdout)
	fs.Var(fact, "output", fact.FlagHelper())

	fs.PrintDefaults()
	fs.Parse([]string{"-output=stdout"})

	log	:= fact.LoggerFound()
	log.Printf("hello %s\n", log)

	// Output:
	// -output value
	//     	choose a flag in : stdout, null (default null)
	// hello stdout
	//
}
