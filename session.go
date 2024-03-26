package vmmanager6

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var Debug = new(bool)

type Response struct {
	Resp *http.Response
	Body []byte
}

type Session struct {
	httpClient *http.Client
	ApiUrl     string
	AuthToken  string
	Headers    http.Header
}

func NewSession(apiUrl string, hclient *http.Client, tls *tls.Config) (session *Session, err error) {
	if hclient == nil {
		tr := &http.Transport{
			TLSClientConfig:    tls,
			DisableCompression: true,
			Proxy:              nil,
		}
		hclient = &http.Client{Transport: tr}
	}
	session = &Session{
		httpClient: hclient,
		ApiUrl:	apiUrl,
		AuthToken: "",
		Headers: http.Header{},
	}
	return session, nil
}

func decodeResponse(resp *http.Response, v interface{}) error {
	if resp.Body == nil {
		return nil
	}
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %s", err)
	}
	if err = json.Unmarshal(rbody, &v); err != nil {
		return err
	}
	return nil
}

func ResponseJSON(resp *http.Response) (jbody map[string]interface{}, err error) {
	err = decodeResponse(resp, &jbody)
	return jbody, err
}

func (s *Session) NewRequest(method, url string, headers *http.Header, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		req.Header = *headers
	}
	if s.AuthToken != "" {
		req.Header.Add("x-xsrf-token", s.AuthToken)
		req.Header.Add("Cookie", fmt.Sprintf("ses6=%s", s.AuthToken))
	}
	return
}
func (s *Session) Do(req *http.Request) (*http.Response, error) {
	// Add session headers
	for k := range s.Headers {
		req.Header.Set(k, s.Headers.Get(k))
	}

	if *Debug {
		d, _ := httputil.DumpRequestOut(req, true)
		log.Printf(">>>>>>>>>> REQUEST:\n%v", string(d))
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// The response body reader needs to be closed, but lots of places call
	// session.Do, and they might not be able to reliably close it themselves.
	// Therefore, read the body out, close the original, then replace it with
	// a NopCloser over the bytes, which does not need to be closed downsteam.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	resp.Body = ioutil.NopCloser(bytes.NewReader(respBody))

	if *Debug {
		dr, _ := httputil.DumpResponse(resp, true)
		log.Printf("<<<<<<<<<< RESULT:\n%v", string(dr))
	}
	if resp.StatusCode == 503 {
		return resp, nil
	}
	if resp.StatusCode == 400 || resp.StatusCode == 500 {
		return resp, fmt.Errorf("%s", string(respBody))
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, fmt.Errorf(resp.Status)
	}

	return resp, nil
}
// Perform a simple get to an endpoint
func (s *Session) Request(
	method string,
	url string,
	params *url.Values,
	headers *http.Header,
	body *[]byte,
) (resp *http.Response, err error) {
	// add params to url here
	url = s.ApiUrl + url
	if params != nil {
		url = url + "?" + params.Encode()
	}

	// Get the body if one is present
	var buf io.Reader
	if body != nil {
		buf = bytes.NewReader(*body)
	}

	req, err := s.NewRequest(method, url, headers, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	return s.Do(req)
}

// Perform a simple get to an endpoint and unmarshall returned JSON
func (s *Session) RequestJSON(
	method string,
	url string,
	params *url.Values,
	headers *http.Header,
	body interface{},
	responseContainer interface{},
) (resp *http.Response, err error) {
	var bodyjson []byte
	if body != nil {
		bodyjson, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	// if headers == nil {
	// 	headers = &http.Header{}
	// 	headers.Add("Content-Type", "application/json")
	// }

	resp, err = s.Request(method, url, params, headers, &bodyjson)
	if err != nil {
		return resp, err
	}

	// err = util.CheckHTTPResponseStatusCode(resp)
	// if err != nil {
	// 	return nil, err
	// }

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, fmt.Errorf("error reading response body")
	}
	if err = json.Unmarshal(rbody, &responseContainer); err != nil {
                return resp, err
        }
	return resp, nil
}
func (s *Session) Delete(
	url string,
	params *url.Values,
	headers *http.Header,
) (resp *http.Response, err error) {
	return s.Request("DELETE", url, params, headers, nil)
}
func (s *Session) Get(
	url string,
	params *url.Values,
	headers *http.Header,
) (resp *http.Response, err error) {
	return s.Request("GET", url, params, headers, nil)
}
func (s *Session) GetJSON(
	url string,
	params *url.Values,
	headers *http.Header,
	responseContainer interface{},
) (resp *http.Response, err error) {
	return s.RequestJSON("GET", url, params, headers, nil, responseContainer)
}

func (s *Session) Head(
	url string,
	params *url.Values,
	headers *http.Header,
) (resp *http.Response, err error) {
	return s.Request("HEAD", url, params, headers, nil)
}
func (s *Session) Post(
	url string,
	params *url.Values,
	headers *http.Header,
	body *[]byte,
) (resp *http.Response, err error) {
	if headers == nil {
		headers = &http.Header{}
		headers.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	return s.Request("POST", url, params, headers, body)
}

func (s *Session) DeleteJSON(
        url string,
        params *url.Values,
        headers *http.Header,
        body interface{},
        responseContainer interface{},
) (resp *http.Response, err error) {
        return s.RequestJSON("DELETE", url, params, headers, body, responseContainer )
}

func (s *Session) PostJSON(
	url string,
	params *url.Values,
	headers *http.Header,
	body interface{},
	responseContainer interface{},
) (resp *http.Response, err error) {
	return s.RequestJSON("POST", url, params, headers, body, responseContainer )
}

func (s *Session) SetAPIToken(token string) {
	s.AuthToken = token
}

func (s *Session) Login(username string, password string) (err error) {
	var resp *http.Response
	reqUser := map[string]interface{}{"email": username, "password": password}
	olddebug := *Debug
	var data map[string]interface{}
	*Debug = false // don't share passwords in debug log
	for ii := 0; ii < 5; ii++ {
		resp, err = s.PostJSON("/auth/v4/public/token", nil , nil, &reqUser, &data)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("Login error reading response")
		}
		if resp.StatusCode != 200 {
			log.Printf("[DEBUG][Login] Sleeping for %d seconds before another login try", ii+1)
			time.Sleep(time.Duration(ii+1) * time.Second)
		} else {
			*Debug = olddebug
			s.AuthToken = data["token"].(string)
			return nil
		}
		
	}
	return fmt.Errorf("Can't login after 5 tries")
}

func ParamsToBody(params map[string]interface{}) (vals url.Values) {
	vals = url.Values{}
	log.Printf("%s", params["email"])
	for k, intrV := range params {
		var v string
                v = fmt.Sprintf("%v", intrV)
		vals.Set(k, v)
		log.Printf("%s", vals["email"])
	}
	return
}


