package rolerecord

type roleRecordMapper struct {
	queryMapper   RoleQueryRecordMapping
	commandMapper RoleCommandRecordMapping
}

type RoleRecordMapper interface {
	QueryMapper() RoleQueryRecordMapping
	CommandMapper() RoleCommandRecordMapping
}

func NewRoleRecordMapper() RoleRecordMapper {
	return &roleRecordMapper{
		queryMapper:   NewRoleQueryRecordMapping(),
		commandMapper: NewRoleCommandRecordMapping(),
	}
}

func (r *roleRecordMapper) QueryMapper() RoleQueryRecordMapping {
	return r.queryMapper
}

func (r *roleRecordMapper) CommandMapper() RoleCommandRecordMapping {
	return r.commandMapper
}
