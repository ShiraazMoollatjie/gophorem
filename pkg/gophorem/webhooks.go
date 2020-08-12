package gophorem

import (
	"context"
	"fmt"
	"time"
)

// Webhooks will return a list of webhooks they have previously registered.
func (c *Client) Webhooks(ctx context.Context) (Webhooks, error) {
	var res Webhooks
	err := c.get(ctx, c.baseURL+"/webhooks", &res)

	return res, err
}

// Webhook will return a single webhook given its id.
func (c *Client) Webhook(ctx context.Context, id int) (*Webhook, error) {
	var res Webhook
	err := c.get(ctx, c.baseURL+fmt.Sprintf("/webhooks/%d", id), &res)

	return &res, err
}

// CreateWebhook will register HTTP endpoints that will be called once a relevant event is triggered inside the web application, events like article_created, article_updated.
func (c *Client) CreateWebhook(ctx context.Context, req CreateWebhookReq) (Webhook, error) {
	var res Webhook
	err := c.post(ctx, c.baseURL+"/webhooks", webhookReq{Webhook: req}, &res)

	return res, err
}

// DeleteWebhook will register HTTP endpoints that will be called once a relevant event is triggered inside the web application, events like article_created, article_updated.
func (c *Client) DeleteWebhook(ctx context.Context, id int) error {
	return c.delete(ctx, c.baseURL+fmt.Sprintf("/webhooks/%d", id), nil)
}

type Webhooks []Webhook

type Webhook struct {
	TypeOf    string    `json:"type_of"`
	ID        int       `json:"id"`
	Source    string    `json:"source"`
	TargetURL string    `json:"target_url"`
	Events    []string  `json:"events"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
}

type CreateWebhookReq struct {
	TargetURL string   `json:"target_url"`
	Source    string   `json:"source"`
	Events    []string `json:"events"`
}

type webhookReq struct {
	Webhook CreateWebhookReq `json:"webhook_endpoint"`
}
