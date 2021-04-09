package file

import (
	"github.com/fsnotify/fsnotify"
)

type watcherCollection map[string]*fsnotify.Watcher

func newWatcherCollection() *watcherCollection {
	watchers := watcherCollection(make(map[string]*fsnotify.Watcher))
	return &watchers
}

//add adds new watcher at the key, if it will be replace the previous one will closed
func (w *watcherCollection) add(key string, watcher *fsnotify.Watcher) bool {
	current, exists := (*w)[key]

	if current == watcher {
		return true
	}

	if exists && current != nil {
		current.Close()
	}

	(*w)[key] = watcher

	return exists
}

//clear clears the collection and closes all watchers
func (w *watcherCollection) clear() {
	w.closeAll()
	*w = *newWatcherCollection()
}

//closeAll closes all watchers
func (w *watcherCollection) closeAll() {
	for _, watcher := range *w {
		if watcher != nil {
			watcher.Close()
		}
	}
}
