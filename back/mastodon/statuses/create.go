package statuses

import "github.com/moonchan-xyz/moonchan/back/mastodon/entities"

func Create(
	id,
	key,
	status string,
	mediaIDs,
	pollOptions []string,
	pollExpiresIn int,
	pollMutiple,
	pollHideTotals *bool,
	inReplyToID *string,
	sensitive *bool,
	spoilerText,
	visibility,
	language,
	scheduledAt *string,
) (*entities.Status, error) {
	return nil, nil
}
