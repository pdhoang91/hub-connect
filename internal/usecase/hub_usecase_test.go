package usecase

import (
	"hub-connect/internal/entities"
	repo "hub-connect/internal/repository"
	mocks "hub-connect/internal/usecase/mocks"
	"reflect"
	"testing"
)

func TestIHubUseCase_CreateHub(t *testing.T) {

	type fields struct {
		hubRepository repo.HubRepository
	}

	type args struct {
		name     string
		location string
	}
	tests := []struct {
		fields  fields
		name    string
		args    args
		want    *entities.Hub
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Happy case create hub",
			fields: fields{
				hubRepository: func() repo.HubRepository {
					mockRepo := &mocks.HubRepository{}
					// mock result here
					mockRepo.On("Create", &entities.Hub{
						Name:     "Hub 1",
						Location: "Location A",
					}).Return(nil)
					return mockRepo
				}(),
			},
			args: args{
				name:     "Hub 1",
				location: "Location A",
			},
			want: &entities.Hub{
				Name:     "Hub 1",
				Location: "Location A",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IHubUseCase{
				HubRepository: tt.fields.hubRepository,
			}
			got, err := s.CreateHub(tt.args.name, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("IHubUseCase.CreateHub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IHubUseCase.CreateHub() = %v, want %v", got, tt.want)
			}
		})
	}
}
