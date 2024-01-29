package tsapi

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// OrderingType is a custom type for ordering options
type OrderingType string

// Define constants for each valid ordering option
const (
	LastUpdated    OrderingType = "last-updated"
	Newest         OrderingType = "newest"
	MostDownloaded OrderingType = "most-downloaded"
	TopRated       OrderingType = "top-rated"
)

type GlobalModView struct {
	ModAuthor  string `json:"mod_author"`
	ModName    string `json:"mod_name"`
	ModPicture string `json:"mod_picture"`
}

func GlobalListMods(ordering OrderingType, page int) ([]GlobalModView, error) {
	reqUrl := "https://thunderstore.io/c/lethal-company/?q=a&ordering=" + string(ordering) + "&section=mods&page=" + fmt.Sprintf("%d", page)
	response, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: %s", response.Status)
	}

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %w", err)
	}

	var mods []GlobalModView
	// Find and iterate over each mod container
	document.Find("div.col-md-4").Each(func(index int, element *goquery.Selection) {
		// Extract mod name
		modName := element.Find("div > h5").Text()

		// Extract mod image
		// div.col-md-4:nth-child(1) > div:nth-child(1) > a:nth-child(2) > img:nth-child(1)
		modPicture, _ := element.Find("div > a > img").Attr("src")

		// Extract mod author
		// div.col-md-4:nth-child(1) > div:nth-child(2) > div:nth-child(3) > a
		modAuthor := strings.Trim(element.Find("div:nth-child(2) > div:nth-child(3) > a").Text(), " \n")

		if modName != "" {
			mods = append(mods, GlobalModView{
				ModAuthor:  modAuthor,
				ModName:    modName,
				ModPicture: modPicture,
			})
		}
	})

	return mods, nil
}
