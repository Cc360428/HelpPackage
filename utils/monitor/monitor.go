package main

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"github.com/fsnotify/fsnotify"
)

// 监听配置文件
func monitor() {
	s.Done()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logs.Error(err.Error())
		panic(err)
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
					logs.Info("创建文件 : ", ev.Name)
				}
				if ev.Op&fsnotify.Write == fsnotify.Write {
					logs.Info("写入文件 : ", ev.Name)
				}
				if ev.Op&fsnotify.Remove == fsnotify.Remove {
					logs.Info("删除文件 : ", ev.Name)
				}
				if ev.Op&fsnotify.Rename == fsnotify.Rename {
					logs.Info("重命名文件 : ", ev.Name)
				}
				if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
					logs.Info("修改权限 : ", ev.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logs.Error(err.Error())
			}
		}
	}()
	err = watcher.Add("utils/monitor/conf")
	if err != nil {
		logs.Error(err)
	}
	<-done
}
