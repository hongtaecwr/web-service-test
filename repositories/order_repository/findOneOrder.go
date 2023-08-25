package order_repository

import (
	"web-service-test/entities/orders_entity"
)

func (orderRepo OrderRepository) OrderDetail(id, userId int) (bool, error, *orders_entity.ResponseOrderDetail) {
	orderDetail := new(orders_entity.ResponseOrderDetail)
	findOneOrder := orderRepo.db.Debug().
		Table("orders").
		Select([]string{
			"orders.id",
			"orders.ref_code",
			"CONCAT(users.first_name, ' ', users.last_name) AS user_name",
			"addresses.name AS address_name",
			"addresses.address_no",
			"addresses.address_moo",
			"addresses.address_road",
			"addresses.address_sub_district",
			"addresses.address_district",
			"addresses.address_province",
			"addresses.address_zip_code",
			"addresses.address_tel",
			"orders.completed_date",
			"orders.shipped_date",
			"orders.cancel_date",
			"order_statuses.status AS order_status",
			"orders.summary_price",
			"orders.created_at",
			"orders.updated_at",
			"orders.updated_by",
			"orders.deleted_by",
		}).
		Where("orders.id = ? AND orders.user_id = ?", id, userId).
		Joins("INNER JOIN users ON users.id = orders.user_id").
		Joins("INNER JOIN addresses ON addresses.user_id = users.id").
		Joins("INNER JOIN order_statuses ON order_statuses.id = orders.order_status_id").
		Scan(&orderDetail)
	if findOneOrder.RowsAffected == 0 {
		return false, findOneOrder.Error, nil
	}
	return true, nil, orderDetail
}
