package product

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetProducts(ctx context.Context, params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int64, error)
	InsertProduct(ctx context.Context, params *getAddProductRequest) (int64, error)
	DeleteProject(ctx context.Context, params *deleteProjectRequest) (int64, error)
	GetProductByID(ctx context.Context, params *getProductByIDRequest) (*Product, error)
	UpdateProduct(ctx context.Context, params *updateProductRequest) (int64, error)
	GetTotalVentas() (float64, error)
	GetBestSellers(ctx context.Context, params *getBestSellersRequest) ([]*ProductTop, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetProducts(ctx context.Context, params *getProductsRequest) ([]*Product, error) {
	const sql = `SELECT id,product_code,product_name,COALESCE(description,''),
				 standard_cost,list_price,COALESCE(reorder_level,0),COALESCE(target_level,0),
				 category
				 FROM products
				 LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)

	if err != nil {
		panic(err.Error())
	}
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description,
			&product.StandardCost, &product.ListPrice, &product.ReorderLevel, &product.TargetLevel,
			&product.Category)

		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *repository) GetTotalProducts() (int64, error) {
	const sql = "SELECT COUNT(*) FROM products"
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err.Error())
	}
	return total, nil
}
func (repo *repository) InsertProduct(ctx context.Context, params *getAddProductRequest) (int64, error) {
	const sql = `
	INSERT INTO products
	(product_code,product_name,category,description,list_price,standard_cost)
	VALUES(?,?,?,?,?,?)`
	result, err := repo.db.Exec(sql, params.Product_Code, params.Product_Name, params.Category, params.Description,
		params.List_Price, params.Standard_Cost)

	if err != nil {
		panic(err.Error())
	}

	id, _ := result.LastInsertId()
	return id, nil
}
func (repo *repository) DeleteProject(ctx context.Context, params *deleteProjectRequest) (int64, error) {
	const sql = `DELETE FROM products WHERE id = ?`
	result, err := repo.db.Exec(sql, params.ProjectID)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return count, nil
}

func (repo *repository) GetProductByID(ctx context.Context, params *getProductByIDRequest) (*Product, error) {
	const sql = `SELECT id,product_code,product_name,COALESCE(description,''),
				standard_cost,list_price,COALESCE(reorder_level,0),COALESCE(target_level,0),
				category
				FROM products
				WHERE id=?`

	row := repo.db.QueryRow(sql, params.ProductID)

	product := &Product{}

	err := row.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description,
		&product.StandardCost, &product.ListPrice, &product.ReorderLevel, &product.TargetLevel,
		&product.Category)

	if err != nil {
		panic(err.Error())
	}
	return product, nil
}
func (repo *repository) UpdateProduct(ctx context.Context, params *updateProductRequest) (int64, error) {
	const sql = `
			UPDATE products
			SET Product_Code = ?,
			Product_Name = ?,
			Category=?,
			Description=?,
			List_Price=?,
			Standard_Cost=?
			WHERE id = ?`
	result, err := repo.db.Exec(sql, params.Product_Code, params.Product_Name, params.Category, params.Description,
		params.List_Price, params.Standard_Cost, params.ID)

	if err != nil {
		panic(err.Error())
	}

	id, _ := result.LastInsertId()
	return id, nil
}
func (repo *repository) GetTotalVentas() (float64, error) {
	const sql = `SELECT SUM(od.quantity*od.unit_price) vendido
				 FROM order_details od`
	var total float64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err.Error())
	}
	return total, nil
}
func (repo *repository) GetBestSellers(ctx context.Context, params *getBestSellersRequest) ([]*ProductTop, error) {
	const sql = `SELECT od.product_id,p.product_name, SUM(od.quantity*od.unit_price) vendido
				FROM northwind.order_details od
				inner join northwind.products p on od.product_id = p.id
				GROUP by od.product_id
				ORDER BY vendido desc
				limit 10`

	results, err := repo.db.Query(sql)

	if err != nil {
		panic(err.Error())
	}
	var products []*ProductTop

	for results.Next() {
		product := &ProductTop{}
		err = results.Scan(&product.ID, &product.ProductName, &product.Vendidos)

		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	return products, nil
}
