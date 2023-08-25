package order_repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
	"web-service-test/entities/database_entity"
	"web-service-test/entities/orders_entity"
	"web-service-test/functions"
)

func (orderRepo OrderRepository) InsertOrder(addressId, summaryPrice, userId int, productIds []int) (bool, error) {
	lastOrder := new(database_entity.Orders)

	refCode := ""
	findLastOrderRef := orderRepo.db.Select("orders.ref_code").Last(&lastOrder)
	if findLastOrderRef.RowsAffected == 0 {
		refCode = "OR-00001"
	} else {
		refCode = functions.GenerateCode(lastOrder.RefCode, "OR-")
	}

	order := &orders_entity.RequestCreateOrder{
		RefCode:       refCode,
		UserId:        userId,
		AddressId:     addressId,
		OrderStatusId: database_entity.IDStatusOrdered,
		SummaryPrice:  summaryPrice,
		CreatedBy:     strconv.Itoa(userId),
		UpdatedBy:     strconv.Itoa(userId),
	}
	transaction := orderRepo.db.Begin()
	insertOrder := transaction.Create(&order)
	if insertOrder.RowsAffected == 0 {
		transaction.Rollback()
		return false, errors.New("no create data")
	}

	for _, productId := range productIds {
		updateProductStock := transaction.
			Table("products").
			Where("products.id = ?", productId).
			Updates(map[string]interface{}{
				"stock":      gorm.Expr("stock - ?", 1),
				"sold":       gorm.Expr("sold + ?", 1),
				"updated_by": strconv.Itoa(userId),
			})
		if updateProductStock.RowsAffected == 0 {
			transaction.Rollback()
			return false, errors.New("no stock update")
		}
		orderProduct := orders_entity.CreateOrderProduct{
			OrderId:   order.ID,
			ProductId: productId,
			CreatedBy: strconv.Itoa(userId),
			UpdatedBy: strconv.Itoa(userId),
		}
		insertOrderProduct := transaction.Create(&orderProduct)
		if insertOrderProduct.RowsAffected == 0 {
			transaction.Rollback()
			return false, errors.New("no update order product data")
		}
	}
	orderLog := &database_entity.OrderLogs{
		RefCode:       refCode,
		UserId:        userId,
		AddressId:     addressId,
		OrderStatusId: database_entity.IDStatusOrdered,
		SummaryPrice:  summaryPrice,
		CreatedBy:     strconv.Itoa(userId),
		UpdatedBy:     strconv.Itoa(userId),
		OrderId:       order.ID,
	}
	insertOrderLog := transaction.Create(&orderLog)
	if insertOrderLog.RowsAffected == 0 {
		transaction.Rollback()
		return false, errors.New("no create log data")
	}

	transaction.Commit()
	return true, nil
}
