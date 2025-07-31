package usecases

import (
	"fmt"
	"integration-dgl-service/modules/dgl/entities"
	"integration-dgl-service/modules/dgl/repositories"
	"strconv"
)

type DglUsecase interface {
	DglMapper(dglRecord *entities.DGLRecord) (*entities.DGLRecord, error)
}

type dglUsecase struct {
	DglRepo repositories.DglRepository
}

func NewDglUsecase(dglRepo repositories.DglRepository) DglUsecase {
	return &dglUsecase{
		DglRepo: dglRepo,
	}
}

func (u *dglUsecase) DglMapper(dglRecord *entities.DGLRecord) (*entities.DGLRecord, error) {
	result := &entities.DGLRecord{}

	var grpID string

	result.Channel = u.DglRepo.GetSubmitChannel(dglRecord.Channel)
	result.MarketingCode, grpID = u.DglRepo.GetMarketingCode(dglRecord.MarketingCode)

	if grpID != "" {
		result.ApplicationType = u.DglRepo.GetApplicationType(grpID)
	}

	result.ProductType = u.DglRepo.GetProductType(dglRecord.ProductType)
	result.LoanPurpose = u.DglRepo.GetLoanPurposeDLG(dglRecord.LoanPurpose)
	result.LoanTerm = u.DglRepo.GetLoanTerm(dglRecord.LoanTerm)

	var countryFromApplicant string
	result.Applicants = make([]*entities.Applicant, 0, len(dglRecord.Applicants))
	for _, y := range dglRecord.Applicants {
		countryFromApplicant = y.Nationality

		newApplicant := y
		newApplicant.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		newApplicant.Sex = u.DglRepo.GetGender(y.Sex)
		newApplicant.MaritalStatus = u.DglRepo.GetMaritalStatus(y.MaritalStatus)
		newApplicant.EducationLevel = u.DglRepo.GetEducation(y.EducationLevel)

		newApplicant.NcbConsentType = ""
		newApplicant.NcbConsentAcceptanceCode = ""
		if len(y.NCB_Consents) > 0 {
			newApplicant.NcbConsentType = y.NCB_Consents[0].NcbConsentType
			newApplicant.NcbConsentAcceptanceCode = y.NCB_Consents[0].NcbConsentAcceptanceCode
		}
		newApplicant.NCB_Consents = []entities.NCBConsents{}

		result.Applicants = append(result.Applicants, newApplicant)
	}

	result.Addresses = make([]*entities.Address, 0, len(dglRecord.Addresses))
	for _, y := range dglRecord.Addresses {
		newAddress := y
		newAddress.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)

		addressCountry := countryFromApplicant
		addressProvinceState := y.Province
		addressSubDistrict, _ := strconv.Atoi(y.SubDistrict)
		parseAddressSubDistrict := fmt.Sprint(addressSubDistrict)
		addressDistrictCity, _ := strconv.Atoi(y.District)
		parseAddressDistrictCity := fmt.Sprint(addressDistrictCity)

		newAddress.SubDistrict = u.DglRepo.GetSubDistrict(parseAddressSubDistrict + parseAddressDistrictCity + addressProvinceState + addressCountry)
		newAddress.District = u.DglRepo.GetDistrict(parseAddressDistrictCity + addressProvinceState + addressCountry)

		newAddress.Province = u.DglRepo.GetProvinceFallBack(addressCountry + addressProvinceState)
		if newAddress.Province == addressCountry+addressProvinceState {
			newAddress.Province = addressProvinceState
		}

		result.Addresses = append(result.Addresses, newAddress)
	}

	result.Employments = make([]*entities.Employment, 0, len(dglRecord.Employments))
	for _, y := range dglRecord.Employments {
		newEmp := y
		newEmp.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		newEmp.Occupation = u.DglRepo.GetOccupation(y.Occupation)
		result.Employments = append(result.Employments, newEmp)
	}

	result.ExistingApplicants = make([]*entities.ExistingApplicant, 0, len(dglRecord.ExistingApplicants))
	for _, y := range dglRecord.ExistingApplicants {
		newItem := y
		newItem.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		result.ExistingApplicants = append(result.ExistingApplicants, newItem)
	}

	result.NCBs = make([]*entities.NCB, 0, len(dglRecord.NCBs))
	for _, y := range dglRecord.NCBs {
		newItem := y
		newItem.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		result.NCBs = append(result.NCBs, newItem)
	}

	result.NCBAccounts = make([]*entities.NCBAccount, 0, len(dglRecord.NCBAccounts))
	for _, y := range dglRecord.NCBAccounts {
		newItem := y
		newItem.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		result.NCBAccounts = append(result.NCBAccounts, newItem)
	}

	result.NCBIDs = make([]*entities.NCBID, 0, len(dglRecord.NCBIDs))
	for _, y := range dglRecord.NCBIDs {
		newItem := y
		newItem.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		result.NCBIDs = append(result.NCBIDs, newItem)
	}

	result.NCBNames = make([]*entities.NCBName, 0, len(dglRecord.NCBNames))
	for _, y := range dglRecord.NCBNames {
		newItem := y
		newItem.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		result.NCBNames = append(result.NCBNames, newItem)
	}

	result.NCBAddresses = make([]*entities.NCBAddress, 0, len(dglRecord.NCBAddresses))
	for _, y := range dglRecord.NCBAddresses {
		newItem := y
		newItem.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		result.NCBAddresses = append(result.NCBAddresses, newItem)
	}

	result.NCBScores = make([]*entities.NCBScore, 0, len(dglRecord.NCBScores))
	for _, y := range dglRecord.NCBScores {
		newItem := y
		newItem.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		result.NCBScores = append(result.NCBScores, newItem)
	}

	return result, nil
}
