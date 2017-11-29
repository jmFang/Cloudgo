package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

//mysql，数据库引擎为innodb事务才有效，myisam引擎是不支持事务的
// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {

	session := myengine.NewSession()
	defer session.Clone()

	//add Begin() before any action
	err := session.Begin()
	checkErr(err)

	_, err = session.Insert(u)

	if err != nil {
		session.Rollback()
		return err
	}
	// add commit after all actions
	err = session.Commit()
	checkErr(err)
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	res := make([]UserInfo, 0)
	err := myengine.Find(&res)
	checkErr(err)
	return res
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	user := new(UserInfo)
	//Gey method
	_, err := myengine.Id(id).Get(user)
	checkErr(err)
	return user
}
