package search

type searchUsecase struct {
	searchRepository Repository
}

func (uc *searchUsecase) GetSearchHistory(userId uint, keyword string, limit int) ([]Domain, error) {
	if keyword == "" {
		return uc.searchRepository.GetLastSearchHistory(userId, limit)
	} else {
		return uc.searchRepository.QuerySearchHistory(userId, keyword, limit)
	}
}

func (uc *searchUsecase) StoreSearchKeyword(userId uint, data *Domain) error {
	return uc.searchRepository.StoreSearchKeyword(userId, data)
}

func InitSearchUsecase(r Repository) Usecase {
	return &searchUsecase{
		searchRepository: r,
	}
}
