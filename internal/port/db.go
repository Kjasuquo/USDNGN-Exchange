package port

import "context"

type DB interface {
	SetNotificationStatus(ctx context.Context)
}
