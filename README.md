# rgwspublic
client library for greek GSIS tax service in [Go](https://golang.org/)

[![GoDoc](https://godoc.org/github.com/kamilakis/rgwspublic?status.svg)](https://godoc.org/github.com/kamhlos/rgwspublic)

## Example

```go
package main

import (
	"fmt"

	"github.com/kamilakis/rgwspublic"
)

func main() {

	// get service version
	v, err := rgwspublic.Version()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)

	// get info using VAT number of a public service
	// replace username and password with the ones you got from
	// https://www1.aade.gr/sgsisapps/tokenservices/protected/displayConsole.htm
	i, err := rgwspublic.GetVATInfo("", "090165560", "username", "password")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(i)
}

```
Note that username and password is supplied from the [service](https://www1.aade.gr/sgsisapps/tokenservices/protected/displayConsole.htm).

`go get -v`

`go run main.go`

Two methods are exposed, `GetVatInfo(string, string, string)` and `Version()`:
GetVatInfo() accepts two vat numbers (strings), and service credentials (username, password).

First VAT number is the callee, second is the one we want information for. The callee can be empty.



### Βήμα - βήμα

1. [x] Εγγραφή στην [υπηρεσία](https://www1.aade.gr/webtax/wspublicreg/faces/pages/wspublicreg/menu.xhtml) κάνοντας χρήση των κωδικών TAXISnet.
2. [x] Απόκτηση ειδικών κωδικών πρόσβασης μέσω της εφαρμογής [Διαχείριση Ειδικών Κωδικών](https://www1.aade.gr/sgsisapps/tokenservices/protected/displayConsole.htm).
3. [x] Χρήση ένος προγράμματος της αρεσκείας σας για την [κλήση της υπηρεσίας](https://www.aade.gr/sites/default/files/2018-07/AadeWebServiceRgWsPublicV401Client.zip).


### Τα βασικά χαρακτηριστικά της υπηρεσίας είναι:

* Η υπηρεσία μπορεί να αξιοποιηθεί απ’ όλους τους πιστοποιημένους χρήστες του TAXISnet.
* Υπάρχει μηνιαίο όριο κλήσεων της υπηρεσίας.
* Ο ΑΦΜ τα στοιχεία του οποίου αναζητούνται, ενημερώνεται με ειδική ειδοποίηση, για το ΑΦΜ / ονοματεπώνυμο που έκανε την αναζήτηση.
* Μέσω της οθόνης εγγραφής στην υπηρεσία μπορεί κάποιος να εξουσιοδοτήσει ένα τρίτο ΑΦΜ να καλεί την υπηρεσία γι’ αυτόν.

**Τα WSDL / ENDPOINT / XSD της αναβαθμισμένης υπηρεσίας είναι:**

* WSDL		: https://www1.gsis.gr/wsaade/RgWsPublic2/RgWsPublic2?WSDL 
* ENDPOINT	: https://www1.gsis.gr/wsaade/RgWsPublic2/RgWsPublic2 
* XSD			: https://www1.gsis.gr/wsaade/RgWsPublic2/RgWsPublic2?xsd=1 

Πρόκειται για Soap JAX-WS 2.0 Web Service (έκδοσης SOAP 1.2).

Για να καλέσει ένας σταθμός εργασίας την υπηρεσία απαιτείται δικτυακή πρόσβαση στο www1.gsis.gr και στο port 443. 

Εφόσον γίνει χρήση Java, απαιτείται χρήση Java 1.8 ή μεταγενέστερη λόγω της χρήσης του πρωτοκόλλου επικοινωνίας TLS1.2.

Περιλαμβάνονται:
a) παραδείγματα κλήσης (Request XML / Response XML) του Web Service, 
b) ένα SoapUI project για να γίνει import στο SoapUI. Προτείνεται χρήση SoapUI Version 5.4.0 ή μεταγενέστερη λόγω της Java 1.8 ( https://www.soapui.org/downloads/latest-release.html ).