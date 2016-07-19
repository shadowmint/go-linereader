package linereader

import (
	"bytes"
	"container/list"
)

// LineReader keeps track of individual lines reader from an arbitrary byte stream.
type LineReader struct {
	MaxLength int
	buffer    *bytes.Buffer
	copy      *bytes.Buffer
	lines     *list.List
}

// New returns a new empty LineReader
func New() *LineReader {
	return &LineReader{
		65535,
		bytes.NewBuffer([]byte{}),
		bytes.NewBuffer([]byte{}),
		list.New(),
	}
}

// Clear discards all held lines and any partial buffer
func (reader *LineReader) Clear() {
	reader.ForcePartial()
	reader.lines.Init()
}

// Write adds a new set of bytes to the line reader
func (reader *LineReader) Write(bytes []byte) {
	if (len(bytes) + reader.buffer.Len()) >= reader.MaxLength {

		// Push bytes in one chunk at a time until we're done.
		root := bytes
		for len(root) > 0 {

			// Bytes for this chunk
			cutoff := reader.MaxLength - reader.buffer.Len()
			if cutoff > len(root) {
				cutoff = len(root)
			}

			// Get chunk
			chunk := root[:cutoff]
			reader.write(chunk)

			// Update root
			root = root[cutoff:]
		}
	} else {
		reader.write(bytes)
	}
}

// write adds a new set of bytes to the line reader
func (reader *LineReader) write(bytes []byte) {
	reader.buffer.Write(bytes)
	reader.update()
}

// Len returns count of lines currently buffered.
func (reader *LineReader) Len() int {
	return reader.lines.Len()
}

// Next returns the next line, or error.
func (reader *LineReader) Next() string {
	if reader.Len() == 0 {
		return ""
	}
	rtn := reader.lines.Front()
	reader.lines.Remove(rtn)
	return rtn.Value.(string)
}

// ForcePartial pushes any partial uncompleted line into the lines buffer
func (reader *LineReader) ForcePartial() {
	if reader.buffer.Len() > 0 {
		value := string(reader.buffer.Bytes())
		reader.lines.PushBack(value)
		reader.buffer.Reset()
	}
}

// Look for any full lines in the buffer and update the lines buffer
func (reader *LineReader) update() {
	offset := 0
	b := reader.buffer.Bytes()
	index := bytes.IndexByte(b[offset:], '\n')
	for index != -1 {

		// Push this part into the strings array
		if index != 0 {
			value := string(b[offset : offset+index])
			reader.lines.PushBack(value)
		}

		// Update offset; if beyond end, break out
		offset = offset + index + 1
		if offset >= len(b) {
			break
		}

		// Find next
		index = bytes.IndexByte(b[offset:], '\n')
	}

	// If the offset moved, update the buffer block
	if offset != 0 {
		leftover := reader.buffer.Len() - offset
		if leftover > 0 {
			reader.copy.Reset()
			reader.copy.Write(b[offset:])
		}
		reader.buffer.Reset()
		if leftover > 0 {
			reader.buffer.Write(reader.copy.Bytes())
		}
	}

	// If this is a full buffer, push it
	if reader.buffer.Len() == reader.MaxLength {
		reader.ForcePartial()
	}
}
