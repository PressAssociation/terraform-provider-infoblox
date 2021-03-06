package records

const wapiVersion = "/wapi/v2.6.1"

// GenericRecord : GenericRecord data structure
type GenericRecord struct {
	Ref       string `json:"_ref,omitempty"`
	Name      string `json:"name,omitempty"`
	View      string `json:"view,omitempty"`
	TTL       uint   `json:"ttl,omitempty"`
	UseTTL    *bool  `json:"use_ttl,omitempty"`
	Comment   string `json:"comment,omitempty"`
	IPv4      string `json:"ipv4addr,omitempty"`
	Canonical string `json:"canonical,omitempty"`
	Text      string `json:"text,omitempty"`
	Port      int    `json:"port,omitempty"`
	Priority  int    `json:"priority,omitempty"`
	Target    string `json:"target,omitempty"`
	Weight    int    `json:"weight,omitempty"`
}

// ARecord : ARecord data structure
type ARecord struct {
	Ref     string `json:"_ref,omitempty"`
	IPv4    string `json:"ipv4addr,omitempty"`
	Name    string `json:"name,omitempty"`
	View    string `json:"view,omitempty"`
	Zone    string `json:"zone,omitempty"`
	TTL     uint   `json:"ttl,omitempty"`
	UseTTL  *bool  `json:"use_ttl,omitempty"`
	Comment string `json:"comment,omitempty"`
}

// CNAMERecord : CNAMERecord data structure
type CNAMERecord struct {
	Ref       string `json:"_ref,omitempty"`
	Canonical string `json:"canonical,omitempty"`
	Name      string `json:"name,omitempty"`
	View      string `json:"view,omitempty"`
	Zone      string `json:"zone,omitempty"`
	TTL       uint   `json:"ttl,omitempty"`
	UseTTL    *bool  `json:"use_ttl,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

// TXTRecord : TXTRecord data structure
type TXTRecord struct {
	Ref     string `json:"_ref,omitempty"`
	Name    string `json:"name,omitempty"`
	Text    string `json:"text,omitempty"`
	View    string `json:"view,omitempty"`
	Zone    string `json:"zone,omitempty"`
	TTL     uint   `json:"ttl,omitempty"`
	UseTTL  *bool  `json:"use_ttl,omitemply"`
	Comment string `json:"comment,omitempty"`
}

// SRVRecord : SRVRecord data structure
type SRVRecord struct {
	Ref      string `json:"_ref,omitempty"`
	Name     string `json:"name,omitempty"`
	Port     int    `json:"port,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Target   string `json:"target,omitempty"`
	View     string `json:"view,omitempty"`
	Weight   int    `json:"weight,omitempty"`
	Zone     string `json:"zone,omitempty"`
	TTL      uint   `json:"ttl,omitempty"`
	UseTTL   *bool  `json:"use_ttl,omitempty"`
	Comment  string `json:"comment,omitempty"`
}
