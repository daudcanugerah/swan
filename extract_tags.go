package swan

import (
	"strings"

	"code.google.com/p/cascadia"
	"github.com/PuerkitoBio/goquery"
)

var (
	tagRelMatcher  = cascadia.MustCompile("a[rel=tag]")
	tagHrefMatcher = cascadia.MustCompile("a[href*='/tag/'], " +
		"a[href*='/tags/'], " +
		"a[href*='/topic/'], " +
		"a[href*='?keyword=']")
)

func extractTags(a *Article) error {
	tags := make(map[string]interface{})

	s := a.Doc.FindMatcher(tagRelMatcher)
	if s.Size() == 0 {
		s = a.Doc.FindMatcher(tagHrefMatcher)
	}

	s.Each(func(i int, s *goquery.Selection) {
		t := strings.TrimSpace(s.Text())
		if t != "" {
			tags[t] = nil
		}
	})

	for t := range tags {
		a.Meta.Tags = append(a.Meta.Tags, t)
	}

	return nil
}
