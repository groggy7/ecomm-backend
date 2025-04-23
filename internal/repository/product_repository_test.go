package repository

import (
	"testing"
)

/* func withTestDB(t *testing.T, fn func(*sql.DB, sqlmock.Sqlmock)) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	fn(db, mock)
} */

func TestCreateProduct(t *testing.T) {
	/*
		 	withTestDB(t, func(db *sql.DB, mock sqlmock.Sqlmock) {
				repo := NewProductRepo(db)

				product := &domain.Product{
					Name:            "Test Product",
					Image:           "test.jpg",
					Category:        "Test Category",
					Description:     "Test Description",
					Rating:          6,
					NumberOfReviews: 10,
					Price:           19.99,
					CountInStock:    100,
				}

				testUUID := uuid.New().String()
				expectedQuery := `
					INSERT INTO
					products(name, image, category, description, rating, num_reviews, price, count_in_stock)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
					RETURNING id
				`

				rows := mock.NewRows([]string{"id"}).AddRow(testUUID)
				mock.ExpectQuery(expectedQuery).
					WithArgs(&product.Name, &product.Image, &product.Category, &product.Description, &product.Rating,
						&product.NumberOfReviews, &product.Price, &product.CountInStock).
					WillReturnRows(rows)

				result, err := repo.CreateProduct(product)
				require.NoError(t, err)
				require.NotNil(t, result)
				require.Equal(t, testUUID, result.ID)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			})
	*/
}

func TestCreateProduct_Error(t *testing.T) {
	/*
		 	withTestDB(t, func(db *sql.DB, mock sqlmock.Sqlmock) {
				repo := NewProductRepo(db)

				product := &domain.Product{
					Name:            "Test Product",
					Image:           "test.jpg",
					Category:        "Test Category",
					Description:     "Test Description",
					Rating:          5,
					NumberOfReviews: 10,
					Price:           19.99,
					CountInStock:    100,
				}

				expectedQuery := `
					INSERT INTO
					products(name, image, category, description, rating, num_reviews, price, count_i	n_stock)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
					RETURNING id
				`

				mock.ExpectQuery(expectedQuery).
					WithArgs(product.Name, product.Image, product.Category, product.Description, product.Rating, product.NumberOfReviews, product.Price, product.CountInStock).
					WillReturnError(fmt.Errorf("error inserting product"))

				result, err := repo.CreateProduct(product)

				require.Error(t, err)
				require.Nil(t, result)
				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			})
	*/
}
