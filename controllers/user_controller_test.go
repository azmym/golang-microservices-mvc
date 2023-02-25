package controllers

//
//func init() {
//	domain.UserRepository = &services.UserRepositoryMock{}
//}
//
//func TestGetUserIfUserNotExist(t *testing.T) {
//	//given
//	services.FindUserByIdFunc = func(userId int64) (*domain.User, *utils.ApplicationError) {
//		return nil, &utils.ApplicationError{
//			Message:    fmt.Sprintf("there is no user with id= %v", userId),
//			StatusCode: http.StatusNotFound,
//			Code:       "not_found",
//		}
//	}
//	req := httptest.NewRequest("GET", "localhost:8080/users?user_id=111", nil)
//	w := httptest.NewRecorder()
//
//	//when
//	UserController.GetUser(w, req)
//
//	//then
//	result := w.Result()
//	var body utils.ApplicationError
//	bytes, _ := io.ReadAll(result.Body)
//	err := json.Unmarshal(bytes, &body)
//	if err != nil {
//		t.Error(err)
//	}
//
//	assert.Equal(t, "404 Not Found", result.Status)
//	assert.Equal(t, "there is no user with id= 111", body.Message)
//	assert.Equal(t, http.StatusNotFound, body.StatusCode)
//	assert.Equal(t, "not_found", body.Code)
//
//}
//
//func TestGetUserIfUserExist(t *testing.T) {
//	//given
//	services.FindUserByIdFunc = func(userId int64) (*domain.User, *utils.ApplicationError) {
//		return &domain.User{
//			Id:        111,
//			FirstName: "Test",
//			LastName:  "Case",
//			Email:     "test@case.de",
//		}, nil
//	}
//	req := httptest.NewRequest("GET", "localhost:8080/users?user_id=111", nil)
//	w := httptest.NewRecorder()
//
//	//when
//	UserController.GetUser(w, req)
//
//	//then
//	result := w.Result()
//	var body domain.User
//	bytes, _ := io.ReadAll(result.Body)
//	err := json.Unmarshal(bytes, &body)
//	if err != nil {
//		t.Error(err)
//	}
//
//	assert.Equal(t, "200 OK", result.Status)
//	assert.Equal(t, int64(111), body.Id)
//	assert.Equal(t, "Test", body.FirstName)
//	assert.Equal(t, "Case", body.LastName)
//	assert.Equal(t, "test@case.de", body.Email)
//
//}
//
//func TestGetUserIfUserNotSend(t *testing.T) {
//	req := httptest.NewRequest("GET", "localhost:8080/users", nil)
//	w := httptest.NewRecorder()
//
//	//when
//	UserController.GetUser(w, req)
//
//	//then
//	result := w.Result()
//	var body utils.ApplicationError
//	bytes, _ := io.ReadAll(result.Body)
//	err := json.Unmarshal(bytes, &body)
//	if err != nil {
//		t.Error(err)
//	}
//
//	assert.Equal(t, "400 Bad Request", result.Status)
//	assert.Equal(t, "user_id should be int.", body.Message)
//	assert.Equal(t, http.StatusBadRequest, body.StatusCode)
//	assert.Equal(t, "bad_request", body.Code)
//
//}
