package authenticate

//var SignupDb []map[string]string
//
//func SignupData() []map[string]string {
//
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "Jo Doe",
//		"email":    "ssss@gmail.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "ali",
//		"email":    "ali@gmail.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "mahdi",
//		"email":    "mahdi@example.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "ali",
//		"email":    "aaa@example.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "reza",
//		"email":    "reza@gmail.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "milad",
//		"email":    "milad@gmail.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "danial",
//		"email":    "danial@gmail.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "mohammad",
//		"email":    "mohamad@example.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "nima",
//		"email":    "reza@example.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "nima",
//		"email":    "nima@example.com",
//		"password": "1234",
//	})
//
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "abas",
//		"email":    "abas@example.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "sanaz",
//		"email":    "sanaz@example.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "melika",
//		"email":    "melika@example.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "mina",
//		"email":    "melika@example.com",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "mina",
//		"email":    "mina@example.com",
//		"password": "1234",
//	})
//	return SignupDb
//
//}
//
//var LoginDb []map[string]string
//
//func LoginData() []map[string]string {
//
//	LoginDb = append(LoginDb, map[string]string{
//		"username": "Jo Doe",
//		"password": "1234",
//	})
//	LoginDb = append(LoginDb, map[string]string{
//		"username": "ali",
//		"password": "1234",
//	})
//	LoginDb = append(LoginDb, map[string]string{
//		"username": "mahdi",
//		"password": "1234",
//	})
//	LoginDb = append(LoginDb, map[string]string{
//		"username": "mamad",
//		"password": "1234",
//	})
//	LoginDb = append(LoginDb, map[string]string{
//		"username": "reza",
//		"password": "1234",
//	})
//	LoginDb = append(LoginDb, map[string]string{
//		"username": "milad",
//		"password": "1234",
//	})
//	LoginDb = append(LoginDb, map[string]string{
//		"username": "danial",
//		"password": "1234",
//	})
//	LoginDb = append(LoginDb, map[string]string{
//		"username": "mohammad",
//		"password": "1234",
//	})
//	LoginDb = append(LoginDb, map[string]string{
//		"username": "sara",
//		"password": "1234",
//	})
//	LoginDb = append(LoginDb, map[string]string{
//		"username": "nima",
//		"password": "1234",
//	})
//	LoginDb = append(LoginDb, map[string]string{
//		"username": "saba",
//		"password": "1234",
//	})
//	SignupDb = append(SignupDb, map[string]string{
//		"username": "mina",
//		"password": "1234",
//	})
//	return LoginDb
//
//}
//func TestCreate(t *testing.T) {
//	database := SignupData()
//	for _, user := range database {
//		postBody, _ := json.Marshal(user)
//		responseBody := bytes.NewBuffer(postBody)
//		//Leverage Go's HTTP Post function to make request
//		resp, err := http.Post("http://localhost:8080/api/auth/signup", "application/json", responseBody)
//		if err != nil {
//			log.Fatalf("An Error Occured %v", err)
//		}
//		defer resp.Body.Close()
//		//Read the response body
//		body, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		sb := string(body)
//		log.Printf(sb)
//	}
//}
//
//func TestLogin(t *testing.T) {
//	database := LoginData()
//	for _, user := range database {
//		postBody, _ := json.Marshal(user)
//		responseBody := bytes.NewBuffer(postBody)
//		//Leverage Go's HTTP Post function to make request
//		resp, err := http.Post("http://localhost:8080/api/auth/login", "application/json", responseBody)
//		if err != nil {
//			log.Fatalf("An Error Occured %v", err)
//		}
//		defer resp.Body.Close()
//		//Read the response body
//		body, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		sb := string(body)
//		log.Printf(sb)
//	}
//}
