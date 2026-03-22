package subcategories

import "database/sql"

type SubcategoryRepository struct {
	db *sql.DB
}

func NewSubcategoryRepository(db *sql.DB) *SubcategoryRepository {
	return &SubcategoryRepository{db: db}
}

func (r *SubcategoryRepository) List() ([]SubCategory, error) {
	rows, err := r.db.Query(`
		SELECT * FROM subcategories
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subcategories []SubCategory

	for rows.Next() {
		var s SubCategory
		err := rows.Scan(
			&s.SubCategoryId,
			&s.SubCategoryName,
		)
		if err != nil {
			return nil, err
		}
		subcategories = append(subcategories, s)
	}

	return subcategories, nil
}

func (r *SubcategoryRepository) GetById(id int) (*SubCategory, error) {
	row := r.db.QueryRow(`
		SELECT * FROM subcategories WHERE sub_category_id = $1`, id)

	var subcategory SubCategory

	err := row.Scan(
		&subcategory.SubCategoryId,
		&subcategory.SubCategoryName,
	)
	if err != nil {
		return nil, err
	}

	return &subcategory, nil
}

func (r *SubcategoryRepository) Create(s *SubCategory) error {
	return r.db.QueryRow(
		`INSERT into subcategories (sub_category_name)
		VALUES ($1)
		RETURNING sub_category_id`,
		s.SubCategoryName,
	).Scan(&s.SubCategoryId)
}

func (r *SubcategoryRepository) Update(s *SubCategory) error {
	_, err := r.db.Exec(`
		UPDATE subcategories
		SET sub_category_name = $1
		WHERE sub_category_id = $2`,
		s.SubCategoryName, s.SubCategoryId,
	)

	return err
}

func (r *SubcategoryRepository) Delete(s *SubCategory) error {
	_, err := r.db.Exec(`
		DELETE FROM subcategories
		WHERE sub_category_id = $1`,
		s.SubCategoryId,
	)

	return err
}

func (r *SubcategoryRepository) ListIncomeTypes() ([]SubCategory, error) {
	rows, err := r.db.Query(`
		SELECT * FROM subcategories WHERE sub_category_type = $1`, "income")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subcategories []SubCategory

	for rows.Next() {
		var s SubCategory
		err := rows.Scan(
			&s.SubCategoryId,
			&s.SubCategoryName,
		)
		if err != nil {
			return nil, err
		}
		subcategories = append(subcategories, s)
	}

	return subcategories, nil
}

func (r *SubcategoryRepository) ListExpenseTypes() ([]SubCategory, error) {
	rows, err := r.db.Query(`
		SELECT * FROM subcategories WHERE sub_category_type = $1`, "expense")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subcategories []SubCategory

	for rows.Next() {
		var s SubCategory
		err := rows.Scan(
			&s.SubCategoryId,
			&s.SubCategoryName,
		)
		if err != nil {
			return nil, err
		}
		subcategories = append(subcategories, s)
	}

	return subcategories, nil
}
