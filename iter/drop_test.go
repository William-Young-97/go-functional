package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/option"
)

func ExampleDrop() {
	counter := iter.Drop[int](iter.Count(), 2)
	fmt.Println(counter.Next().Unwrap())
	// Output: 2
}

func ExampleDrop_method() {
	counter := iter.Count().Drop(2)
	fmt.Println(counter.Next().Unwrap())
	// Output: 2
}

func TestDrop(t *testing.T) {
	counter := iter.Drop[int](iter.Count(), 2)
	assert.Equal(t, counter.Next().Unwrap(), 2)
}

func TestDropExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	iterator := iter.Drop[int](delegate, 5)

	assert.True(t, iterator.Next().IsNone())
	assert.True(t, iterator.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}

func TestDropExhaustedLater(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	iterator := iter.Drop[int](delegate, 1)

	delegate.NextReturnsOnCall(0, option.Some(42))
	delegate.NextReturnsOnCall(1, option.Some(43))
	delegate.NextReturnsOnCall(2, option.None[int]())

	assert.True(t, iterator.Next().IsSome())
	assert.True(t, iterator.Next().IsNone())
	assert.True(t, iterator.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 3)
}
