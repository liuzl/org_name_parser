package main

import (
	"strconv"
	"strings"

	"github.com/bcongdon/corral"
	"github.com/liuzl/ling"
)

const min = 2
const max = 6

var nlp = ling.MustNLP(ling.Norm)

type namemr struct{}

func (w namemr) Map(key, value string, emitter corral.Emitter) {
	d := ling.NewDocument(value)
	if err := nlp.Annotate(d); err != nil {
		return
	}
	tokens := d.XRealTokens(ling.Norm)
	n := len(tokens)
	for i := min; i < max; i++ {
		for j := 0; j+i <= n; j++ {
			var one []string
			for k, token := range tokens {
				if k == j {
					one = append(one, "#")
				} else if k > j && k <= j+i {
				} else {
					one = append(one, token)
				}
			}
			emitter.Emit(strings.Join(one, ""), "1")
		}
	}
}

func (w namemr) Reduce(key string, values corral.ValueIterator, emitter corral.Emitter) {
	count := 0
	for range values.Iter() {
		count++
	}
	emitter.Emit(key, strconv.Itoa(count))
}

func main() {
	job := corral.NewJob(namemr{}, namemr{})

	options := []corral.Option{
		//corral.WithSplitSize(10 * 1024),
		//corral.WithMapBinSize(10 * 1024),
	}

	driver := corral.NewDriver(job, options...)
	driver.Main()
}
