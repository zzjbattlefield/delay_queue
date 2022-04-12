# delay_queue

## 实现原理
> 利用Redis的有序集合，member为JobId, score为任务执行的时间戳,    
每一段时间扫描一次集合，取出执行时间小于等于当前时间的任务.   

## 依赖
* Redis

### 返回值
```json
{
  "code": 200,
  "message": "添加成功",
  "data": null
}
```

|  参数名 |     类型    |     含义     |        备注       |
|:-------:|:-----------:|:------------:|:-----------------:|
|   code  |     int     |    状态码    | |
| message |    string   | 状态描述信息 |                   |
|   data  | object, null |   附加信息   |                   |


### 添加任务   
URL地址 `/push`   
```json
{
    "topic":"phone",
    "id":"project_send_phone_99",
    "delay":120,
    "ttr":5,
    "body":"{\"uid\": 10829378,\"created\": 1498657365 }"
}
```
|  参数名 |     类型    |     含义     |        备注       |
|:-------:|:-----------:|:------------:|:-----------------:|
|   topic  | string     |    Job类型                   |                     |
|   id     | string     |    Job唯一标识                   | 需确保JobID唯一                  |
|   delay  | int        |    Job需要延迟的时间, 单位：秒    |                   |
|   ttr  | int        |    Job执行超时时间, 单位：秒   |                   |
|   body   | string     |    Job的内容，供消费者做具体的业务处理，如果是json格式需转义 |                   |

### 删除任务  
URL地址 `/delete`   

```json
{
  "id": "project_send_phone_99"
}
```

|  参数名 |     类型    |     含义     |        备注       |
|:-------:|:-----------:|:------------:|:-----------------:|
|   id  | string     |    Job唯一标识       |            |