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
	// res := st.getFirehoseCertificateTransparency(1628089450, 1633359850)


	// **************** POST Queries with parameters: ********************************************

	// Search ********************************
	// // Filter and search specific records using this endpoint.
	// // Sample Use Cases: Search for all hostnames that point to your IP address
	// var ips = true
	// var page_number int64 = 1
	// var scroll = true
	
	// queryData := map[string]string{
	// 	"ipv4": "",
	// 	"ipv6": "",
	// 	"apex_domain": "google.com",
	// 	"keyword": "",
	// 	"mx": "",
	// 	"ns": "",
	// 	"cname": "",
	// 	"subdomain": "mail",
	// 	"soa_email": "",
	// 	"tld": "com",
	// 	"whois_email": "",
	// 	"whois_street1": "",
	// 	"whois_telephone": "",
	// 	"whois_postalCode": "",
	// 	"whois_organization": "",
	// 	"whois_name": "",
	// 	"whois_fax": "",
	// 	"whois_city": "",
	// 	"whois_country": "",
	// }

	// res := st.postDomainSearch(ips, page_number, scroll, queryData)


	// Statistics ********************************
	// // Filter and search specific records using this endpoint.
	// // Sample Use Cases: Search for all hostnames that point to your IP address
	
	// queryData := map[string]string{
	// 	"ipv4": "",
	// 	"ipv6": "",
	// 	"apex_domain": "google.com",
	// 	"keyword": "",
	// 	"mx": "",
	// 	"ns": "",
	// 	"cname": "",
	// 	"subdomain": "",
	// 	"soa_email": "",
	// 	"tld": "com",
	// 	"whois_email": "",
	// 	"whois_street1": "",
	// 	"whois_telephone": "",
	// 	"whois_postalCode": "",
	// 	"whois_organization": "",
	// 	"whois_name": "",
	// 	"whois_fax": "",
	// 	"whois_city": "",
	// 	"whois_country": "",
	// }

	// res := st.postDomainStatistics(queryData)


	// Search with DSL ********************************
	// // Search for IP addresses. A maximum of 10000 results can be retrieved.
	// // Use simple JSON like query syntax.
	// // Use advanced SecurityTrails DSL SQL like queries: https://docs.securitytrails.com/docs/how-to-use-the-dsl

	// var page_number int64 = 1
	// dslQuery := `{"query":"ptr_part = 'stackoverflow.com'"}`

	// res := st.postDslSearch(page_number, dslQuery)


	// Statistics with DSL ********************************
	// // Statistics like Reverse DNS pattern identification, ports or total results are returned.
	// // Use simple JSON like query syntax.
	// // Use advanced SecurityTrails DSL SQL like queries: https://docs.securitytrails.com/docs/how-to-use-the-dsl

	dslQuery := `{"query":"ptr_part = 'amazon.com'"}`

	res := st.postDslStatistics(dslQuery)
	
	fmt.Println(res)
}
