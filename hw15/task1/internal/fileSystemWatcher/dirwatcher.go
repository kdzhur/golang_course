package filesystemwatcher

import (
	"fmt"
	"log"
	"os"
	watcherstate "qeueu/task1/internal/watcherState"
	"qeueu/task1/pkg/pubsub"
)

type DirWatcher struct {
	path   string
	Broker *pubsub.Broker
	files  []FileWatcher
}

type FileWatcher struct {
	fileName string
	lmt      int // last modified time
	producer pubsub.Producer
}

func (f *FileWatcher) Watch() {
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
		}
	}()
}

func NewDirWatcher(path string) *DirWatcher {
	d := new(DirWatcher)
	broker := pubsub.NewBroker()
	state := watcherstate.NewGormState()
	state.InitDatabase(path)
	d.walkThroughDirs(path, broker, state)

	for i := range d.files {
		d.files[i].Watch()
	}

	broker.Accept()

	return &DirWatcher{
		path:   path,
		Broker: broker,
		files:  d.files,
	}
}

func (d *DirWatcher) walkThroughDirs(path string, broker *pubsub.Broker, state *watcherstate.GORMState) {
	var files []FileWatcher
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.Type().IsRegular() {
			fileName := path + "/" + e.Name()
			log.Println(fileName, "is ADDED to tracking")
			file := FileWatcher{
				fileName: fileName,
				producer: pubsub.Producer{Broker: broker},
			}
			file.AddFileToWatch(d, state)
		}
		if e.Type().IsDir() {
			dir := DirWatcher{
				path: path + "/" + e.Name(),
			}
			dir.AddDirToWatch(state)
			d.walkThroughDirs(path+"/"+e.Name(), broker, state)
		}
	}
	d.files = append(d.files, files...)
}

func (d *DirWatcher) AddDirToWatch(state *watcherstate.GORMState) {
	// Save to DB
	state.SaveDir(watcherstate.Dir{
		WatcherID: state.Watcher.ID,
		Path:      d.path,
	})
}
