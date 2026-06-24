// Package webpush is a Web Push (VAPID) channel for togo notifications.
// Set VAPID_PUBLIC_KEY / VAPID_PRIVATE_KEY / VAPID_SUBJECT; notifications whose
// recipient RoutePushTokens() returns JSON webpush.Subscription strings get pushed.
package webpush

import (
	"context"
	"encoding/json"
	"os"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/togo-framework/notifications"
	"github.com/togo-framework/togo"
)

func init() {
	notifications.RegisterChannel("webpush", func(k *togo.Kernel) notifications.Channel {
		return &channel{pub: os.Getenv("VAPID_PUBLIC_KEY"), priv: os.Getenv("VAPID_PRIVATE_KEY"), sub: os.Getenv("VAPID_SUBJECT"), ttl: 30}
	})
}

type channel struct {
	pub, priv, sub string
	ttl            int
}

func (c *channel) Send(ctx context.Context, to notifications.Notifiable, n notifications.Notification) error {
	pn, ok := n.(notifications.PushNotification)
	if !ok || c.priv == "" {
		return nil
	}
	msg := pn.ToPush(to)
	payload, _ := json.Marshal(map[string]any{"title": msg.Title, "body": msg.Body, "data": msg.Data})
	for _, tok := range to.RoutePushTokens() {
		var s webpush.Subscription
		if json.Unmarshal([]byte(tok), &s) != nil || s.Endpoint == "" {
			continue
		}
		resp, err := webpush.SendNotification(payload, &s, &webpush.Options{Subscriber: c.sub, VAPIDPublicKey: c.pub, VAPIDPrivateKey: c.priv, TTL: c.ttl})
		if err != nil {
			return err
		}
		_ = resp.Body.Close()
	}
	return nil
}
