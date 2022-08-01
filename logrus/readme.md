## logrus
> 日志使用
### 日志输出格式
```go
logrus.Debug("Useful debugging information.")
logrus.Info("Something noteworthy happened!")
logrus.Warn("You should probably take a look at this.")
logrus.Error("Something failed but I'm not quitting.")
logrus.Fatal("Bye.")         //log之后会调用os.Exit(1)
logrus.Panic("I'm bailing.") //log之后会panic()
```
### 日志输出级别
[级别参考](https://github.com/sirupsen/logrus/blob/6699a89a232f3db797f2e280639854bbc4b89725/logrus.go#L91)