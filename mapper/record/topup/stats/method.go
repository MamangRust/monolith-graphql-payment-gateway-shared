package topupstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// topupStatisticMethodRecordMapper provides methods to map topup method rows to domain models.
type topupStatisticMethodRecordMapper struct {
}

// NewTopupStatisticMethodRecordMapper returns a new instance of TopupStatisticMethodMapper which provides methods to map topup method rows to domain models.
func NewTopupStatisticMethodRecordMapper() TopupStatisticMethodMapper {
	return &topupStatisticMethodRecordMapper{}
}

// ToTopupMonthlyMethod maps a GetMonthlyTopupMethodsRow database row to a TopupMonthMethod domain model.
// It takes a pointer to a GetMonthlyTopupMethodsRow as a parameter and returns a pointer to a TopupMonthMethod.
// The function maps the fields of GetMonthlyTopupMethodsRow to the corresponding fields of TopupMonthMethod.
func (t *topupStatisticMethodRecordMapper) ToTopupMonthlyMethod(s *db.GetMonthlyTopupMethodsRow) *record.TopupMonthMethod {
	return &record.TopupMonthMethod{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyMethods maps a slice of GetMonthlyTopupMethodsRow database rows to a slice of TopupMonthMethod domain models.
// It iterates over the provided slice, converting each element using the ToTopupMonthlyMethod method,
// and appends the result to a new slice. The function returns a slice of pointers to TopupMonthMethod.
func (t *topupStatisticMethodRecordMapper) ToTopupMonthlyMethods(s []*db.GetMonthlyTopupMethodsRow) []*record.TopupMonthMethod {
	var topupRecords []*record.TopupMonthMethod

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupMonthlyMethod(topup))
	}

	return topupRecords
}

// ToTopupYearlyMethod maps a GetYearlyTopupMethodsRow database row to a TopupYearlyMethod domain model.
// It takes a pointer to a GetYearlyTopupMethodsRow as a parameter and returns a pointer to a TopupYearlyMethod.
// The function maps the fields of GetYearlyTopupMethodsRow to the corresponding fields of TopupYearlyMethod.
func (t *topupStatisticMethodRecordMapper) ToTopupYearlyMethod(s *db.GetYearlyTopupMethodsRow) *record.TopupYearlyMethod {
	return &record.TopupYearlyMethod{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyMethods maps a slice of GetYearlyTopupMethodsRow database rows to a slice of TopupYearlyMethod domain models.
// It iterates over the provided slice, converting each element using the ToTopupYearlyMethod method,
// and appends the result to a new slice. The function returns a slice of pointers to TopupYearlyMethod.
func (t *topupStatisticMethodRecordMapper) ToTopupYearlyMethods(s []*db.GetYearlyTopupMethodsRow) []*record.TopupYearlyMethod {
	var topupRecords []*record.TopupYearlyMethod

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupYearlyMethod(topup))
	}

	return topupRecords
}
