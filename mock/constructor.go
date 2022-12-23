package mock

import (
	"github.com/golang/mock/gomock"

	"github.com/conduitio-labs/notionapi"
)

func NewMockNotionClient(ctrl *gomock.Controller) *notionapi.Client {
	return &notionapi.Client{
		Database: NewMockDatabaseService(ctrl),
		Block:    NewMockBlockService(ctrl),
		Page:     NewMockPageService(ctrl),
		User:     NewMockUserService(ctrl),
		Search:   NewMockSearchService(ctrl),
		Comment:  NewMockCommentService(ctrl),
	}
}
