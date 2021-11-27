package lishogi

import (
	"encoding/json"
	"io"
)

const (
	teamBasePath = "/team"
)

type TeamService interface {
	Get(teamID string) (*Team, error)
	GetMembers(teamID string) ([]*TeamMember, error)
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
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Open        bool      `json:"open"`
	Leader      *Leader   `json:"leader"`
	Leaders     []*Leader `json:"leaders"`
	NbMembers   int       `json:"nbMembers"`
	Location    string    `json:"location"`
}

type Leader struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Patron bool   `json:"patron"`
	ID     string `json:"id"`
}

func (t *TeamServiceOp) Get(teamID string) (*Team, error) {
	res, err := t.client.Get(t.teamPath + "/" + teamID)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	team := &Team{}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(team); err != nil {
		return nil, err
	}
	return team, nil
}

type TeamMember struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Online   bool   `json:"online"`
	Perfs    struct {
		Blitz          *GameResult `json:"blitz"`
		Puzzle         *GameResult `json:"puzzle"`
		UltraBullet    *GameResult `json:"ultraBullet"`
		Bullet         *GameResult `json:"bullet"`
		Correspondence *GameResult `json:"correspondence"`
		Classical      *GameResult `json:"classical"`
		Rapid          *GameResult `json:"rapid"`
	} `json:"perfs"`
	CreatedAt int64 `json:"createdAt"`
	Profile   struct {
		Country   string `json:"country"`
		Location  string `json:"location"`
		Bio       string `json:"bio"`
		FirstName string `json:"firstName"`
		Links     string `json:"links"`
	} `json:"profile"`
	SeenAt   int64 `json:"seenAt"`
	Patron   bool  `json:"patron"`
	PlayTime struct {
		Total int `json:"total"`
		Tv    int `json:"tv"`
	} `json:"playTime"`
	Language string `json:"language"`
	URL      string `json:"url"`
}

type GameResult struct {
	Games  int  `json:"games"`
	Rating int  `json:"rating"`
	Rd     int  `json:"rd"`
	Prog   int  `json:"prog"`
	Prov   bool `json:"prov"`
}

func (t *TeamServiceOp) GetMembers(teamID string) ([]*TeamMember, error) {
	res, err := t.client.Get(t.teamPath + "/" + teamID + "/users")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	members := make([]*TeamMember, 0, 100)

	decoder := json.NewDecoder(res.Body)
	for {
		member := &TeamMember{}
		if err := decoder.Decode(member); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}
