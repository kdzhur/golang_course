package watcherstate

import (
	"log"
)

func (g *GORMState) SaveFile(file *File) *File {
	saved := g.db.Save(file)
	if saved.Error != nil {
		log.Fatalln(saved.Error)
	}

	return file
}

func (g *GORMState) SaveWatcher(watcher *Watcher) *Watcher {
	saved := g.db.Save(watcher)
	if saved.Error != nil {
		log.Fatalln(saved.Error)
	}

	return watcher
}

func (g *GORMState) SaveDir(dir *Dir) *Dir {
	saved := g.db.Save(dir)
	if saved.Error != nil {
		log.Fatalln(saved.Error)
	}

	return dir
}

func (g *GORMState) SaveLogger(logger *Logger) *Logger {
	saved := g.db.Save(logger)
	if saved.Error != nil {
		log.Fatalln(saved.Error)
	}

	return logger
}
