package mode_test

import (
	"fmt"

	"github.com/rwxrob/live/mode"
)

func ExampleMode() {
	m := mode.New()
	m.SetName("coding")
	m.SetEmoji("💢")
	m.SetStatus("coding live stream modal utility")
	m.SetSummary("coding bash, go, c, or web")
	m.SetDescription(`Some longer description.`)
	m.SetTags([]string{"some", "tag"})
	m.SetTwitchCategory("1469308723")
	m.SetTwitchTags([]string{
		`15f4833a-1691-4cc1-a4a5-020d130ac94d`,
		`a59f1e4e-257b-4bd0-90c7-189c3efbf917`,
		`96b6073f-450d-4248-8ed4-988e28f3f759`,
	})
	fmt.Println(m)

	// Output:
	// {
	//   "Name": "coding",
	//   "Emoji": "💢",
	//   "Status": "coding live stream modal utility",
	//   "Summary": "coding bash, go, c, or web",
	//   "Description": "Some longer description.",
	//   "Tags": ["some","tag"],
	//   "TwitchCategory": "1469308723",
	//   "TwitchTags": [
	//     "15f4833a-1691-4cc1-a4a5-020d130ac94d",
	//     "a59f1e4e-257b-4bd0-90c7-189c3efbf917",
	//     "96b6073f-450d-4248-8ed4-988e28f3f759"
	//   ]
	// }

}
