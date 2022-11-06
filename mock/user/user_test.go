package user

import (
	mock "deep/mock_gen"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockMale := mock.NewMockMale(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Add().Return(nil),
	)
	user := NewUser(mockMale)
	sum := user.GetUserInfo()
	fmt.Println(sum)
}
