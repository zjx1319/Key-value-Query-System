# Key-value-Query-System
一个简单的键值储存查询系统，使用Go echo开发

# 功能

- 储存键值对
- 查询键值对
- 删除键值对
- 生命周期等

# 使用
编译运行，默认端口为1323
- 查询
```
GET http://{IP}:1323/api/value/{key}
```

- 设置
```
POST http://{IP}:1323/api/value/{key}
Body:
{
    "key": {key},
    "value": {value},
    "life_time": <life_time>
}
```
life_time为可选参数，默认为永久

- 删除
```
DELETE http://{IP}:1323/api/value/{key}
```