package executor

import (
	"fmt"
	"testing"
)

func TestExecPipeline(t *testing.T) {
	fmt.Println()
	commands := [][]string{
		//{"echo", "hello gopher", "hello kotaro"}, // echo で文字列を出力
		{"ls", "-la"},
		{"grep", "go"}, // "go" を含む行をフィルタリング
	}
	ExecPipeline(commands)
}

// func TestExec1pipe(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		words1 []string
// 		words2 []string
// 		wantErr bool
// 	}{
// 		{"Valid command", []string{"ls", "-la"}, []string{"grep", "go"}, false},
// 		{"Invalid first command", []string{"invalidcmd"}, []string{"grep", "go"}, true},
// 		{"Invalid second command", []string{"ls", "-la"}, []string{"invalidcmd"}, true},
// 	}

//		for _, tt := range tests {
//			t.Run(tt.name, func(t *testing.T) {
//				err := Exec1pipe(tt.words1, tt.words2)
//				if (err != nil) != tt.wantErr {
//					t.Errorf("Exec1pipe(%v, %v) error = %v, wantErr %v", tt.words1, tt.words2, err, tt.wantErr)
//				}
//			})
//			fmt.Println()
//		}
//	}
func TestExec1pipe(t *testing.T) {

	words1 := []string{"ls", "-la"}
	words2 := []string{"grep", "pipe"}
	Exec1pipe(words1, words2)
}
