package handlers

import (
	"fmt"
	"net/http"
	"time"

	db "transactions/app/database"
	model "transactions/app/models"

	"github.com/relvacode/iso8601"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Success":    true,
		"StatusCode": "200",
		"Msg":        "Db connected and we are good to go.",
	})
}

func Add(c *gin.Context) {
	var requestBody TransactionRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Error":   "Error in input.",
		})
		return
	}

	currentTime := time.Now().UTC()

	reqTime, err := iso8601.ParseString(requestBody.Timestamp)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"Success": false,
			"Error":   "Error during parsing timestamp.",
		})
		return
	}

	if currentTime.After(reqTime) {
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"Success": false,
			"Error":   "Timestamp is in future.",
		})
		return
	}

	diff := currentTime.Sub(reqTime).Seconds()
	fmt.Println("Time Difference", diff)
	if diff > 60 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	trans := &model.Transactions{
		Amount:    requestBody.Amount,
		Timestamp: requestBody.Timestamp,
		// CreatedAt: time.Now().UTC(),
	}

	res := db.Conn.Create(trans)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Error":   res.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Success": true,
		"Data":    res,
	})

}

func Destroy(c *gin.Context) {
	transObj := []model.Transactions{}
	res := db.Conn.Delete(&transObj)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Error":   res.Error,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Msg":     "All transaction are deleted.",
	})
}

func SetLocation(c *gin.Context) {
	var requestBody LocationSetRequest
	transObj := model.Transactions{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Error":   "Error in input.",
		})
		return
	}
	id := c.Param("id")
	check := db.Conn.Where("id = ?", id).First(&transObj)
	if check.RowsAffected == 1 {
		res := db.Conn.Model(&transObj).Where("id = ?", c.Param("id")).Update("city", requestBody.Location)
		if res.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Success": false,
				"Error":   res.Error,
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Success": false,
			"Error":   "No Data found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Msg":     check,
	})
}

func ResetLocation(c *gin.Context) {
	transObj := model.Transactions{}
	id := c.Param("id")
	check := db.Conn.Where("id = ?", id).First(&transObj)
	if check.RowsAffected == 1 {
		res := db.Conn.Model(&transObj).Where("id = ?", c.Param("id")).Update("city", "")
		if res.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Success": false,
				"Error":   res.Error,
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Success": false,
			"Error":   "No Data found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Msg":     check,
	})
}

func Stats(c *gin.Context) {
	transObj := []model.DbResult{}
	filter := c.Request.Header.Get("Filter")
	if filter != "" {
		db.Conn.Raw("select count(id) as count, sum(amount) as sum, avg(amount) as avg, min(amount) as min, max(amount) as max from transactions where created_at > now() - interval 60 second and city = ?", filter).Scan(&transObj)
	} else {
		db.Conn.Raw("select count(id) as count, SUM(amount) as sum, AVG(amount) as avg, MIN(amount) as min, MAX(amount) as max from transactions where created_at > now() - interval 60 second").Scan(&transObj)
	}
	c.JSON(http.StatusOK, gin.H{
		"Success": true,
		"data":    transObj,
	})
}
