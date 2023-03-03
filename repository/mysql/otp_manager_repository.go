package mysql

import "github.com/tensuqiuwulu/pandora-service/config"

type OtpManagerRepositoryInterface interface {
	UpdatePhoneLimit()
}

type OtpManagerRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewOtpManagerRepository(
	configDatabase *config.Database,
) OtpManagerRepositoryInterface {
	return &OtpManagerRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *OtpManagerRepositoryImplementation) UpdatePhoneLimit() {

}
