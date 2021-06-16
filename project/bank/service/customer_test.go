package service_test

import (
	"bank/errs"
	"bank/mock/mock_repository"
	"bank/service"
	"database/sql"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetCustomerNotFound(t *testing.T) {
	//Arrage
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)

	customerID := 1

	mockRepo.EXPECT().GetById(customerID).Return(nil, sql.ErrNoRows)
	customerService := service.NewCustomerService(mockRepo)

	//Act
	_, err := customerService.GetCustomer(customerID)

	//Assert
	if err == nil {
		t.Error("should be error")
		return
	}

	appErr, ok := err.(errs.AppError)
	if !ok {
		t.Error("should return AppError")
		return
	}

	if appErr.Code != http.StatusNotFound {
		t.Error("invalid error code")
	}

	if appErr.Message != "customer not found" {
		t.Error("invalid error message")
	}
}
