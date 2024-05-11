package hi

import "seregaa020292/StartMonorepoMicroservices/fines/common/pkg/hello"

func Message() []byte {
	return []byte(hello.Message("monitoring"))
}
