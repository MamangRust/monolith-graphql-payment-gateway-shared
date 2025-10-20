package userprotomapper

type userProtoMapper struct {
	UserQueryProtoMapper   UserQueryProtoMapper
	UserCommandProtoMapper UserCommandProtoMapper
}

func NewUserProtoMapper() *userProtoMapper {
	return &userProtoMapper{
		UserQueryProtoMapper:   NewUserQueryProtoMapper(),
		UserCommandProtoMapper: NewUserCommandProtoMapper(),
	}
}
