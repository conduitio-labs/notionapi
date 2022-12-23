HELL           =/bin/bash -o pipefail
MOCKGEN_VERSION ?= v1.6.0

mockgen-install:
	go install github.com/golang/mock/mockgen@$(MOCKGEN_VERSION)

.PHONY: mockgen
mockgen: mockgen-install
		mockgen -source block.go -package mock > mock/mock_block_client.go
		mockgen -source comment.go -package mock > mock/mock_comment_client.go
		mockgen -source database.go -package mock > mock/mock_database_client.go
		mockgen -source page.go -package mock > mock/mock_page_client.go
		mockgen -source search.go -package mock > mock/mock_search_client.go
		mockgen -source user.go -package mock > mock/mock_user_client.go
