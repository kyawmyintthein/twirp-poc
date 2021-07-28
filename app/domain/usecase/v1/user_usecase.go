package ucv1

import (
	"context"
	"time"

	"github.com/kyawmyintthein/twirp-poc/app/config"
	"github.com/kyawmyintthein/twirp-poc/app/domain/dto"
	"github.com/kyawmyintthein/twirp-poc/app/domain/entity"
	"github.com/twitchtv/twirp"
)

type (
	UserID = string
	UserUC interface {
		Onboarding(context.Context, *dto.UserOnboardingRequest) (*dto.LoginSession, error)
		GetProfile(context.Context, UserID, bool) (*entity.User, []entity.Device, error)
		UpdateProfile(context.Context, UserID, *dto.UpdateProfileRequest) error
	}

	userUC struct {
		cfg *config.GlobalCfg
	}
)

func ProvideUserUC(cfg *config.GlobalCfg) UserUC {
	return &userUC{
		cfg: cfg,
	}
}

func (uc *userUC) Onboarding(ctx context.Context, userOnboardingRequest *dto.UserOnboardingRequest) (*dto.LoginSession, error) {
	var (
		loginSession dto.LoginSession
		err          error
	)
	loginSession.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	return &loginSession, err
}

func (uc *userUC) GetProfile(ctx context.Context, userID UserID, includeDevices bool) (*entity.User, []entity.Device, error) {
	var (
		user    entity.User
		devices []entity.Device
		err     error
	)
	user.UUID = "UID_d08a6940-a7db-450a-b2bb-482895110543"
	user.FirstName = "John"
	user.LastName = "Cooper"
	user.DOB = time.Now().Unix()
	user.ISDCode = "65"
	user.MobileNo = "9734747575"
	devices = append(devices, entity.Device{
		DeviceID:   "b79dadd8-ef8d-11eb-9a03-0242ac130003",
		DeviceType: "iOS",
	})
	return &user, devices, err
}

func (uc *userUC) UpdateProfile(ctx context.Context, userID UserID, updateProfileRequest *dto.UpdateProfileRequest) error {
	return twirp.InvalidArgumentError("uuid", "invalid or empty parameter").WithMeta("uuid", "test")
}
