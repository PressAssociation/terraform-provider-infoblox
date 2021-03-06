package dhcprange

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// DeleteDHCPRangeAPI base object.
type DeleteDHCPRangeAPI struct {
	*api.BaseAPI
}

// NewDeleteDHCPRange returns a new object of type DeleteNetworkAPI.
func NewDeleteDHCPRange(objRef string) *DeleteDHCPRangeAPI {
	this := new(DeleteDHCPRangeAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, fmt.Sprintf("%s/%s", wapiVersion, objRef), nil, new(string))
	return this
}

// GetResponse casts the response object and
// returns the single network object
func (gn DeleteDHCPRangeAPI) GetResponse() string {
	return *gn.ResponseObject().(*string)
}
