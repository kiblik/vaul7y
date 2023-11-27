package state

import (
	"github.com/dkyanakiev/vaulty/models"
	"github.com/hashicorp/vault/api"
	"github.com/rivo/tview"
)

type State struct {
	VaultAddress       string
	Mounts             map[string]*models.MountOutput
	SecretsData        []models.SecretPath
	KV2                []models.KVSecret
	Namespace          string
	SelectedNamespace  string
	SelectedMount      string
	SelectedPath       string
	SelectedObject     string
	SelectedPolicyName string
	SelectedSecret     *api.Secret
	PolicyList         []string
	PolicyACL          string

	//Namespaces []*models.Namespace
	Elements *Elements
	Toggle   *Toggle
	Filter   *Filter
}

type Toggle struct {
	Search       bool
	JumpToPolicy bool
}

type Filter struct {
	Policy string
}

type Elements struct {
	DropDownNamespace *tview.DropDown
	TableMain         *tview.Table
	TextMain          *tview.TextView
}

func New() *State {
	return &State{
		Elements: &Elements{},
		Toggle:   &Toggle{},
		Filter:   &Filter{},
	}
}
