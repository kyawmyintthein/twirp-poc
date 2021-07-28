package rpcv1

import (
	"context"

	"github.com/kyawmyintthein/twirp-poc/app/config"
	"github.com/kyawmyintthein/twirp-poc/app/domain/dto"
	ucv1 "github.com/kyawmyintthein/twirp-poc/app/domain/usecase/v1"
	userpbs "github.com/kyawmyintthein/twirp-poc/protopbs/protos/user/v1"
)

type userServiceRPCImpl struct {
	cfg    *config.GlobalCfg
	userUC ucv1.UserUC
}

func NewUserServiceRPCImpl(cfg *config.GlobalCfg, userUC ucv1.UserUC) userpbs.UserServce {
	return &userServiceRPCImpl{
		cfg:    cfg,
		userUC: userUC,
	}
}

func (impl *userServiceRPCImpl) Onboarding(ctx context.Context, req *userpbs.UserOnboardingRequest) (*userpbs.LoginSession, error) {
	var (
		resp userpbs.LoginSession
		err  error
	)

	loginSession, err := impl.userUC.Onboarding(ctx, &dto.UserOnboardingRequest{
		DeviceID:   req.DeviceId,
		DeviceType: req.DeviceType,
	})

	resp.AccessToken = loginSession.Token

	return &resp, err
}

func (impl *userServiceRPCImpl) Profile(ctx context.Context, req *userpbs.GetProfileRequest) (*userpbs.UserProfile, error) {
	var (
		resp userpbs.UserProfile
		err  error
	)

	profile, devices, err := impl.userUC.GetProfile(ctx, ucv1.UserID(req.UserID), req.IncludeDevices)
	if err != nil {
		return &resp, err
	}

	resp.FirstName = profile.FirstName
	resp.LastName = profile.FirstName
	resp.IsdCode = profile.ISDCode
	resp.MobileNo = profile.MobileNo
	resp.Uuid = profile.UUID

	var pbdevices []*userpbs.Device
	if len(devices) == 0 {
		for _, device := range devices {
			pbdevices = append(pbdevices, &userpbs.Device{
				DeviceId:   device.DeviceID,
				DeviceType: device.DeviceType,
			})
		}
	}
	resp.Devices = pbdevices
	return &resp, err
}

func (impl *userServiceRPCImpl) UpdateProfile(ctx context.Context, req *userpbs.UpdateProfileRequest) (*userpbs.Empty, error) {
	var (
		resp userpbs.Empty
		err  error
	)

	err = impl.userUC.UpdateProfile(ctx, ucv1.UserID(req.Uuid), &dto.UpdateProfileRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		return &resp, err
	}

	return &resp, err
}
