/**
 * USAGE
 * Example of your main app file -> main.go:
	``````````
	package main
	import "fmt"

  func main() {
	  st := newST("YOUR_API_KEY_HERE")

		// Get the json response from the API endpoint in the 'res' variable by simply calling the referring sdk method
		// res := st.sdkMethod(optional parametars), example:
		res := st.checkPing()

		// Print res to console:
		fmt.Println(res)
  }
  ````````````
**/

// Change the package name if necessary to corespond to your main app package name.
package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

type stsdk struct {
	key string
}

// Make new sdk instance
func newST(key string) stsdk {
	st := stsdk{
		key: key,
	}

	return st
}

// *************  Receiver functions to match ST Api Reference endpoints funcionalities: https://docs.securitytrails.com/reference

// Ping - You can use this simple endpoint to test your authentication and access to the SecurityTrails API.
func (st *stsdk) checkPing() string{
	url := "https://api.securitytrails.com/v1/ping"
	
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", st.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

// Usage - statistics of the API usage for the current month
func (st *stsdk) checkUsage() string{
	url := "https://api.securitytrails.com/v1/account/usage"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", st.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

// Details - Returns details for a company domain.
func (st *stsdk) getCompanyDetails(domain string) string{
	url := "https://api.securitytrails.com/v1/company/" + domain

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", st.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

// Associated IPs - Returns not paginated nor limited associated IPs for a company domain based on whois data. 
func (st *stsdk) getCompanyAssociatedIPs(domain string) string{
	url := "https://api.securitytrails.com/v1/company/" + domain + "/associated-ips"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", st.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
} 

// Domain Details - Returns the current data about the given hostname with statistics associated with a particular record.
func (st *stsdk) getDomainDetails(domain string) string{
	url := "https://api.securitytrails.com/v1/domain/" + domain

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", st.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

// Subdomains Details - Returns child and sibling subdomains for a given hostname.
func (st *stsdk) getDomainSubdomains(domain string) string{
	url := "https://api.securitytrails.com/v1/domain/" + domain + "/subdomains"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", st.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

// Tags - Returns tags for a given hostname
func (st *stsdk) getDomainTags(domain string) string{
	url := "https://api.securitytrails.com/v1/domain/" + domain + "/tags"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", st.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

// WHOIS - Returns the current WHOIS data about a given hostname with the stats merged together
func (st *stsdk) getDomainWHOIS(domain string) string{
	url := "https://api.securitytrails.com/v1/domain/" + domain + "/whois"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", st.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

// Associated domains - Find all domains that are related to a hostname you input
func (st *stsdk) getAssociatedDomains(domain string) string{
	url := "https://api.securitytrails.com/v1/domain/" + domain + "/associated"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", st.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}


// SSL Certificates (Pages) - Fetch current and historical certificate information for any hostname.
// Query Params:
// include_subdomains:boolean | Default is false.
// status:string | Valid values are "valid", "all", and "expired". Default is valid.
// page:int32 | The page of the returned results, starting at 1. A page returns 100 results.
func (st *stsdk) getSSLCertificatesPages(domain string, include_subdomains bool, status string, page_number int32) string{
	url := "https://api.securitytrails.com/v1/domain/" + domain + "/ssl?include_subdomains=" + strconv.FormatBool(include_subdomains) + "&status=" + status + "&page=" + strconv.FormatInt(int64(page_number), 10)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", st.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

// SSL Certificates (Stream) - Fetch current and historical certificate information for any hostname.
// Query Params:
// include_subdomains:boolean | Default is false.
// status:string | Valid values are "valid", "all", and "expired". Default is valid.
func (st *stsdk) getSSLCertificatesStream(domain string, include_subdomains bool, status string) string{
	url := "https://api.securitytrails.com/v1/domain/" + domain + "/ssl_stream?include_subdomains=" + strconv.FormatBool(include_subdomains) + "&status=" + status

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", st.key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}


