package domain

// User User领域对象（DDD中的聚合根）
// 其他叫法：BO（business object）
type User struct {
	Email    string
	Password string
}
