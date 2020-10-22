package DBStruct

func (TasksTable) TableName() string {
	return "tasks"
}


type TasksTable struct {
	ID				    uint64		`json:"id"                form:"id"                query:"id"                gorm:"column:id;primary_key"                                                                               `
	Task	 			string		`json:"task_name"         form:"task_name"         query:"task_name"         gorm:"column:task_name;type:varchar(255)"                                    validate:"required"           `
	Done		    	bool		`json:"done"              form:"done"              query:"done"              gorm:"column:done"                                                           validate:"required"           `
}
