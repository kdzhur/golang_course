package filesystemwatcher

import (
	"fmt"
	"log"
	"os"
	"qeueu/task1/pkg/pubsub"
)

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
				fmt.Println("WATCH", f.fileName, f.lmt, lmt)
				f.producer.Publish(&pubsub.Message{
					Body: fmt.Sprintf("%v has been modified!", f.fileName),
				})
			}
		}
	}()
}

type DirWatcher struct {
	path   string
	Broker *pubsub.Broker
	files  []FileWatcher
}

func NewDirWatcher(path string) *DirWatcher {
	var files []FileWatcher
	broker := pubsub.NewBroker()

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.Type().IsRegular() {
			fileName := path + "/" + e.Name()
			fmt.Println(fileName, "ADDED")
			stat, _ := os.Stat(fileName)
			files = append(files, FileWatcher{
				fileName: fileName,
				lmt:      stat.ModTime().Nanosecond(),
				producer: pubsub.Producer{Broker: broker},
			})
		}
	}

	for i := range files {
		files[i].Watch()
	}

	broker.Accept()

	return &DirWatcher{
		path:   path,
		Broker: broker,
		files:  files,
	}
}

type Logger struct {
	cons *pubsub.Consumer
}

func NewLogger(broker *pubsub.Broker) *Logger {
	cons := pubsub.NewConsumer()
	broker.Subscribe(cons)

	return &Logger{
		cons: cons,
	}
}

func (l *Logger) NotifyOnModification() {
	l.cons.Consume()
}
