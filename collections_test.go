package atp

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type bitmapTestSuite struct {
	atpTestSuite
}

func (suite *bitmapTestSuite) TestAddLinear() {
	b := newEmptyBitmap(3)

	b.Add(1, nil)
	b.Add(2, nil)

	suite.Equal(uint32(1), b.bitmapData[0])
	suite.Equal(uint32(1), b.bitmapData[1])
	suite.Equal(uint32(0), b.bitmapData[2])
	suite.Equal(uint32(1), b.sequenceNumber)

}

func (suite *bitmapTestSuite) TestMoveTwice() {
	b := newEmptyBitmap(3)

	b.Add(1, nil)
	b.Add(2, nil)
	b.Add(3, nil)

	suite.Equal(uint32(1), b.bitmapData[0])
	suite.Equal(uint32(1), b.bitmapData[1])
	suite.Equal(uint32(1), b.bitmapData[2])

	suite.Equal(uint32(1), b.sequenceNumber)

}

func (suite *bitmapTestSuite) TestAddNonLinear() {
	b := newEmptyBitmap(3)
	b.Add(1, nil)
	b.Add(3, nil)

	suite.Equal(uint32(1), b.bitmapData[0])
	suite.Equal(uint32(0), b.bitmapData[1])
	suite.Equal(uint32(1), b.bitmapData[2])

	suite.Equal(uint32(1), b.sequenceNumber)
}

func (suite *bitmapTestSuite) TestToNumber() {
	b := newEmptyBitmap(3)
	b.Add(1, nil)
	b.Add(3, nil)

	suite.Equal(uint32(5), b.ToNumber())
}

func (suite *bitmapTestSuite) TestInit() {
	b := newBitmap(7, 0, 123)
	suite.Equal(7, len(b.bitmapData))
	suite.Equal(uint32(1), b.bitmapData[0])
	suite.Equal(uint32(1), b.bitmapData[1])
	suite.Equal(uint32(0), b.bitmapData[2])
	suite.Equal(uint32(1), b.bitmapData[3])
	suite.Equal(uint32(1), b.bitmapData[4])
	suite.Equal(uint32(1), b.bitmapData[5])
	suite.Equal(uint32(1), b.bitmapData[6])
}

func TestBitmap(t *testing.T) {
	suite.Run(t, new(bitmapTestSuite))
}

type queueTestSuite struct {
	atpTestSuite
}

func (suite *queueTestSuite) TestEmptyQueue(t *testing.T) {
	q := newQueue()
	suite.Equal(nil, q.Dequeue(), "Empty container value not == nil")
	suite.True(q.IsEmpty(), "IsEmpty() for empty container != true")
}

func (suite *queueTestSuite) TestWithMultipleEntries(t *testing.T) {
	q := newQueue()
	q.Enqueue(3)
	q.Enqueue(5)
	q.Enqueue(2)

	suite.False(q.IsEmpty(), "container with 3 elements shows as empty")
	suite.Equal(3, q.Dequeue().(int))
	suite.Equal(5, q.Dequeue().(int))
	suite.Equal(2, q.Dequeue().(int))
}

func (suite *queueTestSuite) TestPushFront(t *testing.T) {
	q := newQueue()
	q.Enqueue(3)
	q.PushFront(5)

	suite.False(q.IsEmpty(), "container with 2 elements shows as empty")
	suite.Equal(5, q.Dequeue().(int))
	suite.Equal(3, q.Dequeue().(int))
}

func (suite *queueTestSuite) TestPeekRemovesNothing(t *testing.T) {
	q := newQueue()
	q.Enqueue(3)

	suite.False(q.IsEmpty(), "container with 1 element shows as empty")
	suite.Equal(3, q.Peek().(int))
	suite.Equal(3, q.Dequeue().(int))
}

func (suite *queueTestSuite) TestCheckType(t *testing.T) {
	q1 := newQueue()
	q1.Enqueue(segment{sequenceNumber: []byte{0, 0, 0, 1}})
	test := func() { q1.Enqueue(1) }
	assert.Panics(t, test, "")
}

func TestQueue(t *testing.T) {
	suite.Run(t, new(bitmapTestSuite))
}
