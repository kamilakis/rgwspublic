package rgwspublic

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrInvalidVAT         = errors.New("invalid VAT format given")
	ErrInvalidCredentials = errors.New("username or password cannot be less than 6 chars")
)

// Version gets web service version
// returns a string or an error
func Version() (*string, error) {

	body := `<?xml version="1.0" encoding="UTF-8"?>
		<soap:Envelope 
			xmlns:soap="http://www.w3.org/2003/05/soap-envelope" 
			xmlns:rgw="http://rgwspublic2/RgWsPublic2Service">
			<soap:Header/>
			<soap:Body>
	   			<rgw:rgWsPublic2VersionInfo/>
			</soap:Body>
 		</soap:Envelope>`

	req, err := http.NewRequest("POST", Endpoint, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	header := http.Header{}
	header.Set("Content-Type", "application/soap+xml")
	header.Set("Connection", "keep-alive")
	header.Set("Content-Length", strconv.Itoa(len(body)))
	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Status: %d, error: %s", resp.StatusCode, resp.Status)
	}

	// parse response
	xmlBody, err := parseXML(resp)
	if err != nil {
		return nil, err
	}

	err = xmlBody.error()
	if err != nil {
		return nil, err
	}

	xmlBody.Error = nil // to correct parser creating an object
	return xmlBody.Version, nil

}

// AFMInfo gets info associated with a VAT number
// accepts a called by VAT and a called for VAT, username and password
// returns AFMData or an error
func AFMInfo(calledby, calledfor, user, pass string) (*VATInfo, error) {

	// vat number must be between 9 and 12 chars
	if len(calledfor) < 9 || len(calledfor) > 12 {
		return nil, ErrInvalidVAT
	}

	// calledby can be an empty string
	if calledby != "" {
		// duplicate code from above
		// vat number must be between 9 and 12 chars
		if len(calledby) < 9 || len(calledby) > 12 {
			return nil, ErrInvalidVAT
		}
	}

	// same for username/password
	if len(user) < 6 || len(pass) < 6 {
		return nil, ErrInvalidCredentials
	}

	body := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
		<env:Envelope 
			xmlns:env="http://www.w3.org/2003/05/soap-envelope" 
			xmlns:ns1="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" 
			xmlns:ns2="http://rgwspublic2/RgWsPublic2Service" 
			xmlns:ns3="http://rgwspublic2/RgWsPublic2">
			<env:Header>
	   			<ns1:Security>
		  			<ns1:UsernameToken>
			 			<ns1:Username>%v</ns1:Username>
			 			<ns1:Password>%v</ns1:Password>
		  			</ns1:UsernameToken>
	   			</ns1:Security>
			</env:Header>
			<env:Body>
	   			<ns2:rgWsPublic2AfmMethod>
		  			<ns2:INPUT_REC>
			 			<ns3:afm_called_by>%v</ns3:afm_called_by>
			 			<ns3:afm_called_for>%v</ns3:afm_called_for>
		  			</ns2:INPUT_REC>
	   			</ns2:rgWsPublic2AfmMethod>
			</env:Body>
 		</env:Envelope>`, user, pass, calledby, calledfor)

	req, err := http.NewRequest("POST", Endpoint, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	header := http.Header{}
	header.Set("Content-Type", "application/soap+xml")
	header.Set("Connection", "keep-alive")
	header.Set("Content-Length", strconv.Itoa(len(body)))
	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Status: %d, error: %s", resp.StatusCode, resp.Status)
	}

	xmlBody, err := parseXML(resp)
	if err != nil {
		return nil, err
	}

	err = xmlBody.VATInfo.error()
	if err != nil {
		return nil, err
	}

	xmlBody.VATInfo.Error = nil // to correct parser creating an object
	return &xmlBody.VATInfo, nil
}

// helper function to parse xml response
func parseXML(r *http.Response) (*XMLBody, error) {

	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	xmlResp := XMLResponse{}
	err = xml.Unmarshal([]byte(rbody), &xmlResp)
	if err != nil {
		return nil, err
	}

	return &xmlResp.Body, nil
}
