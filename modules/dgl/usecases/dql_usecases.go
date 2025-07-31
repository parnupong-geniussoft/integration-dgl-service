package usecases

import (
	"fmt"
	"integration-dgl-service/modules/dgl/entities"
	"integration-dgl-service/modules/dgl/repositories"
	"integration-dgl-service/pkg/utils"
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

	grpID := ""

	dglRecord.Channel = u.DglRepo.GetSubmitChannel(dglRecord.Channel)
	dglRecord.MarketingCode, grpID = u.DglRepo.GetMarketingCode(dglRecord.MarketingCode)

	if grpID != "" {
		dglRecord.ApplicationType = u.DglRepo.GetApplicationType(grpID)
	}

	dglRecord.ProductType = u.DglRepo.GetProductType(dglRecord.ProductType)
	dglRecord.LoanPurpose = u.DglRepo.GetLoanPurposeDLG(dglRecord.LoanPurpose)
	dglRecord.LoanTerm = u.DglRepo.GetLoanTerm(dglRecord.LoanTerm)

	branch, err := u.DglRepo.GetBranch(dglRecord.Branch)
	if err == nil {
		dglRecord.Branch_District = branch.Brcity
		dglRecord.Branch_Sub_District = branch.Brad3
		if branch.Brstate != "" {
			dglRecord.Branch_Province = u.DglRepo.GetProvinceFallBack(branch.Brstate)
			dglRecord.Branch_Region = u.DglRepo.GetRegion(branch.Brstate)

			// create subDistrictId , districtId for Latitude-Longitude
			subD := u.DglRepo.GetSubDistrictCode(branch.Brstate + branch.Dist)
			di := utils.AddLeadingZeroIfSingleDigit(branch.Dist)
			di = branch.Brstate + di
			subDFormat := utils.FormatSubdistrictCode(subD)
			subDFormat = di + subDFormat
			latiLong, err := u.DglRepo.GetLatitudeLongitude(subDFormat + di + branch.Brstate)
			if err == nil {
				if latiLong.Latitude > 0 && latiLong.Longitude > 0 {
					dglRecord.Branch_Latitude = latiLong.Latitude
					dglRecord.Branch_Longitude = latiLong.Longitude
				}
			}
		}
	}

	var countryFromApplicant string

	for _, y := range dglRecord.Applicants {
		countryFromApplicant = y.Nationality

		y.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		y.Sex = u.DglRepo.GetGender(y.Sex)
		y.MaritalStatus = u.DglRepo.GetMaritalStatus(y.MaritalStatus)
		y.EducationLevel = u.DglRepo.GetEducation(y.EducationLevel)

		y.NcbConsentType = ""
		y.NcbConsentAcceptanceCode = ""
		if len(y.NCB_Consents) > 0 {
			for i, v := range y.NCB_Consents {
				if i == 0 {
					y.NcbConsentType = v.NcbConsentType
					y.NcbConsentAcceptanceCode = v.NcbConsentAcceptanceCode
				}
			}
		}
		y.NCB_Consents = []entities.NCBConsents{}
	}

	for _, y := range dglRecord.Addresses {
		y.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)

		addressCountry := countryFromApplicant
		addressProvinceState := y.Province
		addressSubDistrict, _ := strconv.Atoi(y.SubDistrict)
		parseAddressSubDistrict := fmt.Sprint(addressSubDistrict)
		addressDistrictCity, _ := strconv.Atoi(y.District)
		parseAddressDistrictCity := fmt.Sprint(addressDistrictCity)

		y.SubDistrict = u.DglRepo.GetSubDistrict(parseAddressSubDistrict + parseAddressDistrictCity + addressProvinceState + addressCountry)
		y.District = u.DglRepo.GetDistrict(parseAddressDistrictCity + addressProvinceState + addressCountry)

		y.Province = u.DglRepo.GetProvinceFallBack(addressCountry + addressProvinceState)
		if y.Province == addressCountry+addressProvinceState {
			y.Province = addressProvinceState
		}

		di := utils.AddLeadingZeroIfSingleDigit(parseAddressDistrictCity)
		di = addressProvinceState + di
		subDFormated := utils.FormatSubdistrictCode(parseAddressSubDistrict)
		subDFormated = di + subDFormated
		latiLong, err := u.DglRepo.GetLatitudeLongitude(subDFormated + di + addressProvinceState)
		if err == nil {
			if latiLong.Latitude > 0 && latiLong.Longitude > 0 {
				y.Latitude_N = latiLong.Latitude
				y.Longitude_N = latiLong.Longitude
				y.DistanceFromBranch = utils.CalculateDistance(y.Latitude_N, y.Longitude_N, dglRecord.Branch_Latitude, dglRecord.Branch_Longitude)
			}
		}
		y.Region = u.DglRepo.GetRegion(addressProvinceState)
	}

	for _, y := range dglRecord.Employments {
		y.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		y.Occupation = u.DglRepo.GetOccupation(y.Occupation)
	}

	for _, y := range dglRecord.ExistingApplicants {
		y.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
	}

	for _, y := range dglRecord.NCBs {
		y.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
	}

	for _, y := range dglRecord.NCBAccounts {
		y.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
	}

	for _, y := range dglRecord.NCBIDs {
		y.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
	}

	for _, y := range dglRecord.NCBNames {
		y.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
	}

	for _, y := range dglRecord.NCBAddresses {
		y.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
		subDistrictsTH := utils.RemoveSpecialStrings(y.SubDistrict)
		districtsTH := utils.RemoveSpecialStrings(y.District)
		provinceTH := utils.RemoveSpecialStrings(y.Province)
		regionName := u.DglRepo.GetRegion(provinceTH)
		latLongLookup, err := u.DglRepo.GetLatitudeLongitude(subDistrictsTH + districtsTH + provinceTH)
		if err == nil {
			if latLongLookup.Latitude > 0 && latLongLookup.Longitude > 0 {
				y.Latitude_N = latLongLookup.Latitude
				y.Longitude_N = latLongLookup.Longitude
				y.DistanceFromBranch = utils.CalculateDistance(y.Latitude_N, y.Longitude_N, dglRecord.Branch_Latitude, dglRecord.Branch_Longitude)
			}
		}
		y.Region = regionName
	}

	for _, y := range dglRecord.NCBScores {
		y.ApplicantType = u.DglRepo.GetApplicantType(y.ApplicantType)
	}

	return result, nil
}
