package functions

func CompareData(userPassword, loginPassword string) bool {
	hashLoginPass := HashData(loginPassword)
	if userPassword == hashLoginPass {
		return true
	} else {
		return false
	}
}
