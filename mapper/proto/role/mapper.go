package roleprotomapper

type roleProtoMapper struct {
	RoleQueryProtoMapper   RoleQueryProtoMapper
	RoleCommandProtoMapper RoleCommandProtoMapper
}

func NewRoleProtoMapper() *roleProtoMapper {
	return &roleProtoMapper{
		RoleQueryProtoMapper:   NewRoleQueryProtoMapper(),
		RoleCommandProtoMapper: NewRoleCommandProtoMapper(),
	}
}
