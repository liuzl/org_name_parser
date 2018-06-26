package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb"
	"github.com/dgryski/go-topk"
	"github.com/liuzl/goutil"
	"github.com/liuzl/ling"
)

var (
	k = flag.Int("n", 10000, "k")
	i = flag.String("i", "input.txt.gz", "input file")
	c = flag.Bool("c", false, "use count flag")
	o = flag.String("o", "output_%s.txt", "output file pattern")
)

var nlp = ling.MustNLP(ling.Norm)

func main() {
	flag.Parse()
	count, err := goutil.FileLineCount(*i)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(*i)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var br *bufio.Reader
	if strings.HasSuffix(strings.ToLower(*i), ".gz") {
		if gr, err := gzip.NewReader(f); err != nil {
			log.Fatal(err)
		} else {
			br = bufio.NewReader(gr)
		}
	} else {
		br = bufio.NewReader(f)
	}
	btk := topk.New(*k)
	ftk := topk.New(*k)
	bar := pb.StartNew(count)
	for {
		line, e := br.ReadString('\n')
		if e == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		items := strings.Fields(line)
		line = items[0]
		cnt := 1
		if *c {
			if len(items) > 1 {
				if n, err := strconv.Atoi(items[1]); err == nil {
					cnt = n
				}
			}
		}
		d := ling.NewDocument(line)
		if err := nlp.Annotate(d); err != nil {
			log.Printf("%s, %+v\n", line, err)
			continue
		}
		tokens := d.XRealTokens(ling.Norm)
		for i := 1; i < len(tokens); i++ {
			btk.Insert(strings.Join(tokens[i:], ""), cnt)
			ftk.Insert(strings.Join(tokens[:i], ""), cnt)
		}
		bar.Increment()
	}
	bar.FinishPrint("done!")
	output(ftk, fmt.Sprintf(*o, "prefix"))
	output(btk, fmt.Sprintf(*o, "suffix"))
}

func output(tk *topk.Stream, file string) {
	out, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	for _, v := range tk.Keys() {
		fmt.Fprintln(out, v.Key, v.Count, v.Error)
	}
}
