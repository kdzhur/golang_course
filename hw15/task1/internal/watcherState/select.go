package watcherstate

import "log"

func (g *GORMState) FindDir(path string) *Dir {
	var dir Dir

	finded := g.db.Where("Path = ?", path).Find(&dir)
	if finded.Error != nil {
		log.Fatalln(finded.Error)
	}

	return &dir
}
