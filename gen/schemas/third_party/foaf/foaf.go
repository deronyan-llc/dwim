package generated

// OnlineChatAccount is a generated struct representing the http://xmlns.com/foaf/0.1/OnlineChatAccount class.
type OnlineChatAccount struct {
}

// OnlineGamingAccount is a generated struct representing the http://xmlns.com/foaf/0.1/OnlineGamingAccount class.
type OnlineGamingAccount struct {
}

// Person is a generated struct representing the http://xmlns.com/foaf/0.1/Person class.
type Person struct {
	CurrentProject    *Thing    `json:"currentproject"`    // //http://www.w3.org/2002/07/owl#Thing
	FamilyName        *Literal  `json:"familyname"`        // //http://www.w3.org/2000/01/rdf-schema#Literal
	Familyname        *Literal  `json:"familyname"`        // //http://www.w3.org/2000/01/rdf-schema#Literal
	FirstName         *Literal  `json:"firstname"`         // //http://www.w3.org/2000/01/rdf-schema#Literal
	Geekcode          *Literal  `json:"geekcode"`          // //http://www.w3.org/2000/01/rdf-schema#Literal
	Img               *Image    `json:"img"`               // //http://xmlns.com/foaf/0.1/Image
	Knows             *Person   `json:"knows"`             // //http://xmlns.com/foaf/0.1/Person
	LastName          *Literal  `json:"lastname"`          // //http://www.w3.org/2000/01/rdf-schema#Literal
	MyersBriggs       *Literal  `json:"myersbriggs"`       // //http://www.w3.org/2000/01/rdf-schema#Literal
	PastProject       *Thing    `json:"pastproject"`       // //http://www.w3.org/2002/07/owl#Thing
	Plan              *Literal  `json:"plan"`              // //http://www.w3.org/2000/01/rdf-schema#Literal
	Publications      *Document `json:"publications"`      // //http://xmlns.com/foaf/0.1/Document
	SchoolHomepage    *Document `json:"schoolhomepage"`    // //http://xmlns.com/foaf/0.1/Document
	Surname           *Literal  `json:"surname"`           // //http://www.w3.org/2000/01/rdf-schema#Literal
	WorkInfoHomepage  *Document `json:"workinfohomepage"`  // //http://xmlns.com/foaf/0.1/Document
	WorkplaceHomepage *Document `json:"workplacehomepage"` // //http://xmlns.com/foaf/0.1/Document
}

// Document is a generated struct representing the http://xmlns.com/foaf/0.1/Document class.
type Document struct {
	PrimaryTopic *Thing  `json:"primarytopic"` // //http://www.w3.org/2002/07/owl#Thing
	Sha1         *string `json:"sha1"`         //
	Topic        *Thing  `json:"topic"`        // //http://www.w3.org/2002/07/owl#Thing
}

// Group is a generated struct representing the http://xmlns.com/foaf/0.1/Group class.
type Group struct {
	Member *Agent `json:"member"` // //http://xmlns.com/foaf/0.1/Agent
}

// Image is a generated struct representing the http://xmlns.com/foaf/0.1/Image class.
type Image struct {
	Depicts   *Thing `json:"depicts"`   // //http://www.w3.org/2002/07/owl#Thing
	Thumbnail *Image `json:"thumbnail"` // //http://xmlns.com/foaf/0.1/Image
}

// LabelProperty is a generated struct representing the http://xmlns.com/foaf/0.1/LabelProperty class.
type LabelProperty struct {
}

// Project is a generated struct representing the http://xmlns.com/foaf/0.1/Project class.
type Project struct {
}

// OnlineEcommerceAccount is a generated struct representing the http://xmlns.com/foaf/0.1/OnlineEcommerceAccount class.
type OnlineEcommerceAccount struct {
}

// Organization is a generated struct representing the http://xmlns.com/foaf/0.1/Organization class.
type Organization struct {
}

// Agent is a generated struct representing the http://xmlns.com/foaf/0.1/Agent class.
type Agent struct {
	Account       *OnlineAccount `json:"account"`       // //http://xmlns.com/foaf/0.1/OnlineAccount
	Age           *Literal       `json:"age"`           // //http://www.w3.org/2000/01/rdf-schema#Literal
	AimChatID     *Literal       `json:"aimchatid"`     // //http://www.w3.org/2000/01/rdf-schema#Literal
	Birthday      *Literal       `json:"birthday"`      // //http://www.w3.org/2000/01/rdf-schema#Literal
	Gender        *Literal       `json:"gender"`        // //http://www.w3.org/2000/01/rdf-schema#Literal
	HoldsAccount  *OnlineAccount `json:"holdsaccount"`  // //http://xmlns.com/foaf/0.1/OnlineAccount
	IcqChatID     *Literal       `json:"icqchatid"`     // //http://www.w3.org/2000/01/rdf-schema#Literal
	Interest      *Document      `json:"interest"`      // //http://xmlns.com/foaf/0.1/Document
	JabberID      *Literal       `json:"jabberid"`      // //http://www.w3.org/2000/01/rdf-schema#Literal
	Made          *Thing         `json:"made"`          // //http://www.w3.org/2002/07/owl#Thing
	Mbox          *Thing         `json:"mbox"`          // //http://www.w3.org/2002/07/owl#Thing
	Mboxsha1sum   *Literal       `json:"mboxsha1sum"`   // //http://www.w3.org/2000/01/rdf-schema#Literal
	MsnChatID     *Literal       `json:"msnchatid"`     // //http://www.w3.org/2000/01/rdf-schema#Literal
	Openid        *Document      `json:"openid"`        // //http://xmlns.com/foaf/0.1/Document
	SkypeID       *Literal       `json:"skypeid"`       // //http://www.w3.org/2000/01/rdf-schema#Literal
	Status        *Literal       `json:"status"`        // //http://www.w3.org/2000/01/rdf-schema#Literal
	Tipjar        *Document      `json:"tipjar"`        // //http://xmlns.com/foaf/0.1/Document
	Topicinterest *Thing         `json:"topicinterest"` // //http://www.w3.org/2002/07/owl#Thing
	Weblog        *Document      `json:"weblog"`        // //http://xmlns.com/foaf/0.1/Document
	YahooChatID   *Literal       `json:"yahoochatid"`   // //http://www.w3.org/2000/01/rdf-schema#Literal
}

// OnlineAccount is a generated struct representing the http://xmlns.com/foaf/0.1/OnlineAccount class.
type OnlineAccount struct {
	AccountName            *Literal  `json:"accountname"`            // //http://www.w3.org/2000/01/rdf-schema#Literal
	AccountServiceHomepage *Document `json:"accountservicehomepage"` // //http://xmlns.com/foaf/0.1/Document
}

// PersonalProfileDocument is a generated struct representing the http://xmlns.com/foaf/0.1/PersonalProfileDocument class.
type PersonalProfileDocument struct {
}
