package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Client struct {
	BaseUrl              string
	CurrentToken         Token
	Credential           Credential
	MarshalledCredential []byte
	TokenSuffix          string
}

type Credential struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	CreatedAt   int    `json:"created_at"`
}

func NewClient() *Client {

	var c = Client{
		BaseUrl:     "https://api.producthunt.com/v1/",
		TokenSuffix: "oauth/token",
	}
	var cred, err = c.GetCredential()
	if err != nil {
		fmt.Errorf("The correct environment variables could not be found.")
	}
	marshalled, _ := json.Marshal(cred)
	c.Credential = cred
	c.MarshalledCredential = marshalled
	token, err := c.GetAuthToken()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	c.CurrentToken = token
	return &c
}

func (c *Client) GetCredential() (Credential, error) {
	id := os.Getenv("PH_CLIENT_ID")
	secret := os.Getenv("PH_CLIENT_SECRET")
	var cred Credential
	if (id == "") || (secret == "") {
		return cred, fmt.Errorf("Missing environment key for the id or secret. Set PH_CLIENT_ID and PH_CLIENT_SECRET accordingly.")
	}
	cred = Credential{
		ClientId:     id,
		ClientSecret: secret,
		GrantType:    "client_credentials",
	}
	return cred, nil
}

func (c *Client) Get(url string) ([]byte, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(c.MarshalledCredential))
	req.Header.Set("Content-Type", "application/json")
	req.Close = true
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (c *Client) Post(url string) ([]byte, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(c.MarshalledCredential))
	req.Header.Set("Content-Type", "application/json")
	req.Close = true
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (c *Client) GetAuthToken() (Token, error) {
	url := c.BaseUrl + c.TokenSuffix
	rep, err := c.Post(url)

	var token Token
	if err != nil {
		return token, err
	}

	err = json.Unmarshal(rep, &token)
	token.CreatedAt = int(time.Now().Unix())
	return token, err
}

func main() {
	c := NewClient()
	c.GetAuthToken()
}
