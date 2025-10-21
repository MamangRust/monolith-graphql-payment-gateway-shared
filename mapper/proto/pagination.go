package protomapper

import pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"

// mapPaginationMeta maps a sharedPagination.PaginationMeta to a protobuf PaginationMeta.
//
// Args:
//   - s: A pointer to a sharedPagination.PaginationMeta to be mapped.
//
// Returns:
//   - A pointer to a protobuf PaginationMeta containing the mapped data.
func MapPaginationMeta(s *pb.PaginationMeta) *pb.PaginationMeta {
	return &pb.PaginationMeta{
		CurrentPage:  int32(s.CurrentPage),
		PageSize:     int32(s.PageSize),
		TotalPages:   int32(s.TotalPages),
		TotalRecords: int32(s.TotalRecords),
	}
}
