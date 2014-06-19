package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/message/mass"
	"net/http"
)

// 根据分组群发 ==================================================================

// 根据分组群发消息, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) msgMassSendByGroup(msg interface{}) (msgid int, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return
	}

	_url := clientMessageMassSendByGroupURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch msg.(type) {
		case *mass.GroupNews:
			err = fmt.Errorf("MsgMassSendNewsByGroup: %s", resp.Status)
			return
		case *mass.GroupText:
			err = fmt.Errorf("MsgMassSendTextByGroup: %s", resp.Status)
			return
		case *mass.GroupVoice:
			err = fmt.Errorf("MsgMassSendVoiceByGroup: %s", resp.Status)
			return
		case *mass.GroupImage:
			err = fmt.Errorf("MsgMassSendImageByGroup: %s", resp.Status)
			return
		case *mass.GroupVideo:
			err = fmt.Errorf("MsgMassSendVideoByGroup: %s", resp.Status)
			return
		default:
			err = fmt.Errorf("msgMassSendByGroup: %s", resp.Status)
			return
		}
	}

	var result struct {
		Error
		MsgId int `json:"msg_id"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = &result.Error
		return
	}
	msgid = result.MsgId
	return
}

// 根据分组群发图文消息.
func (c *Client) MsgMassSendNewsByGroup(msg *mass.GroupNews) (msgid int, err error) {
	if msg == nil {
		err = errors.New("MsgMassSendNewsByGroup: msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发文本消息.
func (c *Client) MsgMassSendTextByGroup(msg *mass.GroupText) (msgid int, err error) {
	if msg == nil {
		err = errors.New("MsgMassSendTextByGroup: msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发语音消息.
func (c *Client) MsgMassSendVoiceByGroup(msg *mass.GroupVoice) (msgid int, err error) {
	if msg == nil {
		err = errors.New("MsgMassSendVoiceByGroup: msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发图片消息.
func (c *Client) MsgMassSendImageByGroup(msg *mass.GroupImage) (msgid int, err error) {
	if msg == nil {
		err = errors.New("MsgMassSendImageByGroup: msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发视频消息.
func (c *Client) MsgMassSendVideoByGroup(msg *mass.GroupVideo) (msgid int, err error) {
	if msg == nil {
		err = errors.New("MsgMassSendVideoByGroup: msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据 OpenId 列表群发 ==========================================================

// 根据 OpenId列表 群发消息, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) msgMassSendByOpenId(msg interface{}) (msgid int, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return
	}

	_url := clientMessageMassSendByOpenIdURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch msg.(type) {
		case *mass.OpenIdNews:
			err = fmt.Errorf("MsgMassSendNewsByOpenId: %s", resp.Status)
			return
		case *mass.OpenIdText:
			err = fmt.Errorf("MsgMassSendTextByOpenId: %s", resp.Status)
			return
		case *mass.OpenIdVoice:
			err = fmt.Errorf("MsgMassSendVoiceByOpenId: %s", resp.Status)
			return
		case *mass.OpenIdImage:
			err = fmt.Errorf("MsgMassSendImageByOpenId: %s", resp.Status)
			return
		case *mass.OpenIdVideo:
			err = fmt.Errorf("MsgMassSendVideoByOpenId: %s", resp.Status)
			return
		default:
			err = fmt.Errorf("msgMassSendByOpenId: %s", resp.Status)
			return
		}
	}

	var result struct {
		Error
		MsgId int `json:"msg_id"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = &result.Error
		return
	}
	msgid = result.MsgId
	return
}

// 根据用户列表群发图文消息.
func (c *Client) MsgMassSendNewsByOpenId(msg *mass.OpenIdNews) (msgid int, err error) {
	if msg == nil {
		err = errors.New("MsgMassSendNewsByOpenId: msg == nil")
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发文本消息.
func (c *Client) MsgMassSendTextByOpenId(msg *mass.OpenIdText) (msgid int, err error) {
	if msg == nil {
		err = errors.New("MsgMassSendTextByOpenId: msg == nil")
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发语音消息.
func (c *Client) MsgMassSendVoiceByOpenId(msg *mass.OpenIdVoice) (msgid int, err error) {
	if msg == nil {
		err = errors.New("MsgMassSendVoiceByOpenId: msg == nil")
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发图片消息.
func (c *Client) MsgMassSendImageByOpenId(msg *mass.OpenIdImage) (msgid int, err error) {
	if msg == nil {
		err = errors.New("MsgMassSendImageByOpenId: msg == nil")
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发视频消息.
func (c *Client) MsgMassSendVideoByOpenId(msg *mass.OpenIdVideo) (msgid int, err error) {
	if msg == nil {
		err = errors.New("MsgMassSendVideoByOpenId: msg == nil")
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 删除群发 =====================================================================
//  NOTE: 只有已经发送成功的消息才能删除删除消息只是将消息的图文详情页失效，已经收到的用户，
//  还是能在其本地看到消息卡片。 另外，删除群发消息只能删除图文消息和视频消息，
//  其他类型的消息一经发送，无法删除。
func (c *Client) MsgMassDelete(msgid int) error {
	token, err := c.Token()
	if err != nil {
		return err
	}

	var deleteRequest struct {
		MsgId int `json:"msgid"`
	}
	deleteRequest.MsgId = msgid

	jsonData, err := json.Marshal(deleteRequest)
	if err != nil {
		return err
	}

	_url := clientMessageMassDeleteURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("MsgMassDelete: %s", resp.Status)
	}

	var result Error
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return &result
	}
	return nil
}