package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPatternCustomer(t *testing.T) {
	g := NewGomegaWithT(t)

	fixture := []string{
		"A000000",
		"A5896487",
		"B5896481",
		"C5896486",
		"S5896465",
		"L589645",
		"L589645489",
		"L58964556",
		"M589645",
		"M589645489",
		"M58964556",
		"H589645",
		"H589645489",
		"H58964556",
	}

	for _, pattern := range fixture {
		customer := Customer{
			Name:       "ธารภิรมย์",
			Email:      "thanphirom@gmail.com",
			CustomerID: pattern, //ผิด
		}
		ok, err := govalidator.ValidateStruct(customer)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("รหัสลูกค้าต้องขึ้นต้นด้วย L|M|H ตามด้วยตัวเลขจำนวน 7 ตัว"))
	}
}
