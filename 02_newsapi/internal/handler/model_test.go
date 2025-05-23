package handler_test

import (
	"testing"

	"newsapi/internal/handler"
)

func TestNewsPostReqBody_Validate(t *testing.T) {
	testCases := []struct{
		name string
		req handler.NewsPostRequestBody
		expectedErr bool
	}{
		{
			name: "author empty",
			req: handler.NewsPostRequestBody{},
			expectedErr: true,
		},
		{
			name: "title empty",
			req: handler.NewsPostRequestBody{
				Author: "test-author",
			},
			expectedErr: true,
		},
		{
			name: "summary empty",
			req: handler.NewsPostRequestBody{
				Author: "test-author",
				Title:  "test-title",
			},
			expectedErr: true,
		},
		{
			name: "time invalid",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "invalid",
			},
			expectedErr: true,
		},
		{
			name: "source invalid",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
			},
			expectedErr: true,
		},
		{
			name: "tags empty",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
				Source:    "https://google.com",
			},
			expectedErr: true,
		},
		{
			name: "validate",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
				Source:    "https://gugle.com",
				Tags:      []string{"test-tag"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.req.Validate()

			if tc.expectedErr && err == nil {
				t.Fatal("expected error but got nil")
			}

			if !tc.expectedErr && err != nil {
				t.Fatal("expected nil but got error")
			}
		})
	}
}
