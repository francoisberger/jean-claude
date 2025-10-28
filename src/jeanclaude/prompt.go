package jeanclaude

type PromptCode int64

const (
	// Using iota to set int value for our const
	// No need to specify after first use
    Exit PromptCode = iota
    LoadDictionary
    LoadText
    LoadWordList
    Stats
)

func (p PromptCode) String() string {
	switch p {
	case Exit:
		return "exit"
	case LoadWordList:
		return "load wordlist"
	case LoadText:
		return "load text"
	case LoadDictionary:
		return "load dictionary"
	}
	return "unknown"
}

type Prompt struct {
	code PromptCode
	originalPrompt string
}

func (p Prompt) Code() PromptCode {
	return p.code
}

func (p Prompt) String() string {
	return p.Code().String()
}