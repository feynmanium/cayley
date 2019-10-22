package linkedql

import (
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/query"
	"github.com/cayleygraph/quad"
)

func init() {
	Register(&TagArray{})
	Register(&TagValue{})
	Register(&Value{})
}

// TagArray corresponds to .tagArray()
type TagArray struct {
	From ValueStep `json:"from"`
}

// Type implements Step
func (s *TagArray) Type() quad.IRI {
	return prefix + "TagArray"
}

// BuildIterator implements Step
func (s *TagArray) BuildIterator(qs graph.QuadStore) (query.Iterator, error) {
	fromIt, err := s.From.BuildValueIterator(qs)
	if err != nil {
		return nil, err
	}
	return &TagArrayIterator{fromIt}, nil
}

// TagValue corresponds to .tagValue()
type TagValue struct {
	From ValueStep `json:"from"`
}

// Type implements Step
func (s *TagValue) Type() quad.IRI {
	return prefix + "TagValue"
}

func singleValueIterator(it *ValueIterator) *ValueIterator {
	p := it.path.Limit(1)
	return NewValueIterator(p, it.namer)
}

// BuildIterator implements Step
func (s *TagValue) BuildIterator(qs graph.QuadStore) (query.Iterator, error) {
	it, err := s.From.BuildValueIterator(qs)
	if err != nil {
		return nil, err
	}
	return &TagArrayIterator{singleValueIterator(it)}, nil
}

// Value corresponds to .value()
type Value struct {
	From ValueStep `json:"from"`
}

// Type implements Step
func (s *Value) Type() quad.IRI {
	return prefix + "Value"
}

// BuildIterator implements Step
func (s *Value) BuildIterator(qs graph.QuadStore) (query.Iterator, error) {
	it, err := s.From.BuildValueIterator(qs)
	if err != nil {
		return nil, err
	}
	return singleValueIterator(it), nil
}