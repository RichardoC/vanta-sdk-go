// Vanta service endpoint bindings.

package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// Services is the generated service registry for Vanta APIs.
type Services struct {
	Controls                  *ControlsService
	DiscoveredVendors         *DiscoveredVendorsService
	Documents                 *DocumentsService
	Frameworks                *FrameworksService
	Groups                    *GroupsService
	Integrations              *IntegrationsService
	MonitoredComputers        *MonitoredComputersService
	OAuth                     *OAuthService
	People                    *PeopleService
	Policies                  *PoliciesService
	Resources                 *ResourcesService
	RiskScenarios             *RiskScenariosService
	Tests                     *TestsService
	TrustCenters              *TrustCentersService
	VendorRiskAttributes      *VendorRiskAttributesService
	Vendors                   *VendorsService
	Vulnerabilities           *VulnerabilitiesService
	VulnerabilityRemediations *VulnerabilityRemediationsService
	VulnerableAssets          *VulnerableAssetsService
}

func newGeneratedServices(c *Client) *Services {
	services := &Services{}
	services.Controls = &ControlsService{client: c}
	services.DiscoveredVendors = &DiscoveredVendorsService{client: c}
	services.Documents = &DocumentsService{client: c}
	services.Frameworks = &FrameworksService{client: c}
	services.Groups = &GroupsService{client: c}
	services.Integrations = &IntegrationsService{client: c}
	services.MonitoredComputers = &MonitoredComputersService{client: c}
	services.OAuth = &OAuthService{client: c}
	services.People = &PeopleService{client: c}
	services.Policies = &PoliciesService{client: c}
	services.Resources = &ResourcesService{client: c}
	services.RiskScenarios = &RiskScenariosService{client: c}
	services.Tests = &TestsService{client: c}
	services.TrustCenters = &TrustCentersService{client: c}
	services.VendorRiskAttributes = &VendorRiskAttributesService{client: c}
	services.Vendors = &VendorsService{client: c}
	services.Vulnerabilities = &VulnerabilitiesService{client: c}
	services.VulnerabilityRemediations = &VulnerabilityRemediationsService{client: c}
	services.VulnerableAssets = &VulnerableAssetsService{client: c}
	return services
}

// ControlsService groups 14 endpoint methods under the "Controls" API segment.
type ControlsService struct {
	client *Client
}

type ControlsAddControlFromVantaLibraryRequestBody struct {
	ControlID string `json:"controlId"`
}

type ControlsAddControlFromVantaLibraryResponse struct {
	CreationDate     any              `json:"creationDate"`
	CustomFields     []map[string]any `json:"customFields"`
	Description      string           `json:"description"`
	Domains          []string         `json:"domains"`
	ExternalID       string           `json:"externalId"`
	ID               string           `json:"id"`
	ModificationDate any              `json:"modificationDate"`
	Name             string           `json:"name"`
	Owner            map[string]any   `json:"owner"`
	Role             string           `json:"role"`
	Source           string           `json:"source"`
}

type ControlsAddControlFromVantaLibraryParams struct {
	Body *ControlsAddControlFromVantaLibraryRequestBody
}

// AddControlFromVantaLibrary Add a control from the Vanta library to your organization's controls.
func (s *ControlsService) AddControlFromVantaLibrary(ctx context.Context, params *ControlsAddControlFromVantaLibraryParams) (*ControlsAddControlFromVantaLibraryResponse, error) {
	if params == nil {
		params = &ControlsAddControlFromVantaLibraryParams{}
	}
	path := "/controls/add-from-library"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &ControlsAddControlFromVantaLibraryResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsAddControlToDocumentMappingRequestBody struct {
	DocumentID string `json:"documentId"`
}

type ControlsAddControlToDocumentMappingResponse struct {
	Control  map[string]any `json:"control"`
	Document map[string]any `json:"document"`
}

type ControlsAddControlToDocumentMappingParams struct {
	ControlID string
	Body      *ControlsAddControlToDocumentMappingRequestBody
}

// AddControlToDocumentMapping Add a document to a control.
func (s *ControlsService) AddControlToDocumentMapping(ctx context.Context, params *ControlsAddControlToDocumentMappingParams) (*ControlsAddControlToDocumentMappingResponse, error) {
	if params == nil {
		params = &ControlsAddControlToDocumentMappingParams{}
	}
	path := "/controls/:controlId/add-document-to-control"
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &ControlsAddControlToDocumentMappingResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsAddControlToTestMappingRequestBody struct {
	TestID string `json:"testId"`
}

type ControlsAddControlToTestMappingResponse struct {
	Control map[string]any `json:"control"`
	Test    map[string]any `json:"test"`
}

type ControlsAddControlToTestMappingParams struct {
	ControlID string
	Body      *ControlsAddControlToTestMappingRequestBody
}

// AddControlToTestMapping Add a control to test mapping.
func (s *ControlsService) AddControlToTestMapping(ctx context.Context, params *ControlsAddControlToTestMappingParams) (*ControlsAddControlToTestMappingResponse, error) {
	if params == nil {
		params = &ControlsAddControlToTestMappingParams{}
	}
	path := "/controls/:controlId/add-test-to-control"
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &ControlsAddControlToTestMappingResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsCreateCustomControlRequestBody struct {
	CustomFields  []map[string]any `json:"customFields"`
	Description   string           `json:"description"`
	Domain        string           `json:"domain"`
	EffectiveDate string           `json:"effectiveDate"`
	ExternalID    string           `json:"externalId"`
	Name          string           `json:"name"`
	Role          string           `json:"role"`
	Sections      []map[string]any `json:"sections"`
}

type ControlsCreateCustomControlResponse struct {
	CreationDate     any              `json:"creationDate"`
	CustomFields     []map[string]any `json:"customFields"`
	Description      string           `json:"description"`
	Domains          []string         `json:"domains"`
	ExternalID       string           `json:"externalId"`
	ID               string           `json:"id"`
	ModificationDate any              `json:"modificationDate"`
	Name             string           `json:"name"`
	Owner            map[string]any   `json:"owner"`
	Role             string           `json:"role"`
	Source           string           `json:"source"`
}

type ControlsCreateCustomControlParams struct {
	Body *ControlsCreateCustomControlRequestBody
}

// CreateCustomControl Create a custom control.
func (s *ControlsService) CreateCustomControl(ctx context.Context, params *ControlsCreateCustomControlParams) (*ControlsCreateCustomControlResponse, error) {
	if params == nil {
		params = &ControlsCreateCustomControlParams{}
	}
	path := "/controls"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &ControlsCreateCustomControlResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsGetControlByIDResponse struct {
	CreationDate        any              `json:"creationDate"`
	CustomFields        []map[string]any `json:"customFields"`
	Description         string           `json:"description"`
	Domains             []string         `json:"domains"`
	ExternalID          string           `json:"externalId"`
	ID                  string           `json:"id"`
	ModificationDate    any              `json:"modificationDate"`
	Name                string           `json:"name"`
	Note                string           `json:"note"`
	NumDocumentsPassing float64          `json:"numDocumentsPassing"`
	NumDocumentsTotal   float64          `json:"numDocumentsTotal"`
	NumTestsPassing     float64          `json:"numTestsPassing"`
	NumTestsTotal       float64          `json:"numTestsTotal"`
	Owner               map[string]any   `json:"owner"`
	Role                string           `json:"role"`
	Source              string           `json:"source"`
	Status              string           `json:"status"`
}

type ControlsGetControlByIDParams struct {
	ControlID string
}

// GetControlByID Get a control by an ID.
func (s *ControlsService) GetControlByID(ctx context.Context, params *ControlsGetControlByIDParams) (*ControlsGetControlByIDResponse, error) {
	if params == nil {
		params = &ControlsGetControlByIDParams{}
	}
	path := "/controls/:controlId"
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &ControlsGetControlByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsListControlsResponse struct {
	Results map[string]any `json:"results"`
}

type ControlsListControlsParams struct {
	PageSize            *int
	PageCursor          *string
	FrameworkMatchesAny []string
}

// ListControls List controls.
func (s *ControlsService) ListControls(ctx context.Context, params *ControlsListControlsParams) (*ControlsListControlsResponse, error) {
	if params == nil {
		params = &ControlsListControlsParams{}
	}
	path := "/controls"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	for _, v := range params.FrameworkMatchesAny {
		query.Add("frameworkMatchesAny", v)
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &ControlsListControlsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsListControlsDocumentsResponse struct {
	Results map[string]any `json:"results"`
}

type ControlsListControlsDocumentsParams struct {
	ControlID  string
	PageSize   *int
	PageCursor *string
}

// ListControlsDocuments List a control's documents.
func (s *ControlsService) ListControlsDocuments(ctx context.Context, params *ControlsListControlsDocumentsParams) (*ControlsListControlsDocumentsResponse, error) {
	if params == nil {
		params = &ControlsListControlsDocumentsParams{}
	}
	path := "/controls/:controlId/documents"
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &ControlsListControlsDocumentsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsListControlsTestsResponse struct {
	Results map[string]any `json:"results"`
}

type ControlsListControlsTestsParams struct {
	ControlID  string
	PageSize   *int
	PageCursor *string
}

// ListControlsTests List a control's tests.
func (s *ControlsService) ListControlsTests(ctx context.Context, params *ControlsListControlsTestsParams) (*ControlsListControlsTestsResponse, error) {
	if params == nil {
		params = &ControlsListControlsTestsParams{}
	}
	path := "/controls/:controlId/tests"
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &ControlsListControlsTestsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsListVantaControlsFromLibraryResponse struct {
	Results map[string]any `json:"results"`
}

type ControlsListVantaControlsFromLibraryParams struct {
	PageSize   *int
	PageCursor *string
}

// ListVantaControlsFromLibrary List Vanta controls from the library.
func (s *ControlsService) ListVantaControlsFromLibrary(ctx context.Context, params *ControlsListVantaControlsFromLibraryParams) (*ControlsListVantaControlsFromLibraryResponse, error) {
	if params == nil {
		params = &ControlsListVantaControlsFromLibraryParams{}
	}
	path := "/controls/controls-library"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &ControlsListVantaControlsFromLibraryResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsRemoveControlParams struct {
	ControlID string
}

// RemoveControl Delete a custom control or move a Vanta control back to the library.
func (s *ControlsService) RemoveControl(ctx context.Context, params *ControlsRemoveControlParams) (json.RawMessage, error) {
	if params == nil {
		params = &ControlsRemoveControlParams{}
	}
	path := "/controls/:controlId"
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsRemoveControlFromDocumentMappingParams struct {
	ControlID  string
	DocumentID string
}

// RemoveControlFromDocumentMapping Remove a document by ID from a control.
func (s *ControlsService) RemoveControlFromDocumentMapping(ctx context.Context, params *ControlsRemoveControlFromDocumentMappingParams) (json.RawMessage, error) {
	if params == nil {
		params = &ControlsRemoveControlFromDocumentMappingParams{}
	}
	path := "/controls/:controlId/documents/:documentId"
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsRemoveControlFromTestMappingParams struct {
	ControlID string
	TestID    string
}

// RemoveControlFromTestMapping Remove a control from test mapping.
func (s *ControlsService) RemoveControlFromTestMapping(ctx context.Context, params *ControlsRemoveControlFromTestMappingParams) (json.RawMessage, error) {
	if params == nil {
		params = &ControlsRemoveControlFromTestMappingParams{}
	}
	path := "/controls/:controlId/tests/:testId"
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	if params.TestID == "" {
		return nil, fmt.Errorf("testId is required")
	}
	path = strings.ReplaceAll(path, ":testId", url.PathEscape(params.TestID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsSetOwnerOfControlRequestBody struct {
	UserID string `json:"userId"`
}

type ControlsSetOwnerOfControlResponse struct {
	CreationDate     any              `json:"creationDate"`
	CustomFields     []map[string]any `json:"customFields"`
	Description      string           `json:"description"`
	Domains          []string         `json:"domains"`
	ExternalID       string           `json:"externalId"`
	ID               string           `json:"id"`
	ModificationDate any              `json:"modificationDate"`
	Name             string           `json:"name"`
	Owner            map[string]any   `json:"owner"`
	Role             string           `json:"role"`
	Source           string           `json:"source"`
}

type ControlsSetOwnerOfControlParams struct {
	ControlID string
	Body      *ControlsSetOwnerOfControlRequestBody
}

// SetOwnerOfControl Assign a control to a user or remove an owner from a control.
func (s *ControlsService) SetOwnerOfControl(ctx context.Context, params *ControlsSetOwnerOfControlParams) (*ControlsSetOwnerOfControlResponse, error) {
	if params == nil {
		params = &ControlsSetOwnerOfControlParams{}
	}
	path := "/controls/:controlId/set-owner"
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &ControlsSetOwnerOfControlResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type ControlsUpdateControlsMetadataRequestBody struct {
	CustomFields []map[string]any `json:"customFields"`
	Description  string           `json:"description"`
	Domain       string           `json:"domain"`
	ExternalID   string           `json:"externalId"`
	Name         string           `json:"name"`
	Note         string           `json:"note"`
}

type ControlsUpdateControlsMetadataResponse struct {
	CreationDate     any              `json:"creationDate"`
	CustomFields     []map[string]any `json:"customFields"`
	Description      string           `json:"description"`
	Domains          []string         `json:"domains"`
	ExternalID       string           `json:"externalId"`
	ID               string           `json:"id"`
	ModificationDate any              `json:"modificationDate"`
	Name             string           `json:"name"`
	Owner            map[string]any   `json:"owner"`
	Role             string           `json:"role"`
	Source           string           `json:"source"`
}

type ControlsUpdateControlsMetadataParams struct {
	ControlID string
	Body      *ControlsUpdateControlsMetadataRequestBody
}

// UpdateControlsMetadata Update a control's metadata.
func (s *ControlsService) UpdateControlsMetadata(ctx context.Context, params *ControlsUpdateControlsMetadataParams) (*ControlsUpdateControlsMetadataResponse, error) {
	if params == nil {
		params = &ControlsUpdateControlsMetadataParams{}
	}
	path := "/controls/:controlId"
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &ControlsUpdateControlsMetadataResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveredVendorsService groups 3 endpoint methods under the "DiscoveredVendors" API segment.
type DiscoveredVendorsService struct {
	client *Client
}

type DiscoveredVendorsAddsDiscoveredVendorToManagedVendorByIDResponse struct {
	AccountManagerEmail              string         `json:"accountManagerEmail"`
	AccountManagerName               string         `json:"accountManagerName"`
	AdditionalNotes                  string         `json:"additionalNotes"`
	AuthDetails                      map[string]any `json:"authDetails"`
	BusinessOwnerUserID              string         `json:"businessOwnerUserId"`
	Category                         map[string]any `json:"category"`
	ContractAmount                   map[string]any `json:"contractAmount"`
	ContractRenewalDate              string         `json:"contractRenewalDate"`
	ContractStartDate                string         `json:"contractStartDate"`
	ContractTerminationDate          any            `json:"contractTerminationDate"`
	CustomFields                     any            `json:"customFields"`
	ID                               string         `json:"id"`
	InherentRiskLevel                string         `json:"inherentRiskLevel"`
	IsRiskAutoScored                 bool           `json:"isRiskAutoScored"`
	IsVisibleToAuditors              bool           `json:"isVisibleToAuditors"`
	LastSecurityReviewCompletionDate string         `json:"lastSecurityReviewCompletionDate"`
	LatestDecision                   map[string]any `json:"latestDecision"`
	Name                             string         `json:"name"`
	NextSecurityReviewDueDate        string         `json:"nextSecurityReviewDueDate"`
	ResidualRiskLevel                string         `json:"residualRiskLevel"`
	RiskAttributeIDs                 []string       `json:"riskAttributeIds"`
	SecurityOwnerUserID              string         `json:"securityOwnerUserId"`
	ServicesProvided                 string         `json:"servicesProvided"`
	Status                           string         `json:"status"`
	TagIDentifiers                   any            `json:"tagIdentifiers"`
	VendorHeadquarters               string         `json:"vendorHeadquarters"`
	WebsiteURL                       string         `json:"websiteUrl"`
}

type DiscoveredVendorsAddsDiscoveredVendorToManagedVendorByIDParams struct {
	DiscoveredVendorID string
}

// AddsDiscoveredVendorToManagedVendorByID Add a discovered vendor to managed vendor.
func (s *DiscoveredVendorsService) AddsDiscoveredVendorToManagedVendorByID(ctx context.Context, params *DiscoveredVendorsAddsDiscoveredVendorToManagedVendorByIDParams) (*DiscoveredVendorsAddsDiscoveredVendorToManagedVendorByIDResponse, error) {
	if params == nil {
		params = &DiscoveredVendorsAddsDiscoveredVendorToManagedVendorByIDParams{}
	}
	path := "/discovered-vendors/:discoveredVendorId/add-to-managed"
	if params.DiscoveredVendorID == "" {
		return nil, fmt.Errorf("discoveredVendorId is required")
	}
	path = strings.ReplaceAll(path, ":discoveredVendorId", url.PathEscape(params.DiscoveredVendorID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &DiscoveredVendorsAddsDiscoveredVendorToManagedVendorByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DiscoveredVendorsListDiscoveredVendorsResponse struct {
	Results map[string]any `json:"results"`
}

type DiscoveredVendorsListDiscoveredVendorsParams struct {
	Scope      *string
	PageSize   *int
	PageCursor *string
}

// ListDiscoveredVendors List discovered vendors.
func (s *DiscoveredVendorsService) ListDiscoveredVendors(ctx context.Context, params *DiscoveredVendorsListDiscoveredVendorsParams) (*DiscoveredVendorsListDiscoveredVendorsResponse, error) {
	if params == nil {
		params = &DiscoveredVendorsListDiscoveredVendorsParams{}
	}
	path := "/discovered-vendors"
	query := url.Values{}
	if params.Scope != nil {
		query.Set("scope", fmt.Sprint(*params.Scope))
	}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &DiscoveredVendorsListDiscoveredVendorsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DiscoveredVendorsListOfDiscoveredVendorAccountsResponse struct {
	Results map[string]any `json:"results"`
}

type DiscoveredVendorsListOfDiscoveredVendorAccountsParams struct {
	DiscoveredVendorID string
	PageSize           *int
	PageCursor         *string
}

// ListOfDiscoveredVendorAccounts List of discovered vendor accounts.
func (s *DiscoveredVendorsService) ListOfDiscoveredVendorAccounts(ctx context.Context, params *DiscoveredVendorsListOfDiscoveredVendorAccountsParams) (*DiscoveredVendorsListOfDiscoveredVendorAccountsResponse, error) {
	if params == nil {
		params = &DiscoveredVendorsListOfDiscoveredVendorAccountsParams{}
	}
	path := "/discovered-vendors/:discoveredVendorId/accounts"
	if params.DiscoveredVendorID == "" {
		return nil, fmt.Errorf("discoveredVendorId is required")
	}
	path = strings.ReplaceAll(path, ":discoveredVendorId", url.PathEscape(params.DiscoveredVendorID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &DiscoveredVendorsListOfDiscoveredVendorAccountsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentsService groups 14 endpoint methods under the "Documents" API segment.
type DocumentsService struct {
	client *Client
}

type DocumentsCreateCustomDocumentRequestBody struct {
	Cadence         string `json:"cadence"`
	Description     string `json:"description"`
	IsSensitive     bool   `json:"isSensitive"`
	ReminderWindow  string `json:"reminderWindow"`
	TimeSensitivity string `json:"timeSensitivity"`
	Title           string `json:"title"`
}

type DocumentsCreateCustomDocumentResponse struct {
	Category         string `json:"category"`
	Description      string `json:"description"`
	ID               string `json:"id"`
	IsSensitive      bool   `json:"isSensitive"`
	OwnerID          string `json:"ownerId"`
	Title            string `json:"title"`
	URL              string `json:"url"`
	UploadStatus     string `json:"uploadStatus"`
	UploadStatusDate string `json:"uploadStatusDate"`
}

type DocumentsCreateCustomDocumentParams struct {
	Body *DocumentsCreateCustomDocumentRequestBody
}

// CreateCustomDocument Create a custom document.
func (s *DocumentsService) CreateCustomDocument(ctx context.Context, params *DocumentsCreateCustomDocumentParams) (*DocumentsCreateCustomDocumentResponse, error) {
	if params == nil {
		params = &DocumentsCreateCustomDocumentParams{}
	}
	path := "/documents"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &DocumentsCreateCustomDocumentResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsCreateDocumentLinkRequestBody struct {
	Description   string `json:"description"`
	EffectiveDate string `json:"effectiveDate"`
	Title         string `json:"title"`
	URL           string `json:"url"`
}

type DocumentsCreateDocumentLinkResponse struct {
	CreationDate  string `json:"creationDate"`
	Description   string `json:"description"`
	EffectiveDate string `json:"effectiveDate"`
	ID            string `json:"id"`
	Title         string `json:"title"`
	URL           string `json:"url"`
}

type DocumentsCreateDocumentLinkParams struct {
	DocumentID string
	Body       *DocumentsCreateDocumentLinkRequestBody
}

// CreateDocumentLink Create a link for a document.
func (s *DocumentsService) CreateDocumentLink(ctx context.Context, params *DocumentsCreateDocumentLinkParams) (*DocumentsCreateDocumentLinkResponse, error) {
	if params == nil {
		params = &DocumentsCreateDocumentLinkParams{}
	}
	path := "/documents/:documentId/links"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &DocumentsCreateDocumentLinkResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsDeleteDocumentByIDParams struct {
	DocumentID string
}

// DeleteDocumentByID Delete a document by ID.
func (s *DocumentsService) DeleteDocumentByID(ctx context.Context, params *DocumentsDeleteDocumentByIDParams) (json.RawMessage, error) {
	if params == nil {
		params = &DocumentsDeleteDocumentByIDParams{}
	}
	path := "/documents/:documentId"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsDeleteFileForDocumentParams struct {
	DocumentID     string
	UploadedFileID string
}

// DeleteFileForDocument Delete a file for a document.
func (s *DocumentsService) DeleteFileForDocument(ctx context.Context, params *DocumentsDeleteFileForDocumentParams) (json.RawMessage, error) {
	if params == nil {
		params = &DocumentsDeleteFileForDocumentParams{}
	}
	path := "/documents/:documentId/uploads/:uploadedFileId"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	if params.UploadedFileID == "" {
		return nil, fmt.Errorf("uploadedFileId is required")
	}
	path = strings.ReplaceAll(path, ":uploadedFileId", url.PathEscape(params.UploadedFileID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsDownloadFileForDocumentResponse struct {
	Readable bool `json:"readable"`
}

type DocumentsDownloadFileForDocumentParams struct {
	DocumentID     string
	UploadedFileID string
}

// DownloadFileForDocument Download a file from a document.
func (s *DocumentsService) DownloadFileForDocument(ctx context.Context, params *DocumentsDownloadFileForDocumentParams) (*DocumentsDownloadFileForDocumentResponse, error) {
	if params == nil {
		params = &DocumentsDownloadFileForDocumentParams{}
	}
	path := "/documents/:documentId/uploads/:uploadedFileId/media"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	if params.UploadedFileID == "" {
		return nil, fmt.Errorf("uploadedFileId is required")
	}
	path = strings.ReplaceAll(path, ":uploadedFileId", url.PathEscape(params.UploadedFileID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &DocumentsDownloadFileForDocumentResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsGetDocumentByIDResponse struct {
	Category          string         `json:"category"`
	DeactivatedStatus map[string]any `json:"deactivatedStatus"`
	Description       string         `json:"description"`
	ID                string         `json:"id"`
	IsSensitive       bool           `json:"isSensitive"`
	NextRenewalDate   string         `json:"nextRenewalDate"`
	Note              string         `json:"note"`
	OwnerID           string         `json:"ownerId"`
	ReminderWindow    string         `json:"reminderWindow"`
	RenewalCadence    string         `json:"renewalCadence"`
	Subscribers       []any          `json:"subscribers"`
	Title             string         `json:"title"`
	URL               string         `json:"url"`
	UploadStatus      string         `json:"uploadStatus"`
	UploadStatusDate  string         `json:"uploadStatusDate"`
}

type DocumentsGetDocumentByIDParams struct {
	DocumentID string
}

// GetDocumentByID Get a document by ID.
func (s *DocumentsService) GetDocumentByID(ctx context.Context, params *DocumentsGetDocumentByIDParams) (*DocumentsGetDocumentByIDResponse, error) {
	if params == nil {
		params = &DocumentsGetDocumentByIDParams{}
	}
	path := "/documents/:documentId"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &DocumentsGetDocumentByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsListDocumentsResponse struct {
	Results map[string]any `json:"results"`
}

type DocumentsListDocumentsParams struct {
	PageSize            *int
	PageCursor          *string
	FrameworkMatchesAny []string
	StatusMatchesAny    []string
}

// ListDocuments List documents.
func (s *DocumentsService) ListDocuments(ctx context.Context, params *DocumentsListDocumentsParams) (*DocumentsListDocumentsResponse, error) {
	if params == nil {
		params = &DocumentsListDocumentsParams{}
	}
	path := "/documents"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	for _, v := range params.FrameworkMatchesAny {
		query.Add("frameworkMatchesAny", v)
	}
	for _, v := range params.StatusMatchesAny {
		query.Add("statusMatchesAny", v)
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &DocumentsListDocumentsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsListDocumentsControlsResponse struct {
	Results map[string]any `json:"results"`
}

type DocumentsListDocumentsControlsParams struct {
	DocumentID string
	PageSize   *int
	PageCursor *string
}

// ListDocumentsControls List a document's associated controls.
func (s *DocumentsService) ListDocumentsControls(ctx context.Context, params *DocumentsListDocumentsControlsParams) (*DocumentsListDocumentsControlsResponse, error) {
	if params == nil {
		params = &DocumentsListDocumentsControlsParams{}
	}
	path := "/documents/:documentId/controls"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &DocumentsListDocumentsControlsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsListDocumentsLinksResponse struct {
	Results map[string]any `json:"results"`
}

type DocumentsListDocumentsLinksParams struct {
	DocumentID string
	PageSize   *int
	PageCursor *string
}

// ListDocumentsLinks List the uploaded links for a document.
func (s *DocumentsService) ListDocumentsLinks(ctx context.Context, params *DocumentsListDocumentsLinksParams) (*DocumentsListDocumentsLinksResponse, error) {
	if params == nil {
		params = &DocumentsListDocumentsLinksParams{}
	}
	path := "/documents/:documentId/links"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &DocumentsListDocumentsLinksResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsListDocumentsUploadsResponse struct {
	Results map[string]any `json:"results"`
}

type DocumentsListDocumentsUploadsParams struct {
	DocumentID string
	PageSize   *int
	PageCursor *string
}

// ListDocumentsUploads List the uploaded files for a document.
func (s *DocumentsService) ListDocumentsUploads(ctx context.Context, params *DocumentsListDocumentsUploadsParams) (*DocumentsListDocumentsUploadsResponse, error) {
	if params == nil {
		params = &DocumentsListDocumentsUploadsParams{}
	}
	path := "/documents/:documentId/uploads"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &DocumentsListDocumentsUploadsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsRemoveDocumentLinkParams struct {
	DocumentID string
	LinkID     string
}

// RemoveDocumentLink Remove a link from a document.
func (s *DocumentsService) RemoveDocumentLink(ctx context.Context, params *DocumentsRemoveDocumentLinkParams) (json.RawMessage, error) {
	if params == nil {
		params = &DocumentsRemoveDocumentLinkParams{}
	}
	path := "/documents/:documentId/links/:linkId"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	if params.LinkID == "" {
		return nil, fmt.Errorf("linkId is required")
	}
	path = strings.ReplaceAll(path, ":linkId", url.PathEscape(params.LinkID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsSetDocumentOwnerRequestBody struct {
	UserID string `json:"userId"`
}

type DocumentsSetDocumentOwnerResponse struct {
	Category         string `json:"category"`
	Description      string `json:"description"`
	ID               string `json:"id"`
	IsSensitive      bool   `json:"isSensitive"`
	OwnerID          string `json:"ownerId"`
	Title            string `json:"title"`
	URL              string `json:"url"`
	UploadStatus     string `json:"uploadStatus"`
	UploadStatusDate string `json:"uploadStatusDate"`
}

type DocumentsSetDocumentOwnerParams struct {
	DocumentID string
	Body       *DocumentsSetDocumentOwnerRequestBody
}

// SetDocumentOwner Assign or unassign a user to the document.
func (s *DocumentsService) SetDocumentOwner(ctx context.Context, params *DocumentsSetDocumentOwnerParams) (*DocumentsSetDocumentOwnerResponse, error) {
	if params == nil {
		params = &DocumentsSetDocumentOwnerParams{}
	}
	path := "/documents/:documentId/set-owner"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &DocumentsSetDocumentOwnerResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsSubmitDocumentCollectionParams struct {
	DocumentID string
}

// SubmitDocumentCollection Submit document collection.
func (s *DocumentsService) SubmitDocumentCollection(ctx context.Context, params *DocumentsSubmitDocumentCollectionParams) (json.RawMessage, error) {
	if params == nil {
		params = &DocumentsSubmitDocumentCollectionParams{}
	}
	path := "/documents/:documentId/submit"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type DocumentsUploadFileForDocumentResponse struct {
	CreationDate  string         `json:"creationDate"`
	DeletionDate  any            `json:"deletionDate"`
	Description   string         `json:"description"`
	EffectiveDate string         `json:"effectiveDate"`
	FileName      string         `json:"fileName"`
	ID            string         `json:"id"`
	MimeType      string         `json:"mimeType"`
	Title         string         `json:"title"`
	URL           string         `json:"url"`
	UpdatedDate   string         `json:"updatedDate"`
	UploadedBy    map[string]any `json:"uploadedBy"`
}

type DocumentsUploadFileForDocumentParams struct {
	DocumentID string
	// FormData maps multipart field names to values.
	FormData map[string]string
}

// UploadFileForDocument Upload a file for a document.
func (s *DocumentsService) UploadFileForDocument(ctx context.Context, params *DocumentsUploadFileForDocumentParams) (*DocumentsUploadFileForDocumentResponse, error) {
	if params == nil {
		params = &DocumentsUploadFileForDocumentParams{}
	}
	path := "/documents/:documentId/uploads"
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	query := url.Values{}
	req, err := s.client.newMultipartRequest(ctx, "POST", path, query, params.FormData)
	if err != nil {
		return nil, err
	}
	out := &DocumentsUploadFileForDocumentResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// FrameworksService groups 3 endpoint methods under the "Frameworks" API segment.
type FrameworksService struct {
	client *Client
}

type FrameworksGetFrameworkByIDResponse struct {
	Description           string           `json:"description"`
	DisplayName           string           `json:"displayName"`
	ID                    string           `json:"id"`
	NumControlsCompleted  float64          `json:"numControlsCompleted"`
	NumControlsTotal      float64          `json:"numControlsTotal"`
	NumDocumentsPassing   float64          `json:"numDocumentsPassing"`
	NumDocumentsTotal     float64          `json:"numDocumentsTotal"`
	NumTestsPassing       float64          `json:"numTestsPassing"`
	NumTestsTotal         float64          `json:"numTestsTotal"`
	RequirementCategories []map[string]any `json:"requirementCategories"`
	ShorthandName         string           `json:"shorthandName"`
}

type FrameworksGetFrameworkByIDParams struct {
	FrameworkID string
}

// GetFrameworkByID Get a framework by ID.
func (s *FrameworksService) GetFrameworkByID(ctx context.Context, params *FrameworksGetFrameworkByIDParams) (*FrameworksGetFrameworkByIDResponse, error) {
	if params == nil {
		params = &FrameworksGetFrameworkByIDParams{}
	}
	path := "/frameworks/:frameworkId"
	if params.FrameworkID == "" {
		return nil, fmt.Errorf("frameworkId is required")
	}
	path = strings.ReplaceAll(path, ":frameworkId", url.PathEscape(params.FrameworkID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &FrameworksGetFrameworkByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type FrameworksListAvailableFrameworksResponse struct {
	Results map[string]any `json:"results"`
}

type FrameworksListAvailableFrameworksParams struct {
	PageSize   *int
	PageCursor *string
}

// ListAvailableFrameworks Lists available frameworks.
func (s *FrameworksService) ListAvailableFrameworks(ctx context.Context, params *FrameworksListAvailableFrameworksParams) (*FrameworksListAvailableFrameworksResponse, error) {
	if params == nil {
		params = &FrameworksListAvailableFrameworksParams{}
	}
	path := "/frameworks"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &FrameworksListAvailableFrameworksResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type FrameworksListFrameworksControlsResponse struct {
	Results map[string]any `json:"results"`
}

type FrameworksListFrameworksControlsParams struct {
	FrameworkID string
	PageSize    *int
	PageCursor  *string
}

// ListFrameworksControls List a framework's controls.
func (s *FrameworksService) ListFrameworksControls(ctx context.Context, params *FrameworksListFrameworksControlsParams) (*FrameworksListFrameworksControlsResponse, error) {
	if params == nil {
		params = &FrameworksListFrameworksControlsParams{}
	}
	path := "/frameworks/:frameworkId/controls"
	if params.FrameworkID == "" {
		return nil, fmt.Errorf("frameworkId is required")
	}
	path = strings.ReplaceAll(path, ":frameworkId", url.PathEscape(params.FrameworkID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &FrameworksListFrameworksControlsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GroupsService groups 7 endpoint methods under the "Groups" API segment.
type GroupsService struct {
	client *Client
}

type GroupsAddPeopleToGroupRequestBody struct {
	Updates []map[string]any `json:"updates"`
}

type GroupsAddPeopleToGroupResponse struct {
	Results []map[string]any `json:"results"`
}

type GroupsAddPeopleToGroupParams struct {
	GroupID string
	Body    *GroupsAddPeopleToGroupRequestBody
}

// AddPeopleToGroup Add people to a group.
func (s *GroupsService) AddPeopleToGroup(ctx context.Context, params *GroupsAddPeopleToGroupParams) (*GroupsAddPeopleToGroupResponse, error) {
	if params == nil {
		params = &GroupsAddPeopleToGroupParams{}
	}
	path := "/groups/:groupId/add-people"
	if params.GroupID == "" {
		return nil, fmt.Errorf("groupId is required")
	}
	path = strings.ReplaceAll(path, ":groupId", url.PathEscape(params.GroupID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &GroupsAddPeopleToGroupResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type GroupsAddPersonToGroupRequestBody struct {
	ID string `json:"id"`
}

type GroupsAddPersonToGroupParams struct {
	GroupID string
	Body    *GroupsAddPersonToGroupRequestBody
}

// AddPersonToGroup Add a single person, by ID, to a group.
func (s *GroupsService) AddPersonToGroup(ctx context.Context, params *GroupsAddPersonToGroupParams) (*Person, error) {
	if params == nil {
		params = &GroupsAddPersonToGroupParams{}
	}
	path := "/groups/:groupId/people"
	if params.GroupID == "" {
		return nil, fmt.Errorf("groupId is required")
	}
	path = strings.ReplaceAll(path, ":groupId", url.PathEscape(params.GroupID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &Person{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type GroupsGetGroupByIDResponse struct {
	CreationDate string `json:"creationDate"`
	ID           string `json:"id"`
	Name         string `json:"name"`
}

type GroupsGetGroupByIDParams struct {
	GroupID string
}

// GetGroupByID Get a group by ID.
func (s *GroupsService) GetGroupByID(ctx context.Context, params *GroupsGetGroupByIDParams) (*GroupsGetGroupByIDResponse, error) {
	if params == nil {
		params = &GroupsGetGroupByIDParams{}
	}
	path := "/groups/:groupId"
	if params.GroupID == "" {
		return nil, fmt.Errorf("groupId is required")
	}
	path = strings.ReplaceAll(path, ":groupId", url.PathEscape(params.GroupID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &GroupsGetGroupByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type GroupsListGroupsResponse struct {
	Results map[string]any `json:"results"`
}

type GroupsListGroupsParams struct {
	PageSize   *int
	PageCursor *string
}

// ListGroups Lists all groups by ID.
func (s *GroupsService) ListGroups(ctx context.Context, params *GroupsListGroupsParams) (*GroupsListGroupsResponse, error) {
	if params == nil {
		params = &GroupsListGroupsParams{}
	}
	path := "/groups"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &GroupsListGroupsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type GroupsListPeopleInGroupParams struct {
	GroupID string
}

// ListPeopleInGroup List people in a group.
func (s *GroupsService) ListPeopleInGroup(ctx context.Context, params *GroupsListPeopleInGroupParams) ([]Person, error) {
	if params == nil {
		params = &GroupsListPeopleInGroupParams{}
	}
	path := "/groups/:groupId/people"
	if params.GroupID == "" {
		return nil, fmt.Errorf("groupId is required")
	}
	path = strings.ReplaceAll(path, ":groupId", url.PathEscape(params.GroupID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := make([]Person, 0)
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type GroupsRemovePeopleFromGroupRequestBody struct {
	Updates []map[string]any `json:"updates"`
}

type GroupsRemovePeopleFromGroupResponse struct {
	Results []map[string]any `json:"results"`
}

type GroupsRemovePeopleFromGroupParams struct {
	GroupID string
	Body    *GroupsRemovePeopleFromGroupRequestBody
}

// RemovePeopleFromGroup Remove people from a group.
func (s *GroupsService) RemovePeopleFromGroup(ctx context.Context, params *GroupsRemovePeopleFromGroupParams) (*GroupsRemovePeopleFromGroupResponse, error) {
	if params == nil {
		params = &GroupsRemovePeopleFromGroupParams{}
	}
	path := "/groups/:groupId/remove-people"
	if params.GroupID == "" {
		return nil, fmt.Errorf("groupId is required")
	}
	path = strings.ReplaceAll(path, ":groupId", url.PathEscape(params.GroupID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &GroupsRemovePeopleFromGroupResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type GroupsRemovePersonFromGroupParams struct {
	GroupID  string
	PersonID string
}

// RemovePersonFromGroup Remove a single person, by ID, from a group.
func (s *GroupsService) RemovePersonFromGroup(ctx context.Context, params *GroupsRemovePersonFromGroupParams) (*Person, error) {
	if params == nil {
		params = &GroupsRemovePersonFromGroupParams{}
	}
	path := "/groups/:groupId/people/:personId"
	if params.GroupID == "" {
		return nil, fmt.Errorf("groupId is required")
	}
	path = strings.ReplaceAll(path, ":groupId", url.PathEscape(params.GroupID))
	if params.PersonID == "" {
		return nil, fmt.Errorf("personId is required")
	}
	path = strings.ReplaceAll(path, ":personId", url.PathEscape(params.PersonID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &Person{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// IntegrationsService groups 8 endpoint methods under the "Integrations" API segment.
type IntegrationsService struct {
	client *Client
}

type IntegrationsGetConnectedIntegrationResponse struct {
	Connections   []map[string]any `json:"connections"`
	DisplayName   string           `json:"displayName"`
	IntegrationID string           `json:"integrationId"`
	ResourceKinds []string         `json:"resourceKinds"`
}

type IntegrationsGetConnectedIntegrationParams struct {
	IntegrationID string
}

// GetConnectedIntegration Gets details for a specific integration by connection ID.
func (s *IntegrationsService) GetConnectedIntegration(ctx context.Context, params *IntegrationsGetConnectedIntegrationParams) (*IntegrationsGetConnectedIntegrationResponse, error) {
	if params == nil {
		params = &IntegrationsGetConnectedIntegrationParams{}
	}
	path := "/integrations/:integrationId"
	if params.IntegrationID == "" {
		return nil, fmt.Errorf("integrationId is required")
	}
	path = strings.ReplaceAll(path, ":integrationId", url.PathEscape(params.IntegrationID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &IntegrationsGetConnectedIntegrationResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type IntegrationsGetDetailsForResourceKindResponse struct {
	CanUpdateDescription bool    `json:"canUpdateDescription"`
	CanUpdateOwner       bool    `json:"canUpdateOwner"`
	IntegrationID        string  `json:"integrationId"`
	IsScopable           bool    `json:"isScopable"`
	NumInScope           float64 `json:"numInScope"`
	NumOwned             float64 `json:"numOwned"`
	NumResources         float64 `json:"numResources"`
	NumWithDescription   float64 `json:"numWithDescription"`
	ResourceKind         string  `json:"resourceKind"`
}

type IntegrationsGetDetailsForResourceKindParams struct {
	IntegrationID string
	ResourceKind  string
	ConnectionID  *string
}

// GetDetailsForResourceKind Gets details for a specific resource type (kind) such as S3Bucket or CloudwatchLogGroup.
func (s *IntegrationsService) GetDetailsForResourceKind(ctx context.Context, params *IntegrationsGetDetailsForResourceKindParams) (*IntegrationsGetDetailsForResourceKindResponse, error) {
	if params == nil {
		params = &IntegrationsGetDetailsForResourceKindParams{}
	}
	path := "/integrations/:integrationId/resource-kinds/:resourceKind"
	if params.IntegrationID == "" {
		return nil, fmt.Errorf("integrationId is required")
	}
	path = strings.ReplaceAll(path, ":integrationId", url.PathEscape(params.IntegrationID))
	if params.ResourceKind == "" {
		return nil, fmt.Errorf("resourceKind is required")
	}
	path = strings.ReplaceAll(path, ":resourceKind", url.PathEscape(params.ResourceKind))
	query := url.Values{}
	if params.ConnectionID != nil {
		query.Set("connectionId", fmt.Sprint(*params.ConnectionID))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &IntegrationsGetDetailsForResourceKindResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type IntegrationsGetResourceByIDResponse struct {
	ConnectionID string `json:"connectionId"`
	CreationDate string `json:"creationDate"`
	Description  any    `json:"description"`
	DisplayName  string `json:"displayName"`
	InScope      bool   `json:"inScope"`
	Owner        any    `json:"owner"`
	ResourceID   string `json:"resourceId"`
	ResourceKind string `json:"resourceKind"`
	ResponseType string `json:"responseType"`
}

type IntegrationsGetResourceByIDParams struct {
	IntegrationID string
	ResourceKind  string
	ResourceID    string
}

// GetResourceByID Gets resource by its ID.
func (s *IntegrationsService) GetResourceByID(ctx context.Context, params *IntegrationsGetResourceByIDParams) (*IntegrationsGetResourceByIDResponse, error) {
	if params == nil {
		params = &IntegrationsGetResourceByIDParams{}
	}
	path := "/integrations/:integrationId/resource-kinds/:resourceKind/resources/:resourceId"
	if params.IntegrationID == "" {
		return nil, fmt.Errorf("integrationId is required")
	}
	path = strings.ReplaceAll(path, ":integrationId", url.PathEscape(params.IntegrationID))
	if params.ResourceKind == "" {
		return nil, fmt.Errorf("resourceKind is required")
	}
	path = strings.ReplaceAll(path, ":resourceKind", url.PathEscape(params.ResourceKind))
	if params.ResourceID == "" {
		return nil, fmt.Errorf("resourceId is required")
	}
	path = strings.ReplaceAll(path, ":resourceId", url.PathEscape(params.ResourceID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &IntegrationsGetResourceByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type IntegrationsListConnectedIntegrationsResponse struct {
	Results map[string]any `json:"results"`
}

type IntegrationsListConnectedIntegrationsParams struct {
	PageSize   *int
	PageCursor *string
}

// ListConnectedIntegrations Lists all integrations connected to a Vanta instance.
func (s *IntegrationsService) ListConnectedIntegrations(ctx context.Context, params *IntegrationsListConnectedIntegrationsParams) (*IntegrationsListConnectedIntegrationsResponse, error) {
	if params == nil {
		params = &IntegrationsListConnectedIntegrationsParams{}
	}
	path := "/integrations"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &IntegrationsListConnectedIntegrationsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type IntegrationsListIntegrationResourceKindsParams struct {
	IntegrationID string
}

// ListIntegrationResourceKinds Lists a connected integration's resource types (kinds) such as S3Bucket or CloudwatchLogGroup.
func (s *IntegrationsService) ListIntegrationResourceKinds(ctx context.Context, params *IntegrationsListIntegrationResourceKindsParams) (json.RawMessage, error) {
	if params == nil {
		params = &IntegrationsListIntegrationResourceKindsParams{}
	}
	path := "/integrations/:integrationId/resource-kinds"
	if params.IntegrationID == "" {
		return nil, fmt.Errorf("integrationId is required")
	}
	path = strings.ReplaceAll(path, ":integrationId", url.PathEscape(params.IntegrationID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type IntegrationsListResourcesResponse struct {
	Results map[string]any `json:"results"`
}

type IntegrationsListResourcesParams struct {
	IntegrationID  string
	ResourceKind   string
	ConnectionID   *string
	HasDescription *bool
	HasOwner       *bool
	IsInScope      *bool
	PageSize       *int
	PageCursor     *string
}

// ListResources Lists resources for a specific integration and resource type (kind) such as S3Bucket or CloudwatchLogGroup.
func (s *IntegrationsService) ListResources(ctx context.Context, params *IntegrationsListResourcesParams) (*IntegrationsListResourcesResponse, error) {
	if params == nil {
		params = &IntegrationsListResourcesParams{}
	}
	path := "/integrations/:integrationId/resource-kinds/:resourceKind/resources"
	if params.IntegrationID == "" {
		return nil, fmt.Errorf("integrationId is required")
	}
	path = strings.ReplaceAll(path, ":integrationId", url.PathEscape(params.IntegrationID))
	if params.ResourceKind == "" {
		return nil, fmt.Errorf("resourceKind is required")
	}
	path = strings.ReplaceAll(path, ":resourceKind", url.PathEscape(params.ResourceKind))
	query := url.Values{}
	if params.ConnectionID != nil {
		query.Set("connectionId", fmt.Sprint(*params.ConnectionID))
	}
	if params.HasDescription != nil {
		query.Set("hasDescription", fmt.Sprint(*params.HasDescription))
	}
	if params.HasOwner != nil {
		query.Set("hasOwner", fmt.Sprint(*params.HasOwner))
	}
	if params.IsInScope != nil {
		query.Set("isInScope", fmt.Sprint(*params.IsInScope))
	}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &IntegrationsListResourcesResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type IntegrationsUpdateResourceMetadataRequestBody struct {
	Updates []map[string]any `json:"updates"`
}

type IntegrationsUpdateResourceMetadataResponse struct {
	Results []map[string]any `json:"results"`
}

type IntegrationsUpdateResourceMetadataParams struct {
	IntegrationID string
	ResourceKind  string
	Body          *IntegrationsUpdateResourceMetadataRequestBody
}

// UpdateResourceMetadata Updates metadata for multiple resources.
func (s *IntegrationsService) UpdateResourceMetadata(ctx context.Context, params *IntegrationsUpdateResourceMetadataParams) (*IntegrationsUpdateResourceMetadataResponse, error) {
	if params == nil {
		params = &IntegrationsUpdateResourceMetadataParams{}
	}
	path := "/integrations/:integrationId/resource-kinds/:resourceKind/resources"
	if params.IntegrationID == "" {
		return nil, fmt.Errorf("integrationId is required")
	}
	path = strings.ReplaceAll(path, ":integrationId", url.PathEscape(params.IntegrationID))
	if params.ResourceKind == "" {
		return nil, fmt.Errorf("resourceKind is required")
	}
	path = strings.ReplaceAll(path, ":resourceKind", url.PathEscape(params.ResourceKind))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &IntegrationsUpdateResourceMetadataResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type IntegrationsUpdateResourceMetadataForResourceKindsResourcesRequestBody struct {
	Description string `json:"description"`
	InScope     bool   `json:"inScope"`
	OwnerID     string `json:"ownerId"`
}

type IntegrationsUpdateResourceMetadataForResourceKindsResourcesParams struct {
	IntegrationID string
	ResourceKind  string
	ResourceID    string
	Body          *IntegrationsUpdateResourceMetadataForResourceKindsResourcesRequestBody
}

// UpdateResourceMetadataForResourceKindsResources Updates metadata for a specific resource such as an S3Bucket or CloudwatchLogGroup.
func (s *IntegrationsService) UpdateResourceMetadataForResourceKindsResources(ctx context.Context, params *IntegrationsUpdateResourceMetadataForResourceKindsResourcesParams) (json.RawMessage, error) {
	if params == nil {
		params = &IntegrationsUpdateResourceMetadataForResourceKindsResourcesParams{}
	}
	path := "/integrations/:integrationId/resource-kinds/:resourceKind/resources/:resourceId"
	if params.IntegrationID == "" {
		return nil, fmt.Errorf("integrationId is required")
	}
	path = strings.ReplaceAll(path, ":integrationId", url.PathEscape(params.IntegrationID))
	if params.ResourceKind == "" {
		return nil, fmt.Errorf("resourceKind is required")
	}
	path = strings.ReplaceAll(path, ":resourceKind", url.PathEscape(params.ResourceKind))
	if params.ResourceID == "" {
		return nil, fmt.Errorf("resourceId is required")
	}
	path = strings.ReplaceAll(path, ":resourceId", url.PathEscape(params.ResourceID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// MonitoredComputersService groups 2 endpoint methods under the "MonitoredComputers" API segment.
type MonitoredComputersService struct {
	client *Client
}

type MonitoredComputersGetMonitoredComputerByIDResponse struct {
	AntivirusInstallation map[string]any `json:"antivirusInstallation"`
	DiskEncryption        map[string]any `json:"diskEncryption"`
	ID                    string         `json:"id"`
	IntegrationID         string         `json:"integrationId"`
	LastCheckDate         string         `json:"lastCheckDate"`
	OperatingSystem       map[string]any `json:"operatingSystem"`
	Owner                 map[string]any `json:"owner"`
	PasswordManager       map[string]any `json:"passwordManager"`
	Screenlock            map[string]any `json:"screenlock"`
	SerialNumber          string         `json:"serialNumber"`
	Udid                  string         `json:"udid"`
}

type MonitoredComputersGetMonitoredComputerByIDParams struct {
	ComputerID string
}

// GetMonitoredComputerByID Returns a monitored computer by ID.
func (s *MonitoredComputersService) GetMonitoredComputerByID(ctx context.Context, params *MonitoredComputersGetMonitoredComputerByIDParams) (*MonitoredComputersGetMonitoredComputerByIDResponse, error) {
	if params == nil {
		params = &MonitoredComputersGetMonitoredComputerByIDParams{}
	}
	path := "/monitored-computers/:computerId"
	if params.ComputerID == "" {
		return nil, fmt.Errorf("computerId is required")
	}
	path = strings.ReplaceAll(path, ":computerId", url.PathEscape(params.ComputerID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &MonitoredComputersGetMonitoredComputerByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type MonitoredComputersListMonitoredComputersResponse struct {
	Results map[string]any `json:"results"`
}

type MonitoredComputersListMonitoredComputersParams struct {
	PageSize                         *int
	PageCursor                       *string
	ComplianceStatusFilterMatchesAny []string
}

// ListMonitoredComputers Returns a list of computers monitored by an MDM (with an integration built by Vanta) or by the Vanta Agent. Currently this list does not include resources from partner or customer-built integrations.
func (s *MonitoredComputersService) ListMonitoredComputers(ctx context.Context, params *MonitoredComputersListMonitoredComputersParams) (*MonitoredComputersListMonitoredComputersResponse, error) {
	if params == nil {
		params = &MonitoredComputersListMonitoredComputersParams{}
	}
	path := "/monitored-computers"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	for _, v := range params.ComplianceStatusFilterMatchesAny {
		query.Add("complianceStatusFilterMatchesAny", v)
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &MonitoredComputersListMonitoredComputersResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthService groups 1 endpoint methods under the "OAuth" API segment.
type OAuthService struct {
	client *Client
}

type OAuthCreateTokenRequestBody struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
}

type OAuthCreateTokenParams struct {
	Body *OAuthCreateTokenRequestBody
}

// CreateToken CreateToken performs POST /oauth/token.
func (s *OAuthService) CreateToken(ctx context.Context, params *OAuthCreateTokenParams) (json.RawMessage, error) {
	if params == nil {
		params = &OAuthCreateTokenParams{}
	}
	path := "/oauth/token"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// PeopleService groups 8 endpoint methods under the "People" API segment.
type PeopleService struct {
	client *Client
}

type PeopleGetPersonByIDParams struct {
	PersonID string
}

// GetPersonByID Returns a person by ID.
func (s *PeopleService) GetPersonByID(ctx context.Context, params *PeopleGetPersonByIDParams) (*Person, error) {
	if params == nil {
		params = &PeopleGetPersonByIDParams{}
	}
	path := "/people/:personId"
	if params.PersonID == "" {
		return nil, fmt.Errorf("personId is required")
	}
	path = strings.ReplaceAll(path, ":personId", url.PathEscape(params.PersonID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &Person{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type PeopleListPeopleParams struct {
	PageSize                     *int
	PageCursor                   *string
	TasksSummaryStatusMatchesAny []string
	TaskTypeMatchesAny           []string
	TaskStatusMatchesAny         []string
}

// ListPeople Returns a list of all people.
func (s *PeopleService) ListPeople(ctx context.Context, params *PeopleListPeopleParams) (*PeopleListPeopleResponse, error) {
	if params == nil {
		params = &PeopleListPeopleParams{}
	}
	path := "/people"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	for _, v := range params.TasksSummaryStatusMatchesAny {
		query.Add("tasksSummaryStatusMatchesAny", v)
	}
	for _, v := range params.TaskTypeMatchesAny {
		query.Add("taskTypeMatchesAny", v)
	}
	for _, v := range params.TaskStatusMatchesAny {
		query.Add("taskStatusMatchesAny", v)
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &PeopleListPeopleResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type PeopleMarkAsNotPeopleRequestBody struct {
	Updates []map[string]any `json:"updates"`
}

type PeopleMarkAsNotPeopleResponse struct {
	Results []map[string]any `json:"results"`
}

type PeopleMarkAsNotPeopleParams struct {
	Body *PeopleMarkAsNotPeopleRequestBody
}

// MarkAsNotPeople Mark a set of accounts on the People Page as "not a person." As a result, these accounts will not be treated as people in Vanta, and you will not be able to assign them tasks or use them in tests related to your company's personnel.
func (s *PeopleService) MarkAsNotPeople(ctx context.Context, params *PeopleMarkAsNotPeopleParams) (*PeopleMarkAsNotPeopleResponse, error) {
	if params == nil {
		params = &PeopleMarkAsNotPeopleParams{}
	}
	path := "/people/mark-as-not-people"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &PeopleMarkAsNotPeopleResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type PeopleMarkAsPeopleRequestBody struct {
	Updates []map[string]any `json:"updates"`
}

type PeopleMarkAsPeopleResponse struct {
	Results []map[string]any `json:"results"`
}

type PeopleMarkAsPeopleParams struct {
	Body *PeopleMarkAsPeopleRequestBody
}

// MarkAsPeople Mark a set of accounts on the People Page as "people." As a result, these accounts will be treated as people in Vanta, and you will be able to assign them tasks and use them in tests related to your company's personnel.
func (s *PeopleService) MarkAsPeople(ctx context.Context, params *PeopleMarkAsPeopleParams) (*PeopleMarkAsPeopleResponse, error) {
	if params == nil {
		params = &PeopleMarkAsPeopleParams{}
	}
	path := "/people/mark-as-people"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &PeopleMarkAsPeopleResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type PeopleOffboardPeopleRequestBody struct {
	Updates []map[string]any `json:"updates"`
}

type PeopleOffboardPeopleResponse struct {
	Results []map[string]any `json:"results"`
}

type PeopleOffboardPeopleParams struct {
	Body *PeopleOffboardPeopleRequestBody
}

// OffboardPeople Offboard a list of people. A person is only eligible for offboarding completion when: 1. They are an ex-employee. 2. All of the person's monitored accounts are deactivated or manually overwritten as such. 3. All of a person's custom offboarding tasks have been completed. All of the person's unmonitored accounts will be automatically marked as deactivated when they are offboarded. If the person has unfinished offboarding tasks those will NOT automatically be completed and offboarding them will fail.
func (s *PeopleService) OffboardPeople(ctx context.Context, params *PeopleOffboardPeopleParams) (*PeopleOffboardPeopleResponse, error) {
	if params == nil {
		params = &PeopleOffboardPeopleParams{}
	}
	path := "/people/offboard"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &PeopleOffboardPeopleResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type PeopleRemoveLeaveInformationParams struct {
	PersonID string
}

// RemoveLeaveInformation Remove leave information on a person. The person will become active in Vanta, and will be considered in certain tests related to personnel.
func (s *PeopleService) RemoveLeaveInformation(ctx context.Context, params *PeopleRemoveLeaveInformationParams) (*Person, error) {
	if params == nil {
		params = &PeopleRemoveLeaveInformationParams{}
	}
	path := "/people/:personId/clear-leave"
	if params.PersonID == "" {
		return nil, fmt.Errorf("personId is required")
	}
	path = strings.ReplaceAll(path, ":personId", url.PathEscape(params.PersonID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &Person{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type PeopleSetLeaveInformationRequestBody struct {
	EndDate   string `json:"endDate"`
	StartDate string `json:"startDate"`
}

type PeopleSetLeaveInformationParams struct {
	PersonID string
	Body     *PeopleSetLeaveInformationRequestBody
}

// SetLeaveInformation Set leave information on a person. A person on leave is inactive in Vanta and will not be considered in certain personnel-related tests. If the person has existing leave information, it will be cleared and replaced.
func (s *PeopleService) SetLeaveInformation(ctx context.Context, params *PeopleSetLeaveInformationParams) (*Person, error) {
	if params == nil {
		params = &PeopleSetLeaveInformationParams{}
	}
	path := "/people/:personId/set-leave"
	if params.PersonID == "" {
		return nil, fmt.Errorf("personId is required")
	}
	path = strings.ReplaceAll(path, ":personId", url.PathEscape(params.PersonID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &Person{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type PeopleUpdatePersonMetadataRequestBody struct {
	Employment map[string]any `json:"employment"`
	Name       map[string]any `json:"name"`
}

type PeopleUpdatePersonMetadataParams struct {
	PersonID string
	Body     *PeopleUpdatePersonMetadataRequestBody
}

// UpdatePersonMetadata Update a person's basic information.
func (s *PeopleService) UpdatePersonMetadata(ctx context.Context, params *PeopleUpdatePersonMetadataParams) (*Person, error) {
	if params == nil {
		params = &PeopleUpdatePersonMetadataParams{}
	}
	path := "/people/:personId"
	if params.PersonID == "" {
		return nil, fmt.Errorf("personId is required")
	}
	path = strings.ReplaceAll(path, ":personId", url.PathEscape(params.PersonID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &Person{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// PoliciesService groups 2 endpoint methods under the "Policies" API segment.
type PoliciesService struct {
	client *Client
}

type PoliciesGetPolicyByIDResponse struct {
	ApprovedAtDate string         `json:"approvedAtDate"`
	Description    string         `json:"description"`
	ID             string         `json:"id"`
	LatestVersion  map[string]any `json:"latestVersion"`
	Name           string         `json:"name"`
	Status         string         `json:"status"`
}

type PoliciesGetPolicyByIDParams struct {
	PolicyID string
}

// GetPolicyByID Gets a policy by ID. Policy IDs can be found in Vanta in URL bar after /policies/.
func (s *PoliciesService) GetPolicyByID(ctx context.Context, params *PoliciesGetPolicyByIDParams) (*PoliciesGetPolicyByIDResponse, error) {
	if params == nil {
		params = &PoliciesGetPolicyByIDParams{}
	}
	path := "/policies/:policyId"
	if params.PolicyID == "" {
		return nil, fmt.Errorf("policyId is required")
	}
	path = strings.ReplaceAll(path, ":policyId", url.PathEscape(params.PolicyID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &PoliciesGetPolicyByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type PoliciesListPoliciesResponse struct {
	Results map[string]any `json:"results"`
}

type PoliciesListPoliciesParams struct {
	PageSize   *int
	PageCursor *string
}

// ListPolicies Lists all policies.
func (s *PoliciesService) ListPolicies(ctx context.Context, params *PoliciesListPoliciesParams) (*PoliciesListPoliciesResponse, error) {
	if params == nil {
		params = &PoliciesListPoliciesParams{}
	}
	path := "/policies"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &PoliciesListPoliciesResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// ResourcesService groups 6 endpoint methods under the "Resources" API segment.
type ResourcesService struct {
	client *Client
}

type ResourcesGetComputersParams struct {
	ResourceID *string
}

// GetComputers GetComputers performs GET /v1/resources/macos_user_computer.
func (s *ResourcesService) GetComputers(ctx context.Context, params *ResourcesGetComputersParams) (json.RawMessage, error) {
	if params == nil {
		params = &ResourcesGetComputersParams{}
	}
	path := "/v1/resources/macos_user_computer"
	query := url.Values{}
	if params.ResourceID != nil {
		query.Set("resourceId", fmt.Sprint(*params.ResourceID))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type ResourcesGetCustomResourceServerParams struct {
	ResourceID *string
}

// GetCustomResourceServer GetCustomResourceServer performs GET /v1/resources/custom_resource.
func (s *ResourcesService) GetCustomResourceServer(ctx context.Context, params *ResourcesGetCustomResourceServerParams) (json.RawMessage, error) {
	if params == nil {
		params = &ResourcesGetCustomResourceServerParams{}
	}
	path := "/v1/resources/custom_resource"
	query := url.Values{}
	if params.ResourceID != nil {
		query.Set("resourceId", fmt.Sprint(*params.ResourceID))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type ResourcesGetUserAccountsParams struct {
	ResourceID *string
}

// GetUserAccounts GetUserAccounts performs GET /v1/resources/user_account.
func (s *ResourcesService) GetUserAccounts(ctx context.Context, params *ResourcesGetUserAccountsParams) (json.RawMessage, error) {
	if params == nil {
		params = &ResourcesGetUserAccountsParams{}
	}
	path := "/v1/resources/user_account"
	query := url.Values{}
	if params.ResourceID != nil {
		query.Set("resourceId", fmt.Sprint(*params.ResourceID))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type ResourcesSyncCustomResourceServerRequestBody struct {
	ResourceID string           `json:"resourceId"`
	Resources  []map[string]any `json:"resources"`
}

type ResourcesSyncCustomResourceServerParams struct {
	Body *ResourcesSyncCustomResourceServerRequestBody
}

// SyncCustomResourceServer SyncCustomResourceServer performs PUT /v1/resources/custom_resource.
func (s *ResourcesService) SyncCustomResourceServer(ctx context.Context, params *ResourcesSyncCustomResourceServerParams) (json.RawMessage, error) {
	if params == nil {
		params = &ResourcesSyncCustomResourceServerParams{}
	}
	path := "/v1/resources/custom_resource"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PUT", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type ResourcesSyncMacOsComputersRequestBody struct {
	ResourceID string           `json:"resourceId"`
	Resources  []map[string]any `json:"resources"`
}

type ResourcesSyncMacOsComputersParams struct {
	Body *ResourcesSyncMacOsComputersRequestBody
}

// SyncMacOsComputers SyncMacOsComputers performs PUT /v1/resources/macos_user_computer.
func (s *ResourcesService) SyncMacOsComputers(ctx context.Context, params *ResourcesSyncMacOsComputersParams) (json.RawMessage, error) {
	if params == nil {
		params = &ResourcesSyncMacOsComputersParams{}
	}
	path := "/v1/resources/macos_user_computer"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PUT", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type ResourcesSyncUserAccountsRequestBody struct {
	ResourceID string           `json:"resourceId"`
	Resources  []map[string]any `json:"resources"`
}

type ResourcesSyncUserAccountsParams struct {
	Body *ResourcesSyncUserAccountsRequestBody
}

// SyncUserAccounts SyncUserAccounts performs PUT /v1/resources/user_account.
func (s *ResourcesService) SyncUserAccounts(ctx context.Context, params *ResourcesSyncUserAccountsParams) (json.RawMessage, error) {
	if params == nil {
		params = &ResourcesSyncUserAccountsParams{}
	}
	path := "/v1/resources/user_account"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PUT", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// RiskScenariosService groups 6 endpoint methods under the "RiskScenarios" API segment.
type RiskScenariosService struct {
	client *Client
}

type RiskScenariosCancelRiskScenarioApprovalRequestResponse struct {
	Categories         []string `json:"categories"`
	CiaCategories      []string `json:"ciaCategories"`
	CustomFields       []any    `json:"customFields"`
	Description        string   `json:"description"`
	Impact             float64  `json:"impact"`
	IsArchived         bool     `json:"isArchived"`
	IsSensitive        bool     `json:"isSensitive"`
	Likelihood         float64  `json:"likelihood"`
	Note               any      `json:"note"`
	Owner              any      `json:"owner"`
	RequiredApprovers  []any    `json:"requiredApprovers"`
	ResidualImpact     float64  `json:"residualImpact"`
	ResidualLikelihood float64  `json:"residualLikelihood"`
	ReviewStatus       string   `json:"reviewStatus"`
	RiskID             string   `json:"riskId"`
	RiskRegister       string   `json:"riskRegister"`
	Treatment          string   `json:"treatment"`
	Type               string   `json:"type"`
}

type RiskScenariosCancelRiskScenarioApprovalRequestParams struct {
	RiskScenarioID string
}

// CancelRiskScenarioApprovalRequest Cancel approval request for a risk scenario.
func (s *RiskScenariosService) CancelRiskScenarioApprovalRequest(ctx context.Context, params *RiskScenariosCancelRiskScenarioApprovalRequestParams) (*RiskScenariosCancelRiskScenarioApprovalRequestResponse, error) {
	if params == nil {
		params = &RiskScenariosCancelRiskScenarioApprovalRequestParams{}
	}
	path := "/risk-scenarios/:riskScenarioId/cancel-approval-request"
	if params.RiskScenarioID == "" {
		return nil, fmt.Errorf("riskScenarioId is required")
	}
	path = strings.ReplaceAll(path, ":riskScenarioId", url.PathEscape(params.RiskScenarioID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &RiskScenariosCancelRiskScenarioApprovalRequestResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type RiskScenariosCreateRiskScenarioRequestBody struct {
	Categories         []string         `json:"categories"`
	CiaCategories      []string         `json:"ciaCategories"`
	CustomFields       []map[string]any `json:"customFields"`
	Description        string           `json:"description"`
	Impact             float64          `json:"impact"`
	IsSensitive        bool             `json:"isSensitive"`
	Likelihood         float64          `json:"likelihood"`
	Note               string           `json:"note"`
	Owner              string           `json:"owner"`
	ResidualImpact     float64          `json:"residualImpact"`
	ResidualLikelihood float64          `json:"residualLikelihood"`
	RiskID             string           `json:"riskId"`
	RiskRegister       string           `json:"riskRegister"`
	Treatment          string           `json:"treatment"`
	Type               string           `json:"type"`
}

type RiskScenariosCreateRiskScenarioResponse struct {
	Categories         []string `json:"categories"`
	CiaCategories      []string `json:"ciaCategories"`
	CustomFields       []any    `json:"customFields"`
	Description        string   `json:"description"`
	Impact             float64  `json:"impact"`
	IsArchived         bool     `json:"isArchived"`
	IsSensitive        bool     `json:"isSensitive"`
	Likelihood         float64  `json:"likelihood"`
	Note               any      `json:"note"`
	Owner              any      `json:"owner"`
	RequiredApprovers  []any    `json:"requiredApprovers"`
	ResidualImpact     float64  `json:"residualImpact"`
	ResidualLikelihood float64  `json:"residualLikelihood"`
	ReviewStatus       string   `json:"reviewStatus"`
	RiskID             string   `json:"riskId"`
	RiskRegister       string   `json:"riskRegister"`
	Treatment          string   `json:"treatment"`
	Type               string   `json:"type"`
}

type RiskScenariosCreateRiskScenarioParams struct {
	Body *RiskScenariosCreateRiskScenarioRequestBody
}

// CreateRiskScenario Create a new risk scenario.
func (s *RiskScenariosService) CreateRiskScenario(ctx context.Context, params *RiskScenariosCreateRiskScenarioParams) (*RiskScenariosCreateRiskScenarioResponse, error) {
	if params == nil {
		params = &RiskScenariosCreateRiskScenarioParams{}
	}
	path := "/risk-scenarios"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &RiskScenariosCreateRiskScenarioResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type RiskScenariosGetRiskScenarioByIDResponse struct {
	Categories         []string `json:"categories"`
	CiaCategories      []string `json:"ciaCategories"`
	CustomFields       []any    `json:"customFields"`
	Description        string   `json:"description"`
	Impact             float64  `json:"impact"`
	IsArchived         bool     `json:"isArchived"`
	IsSensitive        bool     `json:"isSensitive"`
	Likelihood         float64  `json:"likelihood"`
	Note               any      `json:"note"`
	Owner              any      `json:"owner"`
	RequiredApprovers  []any    `json:"requiredApprovers"`
	ResidualImpact     float64  `json:"residualImpact"`
	ResidualLikelihood float64  `json:"residualLikelihood"`
	ReviewStatus       string   `json:"reviewStatus"`
	RiskID             string   `json:"riskId"`
	RiskRegister       string   `json:"riskRegister"`
	Treatment          string   `json:"treatment"`
	Type               string   `json:"type"`
}

type RiskScenariosGetRiskScenarioByIDParams struct {
	RiskScenarioID string
}

// GetRiskScenarioByID Get a risk scenario by ID (can be the Risk ID or the object ID).
func (s *RiskScenariosService) GetRiskScenarioByID(ctx context.Context, params *RiskScenariosGetRiskScenarioByIDParams) (*RiskScenariosGetRiskScenarioByIDResponse, error) {
	if params == nil {
		params = &RiskScenariosGetRiskScenarioByIDParams{}
	}
	path := "/risk-scenarios/:riskScenarioId"
	if params.RiskScenarioID == "" {
		return nil, fmt.Errorf("riskScenarioId is required")
	}
	path = strings.ReplaceAll(path, ":riskScenarioId", url.PathEscape(params.RiskScenarioID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &RiskScenariosGetRiskScenarioByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type RiskScenariosListRiskScenariosResponse struct {
	Results map[string]any `json:"results"`
}

type RiskScenariosListRiskScenariosParams struct {
	PageSize                     *int
	PageCursor                   *string
	IncludeIgnored               *bool
	OwnerMatchesAny              []string
	SearchString                 *string
	CategoryMatchesAny           []string
	CiaCategoryMatchesAny        []string
	TreatmentTypeMatchesAny      []string
	InherentScoreGroupMatchesAny []string
	ResidualScoreGroupMatchesAny []string
	ReviewStatusMatchesAny       []string
	Type                         *string
	OrderBy                      *string
}

// ListRiskScenarios List risk scenarios.
func (s *RiskScenariosService) ListRiskScenarios(ctx context.Context, params *RiskScenariosListRiskScenariosParams) (*RiskScenariosListRiskScenariosResponse, error) {
	if params == nil {
		params = &RiskScenariosListRiskScenariosParams{}
	}
	path := "/risk-scenarios"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	if params.IncludeIgnored != nil {
		query.Set("includeIgnored", fmt.Sprint(*params.IncludeIgnored))
	}
	for _, v := range params.OwnerMatchesAny {
		query.Add("ownerMatchesAny", v)
	}
	if params.SearchString != nil {
		query.Set("searchString", fmt.Sprint(*params.SearchString))
	}
	for _, v := range params.CategoryMatchesAny {
		query.Add("categoryMatchesAny", v)
	}
	for _, v := range params.CiaCategoryMatchesAny {
		query.Add("ciaCategoryMatchesAny", v)
	}
	for _, v := range params.TreatmentTypeMatchesAny {
		query.Add("treatmentTypeMatchesAny", v)
	}
	for _, v := range params.InherentScoreGroupMatchesAny {
		query.Add("inherentScoreGroupMatchesAny", v)
	}
	for _, v := range params.ResidualScoreGroupMatchesAny {
		query.Add("residualScoreGroupMatchesAny", v)
	}
	for _, v := range params.ReviewStatusMatchesAny {
		query.Add("reviewStatusMatchesAny", v)
	}
	if params.Type != nil {
		query.Set("type", fmt.Sprint(*params.Type))
	}
	if params.OrderBy != nil {
		query.Set("orderBy", fmt.Sprint(*params.OrderBy))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &RiskScenariosListRiskScenariosResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type RiskScenariosSubmitRiskScenarioForApprovalRequestBody struct {
	Comment string `json:"comment"`
}

type RiskScenariosSubmitRiskScenarioForApprovalResponse struct {
	Categories         []string `json:"categories"`
	CiaCategories      []string `json:"ciaCategories"`
	CustomFields       []any    `json:"customFields"`
	Description        string   `json:"description"`
	Impact             float64  `json:"impact"`
	IsArchived         bool     `json:"isArchived"`
	IsSensitive        bool     `json:"isSensitive"`
	Likelihood         float64  `json:"likelihood"`
	Note               any      `json:"note"`
	Owner              any      `json:"owner"`
	RequiredApprovers  []any    `json:"requiredApprovers"`
	ResidualImpact     float64  `json:"residualImpact"`
	ResidualLikelihood float64  `json:"residualLikelihood"`
	ReviewStatus       string   `json:"reviewStatus"`
	RiskID             string   `json:"riskId"`
	RiskRegister       string   `json:"riskRegister"`
	Treatment          string   `json:"treatment"`
	Type               string   `json:"type"`
}

type RiskScenariosSubmitRiskScenarioForApprovalParams struct {
	RiskScenarioID string
	Body           *RiskScenariosSubmitRiskScenarioForApprovalRequestBody
}

// SubmitRiskScenarioForApproval Submit a risk scenario for approval.
func (s *RiskScenariosService) SubmitRiskScenarioForApproval(ctx context.Context, params *RiskScenariosSubmitRiskScenarioForApprovalParams) (*RiskScenariosSubmitRiskScenarioForApprovalResponse, error) {
	if params == nil {
		params = &RiskScenariosSubmitRiskScenarioForApprovalParams{}
	}
	path := "/risk-scenarios/:riskScenarioId/submit-for-approval"
	if params.RiskScenarioID == "" {
		return nil, fmt.Errorf("riskScenarioId is required")
	}
	path = strings.ReplaceAll(path, ":riskScenarioId", url.PathEscape(params.RiskScenarioID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &RiskScenariosSubmitRiskScenarioForApprovalResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type RiskScenariosUpdateRiskScenarioRequestBody struct {
	Categories         []string         `json:"categories"`
	CiaCategories      []string         `json:"ciaCategories"`
	CustomFields       []map[string]any `json:"customFields"`
	Description        string           `json:"description"`
	Impact             float64          `json:"impact"`
	IsSensitive        bool             `json:"isSensitive"`
	Likelihood         float64          `json:"likelihood"`
	Note               string           `json:"note"`
	Owner              string           `json:"owner"`
	ResidualImpact     float64          `json:"residualImpact"`
	ResidualLikelihood float64          `json:"residualLikelihood"`
	RiskRegister       string           `json:"riskRegister"`
	Treatment          string           `json:"treatment"`
}

type RiskScenariosUpdateRiskScenarioResponse struct {
	Categories         []string `json:"categories"`
	CiaCategories      []string `json:"ciaCategories"`
	CustomFields       []any    `json:"customFields"`
	Description        string   `json:"description"`
	Impact             float64  `json:"impact"`
	IsArchived         bool     `json:"isArchived"`
	IsSensitive        bool     `json:"isSensitive"`
	Likelihood         float64  `json:"likelihood"`
	Note               any      `json:"note"`
	Owner              any      `json:"owner"`
	RequiredApprovers  []any    `json:"requiredApprovers"`
	ResidualImpact     float64  `json:"residualImpact"`
	ResidualLikelihood float64  `json:"residualLikelihood"`
	ReviewStatus       string   `json:"reviewStatus"`
	RiskID             string   `json:"riskId"`
	RiskRegister       string   `json:"riskRegister"`
	Treatment          string   `json:"treatment"`
	Type               string   `json:"type"`
}

type RiskScenariosUpdateRiskScenarioParams struct {
	RiskScenarioID string
	Body           *RiskScenariosUpdateRiskScenarioRequestBody
}

// UpdateRiskScenario Update a risk scenario.
func (s *RiskScenariosService) UpdateRiskScenario(ctx context.Context, params *RiskScenariosUpdateRiskScenarioParams) (*RiskScenariosUpdateRiskScenarioResponse, error) {
	if params == nil {
		params = &RiskScenariosUpdateRiskScenarioParams{}
	}
	path := "/risk-scenarios/:riskScenarioId"
	if params.RiskScenarioID == "" {
		return nil, fmt.Errorf("riskScenarioId is required")
	}
	path = strings.ReplaceAll(path, ":riskScenarioId", url.PathEscape(params.RiskScenarioID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &RiskScenariosUpdateRiskScenarioResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// TestsService groups 5 endpoint methods under the "Tests" API segment.
type TestsService struct {
	client *Client
}

type TestsDeactivateTestEntityRequestBody struct {
	DeactivateReason    string `json:"deactivateReason"`
	DeactivateUntilDate string `json:"deactivateUntilDate"`
}

type TestsDeactivateTestEntityParams struct {
	TestID   string
	EntityID string
	Body     *TestsDeactivateTestEntityRequestBody
}

// DeactivateTestEntity Deactivates a single test item (test entity). There may be a delay in the deactivation of the test entity until the next test run. Use the /vulnerabilities/deactivate endpoint for vulnerabilities.
func (s *TestsService) DeactivateTestEntity(ctx context.Context, params *TestsDeactivateTestEntityParams) (json.RawMessage, error) {
	if params == nil {
		params = &TestsDeactivateTestEntityParams{}
	}
	path := "/tests/:testId/entities/:entityId/deactivate"
	if params.TestID == "" {
		return nil, fmt.Errorf("testId is required")
	}
	path = strings.ReplaceAll(path, ":testId", url.PathEscape(params.TestID))
	if params.EntityID == "" {
		return nil, fmt.Errorf("entityId is required")
	}
	path = strings.ReplaceAll(path, ":entityId", url.PathEscape(params.EntityID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TestsGetTestByIDResponse struct {
	Category               string         `json:"category"`
	DeactivatedStatusInfo  map[string]any `json:"deactivatedStatusInfo"`
	Description            string         `json:"description"`
	FailureDescription     string         `json:"failureDescription"`
	ID                     string         `json:"id"`
	Integrations           []string       `json:"integrations"`
	LastTestRunDate        string         `json:"lastTestRunDate"`
	LatestFlipDate         any            `json:"latestFlipDate"`
	Name                   string         `json:"name"`
	Owner                  any            `json:"owner"`
	RemediationDescription string         `json:"remediationDescription"`
	RemediationStatusInfo  map[string]any `json:"remediationStatusInfo"`
	Status                 string         `json:"status"`
	Version                map[string]any `json:"version"`
}

type TestsGetTestByIDParams struct {
	TestID string
}

// GetTestByID Gets a test by ID. Test IDs can be found in Vanta in URL bar after /tests/.
func (s *TestsService) GetTestByID(ctx context.Context, params *TestsGetTestByIDParams) (*TestsGetTestByIDResponse, error) {
	if params == nil {
		params = &TestsGetTestByIDParams{}
	}
	path := "/tests/:testId"
	if params.TestID == "" {
		return nil, fmt.Errorf("testId is required")
	}
	path = strings.ReplaceAll(path, ":testId", url.PathEscape(params.TestID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TestsGetTestByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TestsGetTestEntitiesByTestIDResponse struct {
	Results map[string]any `json:"results"`
}

type TestsGetTestEntitiesByTestIDParams struct {
	TestID       string
	EntityStatus *string
	PageSize     *int
	PageCursor   *string
}

// GetTestEntitiesByTestID Gets a list of tested items (entities) for a test by test ID. An entity is a tested item that can have its own outcome. For example, for a test that makes sure that all S3 buckets are versioned, an individual S3 bucket would be an entity.
func (s *TestsService) GetTestEntitiesByTestID(ctx context.Context, params *TestsGetTestEntitiesByTestIDParams) (*TestsGetTestEntitiesByTestIDResponse, error) {
	if params == nil {
		params = &TestsGetTestEntitiesByTestIDParams{}
	}
	path := "/tests/:testId/entities"
	if params.TestID == "" {
		return nil, fmt.Errorf("testId is required")
	}
	path = strings.ReplaceAll(path, ":testId", url.PathEscape(params.TestID))
	query := url.Values{}
	if params.EntityStatus != nil {
		query.Set("entityStatus", fmt.Sprint(*params.EntityStatus))
	}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TestsGetTestEntitiesByTestIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TestsListTestsResponse struct {
	Results map[string]any `json:"results"`
}

type TestsListTestsParams struct {
	PageSize          *int
	PageCursor        *string
	StatusFilter      *string
	FrameworkFilter   *string
	IntegrationFilter *string
	ControlFilter     *string
	OwnerFilter       *string
	CategoryFilter    *string
	IsInRollout       *bool
}

// ListTests Lists all tests based on applied filters.
func (s *TestsService) ListTests(ctx context.Context, params *TestsListTestsParams) (*TestsListTestsResponse, error) {
	if params == nil {
		params = &TestsListTestsParams{}
	}
	path := "/tests"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	if params.StatusFilter != nil {
		query.Set("statusFilter", fmt.Sprint(*params.StatusFilter))
	}
	if params.FrameworkFilter != nil {
		query.Set("frameworkFilter", fmt.Sprint(*params.FrameworkFilter))
	}
	if params.IntegrationFilter != nil {
		query.Set("integrationFilter", fmt.Sprint(*params.IntegrationFilter))
	}
	if params.ControlFilter != nil {
		query.Set("controlFilter", fmt.Sprint(*params.ControlFilter))
	}
	if params.OwnerFilter != nil {
		query.Set("ownerFilter", fmt.Sprint(*params.OwnerFilter))
	}
	if params.CategoryFilter != nil {
		query.Set("categoryFilter", fmt.Sprint(*params.CategoryFilter))
	}
	if params.IsInRollout != nil {
		query.Set("isInRollout", fmt.Sprint(*params.IsInRollout))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TestsListTestsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TestsReactivateTestEntityParams struct {
	TestID   string
	EntityID string
}

// ReactivateTestEntity Reactivates a single tested item (test entity). There may be a delay in the reactivation of the test entity until the next test run. Use the /vulnerabilities/reactivate endpoint for vulnerabilities.
func (s *TestsService) ReactivateTestEntity(ctx context.Context, params *TestsReactivateTestEntityParams) (json.RawMessage, error) {
	if params == nil {
		params = &TestsReactivateTestEntityParams{}
	}
	path := "/tests/:testId/entities/:entityId/reactivate"
	if params.TestID == "" {
		return nil, fmt.Errorf("testId is required")
	}
	path = strings.ReplaceAll(path, ":testId", url.PathEscape(params.TestID))
	if params.EntityID == "" {
		return nil, fmt.Errorf("entityId is required")
	}
	path = strings.ReplaceAll(path, ":entityId", url.PathEscape(params.EntityID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TrustCentersService groups 54 endpoint methods under the "TrustCenters" API segment.
type TrustCentersService struct {
	client *Client
}

type TrustCentersAddTrustCenterControlRequestBody struct {
	CategoryIDs []string `json:"categoryIds"`
	ControlID   string   `json:"controlId"`
}

type TrustCentersAddTrustCenterControlResponse struct {
	Categories  []map[string]any `json:"categories"`
	Description string           `json:"description"`
	ID          string           `json:"id"`
	Name        string           `json:"name"`
}

type TrustCentersAddTrustCenterControlParams struct {
	SlugID string
	Body   *TrustCentersAddTrustCenterControlRequestBody
}

// AddTrustCenterControl Adds a control to a Trust Center.
func (s *TrustCentersService) AddTrustCenterControl(ctx context.Context, params *TrustCentersAddTrustCenterControlParams) (*TrustCentersAddTrustCenterControlResponse, error) {
	if params == nil {
		params = &TrustCentersAddTrustCenterControlParams{}
	}
	path := "/trust-centers/:slugId/controls"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersAddTrustCenterControlResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersAddTrustCenterControlCategoryRequestBody struct {
	Name string `json:"name"`
}

type TrustCentersAddTrustCenterControlCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TrustCentersAddTrustCenterControlCategoryParams struct {
	SlugID string
	Body   *TrustCentersAddTrustCenterControlCategoryRequestBody
}

// AddTrustCenterControlCategory Adds a control category to a Trust Center.
func (s *TrustCentersService) AddTrustCenterControlCategory(ctx context.Context, params *TrustCentersAddTrustCenterControlCategoryParams) (*TrustCentersAddTrustCenterControlCategoryResponse, error) {
	if params == nil {
		params = &TrustCentersAddTrustCenterControlCategoryParams{}
	}
	path := "/trust-centers/:slugId/control-categories"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersAddTrustCenterControlCategoryResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersAddTrustCenterViewerRequestBody struct {
	AccessLevel    string   `json:"accessLevel"`
	CompanyName    string   `json:"companyName"`
	Email          string   `json:"email"`
	ExpirationDate string   `json:"expirationDate"`
	IsNdaRequired  bool     `json:"isNdaRequired"`
	Name           string   `json:"name"`
	ResourceIDs    []string `json:"resourceIds"`
}

type TrustCentersAddTrustCenterViewerResponse struct {
	AccessLevel                 string           `json:"accessLevel"`
	AddedByUser                 any              `json:"addedByUser"`
	CompanyName                 string           `json:"companyName"`
	CreationDate                string           `json:"creationDate"`
	Email                       string           `json:"email"`
	ExpirationDate              any              `json:"expirationDate"`
	ExternalServiceAssociations []map[string]any `json:"externalServiceAssociations"`
	ID                          string           `json:"id"`
	Name                        string           `json:"name"`
	NdaInfo                     any              `json:"ndaInfo"`
	ResourceIDs                 any              `json:"resourceIds"`
	UpdatedDate                 string           `json:"updatedDate"`
}

type TrustCentersAddTrustCenterViewerParams struct {
	SlugID string
	Body   *TrustCentersAddTrustCenterViewerRequestBody
}

// AddTrustCenterViewer Adds a viewer and grants them access to a Trust Center.
func (s *TrustCentersService) AddTrustCenterViewer(ctx context.Context, params *TrustCentersAddTrustCenterViewerParams) (*TrustCentersAddTrustCenterViewerResponse, error) {
	if params == nil {
		params = &TrustCentersAddTrustCenterViewerParams{}
	}
	path := "/trust-centers/:slugId/viewers"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersAddTrustCenterViewerResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersApproveTrustCenterAccessRequestRequestBody struct {
	AccessLevel    string   `json:"accessLevel"`
	ExpirationDate string   `json:"expirationDate"`
	IsNdaRequired  bool     `json:"isNdaRequired"`
	ResourceIDs    []string `json:"resourceIds"`
}

type TrustCentersApproveTrustCenterAccessRequestParams struct {
	SlugID          string
	AccessRequestID string
	Body            *TrustCentersApproveTrustCenterAccessRequestRequestBody
}

// ApproveTrustCenterAccessRequest Approves an access request on a Trust Center.
func (s *TrustCentersService) ApproveTrustCenterAccessRequest(ctx context.Context, params *TrustCentersApproveTrustCenterAccessRequestParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersApproveTrustCenterAccessRequestParams{}
	}
	path := "/trust-centers/:slugId/access-requests/:accessRequestId/approve"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.AccessRequestID == "" {
		return nil, fmt.Errorf("accessRequestId is required")
	}
	path = strings.ReplaceAll(path, ":accessRequestId", url.PathEscape(params.AccessRequestID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersCreateTrustCenterDocumentResponse struct {
	CreationDate string `json:"creationDate"`
	Description  string `json:"description"`
	FileName     string `json:"fileName"`
	ID           string `json:"id"`
	IsPublic     bool   `json:"isPublic"`
	MimeType     string `json:"mimeType"`
	Title        string `json:"title"`
	UpdatedDate  string `json:"updatedDate"`
}

type TrustCentersCreateTrustCenterDocumentParams struct {
	SlugID string
	// FormData maps multipart field names to values.
	FormData map[string]string
}

// CreateTrustCenterDocument Adds a document to a Trust Center.
func (s *TrustCentersService) CreateTrustCenterDocument(ctx context.Context, params *TrustCentersCreateTrustCenterDocumentParams) (*TrustCentersCreateTrustCenterDocumentResponse, error) {
	if params == nil {
		params = &TrustCentersCreateTrustCenterDocumentParams{}
	}
	path := "/trust-centers/:slugId/resources"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newMultipartRequest(ctx, "POST", path, query, params.FormData)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersCreateTrustCenterDocumentResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersCreateTrustCenterFaqRequestBody struct {
	Answer   string `json:"answer"`
	Question string `json:"question"`
}

type TrustCentersCreateTrustCenterFaqResponse struct {
	Answer   string `json:"answer"`
	ID       string `json:"id"`
	Question string `json:"question"`
}

type TrustCentersCreateTrustCenterFaqParams struct {
	SlugID string
	Body   *TrustCentersCreateTrustCenterFaqRequestBody
}

// CreateTrustCenterFaq Adds an FAQ to a Trust Center.
func (s *TrustCentersService) CreateTrustCenterFaq(ctx context.Context, params *TrustCentersCreateTrustCenterFaqParams) (*TrustCentersCreateTrustCenterFaqResponse, error) {
	if params == nil {
		params = &TrustCentersCreateTrustCenterFaqParams{}
	}
	path := "/trust-centers/:slugId/faqs"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersCreateTrustCenterFaqResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersCreateTrustCenterSubprocessorRequestBody struct {
	Description string `json:"description"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Purpose     string `json:"purpose"`
	URL         string `json:"url"`
}

type TrustCentersCreateTrustCenterSubprocessorResponse struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Purpose     string `json:"purpose"`
	URL         string `json:"url"`
}

type TrustCentersCreateTrustCenterSubprocessorParams struct {
	SlugID string
	Body   *TrustCentersCreateTrustCenterSubprocessorRequestBody
}

// CreateTrustCenterSubprocessor Adds a subprocessor to a Trust Center.
func (s *TrustCentersService) CreateTrustCenterSubprocessor(ctx context.Context, params *TrustCentersCreateTrustCenterSubprocessorParams) (*TrustCentersCreateTrustCenterSubprocessorResponse, error) {
	if params == nil {
		params = &TrustCentersCreateTrustCenterSubprocessorParams{}
	}
	path := "/trust-centers/:slugId/subprocessors"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersCreateTrustCenterSubprocessorResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersCreateTrustCenterSubscriberRequestBody struct {
	Email string `json:"email"`
}

type TrustCentersCreateTrustCenterSubscriberResponse struct {
	CreationDate    string `json:"creationDate"`
	Email           string `json:"email"`
	ID              string `json:"id"`
	IsEmailVerified bool   `json:"isEmailVerified"`
}

type TrustCentersCreateTrustCenterSubscriberParams struct {
	SlugID string
	Body   *TrustCentersCreateTrustCenterSubscriberRequestBody
}

// CreateTrustCenterSubscriber Adds a subscriber to a Trust Center.
func (s *TrustCentersService) CreateTrustCenterSubscriber(ctx context.Context, params *TrustCentersCreateTrustCenterSubscriberParams) (*TrustCentersCreateTrustCenterSubscriberResponse, error) {
	if params == nil {
		params = &TrustCentersCreateTrustCenterSubscriberParams{}
	}
	path := "/trust-centers/:slugId/subscribers"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersCreateTrustCenterSubscriberResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersCreateTrustCenterSubscriberGroupRequestBody struct {
	Name          string   `json:"name"`
	SubscriberIDs []string `json:"subscriberIds"`
}

type TrustCentersCreateTrustCenterSubscriberGroupResponse struct {
	CreationDate  string   `json:"creationDate"`
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	SubscriberIDs []string `json:"subscriberIds"`
}

type TrustCentersCreateTrustCenterSubscriberGroupParams struct {
	SlugID string
	Body   *TrustCentersCreateTrustCenterSubscriberGroupRequestBody
}

// CreateTrustCenterSubscriberGroup Adds a subscriber group to a Trust Center.
func (s *TrustCentersService) CreateTrustCenterSubscriberGroup(ctx context.Context, params *TrustCentersCreateTrustCenterSubscriberGroupParams) (*TrustCentersCreateTrustCenterSubscriberGroupResponse, error) {
	if params == nil {
		params = &TrustCentersCreateTrustCenterSubscriberGroupParams{}
	}
	path := "/trust-centers/:slugId/subscriber-groups"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersCreateTrustCenterSubscriberGroupResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersCreateTrustCenterUpdateRequestBody struct {
	Category           string   `json:"category"`
	Description        string   `json:"description"`
	NotificationTarget string   `json:"notificationTarget"`
	NotifiedEmails     []string `json:"notifiedEmails"`
	SubscriberGroupIDs []string `json:"subscriberGroupIds"`
	Title              string   `json:"title"`
	VisibilityType     string   `json:"visibilityType"`
}

type TrustCentersCreateTrustCenterUpdateResponse struct {
	Category       string   `json:"category"`
	CreationDate   string   `json:"creationDate"`
	Description    string   `json:"description"`
	ID             string   `json:"id"`
	NotifiedEmails []string `json:"notifiedEmails"`
	Title          string   `json:"title"`
	UpdatedDate    string   `json:"updatedDate"`
	VisibilityType string   `json:"visibilityType"`
}

type TrustCentersCreateTrustCenterUpdateParams struct {
	SlugID string
	Body   *TrustCentersCreateTrustCenterUpdateRequestBody
}

// CreateTrustCenterUpdate Adds an update to a Trust Center.
func (s *TrustCentersService) CreateTrustCenterUpdate(ctx context.Context, params *TrustCentersCreateTrustCenterUpdateParams) (*TrustCentersCreateTrustCenterUpdateResponse, error) {
	if params == nil {
		params = &TrustCentersCreateTrustCenterUpdateParams{}
	}
	path := "/trust-centers/:slugId/updates"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersCreateTrustCenterUpdateResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersDeleteTrustCenterControlParams struct {
	SlugID    string
	ControlID string
}

// DeleteTrustCenterControl Removes a specific control from a Trust Center. This removes the control from all of the control categories that is in.
func (s *TrustCentersService) DeleteTrustCenterControl(ctx context.Context, params *TrustCentersDeleteTrustCenterControlParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersDeleteTrustCenterControlParams{}
	}
	path := "/trust-centers/:slugId/controls/:controlId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersDeleteTrustCenterControlCategoryParams struct {
	SlugID     string
	CategoryID string
}

// DeleteTrustCenterControlCategory Removes a control category from a Trust Center along with all of the controls in the category.
func (s *TrustCentersService) DeleteTrustCenterControlCategory(ctx context.Context, params *TrustCentersDeleteTrustCenterControlCategoryParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersDeleteTrustCenterControlCategoryParams{}
	}
	path := "/trust-centers/:slugId/control-categories/:categoryId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.CategoryID == "" {
		return nil, fmt.Errorf("categoryId is required")
	}
	path = strings.ReplaceAll(path, ":categoryId", url.PathEscape(params.CategoryID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersDeleteTrustCenterDocumentParams struct {
	SlugID     string
	ResourceID string
}

// DeleteTrustCenterDocument Removes a specific document from a Trust Center.
func (s *TrustCentersService) DeleteTrustCenterDocument(ctx context.Context, params *TrustCentersDeleteTrustCenterDocumentParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersDeleteTrustCenterDocumentParams{}
	}
	path := "/trust-centers/:slugId/resources/:resourceId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.ResourceID == "" {
		return nil, fmt.Errorf("resourceId is required")
	}
	path = strings.ReplaceAll(path, ":resourceId", url.PathEscape(params.ResourceID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersDeleteTrustCenterFaqParams struct {
	SlugID string
	FaqID  string
}

// DeleteTrustCenterFaq Remove a specific FAQ from the Trust Center by ID.
func (s *TrustCentersService) DeleteTrustCenterFaq(ctx context.Context, params *TrustCentersDeleteTrustCenterFaqParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersDeleteTrustCenterFaqParams{}
	}
	path := "/trust-centers/:slugId/faqs/:faqId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.FaqID == "" {
		return nil, fmt.Errorf("faqId is required")
	}
	path = strings.ReplaceAll(path, ":faqId", url.PathEscape(params.FaqID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersDeleteTrustCenterSubprocessorParams struct {
	SlugID         string
	SubprocessorID string
}

// DeleteTrustCenterSubprocessor Removes a subprocessor from a Trust Center.
func (s *TrustCentersService) DeleteTrustCenterSubprocessor(ctx context.Context, params *TrustCentersDeleteTrustCenterSubprocessorParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersDeleteTrustCenterSubprocessorParams{}
	}
	path := "/trust-centers/:slugId/subprocessors/:subprocessorId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.SubprocessorID == "" {
		return nil, fmt.Errorf("subprocessorId is required")
	}
	path = strings.ReplaceAll(path, ":subprocessorId", url.PathEscape(params.SubprocessorID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersDeleteTrustCenterSubscriberParams struct {
	SlugID       string
	SubscriberID string
}

// DeleteTrustCenterSubscriber Removes a subscriber from a Trust Center.
func (s *TrustCentersService) DeleteTrustCenterSubscriber(ctx context.Context, params *TrustCentersDeleteTrustCenterSubscriberParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersDeleteTrustCenterSubscriberParams{}
	}
	path := "/trust-centers/:slugId/subscribers/:subscriberId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.SubscriberID == "" {
		return nil, fmt.Errorf("subscriberId is required")
	}
	path = strings.ReplaceAll(path, ":subscriberId", url.PathEscape(params.SubscriberID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersDeleteTrustCenterSubscriberGroupParams struct {
	SlugID            string
	SubscriberGroupID string
}

// DeleteTrustCenterSubscriberGroup Removes a subscriber group from a Trust Center.
func (s *TrustCentersService) DeleteTrustCenterSubscriberGroup(ctx context.Context, params *TrustCentersDeleteTrustCenterSubscriberGroupParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersDeleteTrustCenterSubscriberGroupParams{}
	}
	path := "/trust-centers/:slugId/subscriber-groups/:subscriberGroupId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.SubscriberGroupID == "" {
		return nil, fmt.Errorf("subscriberGroupId is required")
	}
	path = strings.ReplaceAll(path, ":subscriberGroupId", url.PathEscape(params.SubscriberGroupID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersDeleteTrustCenterUpdateParams struct {
	SlugID   string
	UpdateID string
}

// DeleteTrustCenterUpdate Removes an update from a Trust Center.
func (s *TrustCentersService) DeleteTrustCenterUpdate(ctx context.Context, params *TrustCentersDeleteTrustCenterUpdateParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersDeleteTrustCenterUpdateParams{}
	}
	path := "/trust-centers/:slugId/updates/:updateId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.UpdateID == "" {
		return nil, fmt.Errorf("updateId is required")
	}
	path = strings.ReplaceAll(path, ":updateId", url.PathEscape(params.UpdateID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersDenyTrustCenterAccessRequestParams struct {
	SlugID          string
	AccessRequestID string
}

// DenyTrustCenterAccessRequest Denies an access request on a Trust Center.
func (s *TrustCentersService) DenyTrustCenterAccessRequest(ctx context.Context, params *TrustCentersDenyTrustCenterAccessRequestParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersDenyTrustCenterAccessRequestParams{}
	}
	path := "/trust-centers/:slugId/access-requests/:accessRequestId/deny"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.AccessRequestID == "" {
		return nil, fmt.Errorf("accessRequestId is required")
	}
	path = strings.ReplaceAll(path, ":accessRequestId", url.PathEscape(params.AccessRequestID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersEditTrustCenterSubscriberGroupRequestBody struct {
	Name string `json:"name"`
}

type TrustCentersEditTrustCenterSubscriberGroupResponse struct {
	CreationDate  string   `json:"creationDate"`
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	SubscriberIDs []string `json:"subscriberIds"`
}

type TrustCentersEditTrustCenterSubscriberGroupParams struct {
	SlugID            string
	SubscriberGroupID string
	Body              *TrustCentersEditTrustCenterSubscriberGroupRequestBody
}

// EditTrustCenterSubscriberGroup Edits a Trust Center subscriber group.
func (s *TrustCentersService) EditTrustCenterSubscriberGroup(ctx context.Context, params *TrustCentersEditTrustCenterSubscriberGroupParams) (*TrustCentersEditTrustCenterSubscriberGroupResponse, error) {
	if params == nil {
		params = &TrustCentersEditTrustCenterSubscriberGroupParams{}
	}
	path := "/trust-centers/:slugId/subscriber-groups/:subscriberGroupId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.SubscriberGroupID == "" {
		return nil, fmt.Errorf("subscriberGroupId is required")
	}
	path = strings.ReplaceAll(path, ":subscriberGroupId", url.PathEscape(params.SubscriberGroupID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersEditTrustCenterSubscriberGroupResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetTrustCenterResponse struct {
	BannerSetting      map[string]any `json:"bannerSetting"`
	CompanyDescription string         `json:"companyDescription"`
	CreationDate       string         `json:"creationDate"`
	CustomDomain       string         `json:"customDomain"`
	CustomTheme        map[string]any `json:"customTheme"`
	ID                 string         `json:"id"`
	IsPublic           bool           `json:"isPublic"`
	PrivacyPolicy      string         `json:"privacyPolicy"`
	Title              string         `json:"title"`
	UpdatedDate        string         `json:"updatedDate"`
}

type TrustCentersGetTrustCenterParams struct {
	SlugID string
}

// GetTrustCenter Gets a Trust Center by slug ID.
func (s *TrustCentersService) GetTrustCenter(ctx context.Context, params *TrustCentersGetTrustCenterParams) (*TrustCentersGetTrustCenterResponse, error) {
	if params == nil {
		params = &TrustCentersGetTrustCenterParams{}
	}
	path := "/trust-centers/:slugId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetTrustCenterResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetTrustCenterAccessRequestResponse struct {
	AccessLevel        string `json:"accessLevel"`
	CompanyName        string `json:"companyName"`
	CreationDate       string `json:"creationDate"`
	Email              string `json:"email"`
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Reason             string `json:"reason"`
	RequestedResources any    `json:"requestedResources"`
	UpdatedDate        string `json:"updatedDate"`
}

type TrustCentersGetTrustCenterAccessRequestParams struct {
	SlugID          string
	AccessRequestID string
}

// GetTrustCenterAccessRequest Gets a specific access request for a Trust Center.
func (s *TrustCentersService) GetTrustCenterAccessRequest(ctx context.Context, params *TrustCentersGetTrustCenterAccessRequestParams) (*TrustCentersGetTrustCenterAccessRequestResponse, error) {
	if params == nil {
		params = &TrustCentersGetTrustCenterAccessRequestParams{}
	}
	path := "/trust-centers/:slugId/access-requests/:accessRequestId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.AccessRequestID == "" {
		return nil, fmt.Errorf("accessRequestId is required")
	}
	path = strings.ReplaceAll(path, ":accessRequestId", url.PathEscape(params.AccessRequestID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetTrustCenterAccessRequestResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetTrustCenterControlResponse struct {
	Categories  []map[string]any `json:"categories"`
	Description string           `json:"description"`
	ID          string           `json:"id"`
	Name        string           `json:"name"`
}

type TrustCentersGetTrustCenterControlParams struct {
	SlugID    string
	ControlID string
}

// GetTrustCenterControl Gets a specific control on a Trust Center.
func (s *TrustCentersService) GetTrustCenterControl(ctx context.Context, params *TrustCentersGetTrustCenterControlParams) (*TrustCentersGetTrustCenterControlResponse, error) {
	if params == nil {
		params = &TrustCentersGetTrustCenterControlParams{}
	}
	path := "/trust-centers/:slugId/controls/:controlId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.ControlID == "" {
		return nil, fmt.Errorf("controlId is required")
	}
	path = strings.ReplaceAll(path, ":controlId", url.PathEscape(params.ControlID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetTrustCenterControlResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetTrustCenterControlCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TrustCentersGetTrustCenterControlCategoryParams struct {
	SlugID     string
	CategoryID string
}

// GetTrustCenterControlCategory Gets a specific control category on a Trust Center.
func (s *TrustCentersService) GetTrustCenterControlCategory(ctx context.Context, params *TrustCentersGetTrustCenterControlCategoryParams) (*TrustCentersGetTrustCenterControlCategoryResponse, error) {
	if params == nil {
		params = &TrustCentersGetTrustCenterControlCategoryParams{}
	}
	path := "/trust-centers/:slugId/control-categories/:categoryId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.CategoryID == "" {
		return nil, fmt.Errorf("categoryId is required")
	}
	path = strings.ReplaceAll(path, ":categoryId", url.PathEscape(params.CategoryID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetTrustCenterControlCategoryResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetTrustCenterDocumentResponse struct {
	CreationDate string `json:"creationDate"`
	Description  string `json:"description"`
	FileName     string `json:"fileName"`
	ID           string `json:"id"`
	IsPublic     bool   `json:"isPublic"`
	MimeType     string `json:"mimeType"`
	Title        string `json:"title"`
	UpdatedDate  string `json:"updatedDate"`
}

type TrustCentersGetTrustCenterDocumentParams struct {
	SlugID     string
	ResourceID string
}

// GetTrustCenterDocument Gets a specific document on a Trust Center.
func (s *TrustCentersService) GetTrustCenterDocument(ctx context.Context, params *TrustCentersGetTrustCenterDocumentParams) (*TrustCentersGetTrustCenterDocumentResponse, error) {
	if params == nil {
		params = &TrustCentersGetTrustCenterDocumentParams{}
	}
	path := "/trust-centers/:slugId/resources/:resourceId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.ResourceID == "" {
		return nil, fmt.Errorf("resourceId is required")
	}
	path = strings.ReplaceAll(path, ":resourceId", url.PathEscape(params.ResourceID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetTrustCenterDocumentResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetTrustCenterFaqResponse struct {
	Answer   string `json:"answer"`
	ID       string `json:"id"`
	Question string `json:"question"`
}

type TrustCentersGetTrustCenterFaqParams struct {
	SlugID string
	FaqID  string
}

// GetTrustCenterFaq Gets a specific FAQ on the Trust Center by ID.
func (s *TrustCentersService) GetTrustCenterFaq(ctx context.Context, params *TrustCentersGetTrustCenterFaqParams) (*TrustCentersGetTrustCenterFaqResponse, error) {
	if params == nil {
		params = &TrustCentersGetTrustCenterFaqParams{}
	}
	path := "/trust-centers/:slugId/faqs/:faqId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.FaqID == "" {
		return nil, fmt.Errorf("faqId is required")
	}
	path = strings.ReplaceAll(path, ":faqId", url.PathEscape(params.FaqID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetTrustCenterFaqResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetTrustCenterSubprocessorResponse struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Purpose     string `json:"purpose"`
	URL         string `json:"url"`
}

type TrustCentersGetTrustCenterSubprocessorParams struct {
	SlugID         string
	SubprocessorID string
}

// GetTrustCenterSubprocessor Gets a specific subprocessor on a Trust Center.
func (s *TrustCentersService) GetTrustCenterSubprocessor(ctx context.Context, params *TrustCentersGetTrustCenterSubprocessorParams) (*TrustCentersGetTrustCenterSubprocessorResponse, error) {
	if params == nil {
		params = &TrustCentersGetTrustCenterSubprocessorParams{}
	}
	path := "/trust-centers/:slugId/subprocessors/:subprocessorId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.SubprocessorID == "" {
		return nil, fmt.Errorf("subprocessorId is required")
	}
	path = strings.ReplaceAll(path, ":subprocessorId", url.PathEscape(params.SubprocessorID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetTrustCenterSubprocessorResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetTrustCenterSubscriberResponse struct {
	CreationDate    string `json:"creationDate"`
	Email           string `json:"email"`
	ID              string `json:"id"`
	IsEmailVerified bool   `json:"isEmailVerified"`
}

type TrustCentersGetTrustCenterSubscriberParams struct {
	SlugID       string
	SubscriberID string
}

// GetTrustCenterSubscriber Gets a specific subscriber on a Trust Center.
func (s *TrustCentersService) GetTrustCenterSubscriber(ctx context.Context, params *TrustCentersGetTrustCenterSubscriberParams) (*TrustCentersGetTrustCenterSubscriberResponse, error) {
	if params == nil {
		params = &TrustCentersGetTrustCenterSubscriberParams{}
	}
	path := "/trust-centers/:slugId/subscribers/:subscriberId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.SubscriberID == "" {
		return nil, fmt.Errorf("subscriberId is required")
	}
	path = strings.ReplaceAll(path, ":subscriberId", url.PathEscape(params.SubscriberID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetTrustCenterSubscriberResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetTrustCenterSubscriberGroupResponse struct {
	CreationDate  string   `json:"creationDate"`
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	SubscriberIDs []string `json:"subscriberIds"`
}

type TrustCentersGetTrustCenterSubscriberGroupParams struct {
	SlugID            string
	SubscriberGroupID string
}

// GetTrustCenterSubscriberGroup Get a subscriber group by ID.
func (s *TrustCentersService) GetTrustCenterSubscriberGroup(ctx context.Context, params *TrustCentersGetTrustCenterSubscriberGroupParams) (*TrustCentersGetTrustCenterSubscriberGroupResponse, error) {
	if params == nil {
		params = &TrustCentersGetTrustCenterSubscriberGroupParams{}
	}
	path := "/trust-centers/:slugId/subscriber-groups/:subscriberGroupId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.SubscriberGroupID == "" {
		return nil, fmt.Errorf("subscriberGroupId is required")
	}
	path = strings.ReplaceAll(path, ":subscriberGroupId", url.PathEscape(params.SubscriberGroupID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetTrustCenterSubscriberGroupResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetTrustCenterUpdateResponse struct {
	Category       string   `json:"category"`
	CreationDate   string   `json:"creationDate"`
	Description    string   `json:"description"`
	ID             string   `json:"id"`
	NotifiedEmails []string `json:"notifiedEmails"`
	Title          string   `json:"title"`
	UpdatedDate    string   `json:"updatedDate"`
	VisibilityType string   `json:"visibilityType"`
}

type TrustCentersGetTrustCenterUpdateParams struct {
	SlugID   string
	UpdateID string
}

// GetTrustCenterUpdate Gets a specific update on a Trust Center.
func (s *TrustCentersService) GetTrustCenterUpdate(ctx context.Context, params *TrustCentersGetTrustCenterUpdateParams) (*TrustCentersGetTrustCenterUpdateResponse, error) {
	if params == nil {
		params = &TrustCentersGetTrustCenterUpdateParams{}
	}
	path := "/trust-centers/:slugId/updates/:updateId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.UpdateID == "" {
		return nil, fmt.Errorf("updateId is required")
	}
	path = strings.ReplaceAll(path, ":updateId", url.PathEscape(params.UpdateID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetTrustCenterUpdateResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetTrustCenterViewerResponse struct {
	AccessLevel                 string           `json:"accessLevel"`
	AddedByUser                 any              `json:"addedByUser"`
	CompanyName                 string           `json:"companyName"`
	CreationDate                string           `json:"creationDate"`
	Email                       string           `json:"email"`
	ExpirationDate              any              `json:"expirationDate"`
	ExternalServiceAssociations []map[string]any `json:"externalServiceAssociations"`
	ID                          string           `json:"id"`
	Name                        string           `json:"name"`
	NdaInfo                     any              `json:"ndaInfo"`
	ResourceIDs                 any              `json:"resourceIds"`
	UpdatedDate                 string           `json:"updatedDate"`
}

type TrustCentersGetTrustCenterViewerParams struct {
	SlugID   string
	ViewerID string
}

// GetTrustCenterViewer Gets a specific viewer for a Trust Center.
func (s *TrustCentersService) GetTrustCenterViewer(ctx context.Context, params *TrustCentersGetTrustCenterViewerParams) (*TrustCentersGetTrustCenterViewerResponse, error) {
	if params == nil {
		params = &TrustCentersGetTrustCenterViewerParams{}
	}
	path := "/trust-centers/:slugId/viewers/:viewerId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.ViewerID == "" {
		return nil, fmt.Errorf("viewerId is required")
	}
	path = strings.ReplaceAll(path, ":viewerId", url.PathEscape(params.ViewerID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetTrustCenterViewerResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersGetUploadedMediaForTrustCenterDocumentResponse struct {
	Readable bool `json:"readable"`
}

type TrustCentersGetUploadedMediaForTrustCenterDocumentParams struct {
	SlugID     string
	ResourceID string
}

// GetUploadedMediaForTrustCenterDocument Gets the actual given uploaded document for a Trust Center.
func (s *TrustCentersService) GetUploadedMediaForTrustCenterDocument(ctx context.Context, params *TrustCentersGetUploadedMediaForTrustCenterDocumentParams) (*TrustCentersGetUploadedMediaForTrustCenterDocumentResponse, error) {
	if params == nil {
		params = &TrustCentersGetUploadedMediaForTrustCenterDocumentParams{}
	}
	path := "/trust-centers/:slugId/resources/:resourceId/media"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.ResourceID == "" {
		return nil, fmt.Errorf("resourceId is required")
	}
	path = strings.ReplaceAll(path, ":resourceId", url.PathEscape(params.ResourceID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersGetUploadedMediaForTrustCenterDocumentResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListHistoricalTrustCenterAccessRequestsResponse struct {
	Results map[string]any `json:"results"`
}

type TrustCentersListHistoricalTrustCenterAccessRequestsParams struct {
	SlugID     string
	PageSize   *int
	PageCursor *string
}

// ListHistoricalTrustCenterAccessRequests Gets a list of historical (approved or denied) access requests for a Trust Center.
func (s *TrustCentersService) ListHistoricalTrustCenterAccessRequests(ctx context.Context, params *TrustCentersListHistoricalTrustCenterAccessRequestsParams) (*TrustCentersListHistoricalTrustCenterAccessRequestsResponse, error) {
	if params == nil {
		params = &TrustCentersListHistoricalTrustCenterAccessRequestsParams{}
	}
	path := "/trust-centers/:slugId/historical-access-requests"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListHistoricalTrustCenterAccessRequestsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListTrustCenterAccessRequestsResponse struct {
	Results map[string]any `json:"results"`
}

type TrustCentersListTrustCenterAccessRequestsParams struct {
	SlugID     string
	PageSize   *int
	PageCursor *string
}

// ListTrustCenterAccessRequests Gets a list of access requests for a Trust Center.
func (s *TrustCentersService) ListTrustCenterAccessRequests(ctx context.Context, params *TrustCentersListTrustCenterAccessRequestsParams) (*TrustCentersListTrustCenterAccessRequestsResponse, error) {
	if params == nil {
		params = &TrustCentersListTrustCenterAccessRequestsParams{}
	}
	path := "/trust-centers/:slugId/access-requests"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListTrustCenterAccessRequestsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListTrustCenterControlCategoriesResponse struct {
	Results []map[string]any `json:"results"`
}

type TrustCentersListTrustCenterControlCategoriesParams struct {
	SlugID string
}

// ListTrustCenterControlCategories Gets a list of control categories on a Trust Center.
func (s *TrustCentersService) ListTrustCenterControlCategories(ctx context.Context, params *TrustCentersListTrustCenterControlCategoriesParams) (*TrustCentersListTrustCenterControlCategoriesResponse, error) {
	if params == nil {
		params = &TrustCentersListTrustCenterControlCategoriesParams{}
	}
	path := "/trust-centers/:slugId/control-categories"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListTrustCenterControlCategoriesResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListTrustCenterControlsResponse struct {
	Results map[string]any `json:"results"`
}

type TrustCentersListTrustCenterControlsParams struct {
	SlugID     string
	PageSize   *int
	PageCursor *string
}

// ListTrustCenterControls Gets a list of controls on a Trust Center.
func (s *TrustCentersService) ListTrustCenterControls(ctx context.Context, params *TrustCentersListTrustCenterControlsParams) (*TrustCentersListTrustCenterControlsResponse, error) {
	if params == nil {
		params = &TrustCentersListTrustCenterControlsParams{}
	}
	path := "/trust-centers/:slugId/controls"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListTrustCenterControlsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListTrustCenterFaqsResponse struct {
	Results []map[string]any `json:"results"`
}

type TrustCentersListTrustCenterFaqsParams struct {
	SlugID string
}

// ListTrustCenterFaqs Gets a list of FAQs on a Trust Center.
func (s *TrustCentersService) ListTrustCenterFaqs(ctx context.Context, params *TrustCentersListTrustCenterFaqsParams) (*TrustCentersListTrustCenterFaqsResponse, error) {
	if params == nil {
		params = &TrustCentersListTrustCenterFaqsParams{}
	}
	path := "/trust-centers/:slugId/faqs"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListTrustCenterFaqsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListTrustCenterResourcesResponse struct {
	Results []map[string]any `json:"results"`
}

type TrustCentersListTrustCenterResourcesParams struct {
	SlugID string
}

// ListTrustCenterResources Gets a list of resources on a Trust Center.
func (s *TrustCentersService) ListTrustCenterResources(ctx context.Context, params *TrustCentersListTrustCenterResourcesParams) (*TrustCentersListTrustCenterResourcesResponse, error) {
	if params == nil {
		params = &TrustCentersListTrustCenterResourcesParams{}
	}
	path := "/trust-centers/:slugId/resources"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListTrustCenterResourcesResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListTrustCenterSubprocessorsResponse struct {
	Results []map[string]any `json:"results"`
}

type TrustCentersListTrustCenterSubprocessorsParams struct {
	SlugID string
}

// ListTrustCenterSubprocessors Gets the list of subprocessors on a Trust Center.
func (s *TrustCentersService) ListTrustCenterSubprocessors(ctx context.Context, params *TrustCentersListTrustCenterSubprocessorsParams) (*TrustCentersListTrustCenterSubprocessorsResponse, error) {
	if params == nil {
		params = &TrustCentersListTrustCenterSubprocessorsParams{}
	}
	path := "/trust-centers/:slugId/subprocessors"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListTrustCenterSubprocessorsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListTrustCenterSubscriberGroupsResponse struct {
	Results map[string]any `json:"results"`
}

type TrustCentersListTrustCenterSubscriberGroupsParams struct {
	SlugID     string
	PageSize   *int
	PageCursor *string
}

// ListTrustCenterSubscriberGroups Gets a list of subscriber groups on a Trust Center.
func (s *TrustCentersService) ListTrustCenterSubscriberGroups(ctx context.Context, params *TrustCentersListTrustCenterSubscriberGroupsParams) (*TrustCentersListTrustCenterSubscriberGroupsResponse, error) {
	if params == nil {
		params = &TrustCentersListTrustCenterSubscriberGroupsParams{}
	}
	path := "/trust-centers/:slugId/subscriber-groups"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListTrustCenterSubscriberGroupsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListTrustCenterSubscribersResponse struct {
	Results map[string]any `json:"results"`
}

type TrustCentersListTrustCenterSubscribersParams struct {
	SlugID     string
	PageSize   *int
	PageCursor *string
}

// ListTrustCenterSubscribers Gets a list of subscribers on a Trust Center.
func (s *TrustCentersService) ListTrustCenterSubscribers(ctx context.Context, params *TrustCentersListTrustCenterSubscribersParams) (*TrustCentersListTrustCenterSubscribersResponse, error) {
	if params == nil {
		params = &TrustCentersListTrustCenterSubscribersParams{}
	}
	path := "/trust-centers/:slugId/subscribers"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListTrustCenterSubscribersResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListTrustCenterUpdatesResponse struct {
	Results map[string]any `json:"results"`
}

type TrustCentersListTrustCenterUpdatesParams struct {
	SlugID     string
	PageSize   *int
	PageCursor *string
}

// ListTrustCenterUpdates Gets a list of updates on a Trust Center.
func (s *TrustCentersService) ListTrustCenterUpdates(ctx context.Context, params *TrustCentersListTrustCenterUpdatesParams) (*TrustCentersListTrustCenterUpdatesResponse, error) {
	if params == nil {
		params = &TrustCentersListTrustCenterUpdatesParams{}
	}
	path := "/trust-centers/:slugId/updates"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListTrustCenterUpdatesResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListTrustCenterViewerActivityEventsResponse struct {
	Results map[string]any `json:"results"`
}

type TrustCentersListTrustCenterViewerActivityEventsParams struct {
	SlugID               string
	PageSize             *int
	PageCursor           *string
	EventTypesMatchesAny []string
	AfterDate            *string
	BeforeDate           *string
}

// ListTrustCenterViewerActivityEvents Gets a list of viewer activity events on a Trust Center.
func (s *TrustCentersService) ListTrustCenterViewerActivityEvents(ctx context.Context, params *TrustCentersListTrustCenterViewerActivityEventsParams) (*TrustCentersListTrustCenterViewerActivityEventsResponse, error) {
	if params == nil {
		params = &TrustCentersListTrustCenterViewerActivityEventsParams{}
	}
	path := "/trust-centers/:slugId/activity"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	for _, v := range params.EventTypesMatchesAny {
		query.Add("eventTypesMatchesAny", v)
	}
	if params.AfterDate != nil {
		query.Set("afterDate", fmt.Sprint(*params.AfterDate))
	}
	if params.BeforeDate != nil {
		query.Set("beforeDate", fmt.Sprint(*params.BeforeDate))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListTrustCenterViewerActivityEventsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersListTrustCenterViewersResponse struct {
	Results map[string]any `json:"results"`
}

type TrustCentersListTrustCenterViewersParams struct {
	SlugID         string
	PageSize       *int
	PageCursor     *string
	IncludeRemoved *bool
}

// ListTrustCenterViewers Gets a list of viewers that have been granted access to a Trust Center.
func (s *TrustCentersService) ListTrustCenterViewers(ctx context.Context, params *TrustCentersListTrustCenterViewersParams) (*TrustCentersListTrustCenterViewersResponse, error) {
	if params == nil {
		params = &TrustCentersListTrustCenterViewersParams{}
	}
	path := "/trust-centers/:slugId/viewers"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	if params.IncludeRemoved != nil {
		query.Set("includeRemoved", fmt.Sprint(*params.IncludeRemoved))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersListTrustCenterViewersResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersRemoveTrustCenterViewerParams struct {
	SlugID   string
	ViewerID string
}

// RemoveTrustCenterViewer Revokes a viewer's access to a Trust Center.
func (s *TrustCentersService) RemoveTrustCenterViewer(ctx context.Context, params *TrustCentersRemoveTrustCenterViewerParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersRemoveTrustCenterViewerParams{}
	}
	path := "/trust-centers/:slugId/viewers/:viewerId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.ViewerID == "" {
		return nil, fmt.Errorf("viewerId is required")
	}
	path = strings.ReplaceAll(path, ":viewerId", url.PathEscape(params.ViewerID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersSendTrustCenterUpdateNotificationsToAllSubscribersParams struct {
	SlugID   string
	UpdateID string
}

// SendTrustCenterUpdateNotificationsToAllSubscribers Sends notifications for a specific Trust Center update to all subscribers.
func (s *TrustCentersService) SendTrustCenterUpdateNotificationsToAllSubscribers(ctx context.Context, params *TrustCentersSendTrustCenterUpdateNotificationsToAllSubscribersParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersSendTrustCenterUpdateNotificationsToAllSubscribersParams{}
	}
	path := "/trust-centers/:slugId/updates/:updateId/notify-all-subscribers"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.UpdateID == "" {
		return nil, fmt.Errorf("updateId is required")
	}
	path = strings.ReplaceAll(path, ":updateId", url.PathEscape(params.UpdateID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersSendTrustCenterUpdateNotificationsToSpecificSubscribersRequestBody struct {
	Emails             []string `json:"emails"`
	SubscriberGroupIDs []string `json:"subscriberGroupIds"`
}

type TrustCentersSendTrustCenterUpdateNotificationsToSpecificSubscribersParams struct {
	SlugID   string
	UpdateID string
	Body     *TrustCentersSendTrustCenterUpdateNotificationsToSpecificSubscribersRequestBody
}

// SendTrustCenterUpdateNotificationsToSpecificSubscribers Sends notifications for a specific Trust Center update to specific subscribers. At least one subscriber group or email address is required.
func (s *TrustCentersService) SendTrustCenterUpdateNotificationsToSpecificSubscribers(ctx context.Context, params *TrustCentersSendTrustCenterUpdateNotificationsToSpecificSubscribersParams) (json.RawMessage, error) {
	if params == nil {
		params = &TrustCentersSendTrustCenterUpdateNotificationsToSpecificSubscribersParams{}
	}
	path := "/trust-centers/:slugId/updates/:updateId/notify-specific-subscribers"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.UpdateID == "" {
		return nil, fmt.Errorf("updateId is required")
	}
	path = strings.ReplaceAll(path, ":updateId", url.PathEscape(params.UpdateID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersSetGroupsForTrustCenterSubscriberRequestBody struct {
	GroupIDs []string `json:"groupIds"`
}

type TrustCentersSetGroupsForTrustCenterSubscriberResponse struct {
	CreationDate    string `json:"creationDate"`
	Email           string `json:"email"`
	ID              string `json:"id"`
	IsEmailVerified bool   `json:"isEmailVerified"`
}

type TrustCentersSetGroupsForTrustCenterSubscriberParams struct {
	SlugID       string
	SubscriberID string
	Body         *TrustCentersSetGroupsForTrustCenterSubscriberRequestBody
}

// SetGroupsForTrustCenterSubscriber Sets groups on a subscriber.
func (s *TrustCentersService) SetGroupsForTrustCenterSubscriber(ctx context.Context, params *TrustCentersSetGroupsForTrustCenterSubscriberParams) (*TrustCentersSetGroupsForTrustCenterSubscriberResponse, error) {
	if params == nil {
		params = &TrustCentersSetGroupsForTrustCenterSubscriberParams{}
	}
	path := "/trust-centers/:slugId/subscribers/:subscriberId/groups"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.SubscriberID == "" {
		return nil, fmt.Errorf("subscriberId is required")
	}
	path = strings.ReplaceAll(path, ":subscriberId", url.PathEscape(params.SubscriberID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PUT", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersSetGroupsForTrustCenterSubscriberResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersUpdateTrustCenterRequestBody struct {
	BannerSetting      map[string]any `json:"bannerSetting"`
	CompanyDescription string         `json:"companyDescription"`
	CustomTheme        map[string]any `json:"customTheme"`
	IsPublic           bool           `json:"isPublic"`
	PrivacyPolicy      string         `json:"privacyPolicy"`
	Title              string         `json:"title"`
}

type TrustCentersUpdateTrustCenterResponse struct {
	BannerSetting      map[string]any `json:"bannerSetting"`
	CompanyDescription string         `json:"companyDescription"`
	CreationDate       string         `json:"creationDate"`
	CustomDomain       string         `json:"customDomain"`
	CustomTheme        map[string]any `json:"customTheme"`
	ID                 string         `json:"id"`
	IsPublic           bool           `json:"isPublic"`
	PrivacyPolicy      string         `json:"privacyPolicy"`
	Title              string         `json:"title"`
	UpdatedDate        string         `json:"updatedDate"`
}

type TrustCentersUpdateTrustCenterParams struct {
	SlugID string
	Body   *TrustCentersUpdateTrustCenterRequestBody
}

// UpdateTrustCenter Updates a Trust Center by slug ID.
func (s *TrustCentersService) UpdateTrustCenter(ctx context.Context, params *TrustCentersUpdateTrustCenterParams) (*TrustCentersUpdateTrustCenterResponse, error) {
	if params == nil {
		params = &TrustCentersUpdateTrustCenterParams{}
	}
	path := "/trust-centers/:slugId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersUpdateTrustCenterResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersUpdateTrustCenterControlCategoryRequestBody struct {
	Name string `json:"name"`
}

type TrustCentersUpdateTrustCenterControlCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TrustCentersUpdateTrustCenterControlCategoryParams struct {
	SlugID     string
	CategoryID string
	Body       *TrustCentersUpdateTrustCenterControlCategoryRequestBody
}

// UpdateTrustCenterControlCategory Updates a control category on a Trust Center.
func (s *TrustCentersService) UpdateTrustCenterControlCategory(ctx context.Context, params *TrustCentersUpdateTrustCenterControlCategoryParams) (*TrustCentersUpdateTrustCenterControlCategoryResponse, error) {
	if params == nil {
		params = &TrustCentersUpdateTrustCenterControlCategoryParams{}
	}
	path := "/trust-centers/:slugId/control-categories/:categoryId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.CategoryID == "" {
		return nil, fmt.Errorf("categoryId is required")
	}
	path = strings.ReplaceAll(path, ":categoryId", url.PathEscape(params.CategoryID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersUpdateTrustCenterControlCategoryResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersUpdateTrustCenterDocumentRequestBody struct {
	Description string `json:"description"`
	IsPublic    bool   `json:"isPublic"`
	Title       string `json:"title"`
}

type TrustCentersUpdateTrustCenterDocumentResponse struct {
	CreationDate string `json:"creationDate"`
	Description  string `json:"description"`
	FileName     string `json:"fileName"`
	ID           string `json:"id"`
	IsPublic     bool   `json:"isPublic"`
	MimeType     string `json:"mimeType"`
	Title        string `json:"title"`
	UpdatedDate  string `json:"updatedDate"`
}

type TrustCentersUpdateTrustCenterDocumentParams struct {
	SlugID     string
	ResourceID string
	Body       *TrustCentersUpdateTrustCenterDocumentRequestBody
}

// UpdateTrustCenterDocument Updates a specific document on a Trust Center.
func (s *TrustCentersService) UpdateTrustCenterDocument(ctx context.Context, params *TrustCentersUpdateTrustCenterDocumentParams) (*TrustCentersUpdateTrustCenterDocumentResponse, error) {
	if params == nil {
		params = &TrustCentersUpdateTrustCenterDocumentParams{}
	}
	path := "/trust-centers/:slugId/resources/:resourceId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.ResourceID == "" {
		return nil, fmt.Errorf("resourceId is required")
	}
	path = strings.ReplaceAll(path, ":resourceId", url.PathEscape(params.ResourceID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersUpdateTrustCenterDocumentResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersUpdateTrustCenterFaqRequestBody struct {
	Answer   string `json:"answer"`
	Question string `json:"question"`
}

type TrustCentersUpdateTrustCenterFaqResponse struct {
	Answer   string `json:"answer"`
	ID       string `json:"id"`
	Question string `json:"question"`
}

type TrustCentersUpdateTrustCenterFaqParams struct {
	SlugID string
	FaqID  string
	Body   *TrustCentersUpdateTrustCenterFaqRequestBody
}

// UpdateTrustCenterFaq Update a specific FAQ on the Trust Center by ID.
func (s *TrustCentersService) UpdateTrustCenterFaq(ctx context.Context, params *TrustCentersUpdateTrustCenterFaqParams) (*TrustCentersUpdateTrustCenterFaqResponse, error) {
	if params == nil {
		params = &TrustCentersUpdateTrustCenterFaqParams{}
	}
	path := "/trust-centers/:slugId/faqs/:faqId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.FaqID == "" {
		return nil, fmt.Errorf("faqId is required")
	}
	path = strings.ReplaceAll(path, ":faqId", url.PathEscape(params.FaqID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersUpdateTrustCenterFaqResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersUpdateTrustCenterSubprocessorRequestBody struct {
	Description string `json:"description"`
	Location    string `json:"location"`
	Purpose     string `json:"purpose"`
}

type TrustCentersUpdateTrustCenterSubprocessorResponse struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Purpose     string `json:"purpose"`
	URL         string `json:"url"`
}

type TrustCentersUpdateTrustCenterSubprocessorParams struct {
	SlugID         string
	SubprocessorID string
	Body           *TrustCentersUpdateTrustCenterSubprocessorRequestBody
}

// UpdateTrustCenterSubprocessor Updates a subprocessor on a Trust Center.
func (s *TrustCentersService) UpdateTrustCenterSubprocessor(ctx context.Context, params *TrustCentersUpdateTrustCenterSubprocessorParams) (*TrustCentersUpdateTrustCenterSubprocessorResponse, error) {
	if params == nil {
		params = &TrustCentersUpdateTrustCenterSubprocessorParams{}
	}
	path := "/trust-centers/:slugId/subprocessors/:subprocessorId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.SubprocessorID == "" {
		return nil, fmt.Errorf("subprocessorId is required")
	}
	path = strings.ReplaceAll(path, ":subprocessorId", url.PathEscape(params.SubprocessorID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersUpdateTrustCenterSubprocessorResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrustCentersUpdateTrustCenterUpdateRequestBody struct {
	Category       string `json:"category"`
	Description    string `json:"description"`
	Title          string `json:"title"`
	VisibilityType string `json:"visibilityType"`
}

type TrustCentersUpdateTrustCenterUpdateResponse struct {
	Category       string   `json:"category"`
	CreationDate   string   `json:"creationDate"`
	Description    string   `json:"description"`
	ID             string   `json:"id"`
	NotifiedEmails []string `json:"notifiedEmails"`
	Title          string   `json:"title"`
	UpdatedDate    string   `json:"updatedDate"`
	VisibilityType string   `json:"visibilityType"`
}

type TrustCentersUpdateTrustCenterUpdateParams struct {
	SlugID   string
	UpdateID string
	Body     *TrustCentersUpdateTrustCenterUpdateRequestBody
}

// UpdateTrustCenterUpdate Updates an update on a Trust Center.
func (s *TrustCentersService) UpdateTrustCenterUpdate(ctx context.Context, params *TrustCentersUpdateTrustCenterUpdateParams) (*TrustCentersUpdateTrustCenterUpdateResponse, error) {
	if params == nil {
		params = &TrustCentersUpdateTrustCenterUpdateParams{}
	}
	path := "/trust-centers/:slugId/updates/:updateId"
	if params.SlugID == "" {
		return nil, fmt.Errorf("slugId is required")
	}
	path = strings.ReplaceAll(path, ":slugId", url.PathEscape(params.SlugID))
	if params.UpdateID == "" {
		return nil, fmt.Errorf("updateId is required")
	}
	path = strings.ReplaceAll(path, ":updateId", url.PathEscape(params.UpdateID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &TrustCentersUpdateTrustCenterUpdateResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// VendorRiskAttributesService groups 1 endpoint methods under the "VendorRiskAttributes" API segment.
type VendorRiskAttributesService struct {
	client *Client
}

type VendorRiskAttributesListVendorRiskAttributesResponse struct {
	Results map[string]any `json:"results"`
}

type VendorRiskAttributesListVendorRiskAttributesParams struct {
	PageSize   *int
	PageCursor *string
}

// ListVendorRiskAttributes Returns a list of vendor risk attributes.
func (s *VendorRiskAttributesService) ListVendorRiskAttributes(ctx context.Context, params *VendorRiskAttributesListVendorRiskAttributesParams) (*VendorRiskAttributesListVendorRiskAttributesResponse, error) {
	if params == nil {
		params = &VendorRiskAttributesListVendorRiskAttributesParams{}
	}
	path := "/vendor-risk-attributes"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VendorRiskAttributesListVendorRiskAttributesResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// VendorsService groups 17 endpoint methods under the "Vendors" API segment.
type VendorsService struct {
	client *Client
}

type VendorsAddDocumentToSecurityReviewResponse struct {
	CreationDate string         `json:"creationDate"`
	DeletionDate any            `json:"deletionDate"`
	Description  string         `json:"description"`
	FileName     string         `json:"fileName"`
	ID           string         `json:"id"`
	MimeType     string         `json:"mimeType"`
	Title        string         `json:"title"`
	Type         string         `json:"type"`
	URL          string         `json:"url"`
	UpdatedDate  string         `json:"updatedDate"`
	UploadedBy   map[string]any `json:"uploadedBy"`
}

type VendorsAddDocumentToSecurityReviewParams struct {
	VendorID         string
	SecurityReviewID string
	// FormData maps multipart field names to values.
	FormData map[string]string
}

// AddDocumentToSecurityReview Add document to a security review.
func (s *VendorsService) AddDocumentToSecurityReview(ctx context.Context, params *VendorsAddDocumentToSecurityReviewParams) (*VendorsAddDocumentToSecurityReviewResponse, error) {
	if params == nil {
		params = &VendorsAddDocumentToSecurityReviewParams{}
	}
	path := "/vendors/:vendorId/security-reviews/:securityReviewId/documents"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	if params.SecurityReviewID == "" {
		return nil, fmt.Errorf("securityReviewId is required")
	}
	path = strings.ReplaceAll(path, ":securityReviewId", url.PathEscape(params.SecurityReviewID))
	query := url.Values{}
	req, err := s.client.newMultipartRequest(ctx, "POST", path, query, params.FormData)
	if err != nil {
		return nil, err
	}
	out := &VendorsAddDocumentToSecurityReviewResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsAddDocumentToVendorResponse struct {
	CreationDate string         `json:"creationDate"`
	DeletionDate any            `json:"deletionDate"`
	Description  string         `json:"description"`
	FileName     string         `json:"fileName"`
	ID           string         `json:"id"`
	MimeType     string         `json:"mimeType"`
	Title        string         `json:"title"`
	Type         string         `json:"type"`
	URL          string         `json:"url"`
	UpdatedDate  string         `json:"updatedDate"`
	UploadedBy   map[string]any `json:"uploadedBy"`
}

type VendorsAddDocumentToVendorParams struct {
	VendorID string
	// FormData maps multipart field names to values.
	FormData map[string]string
}

// AddDocumentToVendor Add document to a vendor.
func (s *VendorsService) AddDocumentToVendor(ctx context.Context, params *VendorsAddDocumentToVendorParams) (*VendorsAddDocumentToVendorResponse, error) {
	if params == nil {
		params = &VendorsAddDocumentToVendorParams{}
	}
	path := "/vendors/:vendorId/documents"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	query := url.Values{}
	req, err := s.client.newMultipartRequest(ctx, "POST", path, query, params.FormData)
	if err != nil {
		return nil, err
	}
	out := &VendorsAddDocumentToVendorResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsAddVendorFindingRequestBody struct {
	Content          string         `json:"content"`
	DocumentID       string         `json:"documentId"`
	Remediation      map[string]any `json:"remediation"`
	RiskStatus       string         `json:"riskStatus"`
	SecurityReviewID string         `json:"securityReviewId"`
}

type VendorsAddVendorFindingResponse struct {
	Content          string         `json:"content"`
	DocumentID       any            `json:"documentId"`
	ID               string         `json:"id"`
	Remediation      map[string]any `json:"remediation"`
	RiskStatus       string         `json:"riskStatus"`
	SecurityReviewID string         `json:"securityReviewId"`
	VendorID         string         `json:"vendorId"`
}

type VendorsAddVendorFindingParams struct {
	VendorID string
	Body     *VendorsAddVendorFindingRequestBody
}

// AddVendorFinding Add vendor finding.
func (s *VendorsService) AddVendorFinding(ctx context.Context, params *VendorsAddVendorFindingParams) (*VendorsAddVendorFindingResponse, error) {
	if params == nil {
		params = &VendorsAddVendorFindingParams{}
	}
	path := "/vendors/:vendorId/findings"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &VendorsAddVendorFindingResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsCreateVendorRequestBody struct {
	AccountManagerEmail     string           `json:"accountManagerEmail"`
	AccountManagerName      string           `json:"accountManagerName"`
	AdditionalNotes         string           `json:"additionalNotes"`
	AuthDetails             map[string]any   `json:"authDetails"`
	BusinessOwnerUserID     string           `json:"businessOwnerUserId"`
	Category                string           `json:"category"`
	ContractAmount          map[string]any   `json:"contractAmount"`
	ContractRenewalDate     string           `json:"contractRenewalDate"`
	ContractStartDate       string           `json:"contractStartDate"`
	ContractTerminationDate string           `json:"contractTerminationDate"`
	CustomFields            []map[string]any `json:"customFields"`
	FrameworkScope          map[string]any   `json:"frameworkScope"`
	InherentRiskLevel       string           `json:"inherentRiskLevel"`
	IsVisibleToAuditors     bool             `json:"isVisibleToAuditors"`
	Name                    string           `json:"name"`
	ResidualRiskLevel       string           `json:"residualRiskLevel"`
	SecurityOwnerUserID     string           `json:"securityOwnerUserId"`
	ServicesProvided        string           `json:"servicesProvided"`
	Status                  string           `json:"status"`
	VendorHeadquarters      string           `json:"vendorHeadquarters"`
	WebsiteURL              string           `json:"websiteUrl"`
}

type VendorsCreateVendorResponse struct {
	AccountManagerEmail              string         `json:"accountManagerEmail"`
	AccountManagerName               string         `json:"accountManagerName"`
	AdditionalNotes                  string         `json:"additionalNotes"`
	AuthDetails                      map[string]any `json:"authDetails"`
	BusinessOwnerUserID              string         `json:"businessOwnerUserId"`
	Category                         map[string]any `json:"category"`
	ContractAmount                   map[string]any `json:"contractAmount"`
	ContractRenewalDate              string         `json:"contractRenewalDate"`
	ContractStartDate                string         `json:"contractStartDate"`
	ContractTerminationDate          any            `json:"contractTerminationDate"`
	CustomFields                     any            `json:"customFields"`
	ID                               string         `json:"id"`
	InherentRiskLevel                string         `json:"inherentRiskLevel"`
	IsRiskAutoScored                 bool           `json:"isRiskAutoScored"`
	IsVisibleToAuditors              bool           `json:"isVisibleToAuditors"`
	LastSecurityReviewCompletionDate string         `json:"lastSecurityReviewCompletionDate"`
	LatestDecision                   map[string]any `json:"latestDecision"`
	Name                             string         `json:"name"`
	NextSecurityReviewDueDate        string         `json:"nextSecurityReviewDueDate"`
	ResidualRiskLevel                string         `json:"residualRiskLevel"`
	RiskAttributeIDs                 []string       `json:"riskAttributeIds"`
	SecurityOwnerUserID              string         `json:"securityOwnerUserId"`
	ServicesProvided                 string         `json:"servicesProvided"`
	Status                           string         `json:"status"`
	TagIDentifiers                   any            `json:"tagIdentifiers"`
	VendorHeadquarters               string         `json:"vendorHeadquarters"`
	WebsiteURL                       string         `json:"websiteUrl"`
}

type VendorsCreateVendorParams struct {
	Body *VendorsCreateVendorRequestBody
}

// CreateVendor Add vendor with metadata.
func (s *VendorsService) CreateVendor(ctx context.Context, params *VendorsCreateVendorParams) (*VendorsCreateVendorResponse, error) {
	if params == nil {
		params = &VendorsCreateVendorParams{}
	}
	path := "/vendors"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &VendorsCreateVendorResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsDeleteFindingByIDParams struct {
	VendorID  string
	FindingID string
}

// DeleteFindingByID Deletes a finding.
func (s *VendorsService) DeleteFindingByID(ctx context.Context, params *VendorsDeleteFindingByIDParams) (json.RawMessage, error) {
	if params == nil {
		params = &VendorsDeleteFindingByIDParams{}
	}
	path := "/vendors/:vendorId/findings/:findingId"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	if params.FindingID == "" {
		return nil, fmt.Errorf("findingId is required")
	}
	path = strings.ReplaceAll(path, ":findingId", url.PathEscape(params.FindingID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsDeleteSecurityReviewDocumentByIDParams struct {
	VendorID         string
	SecurityReviewID string
	DocumentID       string
}

// DeleteSecurityReviewDocumentByID Delete a security review document.
func (s *VendorsService) DeleteSecurityReviewDocumentByID(ctx context.Context, params *VendorsDeleteSecurityReviewDocumentByIDParams) (json.RawMessage, error) {
	if params == nil {
		params = &VendorsDeleteSecurityReviewDocumentByIDParams{}
	}
	path := "/vendors/:vendorId/security-reviews/:securityReviewId/documents/:documentId"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	if params.SecurityReviewID == "" {
		return nil, fmt.Errorf("securityReviewId is required")
	}
	path = strings.ReplaceAll(path, ":securityReviewId", url.PathEscape(params.SecurityReviewID))
	if params.DocumentID == "" {
		return nil, fmt.Errorf("documentId is required")
	}
	path = strings.ReplaceAll(path, ":documentId", url.PathEscape(params.DocumentID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsDeleteVendorByIDParams struct {
	VendorID string
}

// DeleteVendorByID Deletes a vendor.
func (s *VendorsService) DeleteVendorByID(ctx context.Context, params *VendorsDeleteVendorByIDParams) (json.RawMessage, error) {
	if params == nil {
		params = &VendorsDeleteVendorByIDParams{}
	}
	path := "/vendors/:vendorId"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "DELETE", path, query, nil)
	if err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := s.client.doJSON(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsGetSecurityReviewByIDResponse struct {
	Comments          string         `json:"comments"`
	CompletedByUserID string         `json:"completedByUserId"`
	CompletionDate    string         `json:"completionDate"`
	Decision          map[string]any `json:"decision"`
	DecisionNotes     string         `json:"decisionNotes"`
	DueDate           string         `json:"dueDate"`
	ID                string         `json:"id"`
	OverrideDueDate   string         `json:"overrideDueDate"`
	StartDate         string         `json:"startDate"`
	VendorID          string         `json:"vendorId"`
}

type VendorsGetSecurityReviewByIDParams struct {
	VendorID         string
	SecurityReviewID string
}

// GetSecurityReviewByID Returns a security review.
func (s *VendorsService) GetSecurityReviewByID(ctx context.Context, params *VendorsGetSecurityReviewByIDParams) (*VendorsGetSecurityReviewByIDResponse, error) {
	if params == nil {
		params = &VendorsGetSecurityReviewByIDParams{}
	}
	path := "/vendors/:vendorId/security-reviews/:securityReviewId"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	if params.SecurityReviewID == "" {
		return nil, fmt.Errorf("securityReviewId is required")
	}
	path = strings.ReplaceAll(path, ":securityReviewId", url.PathEscape(params.SecurityReviewID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VendorsGetSecurityReviewByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsGetVendorByIDResponse struct {
	AccountManagerEmail              string         `json:"accountManagerEmail"`
	AccountManagerName               string         `json:"accountManagerName"`
	AdditionalNotes                  string         `json:"additionalNotes"`
	AuthDetails                      map[string]any `json:"authDetails"`
	BusinessOwnerUserID              string         `json:"businessOwnerUserId"`
	Category                         map[string]any `json:"category"`
	ContractAmount                   map[string]any `json:"contractAmount"`
	ContractRenewalDate              string         `json:"contractRenewalDate"`
	ContractStartDate                string         `json:"contractStartDate"`
	ContractTerminationDate          any            `json:"contractTerminationDate"`
	CustomFields                     any            `json:"customFields"`
	ID                               string         `json:"id"`
	InherentRiskLevel                string         `json:"inherentRiskLevel"`
	IsRiskAutoScored                 bool           `json:"isRiskAutoScored"`
	IsVisibleToAuditors              bool           `json:"isVisibleToAuditors"`
	LastSecurityReviewCompletionDate string         `json:"lastSecurityReviewCompletionDate"`
	LatestDecision                   map[string]any `json:"latestDecision"`
	Name                             string         `json:"name"`
	NextSecurityReviewDueDate        string         `json:"nextSecurityReviewDueDate"`
	ResidualRiskLevel                string         `json:"residualRiskLevel"`
	RiskAttributeIDs                 []string       `json:"riskAttributeIds"`
	SecurityOwnerUserID              string         `json:"securityOwnerUserId"`
	ServicesProvided                 string         `json:"servicesProvided"`
	Status                           string         `json:"status"`
	TagIDentifiers                   any            `json:"tagIdentifiers"`
	VendorHeadquarters               string         `json:"vendorHeadquarters"`
	WebsiteURL                       string         `json:"websiteUrl"`
}

type VendorsGetVendorByIDParams struct {
	VendorID string
}

// GetVendorByID Get a vendor.
func (s *VendorsService) GetVendorByID(ctx context.Context, params *VendorsGetVendorByIDParams) (*VendorsGetVendorByIDResponse, error) {
	if params == nil {
		params = &VendorsGetVendorByIDParams{}
	}
	path := "/vendors/:vendorId"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VendorsGetVendorByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsListSecurityReviewDocumentsResponse struct {
	Results map[string]any `json:"results"`
}

type VendorsListSecurityReviewDocumentsParams struct {
	VendorID         string
	SecurityReviewID string
	PageSize         *int
	PageCursor       *string
}

// ListSecurityReviewDocuments Lists a security review's documents.
func (s *VendorsService) ListSecurityReviewDocuments(ctx context.Context, params *VendorsListSecurityReviewDocumentsParams) (*VendorsListSecurityReviewDocumentsResponse, error) {
	if params == nil {
		params = &VendorsListSecurityReviewDocumentsParams{}
	}
	path := "/vendors/:vendorId/security-reviews/:securityReviewId/documents"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	if params.SecurityReviewID == "" {
		return nil, fmt.Errorf("securityReviewId is required")
	}
	path = strings.ReplaceAll(path, ":securityReviewId", url.PathEscape(params.SecurityReviewID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VendorsListSecurityReviewDocumentsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsListSecurityReviewsByVendorIDResponse struct {
	Results map[string]any `json:"results"`
}

type VendorsListSecurityReviewsByVendorIDParams struct {
	VendorID   string
	PageSize   *int
	PageCursor *string
}

// ListSecurityReviewsByVendorID Returns a vendor's security reviews.
func (s *VendorsService) ListSecurityReviewsByVendorID(ctx context.Context, params *VendorsListSecurityReviewsByVendorIDParams) (*VendorsListSecurityReviewsByVendorIDResponse, error) {
	if params == nil {
		params = &VendorsListSecurityReviewsByVendorIDParams{}
	}
	path := "/vendors/:vendorId/security-reviews"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VendorsListSecurityReviewsByVendorIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsListVendorDocumentsResponse struct {
	Results map[string]any `json:"results"`
}

type VendorsListVendorDocumentsParams struct {
	VendorID   string
	PageSize   *int
	PageCursor *string
}

// ListVendorDocuments Returns a vendor's list of documents.
func (s *VendorsService) ListVendorDocuments(ctx context.Context, params *VendorsListVendorDocumentsParams) (*VendorsListVendorDocumentsResponse, error) {
	if params == nil {
		params = &VendorsListVendorDocumentsParams{}
	}
	path := "/vendors/:vendorId/documents"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VendorsListVendorDocumentsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsListVendorFindingsResponse struct {
	Results map[string]any `json:"results"`
}

type VendorsListVendorFindingsParams struct {
	VendorID         string
	PageSize         *int
	PageCursor       *string
	SecurityReviewID *string
	DocumentID       *string
}

// ListVendorFindings Lists a vendor's findings.
func (s *VendorsService) ListVendorFindings(ctx context.Context, params *VendorsListVendorFindingsParams) (*VendorsListVendorFindingsResponse, error) {
	if params == nil {
		params = &VendorsListVendorFindingsParams{}
	}
	path := "/vendors/:vendorId/findings"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	if params.SecurityReviewID != nil {
		query.Set("securityReviewId", fmt.Sprint(*params.SecurityReviewID))
	}
	if params.DocumentID != nil {
		query.Set("documentId", fmt.Sprint(*params.DocumentID))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VendorsListVendorFindingsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsListVendorsResponse struct {
	Results map[string]any `json:"results"`
}

type VendorsListVendorsParams struct {
	PageSize         *int
	PageCursor       *string
	Name             *string
	StatusMatchesAny []string
}

// ListVendors List of vendors.
func (s *VendorsService) ListVendors(ctx context.Context, params *VendorsListVendorsParams) (*VendorsListVendorsResponse, error) {
	if params == nil {
		params = &VendorsListVendorsParams{}
	}
	path := "/vendors"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	if params.Name != nil {
		query.Set("name", fmt.Sprint(*params.Name))
	}
	for _, v := range params.StatusMatchesAny {
		query.Add("statusMatchesAny", v)
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VendorsListVendorsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsSetVendorStatusResponse struct {
	AccountManagerEmail              string         `json:"accountManagerEmail"`
	AccountManagerName               string         `json:"accountManagerName"`
	AdditionalNotes                  string         `json:"additionalNotes"`
	AuthDetails                      map[string]any `json:"authDetails"`
	BusinessOwnerUserID              string         `json:"businessOwnerUserId"`
	Category                         map[string]any `json:"category"`
	ContractAmount                   map[string]any `json:"contractAmount"`
	ContractRenewalDate              string         `json:"contractRenewalDate"`
	ContractStartDate                string         `json:"contractStartDate"`
	ContractTerminationDate          any            `json:"contractTerminationDate"`
	CustomFields                     any            `json:"customFields"`
	ID                               string         `json:"id"`
	InherentRiskLevel                string         `json:"inherentRiskLevel"`
	IsRiskAutoScored                 bool           `json:"isRiskAutoScored"`
	IsVisibleToAuditors              bool           `json:"isVisibleToAuditors"`
	LastSecurityReviewCompletionDate string         `json:"lastSecurityReviewCompletionDate"`
	LatestDecision                   map[string]any `json:"latestDecision"`
	Name                             string         `json:"name"`
	NextSecurityReviewDueDate        string         `json:"nextSecurityReviewDueDate"`
	ResidualRiskLevel                string         `json:"residualRiskLevel"`
	RiskAttributeIDs                 []string       `json:"riskAttributeIds"`
	SecurityOwnerUserID              string         `json:"securityOwnerUserId"`
	ServicesProvided                 string         `json:"servicesProvided"`
	Status                           string         `json:"status"`
	TagIDentifiers                   any            `json:"tagIdentifiers"`
	VendorHeadquarters               string         `json:"vendorHeadquarters"`
	WebsiteURL                       string         `json:"websiteUrl"`
}

type VendorsSetVendorStatusParams struct {
	VendorID string
	// FormData maps multipart field names to values.
	FormData map[string]string
}

// SetVendorStatus Sets the status of a vendor, which can be MANAGED, ARCHIVED, or IN_PROCUREMENT.
func (s *VendorsService) SetVendorStatus(ctx context.Context, params *VendorsSetVendorStatusParams) (*VendorsSetVendorStatusResponse, error) {
	if params == nil {
		params = &VendorsSetVendorStatusParams{}
	}
	path := "/vendors/:vendorId/set-status"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	query := url.Values{}
	req, err := s.client.newMultipartRequest(ctx, "POST", path, query, params.FormData)
	if err != nil {
		return nil, err
	}
	out := &VendorsSetVendorStatusResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsUpdateVendorByIDRequestBody struct {
	AccountManagerEmail     string           `json:"accountManagerEmail"`
	AccountManagerName      string           `json:"accountManagerName"`
	AdditionalNotes         string           `json:"additionalNotes"`
	AuthDetails             map[string]any   `json:"authDetails"`
	BusinessOwnerUserID     string           `json:"businessOwnerUserId"`
	Category                string           `json:"category"`
	ContractAmount          map[string]any   `json:"contractAmount"`
	ContractRenewalDate     string           `json:"contractRenewalDate"`
	ContractStartDate       string           `json:"contractStartDate"`
	ContractTerminationDate string           `json:"contractTerminationDate"`
	CustomFields            []map[string]any `json:"customFields"`
	FrameworkScope          map[string]any   `json:"frameworkScope"`
	InherentRiskLevel       string           `json:"inherentRiskLevel"`
	IsVisibleToAuditors     bool             `json:"isVisibleToAuditors"`
	Name                    string           `json:"name"`
	ResidualRiskLevel       string           `json:"residualRiskLevel"`
	RiskAttributeIDs        []string         `json:"riskAttributeIds"`
	SecurityOwnerUserID     string           `json:"securityOwnerUserId"`
	ServicesProvided        string           `json:"servicesProvided"`
	Status                  string           `json:"status"`
	VendorHeadquarters      string           `json:"vendorHeadquarters"`
	WebsiteURL              string           `json:"websiteUrl"`
}

type VendorsUpdateVendorByIDResponse struct {
	AccountManagerEmail              string         `json:"accountManagerEmail"`
	AccountManagerName               string         `json:"accountManagerName"`
	AdditionalNotes                  string         `json:"additionalNotes"`
	AuthDetails                      map[string]any `json:"authDetails"`
	BusinessOwnerUserID              string         `json:"businessOwnerUserId"`
	Category                         map[string]any `json:"category"`
	ContractAmount                   map[string]any `json:"contractAmount"`
	ContractRenewalDate              string         `json:"contractRenewalDate"`
	ContractStartDate                string         `json:"contractStartDate"`
	ContractTerminationDate          any            `json:"contractTerminationDate"`
	CustomFields                     any            `json:"customFields"`
	ID                               string         `json:"id"`
	InherentRiskLevel                string         `json:"inherentRiskLevel"`
	IsRiskAutoScored                 bool           `json:"isRiskAutoScored"`
	IsVisibleToAuditors              bool           `json:"isVisibleToAuditors"`
	LastSecurityReviewCompletionDate string         `json:"lastSecurityReviewCompletionDate"`
	LatestDecision                   map[string]any `json:"latestDecision"`
	Name                             string         `json:"name"`
	NextSecurityReviewDueDate        string         `json:"nextSecurityReviewDueDate"`
	ResidualRiskLevel                string         `json:"residualRiskLevel"`
	RiskAttributeIDs                 []string       `json:"riskAttributeIds"`
	SecurityOwnerUserID              string         `json:"securityOwnerUserId"`
	ServicesProvided                 string         `json:"servicesProvided"`
	Status                           string         `json:"status"`
	TagIDentifiers                   any            `json:"tagIdentifiers"`
	VendorHeadquarters               string         `json:"vendorHeadquarters"`
	WebsiteURL                       string         `json:"websiteUrl"`
}

type VendorsUpdateVendorByIDParams struct {
	VendorID string
	Body     *VendorsUpdateVendorByIDRequestBody
}

// UpdateVendorByID Update vendor.
func (s *VendorsService) UpdateVendorByID(ctx context.Context, params *VendorsUpdateVendorByIDParams) (*VendorsUpdateVendorByIDResponse, error) {
	if params == nil {
		params = &VendorsUpdateVendorByIDParams{}
	}
	path := "/vendors/:vendorId"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &VendorsUpdateVendorByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VendorsUpdateVendorFindingRequestBody struct {
	Content     string         `json:"content"`
	Remediation map[string]any `json:"remediation"`
	RiskStatus  string         `json:"riskStatus"`
}

type VendorsUpdateVendorFindingResponse struct {
	Content          string         `json:"content"`
	DocumentID       any            `json:"documentId"`
	ID               string         `json:"id"`
	Remediation      map[string]any `json:"remediation"`
	RiskStatus       string         `json:"riskStatus"`
	SecurityReviewID string         `json:"securityReviewId"`
	VendorID         string         `json:"vendorId"`
}

type VendorsUpdateVendorFindingParams struct {
	VendorID  string
	FindingID string
	Body      *VendorsUpdateVendorFindingRequestBody
}

// UpdateVendorFinding Update vendor finding.
func (s *VendorsService) UpdateVendorFinding(ctx context.Context, params *VendorsUpdateVendorFindingParams) (*VendorsUpdateVendorFindingResponse, error) {
	if params == nil {
		params = &VendorsUpdateVendorFindingParams{}
	}
	path := "/vendors/:vendorId/findings/:findingId"
	if params.VendorID == "" {
		return nil, fmt.Errorf("vendorId is required")
	}
	path = strings.ReplaceAll(path, ":vendorId", url.PathEscape(params.VendorID))
	if params.FindingID == "" {
		return nil, fmt.Errorf("findingId is required")
	}
	path = strings.ReplaceAll(path, ":findingId", url.PathEscape(params.FindingID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "PATCH", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &VendorsUpdateVendorFindingResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// VulnerabilitiesService groups 4 endpoint methods under the "Vulnerabilities" API segment.
type VulnerabilitiesService struct {
	client *Client
}

type VulnerabilitiesDeactivateVulnerabilityMonitoringForVulnerabilityRequestBody struct {
	Updates []map[string]any `json:"updates"`
}

type VulnerabilitiesDeactivateVulnerabilityMonitoringForVulnerabilityResponse struct {
	Results []map[string]any `json:"results"`
}

type VulnerabilitiesDeactivateVulnerabilityMonitoringForVulnerabilityParams struct {
	Body *VulnerabilitiesDeactivateVulnerabilityMonitoringForVulnerabilityRequestBody
}

// DeactivateVulnerabilityMonitoringForVulnerability Deactivate monitoring for select vulnerabilities. Vanta will not monitor a deactivated vulnerability until it is reactivated.
func (s *VulnerabilitiesService) DeactivateVulnerabilityMonitoringForVulnerability(ctx context.Context, params *VulnerabilitiesDeactivateVulnerabilityMonitoringForVulnerabilityParams) (*VulnerabilitiesDeactivateVulnerabilityMonitoringForVulnerabilityResponse, error) {
	if params == nil {
		params = &VulnerabilitiesDeactivateVulnerabilityMonitoringForVulnerabilityParams{}
	}
	path := "/vulnerabilities/deactivate"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &VulnerabilitiesDeactivateVulnerabilityMonitoringForVulnerabilityResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VulnerabilitiesGetVulnerabilitiesResponse struct {
	Results map[string]any `json:"results"`
}

type VulnerabilitiesGetVulnerabilitiesParams struct {
	Q                                 *string
	PageSize                          *int
	PageCursor                        *string
	IsDeactivated                     *bool
	ExternalVulnerabilityID           *string
	IsFixAvailable                    *bool
	PackageIDentifier                 *string
	SlaDeadlineAfterDate              *string
	SlaDeadlineBeforeDate             *string
	Severity                          *string
	IntegrationID                     *string
	IncludeVulnerabilitiesWithoutSlas *bool
	VulnerableAssetID                 *string
}

// GetVulnerabilities List all vulnerabilities based on selected filters.
func (s *VulnerabilitiesService) GetVulnerabilities(ctx context.Context, params *VulnerabilitiesGetVulnerabilitiesParams) (*VulnerabilitiesGetVulnerabilitiesResponse, error) {
	if params == nil {
		params = &VulnerabilitiesGetVulnerabilitiesParams{}
	}
	path := "/vulnerabilities"
	query := url.Values{}
	if params.Q != nil {
		query.Set("q", fmt.Sprint(*params.Q))
	}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	if params.IsDeactivated != nil {
		query.Set("isDeactivated", fmt.Sprint(*params.IsDeactivated))
	}
	if params.ExternalVulnerabilityID != nil {
		query.Set("externalVulnerabilityId", fmt.Sprint(*params.ExternalVulnerabilityID))
	}
	if params.IsFixAvailable != nil {
		query.Set("isFixAvailable", fmt.Sprint(*params.IsFixAvailable))
	}
	if params.PackageIDentifier != nil {
		query.Set("packageIdentifier", fmt.Sprint(*params.PackageIDentifier))
	}
	if params.SlaDeadlineAfterDate != nil {
		query.Set("slaDeadlineAfterDate", fmt.Sprint(*params.SlaDeadlineAfterDate))
	}
	if params.SlaDeadlineBeforeDate != nil {
		query.Set("slaDeadlineBeforeDate", fmt.Sprint(*params.SlaDeadlineBeforeDate))
	}
	if params.Severity != nil {
		query.Set("severity", fmt.Sprint(*params.Severity))
	}
	if params.IntegrationID != nil {
		query.Set("integrationId", fmt.Sprint(*params.IntegrationID))
	}
	if params.IncludeVulnerabilitiesWithoutSlas != nil {
		query.Set("includeVulnerabilitiesWithoutSlas", fmt.Sprint(*params.IncludeVulnerabilitiesWithoutSlas))
	}
	if params.VulnerableAssetID != nil {
		query.Set("vulnerableAssetId", fmt.Sprint(*params.VulnerableAssetID))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VulnerabilitiesGetVulnerabilitiesResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VulnerabilitiesGetVulnerabilityByIDResponse struct {
	CvssSeverityScore  float64        `json:"cvssSeverityScore"`
	DeactivateMetadata map[string]any `json:"deactivateMetadata"`
	Description        string         `json:"description"`
	ExternalURL        string         `json:"externalURL"`
	FirstDetectedDate  string         `json:"firstDetectedDate"`
	ID                 string         `json:"id"`
	IntegrationID      string         `json:"integrationId"`
	IsFixable          bool           `json:"isFixable"`
	LastDetectedDate   string         `json:"lastDetectedDate"`
	Name               string         `json:"name"`
	PackageIDentifier  string         `json:"packageIdentifier"`
	RelatedURLs        []string       `json:"relatedUrls"`
	RelatedVulns       []string       `json:"relatedVulns"`
	RemediateByDate    string         `json:"remediateByDate"`
	ScanSource         string         `json:"scanSource"`
	ScannerScore       float64        `json:"scannerScore"`
	Severity           string         `json:"severity"`
	SourceDetectedDate string         `json:"sourceDetectedDate"`
	TargetID           string         `json:"targetId"`
	VulnerabilityType  string         `json:"vulnerabilityType"`
}

type VulnerabilitiesGetVulnerabilityByIDParams struct {
	VulnerabilityID string
}

// GetVulnerabilityByID Gets a vulnerability by an ID.
func (s *VulnerabilitiesService) GetVulnerabilityByID(ctx context.Context, params *VulnerabilitiesGetVulnerabilityByIDParams) (*VulnerabilitiesGetVulnerabilityByIDResponse, error) {
	if params == nil {
		params = &VulnerabilitiesGetVulnerabilityByIDParams{}
	}
	path := "/vulnerabilities/:vulnerabilityId"
	if params.VulnerabilityID == "" {
		return nil, fmt.Errorf("vulnerabilityId is required")
	}
	path = strings.ReplaceAll(path, ":vulnerabilityId", url.PathEscape(params.VulnerabilityID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VulnerabilitiesGetVulnerabilityByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VulnerabilitiesReactivateVulnerabilityMonitoringRequestBody struct {
	Updates []map[string]any `json:"updates"`
}

type VulnerabilitiesReactivateVulnerabilityMonitoringResponse struct {
	Results []map[string]any `json:"results"`
}

type VulnerabilitiesReactivateVulnerabilityMonitoringParams struct {
	Body *VulnerabilitiesReactivateVulnerabilityMonitoringRequestBody
}

// ReactivateVulnerabilityMonitoring Reactivate vulnerabilities and resume Vanta monitoring.
func (s *VulnerabilitiesService) ReactivateVulnerabilityMonitoring(ctx context.Context, params *VulnerabilitiesReactivateVulnerabilityMonitoringParams) (*VulnerabilitiesReactivateVulnerabilityMonitoringResponse, error) {
	if params == nil {
		params = &VulnerabilitiesReactivateVulnerabilityMonitoringParams{}
	}
	path := "/vulnerabilities/reactivate"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &VulnerabilitiesReactivateVulnerabilityMonitoringResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// VulnerabilityRemediationsService groups 2 endpoint methods under the "VulnerabilityRemediations" API segment.
type VulnerabilityRemediationsService struct {
	client *Client
}

type VulnerabilityRemediationsAcknowledgeSlaMissRequestBody struct {
	Updates []map[string]any `json:"updates"`
}

type VulnerabilityRemediationsAcknowledgeSlaMissResponse struct {
	Results []map[string]any `json:"results"`
}

type VulnerabilityRemediationsAcknowledgeSlaMissParams struct {
	Body *VulnerabilityRemediationsAcknowledgeSlaMissRequestBody
}

// AcknowledgeSlaMiss Acknowledge an SLA miss for a vulnerability remediation.
func (s *VulnerabilityRemediationsService) AcknowledgeSlaMiss(ctx context.Context, params *VulnerabilityRemediationsAcknowledgeSlaMissParams) (*VulnerabilityRemediationsAcknowledgeSlaMissResponse, error) {
	if params == nil {
		params = &VulnerabilityRemediationsAcknowledgeSlaMissParams{}
	}
	path := "/vulnerability-remediations/acknowledge-sla-miss"
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "POST", path, query, params.Body)
	if err != nil {
		return nil, err
	}
	out := &VulnerabilityRemediationsAcknowledgeSlaMissResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VulnerabilityRemediationsListVulnerabilityRemediationsResponse struct {
	Results map[string]any `json:"results"`
}

type VulnerabilityRemediationsListVulnerabilityRemediationsParams struct {
	PageSize             *int
	PageCursor           *string
	IntegrationID        *string
	Severity             *string
	IsRemediatedOnTime   *bool
	RemediatedAfterDate  *string
	RemediatedBeforeDate *string
}

// ListVulnerabilityRemediations List all vulnerability remediations based on selected filters.
func (s *VulnerabilityRemediationsService) ListVulnerabilityRemediations(ctx context.Context, params *VulnerabilityRemediationsListVulnerabilityRemediationsParams) (*VulnerabilityRemediationsListVulnerabilityRemediationsResponse, error) {
	if params == nil {
		params = &VulnerabilityRemediationsListVulnerabilityRemediationsParams{}
	}
	path := "/vulnerability-remediations"
	query := url.Values{}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	if params.IntegrationID != nil {
		query.Set("integrationId", fmt.Sprint(*params.IntegrationID))
	}
	if params.Severity != nil {
		query.Set("severity", fmt.Sprint(*params.Severity))
	}
	if params.IsRemediatedOnTime != nil {
		query.Set("isRemediatedOnTime", fmt.Sprint(*params.IsRemediatedOnTime))
	}
	if params.RemediatedAfterDate != nil {
		query.Set("remediatedAfterDate", fmt.Sprint(*params.RemediatedAfterDate))
	}
	if params.RemediatedBeforeDate != nil {
		query.Set("remediatedBeforeDate", fmt.Sprint(*params.RemediatedBeforeDate))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VulnerabilityRemediationsListVulnerabilityRemediationsResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

// VulnerableAssetsService groups 2 endpoint methods under the "VulnerableAssets" API segment.
type VulnerableAssetsService struct {
	client *Client
}

type VulnerableAssetsGetVulnerableAssetByIDResponse struct {
	AssetType      string           `json:"assetType"`
	HasBeenScanned bool             `json:"hasBeenScanned"`
	ID             string           `json:"id"`
	ImageScanTag   string           `json:"imageScanTag"`
	Name           string           `json:"name"`
	Scanners       []map[string]any `json:"scanners"`
}

type VulnerableAssetsGetVulnerableAssetByIDParams struct {
	VulnerableAssetID string
}

// GetVulnerableAssetByID Gets a vulnerable asset by ID.
func (s *VulnerableAssetsService) GetVulnerableAssetByID(ctx context.Context, params *VulnerableAssetsGetVulnerableAssetByIDParams) (*VulnerableAssetsGetVulnerableAssetByIDResponse, error) {
	if params == nil {
		params = &VulnerableAssetsGetVulnerableAssetByIDParams{}
	}
	path := "/vulnerable-assets/:vulnerableAssetId"
	if params.VulnerableAssetID == "" {
		return nil, fmt.Errorf("vulnerableAssetId is required")
	}
	path = strings.ReplaceAll(path, ":vulnerableAssetId", url.PathEscape(params.VulnerableAssetID))
	query := url.Values{}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VulnerableAssetsGetVulnerableAssetByIDResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type VulnerableAssetsListAssetsAssociatedWithVulnerabilitiesResponse struct {
	Results map[string]any `json:"results"`
}

type VulnerableAssetsListAssetsAssociatedWithVulnerabilitiesParams struct {
	Q                      *string
	PageSize               *int
	PageCursor             *string
	IntegrationID          *string
	AssetType              *string
	AssetExternalAccountID *string
}

// ListAssetsAssociatedWithVulnerabilities List assets that Vanta monitors that are associated with vulnerabilities.
func (s *VulnerableAssetsService) ListAssetsAssociatedWithVulnerabilities(ctx context.Context, params *VulnerableAssetsListAssetsAssociatedWithVulnerabilitiesParams) (*VulnerableAssetsListAssetsAssociatedWithVulnerabilitiesResponse, error) {
	if params == nil {
		params = &VulnerableAssetsListAssetsAssociatedWithVulnerabilitiesParams{}
	}
	path := "/vulnerable-assets"
	query := url.Values{}
	if params.Q != nil {
		query.Set("q", fmt.Sprint(*params.Q))
	}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprint(*params.PageSize))
	}
	if params.PageCursor != nil {
		query.Set("pageCursor", fmt.Sprint(*params.PageCursor))
	}
	if params.IntegrationID != nil {
		query.Set("integrationId", fmt.Sprint(*params.IntegrationID))
	}
	if params.AssetType != nil {
		query.Set("assetType", fmt.Sprint(*params.AssetType))
	}
	if params.AssetExternalAccountID != nil {
		query.Set("assetExternalAccountId", fmt.Sprint(*params.AssetExternalAccountID))
	}
	req, err := s.client.newRequest(ctx, "GET", path, query, nil)
	if err != nil {
		return nil, err
	}
	out := &VulnerableAssetsListAssetsAssociatedWithVulnerabilitiesResponse{}
	if err := s.client.doJSON(req, out); err != nil {
		return nil, err
	}
	return out, nil
}
