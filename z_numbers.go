// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package column

import "github.com/kelindar/bitmap"

// --------------------------- float32s ----------------------------

// columnFloat32 represents a generic column
type columnFloat32 struct {
	column
	data []float32 // The actual values
}

// makeFloat32s creates a new vector or float32s
func makeFloat32s() Column {
	return &columnFloat32{
		data: make([]float32, 0, 64),
		column: column{
			fill: make(bitmap.Bitmap, 0, 4),
		},
	}
}

// Grow grows the size of the column until we have enough to store
func (c *columnFloat32) Grow(idx uint32) {
	c.Lock()
	// TODO: also grow the bitmap
	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}
	c.Unlock()
}

// UpdateMany performs a series of updates at once
func (c *columnFloat32) UpdateMany(updates []Update) {
	c.Lock()
	defer c.Unlock()

	// Range over all of the updates, and depending on the operation perform the action
	for i, u := range updates {
		c.fill.Set(u.Index)
		switch u.Kind {
		case UpdatePut:
			c.data[u.Index] = u.Value.(float32)

		// If this is an atomic increment/decrement, we need to change the operation to
		// the final value, since after this update an index needs to be recalculated.
		case UpdateAdd:
			value := c.data[u.Index] + u.Value.(float32)
			c.data[u.Index] = value
			updates[i] = Update{
				Kind:  UpdatePut,
				Index: u.Index,
				Value: value,
			}
		}
	}
}

// Value retrieves a value at a specified index
func (c *columnFloat32) Value(idx uint32) (v interface{}, ok bool) {
	v = float32(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnFloat32) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnFloat32) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnFloat32) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// --------------------------- float64s ----------------------------

// columnFloat64 represents a generic column
type columnFloat64 struct {
	column
	data []float64 // The actual values
}

// makeFloat64s creates a new vector or float64s
func makeFloat64s() Column {
	return &columnFloat64{
		data: make([]float64, 0, 64),
		column: column{
			fill: make(bitmap.Bitmap, 0, 4),
		},
	}
}

// Grow grows the size of the column until we have enough to store
func (c *columnFloat64) Grow(idx uint32) {
	c.Lock()
	// TODO: also grow the bitmap
	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}
	c.Unlock()
}

// UpdateMany performs a series of updates at once
func (c *columnFloat64) UpdateMany(updates []Update) {
	c.Lock()
	defer c.Unlock()

	// Range over all of the updates, and depending on the operation perform the action
	for i, u := range updates {
		c.fill.Set(u.Index)
		switch u.Kind {
		case UpdatePut:
			c.data[u.Index] = u.Value.(float64)

		// If this is an atomic increment/decrement, we need to change the operation to
		// the final value, since after this update an index needs to be recalculated.
		case UpdateAdd:
			value := c.data[u.Index] + u.Value.(float64)
			c.data[u.Index] = value
			updates[i] = Update{
				Kind:  UpdatePut,
				Index: u.Index,
				Value: value,
			}
		}
	}
}

// Value retrieves a value at a specified index
func (c *columnFloat64) Value(idx uint32) (v interface{}, ok bool) {
	v = float64(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnFloat64) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnFloat64) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnFloat64) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// --------------------------- ints ----------------------------

// columnInt represents a generic column
type columnInt struct {
	column
	data []int // The actual values
}

// makeInts creates a new vector or ints
func makeInts() Column {
	return &columnInt{
		data: make([]int, 0, 64),
		column: column{
			fill: make(bitmap.Bitmap, 0, 4),
		},
	}
}

// Grow grows the size of the column until we have enough to store
func (c *columnInt) Grow(idx uint32) {
	c.Lock()
	// TODO: also grow the bitmap
	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}
	c.Unlock()
}

// UpdateMany performs a series of updates at once
func (c *columnInt) UpdateMany(updates []Update) {
	c.Lock()
	defer c.Unlock()

	// Range over all of the updates, and depending on the operation perform the action
	for i, u := range updates {
		c.fill.Set(u.Index)
		switch u.Kind {
		case UpdatePut:
			c.data[u.Index] = u.Value.(int)

		// If this is an atomic increment/decrement, we need to change the operation to
		// the final value, since after this update an index needs to be recalculated.
		case UpdateAdd:
			value := c.data[u.Index] + u.Value.(int)
			c.data[u.Index] = value
			updates[i] = Update{
				Kind:  UpdatePut,
				Index: u.Index,
				Value: value,
			}
		}
	}
}

// Value retrieves a value at a specified index
func (c *columnInt) Value(idx uint32) (v interface{}, ok bool) {
	v = int(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnInt) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnInt) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnInt) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// --------------------------- int16s ----------------------------

// columnInt16 represents a generic column
type columnInt16 struct {
	column
	data []int16 // The actual values
}

// makeInt16s creates a new vector or int16s
func makeInt16s() Column {
	return &columnInt16{
		data: make([]int16, 0, 64),
		column: column{
			fill: make(bitmap.Bitmap, 0, 4),
		},
	}
}

// Grow grows the size of the column until we have enough to store
func (c *columnInt16) Grow(idx uint32) {
	c.Lock()
	// TODO: also grow the bitmap
	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}
	c.Unlock()
}

// UpdateMany performs a series of updates at once
func (c *columnInt16) UpdateMany(updates []Update) {
	c.Lock()
	defer c.Unlock()

	// Range over all of the updates, and depending on the operation perform the action
	for i, u := range updates {
		c.fill.Set(u.Index)
		switch u.Kind {
		case UpdatePut:
			c.data[u.Index] = u.Value.(int16)

		// If this is an atomic increment/decrement, we need to change the operation to
		// the final value, since after this update an index needs to be recalculated.
		case UpdateAdd:
			value := c.data[u.Index] + u.Value.(int16)
			c.data[u.Index] = value
			updates[i] = Update{
				Kind:  UpdatePut,
				Index: u.Index,
				Value: value,
			}
		}
	}
}

// Value retrieves a value at a specified index
func (c *columnInt16) Value(idx uint32) (v interface{}, ok bool) {
	v = int16(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnInt16) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnInt16) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnInt16) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// --------------------------- int32s ----------------------------

// columnInt32 represents a generic column
type columnInt32 struct {
	column
	data []int32 // The actual values
}

// makeInt32s creates a new vector or int32s
func makeInt32s() Column {
	return &columnInt32{
		data: make([]int32, 0, 64),
		column: column{
			fill: make(bitmap.Bitmap, 0, 4),
		},
	}
}

// Grow grows the size of the column until we have enough to store
func (c *columnInt32) Grow(idx uint32) {
	c.Lock()
	// TODO: also grow the bitmap
	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}
	c.Unlock()
}

// UpdateMany performs a series of updates at once
func (c *columnInt32) UpdateMany(updates []Update) {
	c.Lock()
	defer c.Unlock()

	// Range over all of the updates, and depending on the operation perform the action
	for i, u := range updates {
		c.fill.Set(u.Index)
		switch u.Kind {
		case UpdatePut:
			c.data[u.Index] = u.Value.(int32)

		// If this is an atomic increment/decrement, we need to change the operation to
		// the final value, since after this update an index needs to be recalculated.
		case UpdateAdd:
			value := c.data[u.Index] + u.Value.(int32)
			c.data[u.Index] = value
			updates[i] = Update{
				Kind:  UpdatePut,
				Index: u.Index,
				Value: value,
			}
		}
	}
}

// Value retrieves a value at a specified index
func (c *columnInt32) Value(idx uint32) (v interface{}, ok bool) {
	v = int32(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnInt32) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnInt32) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnInt32) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// --------------------------- int64s ----------------------------

// columnInt64 represents a generic column
type columnInt64 struct {
	column
	data []int64 // The actual values
}

// makeInt64s creates a new vector or int64s
func makeInt64s() Column {
	return &columnInt64{
		data: make([]int64, 0, 64),
		column: column{
			fill: make(bitmap.Bitmap, 0, 4),
		},
	}
}

// Grow grows the size of the column until we have enough to store
func (c *columnInt64) Grow(idx uint32) {
	c.Lock()
	// TODO: also grow the bitmap
	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}
	c.Unlock()
}

// UpdateMany performs a series of updates at once
func (c *columnInt64) UpdateMany(updates []Update) {
	c.Lock()
	defer c.Unlock()

	// Range over all of the updates, and depending on the operation perform the action
	for i, u := range updates {
		c.fill.Set(u.Index)
		switch u.Kind {
		case UpdatePut:
			c.data[u.Index] = u.Value.(int64)

		// If this is an atomic increment/decrement, we need to change the operation to
		// the final value, since after this update an index needs to be recalculated.
		case UpdateAdd:
			value := c.data[u.Index] + u.Value.(int64)
			c.data[u.Index] = value
			updates[i] = Update{
				Kind:  UpdatePut,
				Index: u.Index,
				Value: value,
			}
		}
	}
}

// Value retrieves a value at a specified index
func (c *columnInt64) Value(idx uint32) (v interface{}, ok bool) {
	v = int64(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnInt64) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnInt64) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnInt64) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// --------------------------- uints ----------------------------

// columnUint represents a generic column
type columnUint struct {
	column
	data []uint // The actual values
}

// makeUints creates a new vector or uints
func makeUints() Column {
	return &columnUint{
		data: make([]uint, 0, 64),
		column: column{
			fill: make(bitmap.Bitmap, 0, 4),
		},
	}
}

// Grow grows the size of the column until we have enough to store
func (c *columnUint) Grow(idx uint32) {
	c.Lock()
	// TODO: also grow the bitmap
	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}
	c.Unlock()
}

// UpdateMany performs a series of updates at once
func (c *columnUint) UpdateMany(updates []Update) {
	c.Lock()
	defer c.Unlock()

	// Range over all of the updates, and depending on the operation perform the action
	for i, u := range updates {
		c.fill.Set(u.Index)
		switch u.Kind {
		case UpdatePut:
			c.data[u.Index] = u.Value.(uint)

		// If this is an atomic increment/decrement, we need to change the operation to
		// the final value, since after this update an index needs to be recalculated.
		case UpdateAdd:
			value := c.data[u.Index] + u.Value.(uint)
			c.data[u.Index] = value
			updates[i] = Update{
				Kind:  UpdatePut,
				Index: u.Index,
				Value: value,
			}
		}
	}
}

// Value retrieves a value at a specified index
func (c *columnUint) Value(idx uint32) (v interface{}, ok bool) {
	v = uint(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnUint) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnUint) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnUint) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// --------------------------- uint16s ----------------------------

// columnUint16 represents a generic column
type columnUint16 struct {
	column
	data []uint16 // The actual values
}

// makeUint16s creates a new vector or uint16s
func makeUint16s() Column {
	return &columnUint16{
		data: make([]uint16, 0, 64),
		column: column{
			fill: make(bitmap.Bitmap, 0, 4),
		},
	}
}

// Grow grows the size of the column until we have enough to store
func (c *columnUint16) Grow(idx uint32) {
	c.Lock()
	// TODO: also grow the bitmap
	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}
	c.Unlock()
}

// UpdateMany performs a series of updates at once
func (c *columnUint16) UpdateMany(updates []Update) {
	c.Lock()
	defer c.Unlock()

	// Range over all of the updates, and depending on the operation perform the action
	for i, u := range updates {
		c.fill.Set(u.Index)
		switch u.Kind {
		case UpdatePut:
			c.data[u.Index] = u.Value.(uint16)

		// If this is an atomic increment/decrement, we need to change the operation to
		// the final value, since after this update an index needs to be recalculated.
		case UpdateAdd:
			value := c.data[u.Index] + u.Value.(uint16)
			c.data[u.Index] = value
			updates[i] = Update{
				Kind:  UpdatePut,
				Index: u.Index,
				Value: value,
			}
		}
	}
}

// Value retrieves a value at a specified index
func (c *columnUint16) Value(idx uint32) (v interface{}, ok bool) {
	v = uint16(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnUint16) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnUint16) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnUint16) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// --------------------------- uint32s ----------------------------

// columnUint32 represents a generic column
type columnUint32 struct {
	column
	data []uint32 // The actual values
}

// makeUint32s creates a new vector or uint32s
func makeUint32s() Column {
	return &columnUint32{
		data: make([]uint32, 0, 64),
		column: column{
			fill: make(bitmap.Bitmap, 0, 4),
		},
	}
}

// Grow grows the size of the column until we have enough to store
func (c *columnUint32) Grow(idx uint32) {
	c.Lock()
	// TODO: also grow the bitmap
	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}
	c.Unlock()
}

// UpdateMany performs a series of updates at once
func (c *columnUint32) UpdateMany(updates []Update) {
	c.Lock()
	defer c.Unlock()

	// Range over all of the updates, and depending on the operation perform the action
	for i, u := range updates {
		c.fill.Set(u.Index)
		switch u.Kind {
		case UpdatePut:
			c.data[u.Index] = u.Value.(uint32)

		// If this is an atomic increment/decrement, we need to change the operation to
		// the final value, since after this update an index needs to be recalculated.
		case UpdateAdd:
			value := c.data[u.Index] + u.Value.(uint32)
			c.data[u.Index] = value
			updates[i] = Update{
				Kind:  UpdatePut,
				Index: u.Index,
				Value: value,
			}
		}
	}
}

// Value retrieves a value at a specified index
func (c *columnUint32) Value(idx uint32) (v interface{}, ok bool) {
	v = uint32(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnUint32) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnUint32) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnUint32) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// --------------------------- uint64s ----------------------------

// columnUint64 represents a generic column
type columnUint64 struct {
	column
	data []uint64 // The actual values
}

// makeUint64s creates a new vector or uint64s
func makeUint64s() Column {
	return &columnUint64{
		data: make([]uint64, 0, 64),
		column: column{
			fill: make(bitmap.Bitmap, 0, 4),
		},
	}
}

// Grow grows the size of the column until we have enough to store
func (c *columnUint64) Grow(idx uint32) {
	c.Lock()
	// TODO: also grow the bitmap
	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}
	c.Unlock()
}

// UpdateMany performs a series of updates at once
func (c *columnUint64) UpdateMany(updates []Update) {
	c.Lock()
	defer c.Unlock()

	// Range over all of the updates, and depending on the operation perform the action
	for i, u := range updates {
		c.fill.Set(u.Index)
		switch u.Kind {
		case UpdatePut:
			c.data[u.Index] = u.Value.(uint64)

		// If this is an atomic increment/decrement, we need to change the operation to
		// the final value, since after this update an index needs to be recalculated.
		case UpdateAdd:
			value := c.data[u.Index] + u.Value.(uint64)
			c.data[u.Index] = value
			updates[i] = Update{
				Kind:  UpdatePut,
				Index: u.Index,
				Value: value,
			}
		}
	}
}

// Value retrieves a value at a specified index
func (c *columnUint64) Value(idx uint32) (v interface{}, ok bool) {
	v = uint64(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnUint64) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnUint64) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnUint64) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}
