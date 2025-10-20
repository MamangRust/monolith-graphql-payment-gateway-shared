package topupstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TopupStatisticMethodByCardNumberMapper maps GetMonthlyTopupMethodsByCardNumberRow database rows to TopupMonthMethod domain models.
type topupStatisticMethodByCardNumberMapper struct {
}

// NewTopupStatisticMethodByCardNumberMapper returns a new instance of TopupStatisticMethodByCardNumberMapper
// which provides methods to map GetMonthlyTopupMethodsByCardNumberRow and GetYearlyTopupMethodsByCardNumberRow
// database rows to TopupMonthMethod and TopupYearlyMethod domain models.
func NewTopupStatisticMethodByCardNumberMapper() TopupStatisticMethodByCardNumberMapper {
	return &topupStatisticMethodByCardNumberMapper{}
}

// ToTopupMonthlyMethodByCardNumber maps a GetMonthlyTopupMethodsByCardNumberRow database row
// to a TopupMonthMethod domain model. It takes a pointer to a
// GetMonthlyTopupMethodsByCardNumberRow as a parameter and returns a pointer to a
// TopupMonthMethod. The function maps the fields of GetMonthlyTopupMethodsByCardNumberRow
// to the corresponding fields of TopupMonthMethod.
func (t *topupStatisticMethodByCardNumberMapper) ToTopupMonthlyMethodByCardNumber(s *db.GetMonthlyTopupMethodsByCardNumberRow) *record.TopupMonthMethod {
	return &record.TopupMonthMethod{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyMethodsByCardNumber maps a slice of GetMonthlyTopupMethodsByCardNumberRow database rows
// to a slice of TopupMonthMethod domain models. It iterates over the provided slice, converting
// each element using the ToTopupMonthlyMethodByCardNumber method, and appends the result to a new
// slice. The function returns a slice of pointers to TopupMonthMethod.
func (t *topupStatisticMethodByCardNumberMapper) ToTopupMonthlyMethodsByCardNumber(s []*db.GetMonthlyTopupMethodsByCardNumberRow) []*record.TopupMonthMethod {
	var topupRecords []*record.TopupMonthMethod

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupMonthlyMethodByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupYearlyMethodByCardNumber maps a GetYearlyTopupMethodsByCardNumberRow database row to a TopupYearlyMethod domain model.
// It takes a pointer to a GetYearlyTopupMethodsByCardNumberRow as a parameter and returns a pointer to a TopupYearlyMethod.
// The function maps the fields of GetYearlyTopupMethodsByCardNumberRow to the corresponding fields of TopupYearlyMethod.
func (t *topupStatisticMethodByCardNumberMapper) ToTopupYearlyMethodByCardNumber(s *db.GetYearlyTopupMethodsByCardNumberRow) *record.TopupYearlyMethod {
	return &record.TopupYearlyMethod{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyMethodsByCardNumber maps a slice of GetYearlyTopupMethodsByCardNumberRow database rows
// to a slice of TopupYearlyMethod domain models. It iterates over the provided slice, converting
// each element using the ToTopupYearlyMethodByCardNumber method, and appends the result to a new
// slice. The function returns a slice of pointers to TopupYearlyMethod.
func (t *topupStatisticMethodByCardNumberMapper) ToTopupYearlyMethodsByCardNumber(s []*db.GetYearlyTopupMethodsByCardNumberRow) []*record.TopupYearlyMethod {
	var topupRecords []*record.TopupYearlyMethod

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupYearlyMethodByCardNumber(topup))
	}

	return topupRecords
}
