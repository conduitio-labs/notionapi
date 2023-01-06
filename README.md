# notionapi
[![Go Reference](https://pkg.go.dev/badge/https://github.com/jomei/notionapi.svg)](https://pkg.go.dev/https://github.com/jomei/notionapi)
[![Test](https://https://github.com/jomei/notionapi/actions/workflows/test.yml/badge.svg)](https://https://github.com/jomei/notionapi/actions/workflows/test.yml)

An API client for the [Notion API](https://developers.notion.com/) implemented in Golang

# Supported APIs
It supports all APIs for Notion API version `2022-06-28`

# Installation

```
$ go get https://github.com/jomei/notionapi
```

# Getting started
Follow Notion’s [getting started guide](https://developers.notion.com/docs/getting-started) to obtain an Integration Token.

## Example

Make a new `Client`

```go
import "https://github.com/jomei/notionapi"


client := notionapi.NewClient("your-integration-token")
```
Then, use client's methods to retrieve or update your content

```go
page, err := client.Page.Get(context.Background(), "your-page-id")
if err != nil {
	// do something
}
```
