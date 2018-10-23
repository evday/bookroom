package service

import (
	"sync"
	"bookroom/models"
)

func ListUser(username string,offset,limit int)([]*models.User,int64,error){
	infos := make([]*models.User,0)
	
	users,count,err := models.ListUser(username,offset,limit)
	if err != nil {
		return nil,count,err
	}

	ids := []int64{}
	for _,user := range users {
		ids = append(ids,user.ID)
	}

	wg := sync.WaitGroup{}
	userList := models.UserList{
		Lock: new(sync.Mutex),
		IdMap: make(map[int64]*models.User,len(users)),
	}

	errChan := make(chan error,1)
	finished := make(chan bool,1)

	for _,u := range users {
		wg.Add(1)
		go func(u *models.User) {
			defer wg.Done()

			userList.Lock.Lock()
			defer userList.Lock.Unlock()

			userList.IdMap[u.ID] = &models.User{
				Model:models.Model{ID:u.ID,CreateAt:u.CreateAt},
				Name:u.Name,
				Password:u.Password,
				IsAdmin:u.IsAdmin,
				
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}