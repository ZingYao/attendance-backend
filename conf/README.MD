# 配置文件说明
```yaml
App:
  Addr: :8000   # 监听地址
  Name: 签到考勤系统 # 系统名
  Timeout: 10000 # 接口超时时间 单位 毫秒

Log:
  Level: debug # disable fatal error warn info debug
  Output: file # file console
  FilePath: ./log # 日志输出路径

Database:
  PostgreSql:
    Host: localhost # 地址 
    Port: 5432 # 端口
    User: attendance # 账号
    Pass: localpass # 密码
    DB: attendance # 库名
  Redis:
    Host: 6 # 地址
    Port: 7 # 端口
    Pass: 8 # 密码 无密码填写空字符串
```