## mikilin-go

mikilin-go是核查框架，该框架是java版本的 [核查框架Mikilin](https://github.com/simonAlong/Mikilin) 的go版本实现

## 下载

```shell
go get github.com/simonalong/mikilin-go
```

## 快速使用

这里举个例子，快速使用

```go
package main

import (
    "github.com/simonalong/mikilin-go"
    log "github.com/sirupsen/logrus"
)

type ValueBaseEntityOne struct {
    // 对属性修饰 
    Name string `match:"value=zhou"`
    Age  int
}

func main() {
    var value ValueBaseEntityOne
    var result bool
    var err string

    value = ValueBaseEntityOne{Name: "zhou"}

    // 核查
    result, err = mikilin.Check(value)
    if !result {
        log.Errorf(err)
    }

    value = ValueBaseEntityOne{Name: "chen"}
    // 核查
    result, err = mikilin.Check(value)
    if !result { 
        // 核查错误：属性 Name 的值 chen 不在只可用列表 [zhou] 中 
        log.Errorf(err)
    }
}

```

说明：<br/>

1. 这里提供方法Check，用于核查属性
2. 提供标签match，标签内容中提供匹配器：value，该匹配器表示匹配的具体的一些值

提示：<br/>
value除了匹配一个值，还可以修饰多个值，更多的功能，以及更多匹配器见更多功能

```go
type ValueBaseEntity struct {
    Name string `match:"value={zhou, 宋江}"`
    Age  int    `match:"value={12, 13}"`
}
```

## 更多功能

#### 匹配器

- value：匹配指定的值
- isBlank：值是否为空字符
- range：匹配数值的范围（最大值和最小值，用法是数学表达式）：数值（整数和浮点数）的大小、字符串的长度、时间的范围、时间的移动
- model：匹配指定的类型：
    - id_card：身份证
    - phone: 手机号
    - fixed_phone:固定电话
    - mail: 邮件地址
    - ip: ip地址
- condition：修饰的属性的表达式的匹配，提供#current和#root占位符，用于获取相邻属性的值
- regex：匹配正则表达式
- customize：匹配自定义的回调函数

#### 处理模块

- errMsg: 自定义的异常
- accept: 匹配后接受还是拒绝
- disable: 是否启用匹配，默认启用

### 1. 匹配器：value
匹配指定的一些值，可以修饰一个，也可以修饰多个值，可以修饰字符，也可修饰整数（int、int8、int16、int32、int64）、无符号整数（uint、uint8、uint16、uint32、uint64）、浮点数（float32、float64）、bool类型和string类型。<br/>

提示：
 - 中间逗号也可以为中文，为了防止某些手误写错为中文字符


```go
// 修饰一个值
type ValueBaseEntityOne struct {
    Name string `match:"value=zhou"`
    Age  int    `match:"value=12"`
}

// 修饰一个值
type ValueBaseEntity struct {
    Name string `match:"value={zhou, 宋江}"`
    Age  int    `match:"value={12, 13}"`
}
```

如果有自定义类型嵌套，则可以使用标签`check`，用于解析复杂结构
```go
type ValueInnerEntity struct {
    InnerName string `match:"value={inner_zhou, inner_宋江}"`
    InnerAge  int    `match:"value={2212, 2213}"`
}

type ValueStructEntity struct {
    Name string `match:"value={zhou, 宋江}"`
    Age  int    `match:"value={12, 13}"`

    Inner ValueInnerEntity `match:"check"`
}
```
修饰的结构可以有如下
- 自定义结构
- 数组/分片：对应类型只有为复杂结构才会核查
- map：其中的key和value类型只有是复杂结构才会核查

### 2. 匹配器：isBlank
### 3. 匹配器：range
### 4. 匹配器：model
### 5. 匹配器：condition
### 6. 匹配器：regex
### 7. 匹配器：customize
### 8. 自定义异常：errMsg
### 9. 匹配上接受/拒绝：accept
### 10. 启用：disable
