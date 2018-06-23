package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/cheggaaa/pb"
	"github.com/dgryski/go-topk"
	"github.com/liuzl/goutil"
	"github.com/liuzl/ling"
)

var (
	k = flag.Int("n", 10000, "k")
	i = flag.String("i", "input.txt.gz", "input file")
	o = flag.String("o", "output.txt", "output file")
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
	tk := topk.New(*k)
	bar := pb.StartNew(count)
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		d := ling.NewDocument(line)
		if err := nlp.Annotate(d); err != nil {
			log.Printf("%s, %+v\n", line, err)
			continue
		}
		tokens := d.XRealTokens(ling.Norm)
		for i := 1; i < len(tokens); i++ {
			tk.Insert(strings.Join(tokens[i:], ""), 1)
		}
		bar.Increment()
	}
	bar.FinishPrint("done!")
	fmt.Println(len(tk.Keys()))
	out, err := os.Create(*o)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	for _, v := range tk.Keys() {
		fmt.Fprintln(out, v.Key, v.Count, v.Error)
	}
}
