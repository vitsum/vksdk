package api

import (
	"log"
	"os"
	"testing"
)

func needUserToken(t *testing.T) {
	t.Helper()
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}
}

func needGroupToken(t *testing.T) {
	t.Helper()
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
}

func needServiceToken(t *testing.T) {
	t.Helper()
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
}

var vkGroup, vkService, vkUser *VK    // nolint:gochecknoglobals
var vkUserID, vkGroupID, vkChatID int // nolint:gochecknoglobals

func TestMain(m *testing.M) {
	vkGroup = Init(os.Getenv("GROUP_TOKEN"))
	if vkGroup.AccessToken != "" {
		group, err := vkGroup.GroupsGetByID(map[string]string{})
		if err != nil {
			log.Fatal(err)
		}
		vkGroupID = group[0].ID
	}
	vkService = Init(os.Getenv("SERVICE_TOKEN"))
	vkService.Limit = 3
	vkUser = Init(os.Getenv("USER_TOKEN"))
	vkUser.Limit = 3
	if vkUser.AccessToken != "" {
		user, err := vkUser.UsersGet(map[string]string{})
		if err != nil {
			log.Fatal(err)
		}
		vkUserID = user[0].ID

		vkChatID, err = vkUser.MessagesCreateChat(map[string]string{
			"title": "TestChat",
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	runTests := m.Run()
	os.Exit(runTests)
}

func TestVK_Request(t *testing.T) {
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(groupToken)

	t.Run("Request 403 error", func(t *testing.T) {
		_, err := vk.Request("", map[string]string{})
		if err == nil {
			t.Errorf("VK.Request() got1 = %v, want -1", err)
		}
	})
	vk.MethodURL = ""
	t.Run("Client error", func(t *testing.T) {
		_, err := vk.Request("test", map[string]string{"test": "test"})
		if err == nil {
			t.Errorf("VK.Request() got1 = %v, want -1", err)
		}
	})
}

func TestVK_RequestLimit(t *testing.T) {
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(groupToken)
	vk.Limit = 2

	t.Run("vk.Limit", func(t *testing.T) {
		// TODO: check err
		go vk.UsersGet(map[string]string{}) // nolint: errcheck
		for i := 0; i < 2; i++ {
			vk.UsersGet(map[string]string{}) // nolint: errcheck
		}
	})
}

func TestVK_Execute(t *testing.T) {
	needGroupToken(t)

	t.Run("Execute test", func(t *testing.T) {
		var response int
		err := vkGroup.Execute(`return 1;`, &response)
		if err != nil {
			t.Errorf("VK.Execute() err = %v, want 0", err)
		}
		if response != 1 {
			t.Error("Execute response error")
		}
	})
}

func TestVK_RequestUnmarshal(t *testing.T) {
	needGroupToken(t)

	var testObj string
	type args struct {
		method string
		params map[string]string
		obj    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "execute error",
			args: args{
				method: "execute",
				params: map[string]string{"code": "return 1;"},
				obj:    &testObj,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := vkGroup.RequestUnmarshal(tt.args.method, tt.args.params, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("VK.RequestUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
