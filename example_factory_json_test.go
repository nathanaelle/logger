package	logger	// import "github.com/nathanaelle/logger"

import	(
	"log"
	"encoding/json"
)


func ExampleFactory_UnmarshalText() {
	var t  struct {
		Log   *Factory	`json:"output"`
	}

	data	:= []byte(`{"output":"stdout"}`)

	if err := json.Unmarshal(data, &t); err != nil {
		log.Fatal(err)
	}

	if t.Log == nil {
		log.Fatal("no output parsed")
	}

	output := t.Log.LoggerFound()
	if output == nil {
		log.Fatal("no output found")
	}

	output.Printf("hello %s\n", output)

	// Output:
	// hello stdout
	//
}
