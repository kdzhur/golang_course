package editor

import (
	"hw_patterns/utils/filehandler"
	"os"
)

type File struct {
	Content []byte
	Config  *Config
}

type Config struct {
	Action Action
}

type Action interface {
	WorkOnText(f *File) []byte
}

func NewFile(file *os.File) *File {
	return &File{
		Content: filehandler.ReadAll(file),
		Config: &Config{
			Action: nil,
		},
	}
}

func (f *File) SetAction(a Action) {
	f.Config.Action = a
}

func (f *File) ApplyAction(decorators ...func(text string) string) []byte {
	result := string(f.Config.Action.WorkOnText(f))

	for _, fu := range decorators {
		result = fu(result)
	}
	return []byte(result)
}
