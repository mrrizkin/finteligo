package api_token

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/app/utils"
	"github.com/mrrizkin/finteligo/system/types"
)

func NewService(repo *Repo) *Service {
	return &Service{repo}
}

func (s *Service) Create(user *models.User, apiToken *models.ApiToken) (*models.ApiToken, error) {
	token, err := generateToken()
	if err != nil {
		return nil, err
	}

	apiToken.Token = token
	apiToken.UserId = user.ID

	err = s.repo.Create(apiToken)
	if err != nil {
		return nil, err
	}

	return apiToken, nil
}

func (s *Service) FindAll(
	user *models.User,
	pagination types.Pagination,
) (*PaginatedApiToken, error) {
	wb := utils.NewWhereBuilder()

	wb.And("user_id", user.ID)

	where, whereArgs := wb.Get()

	filter := types.Filter{
		Where:     where,
		WhereArgs: whereArgs,
	}

	apiTokens, err := s.repo.FindAll(pagination, filter)
	if err != nil {
		return nil, err
	}

	for i := range apiTokens {
		apiTokens[i].Token = apiTokens[i].Token[:12] + "..." + apiTokens[i].Token[len(apiTokens[i].Token)-4:]
	}

	apiTokensCount, err := s.repo.FindAllCount(filter)
	if err != nil {
		return nil, err
	}

	return &PaginatedApiToken{
		Result: apiTokens,
		Total:  int(apiTokensCount),
	}, nil
}

func (s *Service) FindByID(id uint) (*models.ApiToken, error) {
	apiToken, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return apiToken, nil
}

func (s *Service) Update(id uint, apiToken *models.ApiToken) (*models.ApiToken, error) {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(apiToken)
	if err != nil {
		return nil, err
	}

	return apiToken, nil
}

func (s *Service) Enable(id uint) (*models.ApiToken, error) {
	return s.repo.Enable(id)
}

func (s *Service) Disable(id uint) (*models.ApiToken, error) {
	return s.repo.Disable(id)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}

func generateToken() (string, error) {
	token, err := gonanoid.Generate(
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-",
		32,
	)
	return "fin_sk_" + token, err
}
