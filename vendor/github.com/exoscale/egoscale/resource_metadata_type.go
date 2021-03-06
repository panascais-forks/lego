package egoscale

// ListResourceDetails lists the resource tag(s) (but different from listTags...)
//
// CloudStack API: http://cloudstack.apache.org/api/apidocs-4.10/apis/listResourceDetails.html
type ListResourceDetails struct {
	ResourceType string `json:"resourcetype" doc:"list by resource type"`
	Account      string `json:"account,omitempty" doc:"list resources by account. Must be used with the domainId parameter."`
	DomainID     string `json:"domainid,omitempty" doc:"list only resources belonging to the domain specified"`
	ForDisplay   bool   `json:"fordisplay,omitempty" doc:"if set to true, only details marked with display=true, are returned. False by default"`
	Key          string `json:"key,omitempty" doc:"list by key"`
	Keyword      string `json:"keyword,omitempty" doc:"List by keyword"`
	ListAll      bool   `json:"listall,omitempty" doc:"If set to false, list only resources belonging to the command's caller; if set to true - list resources that the caller is authorized to see. Default value is false"`
	Page         int    `json:"page,omitempty"`
	PageSize     int    `json:"pagesize,omitempty"`
	ResourceID   string `json:"resourceid,omitempty" doc:"list by resource id"`
	Value        string `json:"value,omitempty" doc:"list by key, value. Needs to be passed only along with key"`
	IsRecursive  bool   `json:"isrecursive,omitempty" doc:"defaults to false, but if true, lists all resources from the parent specified by the domainId till leaves."`
}

// ListResourceDetailsResponse represents a list of resource details
type ListResourceDetailsResponse struct {
	Count          int           `json:"count"`
	ResourceDetail []ResourceTag `json:"resourcedetail"`
}
