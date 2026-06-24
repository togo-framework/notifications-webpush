package webpush

import (
	"context"
	"testing"

	"github.com/togo-framework/notifications"
)

type nfy struct{ tokens []string }

func (n nfy) RouteID() string            { return "1" }
func (n nfy) RouteEmail() string         { return "" }
func (n nfy) RoutePushTokens() []string  { return n.tokens }

type pushNote struct{ t, b string }

func (pushNote) Via(notifications.Notifiable) []string { return []string{"webpush"} }
func (p pushNote) ToPush(notifications.Notifiable) notifications.PushMessage {
	return notifications.PushMessage{Title: p.t, Body: p.b}
}

type plainNote struct{}

func (plainNote) Via(notifications.Notifiable) []string { return []string{"webpush"} }

func TestNoConfigIsNoop(t *testing.T) {
	c := &channel{}
	if err := c.Send(context.Background(), nfy{tokens: []string{`{"endpoint":"x"}`}}, pushNote{t: "Hi", b: "There"}); err != nil {
		t.Fatalf("no-config should be a no-op, got %v", err)
	}
}

func TestIgnoresNonPushNotification(t *testing.T) {
	c := &channel{priv: "set"}
	if err := c.Send(context.Background(), nfy{tokens: []string{`{"endpoint":"x"}`}}, plainNote{}); err != nil {
		t.Fatalf("non-push notification should be ignored, got %v", err)
	}
}

func TestSkipsBadSubscription(t *testing.T) {
	c := &channel{priv: "set", pub: "set"}
	if err := c.Send(context.Background(), nfy{tokens: []string{"not-json", `{"endpoint":""}`}}, pushNote{t: "Hi", b: "B"}); err != nil {
		t.Fatalf("bad subscriptions should be skipped, got %v", err)
	}
}
