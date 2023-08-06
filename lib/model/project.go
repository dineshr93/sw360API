package model

type LinkedProjects struct {
	Project   string `json:"project"`
	EnableSvm string `json:"enableSvm"`
	Relation  string `json:"relation"`
}

type LinkedReleases struct {
	CreatedBy     string `json:"createdBy"`
	Release       string `json:"release"`
	MainlineState string `json:"mainlineState"`
	Comment       string `json:"comment"`
	CreatedOn     string `json:"createdOn"`
	Relation      string `json:"relation"`
}

type ProjectDetail struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version"`
	Domain      string `json:"domain,omitempty"`
	ExternalIds struct {
	} `json:"externalIds,omitempty"`
	AdditionalData struct {
	} `json:"additionalData,omitempty"`
	CreatedOn          string `json:"createdOn,omitempty"`
	BusinessUnit       string `json:"businessUnit,omitempty"`
	State              string `json:"state,omitempty"`
	Tag                string `json:"tag,omitempty"`
	ClearingState      string `json:"clearingState,omitempty"`
	ProjectResponsible string `json:"projectResponsible,omitempty"`
	Roles              struct {
	} `json:"roles,omitempty"`
	SecurityResponsibles          []interface{}    `json:"securityResponsibles"`
	ProjectOwner                  string           `json:"projectOwner,omitempty"`
	OwnerAccountingUnit           string           `json:"ownerAccountingUnit,omitempty"`
	OwnerGroup                    string           `json:"ownerGroup,omitempty"`
	OwnerCountry                  string           `json:"ownerCountry,omitempty"`
	PreevaluationDeadline         string           `json:"preevaluationDeadline,omitempty"`
	SystemTestStart               string           `json:"systemTestStart,omitempty"`
	SystemTestEnd                 string           `json:"systemTestEnd,omitempty"`
	DeliveryStart                 string           `json:"deliveryStart,omitempty"`
	PhaseOutSince                 string           `json:"phaseOutSince,omitempty"`
	EnableSvm                     bool             `json:"enableSvm,omitempty"`
	EnableVulnerabilitiesDisplay  bool             `json:"enableVulnerabilitiesDisplay"`
	ClearingSummary               string           `json:"clearingSummary,omitempty"`
	SpecialRisksOSS               string           `json:"specialRisksOSS,omitempty"`
	GeneralRisks3RdParty          string           `json:"generalRisks3rdParty,omitempty"`
	SpecialRisks3RdParty          string           `json:"specialRisks3rdParty,omitempty"`
	DeliveryChannels              string           `json:"deliveryChannels,omitempty"`
	RemarksAdditionalRequirements string           `json:"remarksAdditionalRequirements,omitempty"`
	ProjectType                   string           `json:"projectType,omitempty"`
	Visibility                    string           `json:"visibility,omitempty"`
	LinkedProjects                []LinkedProjects `json:"linkedProjects,omitempty"`
	LinkedReleases                []LinkedReleases `json:"linkedReleases,omitempty"`
	Links                         struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Embedded struct {
		ClearingTeam string `json:"clearingTeam"`
		CreatedBy    string `json:"createdBy"`
	} `json:"_embedded,omitempty"`
	VendorID string `json:"vendorId,omitempty"`
}

type Project struct {
	Embedded struct {
		Sw360Projects []ProjectDetail `json:"sw360:projects,omitempty"`
	} `json:"_embedded,omitempty"`
	Links struct {
		Curies []struct {
			Href      string `json:"href,omitempty"`
			Name      string `json:"name,omitempty"`
			Templated bool   `json:"templated,omitempty"`
		} `json:"curies,omitempty"`
	} `json:"_links,omitempty"`
}
