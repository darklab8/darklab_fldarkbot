package viewer_msg

import (
	"github.com/darklab8/darklab_goutils/goutils/logus"
)

func LogusMsg(value *Msg) logus.SlogParam {
	return func(c *logus.SlogGroup) {
		c.Params["msg_message_id"] = string(value.messageID)
		c.Params["msg_view_id"] = string(value.viewID)
		c.Params["msg_view_enumerated_id"] = string(value.viewEnumeratedID)
	}
}