package factual

// PlaceWrap allows unmarshalling of API response metadata
type PlaceWrap struct {
	Version  int           `json:"version"`
	Status   string        `json:"status"`
	Response PlaceResponse `json:"response"`
}

// PlaceResponse is schema returned by the Factual API
type PlaceResponse struct {
	Data         []Place `json:"data"`
	IncludedRows int     `json:"included_rows"`
}

// Hours describes the open hours for a place
type Hours struct {
	Monday    [][]string `json:"monday"`
	Tuesday   [][]string `json:"tuesday"`
	Wednesday [][]string `json:"wednesday"`
	Thursday  [][]string `json:"thursday"`
	Friday    [][]string `json:"friday"`
	Saturday  [][]string `json:"saturday"`
	Sunday    [][]string `json:"sunday"`
}

// Place contains the properties of a Place
type Place struct {
	FactualID      string     `json:"factual_id"`
	CategoryIDs    []int      `json:"category_ids"`
	CategoryLabels [][]string `json:"category_labels"`
	ChainID        string     `json:"chain_id,omitempty"`
	ChainName      string     `json:"chain_name,omitempty"`

	Address  string `json:"address"`
	Country  string `json:"country"`
	Postcode string `json:"postcode"`
	Fax      string `json:"fac"`
	Region   string `json:"region"`
	Tel      string `json:"tel"`

	Hours        Hours    `json:"hours"`
	HoursDisplay string   `json:"hours_display"`
	Latitude     float64  `json:"latitude"`
	Longitude    float64  `json:"latitude"`
	Locality     string   `json:"locality"`
	Name         string   `json:"name"`
	Neighborhood []string `json:"neighborhood"`
	Website      string   `json:"website"`
}
