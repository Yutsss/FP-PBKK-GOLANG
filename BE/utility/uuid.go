package utility

import (
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"github.com/google/uuid"
)

func StringToUUID(str string) (uuid.UUID, errorUtils.CustomError) {
	id, err := uuid.Parse(str)

	if err != nil {
		return uuid.UUID{}, errorUtils.ErrInternalServer
	}

	return id, nil
}
