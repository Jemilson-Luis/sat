package models

type Task struct {
	Id string `json:id`
	FkUser string `json:fk_user`
	Task string `json:task`
	Desc string `json:desc`
	Status string `json:status`
	Date string `json:date`
	Script string `json:script`
	From string `json:from`
	To string `json:to`
}

type TaskDTOInput struct {

}
