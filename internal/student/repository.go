package student

const CtxUserKey = "userId"

type StudentRepository interface {
	findById(id int)
}
