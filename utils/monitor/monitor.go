package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// 监听配置文件
func monitor() {
	s.Done()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Panic(err.Error())
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case ev, ok := <-watcher.Events:
				if !ok {
					return
				}
				if ev.Op&fsnotify.Create == fsnotify.Create {
					log.Println("创建文件 : ", ev.Name)
				}
				if ev.Op&fsnotify.Write == fsnotify.Write {
					log.Println("写入文件 : ", ev.Name)
				}
				if ev.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("删除文件 : ", ev.Name)
				}
				if ev.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("重命名文件 : ", ev.Name)
				}
				if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
					log.Println("修改权限 : ", ev.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println(err.Error())
			}
		}
	}()
	err = watcher.Add("utils/monitor/conf")
	if err != nil {
		log.Println(err)
	}
	<-done
}
