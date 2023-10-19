package accounts

import "github.com/moonchan-xyz/moonchan/back/mastodon/entities"

func Status(
	id,
	maxID,
	sinceID,
	minID string,
	limit int,
	onlyMedia,
	excludeReplies,
	excludeReblogs,
	pinned bool,
	tagged string,
) ([]entities.Status, error) {

	return nil, nil // todo
}
