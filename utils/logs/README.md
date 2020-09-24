## logs
## Cc 
```golang
    var err error
    uploadDir := "static/logger/" + time.Now().Format("2006/01/02/")
    err = os.MkdirAll(uploadDir, 777)
    if err != nil {
        logs.Error(err.Error())
    }
    err = logs.SetLogger(logs.AdapterMultiFile, `{"filename":"`+uploadDir+`/33.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
    if err != nil {
        logs.Error("日志输出错误!", err.Error())
    }
    //logs := logs.NewLogger()
    logs.Info("错误")
    //logs.Info("你是")
    logs.Error("error", errors.New("错误"))
```