package composition

import (
	"github.com/silvalnk/orihon/internal/application"
	"github.com/silvalnk/orihon/internal/infrastructure/jsonshelf"
	"github.com/silvalnk/orihon/internal/infrastructure/mangadex"
)

func Wire(dir string) (*application.Service, *mangadex.Client, error) {
	shelf, err := jsonshelf.Open(dir)
	if err != nil {
		return nil, nil, err
	}
	client := mangadex.NewClient()
	return application.New(client, shelf), client, nil
}
