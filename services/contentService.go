package services

import (
	cm "UlboraCmsAdmin/common"
	b64 "encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

//ContentService service
type ContentService struct {
	Token     string
	ClientID  string
	APIClient string
	APIKey    string
	UserID    string
	Hashed    string
	Host      string
}

//Content content
type Content struct {
	ID                int64     `json:"id"`
	Title             string    `json:"title"`
	Category          string    `json:"category"`
	CreateDate        time.Time `json:"createDate"`
	ModifiedDate      time.Time `json:"modifiedDate"`
	UseModifiedDate   bool      `json:"useModifiedDate"`
	Hits              int64     `json:"hits"`
	MetaAuthorName    string    `json:"metaAuthorName"`
	MetaDesc          string    `json:"metaDesc"`
	MetaKeyWords      string    `json:"metaKeyWords"`
	MetaRobotKeyWords string    `json:"metaRobotKeyWords"`
	Text              string    `json:"text"`
	TextHTML          template.HTML
	SortOrder         int   `json:"sortOrder"`
	Archived          bool  `json:"archived"`
	ClientID          int64 `json:"clientId"`
}

// PageHead used for page head
type PageHead struct {
	Title        string
	MetaAuthor   string
	MetaDesc     string
	MetaKeyWords string
}

//Response res
type Response struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
	Code    int   `json:"code"`
}

//AddContent add content
func (c *ContentService) AddContent(content *Content) *Response {
	var rtn = new(Response)
	var addURL = c.Host + "/rs/content/add"
	content.Text = b64.StdEncoding.EncodeToString([]byte(content.Text))
	//fmt.Println(content.Text)
	aJSON := cm.GetJSONEncode(content)
	reqa, fail := cm.GetRequest(addURL, http.MethodPost, aJSON)

	if !fail {
		reqa.Header.Set("Content-Type", "application/json")
		reqa.Header.Set("Authorization", "Bearer "+c.Token)
		reqa.Header.Set("u-client-id", c.APIClient)
		reqa.Header.Set("u-api-key", c.APIKey)
		reqa.Header.Set("clientId", c.ClientID)
		reqa.Header.Set("userId", c.UserID)
		reqa.Header.Set("hashed", c.Hashed)

		code := cm.ProcessServiceCall(reqa, &rtn)
		rtn.Code = code
	}
	return rtn
}

//UpdateContent update content
func (c *ContentService) UpdateContent(content *Content) *Response {
	//fmt.Print("Content Service at start: ")
	var rtn = new(Response)
	var upURL = c.Host + "/rs/content/update"
	content.Text = b64.StdEncoding.EncodeToString([]byte(content.Text))
	//fmt.Println(content.Text)
	aJSON := cm.GetJSONEncode(content)
	requ, fail := cm.GetRequest(upURL, http.MethodPut, aJSON)
	if !fail {
		requ.Header.Set("Content-Type", "application/json")
		requ.Header.Set("Authorization", "Bearer "+c.Token)
		requ.Header.Set("u-client-id", c.APIClient)
		requ.Header.Set("u-api-key", c.APIKey)
		requ.Header.Set("clientId", c.ClientID)
		requ.Header.Set("userId", c.UserID)
		requ.Header.Set("hashed", c.Hashed)
		code := cm.ProcessServiceCall(requ, &rtn)
		rtn.Code = code
	}
	//fmt.Print("Content Service Update leaving: ")
	return rtn
}

//UpdateContentHits update content hits
func (c *ContentService) UpdateContentHits(content *Content) *Response {
	var rtn = new(Response)
	var upURL = c.Host + "/rs/content/hits"
	content.Text = b64.StdEncoding.EncodeToString([]byte(content.Text))
	//fmt.Println(content.Text)
	aJSON := cm.GetJSONEncode(content)
	reqh, fail := cm.GetRequest(upURL, http.MethodPut, aJSON)
	if !fail {
		reqh.Header.Set("Content-Type", "application/json")
		reqh.Header.Set("Authorization", "Bearer "+c.Token)
		reqh.Header.Set("u-client-id", c.APIClient)
		reqh.Header.Set("u-api-key", c.APIKey)
		reqh.Header.Set("clientId", c.ClientID)
		reqh.Header.Set("userId", c.UserID)
		reqh.Header.Set("hashed", c.Hashed)
		code := cm.ProcessServiceCall(reqh, &rtn)
		rtn.Code = code
	}
	return rtn
}

// GetContent get content
func (c *ContentService) GetContent(id string, clientID string) *Content {
	var rtn = new(Content)
	var gURL = c.Host + "/rs/content/get/" + id + "/" + clientID
	req, fail := cm.GetRequest(gURL, http.MethodGet, nil)
	if !fail {
		req.Header.Set("u-client-id", c.APIClient)
		req.Header.Set("u-api-key", c.APIKey)
		req.Header.Set("clientId", c.ClientID)
		cm.ProcessServiceCall(req, &rtn)
		rtn.Text = decodeB64String(rtn.Text)
	}
	return rtn
}

// GetContentList get content list by client
func (c *ContentService) GetContentList(clientID string) *[]Content {
	var rtn = make([]Content, 0)
	var gURL = c.Host + "/rs/content/list/" + clientID
	//fmt.Println(gURL)
	req, fail := cm.GetRequest(gURL, http.MethodGet, nil)
	if !fail {
		req.Header.Set("u-client-id", c.APIClient)
		req.Header.Set("u-api-key", c.APIKey)
		req.Header.Set("clientId", c.ClientID)
		cm.ProcessServiceCall(req, &rtn)
		for r := range rtn {
			rtn[r].Text = decodeB64String(rtn[r].Text)
			//fmt.Println(rtn[r].ModifiedDate.Year())
			if rtn[r].ModifiedDate.Year() != 1 {
				rtn[r].UseModifiedDate = true
			}
		}
	}
	return &rtn
}

// GetContentListCategory get content list by client
func (c *ContentService) GetContentListCategory(clientID string, category string) (*PageHead, *[]Content) {
	var rtn = make([]Content, 0)
	var pghead = new(PageHead)
	var gURL = c.Host + "/rs/content/list/" + clientID + "/" + category
	req, fail := cm.GetRequest(gURL, http.MethodGet, nil)
	if !fail {
		req.Header.Set("u-client-id", c.APIClient)
		req.Header.Set("u-api-key", c.APIKey)
		req.Header.Set("clientId", c.ClientID)
		cm.ProcessServiceCall(req, &rtn)

		for r := range rtn {
			rtn[r].Text = decodeB64String(rtn[r].Text)
			//txt, err := b64.StdEncoding.DecodeString(rtn[r].Text)
			rtn[r].TextHTML = template.HTML(rtn[r].Text)
			if r == 0 {
				pghead.MetaAuthor = rtn[r].MetaAuthorName
				pghead.MetaDesc = rtn[r].MetaDesc
				pghead.MetaKeyWords = rtn[r].MetaKeyWords
				pghead.Title = rtn[r].Title
			}
			if rtn[r].ModifiedDate.Year() != 1 {
				rtn[r].UseModifiedDate = true
			}
		}
	}
	return pghead, &rtn
}

// DeleteContent delete content
func (c *ContentService) DeleteContent(id string) *Response {
	var rtn = new(Response)
	var gURL = c.Host + "/rs/content/delete/" + id

	reqd, fail := cm.GetRequest(gURL, http.MethodDelete, nil)
	if !fail {
		reqd.Header.Set("Content-Type", "application/json")
		reqd.Header.Set("Authorization", "Bearer "+c.Token)
		reqd.Header.Set("u-client-id", c.APIClient)
		reqd.Header.Set("u-api-key", c.APIKey)
		reqd.Header.Set("clientId", c.ClientID)
		reqd.Header.Set("userId", c.UserID)
		reqd.Header.Set("hashed", c.Hashed)
		code := cm.ProcessServiceCall(reqd, &rtn)
		rtn.Code = code
	}
	return rtn
}

func decodeB64String(cont string) string {
	var rtn string
	txt, err := b64.StdEncoding.DecodeString(cont)
	if err != nil {
		fmt.Println(err)
	} else {
		rtn = string(txt)
	}
	return rtn
}
