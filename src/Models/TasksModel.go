package Models

import (
	"../Sys"
	"../DBStruct"
)

func addNewTask(taskName string) bool  {
	Conn :=  Sys.DB()
	defer Conn.Close()
	Task  := DBStruct.TasksTable{}
	Task.Task = taskName
	err := Conn.Create(&Task).Error
	Conn.Close()
	if err != nil{
		return  false
	}


	return true
}


func editTask(taskID string, Task string) bool  {
	Conn :=  Sys.DB()
	defer Conn.Close()
	err := Conn.Model(DBStruct.TasksTable{}).Where("id = ? ", taskID).Update("task_name", Task).Error
	Conn.Close()
	if err != nil{
		return false
	}
	return true
}

func deleteTask(taskID string) bool  {
	Conn :=  Sys.DB()
	defer Conn.Close()
	err := Conn.Delete(SiteStruct.OrderItemsTable{}, "id = ?", OrderItemID).Error
	Conn.Close()
	if err != nil{
		return false
	}
	return true
}


func doneTask(taskID string, isDone bool) bool  {
	Conn :=  Sys.DB()
	defer Conn.Close()
	err := Conn.Model(DBStruct.TasksTable{}).Where("id = ? ", taskID).Update("task_name", isDone).Error
	Conn.Close()
	if err != nil{
		return false
	}
	return true
}

