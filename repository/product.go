package repository

import (
	"nutech/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProduct() ([]models.Products,error)
	CreateProduct(products models.Products)(models.Products,error)
	UpdateProduct(products models.Products) (models.Products,error)
	DeleteProduct(products models.Products,ID int) (models.Products,error)
	GetProduct(ID int) (models.Products, error)
}
type repository struct {
	db *gorm.DB
}
func RepositoryProduct(db *gorm.DB) *repository{
	return &repository{db}
}
func(r *repository) FindProduct()([]models.Products,error){
	var products []models.Products
	err := r.db.Find(&products).Error

	return products,err
}
func (r *repository) GetProduct(ID int) (models.Products, error) {
	var products models.Products

	err := r.db.First(&products, ID).Error

	return products, err
}
func (r *repository)CreateProduct(products models.Products) (models.Products,error){
err := r.db.Create(&products).Error

return products,err
}
func (r *repository)DeleteProduct(products models.Products,ID int) (models.Products,error){
	err := r.db.Delete(&products).Error

	return products,err
}
func (r *repository) UpdateProduct(products models.Products) (models.Products, error) {
	err := r.db.Save(&products).Error

	return products, err
}

