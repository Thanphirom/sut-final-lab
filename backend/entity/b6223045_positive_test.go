package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestAllCustomerCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	customer := Customer{
		Name:       "ธารภิรมย์",
		Email:      "thanphirom@gmail.com",
		CustomerID: "L6223045",
	}
	ok, err := govalidator.ValidateStruct(customer)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
