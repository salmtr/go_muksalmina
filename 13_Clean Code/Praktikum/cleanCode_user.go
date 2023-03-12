package main

type user struct {
	id       int
	username int
	password int
}

// untuk penamaan struct alangkah baiknya ditulis menggunakan kombinasi huruf kapital
type userService struct {
	t []user
}

// untuk penamaan struct alangkah baiknya ditulis menggunakan kombinasi huruf kapital
func (u userService) getallusers() []user {
	return u.t
}

// untuk penamaan struct alangkah baiknya ditulis menggunakan kombinasi huruf kapital
func (u userService) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
