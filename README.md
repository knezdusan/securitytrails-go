# SecurityTrails-GOlang-Wrapper
A GOlang based wrapper for SecurityTrails

Based off https://docs.securitytrails.com/reference

<!-- Work in progress: API Wrapper Function Documentation: https://github.com/knezdusan/SecurityTrails-GOlang-Wrapper/wiki -->

## Usage examples

### checkPing: Test your authentication and access to the SecurityTrails API.

```
package main
import "fmt"

func main() {
  st := newST("YOUR_API_KEY_HERE")
  res := st.checkPing()
  fmt.Println(res)
}
```

### getDomainWHOIS: Returns the current WHOIS data about a given hostname

```
package main
import "fmt"

func main() {
  st := newST("YOUR_API_KEY_HERE")

  // Get the json response from the API endpoint in the 'res' variable by simply calling the referring sdk method
  // res := st.sdkMethod(optional parametars), example:
  res := st.getDomainWHOIS("google.com")

  fmt.Println(res)
}
```

## Implemented methods

### General

- [X] Ping
- [X] Usage

### Company

- [X] Details
- [X] Associated IPs

### Domain

- [X] Details
- [X] Subdomains
- [X] Tags
- [X] WHOIS
- [ ] Search
- [ ] Statistics
- [X] Associated domains
- [X] SSL Certificates (Pages)
- [X] SSL Certificates (Stream)

### History

- [X] DNS History
- [ ] WHOIS History

### IPs

- [X] Neighbors
- [ ] Search with DSL
- [ ] Statistics
- [X] Whois
- [X] Useragents

### Feeds

- [ ] Domains
- [ ] DMARC
- [ ] Subdomains

### Firehose

- [X] Certificate Transparency
