package legacy

const (
	marketingEmailsBasePath = "marketing-emails/v1/emails"
)

type (
	ChildCSS = map[string]interface{}
	CSS      = map[string]interface{}
)

type BlogRssSettings struct {
	BlogLayout string `json:"blog_layout,omitempty"`
}

type RssToEmailTiming struct {
	Repeats          string `json:"repeats,omitempty"`
	RepeatsOnMonthly int64  `json:"repeats_on_monthly,omitempty"`
	RepeatsOnWeekly  int64  `json:"repeats_on_weekly,omitempty"`
	Summary          string `json:"summary,omitempty"`
	Time             string `json:"time,omitempty"`
}

type (
	SmartEmailFields struct{}
	Counters         struct {
		Bounce       int64 `json:"bounce,omitempty"`
		Click        int64 `json:"click,omitempty"`
		Contactslost int64 `json:"contactslost,omitempty"`
		Delivered    int64 `json:"delivered,omitempty"`
		Dropped      int64 `json:"dropped,omitempty"`
		HardBounced  int64 `json:"hardbounced,omitempty"`
		NotSent      int64 `json:"notsent,omitempty"`
		Open         int64 `json:"open,omitempty"`
		Pending      int64 `json:"pending,omitempty"`
		Selected     int64 `json:"selected,omitempty"`
		Sent         int64 `json:"sent,omitempty"`
		SoftBounced  int64 `json:"softbounced,omitempty"`
		SpamReport   int64 `json:"spamreport,omitempty"`
		Suppressed   int64 `json:"suppressed,omitempty"`
		Unsubscribed int64 `json:"unsubscribed,omitempty"`
	}
)

type ClickDeviceType struct {
	Computer int64 `json:"computer,omitempty"`
	Mobile   int64 `json:"mobile,omitempty"`
	Unknown  int64 `json:"unknown,omitempty"`
}

type OpenDeviceType struct {
	Computer int64 `json:"computer,omitempty"`
	Mobile   int64 `json:"mobile,omitempty"`
	Unknown  int64 `json:"unknown,omitempty"`
}

type DeviceBreakdown struct {
	ClickDeviceType *ClickDeviceType `json:"click_device_type,omitempty"`
	OpenDeviceType  *OpenDeviceType  `json:"open_device_type,omitempty"`
}

type Ratios struct {
	BounceRatio       float64 `json:"bounceratio,omitempty"`
	ClickRatio        float64 `json:"clickratio,omitempty"`
	ClickThroughRatio float64 `json:"clickthroughratio,omitempty"`
	ContactSlostRatio float64 `json:"contactslostratio,omitempty"`
	DeliveredRatio    float64 `json:"deliveredratio,omitempty"`
	HardBounceRatio   float64 `json:"hardbounceratio,omitempty"`
	NotSentRatio      float64 `json:"notsentratio,omitempty"`
	OpenRatio         float64 `json:"openratio,omitempty"`
	PendingRatio      float64 `json:"pendingratio,omitempty"`
	SoftBounceRatio   float64 `json:"softbounceratio,omitempty"`
	SpamReportRatio   float64 `json:"spamreportratio,omitempty"`
	UnsubscribedRatio float64 `json:"unsubscribedratio,omitempty"`
}

type Stats struct {
	Counters        *Counters              `json:"counters,omitempty"`
	DeviceBreakdown *DeviceBreakdown       `json:"deviceBreakdown,omitempty"`
	FailedToLoad    bool                   `json:"failedToLoad,omitempty"`
	QualifierStats  map[string]interface{} `json:"qualifierStats,omitempty"`
	Ratios          *Ratios                `json:"ratios,omitempty"`
}

type Body struct {
	HTML string `json:"html,omitempty"`
}

type Column1 struct {
	Body      *Body       `json:"body,omitempty"`
	ChildCSS  *ChildCSS   `json:"child_css,omitempty"`
	CSS       *CSS        `json:"css,omitempty"`
	ID        string      `json:"id,omitempty"`
	Label     string      `json:"label,omitempty"`
	Name      string      `json:"name,omitempty"`
	SmartType interface{} `json:"smart_type,omitempty"`
	Type      string      `json:"type,omitempty"`
}

type HsEmailBody struct {
	Body      *Body       `json:"body,omitempty"`
	ChildCSS  *ChildCSS   `json:"child_css,omitempty"`
	CSS       *CSS        `json:"css,omitempty"`
	ID        string      `json:"id,omitempty"`
	Label     string      `json:"label,omitempty"`
	Name      string      `json:"name,omitempty"`
	SmartType interface{} `json:"smart_type,omitempty"`
	Type      string      `json:"type,omitempty"`
}

type LogoImage struct {
	Body      *Body       `json:"body,omitempty"`
	ChildCSS  *ChildCSS   `json:"child_css,omitempty"`
	CSS       *CSS        `json:"css,omitempty"`
	ID        string      `json:"id,omitempty"`
	Label     string      `json:"label,omitempty"`
	Name      string      `json:"name,omitempty"`
	SmartType interface{} `json:"smart_type,omitempty"`
	Type      string      `json:"type,omitempty"`
}

type PreviewText struct {
	Body      *Body       `json:"body,omitempty"`
	ChildCSS  *ChildCSS   `json:"child_css,omitempty"`
	CSS       *CSS        `json:"css,omitempty"`
	ID        string      `json:"id,omitempty"`
	Label     string      `json:"label,omitempty"`
	Name      string      `json:"name,omitempty"`
	SmartType interface{} `json:"smart_type,omitempty"`
	Type      string      `json:"type,omitempty"`
}

type Widgets struct {
	Column1     *Column1     `json:"column1,omitempty"`
	HsEmailBody *HsEmailBody `json:"hs_email_body,omitempty"`
	LogoImage   *LogoImage   `json:"logo_image,omitempty"`
	PreviewText *PreviewText `json:"preview_text,omitempty"`
}

type StatisticsResponse struct {
	Ab                                   bool                   `json:"ab,omitempty"`
	AbHoursToWait                        int64                  `json:"abHoursToWait,omitempty"`
	AbSampleSizeDefault                  interface{}            `json:"abSampleSizeDefault,omitempty"`
	AbSamplingDefault                    interface{}            `json:"abSamplingDefault,omitempty"`
	AbSuccessMetric                      interface{}            `json:"abSuccessMetric,omitempty"`
	AbTestPercentage                     int64                  `json:"abTestPercentage,omitempty"`
	AbVariation                          bool                   `json:"abVariation,omitempty"`
	AbsoluteURL                          string                 `json:"absoluteURL,omitempty"`
	AllEmailCampaignIDs                  []int64                `json:"allEmailCampaignIDs,omitempty"`
	AnalyticsPageID                      string                 `json:"analyticsPageID,omitempty"`
	AnalyticsPageType                    string                 `json:"analyticsPageType,omitempty"`
	Archived                             bool                   `json:"archived,omitempty"`
	AuthorAt                             int64                  `json:"authorAt,omitempty"`
	AuthorName                           string                 `json:"authorName,omitempty"`
	AuthorUserID                         int64                  `json:"authorUserID,omitempty"`
	BlogEmailType                        string                 `json:"blogEmailType,omitempty"`
	BlogRssSettings                      *BlogRssSettings       `json:"blogRssSettings,omitempty"`
	CanSpamSettingsID                    int64                  `json:"canSpamSettingsID,omitempty"`
	CategoryID                           int64                  `json:"categoryID,omitempty"`
	ContentTypeCategory                  int64                  `json:"contentTypeCategory,omitempty"`
	CreatePage                           bool                   `json:"createPage,omitempty"`
	Created                              int64                  `json:"created,omitempty"`
	CreatedByID                          int64                  `json:"createdByID,omitempty"`
	CurrentState                         string                 `json:"currentState,omitempty"`
	CurrentlyPublished                   bool                   `json:"currentlyPublished,omitempty"`
	CustomReplyTo                        string                 `json:"customReplyTo,omitempty"`
	CustomReplyToEnabled                 bool                   `json:"customReplyToEnabled,omitempty"`
	Domain                               string                 `json:"domain,omitempty"`
	EmailBody                            string                 `json:"emailBody,omitempty"`
	EmailNote                            string                 `json:"emailNote,omitempty"`
	EmailTemplateMode                    string                 `json:"emailTemplateMode,omitempty"`
	EmailType                            string                 `json:"emailType,omitempty"`
	EmailbodyPlaintext                   string                 `json:"emailbodyPlaintext,omitempty"`
	FeedbackEmailCategory                interface{}            `json:"feedbackEmailCategory,omitempty"`
	FeedbackSurveyID                     interface{}            `json:"feedbackSurveyID,omitempty"`
	FlexAreas                            map[string]interface{} `json:"flexAreas,omitempty"`
	FreezeDate                           int64                  `json:"freezeDate,omitempty"`
	FromName                             string                 `json:"fromName,omitempty"`
	HsEmailBody                          string                 `json:"hsEmailBody,omitempty"`
	HTMLTitle                            string                 `json:"htmlTitle,omitempty"`
	ID                                   int64                  `json:"id,omitempty"`
	IsGraymailSuppressionEnabled         bool                   `json:"isGraymailSuppressionEnabled,omitempty"`
	IsLocalTimezoneSend                  bool                   `json:"isLocalTimezoneSend,omitempty"`
	IsPublished                          bool                   `json:"isPublished,omitempty"`
	IsRecipientFatigueSuppressionEnabled interface{}            `json:"isRecipientFatigueSuppressionEnabled,omitempty"`
	LayoutSections                       map[string]interface{} `json:"layoutSections,omitempty"`
	LeadFlowID                           interface{}            `json:"leadFlowID,omitempty"`
	LiveDomain                           string                 `json:"liveDomain,omitempty"`
	MailingListsExcluded                 []interface{}          `json:"mailingListsExcluded,omitempty"`
	MailingListsIncluded                 []int64                `json:"mailingListsIncluded,omitempty"`
	MaxRssEntries                        int64                  `json:"maxRssEntries,omitempty"`
	MetaDescription                      string                 `json:"metaDescription,omitempty"`
	Name                                 string                 `json:"name,omitempty"`
	PageExpiryEnabled                    bool                   `json:"pageExpiryEnabled,omitempty"`
	PageRedirected                       bool                   `json:"pageRedirected,omitempty"`
	PastMabExperimentIDs                 []interface{}          `json:"pastMabExperimentIDs,omitempty"`
	PortalID                             int64                  `json:"portalID,omitempty"`
	PreviewKey                           string                 `json:"previewKey,omitempty"`
	PrimaryEmailCampaignID               int64                  `json:"primaryEmailCampaignID,omitempty"`
	PrimaryRichTextModuleHTML            string                 `json:"primaryRichTextModuleHTML,omitempty"`
	ProcessingStatus                     string                 `json:"processingStatus,omitempty"`
	PublishDate                          int64                  `json:"publishDate,omitempty"`
	PublishedAt                          int64                  `json:"publishedAt,omitempty"`
	PublishedByID                        int64                  `json:"publishedByID,omitempty"`
	PublishedURL                         string                 `json:"publishedURL,omitempty"`
	ReplyTo                              string                 `json:"replyTo,omitempty"`
	ResolvedDomain                       string                 `json:"resolvedDomain,omitempty"`
	RssEmailAuthorLineTemplate           string                 `json:"rssEmailAuthorLineTemplate,omitempty"`
	RssEmailBlogImageMaxWidth            int64                  `json:"rssEmailBlogImageMaxWidth,omitempty"`
	RssEmailByText                       string                 `json:"rssEmailByText,omitempty"`
	RssEmailClickThroughText             string                 `json:"rssEmailClickThroughText,omitempty"`
	RssEmailCommentText                  string                 `json:"rssEmailCommentText,omitempty"`
	RssEmailEntryTemplate                string                 `json:"rssEmailEntryTemplate,omitempty"`
	RssEmailEntryTemplateEnabled         bool                   `json:"rssEmailEntryTemplateEnabled,omitempty"`
	RssEmailImageMaxWidth                int64                  `json:"rssEmailImageMaxWidth,omitempty"`
	RssEmailURL                          string                 `json:"rssEmailURL,omitempty"`
	RssToEmailTiming                     *RssToEmailTiming      `json:"rssToEmailTiming,omitempty"`
	RssURL                               string                 `json:"rssURL,omitempty"`
	Selected                             int64                  `json:"selected,omitempty"`
	Slug                                 string                 `json:"slug,omitempty"`
	SmartEmailFields                     *SmartEmailFields      `json:"smartEmailFields,omitempty"`
	State                                string                 `json:"state,omitempty"`
	Stats                                *Stats                 `json:"stats,omitempty"`
	StyleSettings                        interface{}            `json:"styleSettings,omitempty"`
	Subcategory                          string                 `json:"subcategory,omitempty"`
	Subject                              string                 `json:"subject,omitempty"`
	Subscription                         int64                  `json:"subscription,omitempty"`
	SubscriptionBlogID                   int64                  `json:"subscriptionBlogID,omitempty"`
	SubscriptionName                     string                 `json:"subscriptionName,omitempty"`
	TeamPerms                            []interface{}          `json:"teamPerms,omitempty"`
	TemplatePath                         string                 `json:"templatePath,omitempty"`
	Transactional                        bool                   `json:"transactional,omitempty"`
	UnpublishedAt                        int64                  `json:"unpublishedAt,omitempty"`
	Updated                              int64                  `json:"updated,omitempty"`
	UpdatedByID                          int64                  `json:"updatedByID,omitempty"`
	URL                                  string                 `json:"url,omitempty"`
	UseRssHeadlineAsSubject              bool                   `json:"useRssHeadlineAsSubject,omitempty"`
	UserPerms                            []interface{}          `json:"userPerms,omitempty"`
	VidsExcluded                         []interface{}          `json:"vidsExcluded,omitempty"`
	VidsIncluded                         []interface{}          `json:"vidsIncluded,omitempty"`
	Widgets                              *Widgets               `json:"widgets,omitempty"`
}

// MarketingEmailHelper is a helper to handle MarketingEmail API v1.
type MarketingEmailHelper interface {
	GetStatisticsPath() string
}

type marketingEmailHelperOp struct {
	marketingEmailsBasePath string
}

func (m *marketingEmailHelperOp) GetStatisticsPath() string {
	return m.marketingEmailsBasePath + "/with-statistics"
}

var _ MarketingEmailHelper = (*marketingEmailHelperOp)(nil)

func NewMarketingEmailHelper() MarketingEmailHelper {
	return &marketingEmailHelperOp{marketingEmailsBasePath: marketingEmailsBasePath}
}
