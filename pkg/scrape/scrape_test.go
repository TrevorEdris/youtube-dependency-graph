package scrape

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScrape_ScrapeNoErr(t *testing.T) {
	s := New()
	err := s.Scrape("https://fake.url/ytLinkLol")
	require.NoError(t, err, "s.Scrape produced an unexpected error")
}
