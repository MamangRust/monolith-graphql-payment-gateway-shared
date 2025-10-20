package userrecord

type UserRecordMapper interface {
	QueryMapper() UserQueryRecordMapper
	CommandMapper() UserCommandRecordMapper
}

type userRecordMapper struct {
	queryMapper   UserQueryRecordMapper
	commandMapper UserCommandRecordMapper
}

func NewUserRecordMapper() UserRecordMapper {
	return &userRecordMapper{
		queryMapper:   NewUserQueryRecordMapper(),
		commandMapper: NewUserCommandRecordMapper(),
	}
}

func (u *userRecordMapper) QueryMapper() UserQueryRecordMapper {
	return u.queryMapper
}

func (u *userRecordMapper) CommandMapper() UserCommandRecordMapper {
	return u.commandMapper
}
