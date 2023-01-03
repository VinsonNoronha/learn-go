package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//format
// <rss version=”2.0">
// <channel>
// <title>Florida Lottery Winning Numbers</title>
// <link>http://www.flalottery.com</link>
// <description>Florida Lottery winning number drawing results for FLORIDA LOTTO, MEGA MILLIONS, POWERBALL, CASH4LIFE, JACKPOT TRIPLE PLAY, LUCKY MONEY, FANTASY 5, Midday and Evening PICK 2, PICK3, PICK 4 and PICK 5</description>
// <item game=”powerball” nextjp=”$140 Million” nextdd=”Saturday, March 21, 2020" winnum=”15–27–44–59–63 PB8 X4" windd=”Wednesday, March 18, 2020" myflv=”widget_pb.flv” winnumNM=”15–27–44–59–63 PB8" name=”POWERBALL”>
// <title>POWERBALL Drawing Results</title>
// <link>http://www.flalottery.com/powerball.do</link>
// <description>POWERBALL winning numbers are 15–27–44–59–63 PB8 X4 for 03/18/2020 and the next estimated Jackpot is $140 Million for 03/21/2020</description>
// <pubDate>Fri, 20 Mar 2020 01:30:30 EST</pubDate>
// <guid>http://www.flalottery.com/powerball.do</guid>
// </item>
// </channel>
// </rss>

type Rss struct {
	XMLName     xml.Name `xml:"rss"`
	Version     string   `xml:"version,attr"`
	Channel     Channel  `xml:"channel"`
	Description string   `xml:"description"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
 }

 type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Items       []Item   `xml:"item"`
 }

 type Item struct {
	XMLName  xml.Name `xml:"item"`
	Game     string   `xml:"game,attr"`
	NextJp   string   `xml:"nextjp,attr"`
	NextDd   string   `xml:"nextdd,attr"`
	WinNums  string   `xml:"winnum,attr"`
	WinDd    string   `xml:"windd,attr"`
	Myflv    string   `xml:"myflv,attr"`
	WinnumNM string   `xml:"winnumNM,attr"`
	Name     string   `xml:"name,attr"`   
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubdate"`
	Guid        string `xml:"guid"`
 }

func main() {
	// Open the xmlFile
	xmlFile, err := os.Open("data/winners.xml")
	// if os.Open returns an error then handle it
	if err != nil {
	   fmt.Println(err)
	   return
	}
	fmt.Println("\tSuccessfully opened winners.xml")
	// defer the closing of xmlFile so that we can parse it.
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)   // Unmarshal takes a []byte and fills the rss struct with the values found in the xmlFile
	rss := Rss{}
	xml.Unmarshal(byteValue, &rss)
	fmt.Println("Rss version: " + rss.Version)
	for _ , item := range rss.Channel.Items {
	   fmt.Println(item.Description)
	}


}