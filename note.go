package hubspot

const (
	noteBasePath = "notes"
)

// NoteService is an interface of note endpoints of the HubSpot API.
// HubSpot notes store information about interactions and communications.
// It can be associated with other CRM objects such as contact, company, and deal.
// Reference: https://developers.hubspot.com/docs/api/crm/notes
type NoteService interface {
	Get(noteID string, note interface{}, option *RequestQueryOption) (*ResponseResource, error)
	Create(note interface{}) (*ResponseResource, error)
	Update(noteID string, note interface{}) (*ResponseResource, error)
	Delete(noteID string) error
	AssociateAnotherObj(noteID string, conf *AssociationConfig) (*ResponseResource, error)
}

// NoteServiceOp handles communication with the note related methods of the HubSpot API.
type NoteServiceOp struct {
	notePath string
	client   *Client
}

var _ NoteService = (*NoteServiceOp)(nil)

// Note represents a note in HubSpot.
type Note struct {
	HsCreateDate   *HsTime `json:"hs_createdate,omitempty"`
	HsObjectID     *HsStr  `json:"hs_object_id,omitempty"`
	HsNoteBody     *HsStr  `json:"hs_note_body,omitempty"`
	HsNoteTitle    *HsStr  `json:"hs_note_title,omitempty"`
	HsTimestamp    *HsStr  `json:"hs_timestamp,omitempty"`
	HubspotOwnerID *HsStr  `json:"hubspot_owner_id,omitempty"`
}

var defaultNoteFields = []string{
	"hs_all_accessible_team_ids",
	"hs_all_assigned_business_unit_ids",
	"hs_all_owner_ids",
	"hs_all_team_ids",
	"hs_at_mentioned_owner_ids",
	"hs_attachment_ids",
	"hs_body_preview",
	"hs_body_preview_html",
	"hs_body_preview_is_truncated",
	"hs_created_by",
	"hs_created_by_user_id",
	"hs_createdate",
	"hs_engagement_source",
	"hs_engagement_source_id",
	"hs_follow_up_action",
	"hs_gdpr_deleted",
	"hs_hd_ticket_ids",
	"hs_lastmodifieddate",
	"hs_merged_object_ids",
	"hs_modified_by",
	"hs_note_body",
	"hs_note_ms_teams_payload",
	"hs_object_id",
	"hs_object_source",
	"hs_object_source_detail_1",
	"hs_object_source_detail_2",
	"hs_object_source_detail_3",
	"hs_object_source_id",
	"hs_object_source_label",
	"hs_object_source_user_id",
	"hs_product_name",
	"hs_queue_membership_ids",
	"hs_read_only",
	"hs_shared_team_ids",
	"hs_shared_user_ids",
	"hs_timestamp",
	"hs_unique_creation_key",
	"hs_unique_id",
	"hs_updated_by_user_id",
	"hs_user_ids_of_all_notification_followers",
	"hs_user_ids_of_all_notification_unfollowers",
	"hs_user_ids_of_all_owners",
	"hs_was_imported",
	"hubspot_owner_assigneddate",
	"hubspot_owner_id",
	"hubspot_team_id",
}

// Get gets a note.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *NoteServiceOp) Get(noteID string, note interface{}, option *RequestQueryOption) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: note}
	if err := s.client.Get(s.notePath+"/"+noteID, resource, option.setupProperties(defaultNoteFields)); err != nil {
		return nil, err
	}
	return resource, nil
}

// Create creates a new note.
// In order to bind the created content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Note in your own structure.
func (s *NoteServiceOp) Create(note interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: note}
	resource := &ResponseResource{Properties: note}
	if err := s.client.Post(s.notePath, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Update updates a note.
// In order to bind the updated content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Note in your own structure.
func (s *NoteServiceOp) Update(noteID string, note interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: note}
	resource := &ResponseResource{Properties: note}
	if err := s.client.Patch(s.notePath+"/"+noteID, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Delete deletes a note.
func (s *NoteServiceOp) Delete(noteID string) error {
	return s.client.Delete(s.notePath+"/"+noteID, nil)
}

// AssociateAnotherObj associates Note with another HubSpot objects.
// If you want to associate a custom object, please use a defined value in HubSpot.
func (s *NoteServiceOp) AssociateAnotherObj(noteID string, conf *AssociationConfig) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: &Note{}}
	if err := s.client.Put(s.notePath+"/"+noteID+"/"+conf.makeAssociationPath(), nil, resource); err != nil {
		return nil, err
	}
	return resource, nil
}
