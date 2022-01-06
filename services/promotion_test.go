package services

import (
	"errors"
	"gotest/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotion(t *testing.T) {

	type testcase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	testCases := []testcase{
		{name: "applied 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "applied 200", purchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
		{name: "applied 300", purchaseMin: 100, discountPercent: 20, amount: 300, expected: 240},
		{name: "not apply 50", purchaseMin: 100, discountPercent: 20, amount: 50, expected: 50},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			//Arrage
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)
			promoService := NewPromotionService(promoRepo)

			//Act
			discount, _ := promoService.CalculateDiscount(c.amount)

			//Assert
			assert.Equal(t, c.expected, discount)
		})
	}

	t.Run("TestErrorZeroAmount", func(t *testing.T) {
		//Arrage
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, nil)
		promoService := NewPromotionService(promoRepo)

		//Act
		_, err := promoService.CalculateDiscount(0)

		//Assert
		assert.ErrorIs(t, ErrZeroAmount, err)
		promoRepo.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("TestErrorRepository", func(t *testing.T) {
		//Arrage
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New("Can't connect database"))
		promoService := NewPromotionService(promoRepo)

		//Act
		_, err := promoService.CalculateDiscount(50)

		//Assert
		assert.ErrorIs(t, ErrRepository, err)
	})

}
