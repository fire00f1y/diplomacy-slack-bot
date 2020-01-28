// Package p contains an HTTP Cloud Function.
package diplomacy_slack_bot

import (
	"fmt"
	"html"
	"net/http"
)

var block = `{
    "blocks": [
        {
            "type": "section",
            "text": {
                "type": "mrkdwn",
                "text": "*It's 80 degrees right now.*"
            }
        },
        {
            "type": "section",
            "text": {
                "type": "mrkdwn",
                "text": "Partly cloudy today and tomorrow"
            }
        }
    ]
}`

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, html.EscapeString(block))
}
