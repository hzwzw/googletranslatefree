package translategooglefree

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
)

func Translate(source, sourceLang, targetLang string) (string, error) {
	var text []string
	var result []interface{}

	myurl := "https://translate.google.cn/translate_a/single?client=gtx&sl=" +
	url.QueryEscape(sourceLang) + "&tl=" + url.QueryEscape(targetLang) + "&dt=t&q=" + url.QueryEscape(source)
	
	fmt.Println("%v", myurl)
	r, err := http.Get(myurl)
	if err != nil {
		return "err", errors.New("Error getting translate.google.cn")
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "err", errors.New("Error reading response body")
	}
	fmt.Println(body)
	bReq := strings.Contains(string(body), `<title>Error 400 (Bad Request)`)
	if bReq {
		return "err", errors.New("Error 400 (Bad Request)")
	}
	
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "err", errors.New("Error unmarshaling data")
	}
	fmt.Println(result)
	if len(result) > 0 {
		inner := result[0]
		for _, slice := range inner.([]interface{}) {
			for _, translatedText := range slice.([]interface{}) {
				text = append(text, fmt.Sprintf("%v", translatedText))
				fmt.Println(text)
				fmt.Println(translatedText)
				break
			}
		}
		cText := strings.Join(text, "")

		return cText, nil
	} else {
		return "err", errors.New("No translated data in responce")
	}
}
