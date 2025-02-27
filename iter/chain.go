package iter

import "github.com/BooleanCat/go-functional/option"

// ChainIter iterator, see [Chain].
type ChainIter[T any] struct {
	BaseIter[T]
	iterators     []Iterator[T]
	iteratorIndex int
}

// Chain instantiates a [*ChainIter] that will yield all items in the provided
// iterators to exhaustion first to last.
func Chain[T any](iterators ...Iterator[T]) *ChainIter[T] {
	iter := &ChainIter[T]{iterators: iterators}
	iter.BaseIter = BaseIter[T]{iter}
	return iter
}

// Next implements the [Iterator] interface.
func (iter *ChainIter[T]) Next() option.Option[T] {
	for {
		if iter.iteratorIndex == len(iter.iterators) {
			return option.None[T]()
		}

		if value, ok := iter.iterators[iter.iteratorIndex].Next().Value(); ok {
			return option.Some(value)
		}

		iter.iteratorIndex++
	}

}

var _ Iterator[struct{}] = new(ChainIter[struct{}])
