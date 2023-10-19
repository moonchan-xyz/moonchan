package accounts

import (
	"log"

	"github.com/moonchan-xyz/moonchan/back/mastodon/entities"
)

func UpdateCredentials(
	id string,
	displayName *string,
	note *string,
	avatar *entities.URLstring,
	header *entities.URLstring,
	locked *bool,
	bot *bool,
	discoverable *string,
	fieldsAttributes entities.Hash,
	sourcePrivacy *string,
	sourceSensitive *bool,
	sourceLanguage *string,
) (account *entities.Account, err error) {
	// todolog.Println(	id)
	if displayName != nil {
		log.Println(*displayName)
	}
	if note != nil {
		log.Println(*note)
	}
	if avatar != nil {
		log.Println(*avatar)
	}
	if header != nil {
		log.Println(*header)
	}
	if locked != nil {
		log.Println(*locked)
	}
	if bot != nil {
		log.Println(*bot)
	}
	if discoverable != nil {
		log.Println(*discoverable)
	}
	log.Println(fieldsAttributes)
	if sourcePrivacy != nil {
		log.Println(*sourcePrivacy)
	}
	if sourceSensitive != nil {
		log.Println(*sourceSensitive)
	}
	if sourceLanguage != nil {
		log.Println(*sourceLanguage)
	}
	return
}
