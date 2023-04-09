package controllers

import (
	"io/ioutil"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestSatu(c *gin.Context) {
	var dataSatu struct {
		Data int `json:"data"`
	}
	var result []int

	if c.Bind(&dataSatu) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data failed"})
		return
	}

	for i := 2; i < dataSatu.Data; i++ {
		if BilanganPrima(i) {
			result = append(result, i)
		}
	}

	c.JSON(http.StatusOK, gin.H{"bilangan prima nya adalah": result})
}

func BilanganPrima(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func TestDua(c *gin.Context) {
	file, _ := c.FormFile("data")

	content, err := ioutil.ReadFile(file.Filename)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to read file"})
		return
	}
	c.JSON(200, gin.H{"terbaca : ": string(content)})
}

func TestTiga(c *gin.Context) {
	var fiboGenap []int
	n1, n2 := 1, 2
	for n2 < 4000000 {
		if n2%2 == 0 {
			fiboGenap = append(fiboGenap, n2)
		}
		n1, n2 = n2, n1+n2
	}

	c.JSON(200, gin.H{"fibonaci genap : ": fiboGenap})
}

func TestEmpat(c *gin.Context) {

	data := c.Request.FormValue("data")

	palindrome := checkPalindrome(data)

	c.JSON(200, gin.H{"data yang di input : ": data, "apakah palindrome ?": palindrome})
}

func checkPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func TestLima(c *gin.Context) {
	var dataLima struct {
		Data int `json:"data"`
	}
	if c.Bind(&dataLima) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data failed"})
		return
	}

	if dataLima.Data < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data harus nilai positif"})
		return
	}

	c.JSON(200, gin.H{"data yang di input : ": dataLima.Data, "nilai faktorialnya ?": faktorial(dataLima.Data)})
}

func faktorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}
