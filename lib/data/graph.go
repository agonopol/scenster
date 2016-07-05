package data

import (
	"github.com/qedus/osmpbf"
	"os"
	"runtime"
	"log"
	"fmt"
	"io"
)

type Graph struct {
	links map[string]*Link
}

func (g *Graph) link(id string) *Link {
	return nil
}

func (g *Graph) next(link *Link) []*Link {
	return []*Link{}
}

func (g *Graph) prev(link *Link) []*Link {
	return []*Link{}
}

func NewGraphFromPBF(f *os.File) *Graph {

	d := osmpbf.NewDecoder(f)
	err := d.Start(runtime.GOMAXPROCS(-1)) // use several goroutines for faster decoding
	if err != nil {
		log.Fatal(err)
	}

	var nc, wc, rc uint64

	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch v := v.(type) {
			case *osmpbf.Node:
				// Process Node v.
				nc++
			case *osmpbf.Way:
				// Process Way v.
				wc++
			case *osmpbf.Relation:
				// Process Relation v.
				rc++
			default:
				log.Fatalf("unknown type %T\n", v)
			}
		}
	}

	fmt.Printf("Nodes: %d, Ways: %d, Relations: %d\n", nc, wc, rc)
	return nil
}
