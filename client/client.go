package client

import (
	"action-cable-cli/helpers"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Config       *Config
	response     *http.Response
	conn         *websocket.Conn
	UIChan chan helpers.UIMsg
}


var mutex sync.Mutex
var clientInstance *Client

func GetInstance() *Client  {
	mutex.Lock()
	defer mutex.Unlock()
	if clientInstance == nil {
		clientInstance = &Client{Config:&Config{}, UIChan:make(chan helpers.UIMsg)}
	}
	return clientInstance
}

func (self *Client) Connect() {
	var err error
	self.conn, self.response, err = websocket.DefaultDialer.Dial(self.Config.Url, nil)
	if err !=nil {
		self.UIChan <- helpers.UIMsg{MsgType:helpers.UI_INFO, Msg: err.Error()}
	}
	if self.response != nil {
		data, err2 := ioutil.ReadAll(self.response.Body)
		if err2 != nil {
			self.UIChan <- helpers.UIMsg{MsgType:helpers.UI_INFO, Msg: err2.Error()}
		}else{
			defer self.response.Body.Close()
			var sb strings.Builder
			for k, v := range self.response.Header {
				sb.WriteString(fmt.Sprintf("%s:%s\n", k,v))
			}
			sb.WriteString(fmt.Sprintf("\n%s", string(data)))
			self.UIChan <- helpers.UIMsg{MsgType:helpers.UI_INFO, Msg: sb.String(), Method: helpers.METHOD_REPLACE}
		}
	}
}

func (self *Client) WriteMessage(msg string) {
	self.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

