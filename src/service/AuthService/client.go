package AuthService

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	Aes "github.com/CHYinYiHang/gocron/pkg/Crypt_x"
	"github.com/CHYinYiHang/gocron/pkg/logging"
	"github.com/CHYinYiHang/gocron/pkg/utils"
)

type ClientInfo struct {
	ClientName string `json:"client_name";form:"client_name"`
}

//创建客户端
//TODO 创建客户端入口-server
func GenerateClient(clientName string) (clientInfo ClientInfo, err error) {
	client := ClientInfo{
		ClientName: clientName,
	}
	return client, nil
}

//生成客户端令牌标识
//TODO 生成客户端令牌-server
func (c *ClientInfo) GenerateToken() (token string, err error) {
	origData, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("JsonInitClientTokenError - " + err.Error())
	}

	key := utils.RandOmStr(24)
	nowData, err := Aes.EncryptAES(origData, []byte(key))
	if err != nil {
		return "", fmt.Errorf("AESInitClientTokenError - " + err.Error())
	}

	//AES加密后的数据转成base64编码
	token = base64.StdEncoding.EncodeToString(nowData)
	logging.Info(fmt.Sprintf("Client name is: '%s' Init success token is '%s' and key is '%s'", c.ClientName, token, key))
	return token, nil
}

//解析客户端令牌获取信息
//TODO 解析客户端令牌获取客户端信息-server
func ParseClientToken() (clientInfo ClientInfo, err error) {
	return
}
