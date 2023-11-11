package filesystemwatcher

import (
	"fmt"
	"log"
	"os"
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

	d.walkThroughDirs(path, broker)

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

func (d *DirWatcher) walkThroughDirs(path string, broker *pubsub.Broker) {
	var files []FileWatcher
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.Type().IsDir() {
			d.walkThroughDirs(path+"/"+e.Name(), broker)
		}
		if e.Type().IsRegular() {
			fileName := path + "/" + e.Name()
			fmt.Println(fileName, " is ADDED to tracking")
			stat, _ := os.Stat(fileName)
			files = append(files, FileWatcher{
				fileName: fileName,
				lmt:      stat.ModTime().Nanosecond(),
				producer: pubsub.Producer{Broker: broker},
			})
		}
	}
	d.files = append(d.files, files...)
}
