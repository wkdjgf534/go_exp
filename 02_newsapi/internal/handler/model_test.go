package handler_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"newsapi/internal/handler"
	"newsapi/internal/store"
)

func TestNewsPostReqBody_Validate(t *testing.T) {
	type expectations struct {
		err  string
		news store.News
	}

	testCases := []struct {
		name         string
		req          handler.NewsPostReqBody
		expectations expectations
	}{
		{
			name: "author empty",
			req:  handler.NewsPostReqBody{},
			expectations: expectations{
				err: "author is empty",
			},
		},
		{
			name: "title empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
			},
			expectations: expectations{
				err: "title is empty",
			},
		},
		{
			name: "content empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
				Title:  "test-title",
			},
			expectations: expectations{
				err: "content is empty",
			},
		},
		{
			name: "summary empty",
			req: handler.NewsPostReqBody{
				Author:  "test-author",
				Title:   "test-title",
				Content: "test-content",
			},
			expectations: expectations{
				err: "summary is empty",
			},
		},
		{
			name: "time invalid",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Content:   "test-content",
				Summary:   "test-summary",
				CreatedAt: "invalid",
			},
			expectations: expectations{
				err: `parsing time "invalid"`,
			},
		},
		{
			name: "source invalid",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Content:   "test-content",
				Summary:   "test-summary",
				CreatedAt: "2025-01-01T06:06:06+00:00",
			},
			expectations: expectations{
				err: "source is empty",
			},
		},
		{
			name: "tags empty",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Content:   "test-content",
				Summary:   "test-summary",
				CreatedAt: "2025-01-01T06:06:06+00:00",
				Source:    "https://test-news.com",
			},
			expectations: expectations{
				err: "tags cannot be empty",
			},
		},
		{
			name: "validate",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Content:   "test-content",
				Summary:   "test-summary",
				CreatedAt: "2025-01-01T06:06:06+00:00",
				Source:    "https://test-news.com",
				Tags:      []string{"test-tag"},
			},
			expectations: expectations{
				news: store.News{
					Author:  "test-author",
					Title:   "test-title",
					Content: "test-content",
					Summary: "test-summary",
					Tags:    []string{"test-tag"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			news, err := tc.req.Validate()

			if tc.expectations.err != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectations.err)
			} else {
				assert.NoError(t, err)
				parsedTime, parseErr := time.Parse(time.RFC3339, tc.req.CreatedAt)
				require.NoError(t, parseErr)

				parsedSource, parseErr := url.Parse(tc.req.Source)
				require.NoError(t, parseErr)

				tc.expectations.news.CreatedAt = parsedTime
				tc.expectations.news.Source = parsedSource
				assert.Equal(t, tc.expectations.news, news)
			}
		})
	}
}
