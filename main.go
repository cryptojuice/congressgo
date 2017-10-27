package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	data, err := ioutil.ReadFile("MemberData.xml")
	if err != nil {
		fmt.Printf("error reading file: %v", err)
	}

	h := house{}
	err = xml.Unmarshal(data, &h)
	if err != nil {
		fmt.Printf("error unmarshaling xml: %v", err)
	}

	data, err = ioutil.ReadFile("senators_cfm.xml")
	if err != nil {
		fmt.Printf("error reading senators_cfm file: %v", err)
	}

	s := senate{}
	err = xml.Unmarshal(data, &s)
	if err != nil {
		fmt.Printf("eror unmarshaling xml: %v", err)
	}
}

type house struct {
	XMLName xml.Name `xml:"MemberData"`
	Members []struct {
		XMLName    xml.Name `xml:"member"`
		FirstName  string   `xml:"member-info>firstname"`
		LastName   string   `xml:"member-info>lastname"`
		Party      string   `xml:"member-info>party"`
		State      string   `xml:"member-info>state>state-fullname"`
		BioguideID string   `xml:"member-info>bioguideID"`
	} `xml:"members>member"`
}

type senate struct {
	XMLName xml.Name `xml:"contact_information"`
	Members []struct {
		XMLName    xml.Name `xml:"member"`
		FirstName  string   `xml:"first_name"`
		LastName   string   `xml:"last_name"`
		Party      string   `xml:"party"`
		State      string   `xml:"state"`
		BioguideID string   `xml:"bioguide_id"`
	} `xml:"member"`
}

func fetchRemoteHouseXML() ([]byte, error) {
	response, err := http.Get("http://clerk.house.gov/xml/lists/MemberData.xml")
	bytes, err := ioutil.ReadAll(response.Body)
	return bytes, err
}

// response, err := http.Get("http://clerk.house.gov/xml/lists/MemberData.xml")
func fetchRemoteXMLFile(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	bytes, err := ioutil.ReadAll(response.Body)
	return bytes, err
}
