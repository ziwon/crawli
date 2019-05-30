package crawli

import (
	// "github.com/PuerkitoBio/goquery"

	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/ziwon/crawli/config"
)

const (
	appName = "crawli"
)

// default home directories for expected binaries
var (
	DefaultAppHome       = os.ExpandEnv("$HOME/.crawli")
	DefaultWorksheetHome = os.ExpandEnv("$HOME/.crawli/worksheets")
	DefaultDataHome      = os.ExpandEnv("$HOME/.crawli/data")
)

const (
	CommandVisit   = "visit"
	CommandCollect = "collect"
)

var (
	worksheets []*Worksheet
)

type Crawli struct {
	*colly.Collector
	worksheet *Worksheet
	data      map[string][]string
}

func NewCrawli(worksheet *Worksheet) *Crawli {
	c := initColly(worksheet)

	crawli := &Crawli{
		c,
		worksheet,
		make(map[string][]string),
	}

	crawli.OnHTML(worksheet.Task.Trigger, func(el *colly.HTMLElement) {
		crawli.collect(el, worksheet.Task.Columns)
	})

	return crawli
}

func initColly(worksheet *Worksheet) *colly.Collector {
	userAgent := appName
	if worksheet.Task.UserAgent != "" {
		userAgent = worksheet.Task.UserAgent
	}

	async := !(worksheet.Task.Async == 0)
	c := colly.NewCollector(
		colly.Async(async),
		colly.UserAgent(userAgent),
		colly.AllowedDomains(worksheet.Task.AllowedDomains...),
	)

	if worksheet.Task.Delay != 0 {
		fmt.Println("Adding delay: " + strconv.Itoa(worksheet.Task.Delay))
		c.Limit(&colly.LimitRule{
			DomainGlob: ".*",
			Delay:      time.Duration(worksheet.Task.Delay) * time.Second,
		})
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Attempting to load:", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Loaded page from:", r.Request.URL)
	})

	return c
}

func (c Crawli) collect(el *colly.HTMLElement, columns []*ColumnItem) {
	key := ""
	rows := make([]string, len(columns))
	for _, col := range columns {
		var value string
		switch col.Type {
		case "text":
			value = el.ChildText(col.Selector)
		case "attr":
			value = el.ChildAttr(col.Selector, col.Attr)
		}

		if col.Primary {
			key = value
		} else {
			rows = append(rows, value)
		}
	}
	c.data[key] = rows
}

func (c Crawli) Run() {
	fmt.Printf("Crawliing... %s", c.worksheet.Task.URL)
	err := c.Visit(c.worksheet.Task.URL)
	if err != nil {
		panic(err)
	}
}

func (c Crawli) Result() {
	fmt.Println(c.data)
}

func Collect() {
	fmt.Println("Collect...!")
	for _, worksheet := range worksheets {
		fmt.Println(worksheet)
		crawli := NewCrawli(worksheet)
		crawli.Run()
		crawli.Result()
	}

}

func init() {
	fmt.Println("init..")

	if config.Config().InConfig("default") {
		home := config.Config().GetString("default.home")
		worksheetRoot := path.Join(home, "worksheets")

		var worksheetFiles []string
		err := filepath.Walk(worksheetRoot,
			func(path string, file os.FileInfo, err error) error {
				if file.IsDir() != true {
					worksheetFiles = append(worksheetFiles, path)
				}
				return nil
			})

		if err != nil {
			panic(err)
		}

		parser := NewWorksheetParser()
		for _, file := range worksheetFiles {
			worksheet, err := parser.Parse(file)
			if err == nil {
				worksheets = append(worksheets, worksheet)
			}
		}
	}
}
