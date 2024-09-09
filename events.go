package coreclient

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/datarhei/core-client-go/v16/api"

	"github.com/goccy/go-json"
)

func (r *restclient) Events(ctx context.Context, filters api.EventFilters) (<-chan api.Event, error) {
	var buf bytes.Buffer

	e := json.NewEncoder(&buf)
	e.Encode(filters)

	header := make(http.Header)
	header.Set("Accept", "application/x-json-stream")
	header.Set("Connection", "close")

	stream, err := r.stream(ctx, "POST", "/v3/events", nil, header, "application/json", &buf)
	if err != nil {
		return nil, err
	}

	channel := make(chan api.Event, 128)

	go func(stream io.ReadCloser, ch chan<- api.Event) {
		defer stream.Close()
		defer close(channel)

		decoder := json.NewDecoder(stream)

		for decoder.More() {
			var event api.Event
			if err := decoder.Decode(&event); err == io.EOF {
				return
			} else if err != nil {
				event.Component = "error"
				event.Message = err.Error()
			}

			// Don't emit keepalives
			if event.Component == "keepalive" {
				continue
			}

			ch <- event

			if event.Component == "" || event.Component == "error" {
				return
			}
		}
	}(stream, channel)

	return channel, nil
}
