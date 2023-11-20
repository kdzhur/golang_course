package filesystemwatcher

import (
	"fmt"
	"os"
	watcherstate "qeueu/task1/internal/watcherState"
	"qeueu/task1/pkg/pubsub"
	"strings"
	"time"
)

type FileWatcher struct {
	fileName string
	lmt      int // last modified time
	producer pubsub.Producer
}

func (f *FileWatcher) Watch(rateLimit time.Duration) {
	go func() {
		for {
			stat, _ := os.Stat(f.fileName)
			lmt := stat.ModTime().Nanosecond()
			if f.lmt != lmt {
				f.lmt = lmt
				f.producer.Publish(&pubsub.Message{
					Body: fmt.Sprintf("%v has been modified!", f.fileName),
				})
			}
			time.Sleep(time.Second * rateLimit)
		}
	}()
}

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
