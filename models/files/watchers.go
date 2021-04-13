package files

import (
	"github.com/fsnotify/fsnotify"
)

//WatcherCollection is map with the path and the instance of watcher
type WatcherCollection map[string]*fsnotify.Watcher

//NewWatcherCollection creates a new instance of the collection of watchers
func NewWatcherCollection() *WatcherCollection {
	watchers := WatcherCollection(make(map[string]*fsnotify.Watcher))
	return &watchers
}

//Add adds new watcher at the key, if it will be replace the previous one will closed
func (w *WatcherCollection) Add(key string, watcher *fsnotify.Watcher) bool {
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

//Clear clears the collection and closes all watchers
func (w *WatcherCollection) Clear() {
	w.CloseAll()
	*w = *NewWatcherCollection()
}

//CloseAll closes all watchers
func (w *WatcherCollection) CloseAll() {
	for _, watcher := range *w {
		if watcher != nil {
			watcher.Close()
		}
	}
}

func (w *WatcherCollection) exists(path string) bool {
	_, exists := (*w)[path]
	return exists
}

func (w *WatcherCollection) get(path string) (*fsnotify.Watcher, bool) {
	watcher, exists := (*w)[path]
	return watcher, exists
}
