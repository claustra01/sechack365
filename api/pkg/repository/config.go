package repository

import (
	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type NostrRelayRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *NostrRelayRepository) FindAll() ([]*model.NostrRelay, error) {
	var nostrRelays []*model.NostrRelay
	if err := r.SqlHandler.Select(&nostrRelays, `SELECT * FROM nostr_relays;`); err != nil {
		return nil, cerror.Wrap(err, "failed to get nostr relays")
	}
	return nostrRelays, nil
}

func (r *NostrRelayRepository) Create(url string) error {
	id := util.NewUuid()
	query := `
		INSERT INTO nostr_relays (id, url)
		VALUES ($1, $2);
	`
	if _, err := r.SqlHandler.Exec(query, id, url); err != nil {
		return cerror.Wrap(err, "failed to create nostr relay")
	}
	return nil
}

func (r *NostrRelayRepository) Delete(id string) error {
	query := `DELETE FROM nostr_relays WHERE id = $1;`
	if _, err := r.SqlHandler.Exec(query, id); err != nil {
		return cerror.Wrap(err, "failed to delete nostr relay")
	}
	return nil
}
