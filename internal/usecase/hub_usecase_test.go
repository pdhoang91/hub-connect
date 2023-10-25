package usecase

import (
	"fmt"
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
		{
			name: "case insert error",
			fields: fields{
				hubRepository: func() repo.HubRepository {
					mockRepo := &mocks.HubRepository{}
					// mock result here
					mockRepo.On("Create", &entities.Hub{
						Name:     "Hub 1",
						Location: "Location A",
					}).Return(fmt.Errorf("error insert"))
					return mockRepo
				}(),
			},
			args: args{
				name:     "Hub 1",
				location: "Location A",
			},
			want:    nil,
			wantErr: true,
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

func TestIHubUseCase_GetHubByID(t *testing.T) {

	type fields struct {
		hubRepository repo.HubRepository
	}

	type args struct {
		id int
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
			name: "Happy case GetHubByID",
			fields: fields{
				hubRepository: func() repo.HubRepository {
					mockRepo := &mocks.HubRepository{}
					// mock result here
					mockRepo.On("FindByID", 1).Return(&entities.Hub{
						Name:     "Hub 1",
						Location: "Location A",
					}, nil)
					return mockRepo
				}(),
			},
			args: args{
				id: 1,
			},
			want: &entities.Hub{
				Name:     "Hub 1",
				Location: "Location A",
			},
			wantErr: false,
		},
		{
			name: "case GetHubByID not found",
			fields: fields{
				hubRepository: func() repo.HubRepository {
					mockRepo := &mocks.HubRepository{}
					// mock result here
					mockRepo.On("FindByID", 1).Return(nil, fmt.Errorf("not found"))
					return mockRepo
				}(),
			},
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IHubUseCase{
				HubRepository: tt.fields.hubRepository,
			}
			got, err := s.GetHubByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("IHubUseCase.GetHubByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IHubUseCase.GetHubByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIHubUseCase_SearchHubs(t *testing.T) {

	type fields struct {
		hubRepository repo.HubRepository
	}

	type args struct {
		keyword string
	}
	tests := []struct {
		fields  fields
		name    string
		hi      *IHubUseCase
		args    args
		want    []*entities.Hub
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Happy case SearchHubs",
			fields: fields{
				hubRepository: func() repo.HubRepository {
					mockRepo := &mocks.HubRepository{}
					// mock result here
					mockRepo.On("SearchHubs", "").Return([]*entities.Hub{
						{Name: "Hub 1",
							Location: "Location A"},
					}, nil)
					return mockRepo
				}(),
			},
			args: args{
				keyword: "",
			},
			want: []*entities.Hub{
				{Name: "Hub 1",
					Location: "Location A"},
			},
			wantErr: false,
		},
		{
			name: "case SearchHubs not found",
			fields: fields{
				hubRepository: func() repo.HubRepository {
					mockRepo := &mocks.HubRepository{}
					// mock result here
					mockRepo.On("SearchHubs", "").Return(nil, fmt.Errorf("not found"))
					return mockRepo
				}(),
			},
			args: args{
				keyword: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IHubUseCase{
				HubRepository: tt.fields.hubRepository,
			}
			got, err := s.SearchHubs(tt.args.keyword)
			if (err != nil) != tt.wantErr {
				t.Errorf("IHubUseCase.SearchHubs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IHubUseCase.SearchHubs() = %v, want %v", got, tt.want)
			}
		})
	}
}
