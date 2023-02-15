package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	customer := Customer{
		Name:       "", //ผิด
		Email:      "thanphirom@gmail.com",
		CustomerID: "L6223045",
	}
	ok, err := govalidator.ValidateStruct(customer)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("กรุณากรอกข้อมูลชื่อ"))
}
