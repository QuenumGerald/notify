package runner

import (
	"context"
	"notify/internal/config"
	"notify/internal/sink"
	"notify/internal/subscriber"
)

type Runner struct {
	Subs []config.Subscription
}

func (r *Runner) Start(ctx context.Context) error {
	for _, sub := range r.Subs {
		var s sink.Sink
		switch sub.Sink {
		case "stdout":
			s = &sink.StdoutSink{}
		case "slack":
			s = &sink.SlackSink{Webhook: sub.Webhook}
		case "telegram":
			s = &sink.TelegramSink{APIURL: sub.Webhook}
		default:
			s = &sink.StdoutSink{}
		}
		runner := &subscriber.SubscriptionRunner{Sub: sub, Sink: s}
		go runner.Run(ctx)
	}
	<-ctx.Done()
	return nil
}
