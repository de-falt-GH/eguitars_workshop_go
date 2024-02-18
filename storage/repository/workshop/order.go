package workshop

import (
	"errors"
	"kursarbeit/storage/models/workshop"
	common_repo "kursarbeit/storage/repository/common"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderRepo struct {
	common_repo.CommonRepo
}

func (repo *OrderRepo) GetAll(ctx *gin.Context) []workshop.Order {
	var orders []workshop.Order
	query := repo.DB.InnerJoins("Customer").
		InnerJoins("Master").
		InnerJoins("OrderType").
		InnerJoins("OrderStatus").
		Joins("Guitar").
		Preload("RequiredComponents.Component")

	searchBy := ctx.Query("search_by")
	search := ctx.Query("search")
	if search != "" {
		if searchBy == "master_id" {
			query = query.Where("\"order\".master_id = ?", search)
		} else if searchBy == "customer_id" {
			query = query.Where("\"order\".customer_id = ?", search)
		} else if searchBy == "order_status_id" {
			query = query.Where("\"order\".order_status_id = ?", search)
		} else if searchBy == "order_type_id" {
			query = query.Where("\"order\".order_type_id = ?", search)
		}
	}

	sort_by := ctx.Query("sort_by")
	if sort_by == "master_id" || sort_by == "customer_id" ||
		sort_by == "order_status_id" || sort_by == "order_type_id" ||
		sort_by == "id" {
		order := "\"order\"." + sort_by

		if ctx.Query("desc") == "true" {
			order += " DESC"
		}
		query = query.Order(order)
	}

	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	query.Offset(offset).Limit(limit).Find(&orders)

	return orders
}

func (repo *OrderRepo) GetAllComponents() []workshop.Component {
	var components []workshop.Component
	repo.DB.Find(&components)

	return components
}

func (repo *OrderRepo) GetByID(ID int) *workshop.Order {
	var order workshop.Order
	query := repo.DB.InnerJoins("Customer").
		InnerJoins("Master").
		InnerJoins("OrderType").
		// InnerJoins("OrderStatus").
		Joins("Guitar").
		Preload("RequiredComponents.Component")
	err := query.Where(`order.id`, ID).First(&order).Error
	if err != nil {
		return nil
	}

	return &order
}

func (repo *OrderRepo) GetByCustomerID(ID int) []workshop.Order {
	var orders []workshop.Order
	repo.DB.Where("customer_id = ?", ID).Find(orders)

	return orders
}

func (repo *OrderRepo) GetByMasterID(ID int) []workshop.Order {
	var orders []workshop.Order
	repo.DB.Where("master_id = ?", ID).Find(orders)

	return orders
}

func (repo *OrderRepo) DeleteByID(ID int) error {
	db := repo.DB.Delete(&workshop.Order{}, ID)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected < 1 {
		return errors.New("entries to delete not found")
	}

	return nil
}

func (repo *OrderRepo) GetDB() *gorm.DB {
	return repo.DB
}
