package main
import (
    "context"
    "fmt"
    "os"
    "time"
    "github.com/oracle/oci-go-sdk/common"
    "github.com/oracle/oci-go-sdk/identity"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {

    tenancy := os.Getenv("tenancy")
    fmt.Printf("oci tenancy ocid: %s\n", tenancy)

    user := os.Getenv("user")
    fmt.Printf("oci user ocid: %s\n", user)

    region := os.Getenv("region")
    fmt.Printf("oci region: %s\n", region)

    fingerprint := os.Getenv("fingerprint")
    fmt.Printf("fingerprint: %s\n", fingerprint)

    ppkcontent := os.Getenv("ppkfile")

	ppkpasswd := os.Getenv("ppkpassword")
    fmt.Printf("ppk passwd: %s\n", "****")

    fmt.Printf("1\n")
	nrcp := common.NewRawConfigurationProvider(tenancy, user, region, fingerprint, string(ppkcontent), &ppkpasswd)
	c, err := identity.NewIdentityClientWithConfigurationProvider(nrcp)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
  
    fmt.Printf("2\n")
    tenancyID, err := nrcp.TenancyOCID()
    if err != nil {
        fmt.Println("Error:", err)
        return
        }
  
    request := identity.ListAvailabilityDomainsRequest{
        CompartmentId: &tenancyID,
    }
    fmt.Printf("3\n")
    r, err := c.ListAvailabilityDomains(context.Background(), request)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("4\n")
    fmt.Printf("List of available domains: %v\n", r.Items)

    fmt.Printf("sleeping...\n")
    time.Sleep(time.Second * 8640000000)// cien mil días

    return
}