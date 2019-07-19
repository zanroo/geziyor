package export

import (
	"encoding/csv"
	"fmt"
	"github.com/zanroo/geziyor/internal"
	"log"
	"os"
	"reflect"
	"sort"
)

// CSV exports response data as CSV streaming file
type CSV struct {
	FileName string
	Comma    rune
	UseCRLF  bool
}

// Export exports response data as CSV streaming file
func (e *CSV) Export(exports chan interface{}) {

	// Create or append file
	file, err := os.OpenFile(internal.DefaultString(e.FileName, "out.csv"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Output file creation error: %v\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = internal.DefaultRune(e.Comma, ',')
	writer.UseCRLF = e.UseCRLF

	// Export data as responses came
	for res := range exports {
		var values []string

		// Detect type and extract CSV values
		val := reflect.ValueOf(res)
		switch val.Kind() {
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				values = append(values, fmt.Sprint(val.Index(i)))
			}
		case reflect.Map:
			iter := val.MapRange()
			for iter.Next() {
				values = append(values, fmt.Sprint(iter.Value()))
			}
			sort.Strings(values)
		}
		if err := writer.Write(values); err != nil {
			log.Printf("CSV writing error on exporter: %v\n", err)
		}
	}
	writer.Flush()
}
