package live_test

import (
	"fmt"

	"github.com/rwxrob/live-go"
	"github.com/rwxrob/live-go/twitch"
)

func ExampleMode() {
	m := new(live.TwitchMode)

	m.Name = "coding"
	m.Emoji = '💢'
	m.Status = "coding live stream modal utility"
	m.Summary = "coding bash, go, c, or web"
	m.Description = `Some longer description.`
	m.TwitchCategory = twitch.Category{ID: "1469308723"}
	m.TwitchTags = []twitch.Tag{
		{ID: `15f4833a-1691-4cc1-a4a5-020d130ac94d`},
		{ID: `a59f1e4e-257b-4bd0-90c7-189c3efbf917`},
		{ID: `96b6073f-450d-4248-8ed4-988e28f3f759`},
	}

	fmt.Println(m)

	// Output:
	// {"Name":"coding","Emoji":128162,"Status":"coding live stream modal utility","Summary":"coding bash, go, c, or web","Description":"Some longer description.","TwitchCategory":{"ID":"1469308723"},"TwitchTags":[{"ID":"15f4833a-1691-4cc1-a4a5-020d130ac94d"},{"ID":"a59f1e4e-257b-4bd0-90c7-189c3efbf917"},{"ID":"96b6073f-450d-4248-8ed4-988e28f3f759"}]}

}
