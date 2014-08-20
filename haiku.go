package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/mrjones/oauth"
)

var (
	consumerKey       = os.Getenv("HAIKU_CONSUMER_KEY")
	consumerSecret    = os.Getenv("HAIKU_CONSUMER_SECRET")
	accessTokenKey    = os.Getenv("HAIKU_ACCESSTOKEN_KEY")
	accessTokenSecret = os.Getenv("HAIKU_ACCESSTOKEN_SECRET")
)

func main() {
	c := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	for {
		accessToken := &oauth.AccessToken{Token: accessTokenKey, Secret: accessTokenSecret}
		haiku := generateHaiku()
		_, err := c.Post(
			"https://api.twitter.com/1.1/statuses/update.json",
			map[string]string{
				"status": haiku,
			},
			accessToken,
		)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Minute)
	}
}

var (
	one   = []string{}
	two   = []string{}
	three = []string{}
	four  = []string{}
	five  = []string{}
)

func init() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			//fmt.Println(err)
			break
		}
		line = strings.Trim(line, "\n")
		syll := strings.Split(line, "Ĩ")
		stripped := strings.Replace(line, "Ĩ", "", -1)
		//fmt.Println(stripped, len(syll))
		switch len(syll) {
		case 1:
			one = append(one, stripped)
		case 2:
			two = append(two, stripped)
		case 3:
			three = append(three, stripped)
		case 4:
			four = append(four, stripped)
		case 5:
			five = append(five, stripped)
		}
	}
}

func generateHaiku() string {
	ret := fmt.Sprintf("%s\n", getLine(5))
	ret += fmt.Sprintf("%s\n", getLine(7))
	ret += fmt.Sprintf("%s", getLine(5))
	return ret
}

func getLine(n int) string {
	return getLineHelper("", n)
}

func getLineHelper(ret string, left int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := r.Intn(left) + 1
	if x > 5 {
		x = 5
	}
	switch x {
	case 1:
		ret += one[r.Intn(len(one))] + " "
	case 2:
		ret += two[r.Intn(len(two))] + " "
	case 3:
		ret += three[r.Intn(len(three))] + " "
	case 4:
		ret += four[r.Intn(len(four))] + " "
	case 5:
		ret += five[r.Intn(len(five))] + " "
	}
	if left-x == 0 {
		return ret
	}
	return getLineHelper(ret, left-x)
}
