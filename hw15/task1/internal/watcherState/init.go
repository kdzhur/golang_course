// Додати роботу з базою даних до домашньої роботи із домашнього завдання вебінару 15:
// • створити базу і щонайменше 3 таблиці;
// • у коді виконаного завдання підʼєднатися до бази;
// • кодом додавати записи до таблиць і вибирати їх.

package watcherstate

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Dir struct {
	gorm.Model
	WatcherID uint
	Path      string
	Files     []File `gorm:"foreignKey:DirID"`
}

type File struct {
	gorm.Model
	FileName string
	Path     string
	Dir      Dir
	DirID    uint
	LMT      time.Time
}

type Watcher struct {
	gorm.Model
	Dirs   []Dir  `gorm:"foreignKey:WatcherID"`
	Logger Logger `gorm:"foreignKey:WatcherID"`
}

type Logger struct {
	gorm.Model
	WatcherID uint
}

type GORMState struct {
	db      *gorm.DB
	RootDir *Dir
	Watcher *Watcher
	Redis   *RedisState
}

func NewGormState() *GORMState {
	db, err := gorm.Open(sqlite.Open("watcher.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Watcher{}, &Logger{}, &Dir{}, &File{})

	return &GORMState{
		db: db,
	}
}
func (g *GORMState) InitDatabase(rootPath string) {
	g.Watcher = g.SaveWatcher(&Watcher{})
	g.RootDir = g.SaveDir(&Dir{
		WatcherID: g.Watcher.ID,
		Path:      rootPath,
	})
	g.Redis = NewRedisState()
	g.Redis.Set(rootPath, 1)
}
