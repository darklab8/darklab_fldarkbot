package viewer

import (
	"darkbot/app/discorder"
	"darkbot/app/settings/logus"
	"darkbot/app/settings/types"
	"darkbot/app/viewer/apis"
	"darkbot/app/viewer/templ"
	"strings"
	"time"
)

type ChannelView struct {
	api         apis.API
	BaseView    templ.TemplateBase
	Msgs        []discorder.DiscordMessage
	PlayersView templ.PlayersTemplates
	ChannelID   types.DiscordChannelID
}

func NewChannelView(dbpath types.Dbpath) ChannelView {
	view := ChannelView{}
	view.ChannelID = ""
	view.api = apis.NewAPI(view.ChannelID, dbpath)
	view.BaseView = templ.NewTemplateBase(view.ChannelID, dbpath)
	view.PlayersView = templ.NewTemplatePlayers(view.ChannelID, dbpath)

	return view
}

// Query all Discord messages
// Try to grab already sent message by ID, if yes, assign to found objects with message ID.
func (v *ChannelView) Setup(channelID types.DiscordChannelID) {
	v.ChannelID = channelID
	v.api.ChannelID = channelID
	v.BaseView.Setup(channelID)
	v.PlayersView.Setup(channelID)
}

// Query all Discord messages
// Try to grab already sent message by ID, if yes, assign to found objects with message ID.
func (v *ChannelView) Discover() error {
	logus.Info("viewer.Init.channelID=", logus.ChannelID(v.ChannelID))
	msgs, err := v.api.Discorder.GetLatestMessages(v.ChannelID)
	if logus.CheckWarn(err, "unable to grab latst msg", logus.ChannelID(v.ChannelID)) {
		return err
	}

	for _, msg := range msgs {
		v.BaseView.DiscoverMessageID(msg.Content, msg.ID)
		v.PlayersView.DiscoverMessageID(msg.Content, msg.ID)
	}

	v.Msgs = msgs

	return nil
}

// Render new messages (ensure preserved Message ID)
func (v *ChannelView) Render() {
	v.BaseView.Render()
	v.PlayersView.Render()
}

// Edit if message ID is present.
// Send if not present.
func (v ChannelView) Send() {
	v.BaseView.Send()
	v.PlayersView.Send()
}

func (v ChannelView) DeleteOld() {
	deleteLimit := 10
	for _, msg := range v.Msgs {

		if v.BaseView.MatchMessageID(msg.ID) {
			continue
		}

		if v.PlayersView.MatchMessageID(msg.ID) {
			continue
		}

		// forbidding to delete messages that aren't having their own template renderer
		if strings.Contains(msg.Content, templ.MsgViewHeader) {
			continue
		}

		if deleteLimit <= 0 {
			break
		}

		timeDiff := time.Now().Sub(msg.Timestamp)
		if timeDiff.Seconds() < 40 {
			continue
		}

		v.api.Discorder.DeleteMessage(v.ChannelID, msg.ID)
		logus.Info("deleted message with id", logus.MessageID(msg.ID))
		deleteLimit--
	}
}