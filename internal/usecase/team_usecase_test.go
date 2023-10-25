package usecase

import (
	"fmt"
	"hub-connect/internal/entities"
	repo "hub-connect/internal/repository"
	mocks "hub-connect/internal/usecase/mocks"
	"reflect"
	"testing"
)

func TestITeamUseCase_CreateTeam(t *testing.T) {

	type fields struct {
		teamRepository repo.TeamRepository
		hubRepository  repo.HubRepository
	}
	type args struct {
		name     string
		teamType string
	}
	tests := []struct {
		fields  fields
		name    string
		args    args
		want    *entities.Team
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Happy case create team",
			fields: fields{
				teamRepository: func() repo.TeamRepository {
					mockRepo := &mocks.TeamRepository{}
					// mock result here
					mockRepo.On("Create", &entities.Team{
						Name: "Team 1",
						Type: "Type A",
					}).Return(nil)
					return mockRepo
				}(),

				hubRepository: nil,
			},
			args: args{
				name:     "Team 1",
				teamType: "Type A",
			},
			want: &entities.Team{
				Name: "Team 1",
				Type: "Type A",
			},
			wantErr: false,
		},
		{
			name: "case create error",
			fields: fields{
				teamRepository: func() repo.TeamRepository {
					mockRepo := &mocks.TeamRepository{}
					// mock result here
					mockRepo.On("Create", &entities.Team{
						Name: "Team 1",
						Type: "Type A",
					}).Return(fmt.Errorf("error insert"))
					return mockRepo
				}(),

				hubRepository: nil,
			},
			args: args{
				name:     "Team 1",
				teamType: "Type A",
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ITeamUseCase{
				TeamRepository: tt.fields.teamRepository,
				HubRepository:  tt.fields.hubRepository,
			}
			got, err := s.CreateTeam(tt.args.name, tt.args.teamType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ITeamUseCase.CreateTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ITeamUseCase.CreateTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}
