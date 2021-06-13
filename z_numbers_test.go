// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package column

import (
	"testing"

	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
)

func TestOfFloat32s(t *testing.T) {
	c := makeFloat32s().(*columnFloat32)

	{ // Set the value at index
		c.Update(9, float32(99))
		assert.Equal(t, 10, len(c.data))
		assert.True(t, c.Contains(9))
	}

	{ // Get the values
		v, ok := c.Value(9)
		assert.Equal(t, float32(99), v)
		assert.True(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(99), f)
		assert.True(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(99), i)
		assert.True(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(99), u)
		assert.True(t, ok)
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Intersect(&other)
		assert.Equal(t, uint64(0b1000000000), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Difference(&other)
		assert.Equal(t, uint64(0xfffffffffffffdff), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Union(&other)
		assert.Equal(t, uint64(0xffffffffffffffff), other[0])
	}

	{ // Remove the value
		c.Delete(9)
		c.DeleteMany(&bitmap.Bitmap{0xffffffffffffffff})

		v, ok := c.Value(9)
		assert.Equal(t, float32(0), v)
		assert.False(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(0), f)
		assert.False(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(0), i)
		assert.False(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(0), u)
		assert.False(t, ok)
	}

	{ // Update several items at once
		c.UpdateMany([]Update{{1, float32(1)}, {2, float32(2)}})
		assert.True(t, c.Contains(1))
		assert.True(t, c.Contains(2))
	}

}

func TestOfFloat64s(t *testing.T) {
	c := makeFloat64s().(*columnFloat64)

	{ // Set the value at index
		c.Update(9, float64(99))
		assert.Equal(t, 10, len(c.data))
		assert.True(t, c.Contains(9))
	}

	{ // Get the values
		v, ok := c.Value(9)
		assert.Equal(t, float64(99), v)
		assert.True(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(99), f)
		assert.True(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(99), i)
		assert.True(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(99), u)
		assert.True(t, ok)
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Intersect(&other)
		assert.Equal(t, uint64(0b1000000000), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Difference(&other)
		assert.Equal(t, uint64(0xfffffffffffffdff), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Union(&other)
		assert.Equal(t, uint64(0xffffffffffffffff), other[0])
	}

	{ // Remove the value
		c.Delete(9)
		c.DeleteMany(&bitmap.Bitmap{0xffffffffffffffff})

		v, ok := c.Value(9)
		assert.Equal(t, float64(0), v)
		assert.False(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(0), f)
		assert.False(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(0), i)
		assert.False(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(0), u)
		assert.False(t, ok)
	}

	{ // Update several items at once
		c.UpdateMany([]Update{{1, float64(1)}, {2, float64(2)}})
		assert.True(t, c.Contains(1))
		assert.True(t, c.Contains(2))
	}

}

func TestOfInts(t *testing.T) {
	c := makeInts().(*columnInt)

	{ // Set the value at index
		c.Update(9, int(99))
		assert.Equal(t, 10, len(c.data))
		assert.True(t, c.Contains(9))
	}

	{ // Get the values
		v, ok := c.Value(9)
		assert.Equal(t, int(99), v)
		assert.True(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(99), f)
		assert.True(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(99), i)
		assert.True(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(99), u)
		assert.True(t, ok)
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Intersect(&other)
		assert.Equal(t, uint64(0b1000000000), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Difference(&other)
		assert.Equal(t, uint64(0xfffffffffffffdff), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Union(&other)
		assert.Equal(t, uint64(0xffffffffffffffff), other[0])
	}

	{ // Remove the value
		c.Delete(9)
		c.DeleteMany(&bitmap.Bitmap{0xffffffffffffffff})

		v, ok := c.Value(9)
		assert.Equal(t, int(0), v)
		assert.False(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(0), f)
		assert.False(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(0), i)
		assert.False(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(0), u)
		assert.False(t, ok)
	}

	{ // Update several items at once
		c.UpdateMany([]Update{{1, int(1)}, {2, int(2)}})
		assert.True(t, c.Contains(1))
		assert.True(t, c.Contains(2))
	}

}

func TestOfInt16s(t *testing.T) {
	c := makeInt16s().(*columnInt16)

	{ // Set the value at index
		c.Update(9, int16(99))
		assert.Equal(t, 10, len(c.data))
		assert.True(t, c.Contains(9))
	}

	{ // Get the values
		v, ok := c.Value(9)
		assert.Equal(t, int16(99), v)
		assert.True(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(99), f)
		assert.True(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(99), i)
		assert.True(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(99), u)
		assert.True(t, ok)
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Intersect(&other)
		assert.Equal(t, uint64(0b1000000000), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Difference(&other)
		assert.Equal(t, uint64(0xfffffffffffffdff), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Union(&other)
		assert.Equal(t, uint64(0xffffffffffffffff), other[0])
	}

	{ // Remove the value
		c.Delete(9)
		c.DeleteMany(&bitmap.Bitmap{0xffffffffffffffff})

		v, ok := c.Value(9)
		assert.Equal(t, int16(0), v)
		assert.False(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(0), f)
		assert.False(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(0), i)
		assert.False(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(0), u)
		assert.False(t, ok)
	}

	{ // Update several items at once
		c.UpdateMany([]Update{{1, int16(1)}, {2, int16(2)}})
		assert.True(t, c.Contains(1))
		assert.True(t, c.Contains(2))
	}

}

func TestOfInt32s(t *testing.T) {
	c := makeInt32s().(*columnInt32)

	{ // Set the value at index
		c.Update(9, int32(99))
		assert.Equal(t, 10, len(c.data))
		assert.True(t, c.Contains(9))
	}

	{ // Get the values
		v, ok := c.Value(9)
		assert.Equal(t, int32(99), v)
		assert.True(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(99), f)
		assert.True(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(99), i)
		assert.True(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(99), u)
		assert.True(t, ok)
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Intersect(&other)
		assert.Equal(t, uint64(0b1000000000), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Difference(&other)
		assert.Equal(t, uint64(0xfffffffffffffdff), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Union(&other)
		assert.Equal(t, uint64(0xffffffffffffffff), other[0])
	}

	{ // Remove the value
		c.Delete(9)
		c.DeleteMany(&bitmap.Bitmap{0xffffffffffffffff})

		v, ok := c.Value(9)
		assert.Equal(t, int32(0), v)
		assert.False(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(0), f)
		assert.False(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(0), i)
		assert.False(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(0), u)
		assert.False(t, ok)
	}

	{ // Update several items at once
		c.UpdateMany([]Update{{1, int32(1)}, {2, int32(2)}})
		assert.True(t, c.Contains(1))
		assert.True(t, c.Contains(2))
	}

}

func TestOfInt64s(t *testing.T) {
	c := makeInt64s().(*columnInt64)

	{ // Set the value at index
		c.Update(9, int64(99))
		assert.Equal(t, 10, len(c.data))
		assert.True(t, c.Contains(9))
	}

	{ // Get the values
		v, ok := c.Value(9)
		assert.Equal(t, int64(99), v)
		assert.True(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(99), f)
		assert.True(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(99), i)
		assert.True(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(99), u)
		assert.True(t, ok)
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Intersect(&other)
		assert.Equal(t, uint64(0b1000000000), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Difference(&other)
		assert.Equal(t, uint64(0xfffffffffffffdff), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Union(&other)
		assert.Equal(t, uint64(0xffffffffffffffff), other[0])
	}

	{ // Remove the value
		c.Delete(9)
		c.DeleteMany(&bitmap.Bitmap{0xffffffffffffffff})

		v, ok := c.Value(9)
		assert.Equal(t, int64(0), v)
		assert.False(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(0), f)
		assert.False(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(0), i)
		assert.False(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(0), u)
		assert.False(t, ok)
	}

	{ // Update several items at once
		c.UpdateMany([]Update{{1, int64(1)}, {2, int64(2)}})
		assert.True(t, c.Contains(1))
		assert.True(t, c.Contains(2))
	}

}

func TestOfUints(t *testing.T) {
	c := makeUints().(*columnUint)

	{ // Set the value at index
		c.Update(9, uint(99))
		assert.Equal(t, 10, len(c.data))
		assert.True(t, c.Contains(9))
	}

	{ // Get the values
		v, ok := c.Value(9)
		assert.Equal(t, uint(99), v)
		assert.True(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(99), f)
		assert.True(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(99), i)
		assert.True(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(99), u)
		assert.True(t, ok)
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Intersect(&other)
		assert.Equal(t, uint64(0b1000000000), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Difference(&other)
		assert.Equal(t, uint64(0xfffffffffffffdff), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Union(&other)
		assert.Equal(t, uint64(0xffffffffffffffff), other[0])
	}

	{ // Remove the value
		c.Delete(9)
		c.DeleteMany(&bitmap.Bitmap{0xffffffffffffffff})

		v, ok := c.Value(9)
		assert.Equal(t, uint(0), v)
		assert.False(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(0), f)
		assert.False(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(0), i)
		assert.False(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(0), u)
		assert.False(t, ok)
	}

	{ // Update several items at once
		c.UpdateMany([]Update{{1, uint(1)}, {2, uint(2)}})
		assert.True(t, c.Contains(1))
		assert.True(t, c.Contains(2))
	}

}

func TestOfUint16s(t *testing.T) {
	c := makeUint16s().(*columnUint16)

	{ // Set the value at index
		c.Update(9, uint16(99))
		assert.Equal(t, 10, len(c.data))
		assert.True(t, c.Contains(9))
	}

	{ // Get the values
		v, ok := c.Value(9)
		assert.Equal(t, uint16(99), v)
		assert.True(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(99), f)
		assert.True(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(99), i)
		assert.True(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(99), u)
		assert.True(t, ok)
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Intersect(&other)
		assert.Equal(t, uint64(0b1000000000), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Difference(&other)
		assert.Equal(t, uint64(0xfffffffffffffdff), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Union(&other)
		assert.Equal(t, uint64(0xffffffffffffffff), other[0])
	}

	{ // Remove the value
		c.Delete(9)
		c.DeleteMany(&bitmap.Bitmap{0xffffffffffffffff})

		v, ok := c.Value(9)
		assert.Equal(t, uint16(0), v)
		assert.False(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(0), f)
		assert.False(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(0), i)
		assert.False(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(0), u)
		assert.False(t, ok)
	}

	{ // Update several items at once
		c.UpdateMany([]Update{{1, uint16(1)}, {2, uint16(2)}})
		assert.True(t, c.Contains(1))
		assert.True(t, c.Contains(2))
	}

}

func TestOfUint32s(t *testing.T) {
	c := makeUint32s().(*columnUint32)

	{ // Set the value at index
		c.Update(9, uint32(99))
		assert.Equal(t, 10, len(c.data))
		assert.True(t, c.Contains(9))
	}

	{ // Get the values
		v, ok := c.Value(9)
		assert.Equal(t, uint32(99), v)
		assert.True(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(99), f)
		assert.True(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(99), i)
		assert.True(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(99), u)
		assert.True(t, ok)
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Intersect(&other)
		assert.Equal(t, uint64(0b1000000000), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Difference(&other)
		assert.Equal(t, uint64(0xfffffffffffffdff), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Union(&other)
		assert.Equal(t, uint64(0xffffffffffffffff), other[0])
	}

	{ // Remove the value
		c.Delete(9)
		c.DeleteMany(&bitmap.Bitmap{0xffffffffffffffff})

		v, ok := c.Value(9)
		assert.Equal(t, uint32(0), v)
		assert.False(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(0), f)
		assert.False(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(0), i)
		assert.False(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(0), u)
		assert.False(t, ok)
	}

	{ // Update several items at once
		c.UpdateMany([]Update{{1, uint32(1)}, {2, uint32(2)}})
		assert.True(t, c.Contains(1))
		assert.True(t, c.Contains(2))
	}

}

func TestOfUint64s(t *testing.T) {
	c := makeUint64s().(*columnUint64)

	{ // Set the value at index
		c.Update(9, uint64(99))
		assert.Equal(t, 10, len(c.data))
		assert.True(t, c.Contains(9))
	}

	{ // Get the values
		v, ok := c.Value(9)
		assert.Equal(t, uint64(99), v)
		assert.True(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(99), f)
		assert.True(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(99), i)
		assert.True(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(99), u)
		assert.True(t, ok)
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Intersect(&other)
		assert.Equal(t, uint64(0b1000000000), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Difference(&other)
		assert.Equal(t, uint64(0xfffffffffffffdff), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Union(&other)
		assert.Equal(t, uint64(0xffffffffffffffff), other[0])
	}

	{ // Remove the value
		c.Delete(9)
		c.DeleteMany(&bitmap.Bitmap{0xffffffffffffffff})

		v, ok := c.Value(9)
		assert.Equal(t, uint64(0), v)
		assert.False(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(0), f)
		assert.False(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(0), i)
		assert.False(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(0), u)
		assert.False(t, ok)
	}

	{ // Update several items at once
		c.UpdateMany([]Update{{1, uint64(1)}, {2, uint64(2)}})
		assert.True(t, c.Contains(1))
		assert.True(t, c.Contains(2))
	}

}
