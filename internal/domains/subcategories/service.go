package subcategories

type SubcategoryService struct {
	repo *SubcategoryRepository
}

func NewSubCategoryService(repo *SubcategoryRepository) *SubcategoryService {
	return &SubcategoryService{repo: repo}
}

func (s *SubcategoryService) CreateSubcategory(req SubCategory) (*SubCategory, error) {
	newSubcategory := &SubCategory{
		SubCategoryName: req.SubCategoryName,
	}

	err := s.repo.Create(newSubcategory)
	if err != nil {
		return nil, err
	}
	return newSubcategory, nil
}

func (s *SubcategoryService) UpdateSubcategory(id int, req SubCategoryPatch) (*SubCategory, error) {
	selectedSubcategory, err := s.repo.GetById(*req.SubCategoryId)
	if err != nil {
		return nil, err
	}

	if req.SubCategoryName != nil {
		selectedSubcategory.SubCategoryName = *req.SubCategoryName
	}

	err = s.repo.Update(selectedSubcategory)
	if err != nil {
		return nil, err
	}

	return selectedSubcategory, nil
}

func (s *SubcategoryService) DeleteSubcategory(id int) (*SubCategory, error) {
	selectedSubcategory, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Delete(selectedSubcategory)
	if err != nil {
		return nil, err
	}

	return selectedSubcategory, nil
}

func (s *SubcategoryService) ListExpenseTypes() ([]SubCategory, error) {
	subcategories, err := s.repo.ListExpenseTypes()
	if err != nil {
		return nil, err
	}
	return subcategories, err
}

func (s *SubcategoryService) ListIncomeTypes() ([]SubCategory, error) {
	subcategories, err := s.repo.ListIncomeTypes()
	if err != nil {
		return nil, err
	}
	return subcategories, err
}

func (s *SubcategoryService) GetById(id int) (*SubCategory, error) {
	subcategory, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return subcategory, nil
}
