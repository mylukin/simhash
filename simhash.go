package simhash

import (
	gdSimhash "github.com/go-dedup/simhash"
	"math"
	"regexp"
)

type Simhash struct {
	gdSimhash.SimhashBase
}

type WordFeatureSet struct {
	gdSimhash.WordFeatureSet
}

func NewSimhash() *Simhash {
	return &Simhash{}
}

func (st *Simhash) NewWordFeatureSet(b []byte) *WordFeatureSet {
	fs := &WordFeatureSet{gdSimhash.WordFeatureSet{b}}
	fs.Normalize()
	return fs
}

var boundaries = regexp.MustCompile(`[\w']+(?:\://[\w\./]+){0,1}`)

func (w *WordFeatureSet) GetFeatures() []gdSimhash.Feature {
	var features []gdSimhash.Feature
	words := string(w.B)
	for _, w := range words {
		if len(string(w)) > 1 {
			feature := gdSimhash.NewFeature([]byte(string(w)))
			features = append(features, feature)
		}
	}
	bWords := boundaries.FindAll(w.B, -1)
	for _, w := range bWords {
		feature := gdSimhash.NewFeature(w)
		features = append(features, feature)
	}
	return features
}

func Compare(a uint64, b uint64) float64 {
	x := gdSimhash.Compare(a, b)
	y := - (1 / 2) * math.Pow(float64(x)/3, 2)
	y = math.Exp(y)
	y = (1 / math.Sqrt(2*math.Pi)) * y
	return y
}
