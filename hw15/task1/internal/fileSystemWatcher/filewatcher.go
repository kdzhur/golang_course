package filesystemwatcher

import (
	"log"
	"os"
	watcherstate "qeueu/task1/internal/watcherState"
	"strings"
)

func (f *FileWatcher) AddFileToWatch(dir *DirWatcher, state *watcherstate.GORMState) {
	stat, _ := os.Stat(f.fileName)

	f.lmt = stat.ModTime().Nanosecond()
	dir.files = append(dir.files, *f)

	lastInd := strings.LastIndex(f.fileName, "/")

	// GET from redis
	currentDirID := state.Redis.Get(f.fileName[:lastInd])
	if currentDirID == 0 {
		// GET from DB
		log.Println("Directory ID", currentDirID, "was not found in cache. Quering DB...")
		currentDir := state.FindDir(f.fileName[:lastInd])
		currentDirID = currentDir.ID
	} else {
		log.Println("Directory ID", currentDirID, "was found in cache")
	}

	// Save to DB
	state.SaveFile(&watcherstate.File{
		FileName: f.fileName,
		Path:     f.fileName,
		DirID:    currentDirID,
		LMT:      stat.ModTime(),
	})
}
