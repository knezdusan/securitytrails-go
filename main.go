package main

import "fmt"

func main() {
	st := newST("c0LxAP033JLunaGxxzkdNUU7OoF8J70M")

	// General *******************************
	// res := st.checkPing()
	// res := st.checkUsage()
	
	// Company ******************************* 
	// res := st.getCompanyDetails("google.com")
	// res := st.getCompanyAssociatedIPs("google.com")
	
	// Domains ******************************* 
	// res := st.getDomainDetails("google.com")
	// res := st.getDomainSubdomains("google.com")
	// res := st.getDomainTags("google.com")
	// res := st.getDomainWHOIS("google.com")
	// res := st.getAssociatedDomains("securitytrails.com")

	// SSL *************************************
	// // Query Params:
	// // include_subdomains:boolean | Default is false.
	// // status:string | Valid values are "valid", "all", and "expired". Default is valid.
	// // page:int32 | The page of the returned results, starting at 1. A page returns 100 results.
	// // https://api.securitytrails.com/v1/domain/stackoverflow.com/ssl?include_subdomains=false&status=valid&page=1
	// res := st.getSSLCertificatesPages("stackoverflow.com", false, "valid", 1)

	// // Query Params:
	// // include_subdomains:boolean | Default is false.
	// // status:string | Valid values are "valid", "all", and "expired". Default is valid.
	// // page:int32 | The page of the returned results, starting at 1. A page returns 100 results.
	// // https://api.securitytrails.com/v1/domain/stackoverflow.com/ssl?include_subdomains=false&status=valid&page=1
	res := st.getSSLCertificatesStream("stackoverflow.com", false, "valid")


	
	fmt.Println(res)
}