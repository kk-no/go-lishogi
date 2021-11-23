package lishogi

const (
	teamBasePath = "/team"
)

type TeamService interface {
	Get(teamID string) (*Team, error)
}

type TeamServiceOp struct {
	teamPath string
	client   *Client
}

var _ TeamService = (*TeamServiceOp)(nil)

func NewTeamService(pathPrefix string, client *Client) TeamService {
	return &TeamServiceOp{
		teamPath: pathPrefix + teamBasePath,
		client:   client,
	}
}

type Team struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Open        bool     `json:"open"`
	Leader      Leader   `json:"leader"`
	Leaders     []Leader `json:"leaders"`
	NbMembers   int      `json:"nbMembers"`
	Location    string   `json:"location"`
}

type Leader struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Patron bool   `json:"patron"`
	ID     string `json:"id"`
}

func (t *TeamServiceOp) Get(teamID string) (*Team, error) {
	team := &Team{}
	if err := t.client.Get(t.teamPath+"/"+teamID, team); err != nil {
		return nil, err
	}
	return team, nil
}
