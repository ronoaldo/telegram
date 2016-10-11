package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	TelegramBotEndpoint          = "https://api.telegram.org/bot"
	TelegramFileDownloadEndpoint = "https://api.telegram.org/file/bot"
)

type DebugFunc func(msg string)

type ApiClient struct {
	client *http.Client
	token  string
	debug  DebugFunc
}

func stderrDebug(msg string) {
	log.Println(msg)
}

// NewApiClient returns an instance of a Telegram Bot API client.
func NewApiClient(c *http.Client, token string) *ApiClient {
	return &ApiClient{
		client: c,
		token:  token,
		debug:  stderrDebug,
	}
}

func (t *ApiClient) Debug(dbg DebugFunc) {
	t.debug = dbg
}

func (t *ApiClient) Debugf(msg string, args ...interface{}) {
	if t.debug != nil {
		t.debug(fmt.Sprintf(msg, args...))
	}
}

func (t *ApiClient) Call(httpMethod, apiMethod string, in, out interface{}) error {
	var buff bytes.Buffer

	if err := json.NewEncoder(&buff).Encode(in); err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s/%s", TelegramBotEndpoint, t.token, apiMethod)

	req, err := http.NewRequest(httpMethod, url, &buff)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	t.Debugf("* API Response: %v", string(b))

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("telegram: unexpected response: %v: %v [%s]", resp.StatusCode, resp.Status, string(b))
	}

	// Check if operation suceeded
	apiResp := new(ApiResponse)
	if err := json.Unmarshal(b, apiResp); err != nil {
		return fmt.Errorf("telegram: unable to parse response %v", err)
	}
	if !apiResp.OK {
		return fmt.Errorf("telegram: operation failed: %s", string(b))
	}

	return json.Unmarshal(apiResp.Result, out)
}

// GetFile prepares a file for download.
func (t *ApiClient) GetFile(fileId string) (*File, error) {
	params := map[string]string{
		"file_id": fileId,
	}
	t.Debugf("Fetching file with getFile: %v", fileId)
	file := new(File)
	if err := t.Call("GET", "getFile", params, file); err != nil {
		return nil, err
	}
	return file, nil
}

// DownloadFile fetches the file from f.FilePath and writes the content into w.
func (t *ApiClient) DownloadFile(f *File, w io.Writer) error {
	// https://api.telegram.org/file/bot<token>/<file_path>
	url := fmt.Sprintf("%s%s/%s", TelegramFileDownloadEndpoint, t.token, f.FilePath)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, resp.Body)
	return err
}

// SendMessage sends a plain text message to the provided recipient.
func (t *ApiClient) SendMessage(to, text string) (*Message, error) {
	params := map[string]string{
		"chat_id": to,
		"text":    text,
	}
	msg := new(Message)
	if err := t.Call("POST", "sendMessage", params, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

// SendFormattedMessage sends a formatted message in either HTML or Markdown.
func (t *ApiClient) SendFormattedMessage(to, text string, parseMode ParseMode) (*Message, error) {
	params := map[string]string{
		"chat_id":    to,
		"text":       text,
		"parse_mode": string(parseMode),
	}
	msg := new(Message)
	if err := t.Call("POST", "sendMessage", params, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

// SendMessagef calls fmt.Sprintf and passes the resulting message to SendMessage.
func (t *ApiClient) SendMessagef(to, formatText string, args ...interface{}) (*Message, error) {
	return t.SendMessage(to, fmt.Sprintf(formatText, args...))
}

// SetWebhook method configures the provided HTTPS endpoint as the bot callback.
func (t *ApiClient) SetWebhook(httpsURL string) error {
	out := make(json.RawMessage, 0)
	if err := t.Call("POST", fmt.Sprintf("setWebhook?url=%s", httpsURL), nil, out); err != nil {
		return err
	}
	return nil
}

// ApiResponse is the response API wrapper.
type ApiResponse struct {
	OK      bool            `json:"ok"`
	Result  json.RawMessage `json:"result"`
	Message string          `json:"message"`
}
