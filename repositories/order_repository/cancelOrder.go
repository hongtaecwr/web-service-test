package order_repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
	"web-service-test/entities/database_entity"
)

func (orderRepo OrderRepository) CancelOrder(id, userId int) (bool, error) {

	transaction := orderRepo.db.Begin()
	order := new(database_entity.Orders)
	insertOrder := transaction.Table("orders").
		Where("orders.id = ? AND orders.user_id = ? AND order_status_id NOT IN (?)", id, userId, database_entity.IDStatusCancel).
		Updates(map[string]interface{}{
			"order_status_id": database_entity.IDStatusCancel,
			"cancel_date":     time.Now(),
			"updated_by":      strconv.Itoa(userId),
			"updated_at":      time.Now(),
		})
	if insertOrder.RowsAffected == 0 {
		transaction.Rollback()
		return false, errors.New("no update data")
	}

	orderProducts := new([]database_entity.OrderProducts)
	findOrderProducts := transaction.Find(&orderProducts)
	if findOrderProducts.RowsAffected == 0 {
		transaction.Rollback()
		return false, errors.New("product not found")
	}

	for _, data := range *orderProducts {
		updateProductStock := transaction.
			Table("products").
			Where("products.id = ?", data.ProductId).
			Updates(map[string]interface{}{
				"stock":      gorm.Expr("stock + ?", 1),
				"sold":       gorm.Expr("sold - ?", 1),
				"updated_by": strconv.Itoa(userId),
			})
		if updateProductStock.RowsAffected == 0 {
			transaction.Rollback()
			return false, errors.New("no stock update")
		}
	}
	orderLog := &database_entity.OrderLogs{
		RefCode:       order.RefCode,
		UserId:        userId,
		AddressId:     order.AddressId,
		OrderStatusId: database_entity.IDStatusCancel,
		SummaryPrice:  order.SummaryPrice,
		CreatedBy:     strconv.Itoa(userId),
		UpdatedBy:     strconv.Itoa(userId),
		OrderId:       id,
	}
	insertOrderLog := transaction.Create(&orderLog)
	if insertOrderLog.RowsAffected == 0 {
		transaction.Rollback()
		return false, errors.New("no create log data")
	}

	transaction.Commit()
	return true, nil
}
