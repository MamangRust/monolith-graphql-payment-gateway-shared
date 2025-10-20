package userservicemapper

type UserResponseMapper interface {
	QueryMapper() UserQueryResponseMapper
	CommandMapper() UserCommandResponseMapper
}

type userResponseMapper struct {
	queryMapper   UserQueryResponseMapper
	commandMapper UserCommandResponseMapper
}

func NewUserResponseMapper() UserResponseMapper {
	return &userResponseMapper{
		queryMapper:   NewUserQueryResponseMapper(),
		commandMapper: NewUserCommandResponseMapper(),
	}
}

func (m *userResponseMapper) QueryMapper() UserQueryResponseMapper {
	return m.queryMapper
}

func (m *userResponseMapper) CommandMapper() UserCommandResponseMapper {
	return m.commandMapper
}
