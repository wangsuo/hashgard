# hashgardcli gov deposits

## 描述

查询指定提议的保证金详细情况

## 使用方式

```
  hashgardcli gov deposits [proposal-id] [flags]

```
打印帮助信息:

```
hashgardcli gov deposits --help
```
## 标志

| 名称, 缩写       | 默认值                      | 描述                                                                                                                                                 | 是否必须  |
| --------------- | -------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| --proposal-id   |                            | [string] 提议ID                                                                                                        | Yes      |

## 例子

### 查询所有保证金

```shell
hashgardcli gov deposits 1 --trust-node

```

你可以查询到指定提议的所有保证金代币，包括每个存款人的充值详情。

```txt
[
  {
    "depositor": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
    "proposal_id": "1",
    "amount": [
      {
        "denom": "apple",
        "amount": "50"
      }
    ]
  }
]

```