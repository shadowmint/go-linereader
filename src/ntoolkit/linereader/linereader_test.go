package linereader_test

import (
	"ntoolkit/assert"
	"ntoolkit/linereader"
	"testing"
)

func TestResolve(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		T.Assert(linereader.New() != nil)
	})
}

func TestWriteSimple(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		reader := linereader.New()
		reader.Write([]byte("Hello World\nHello World\nHello World"))
		T.Assert(reader.Len() == 2)

		reader.ForcePartial()

		T.Assert(reader.Len() == 3)
		T.Assert(reader.Next() == "Hello World")
		T.Assert(reader.Next() == "Hello World")
		T.Assert(reader.Next() == "Hello World")

		T.Assert(reader.Len() == 0)
		T.Assert(reader.Next() == "")
	})
}

func TestWriteComplexManyNewLines(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		reader := linereader.New()
		reader.Write([]byte("\n\nHello\n\nWorld\n\n"))
		T.Assert(reader.Len() == 2)
		T.Assert(reader.Next() == "Hello")
		T.Assert(reader.Next() == "World")
	})
}

func TestWriteComplexLongChunks(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		reader := linereader.New()
		reader.Write([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."))
		reader.Write([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."))
		reader.Write([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."))
		reader.Write([]byte("\nLorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.\n"))
		T.Assert(reader.Len() == 2)
	})
}

func TestMaxLenght(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		reader := linereader.New()
		reader.MaxLength = 10
		reader.Write([]byte("123456789\n"))
		reader.Write([]byte("1234567890"))
		reader.Write([]byte("12345"))
		reader.Write([]byte("1234567890222222222233333333334444444444"))
		T.Assert(reader.Len() == 6)
	})
}
