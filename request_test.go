package rgwspublic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestJSON(t *testing.T) {

	// a mock http response
	body := fmt.Sprintf(`<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<env:Header/>
	<env:Body>
		<m:rgWsPublicAfmMethodResponse xmlns:m="http://gr/gsis/rgwspublic/RgWsPublic.wsdl">
			<RgWsPublicBasicRt_out>
				<m:afm>094014298   </m:afm>
				<m:stopDate xsi:nil="true"/>
				<m:postalAddressNo>4        </m:postalAddressNo>
				<m:doyDescr>Φ.Α.Ε. ΑΘΗΝΩΝ</m:doyDescr>
				<m:doy>1159</m:doy>
				<m:onomasia>ΤΡΑΠΕΖΑ ΠΕΙΡΑΙΩΣ Α Ε</m:onomasia>
				<m:legalStatusDescr>ΑΕ</m:legalStatusDescr>
				<m:registDate>1916-01-01T00:00:00.000+01:34</m:registDate>
				<m:deactivationFlag>1</m:deactivationFlag>
				<m:deactivationFlagDescr>ΕΝΕΡΓΟΣ ΑΦΜ          </m:deactivationFlagDescr>
				<m:postalAddress>ΑΜΕΡΙΚΗΣ</m:postalAddress>
				<m:firmFlagDescr>ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ      </m:firmFlagDescr>
				<m:commerTitle xsi:nil="true"/>
				<m:postalAreaDescription>ΑΘΗΝΑ</m:postalAreaDescription>
				<m:INiFlagDescr>ΜΗ ΦΠ</m:INiFlagDescr>
				<m:postalZipCode>10564</m:postalZipCode>
			</RgWsPublicBasicRt_out>
			<arrayOfRgWsPublicFirmActRt_out>
				<m:RgWsPublicFirmActRtUser>
					<m:firmActDescr>ΥΠΗΡΕΣΙΕΣ ΤΡΑΠΕΖΩΝ</m:firmActDescr>
					<m:firmActKind>1</m:firmActKind>
					<m:firmActKindDescr>ΚΥΡΙΑ</m:firmActKindDescr>
					<m:firmActCode>64191204</m:firmActCode>
				</m:RgWsPublicFirmActRtUser>
			</arrayOfRgWsPublicFirmActRt_out>
			<pCallSeqId_out>709330921</pCallSeqId_out>
			<pErrorRec_out>
				<m:errorDescr xsi:nil="true"/>
				<m:errorCode xsi:nil="true"/>
			</pErrorRec_out>
		</m:rgWsPublicAfmMethodResponse>
	</env:Body>
</env:Envelope>`)

	r := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header, 0),
	}

	i, err := parseXML(r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(json.Marshal(i))
}

func TestVersion(t *testing.T) {

	version, err := Version()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("got version: %s\n", *version)
}

func TestInvalids(t *testing.T) {

	// some invalid input to test returned service errors
	inputs := []map[string]string{
		{
			"vat":      "1234567890",
			"username": "someuser",
			"password": "somepass",
			"error":    "RG_WS_PUBLIC_TOKEN_USERNAME_NOT_AUTHENTICATED",
		},
		{
			"vat":      "104807035",
			"username": "someuser",
			"password": "somepass",
			"error":    "RG_WS_PUBLIC_TOKEN_USERNAME_NOT_AUTHENTICATED",
		},
		{
			"vat":      "1234567890",
			"username": "KAMILAKIS1", // valid user but,
			"password": "hyhyhhyh!",  // wrong pass
			"error":    "RG_WS_PUBLIC_TOKEN_USERNAME_NOT_AUTHENTICATED",
		},
		// {
		// 	"vat":      "1234567890",
		// 	"username": "KAMILAKIS1", // put valid credentials here
		// 	"password": "hyhyhhyh!",  // for this test to pass
		// 	"error":    "RG_WS_PUBLIC_WRONG_AFM",
		// },
	}

	for k, v := range inputs {
		t.Logf("testing input #%d, vat:%s, user:%s, pass:%s", k, v["vat"], v["username"], v["password"])
		i, err := AFMInfo("", v["vat"], v["username"], v["password"])
		if err != nil {
			t.Errorf("error getting AFM info: %s", err.Error())
			continue
		}

		if i.Error.Code != v["error"] {
			t.Errorf("error code returned not expected, got: %s, wanted: %s", i.Error.Code, v["error"])
		}
	}

}

func TestPublicInfo(t *testing.T) {

	// VAT number of InfoQuest
	// replace username and password with the ones you got from
	// http://www.gsis.gr/gsis/info/gsis_site/PublicIssue/wnsp/wnsp_pages/wnsp.html
	i, err := AFMInfo("", "998184801", "username", "password")
	if err != nil {
		t.Errorf("error getting AFM info: %s", err.Error())
	}

	fmt.Println(i.String())
}

func TestParseAFMInfo(t *testing.T) {

	// a mock http response
	body := fmt.Sprintf(`<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<env:Header/>
	<env:Body>
		<m:rgWsPublicAfmMethodResponse xmlns:m="http://gr/gsis/rgwspublic/RgWsPublic.wsdl">
			<RgWsPublicBasicRt_out>
				<m:afm>094014298   </m:afm>
				<m:stopDate xsi:nil="true"/>
				<m:postalAddressNo>4        </m:postalAddressNo>
				<m:doyDescr>Φ.Α.Ε. ΑΘΗΝΩΝ</m:doyDescr>
				<m:doy>1159</m:doy>
				<m:onomasia>ΤΡΑΠΕΖΑ ΠΕΙΡΑΙΩΣ Α Ε</m:onomasia>
				<m:legalStatusDescr>ΑΕ</m:legalStatusDescr>
				<m:registDate>1916-01-01T00:00:00.000+01:34</m:registDate>
				<m:deactivationFlag>1</m:deactivationFlag>
				<m:deactivationFlagDescr>ΕΝΕΡΓΟΣ ΑΦΜ          </m:deactivationFlagDescr>
				<m:postalAddress>ΑΜΕΡΙΚΗΣ</m:postalAddress>
				<m:firmFlagDescr>ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ      </m:firmFlagDescr>
				<m:commerTitle xsi:nil="true"/>
				<m:postalAreaDescription>ΑΘΗΝΑ</m:postalAreaDescription>
				<m:INiFlagDescr>ΜΗ ΦΠ</m:INiFlagDescr>
				<m:postalZipCode>10564</m:postalZipCode>
			</RgWsPublicBasicRt_out>
			<arrayOfRgWsPublicFirmActRt_out>
				<m:RgWsPublicFirmActRtUser>
					<m:firmActDescr>ΥΠΗΡΕΣΙΕΣ ΤΡΑΠΕΖΩΝ</m:firmActDescr>
					<m:firmActKind>1</m:firmActKind>
					<m:firmActKindDescr>ΚΥΡΙΑ</m:firmActKindDescr>
					<m:firmActCode>64191204</m:firmActCode>
				</m:RgWsPublicFirmActRtUser>
			</arrayOfRgWsPublicFirmActRt_out>
			<pCallSeqId_out>709330921</pCallSeqId_out>
			<pErrorRec_out>
				<m:errorDescr xsi:nil="true"/>
				<m:errorCode xsi:nil="true"/>
			</pErrorRec_out>
		</m:rgWsPublicAfmMethodResponse>
	</env:Body>
</env:Envelope>`)

	r := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header, 0),
	}

	i, err := parseXML(r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(i.VATInfo.String())
}

func TestParseVersion(t *testing.T) {

	// a mock http response
	body := fmt.Sprintf(`<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
		<env:Header/>
			<env:Body>
				<m:rgWsPublicVersionInfoResponse xmlns:m="http://gr/gsis/rgwspublic/RgWsPublic.wsdl">
					<result>Version: 3.1.0, 11/04/2014, Copyright Γ.Γ.Δ.Ε. / Γ.Γ.Π.Σ. 2014.Υπηρεσία &quot;Βασικά στοιχεία μητρώου για νομικά πρόσωπα, νομικές οντότητες, και φυσικά πρόσωπα με εισόδημα από επιχειρηματική δραστηριότητα» με όριο κλήσεων και ταυτοποίηση χρήστη.</result>
				</m:rgWsPublicVersionInfoResponse>
			</env:Body>
		</env:Envelope>`)

	r := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header, 0),
	}

	v, err := parseXML(r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(v)
}
