package service

import (
	"errors"
	"gorm.io/gorm"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/service/zan"
	"time"
)

type ZanService struct {
	cli *zan.ZanClient
	db  *gorm.DB
}

func NewZanService(cli *zan.ZanClient, db *gorm.DB) *ZanService {
	return &ZanService{
		cli: cli,
		db:  db,
	}
}

func (s *ZanService) GetUserAuthed(user aline.UserPrincipal) bool {
	var zanUser models.ZanUser
	err := s.db.Model(models.ZanUser{}).Where("user_id = ?", user.GetUserId()).First(&zanUser).Error
	if err != nil {
		return false
	}
	return zanUser.AccessToken != ""
}

func (s *ZanService) GetAuthUrl() (string, error) {
	url, err := s.cli.AuthUrl()
	if err != nil {
		return "", err
	}
	return url.Data.AuthUrl, err
}

func (s *ZanService) ExchangeAccessToken(user aline.UserPrincipal, authCode string) error {
	resp, err := s.cli.AccessToken(authCode)
	if err != nil {
		return err
	}

	var zanUser models.ZanUser
	err = s.db.Model(models.ZanUser{}).Where("user_id = ?", user.GetUserId()).First(&zanUser).Error
	if err != nil {
		zanUser.UserId = user.GetUserId()
		zanUser.AccessToken = resp.Data.AccessToken
		zanUser.Created = time.Now()
	} else {
		zanUser.AccessToken = resp.Data.AccessToken
	}
	err = s.db.Model(zanUser).Save(&zanUser).Error

	if err != nil {
		return err
	}

	// 初始化plan
	err = s.cli.InitFreeSpec(resp.Data.AccessToken)

	return err
}

func (s *ZanService) GetUserAccessToken(u aline.UserPrincipal) (string, error) {
	var zanUser models.ZanUser
	err := s.db.Model(models.ZanUser{}).Where("user_id = ?", u.GetUserId()).First(&zanUser).Error
	if err != nil {
		return "", err
	}
	return zanUser.AccessToken, nil
}

func (s *ZanService) CreateApiKey(u aline.UserPrincipal, req zan.ApiKeyCreateReq) (*zan.ApiKeyBase, error) {
	token, err := s.GetUserAccessToken(u)
	if err != nil {
		return nil, err
	}

	created, err := s.cli.ApiKeyCreate(req.Name, token)
	if err != nil {
		return nil, err
	}
	if created.Success {
		return &created.Data, nil
	} else {
		return nil, errors.New(*created.Message)
	}

}

func (s *ZanService) ApiKeyList(u aline.UserPrincipal, page int, size int) (zan.PageResponse[zan.ApiKeyDigestInfo], error) {
	token, err := s.GetUserAccessToken(u)
	if err != nil {
		return zan.PageResponse[zan.ApiKeyDigestInfo]{}, err
	}
	resp, err := s.cli.ApiKeyList(page, size, token)
	if err != nil {
		return zan.PageResponse[zan.ApiKeyDigestInfo]{}, err
	}
	return resp.Data, nil
}

func (s *ZanService) ApiKeyDetail(u aline.UserPrincipal, apiKeyId string) (zan.ApiKeyDetailInfo, error) {
	token, err := s.GetUserAccessToken(u)
	if err != nil {
		return zan.ApiKeyDetailInfo{}, err
	}
	resp, err := s.cli.ApiKeyDetail(apiKeyId, token)
	if err != nil {
		return zan.ApiKeyDetailInfo{}, err
	}
	return resp.Data, nil
}

func (s *ZanService) ApiKeyCreditCost(u aline.UserPrincipal, apiKeyId string) ([]zan.StatCreditCostItem, error) {
	token, err := s.GetUserAccessToken(u)
	if err != nil {
		return []zan.StatCreditCostItem{}, err
	}
	resp, err := s.cli.ApiKeyCreditCost(apiKeyId, token)
	if err != nil {
		return []zan.StatCreditCostItem{}, err
	}
	return resp.Data, nil
}

func (s *ZanService) ApiKeyRequestStats(u aline.UserPrincipal, apiKeyId string, timeInterval string, ecosystem string) ([]zan.StatMethodCountItem, error) {
	token, err := s.GetUserAccessToken(u)
	if err != nil {
		return []zan.StatMethodCountItem{}, err
	}
	resp, err := s.cli.ApiKeyRequestStats(apiKeyId, timeInterval, ecosystem, token)
	if err != nil {
		return []zan.StatMethodCountItem{}, err
	}
	return resp.Data, nil
}

func (s *ZanService) ApiKeyRequestActivityStats(u aline.UserPrincipal, apiKeyId string, timeInterval string, ecosystem string) ([]zan.StatMethodRequestActivityDetail, error) {
	token, err := s.GetUserAccessToken(u)
	if err != nil {
		return []zan.StatMethodRequestActivityDetail{}, err
	}
	resp, err := s.cli.ApiKeyRequestActivityStats(apiKeyId, timeInterval, ecosystem, token)
	if err != nil {
		return []zan.StatMethodRequestActivityDetail{}, err
	}
	return resp.Data, nil
}

func (s *ZanService) ApiKeyRequestOriginStats(u aline.UserPrincipal, apiKeyId string, timeInterval string) ([]zan.StatCreditCostOrigin, error) {
	token, err := s.GetUserAccessToken(u)
	if err != nil {
		return []zan.StatCreditCostOrigin{}, err
	}
	resp, err := s.cli.ApiKeyRequestOriginStats(apiKeyId, timeInterval, token)
	if err != nil {
		return []zan.StatCreditCostOrigin{}, err
	}
	return resp.Data, nil
}

func (s *ZanService) EcosystemsDigest() ([]zan.EcosystemDigestInfo, error) {
	resp, err := s.cli.EcosystemsDigest()
	if err != nil {
		return []zan.EcosystemDigestInfo{}, err
	}
	return resp.Data, nil
}

func (s *ZanService) UserPlan(u aline.UserPrincipal) (*zan.PlanDetailInfo, error) {
	token, err := s.GetUserAccessToken(u)
	if err != nil {
		return nil, err
	}
	resp, err := s.cli.Plan(token)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (s *ZanService) ApiKeyRequestActivityStatsFail(u aline.UserPrincipal, apiKeyId string, timeInterval string, ecosystem string, method string) ([]zan.StatMethodRequestActivityFailedDetailGwInfo, error) {
	token, err := s.GetUserAccessToken(u)
	if err != nil {
		return []zan.StatMethodRequestActivityFailedDetailGwInfo{}, err
	}
	resp, err := s.cli.ApiKeyRequestActivityStatsFail(token, apiKeyId, timeInterval, ecosystem, method)
	if err != nil {
		return []zan.StatMethodRequestActivityFailedDetailGwInfo{}, err
	}
	return resp.Data, nil
}
