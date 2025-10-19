Pachong — Go 爬虫项目

一个使用 **Golang** 编写的简易并发爬虫，用于练习 Go 的网络请求、正则解析、JSON 存储与模块化开发。

---

功能简介

* 支持抓取网页 HTML 内容
* 使用正则表达式提取 `<title>` 标签
* 结果以 JSON 格式保存到本地文件
* 代码模块化：分为 `fetch`、`parse`、`storage` 三个包
* 后续可拓展并发抓取与错误重试机制

---

项目结构

```
Pachong/
├── go.mod
├── main.go              # 程序入口
├── result.json          # 输出的结果文件
│
├── fetch/               # 网络请求模块
│   └── fetch.go
│
├── parse/               # HTML 解析模块
│   └── parse.go
│
└── storage/             # 数据存储模块
    └── storage.go
```
---

示例输出

### 控制台输出：

```
正在爬取: https://www.baidu.com
页面标题: 百度一下，你就知道
```

### result.json 文件内容：

```json
[
  {
    "url": "https://www.baidu.com",
    "title": "百度一下，你就知道"
  }
]
```

---

学到的知识点

* Go 模块化 (`go mod init`, `package`)
* 网络请求 (`net/http`)
* 数据读取 (`io.ReadAll`)
* 正则表达式 (`regexp`)
* JSON 编码解码 (`encoding/json`)
* 文件操作 (`os.WriteFile`)
* Goroutine 与 Channel（并发爬取可扩展）

---

作者

**lkwfej**
