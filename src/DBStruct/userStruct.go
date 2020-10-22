package DBStruct

func (UserTable) TableName() string {
	return "user"
}


type UserTable struct {
	ID				    uint64		`json:"id"                form:"id"                query:"id"                gorm:"column:id;primary_key"                                                                               `
	userName	 		string		`json:"user_name"         form:"user_name"         query:"user_name"         gorm:"column:user_name;type:varchar(255)"                                    validate:"required"           `
	Password		    string		`json:"password"          form:"password"          query:"password"          gorm:"column:password;type:varchar(255)"                                     validate:"required"           `
	Email			    string		`json:"email"             form:"email"             query:"email"             ggorm:"column:email;type:varchar(255)"                                       validate:"required"           `
}
