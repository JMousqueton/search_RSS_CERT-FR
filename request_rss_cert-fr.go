// Version : 1.0

package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"runtime"
)

// array of all Item in the url
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Items   []Item   `xml:"channel>item"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	PubDate string   `xml:"pubDate"`
	Desc    string   `xml:"description"`
}

func main() {

	url := "https://www.cert.ssi.gouv.fr/avis/feed/"

	resp, err := http.Get(url)

	// Open search keys file
	search_config_file, err := os.Open("search.conf")

	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// Close files
	defer resp.Body.Close()
	defer search_config_file.Close()

	// reads html as a slice of bytes
	xmlpage, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// initialize RSS array
	var items RSS
	// unmarshal xmlpage
	xml.Unmarshal(xmlpage, &items)

	// Split elements of the title with REGEX
	pattern := regexp.MustCompile(`(?P<AVI>([^\s]+)).+?\: (?P<title>(.*?)) \((?P<date>(.*?))\)`)

	// Split each line of file search_config_file
	scanner := bufio.NewScanner(search_config_file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	//List "item" one by one
	for i := 0; i < len(items.Items); i++ {
		//List "title" one by one
		str_title := items.Items[i].Title
		//Search each line of scanner in title heading
		for _, each_ln := range text {
			//re_search : REGEX of each line
			re_search := regexp.MustCompile(each_ln)
			match := re_search.Match([]byte(str_title))
			if match {
				//Enumerate each element of pattern
				for j := 1; j < 8; j++ {
					//Fix odds
					if j%2 == 0 {
						//Print each element of pattern with match and add ";" for future parsing
						str_title := items.Items[i].Title
						fmt.Print(pattern.FindStringSubmatch(str_title)[j], ";")
					}
				}
				//Add the link of the AVI number for research
				fmt.Print(items.Items[i].Link)
				println()
			}
		}
	}
	// Close CLI on Windows
	if runtime.GOOS == "windows" {
		fmt.Print("Press 'Enter' to close...")
		_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
