package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"test-golang/config/mysql"
	"test-golang/models"

	"github.com/gin-gonic/gin"
)

var dataCreateProduct struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Price       int64  `form:"price"`
	Image       string
}

var dataUpdateProduct struct {
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
	Price       int64  `form:"price" json:"price"`
	Image       string
}

func ProductIndex(c *gin.Context) {
	var products []models.Product
	var page, limit int
	// paramater page
	page, _ = strconv.Atoi(c.Param("page"))
	// jumlah data setiap halaman
	limit = 10
	// offset data
	offset := (page - 1) * limit
	// query berdasarkan "Name / Price"
	sortBy := c.Query("sort")
	// query "asc" atau "desc"
	order := c.Query("order")
	if sortBy == "" {
		sortBy = "id" // default id
	}
	if order == "" {
		order = "asc" // default asc
	}
	// mengambil data produk dari database berdasarkan limit dan offset
	mysql.DB.Offset(offset).Limit(limit).Order(sortBy + " " + order).Find(&products)

	baseURL := "http://localhost:8080/images/"
	for i, product := range products {
		products[i].Image = baseURL + product.Image
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func ProductShow(c *gin.Context) {
	var product models.Product

	// paramater page
	id, _ := strconv.Atoi(c.Param("id"))

	mysql.DB.First(&product, id)
	// .Offset(offset).Limit(limit).Order(sortBy + " " + order).Find(&products)

	baseURL := "http://localhost:8080/images/"
	product.Image = baseURL + product.Image
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func ProductCreate(c *gin.Context) {
	// Bind data dari request ke struct Product
	if err := c.Bind(&dataCreateProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Simpan gambar ke dalam folder "images"
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read image file"})
		return
	}
	// ganti nama gambar
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)

	if err := c.SaveUploadedFile(file, "images/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image file"})
		return
	}

	productsVal := models.Product{Name: dataCreateProduct.Name, Description: dataCreateProduct.Description, Price: dataCreateProduct.Price, Image: filename}
	// product.Image = filename

	// Simpan data produk ke dalam database
	if err := mysql.DB.Create(&productsVal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save product to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": productsVal})
}

func ProductUpdate(c *gin.Context) {

	// Set Content-Type header to multipart/form-data
	// c.Request.Header.Set("Content-Type", "multipart/form-data")

	// Get model if exist
	id, _ := strconv.Atoi(c.Param("id"))

	// var filename string
	var product models.Product

	if err := mysql.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// data binding dari form
	if err := c.ShouldBind(&dataUpdateProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	file, _ := c.FormFile("image")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read image file"})
	// 	return
	// }

	// jika ada file baru yang diupload
	if file != nil {
		// hapus file lama jika ada
		if err := os.Remove("images/" + product.Image); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove old image file"})
		}
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read image file"})
		// 	return
		// }
		// ganti nama gambar
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)

		// simpan gambar
		if err := c.SaveUploadedFile(file, "images/"+filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image file"})
			return
		}

		dataUpdateProduct.Image = filename
	} else {
		dataUpdateProduct.Image = product.Image
	}

	product.Name = dataUpdateProduct.Name
	product.Description = dataUpdateProduct.Description
	product.Price = dataUpdateProduct.Price
	product.Image = dataUpdateProduct.Image

	mysql.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func ProductDelete(c *gin.Context) {
	var product models.Product
	id, _ := strconv.Atoi(c.Param("id"))

	if err := mysql.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Cek apakah file image ada atau tidak
	if _, err := os.Stat("images/" + product.Image); err == nil {
		os.Remove("images/" + product.Image)
	}

	mysql.DB.Delete(&product, id)

	c.JSON(http.StatusOK, gin.H{"data": product})
}
