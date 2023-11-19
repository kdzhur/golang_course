package filesystemwatcher

import (
	"os"
	watcherstate "qeueu/task1/internal/watcherState"
	"strings"
)

func (f *FileWatcher) AddFileToWatch(dir *DirWatcher, state *watcherstate.GORMState) {
	stat, _ := os.Stat(f.fileName)

	f.lmt = stat.ModTime().Nanosecond()
	dir.files = append(dir.files, *f)

	lastInd := strings.LastIndex(f.fileName, "/")
	// Save to DB
	currentDir := state.FindDir(f.fileName[:lastInd])
	state.SaveFile(watcherstate.File{
		FileName: f.fileName,
		Path:     f.fileName,
		DirID:    currentDir.ID,
		LMT:      stat.ModTime(),
	})
}
