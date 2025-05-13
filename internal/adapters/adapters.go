package adapters

import (
	"ecomm/internal/domain"
	"ecomm/proto"
)

func ToProtoProduct(product domain.Product) *proto.Product {
	return &proto.Product{
		Id:              product.ID,
		Name:            product.Name,
		Image:           product.Image,
		Category:        product.Category,
		Description:     product.Description,
		Rating:          int32(product.Rating),
		NumberOfReviews: int32(product.NumberOfReviews),
		Price:           float64(product.Price),
		CountInStock:    int32(product.CountInStock),
		CreatedAt:       product.CreatedAt,
		UpdatedAt:       product.UpdatedAt,
	}
}

func ToProtoCreateProductRequest(product *domain.CreateProductRequest) *proto.CreateProductRequest {
	return &proto.CreateProductRequest{
		Name:            product.Name,
		Image:           product.Image,
		Category:        product.Category,
		Description:     product.Description,
		Rating:          int32(product.Rating),
		NumberOfReviews: int32(product.NumberOfReviews),
		Price:           float64(product.Price),
		CountInStock:    int32(product.CountInStock),
	}
}

func ToProtoUpdateProductRequest(product *domain.UpdateProductRequest) *proto.UpdateProductRequest {
	return &proto.UpdateProductRequest{
		Id:              product.ID,
		Name:            product.Name,
		Image:           product.Image,
		Category:        product.Category,
		Description:     product.Description,
		Rating:          int32(product.Rating),
		NumberOfReviews: int32(product.NumberOfReviews),
		Price:           product.Price,
		CountInStock:    int32(product.CountInStock),
	}
}

func ToProtoProducts(products []*domain.Product) []*proto.Product {
	protoProducts := make([]*proto.Product, len(products))
	for i, product := range products {
		protoProducts[i] = ToProtoProduct(*product)
	}
	return protoProducts
}

func ToProtoOrder(order domain.Order) *proto.Order {
	orderItems := make([]*proto.OrderItem, len(order.OrderItems))
	for i, item := range order.OrderItems {
		orderItems[i] = &proto.OrderItem{
			Id:        item.ID,
			ProductId: item.ProductID,
			Name:      item.Name,
			Quantity:  int32(item.Quantity),
			Image:     item.Image,
			Price:     float64(item.Price),
		}
	}

	return &proto.Order{
		Id:            order.ID,
		PaymentMethod: order.PaymentMethod,
		TaxPrice:      float64(order.TaxPrice),
		ShippingPrice: float64(order.ShippingPrice),
		TotalPrice:    float64(order.TotalPrice),
		OrderItems:    orderItems,
		UserId:        order.UserID,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}
}

func ToProtoOrders(orders []*domain.Order) []*proto.Order {
	protoOrders := make([]*proto.Order, len(orders))
	for i, order := range orders {
		protoOrders[i] = ToProtoOrder(*order)
	}
	return protoOrders
}

func ToProtoCreateOrderRequest(order *domain.CreateOrderRequest) *proto.CreateOrderRequest {
	orderItems := make([]*proto.OrderItem, len(order.OrderItems))
	for i, item := range order.OrderItems {
		orderItems[i] = &proto.OrderItem{
			Id:        item.ID,
			ProductId: item.ProductID,
			Name:      item.Name,
			Quantity:  int32(item.Quantity),
			Image:     item.Image,
			Price:     float64(item.Price),
		}
	}

	return &proto.CreateOrderRequest{
		PaymentMethod: order.PaymentMethod,
		TaxPrice:      order.TaxPrice,
		ShippingPrice: order.ShippingPrice,
		TotalPrice:    order.TotalPrice,
		OrderItems:    orderItems,
		UserId:        order.UserID,
	}
}

func ToProtoUser(user domain.User) *proto.User {
	return &proto.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToProtoUsers(users []*domain.User) []*proto.User {
	protoUsers := make([]*proto.User, len(users))
	for i, user := range users {
		protoUsers[i] = ToProtoUser(*user)
	}
	return protoUsers
}

func ToProtoCreateUserRequest(user *domain.CreateUserRequest) *proto.CreateUserRequest {
	return &proto.CreateUserRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}
}

func ToProtoUpdateUserRequest(user *domain.UpdateUserRequest) *proto.UpdateUserRequest {
	return &proto.UpdateUserRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}
}

func ToProtoLoginUserRequest(req *domain.LoginRequest) *proto.LoginRequest {
	return &proto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func ToProtoRefreshTokenRequest(req *domain.RefreshAccessTokenRequest) *proto.RefreshAccessTokenRequest {
	return &proto.RefreshAccessTokenRequest{
		RefreshToken: req.RefreshToken,
	}
}
