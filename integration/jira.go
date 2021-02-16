package integration

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)


const (
	jiraReport = "https://seu_dominio_aqui/rest/api/2/search"
)


// Jira - struct
type Jira struct {
	Hours int
}

// ResponseJiraIssue - struct
type ResponseJiraIssue struct {
	Expand string `json:"expand"`
	Issues []struct {
		Expand string `json:"expand"`
		Fields struct {
			Worklog struct {
				MaxResults int64 `json:"maxResults"`
				StartAt    int64 `json:"startAt"`
				Total      int64 `json:"total"`
				Worklogs   []struct {
					Author struct {
						AccountID   string `json:"accountId"`
						AccountType string `json:"accountType"`
						Active      bool   `json:"active"`
						AvatarUrls  struct {
							One6x16   string `json:"16x16"`
							Two4x24   string `json:"24x24"`
							Three2x32 string `json:"32x32"`
							Four8x48  string `json:"48x48"`
						} `json:"avatarUrls"`
						DisplayName string `json:"displayName"`
						Self        string `json:"self"`
						TimeZone    string `json:"timeZone"`
					} `json:"author"`
					Comment struct {
						Content []struct {
							Content []struct {
								Text string `json:"text"`
								Type string `json:"type"`
							} `json:"content"`
							Type string `json:"type"`
						} `json:"content"`
						Type    string `json:"type"`
						Version int64  `json:"version"`
					} `json:"comment"`
					Created          string `json:"created"`
					ID               string `json:"id"`
					IssueID          string `json:"issueId"`
					Self             string `json:"self"`
					Started          string `json:"started"`
					TimeSpent        string `json:"timeSpent"`
					TimeSpentSeconds int64  `json:"timeSpentSeconds"`
					UpdateAuthor     struct {
						AccountID   string `json:"accountId"`
						AccountType string `json:"accountType"`
						Active      bool   `json:"active"`
						AvatarUrls  struct {
							One6x16   string `json:"16x16"`
							Two4x24   string `json:"24x24"`
							Three2x32 string `json:"32x32"`
							Four8x48  string `json:"48x48"`
						} `json:"avatarUrls"`
						DisplayName string `json:"displayName"`
						Self        string `json:"self"`
						TimeZone    string `json:"timeZone"`
					} `json:"updateAuthor"`
					Updated string `json:"updated"`
				} `json:"worklogs"`
			} `json:"worklog"`
		} `json:"fields"`
		ID   string `json:"id"`
		Key  string `json:"key"`
		Self string `json:"self"`
	} `json:"issues"`
	MaxResults int64 `json:"maxResults"`
	StartAt    int64 `json:"startAt"`
	Total      int64 `json:"total"`
}


func mountedHttp (url, authorization, method string, body *strings.Reader) (*http.Response, error) {
	timeout := 5 * time.Second

	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest(method, url, body)


	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", authorization)

	response, e := client.Do(request)

	if e != nil {
		return nil, e
	}

	return response, nil
}

// RequestHttpJiraReport - chamada http para request de issues do jira.
func RequestHttpJiraReport (authorization string, body *strings.Reader) ([]byte, error) {

	response, err := mountedHttp(jiraReport, authorization, "POST", body)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()


	data, er := ioutil.ReadAll(response.Body)

	if er != nil {
		return nil, er
	}

	return data, nil
}
