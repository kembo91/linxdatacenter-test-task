package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var stateMap = map[string]string{
	"AZ": "Arizona",
	"CA": "California",
	"ID": "Idaho",
	"IN": "Indiana",
	"MA": "Massachusetts",
	"OK": "Oklahoma",
	"PA": "Pennsylvania",
	"VA": "Virginia",
}

type dataItem struct {
	Item string `json:"item"`
}

//requestSchema is a struct to parse email data
type requestSchema struct {
	ReqType string     `json:"req_type"`
	Data    []dataItem `json:"data"`
}

//responseSchema is a struct to respond with
type responseSchema struct {
	ResType string `json:"res_type"`
	Result  string `json:"result"`
	Data    string `json:"data"`
}

func parseMail(r []dataItem) (map[string][]string, error) {
	m := make(map[string][]string)
	if len(r) == 0 {
		return m, fmt.Errorf("Empty \"Data field\"")
	}
	for _, v := range r {
		item := strings.ReplaceAll(v.Item, ",", "")
		splitItem := strings.Split(item, " ")
		state := splitItem[len(splitItem)-1]
		stateName := stateMap[state]
		if stateName == "" {
			return m, fmt.Errorf(`Wrong state %v`, state)
		}
		splitItem[len(splitItem)-1] = stateName
		m[stateName] = append(m[stateName], strings.Join(splitItem, " "))
	}
	return m, nil
}

func processMail(m map[string][]string) string {
	var resStr string
	for k, v := range m {
		resStr += k + "\n"
		for _, p := range v {
			resStr += "..... " + p + "\n"
		}
	}
	return resStr
}

func methodCheckPost(r *http.Request) error {
	if r.Method == http.MethodPost {
		return nil
	}
	return fmt.Errorf("This endpoint only accepts POST requests")
}

//EmailHandler handles incoming email data requests
func EmailHandler(w http.ResponseWriter, r *http.Request) {
	var d requestSchema
	err := methodCheckPost(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m, err := parseMail(d.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := processMail(m)
	response := responseSchema{
		ResType: d.ReqType,
		Data:    data,
		Result:  "success",
	}
	json.NewEncoder(w).Encode(response)
}
