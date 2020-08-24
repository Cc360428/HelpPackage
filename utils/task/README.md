## Cron
>用户们可以通过cron在固定时间、日期、间隔下，运行定期任务（可以是命令和脚本）

### 坑
>linux 的应该对 cron 有所了解。linux 中可以通过 crontab -e 来配置定时任务。不过，linux 中的 cron 只能精确到分钟。
https://www.cnblogs.com/liuzhongchao/p/9521897.html
>而 github.com/robfig/cron 也按照操作系统进行判断，默认支持到分钟级别


### crontab 文件（文件格式）
#### 用户文件
>/var/spool/cron/
- 例子
```shell script
# 文件格式说明
# ┌──分钟（0 - 59）
# │  ┌──小时（0 - 23）
# │  │  ┌──日（1 - 31）
# │  │  │  ┌─月（1 - 12）
# │  │  │  │  ┌─星期（0 - 6，表示从周日到周六）
# │  │  │  │  │
# *  *  *  *  * 被执行的命令
```
- 注意⚠️
```shell script
注：
    在某些系统里，星期日也可以为7
    不很直观的用法：如果日期和星期同时被设定，那么其中的一个条件被满足时，指令便会被执行。请参考下例。
    前5个域称之分时日月周，可方便个人记忆。
```

#### 系统文件
>/etc/crontab和/etc/cron.d/
- 例子
```shell script
# 文件格式说明
# ┌──分钟（0 - 59）
# │  ┌──小时（0 - 23）
# │  │  ┌──日（1 - 31）
# │  │  │  ┌─月（1 - 12）
# │  │  │  │  ┌─星期（0 - 6，表示从周日到周六）
# │  │  │  │  │
# *  *  *  *  *  用户名 被执行的命令
```
### 表达式
```shell script
逗号（,）表示列举，例如： 1,3,4,7 * * * * echo hello world 表示，在每小时的1、3、4、7分时，打印"hello world"。
连词符（-）表示范围，例如：1-6 * * * * echo hello world ，表示，每小时的1到6分钟内，每分钟都会打印"hello world"。
星号（*）代表任何可能的值。例如：在“小时域”里的星号等于是“每一个小时”。
百分号(%) 表示“每"。例如：*%10 * * * * echo hello world 表示，每10分钟打印一回"hello world"。

1）星号(*)
表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月

2）斜线(/)
表示增长间隔，如第1个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后每隔 15 分钟执行一次（即 3、18、33、48 这些时间点执行），这里也可以表示为：3/15

3）逗号(,)
用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行

4）连字号(-)
表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）

5）问号(?)
只用于日(Day of month)和星期(Day of week)，\表示不指定值，可以用于代替 *
```
| 字段名            | 是否必须 | 允许的值        | 允许的特定字符 |
| :---------------- | :------- | :-------------- | :------------- |
| 秒(Seconds)       | 是       | 0-59            | * / , -        |
| 分(Minutes)       | 是       | 0-59            | * / , -        |
| 时(Hours)         | 是       | 0-23            | * / , -        |
| 日(Day of month)  | 是       | 1-31            | * / , – ?      |
| 月(Month)         | 是       | 1-12 or JAN-DEC | * / , -        |
| 星期(Day of week) | 否       | 0-6 or SUM-SAT  | * / , – ?      |

- 例子
```shell script
 #=================================================================
 #      SYSTEM ACTIVITY REPORTS
 #  8am-5pm activity reports every 20 mins during weekdays.
 #  activity reports every hour on Saturday and Sunday.
 #  6pm-7am activity reports every hour during weekdays.
 #  summary prepared at 18:05 every weekday.
 #=================================================================
 0,20,40 8-17 * * 1-5 /usr/lib/sa/sa1 1200 3 &
 0 * * * 0,6 /usr/lib/sa/sa1 &
 0 18-7 * * 1-5 /usr/lib/sa/sa1 &
 5 18 * * 1-5 /usr/lib/sa/sa2 -s 8:00 -e 18:01 -i 3600 -ubcwyaqvm &
```