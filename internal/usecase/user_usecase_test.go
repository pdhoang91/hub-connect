// Package usecase defines the business logic for managing users.

package usecase

import (
	"fmt"
	"hub-connect/internal/entities"
	repo "hub-connect/internal/repository"
	mocks "hub-connect/internal/usecase/mocks"
	"reflect"
	"testing"
)

func TestIUserUseCase_CreateUser(t *testing.T) {
	type fields struct {
		userRepository repo.UserRepository
		teamRepository repo.TeamRepository
	}

	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Happy case create user",
			fields: fields{
				userRepository: func() repo.UserRepository {
					mockRepo := &mocks.UserRepository{}
					// mock result here
					mockRepo.On("Create", &entities.User{
						Name:  "Name 1",
						Email: "email@gmail.com",
					}).Return(nil)
					return mockRepo
				}(),

				teamRepository: nil,
			},
			args: args{
				name:  "Name 1",
				email: "email@gmail.com",
			},
			want: &entities.User{
				Name:  "Name 1",
				Email: "email@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "case insert error",
			fields: fields{
				userRepository: func() repo.UserRepository {
					mockRepo := &mocks.UserRepository{}
					// mock result here
					mockRepo.On("Create", &entities.User{
						Name:  "Name 1",
						Email: "email@gmail.com",
					}).Return(fmt.Errorf("error insert"))
					return mockRepo
				}(),

				teamRepository: nil,
			},
			args: args{
				name:  "Name 1",
				email: "email@gmail.com",
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IUserUseCase{
				UserRepository: tt.fields.userRepository,
				TeamRepository: tt.fields.teamRepository,
			}
			got, err := s.CreateUser(tt.args.name, tt.args.email)
			fmt.Println("got", got)
			fmt.Println("err", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("IUserUseCase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IUserUseCase.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
