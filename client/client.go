package client

import (
	"action-cable-cli/helpers"
	"io/ioutil"
	"net/http"
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
			self.UIChan <- helpers.UIMsg{MsgType:helpers.UI_INFO, Msg: string(data)}
		}
	}
}

func (self *Client) WriteMessage(msg string) {
	self.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

