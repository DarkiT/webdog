package context

import (
	"edboffical/webdog/config"
	"edboffical/webdog/message"
	"edboffical/webdog/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Cfg     config.BaseCfg
}

// get new context
func GetCtx(cfg config.BaseCfg) *Context {
	var ctx Context
	ctx.Cfg = cfg

	return &ctx
}

// set cfg
func (c *Context) SetCfg(cfg config.BaseCfg) {
	c.Cfg = cfg
}

// set http.ResponseWriter and *http.Request
func (c *Context) SetWr(w http.ResponseWriter, r *http.Request) {
	c.Writer = w
	c.Request = r
}

// command handler
func (c *Context) CommandHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c.SetWr(w, r)

	reqData := c.GetReqData()
	reqStr := ""
	for k, v := range reqData {
		reqStr += "--" + k + "=" + v[0] + " "
	}
	// command
	command := c.Cfg.Property
	if reqStr != "" {
		reqStr = utils.RemoveLastRune(reqStr)
		command = command + " " + reqStr
	}
	reqSlice := strings.Split(command, " ")
	out, err := exec.Command(reqSlice[0], reqSlice[1:]...).Output()
	if err != nil {
		log.Println("commandHandler error:", err.Error())
		// return error
		var resp message.Response
		resp.Code = -2
		resp.Msg = "command exec error:" + err.Error()
		data, _ := json.Marshal(resp)
		c.Render(data, "application/json")
		return
	}
	outSlice := strings.Split(string(out), c.Cfg.Split)

	tmpl, _ := template.New("test").Parse(c.Cfg.Resp)
	builder := &strings.Builder{}
	err = tmpl.Execute(builder, outSlice)
	if err != nil {
		var resp message.Response
		resp.Code = -2
		resp.Msg = "template error:" + err.Error()
		data, _ := json.Marshal(resp)
		c.Render(data, "application/json")
		return
	}
	c.Render([]byte(builder.String()), c.Cfg.Type)
}

// return response with json
func (c *Context) Render(data []byte, contentType string) {
	c.Writer.Header().Set("Content-Type", contentType)
	c.Writer.Write(data)
}

// get request data form http.Request
func (c *Context) GetReqData() map[string][]string {
	var data map[string][]string
	switch c.Request.Method {
	case "GET":
		return c.Request.URL.Query()
	case "POST":
		c.Request.ParseForm()
		return c.Request.PostForm
	default:
		return data
	}
}

// commonHandler
func (c *Context) CommonHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c.SetWr(w, r)

	c.Render([]byte(c.Cfg.Resp), c.Cfg.Type)
}

// contentHandler
func (c *Context) ContentHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c.SetWr(w, r)

	// filter ..
	if utils.ContainsDotDot(r.URL.Path) {
		var resp message.Response
		resp.Code = -3
		resp.Msg = "illegal request"
		data, _ := json.Marshal(resp)
		c.Render(data, "application/json")
		return
	}

	filePath := c.Cfg.Property + r.URL.Path
	if !utils.IsFileExisted(filePath) {
		var resp message.Response
		resp.Code = -2
		resp.Msg = "cant reach this resource"
		data, _ := json.Marshal(resp)
		c.Render(data, "application/json")
		return
	}

	file, _ := os.Open(filePath)
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	file.Read(buffer)
	contentType := http.DetectContentType(buffer)
	// reset the offset
	file.Seek(0, 0)
	// set the content-type header
	w.Header().Set("Content-Type", contentType)
	io.Copy(w, file)
}
