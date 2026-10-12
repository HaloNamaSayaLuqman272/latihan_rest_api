package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	Exchange = "app.events"
	DLX      = "app.events.dx"
)

type Publisher interface {
	Publish(ctx context.Context, routingKey string, payload any) error
}

type HandlerFunc func(ctx context.Context, body []byte) error

type Client struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	mu   sync.Mutex
}

func New(url string) (*Client, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("dial rabbitmq: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("open channel: %w", err)
	}

	if err := ch.ExchangeDeclare(Exchange, "topic", true, false, false, false, nil); err != nil {
		conn.Close()
		return nil, err
	}

	if err := ch.ExchangeDeclare(DLX, "direct", true, false, false, false, nil); err != nil {
		conn.Close()
		return nil, err
	}

	return &Client{conn: conn, ch: ch}, nil
}

func (c *Client) Publish(ctx context.Context, routingKey string, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	c.mu.Lock() // amqp.Channel tidak aman dipakai banyak goroutine
	defer c.mu.Unlock()
	return c.ch.PublishWithContext(ctx, Exchange, routingKey, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		MessageId:    uuid.NewString(),
		Timestamp:    time.Now(),
		Body:         body,
	})
}

// consume memblokir sampai ctx selesai. Pesan yg gagal diproses masuk ke <queue>.dlq
func (c *Client) Consume(ctx context.Context, queue, routingKey string, h HandlerFunc) error {
	dlq := queue + ".dlq"
	if _, err := c.ch.QueueDeclare(dlq, true, false, false, false, nil); err != nil {
		return err
	}

	if err := c.ch.QueueBind(dlq, queue, DLX, false, nil); err != nil {
		return err
	}

	if _, err := c.ch.QueueDeclare(queue, true, false, false, false, amqp.Table{
		"x-dead-letter-exchange":    DLX,
		"x-dead-letter-routing-key": queue,
	}); err != nil {
		return err
	}

	if err := c.ch.QueueBind(queue, routingKey, Exchange, false, nil); err != nil {
		return err
	}

	if err := c.ch.Qos(10, 0, false); err != nil {
		return err
	}

	msgs, err := c.ch.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case m, ok := <-msgs:
			if !ok {
				return errors.New("rabbitmq channel closed")
			}
			if err := h(ctx, m.Body); err != nil {
				m.Nack(false, false) // ke DLQ
				continue
			}

			m.Ack(false)
		}
	}
}

func (c *Client) Close() error {
	c.ch.Close()
	return c.conn.Close()
}
