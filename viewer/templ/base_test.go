package templ

import (
	"darkbot/configurator"
	"darkbot/scrappy"
	"darkbot/scrappy/base"
	"darkbot/scrappy/shared/records"
	"darkbot/settings"
	"darkbot/viewer/apis"
	"fmt"
	"os"
	"testing"
)

func TestBaseViewer(t *testing.T) {
	os.Remove(settings.Dbpath)
	channelID := "123"
	configurator.ConfiguratorChannel{Configurator: configurator.NewConfigurator()}.Add(channelID)

	cg := configurator.ConfiguratorBase{Configurator: configurator.NewConfigurator().Migrate()}
	cg.TagsAdd(channelID, []string{"Station"}...)

	bases := base.BaseStorage{}
	scrappy.Storage = &scrappy.ScrappyStorage{BaseStorage: &bases}
	record := records.StampedObjects[base.Base]{}.New()
	record.Add("Station1", base.Base{Name: "Station1", Affiliation: "Abc", Health: 100})
	record.Add("Station2", base.Base{Name: "Station2", Affiliation: "Qwe", Health: 100})
	bases.Add(record)

	base := TemplateBase{}
	base.API = apis.NewAPI(channelID)
	base.Render()
	fmt.Println(base.Content)
}

// func TestIntegrationTesting(t *testing.T) {
// 	os.Remove(settings.Dbpath)
// 	channelID := "838802002582175756"

// 	cg := configurator.ConfiguratorBase{Configurator: configurator.NewConfigurator()}
// 	cg.TagsAdd(channelID, []string{"Station"}...)

// 	scrappy.Storage.Update()

// 	base := BaseView{}
// 	base.ViewConfig = apis.NewAPI(channelID)
// 	base.Render()
// 	fmt.Println(base.Content)
// }