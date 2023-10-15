package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mattn/go-sqlite3"
)

var (
	db  *sql.DB
	err error
)

const dbPath = "product.db"

type Product struct {
	ID         int
	Name       string
	Created_at time.Time
	Updated_at time.Time
}

type Variant struct {
	ID           int
	Variant_name string
	Quantity     int
	Product_id   int
	Product      Product
	Created_at   time.Time
	Updated_at   time.Time
}

type ProductWithVariants struct {
	Product
	Variants []Variant
}

func main() {
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// Create table if not exists
	createTables()

	fmt.Println("\n=======================================")
	fmt.Println("1. Create Product")
	id := createProduct("Kemeja")

	fmt.Println("\n=======================================")
	fmt.Println("2. Update Product")
	updateProduct(id, "kemeja pria")

	fmt.Println("\n=======================================")
	fmt.Println("3. get product")
	p := getProductById(id)
	fmt.Printf("%+v\n", p)
	fmt.Printf("product_id %+v\n\n", id)

	fmt.Println("\n=======================================")
	fmt.Println("4. create Variant")
	vid := createVariant("flanel", 20, id)
	fmt.Printf("Variant id %+v\n\n", vid)
	v_before := getVariantById(vid)

	fmt.Println("\n=======================================")
	fmt.Println("5. Update Variant")
	fmt.Printf("Variant before update:\n%+v\n\n", v_before)
	v_before.Quantity = 10
	v_before.Variant_name = "flanel merah"
	updateVariantById(v_before)
	v := getVariantById(vid)
	fmt.Printf("Variant after update:\n%+v\n\n", v)

	fmt.Println("\n=======================================")
	fmt.Println("6. Get Product with variant")
	_ = createVariant("putih", 40, id)
	vid_del := createVariant("coklat", 100, id)

	pv := getProductWithVariant(id)
	fmt.Printf("Product with Variants:\n%+v\n\n", pv)

	fmt.Println("\n=======================================")
	fmt.Println("6. Delete variant by ID")
	deleteVariantById(vid_del)

}

func createTables() {
	pTable := `
		CREATE TABLE IF NOT EXISTS product (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			created_at DATETIME,
			updated_at DATETIME
		);
	`
	vTable := `
		CREATE TABLE IF NOT EXISTS variant (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			variant_name TEXT,
			quantity INTEGER,
			product_id INTEGER,
			created_at DATETIME,
			updated_at DATETIME,
			FOREIGN KEY(product_id) REFERENCES product (id)
		);
	`
	_, err = db.Exec(pTable)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(vTable)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		panic(err)
	}
}

func createProduct(name string) int {
	// var product = Product{}

	sqlStmt := `
		INSERT INTO product(name, created_at, updated_at) VALUES(?, ?, ?)
	`
	now := time.Now()
	result, err := db.Exec(sqlStmt, name, now, now)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	product := getProductById(int(id))

	fmt.Printf("Product created %+v\n", product)

	return int(id)
}

func getProductById(id int) Product {
	product := Product{}
	sqlStmt := `
	SELECT * FROM product WHERE product.id = ?;
	`
	row := db.QueryRow(sqlStmt, id)
	err := row.Scan(&product.ID, &product.Name, &product.Created_at, &product.Updated_at)
	if err != nil {
		panic(err)
	}
	return product
}

func updateProduct(ID int, Name string) {
	sqlStmt := `
	UPDATE product SET name=?, updated_at=? WHERE id=?
	`
	now := time.Now()
	_, err = db.Exec(sqlStmt, Name, now, ID)
	if err != nil {
		panic(err)
	}
}

func createVariant(variant_name string, quantity int, product_id int) int {
	sqlStmt := `
		INSERT INTO variant(variant_name, quantity, product_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?)
	`
	now := time.Now()
	result, err := db.Exec(sqlStmt, variant_name, quantity, product_id, now, now)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.Code == sqlite3.ErrConstraint {
				fmt.Printf("Foreign key constraint violated. Tidak ada product_id %+v di database.\n", product_id)
			}
		}
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	variant := getVariantById(int(id))

	fmt.Printf("Variant created:\n%+v\n\n", variant)

	return int(id)
}

func getVariantById(id int) Variant {
	variant := Variant{}
	sqlStmt := `
	SELECT variant.*, product.* FROM variant LEFT JOIN product ON variant.product_id=product.id WHERE variant.id = ?;
	`
	row := db.QueryRow(sqlStmt, id)
	err := row.Scan(&variant.ID, &variant.Variant_name, &variant.Quantity, &variant.Product_id, &variant.Created_at, &variant.Updated_at, &variant.Product.ID, &variant.Product.Name, &variant.Product.Created_at, &variant.Product.Updated_at)
	if err != nil {
		panic(err)
	}
	return variant
}

func updateVariantById(variant Variant) Variant {
	sqlStmt := `
	UPDATE variant SET variant_name=?, quantity=?, product_id=?, updated_at=? WHERE id=?
	`
	now := time.Now()
	_, err = db.Exec(sqlStmt, variant.Variant_name, variant.Quantity, variant.Product_id, now, variant.ID)
	if err != nil {
		panic(err)
	}
	return variant
}

func getProductWithVariant(product_id int) ProductWithVariants {
	var pv ProductWithVariants
	sqlStmt := `
		SELECT p.id, p.name, p.created_at, p.updated_at, v.id, v.variant_name, v.quantity, v.created_at, v.updated_at, v.product_id
		FROM product AS p
		LEFT JOIN variant AS v ON p.id = v.product_id
		WHERE p.id=?;
	`
	rows, err := db.Query(sqlStmt, product_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		var v Variant
		err := rows.Scan(&p.ID, &p.Name, &p.Created_at, &p.Updated_at, &v.ID, &v.Variant_name, &v.Quantity, &v.Created_at, &v.Updated_at, &v.Product_id)
		if err != nil {
			panic(err)
		}
		pv.Product = p
		v.Product = p
		pv.Variants = append(pv.Variants, v)
	}
	return pv
}

func deleteVariantById(vID int) error {
	sqlstmt := `
	DELETE FROM variant WHERE id = ?
	`
	_, err = db.Exec(sqlstmt, vID)
	if err == nil {
		fmt.Printf("Variant %v has been deleted!\n", vID)
	}
	return err
}
