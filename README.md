# GradeManager

## 简介

本文是基于Golang的学生成绩管理系统，项目基于B/S架构。

## 项目依赖
* golang 1.12+

## Permission level

|  功能列表  | 学生 | 教师 | 管理员 |
| :--------: | :--: | :--: | :----: |
|   读公告   |  √   |  √   |   √    |
|   写公告   |      |      |   √    |
|    登录    |  √   |  √   |   √    |
|   读课程   |  √   |  √   |   √    |
|   写课程   |      |      |   √    |
|   读成绩   |  √   |  √   |   √    |
|   写成绩   |      |  √   |   √    |

## 项目配置

```toml
# GradeManager/config/config.toml
# db info
[GradeManagerDB]
Host="xxx.xxx.xxx.xxx" # 数据库IP
Port=3306              # 数据库端口号
User="xxxxxx"          # 数据库用户名
Password="xxxxxx"      # 数据库密码
DataBaseName="xxxxxx"  # 数据库DataBase名

# 方糖服务url，用于服务告警
[Alarm]
Url="https://sc.ftqq.com/xxxxxxxxxxxxx.send"


# 腾讯地图API
[IpToAddr]
Key="xxxxxxx"
SK="xxxxxx"
Url="xxxxxx"
Path="xxxxxx"

[LocationToAddr]
Key="xxxxxx"
SK="xxxxxx"
Url="xxxxxx"
Path="xxxxxx"
```

## 项目展示

项目地址： [请点击此处](http://101.37.175.110:8080).

## API

IP地址转换：URL[101.37.175.110:8080/ip-to-addr?ip=xxx.xxx.xxx.xxx], [example](http://101.37.175.110:8080/ip-to-addr?ip=101.37.175.110).

坐标地址转换：URL[101.37.175.110:8080/location-to-addr?location=xx.xx.xx,xx.xx.xx], [example](http://101.37.175.110:8080/location-to-addr?location=30.32,45.65).
