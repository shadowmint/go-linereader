# linereader

Reads newline terminated lines from a byte stream with a maximum length.

Notice specifically this reader makes no distinction between lines which exceed `MaxLength` in
length, and lines of `MaxLength` with a terminating '\n'.

All lines are buffered internally until requested.

# Usage

    import "ntoolkit/linereader"

    reader := linereader.New()
    reader.MaxLength = 10

    reader.Write([]byte("123456789\n"))
    reader.Write([]byte("1234567890"))
    reader.Write([]byte("12345"))
    reader.Write([]byte("1234567890222222222233333333334444444444"))

    for reader.Len() > 0 {
      fmt.Printf("%v\n", reader.Next())
    }
