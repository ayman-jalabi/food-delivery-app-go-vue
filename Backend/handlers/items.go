package handlers

import (
	"main/repos"
)

type ItemHandler struct {
	Repo *repos.ItemRepo
}

func NewItemHandler(repo *repos.ItemRepo) ItemHandler {
	return ItemHandler{Repo: repo}
}

//func (ih ItemHandler) GetAllItems(writer http.ResponseWriter, _ *http.Request) {
//	allItems, err := ih.Repo.GetAll()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	byteSlice, err := json.Marshal(allItems)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	writer.Header().Set("Content-Type", "application/json")
//	writer.Write(byteSlice)
//}
