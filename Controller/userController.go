package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSahams(c *gin.Context) {

	db := Connect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM saham")
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	var saham Saham
	var sahams []Saham

	for results.Next() {
		err = results.Scan(&saham.ID, &saham.Name, &saham.Price)
		if err != nil {
			panic(err.Error())
		}
		sahams = append(sahams, saham)
	}
	if len(sahams) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, sahams)
	}
}

func AddSahams(c *gin.Context) {
	db := Connect()
	defer db.Close()

	var saham Saham

	if err := c.Bind(&saham); err != nil {
		fmt.Println(err)
		return
	}
	db.Exec("INSERT INTO saham (id, name, price) VALUES (?,?,?)", saham.ID, saham.Name, saham.Price)

	c.IndentedJSON(http.StatusOK, saham)
}

func UpdateSaham(c *gin.Context) {
	db := Connect()
	defer db.Close()

	var saham Saham

	if err := c.Bind(&saham); err != nil {
		fmt.Println(err)
		return
	}

	result, errQuery := db.Exec("UPDATE saham SET name=?, price=? WHERE id=?", saham.ID, saham.Name, saham.Price)

	num, _ := result.RowsAffected()

	if errQuery == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Failed to Update")
			return
		} else {
			c.IndentedJSON(http.StatusOK, saham)
		}
	}
}

func DeleteSaham(c *gin.Context) {
	db := Connect()
	defer db.Close()

	id := c.Query("id")

	result, errQuery := db.Exec("DELETE FROM saham WHERE id=?", id)

	num, _ := result.RowsAffected()

	if errQuery == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Failed to Delete")
		} else {
			c.IndentedJSON(http.StatusOK, id)
		}
	}
}
