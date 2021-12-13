package repositories

import (
	"letsgo/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListTeamsByCompany(t *testing.T) {
	teams := testTeamRepo.ListTeamsByCompany(1)
	assert.Equal(t, 1, len(teams))
	assert.Equal(t, helpers.TestTm.Name, teams[0].Name)
}

func TestFindCompanyTeamByName(t *testing.T) {
	team, err := testTeamRepo.FindCompanyTeamByName(1, helpers.TestTm.Name)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, helpers.TestTm.Name, team.Name)
}
