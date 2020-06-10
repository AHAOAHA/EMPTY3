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
Port=xxxx              # 数据库端口号
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
Path="/ws/geocoder/v1"
```
