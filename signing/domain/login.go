package domain

func NewLogin(lid string) Login {
	return Login{Id: lid}
}

type Login struct {
	Id        string
	Frozen    bool
	FailedNum int
}

func (l *Login) Fail() bool {
	l.FailedNum++

	if l.FailedNum >= config.MaxNumOfFailedLogin {
		l.Frozen = true

		return true
	}

	return false
}

func (l *Login) RetryNum() int {
	if l.Frozen {
		return 0
	}

	return config.MaxNumOfFailedLogin - l.FailedNum
}

func (l *Login) HasFailure() bool {
	return l.FailedNum > 0
}
