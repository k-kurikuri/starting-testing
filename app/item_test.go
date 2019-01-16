package app_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/pkg/errors"

	"github.com/k-kurikuri/starting-testing/app"
)

type itemMock struct {
	app.ItemService
	mockListFunc func() ([]app.ItemData, error)
}

func (m *itemMock) List() ([]app.ItemData, error) {
	return m.mockListFunc()
}

func TestMain(m *testing.M) {
	fmt.Println("setUp")
	code := m.Run()
	if code == 0 {
		fmt.Println("tearDown")
	}

	os.Exit(code)
}

func TestItem_List(t *testing.T) {
	cases := map[string]struct {
		mockFuncList func() ([]app.ItemData, error)
		expected     int
		wantError    bool
	}{
		"success": {
			mockFuncList: func() ([]app.ItemData, error) {
				items := make([]app.ItemData, 0)
				items = append(items, app.ItemData{ID: 1, Name: "Item1"})
				items = append(items, app.ItemData{ID: 2, Name: "Item2"})

				return items, nil
			},
			wantError: false,
			expected:  2,
		},
		"error": {
			mockFuncList: func() ([]app.ItemData, error) {
				return nil, errors.New("something wrong...")
			},
			wantError: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			mock := &itemMock{mockListFunc: tc.mockFuncList}
			client := &app.Client{ItemSrv: mock}

			items, err := client.ItemSrv.List()
			if !tc.wantError && err != nil {
				t.Error("item List returns error")
			}

			if tc.wantError {
				if err == nil {
					t.Error("want error")
				}
				return
			}

			if items == nil {
				t.Error("item must not nil")
			}

			if got, want := len(items), tc.expected; got != want {
				t.Errorf("items length want %d, but %d", want, got)
			}
		})
	}
}
