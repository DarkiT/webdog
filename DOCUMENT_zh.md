# webdog 使用文档

## 配置文件
在 webdog 可执行文件目录中，必须正确配置 `config.yml` 配置文件，可参照项目目录里的配置文件。

```yaml
---
normal:
  "/getInfo":
    mode: command
    method: GET
    split: "|"
    type: application/json
    property: "sh ./resource/resp.sh"
    resp: '{"os":{{index . 0}}, "who":{{index . 1}}}'
  "/getName":
    type: text/html; charset=utf-8
    resp: "webdog"
  "/static/*filename":
    mode: content
    method: GET
    property: "./resource"
server:
  port: 1995
```
- 字段 `server` 中配置 `port` ，这是 webdog 运行时监听的端口

- 字段 `normal` 中配置服务相关内容，该字段下键值名即为路由，如参考配置文件中的 `/getName`，在 webdog 运行后，你将可以通过 `http://localhost:1995/getName` 访问到该配置内容

## 具体配置项

- `mode` 字段说明：
`mode` 用来指定该路由配置以何种方式来执行，目前支持 `command` 和 `content`

### command 模式
用户的请求和参数将会通过 webdog 转发给 `property` 中配置的脚本

webdog 会将 GET 请求的参数和 POST 的 **PostForm** 数据以 **getopt** 长参数的形式传递给配置脚本，并期望脚本响应 `resp` 字段中配置的内容

在参考配置中，如果你访问 `/getInfo?name=callous&pass=123456`，webdog 最终获取响应的脚本即为 `sh ./resource/resp.sh --name=callous --pass=123456`

脚本在返回结果的时必须按照 `split` 的字符进行结果的分割，如 **结果1｜结果2** ，`resp` 中 **{{index . 0}}** 和 **{{index .1}}**，是固定的模版写法，前者会被**结果1**填充，后者会被**结果2**填充。

### content 模式
webdog 会将用户请求以文件服务的形式进行处理，在参考配置中，如果你访问 `/static/index.html` ，webdog 将会在 `property` 配置的路径中寻找对应资源。所以你最终将会访问到 `./resource/static/index.html`。`property` 中可以配置相对路径和绝对路径。

### 默认模式
如果你没有指定 `mode`，那么 webdog 会以给定的 `type` 作为 **Content-Type** ，返回 `resp` 中的内容。

## 启动
正确配置 `config.yml` 后直接运行 webdog 可执行文件即可

## 其他
- 路由匹配规则可参照 [httprouter](https://github.com/julienschmidt/httprouter)
- webdog 运行后，你可以随时修改 `config.yml` 而无需重启 webdog，配置将会动态生效