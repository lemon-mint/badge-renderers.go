package main

import (
	"strings"

	"github.com/go-www/silverlining"
	badgerenderers "github.com/lemon-mint/badge-renderers.go"
	"github.com/lemon-mint/envaddr"
	"github.com/valyala/bytebufferpool"
)

func main() {
	err := silverlining.ListenAndServe(
		envaddr.Get(":8080"),
		func(r *silverlining.Context) {
			switch string(r.Path()) {
			case "/":
				r.ResponseHeaders().Set("Content-Type", "text/plain")
				r.WriteFullBodyString(200, "Badge API Server")
			case "/badge/forthebadge":
				r.ResponseHeaders().Set("Content-Type", "image/svg+xml")
				label, _ := r.GetQueryParamString("label")
				message, _ := r.GetQueryParamString("message")
				messageFillColor, _ := r.GetQueryParamString("color")
				messageTextColor, _ := r.GetQueryParamString("textcolor")
				labelFillColor, _ := r.GetQueryParamString("labelcolor")
				labelTextColor, _ := r.GetQueryParamString("labeltextcolor")

				if messageFillColor == "" {
					messageFillColor = "4bc51d"
				}

				if messageTextColor == "" {
					messageTextColor = "ffffff"
				}

				if labelFillColor == "" {
					labelFillColor = "555555"
				}

				if labelTextColor == "" {
					labelTextColor = "ffffff"
				}

				if message == "" {
					message = "INVALID MESSAGE"
					messageFillColor = "eb4511"
					messageTextColor = "ffffff"
					label = "ERROR"
					labelFillColor = "555555"
					labelTextColor = "ffffff"
				}

				label = strings.ToUpper(label)
				message = strings.ToUpper(message)

				r.ResponseHeaders().Set("Cache-Control", "max-age=86400, public")

				buf := bytebufferpool.Get()
				defer bytebufferpool.Put(buf)
				badgerenderers.WriteForTheBadge(buf, label, message, messageFillColor, messageTextColor, labelFillColor, labelTextColor)

				r.WriteFullBody(200, buf.Bytes())
			default:
				r.ResponseHeaders().Set("Content-Type", "text/plain")
				r.WriteFullBodyString(404, "Not Found")
			}
		},
	)
	if err != nil {
		panic(err)
	}
}
