package main

import "fmt"

func main() {
	st := newST("YOUR.APP.KEY.HERE")

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
	// res := st.getSSLCertificatesStream("stackoverflow.com", false, "valid")


	// History *************************************
	// // Path Param:
	// // type:string | Allowed values: a, aaaa, mx, ns, soa or txt
	// // Query Param:
	// // page:int32 | The page of the returned results, starting at 1. A page returns 100 results.
	// res := st.getDNSHistory("oracle.com", "a", 1)

	// // Query Param:
	// // page:int32 | The page of the returned results, starting at 1. A page returns 100 results.
	// res := st.getWHOISHistory("oracle.com", 1)
	
	
	// IPs *************************************
	// res := st.getIPNeighbors("8.8.8.8")
	// res := st.getIPWHOIS("1.1.1.1")
	// res := st.getIPUseragents("93.80.234.159", 1)


	// Firehose *************************************
	res := st.getFirehoseCertificateTransparency(1628089450, 1633359850)
	
	fmt.Println(res)
}
