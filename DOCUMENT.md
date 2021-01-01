# Webdog Documentation

## Configuration File
The `config.yml` configuration file must be properly configured in the webdog executable directory, which can be found in the project directory.

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
- Configure `port` in the field `server`, which is the port that webdog listens on

- Configure the service content in the field `normal`, the keyname under this field is the route, such as `/getName` in the reference configuration file, after webdog runs, you will be able to request the configuration content through `http://localhost:1995/getName`

## Specific Configuration

- `mode` description：
`mode` is used to specify the way that route configuration will be executed, currently `command` and `content` are supported

### command mode
User's request and parameters will be forwarded via webdog to the script configured in `property`

Webdog will pass the parameters of the GET request and the POST **PostForm** data to the configuration script as **getopt** long parameters and expects the script to respond with the content configured in the `resp` field

In the reference configuration, if you request `/getInfo?name=callous&pass=123456`, the command that webdog finally gets the response is `sh. /resource/resp.sh --name=callous --pass=123456`

The script must split the results according to the character of `split` when returning the results, such as **result 1｜result 2** , **{{index . 0}}** and **{{index .1}}** are fixed templates, the former will be filled by **result 1** and the latter will be filled by **result 2**.

### content mode
Webdog will process the user request as a file service，in the reference configuration, if you request `/static/index.html`, webdog will look for the corresponding resource in the path configured by `property`. So you will eventually to request `. /resource/static/index.html`. Relative and absolute paths can be configured in `property`.

### default mode
If you do not specify `mode`, then webdog returns the content in `resp` with the given `type` as **Content-Type**.

## Start Service
Execute the webdog executable directly after configuring `config.yml` correctly

## Others
- Routing matching rules can be found in [httprouter](https://github.com/julienschmidt/httprouter)
- Once webdog is running, you can change `config.yml` at any time without restarting webdog and the configuration will take effect dynamically