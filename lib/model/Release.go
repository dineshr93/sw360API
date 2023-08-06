package model

type SW360ReleaseShort struct {
	Embedded struct {
		Sw360Releases []struct {
			Name    string `json:"name,omitempty"`
			Version string `json:"version,omitempty"`
			Links   struct {
				Self struct {
					Href string `json:"href,omitempty"`
				} `json:"self,omitempty"`
			} `json:"_links,omitempty"`
		} `json:"sw360:releases,omitempty"`
	} `json:"_embedded,omitempty"`
	Links struct {
		Curies []struct {
			Href      string `json:"href,omitempty"`
			Name      string `json:"name,omitempty"`
			Templated bool   `json:"templated,omitempty"`
		} `json:"curies,omitempty"`
	} `json:"_links,omitempty"`
}

type Release struct {
	Name                  string         `json:"name,omitempty"`
	Version               string         `json:"version,omitempty"`
	ReleaseDate           string         `json:"releaseDate,omitempty"`
	ComponentType         string         `json:"componentType,omitempty"`
	ExternalIds           ExternalIds    `json:"externalIds,omitempty"`
	AdditionalData        AdditionalData `json:"additionalData,omitempty"`
	CreatedOn             string         `json:"createdOn,omitempty"`
	MainlineState         string         `json:"mainlineState,omitempty"`
	ClearingState         string         `json:"clearingState,omitempty"`
	CreatedBy             string         `json:"createdBy,omitempty"`
	Contributors          []any          `json:"contributors,omitempty"`
	Subscribers           []any          `json:"subscribers,omitempty"`
	Roles                 Roles          `json:"roles,omitempty"`
	OtherLicenseIds       []any          `json:"otherLicenseIds,omitempty"`
	Languages             []any          `json:"languages,omitempty"`
	OperatingSystems      []any          `json:"operatingSystems,omitempty"`
	SoftwarePlatforms     []any          `json:"softwarePlatforms,omitempty"`
	SourceCodeDownloadurl string         `json:"sourceCodeDownloadurl,omitempty"`
	BinaryDownloadurl     string         `json:"binaryDownloadurl,omitempty"`
	CpeID                 string         `json:"cpeId,omitempty"`
	EccInformation        EccInformation `json:"eccInformation,omitempty"`
	Links                 Links          `json:"_links,omitempty"`
	Embedded              Embedded       `json:"_embedded,omitempty"`
}
type ExternalIds struct {
}
type AdditionalData struct {
}
type Roles struct {
}
type EccInformation struct {
	Al        string `json:"al,omitempty"`
	Eccn      string `json:"eccn,omitempty"`
	EccStatus string `json:"eccStatus,omitempty"`
}
type Sw360Component struct {
	Href string `json:"href,omitempty"`
}
type Self struct {
	Href string `json:"href,omitempty"`
}
type Curies struct {
	Href      string `json:"href,omitempty"`
	Name      string `json:"name,omitempty"`
	Templated bool   `json:"templated,omitempty"`
}
type Links struct {
	Sw360Component Sw360Component `json:"sw360:component,omitempty"`
	Self           Self           `json:"self,omitempty"`
	Curies         []Curies       `json:"curies,omitempty"`
}

type Sw360Attachments struct {
	Filename       string `json:"filename,omitempty"`
	Sha1           string `json:"sha1,omitempty"`
	AttachmentType string `json:"attachmentType,omitempty"`
	Links          Links  `json:"_links,omitempty"`
}
type Embedded struct {
	Sw360Attachments []Sw360Attachments `json:"sw360:attachments,omitempty"`
}
