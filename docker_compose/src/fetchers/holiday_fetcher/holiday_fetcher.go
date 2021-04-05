package holiday_fetcher

import (
	"github.com/anaskhan96/soup"
	"math/rand"
	"strings"
	"time"
)

const sourceURL = "https://prazdniki-segodnya.ru/"

type HolidayFetcherInterface interface {
	GetHolidaysList() []string
	GetFirstHoliday() string
	GetRandomHoliday() string
}

type HolidayFetcher struct {
	sourceURL string
	document  soup.Root
}

func (h *HolidayFetcher) getDocument() {
	url := h.sourceURL

	response, err := soup.Get(url)
	if err != nil {
		panic(err)
	}

	h.document = soup.HTMLParse(response)
}

var holidayFetcher *HolidayFetcher

func newHolidayFetcher() *HolidayFetcher {
	var fetcher HolidayFetcher

	fetcher.sourceURL = sourceURL
	fetcher.getDocument()

	return &fetcher
}

func GetHolidayFetcher() *HolidayFetcher {
	if holidayFetcher == nil {
		holidayFetcher = newHolidayFetcher()
	}

	return holidayFetcher
}

func (h *HolidayFetcher) elementToHoliday(element soup.Root) string {
	text := element.Text()
	return strings.TrimSpace(text)
}

func (h *HolidayFetcher) GetHolidaysList() []string {
	var accumulator []string
	root := h.document

	elements := root.FindAll("div", "class", "list-group-item")
	for _, element := range elements {
		holiday := h.elementToHoliday(element)
		accumulator = append(accumulator, holiday)
	}

	return accumulator
}

func (h *HolidayFetcher) GetFirstHoliday() string {
	root := h.document
	element := root.Find("div", "class", "list-group-item")

	return h.elementToHoliday(element)
}

func (h *HolidayFetcher) GetRandomHoliday() string {
	list := h.GetHolidaysList()
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(list))

	return list[index]
}
