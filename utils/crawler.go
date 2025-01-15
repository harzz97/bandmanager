package utils

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/harzz97/bandmanager/models"
)

// FetchSongDetails fetches song details from the given URL
func FetchSongDetails(url string) (*models.Song, error) {
	// Fetch the HTML page
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the URL: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status: %s", res.Status)
	}

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the HTML: %v", err)
	}

	// Extract song details
	song := &models.Song{}

	// Extract title
	song.Title = strings.TrimSpace(doc.Find("title").Text())

	// Extract artist & lyrics
	var lyricsLines []string
	doc.Find(".lyric-text p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		if strings.HasPrefix(strings.ToLower(text), "singer") {
			// regex := `(?i)(?:Singer|Singers)\s*[:：]\s*([a-zA-Z\s,]+)`
			regex := `(?i)(?:Singer|Singers)\s*[:：]\s*([a-zA-Z\s.,]+)`
			re, err := regexp.Compile(regex)
			if err != nil {
				log.Fatal(err)
				// return song, err
			}

			// Find the match for singers
			var singers []string
			matches := re.FindStringSubmatch(text)
			if len(matches) > 1 {
				// Split the matched singers into CSV values
				singers = strings.Split(matches[1], ",")
				// Clean up extra spaces around names
				for i := range singers {
					singers[i] = strings.TrimSpace(singers[i])
				}
				// return singers, nil
			}

			song.Artist = strings.Join(singers, ",")
		} else if strings.HasPrefix(strings.ToLower(text), "music") {
			song.MusicBy = cleanText(strings.TrimPrefix(text, "Music by :"))
		} else if strings.HasPrefix(strings.ToLower(text), "lyrics by") {
			song.LyricsBy = cleanText(strings.TrimPrefix(text, "Lyrics by :"))
		} else {
			ln := extractEnglishLyrics(strings.TrimSpace(text))
			if len(ln) != 0 {
				lyricsLines = append(lyricsLines, ln)
			}
		}

	})
	song.Lyrics = strings.Join(lyricsLines, "\n")

	return song, nil
}

// Filter and extract only English lyrics from a mixture of languages
func extractEnglishLyrics(text string) string {
	lines := strings.Split(text, "\n")
	var englishLyrics []string
	for _, line := range lines {
		if strings.Contains(line, "Singer") || strings.Contains(line, "Music") || strings.Contains(line, "Lyrics") {
			continue
		}
		// Check if the line contains alphabetic characters (assuming English text)
		if containsEnglishChars(line) {
			englishLyrics = append(englishLyrics, line)
		}
	}
	return strings.Join(englishLyrics, "\n")
}

// Check if a line contains English characters
func containsEnglishChars(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}

// Clean and trim unnecessary text
func cleanText(text string) string {
	return strings.TrimSpace(text)
}
