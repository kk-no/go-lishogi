package lishogi

import (
	"encoding/json"
	"time"
)

const (
	tournamentBasePath = "/tournament"
)

type TournamentService interface {
	Get(tournamentID string) (*Tournament, error)
	GetStanding(tournamentID, page string) (*TournamentStanding, error)
}

type TournamentServiceOp struct {
	tournamentPath string
	client         *Client
}

var _ TournamentService = (*TournamentServiceOp)(nil)

func NewTournamentService(pathPrefix string, client *Client) TournamentService {
	return &TournamentServiceOp{
		tournamentPath: pathPrefix + tournamentBasePath,
		client:         client,
	}
}

type Tournament struct {
	NbPlayers int           `json:"nbPlayers"`
	Duels     []interface{} `json:"duels"`
	Standing  struct {
		Page    int `json:"page"`
		Players []struct {
			Name   string `json:"name"`
			Rank   int    `json:"rank"`
			Rating int    `json:"rating"`
			Score  int    `json:"score"`
			Sheet  struct {
				Scores []interface{} `json:"scores"`
				Total  int           `json:"total"`
				Fire   bool          `json:"fire"`
			} `json:"sheet,omitempty"`
			Team        string `json:"team"`
			Provisional bool   `json:"provisional,omitempty"`
			Title       string `json:"title,omitempty"`
		} `json:"players"`
	} `json:"standing"`
	IsFinished bool `json:"isFinished"`
	Podium     []struct {
		Name   string `json:"name"`
		Rank   int    `json:"rank"`
		Rating int    `json:"rating"`
		Score  int    `json:"score"`
		Sheet  struct {
			Scores []interface{} `json:"scores"`
			Total  int           `json:"total"`
			Fire   bool          `json:"fire"`
		} `json:"sheet"`
		Team string `json:"team"`
		Nb   struct {
			Game    int `json:"game"`
			Berserk int `json:"berserk"`
			Win     int `json:"win"`
		} `json:"nb"`
		Performance int  `json:"performance"`
		Provisional bool `json:"provisional,omitempty"`
	} `json:"podium"`
	PairingsClosed bool `json:"pairingsClosed"`
	Stats          struct {
		Games         int `json:"games"`
		Moves         int `json:"moves"`
		SenteWins     int `json:"senteWins"`
		GoteWins      int `json:"goteWins"`
		Draws         int `json:"draws"`
		Berserks      int `json:"berserks"`
		AverageRating int `json:"averageRating"`
	} `json:"stats"`
	TeamStanding []struct {
		Rank    int    `json:"rank"`
		ID      string `json:"id"`
		Score   int    `json:"score"`
		Players []struct {
			User struct {
				Name   string `json:"name"`
				Patron bool   `json:"patron"`
				Title  string `json:"title"`
				ID     string `json:"id"`
			} `json:"user,omitempty"`
			Score int `json:"score"`
		} `json:"players"`
	} `json:"teamStanding"`
	DuelTeams struct{}  `json:"duelTeams"`
	ID        string    `json:"id"`
	CreatedBy string    `json:"createdBy"`
	StartsAt  time.Time `json:"startsAt"`
	System    string    `json:"system"`
	FullName  string    `json:"fullName"`
	Minutes   int       `json:"minutes"`
	Perf      struct {
		Icon string `json:"icon"`
		Name string `json:"name"`
	} `json:"perf"`
	Clock struct {
		Limit     int `json:"limit"`
		Increment int `json:"increment"`
		Byoyomi   int `json:"byoyomi"`
	} `json:"clock"`
	Variant     string `json:"variant"`
	Berserkable bool   `json:"berserkable"`
	Verdicts    struct {
		List     []interface{} `json:"list"`
		Accepted bool          `json:"accepted"`
	} `json:"verdicts"`
}

func (t *TournamentServiceOp) Get(tournamentID string) (*Tournament, error) {
	res, err := t.client.Get(t.tournamentPath + "/" + tournamentID)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	tournament := &Tournament{}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(tournament); err != nil {
		return nil, err
	}

	return tournament, nil
}

type TournamentStanding struct {
	Page    int `json:"page"`
	Players []struct {
		Name   string `json:"name"`
		Rank   int    `json:"rank"`
		Rating int    `json:"rating"`
		Score  int    `json:"score"`
		Sheet  struct {
			Scores []interface{} `json:"scores"`
			Total  int           `json:"total"`
			Fire   bool          `json:"fire"`
		} `json:"sheet,omitempty"`
		Team        string `json:"team"`
		Provisional bool   `json:"provisional,omitempty"`
		Title       string `json:"title,omitempty"`
	} `json:"players"`
}

func (t *TournamentServiceOp) GetStanding(tournamentID, page string) (*TournamentStanding, error) {
	res, err := t.client.Get(t.tournamentPath + "/" + tournamentID + "/standing/" + page)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	standing := &TournamentStanding{}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(standing); err != nil {
		return nil, err
	}

	return standing, nil
}
