// Package p contains an HTTP Cloud Function.
package diplomacy_slack_bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var block = `{
	"blocks": [
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "This app does not do anything yet. The following link will send you to its Board Game Geek page."
			}
		},
		{
			"type": "section",
			"block_id": "section567",
			"text": {
				"type": "mrkdwn",
				"text": "<https://boardgamegeek.com/boardgame/483/diplomacy|Diplomacy> \n :star: :star: :star: :star: :star: \n At the turn of the 20th century, the seven Great European Powers engage in an intricate struggle for supremacy. Military forces invade and withdraw, shifting borders and altering empires with subtle maneuvers and daring gambits."
			},
			"accessory": {
				"type": "image",
				"image_url": "https://cf.geekdo-images.com/imagepagezoom/img/HStBe0j_SLNQxsC25YI_Io_MjZc=/fit-in/1200x900/filters:no_upscale()/pic288149.jpg",
				"alt_text": "Diplomacy Board Game"
			}
		},
		{
			"type": "section",
			"block_id": "section789",
			"fields": [
				{
					"type": "mrkdwn",
					"text": "*Average Rating*\nDiplomacy > 9000"
				}
			]
		}
	]
}`

type Request struct {
	// deprecated. you should not use this anymore
	Token string `json:"token,omitempty"`

	// The comment that was typed into the request trigger AKA the actual slash command
	Command string `json:"command,omitempty"`

	// The part after the slash command.
	Text string  `json:"text,omitempty"`

	// A temporary webhook URL that you can use to generate messages responses
	// usng https://api.slack.com/interactivity/handling#message_responses
	ResponseUrl string `json:"response_url,omitempty"`

	// A short-lived ID that will let your app open a modal
	TriggerId string `json:"trigger_id,omitempty"`

	//The ID of the user who triggerd the command
	UserId string `json:"user_id,omitempty"`

	// The plain text name of the user who triggered the command. This is being phased out
	// in favor of the user_id
	UserName string `json:"user_name,omitempty"`

	// Workspace ID where the command was triggered
	TeamId string `json:"team_id,omitempty"`
	TeamName string `json:"team_name,omitempty"`

	// Enterprise Grid ID
	EnterpriseId string `json:"enterprise_id,omitempty"`
	EnterpriseName string `json:"enterprise_name,omitempty"`

	// Channel ID where the command was triggered
	ChannelId string `json:"channel_id,omitempty"`
	ChannelName string `json:"channel_name,omitempty"`


}

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var request Request
	if e := json.NewDecoder(r.Body).Decode(&request); e != nil {
		fmt.Fprint(w, e)
	}

	log.Printf("incoming request: %q\n", request)
	fmt.Fprint(w, block)
}
