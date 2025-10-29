package commandline

type Code int64

const (
	// Using iota to set int value for our const
	// No need to specify after first use
	Unknown Code = iota
    Exit
    LoadDictionary
    LoadText
    LoadWordList
    Stats
)

// Returns a string representation of the commandline code.
func (c Code) String() string {
	switch c {
	case Exit:
		return "exit"
	case LoadWordList:
		return "load wordlist"
	case LoadText:
		return "load text"
	case LoadDictionary:
		return "load dictionary"
	case Unknown:
		return "unknown"
	}
	// We should not reach this point
	panic ("Reaching unknown of the unknown.")
}