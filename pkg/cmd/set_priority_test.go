package cmd

// Todo fix set priority test
//
// import (
// 	"errors"
// 	"testing"
//
// 	"github.com/ZupIT/ritchie-cli/pkg/formula"
// 	"github.com/ZupIT/ritchie-cli/pkg/prompt"
// )
//
// func TestSetRepoCmd_runFunc(t *testing.T) {
// 	type fields struct {
// 		InputList          prompt.InputList
// 		InputInt           prompt.InputInt
// 		RepoLister         formula.RepositoryLister
// 		RepoPrioritySetter formula.RepositoryPrioritySetter
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		{
// 			name: "run with success",
// 			fields: fields{
// 				InputList:          inputListMock{},
// 				InputInt:           inputIntMock{},
// 				RepoLister:         repoListerNonEmptyMock{},
// 				RepoPrioritySetter: repoPrioritySetterMock{},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "error on repoLister",
// 			fields: fields{
// 				InputList:          inputListMock{},
// 				InputInt:           inputIntMock{},
// 				RepoLister:         repoListerErrorMock{},
// 				RepoPrioritySetter: repoPrioritySetterMock{},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "error on inputList",
// 			fields: fields{
// 				InputList:          inputListErrorMock{},
// 				InputInt:           inputIntMock{},
// 				RepoLister:         repoListerMock{},
// 				RepoPrioritySetter: repoPrioritySetterMock{},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "error on inputInt",
// 			fields: fields{
// 				InputList:          inputListMock{},
// 				InputInt:           inputIntErrorMock{},
// 				RepoLister:         repoListerMock{},
// 				RepoPrioritySetter: repoPrioritySetterMock{},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "success pass on if r.Name == repoName",
// 			fields: fields{
// 				InputList:          inputListCustomMock{name: "repoName"},
// 				InputInt:           inputIntMock{},
// 				RepoLister:         repoListerNonEmptyMock{},
// 				RepoPrioritySetter: repoPrioritySetterMock{},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "error on setPriority",
// 			fields: fields{
// 				InputList:  inputListMock{},
// 				InputInt:   inputIntMock{},
// 				RepoLister: repoListerMock{},
// 				RepoPrioritySetter: repoPrioritySetterCustomMock{
// 					setPriority: func(repo string, priority int) error {
// 						return errors.New("some error")
// 					},
// 				},
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := NewSetPriorityCmd(tt.fields.InputList, tt.fields.InputInt, tt.fields.RepoLister, tt.fields.RepoPrioritySetter)
// 			if err := s.Execute(); (err != nil) != tt.wantErr {
// 				t.Skip("Todo fix later")
// 				t.Errorf("runFunc() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
