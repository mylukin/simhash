package simhash

import (
	gdSimhash "github.com/go-dedup/simhash"
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
	//bWords := boundaries.FindAll(w.B, -1)
	//for _, w := range bWords {
	//	feature := gdSimhash.NewFeature(w)
	//	features = append(features, feature)
	//}
	return features
}
