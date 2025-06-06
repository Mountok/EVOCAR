package handler

import (
	"fmt"
	"net/http"
	"todoapp/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) getOrders(c *gin.Context) {
	orders, err := h.services.GetOrders()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, orders)
}

var autoTypes = map[string]string{
	"sedan":     "Легковой",
	"crossover": "Кроссовер",
	"moto":      "Мото",
	"truck":     "Грузовой",
	"cuv":       "Внедорожник",
}
var orderTypes = map[string]string{
	"delivery":  "ДОСТАВКА",
	"accident":  "АВАРИЯ",
	"breakdown": "ПОЛОМКА",
}

func (h *Handler) activeOrders(c *gin.Context) {
	orders, err := h.services.GetActiveOrders()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *Handler) createOrder(c *gin.Context) {
	var input models.Order
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateOrder(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Println("Отправка смс уведомления")
	resultUrl := fmt.Sprintf("https://sms.ru/sms/send?api_id=C1670AE3-B8D7-C652-5245-6EF19C9946A5&to=79626547706&msg=Пришел+заказ:+Тип:+%s.+Тип+авто:+%s.+Зайдите+на+сайте&json=1", orderTypes[input.TypeOfOrder], autoTypes[input.TypeOfAuto])
	resp, err := http.Post(resultUrl, "application/json", nil)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Println("Ошибка при отправке смс уведомления:", err)
		return
	}
	defer resp.Body.Close()
	logrus.Println(resp)
	logrus.Println("Смс уведомление отправлено")

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) acceptOrder(c *gin.Context) {
	id := c.Param("id")
	// Получить из тела запроса номер телефона
	phoneNumber := c.Query("phoneNumber")
	if phoneNumber == "" {
		newErrorResponse(c, http.StatusBadRequest, "номер телефона обезателен")
		return
	}

	if err := h.services.AcceptOrder(id, phoneNumber); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"message": "Заказ принят"})
}
func (h *Handler) completeOrder(c *gin.Context) {
	id := c.Param("id")
	err := h.services.CompleteOrder(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	wrapOkJSON(c, map[string]interface{}{
		"message": "Заказ выполнен",
	})

}
func (h *Handler) cancleOrder(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.CancleOrder(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	wrapOkJSON(c, map[string]interface{}{"message": "Заказ отменен"})
}

func (h *Handler) getOrdersByPhoneNumber(c *gin.Context) {
	phoneNumber := c.Query("phoneNumber")
	if phoneNumber == "" {
		newErrorResponse(c, http.StatusBadRequest, "номер телефона обезателен")
		return
	}
	orders, err := h.services.GetOrdersByPhoneNumber(phoneNumber)
	fmt.Println(orders)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if orders == nil {
		c.JSON(http.StatusOK, []models.Order{})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *Handler) getExecutorsHistory(c *gin.Context) {
	phoneNumber := c.Query("phoneNumber")
	if phoneNumber == "" {
		newErrorResponse(c, http.StatusBadRequest, "номер телефона обезателен")
		return
	}
	history, err := h.services.GetExecutorsHistory(phoneNumber)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(history) == 0 {
		c.JSON(http.StatusOK, []models.ExecutorHistory{})
		return
	}
	c.JSON(http.StatusOK, history)
}

func (h *Handler) CheckOrderStatus(c *gin.Context) {
	id := c.Param("id")
	data, err := h.services.CheckOrderStatus(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	wrapOkJSON(c, data)
}
