package utility

import (
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"github.com/google/uuid"
	"strconv"
)

func StringToUUID(str string) (uuid.UUID, errorUtils.CustomError) {
	id, err := uuid.Parse(str)

	if err != nil {
		return uuid.UUID{}, errorUtils.ErrInternalServer
	}

	return id, nil
}

func StringToInt64(str string) (int64, errorUtils.CustomError) {
	id, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return 0, errorUtils.ErrInternalServer
	}

	return id, nil
}
