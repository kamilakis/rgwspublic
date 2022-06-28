package rgwspublic

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// XMLResponse is where we parse an http response
type XMLResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    XMLBody
}

// XMLBody is the body of a response
type XMLBody struct {
	XMLName                   xml.Name `xml:"Body"`
	AFMMethodResponse         AFMMethodResponseData
	PublicVersionInfoResponse PublicVersionInfoResponseData
}

// XMLRGWSPublicAfmMethodResponse is the response to a publicAFM method
type AFMMethodResponseData struct {
	XMLName xml.Name `xml:"rgWsPublic2AfmMethodResponse" json:"-"`
	Result  ResultData
}

// PublicVersionInfoResponseData holds version info
type PublicVersionInfoResponseData struct {
	XMLName xml.Name `xml:"rgWsPublic2VersionInfoResponse"`
	Result  string   `xml:"result"`
}

type ResultData struct {
	XMLName    xml.Name `xml:"result" json:"-"`
	ResultType ResultTypeData
}

type ResultTypeData struct {
	XMLName        xml.Name `xml:"rg_ws_public2_result_rtType" json:"-"`
	CallSeqID      string   `xml:"call_seq_id" json:"call_seq_id"`
	AFMCalledByRec AFMCalledByRecData
	BasicRec       BasicRecData
	Activities     FirmActivityTabData `json:"-"`
	Error          ErrorRecData
}

type CallSeqIDData struct {
	CallSeqID string `xml:"call_seq_id" json:"call_seq_id"`
}

type ErrorRecData struct {
	XMLName          xml.Name `xml:"error_rec" json:"-"`
	ErrorCode        string   `xml:"error_code" json:"error_code"`
	ErrorDescription string   `xml:"error_descr" json:"error_description"`
}

// AFMDCalledByData is the data relative to who did the search
type AFMCalledByRecData struct {
	XMLName             xml.Name `xml:"afm_called_by_rec" json:"-"`
	TokenUsername       string   `xml:"token_username" json:"token_username"`
	TokenAFM            string   `xml:"token_afm" json:"token_afm"`
	TokenAFMFullName    string   `xml:"token_afm_fullname" json:"token_afm_fullname"`
	AFMCalledBy         string   `xml:"afm_called_by" json:"afm_called_by"`
	AFMCalledByFullName string   `xml:"afm_called_by_fullname" json:"afm_called_by_fullname"`
	AsOnDate            string   `xml:"as_on_date" json:"as_on_date"`
}

// BasicRecData is the data relative to an entity's VAT search
type BasicRecData struct {
	XMLName                     xml.Name `xml:"basic_rec" json:"-"`
	AFM                         string   `xml:"afm" json:"afm"`                                              // ΑΦΜ
	DOY                         string   `xml:"doy" json:"doy"`                                              // ΚΩΔΙΚΟΣ ΔΟΥ
	DOYDescription              string   `xml:"doy_descr" json:"doy_description"`                            // ΠΕΡΙΓΡΑΦΗ ΔΟΥ
	InitialFlagDescription      string   `xml:"i_ni_flag_descr" json:"initial_flag_description"`             // ΦΠ /ΜΗ ΦΠ
	DeactivationFlag            string   `xml:"deactivation_flag" json:"deactivation_flag"`                  // ΕΝΔΕΙΞΗ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ:1=ΕΝΕΡΓΟΣ ΑΦΜ 2=ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ
	DeactivationFlagDescription string   `xml:"deactivation_flag_desc" json:"deactivation_flag_description"` // ΕΝΔΕΙΞΗ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ(ΠΕΡΙΓΡΑΦΗ): ΕΝΕΡΓΟΣ ΑΦΜ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ
	FirmFlagDescription         string   `xml:"firm_flag_descr" json:"firm_flag_description"`                // ΤΙΜΕΣ: ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ, ΜΗ ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ, ΠΡΩΗΝ ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ
	Onomasia                    string   `xml:"onomasia" json:"onomasia"`                                    // ΕΠΩΝΥΜΙΑ
	CommercialTitle             string   `xml:"commer_title" json:"commercial_title"`                        // ΤΙΤΛΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	LegalStatusDescription      string   `xml:"legal_status_descr" json:"legal_status_descr"`                // ΠΕΡΙΓΡΑΦΗ ΜΟΡΦΗΣ ΜΗ Φ.Π.
	PostalAddress               string   `xml:"postal_address" json:"postal_address"`                        // ΟΔΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	PostalAddressNo             string   `xml:"postal_address_no" json:"postal_address_no"`                  // ΑΡΙΘΜΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	PostalZipCode               string   `xml:"postal_zip_code" json:"postal_zip_code"`                      // ΤΑΧ. ΚΩΔ. ΕΠΙΧΕΙΡΗΣΗΣ
	PostalAreaDescription       string   `xml:"postal_area_description" json:"postal_area_description"`      // ΠΕΡΙΟΧΗ ΕΠΙΧΕΙΡΗΣΗΣ
	RegistrationDate            string   `xml:"regist_date" json:"registration_date"`                        // ΗΜ/ΝΙΑ ΕΝΑΡΞΗΣ
	StopDate                    string   `xml:"stop_date" json:"stop_date"`                                  // ΗΜ/ΝΙΑ ΔΙΑΚΟΠΗΣ
	NormalVATSystemFlag         string   `xml:"normal_vat_system_flag" json:"normal_vat_system_flag"`
	Activities                  []FirmActivities
}

// FirmActivities is the activities of the entity
type FirmActivityTabData struct {
	XMLName    xml.Name         `xml:"firm_act_tab" json:"-"`
	Activities []FirmActivities `xml:"item" json:"-"`
}

type FirmActivities struct {
	// XMLName             xml.Name `xml:"item" json:"-"`
	FirmActCode         string `xml:"firm_act_code" json:"firm_act_code"`                   // ΚΩΔΙΚΟΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ
	FirmActDescriptionn string `xml:"firm_act_descr" json:"firm_act_description"`           // ΠΕΡΙΓΡΑΦΗ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ
	FirmActKind         string `xml:"firm_act_kind" json:"firm_activity_kind"`              // ΕΙΔΟΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ: 1=ΚΥΡΙΑ, 2=ΔΕΥΤΕΡΕΥΟΥΣΑ, 3=ΛΟΙΠΗ, 4=ΒΟΗΘΗΤΙΚΗ
	FirmActKindDescr    string `xml:"firm_act_kind_descr" json:"firm_act_kind_description"` // ΠΕΡΙΓΡΑΦΗ ΕΙΔΟΥΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ: ΚΥΡΙΑ, ΔΕΥΤΕΡΕΥΟΥΣΑ, ΛΟΙΠΗ, ΒΟΗΘΗΤΙΚΗ
}

const (
	// Endpoint is the url for WSDL service
	Endpoint                                       = "https://www1.gsis.gr/wsaade/RgWsPublic2/RgWsPublic2"
	RG_WS_PUBLIC_AFM_CALLED_BY_BLOCKED             = "Ο χρήστης που καλεί την υπηρεσία έχει προσωρινά αποκλειστεί από τη χρήση της."
	RG_WS_PUBLIC_AFM_CALLED_BY_NOT_FOUND           = "Ο Α.Φ.Μ. για τον οποίο γίνεται η κλήση δε βρέθηκε στους έγκυρους Α.Φ.Μ. του Μητρώου TAXIS."
	RG_WS_PUBLIC_EPIT_NF                           = "O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δεν ανήκει και δεν ανήκε ποτέ σε νομικό πρόσωπο, νομική οντότητα, ή φυσικό πρόσωπο με εισόδημα από επιχειρηματική δραστηριότητα."
	RG_WS_PUBLIC_FAILURES_TOLERATED_EXCEEDED       = "Υπέρβαση μέγιστου επιτρεπτού ορίου πρόσφατων αποτυχημένων κλήσεων. Προσπαθήστε εκ νέου σε μερικές ώρες."
	RG_WS_PUBLIC_MAX_DAILY_USERNAME_CALLS_EXCEEDED = "Υπέρβαση μέγιστου επιτρεπτού ορίου ημερήσιων κλήσεων ανά χρήστη (ανεξαρτήτως εξουσιοδοτήσεων)."
	RG_WS_PUBLIC_MONTHLY_LIMIT_EXCEEDED            = "Υπέρβαση του Μέγιστου Επιτρεπτού Μηνιαίου Ορίου Κλήσεων."
	RG_WS_PUBLIC_MSG_TO_TAXISNET_ERROR             = "Δημιουργήθηκε πρόβλημα κατά την ενημέρωση των εισερχόμενων μηνυμάτων στο MyTAXISnet."
	RG_WS_PUBLIC_NO_INPUT_PARAMETERS               = "Δε δόθηκαν υποχρεωτικές παράμετροι εισόδου για την κλήση της υπηρεσίας."
	RG_WS_PUBLIC_SERVICE_NOT_ACTIVE                = "Η υπηρεσία δεν είναι ενεργή."
	RG_WS_PUBLIC_TAXPAYER_NF                       = "O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δε βρέθηκε στους έγκυρους Α.Φ.Μ. του Μητρώου TAXIS."
	RG_WS_PUBLIC_TOKEN_AFM_BLOCKED                 = "Ο χρήστης (ή ο εξουσιοδοτημένος τρίτος) που καλεί την υπηρεσία έχει προσωρινά αποκλειστεί από τη χρήση της."
	RG_WS_PUBLIC_TOKEN_AFM_NOT_AUTHORIZED          = "Ο τρέχον χρήστης δεν έχει εξουσιοδοτηθεί από τον Α.Φ.Μ. για χρήση της υπηρεσίας."
	RG_WS_PUBLIC_TOKEN_AFM_NOT_FOUND               = "Ο Α.Φ.Μ. του τρέχοντος χρήστη δε βρέθηκε στους έγκυρους Α.Φ.Μ. του Μητρώου TAXIS."
	RG_WS_PUBLIC_TOKEN_AFM_NOT_REGISTERED          = "Ο τρέχον χρήστης δεν έχει εγγραφεί για χρήση της υπηρεσίας."
	RG_WS_PUBLIC_TOKEN_USERNAME_NOT_ACTIVE         = "Ο κωδικός χρήστη (username) που χρησιμοποιήθηκε έχει ανακληθεί."
	RG_WS_PUBLIC_TOKEN_USERNAME_NOT_AUTHENTICATED  = "Ο συνδυασμός χρήστη/κωδικού πρόσβασης που δόθηκε δεν είναι έγκυρος."
	RG_WS_PUBLIC_TOKEN_USERNAME_NOT_DEFINED        = "Δεν ορίσθηκε ο χρήστης που καλεί την υπηρεσία."
	RG_WS_PUBLIC_TOKEN_USERNAME_TOO_LONG           = "Διαπιστώθηκε υπέρβαση του μήκους του ονόματος του χρήστη (username) της υπηρεσίας"
	RG_WS_PUBLIC_WRONG_AFM                         = "O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δεν είναι έγκυρος."
)

func (a *ResultTypeData) JSON() (string, error) {
	var j []byte
	j, err := json.MarshalIndent(&a, "", "\t")
	if err != nil {
		return "", err
	}

	return string(j), nil
}

func (a *ResultTypeData) String() string {
	var s string

	s = fmt.Sprintf("XMLName:%s\n", a.XMLName)
	s += fmt.Sprintf("afm:%s\n", a.BasicRec.AFM)
	s += fmt.Sprintf("doy:%s\n", a.BasicRec.DOY)
	s += fmt.Sprintf("doy_descr:%s\n", a.BasicRec.DOYDescription)
	s += fmt.Sprintf("i_ni_flag_descr:%s\n", a.BasicRec.InitialFlagDescription)
	s += fmt.Sprintf("deactivation_flag:%s\n", a.BasicRec.DeactivationFlag)
	s += fmt.Sprintf("deactivation_flag_desc:%s\n", a.BasicRec.DeactivationFlagDescription)
	s += fmt.Sprintf("firm_flag_descr:%s\n", a.BasicRec.FirmFlagDescription)
	s += fmt.Sprintf("onomasia:%s\n", a.BasicRec.Onomasia)
	s += fmt.Sprintf("commer_title:%s\n", a.BasicRec.CommercialTitle)
	s += fmt.Sprintf("legal_status_descr:%s\n", a.BasicRec.LegalStatusDescription)
	s += fmt.Sprintf("postal_address:%s\n", a.BasicRec.PostalAddress)
	s += fmt.Sprintf("postal_address_no:%s\n", a.BasicRec.PostalAddressNo)
	s += fmt.Sprintf("postal_zip_code:%s\n", a.BasicRec.PostalZipCode)
	s += fmt.Sprintf("postal_area_description:%s\n", a.BasicRec.PostalAreaDescription)
	s += fmt.Sprintf("regist_date:%s\n", a.BasicRec.RegistrationDate)
	s += fmt.Sprintf("stop_date:%s\n", a.BasicRec.StopDate)
	s += fmt.Sprintf("normal_vat_system_flag:%s\n", a.BasicRec.NormalVATSystemFlag)

	s += fmt.Sprintf("ACTIVITIES:--------------------\n")
	for k, v := range a.BasicRec.Activities {
		s += fmt.Sprintf("ACTIVITY #%d\n", k)
		s += fmt.Sprintf("FirmActCode: %s\n", v.FirmActCode)
		s += fmt.Sprintf("FirmActDescr: %s\n", v.FirmActDescriptionn)
		s += fmt.Sprintf("FirmActKind: %s\n", v.FirmActKind)
		s += fmt.Sprintf("FirmActKindDescr: %s\n", v.FirmActKindDescr)
	}

	s += fmt.Sprintf("ErrorDescr: %s\n", a.Error.ErrorDescription)
	s += fmt.Sprintf("ErrorCode: %s\n", a.Error.ErrorCode)

	return s
}
