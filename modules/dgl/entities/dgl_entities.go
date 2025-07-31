package entities

type Loan struct {
	LoanTypeCode      string `json:"Loan_Type_Code"`
	LoanAmount        string `json:"Loan_Amount"`
	LoanPurpose       string `json:"Loan_Purpose"`
	ProductCode       string `json:"Product_Code"`
	MarketingCode     string `json:"Marketing_Code"`
	InstallmentPeriod string `json:"Installment_Period"`
	InstallmentAmount string `json:"Installment_Amount"`
}

type NCBConsents struct {
	NcbConsentType           string `json:"Ncb_Consent_Type"`
	NcbConsentAcceptanceCode string `json:"Ncb_Consent_Acceptance_Code"`
}

type Applicant struct {
	ApplicantType            string        `json:"Applicant_Type"`
	ThaiName                 string        `json:"Thai_Name"`
	ThaiMiddleName           string        `json:"Thai_Middle_Name"`
	ThaiSurname              string        `json:"Thai_Surname"`
	CIFNumber                string        `json:"CIF_Number"`
	Sex                      string        `json:"Sex"`
	DateOfBirth              string        `json:"Date_Of_Birth"`
	Age                      int           `json:"Age"`
	Nationality              string        `json:"Nationality"`
	MaritalStatus            string        `json:"Marital_Status"`
	EducationLevel           string        `json:"Education_Level"`
	CitizenIdNumber          string        `json:"Citizen_Id_Number"`
	MobilePhone              string        `json:"Mobile_Phone"`
	Email                    string        `json:"Email"`
	ResidentType             string        `json:"Resident_Type"`
	LivingPlacePeriodMonths  int           `json:"Living_Place_Period_Months"`
	NCB_Consents             []NCBConsents `json:"NCB_Consents"`
	NcbConsentType           string        `json:"Ncb_Consent_Type"`
	NcbConsentAcceptanceCode string        `json:"Ncb_Consent_Acceptance_Code"`
}

type Address struct {
	ApplicantType      string  `json:"Applicant_Type"`
	ThaiName           string  `json:"Thai_Name"`
	ThaiSurname        string  `json:"Thai_Surname"`
	CitizenIdNumber    string  `json:"Citizen_Id_Number"`
	YearsOfResidence   int     `json:"Years_Of_Residence"`
	AddressType        string  `json:"Address_Type"`
	Address1           string  `json:"Address_1"`
	Address2           string  `json:"Address_2"`
	Address3           string  `json:"Address_3"`
	Address4           string  `json:"Address_4"`
	Province           string  `json:"Province"`
	District           string  `json:"District"`
	SubDistrict        string  `json:"Sub_District"`
	ZipCode            string  `json:"Zip_Code"`
	Longitude_N        float64 `json:"Longitude_N,omitempty"`
	Latitude_N         float64 `json:"Latitude_N,omitempty"`
	DistanceFromBranch float64 `json:"Distance_From_Branch,omitempty"`
	Resident_Type      string  `json:"Resident_Type"`
	Region             string  `json:"Region"`
}

type Employment struct {
	ApplicantType   string  `json:"Applicant_Type"`
	ThaiName        string  `json:"Thai_Name"`
	ThaiSurname     string  `json:"Thai_Surname"`
	CitizenIdNumber string  `json:"Citizen_Id_Number"`
	CompanyName     string  `json:"Company_Name"`
	OccupationType  string  `json:"Occupation_Type"`
	Occupation      string  `json:"Occupation"`
	MonthlyIncome   float64 `json:"Monthly_Income"`
}

type ExistingApplicant struct {
	ApplicantType        string `json:"Applicant_Type"`
	ThaiName             string `json:"Thai_Name"`
	ThaiSurname          string `json:"Thai_Surname"`
	CitizenIdNumber      string `json:"Citizen_Id_Number"`
	CIFNumber            string `json:"CIF_Number"`
	TitleTH              string `json:"Title_TH"`
	FirstNameTH          string `json:"First_Name_TH"`
	LastNameTH           string `json:"Last_Name_TH"`
	MiddleNameTH         string `json:"Middle_Name_TH"`
	TitleEN              string `json:"Title_EN"`
	FirstNameEN          string `json:"First_Name_EN"`
	LastNameEN           string `json:"Last_Name_EN"`
	MiddleNameEN         string `json:"Middle_Name_EN"`
	DateOfBirth          string `json:"Date_of_Birth"`
	Nationality          string `json:"Nationality"`
	EducationLevel       string `json:"Education_Level"`
	OccupationSector     string `json:"Occupation_Sector"`
	OccupationGroup      string `json:"Occupation_Group"`
	SubOccupation        string `json:"Sub_Occupation"`
	Employee             string `json:"Employee"`
	HomePhone            string `json:"Home_Phone"`
	MobilePhone          string `json:"Mobile_Phone"`
	OfficePhone          string `json:"Office_Phone"`
	Extension            string `json:"Extension"`
	FaxNumber            string `json:"Fax_Number"`
	Date_of_Death        string `json:"Date_of_Death"`
	Email                string `json:"Email"`
	KYCCDDStatus         string `json:"KYC_CDD_Status"`
	Sex                  string `json:"Sex"`
	CountryOfBirth       string `json:"Country_Of_Birth"`
	Nationality2         string `json:"Nationality_2"`
	Employer             string `json:"Employer"`
	IncomeCode           string `json:"Income_Code"`
	LegalNo              string `json:"Legal_No"`
	LegalSoi             string `json:"Legal_Soi"`
	LegalRoad            string `json:"Legal_Road"`
	LegalSubDistrict     string `json:"Legal_Sub-district"`
	LegalCountry         string `json:"Legal_Country"`
	LegalDistrictCity    string `json:"Legal_District_City"`
	LegalPostalCode      string `json:"Legal_Postal_Code"`
	LegalProvinceState   string `json:"Legal_Province_State"`
	OfficeNo             string `json:"Office_No"`
	OfficeSoi            string `json:"Office_Soi"`
	OfficeRoad           string `json:"Office_Road"`
	OfficeSubDistrict    string `json:"Office_Sub-district"`
	OfficeCountry        string `json:"Office_Country"`
	OfficeDistrictCity   string `json:"Office_District_City"`
	OfficePostalCode     string `json:"Office_Postal_Code"`
	OfficeProvinceState  string `json:"Office_Province_State"`
	CurrentNo            string `json:"Current_No"`
	CurrentSoi           string `json:"Current_Soi"`
	CurrentRoad          string `json:"Current_Road"`
	CurrentSubDistrict   string `json:"Current_Sub_District"`
	CurrentCountry       string `json:"Current_Country"`
	CurrentDistrictCity  string `json:"Current_District_City"`
	CurrentPostalCode    string `json:"Current_Postal_Code"`
	CurrentProvinceState string `json:"Current_Province_State"`
}

type NCB struct {
	ApplicantType                 string `json:"Applicant_Type"`
	ThaiName                      string `json:"Thai_Name"`
	ThaiSurname                   string `json:"Thai_Surname"`
	CitizenIDNumber               string `json:"Citizen_Id_Number"`
	SubjectReturnCode             string `json:"Subject_Return_Code"`
	NoOfApprovedCreditL30D        int    `json:"No_Of_Approved_Credit_L30D"`
	CreditLimitApprovedAllAccount int    `json:"Credit_Limit_Approved_All_Account"`
	// MaxNCBApprovedCreditLimit     int    `json:"Max_NCB_Approved_Credit_Limit"`
	MaxNCBApprovedCreditLimit int `json:"-"`
}

type NCBAccount struct {
	ApplicantType               string `json:"Applicant_Type"`
	ThaiName                    string `json:"Thai_Name"`
	ThaiSurname                 string `json:"Thai_Surname"`
	CitizenIDNumber             string `json:"Citizen_Id_Number"`
	SubjectNo                   int    `json:"Subject_No"`
	MemberShortName             string `json:"Member_Short_Name"`
	AccountType                 string `json:"Account_Type"`
	OwnershipIndicator          string `json:"Ownership_Indicator"`
	CurrencyCode                string `json:"Currency_Code"`
	DateAccountOpened           string `json:"Date_Account_Opened"`
	DateOfLastPayment           string `json:"Date_Of_Last_Payment"`
	AsOfDate                    string `json:"As_Of_Date"`
	CreditLimit                 int    `json:"Credit_Limit"`
	AmountOwed                  int    `json:"Amount_Owed"`
	AmountPastDue               int    `json:"Amount_Past_Due"`
	InstallmentFrequency        string `json:"Installment_Frequency"`
	InstallmentAmount           int    `json:"Installment_Amount"`
	InstallmentNumberOfPayments string `json:"Installment_Number_Of_Payments"`
	AccountStatus               string `json:"Account_Status"`
	LoanClass                   string `json:"Loan_Class"`
	PaymentHistory1             string `json:"Payment_History_1"`
	PaymentHistory2             string `json:"Payment_History_2"`
	PaymentHistoryStartDate     string `json:"Payment_History_Start_Date"`
	PaymentHistoryEndDate       string `json:"Payment_History_End_Date"`
	LoanObjective               string `json:"Loan_Objective,omitempty"`
	Collateral1                 string `json:"Collateral_1"`
	Collateral2                 string `json:"Collateral_2"`
	Collateral3                 string `json:"Collateral_3"`
	PercentPayment              string `json:"Percent_Payment,omitempty"`
	TypeOfCreditCard            string `json:"Type_Of_Credit_Card"`
	DefaultDate                 string `json:"Default_Date,omitempty"`
	DateAccountClosed           string `json:"Date_Account_Closed"`
	UnitMake                    string `json:"Unit_Make,omitempty"`
	UnitModel                   string `json:"Unit_Model,omitempty"`
}

type NCBID struct {
	ThaiName        string `json:"Thai_Name"`
	ThaiSurname     string `json:"Thai_Surname"`
	CitizenIDNumber string `json:"Citizen_Id_Number"`
	ApplicantType   string `json:"Applicant_Type"`
	SubjectNo       int    `json:"Subject_No"`
	IdType          string `json:"Id_Type"`
	IdNumber        string `json:"Id_Number"`
}

type NCBName struct {
	ThaiName         string `json:"Thai_Name"`
	ThaiSurname      string `json:"Thai_Surname"`
	CitizenIDNumber  string `json:"Citizen_Id_Number"`
	ApplicantType    string `json:"Applicant_Type"`
	SubjectNo        int    `json:"Subject_No"`
	FamilyName1      string `json:"Family_Name_1"`
	FirstName        string `json:"First_Name"`
	MaritalStatus    string `json:"Marital_Status"`
	DateOfBirth      string `json:"Date_Of_Birth"`
	Gender           string `json:"Gender"`
	Title            string `json:"Title"`
	Nationality      string `json:"Nationality"`
	NumberOfChildren int    `json:"Number_Of_Children"`
	Occupation       string `json:"Occupation"`
	ConsentToEnquire string `json:"Consent_To_Enquire"`
}

type NCBAddress struct {
	ThaiName                string  `json:"Thai_Name"`
	ThaiSurname             string  `json:"Thai_Surname"`
	CitizenIDNumber         string  `json:"Citizen_Id_Number"`
	ApplicantType           string  `json:"Applicant_Type"`
	SubjectNo               int     `json:"Subject_No"`
	AddressLine1            string  `json:"Address_Line_1"`
	AddressLine2            string  `json:"Address_Line_2"`
	AddressLine3            string  `json:"Address_Line_3"`
	SubDistrict             string  `json:"Sub_District"`
	District                string  `json:"District"`
	Province                string  `json:"Province"`
	Country                 string  `json:"Country"`
	PostalCode              string  `json:"Postal_Code"`
	ReportedDate            string  `json:"Reported_Date"`
	Telephone               string  `json:"Telephone"`
	TelephoneType           string  `json:"Telephone_Type"`
	AddressType             string  `json:"Address_Type"`
	ResidentialStatus       string  `json:"Residential_Status"`
	Longitude_N             float64 `json:"Longitude_N,omitempty"`
	Latitude_N              float64 `json:"Latitude_N,omitempty"`
	DistanceFromBranch      float64 `json:"Distance_From_Branch,omitempty"`
	DistanceFromCollateral1 float64 `json:"Distance_From_Collateral_1,omitempty"`
	Region                  string  `json:"Region"`
}

type NCBScore struct {
	ThaiName             string `json:"Thai_Name"`
	ThaiSurname          string `json:"Thai_Surname"`
	CitizenIDNumber      string `json:"Citizen_Id_Number"`
	ApplicantType        string `json:"Applicant_Type"`
	SubjectNo            string `json:"Subject_No"`
	ScoreVersion         string `json:"Score_Version"`
	ScoreSegment         string `json:"Score_Segment"`
	RecalibrationVersion string `json:"Recalibration_Version"`
	ScoreVersionDate     string `json:"Score_Version_Date"`
	Score                int    `json:"Score"`
	ScoreGrade           string `json:"Score_Grade"`
	Odds                 int    `json:"Odds"`
	ReasonCode1          string `json:"Reason_Code_1"`
	ReasonCode2          string `json:"Reason_Code_2"`
	ReasonCode3          string `json:"Reason_Code_3"`
	ReasonCode4          string `json:"Reason_Code_4"`
	ReasonCode5          string `json:"Reason_Code_5"`
	ErrorCode            string `json:"Error_Code"`
}

type DeviceInformation struct {
	Deviceinfo  string `json:"Device_info,omitempty"`
	DeviceID    string `json:"Device_ID,omitempty"`
	DeviceIMEI  string `json:"Device_IMEI,omitempty"`
	DeviceOS    string `json:"Device_OS,omitempty"`
	IPAddress   string `json:"IP_Address,omitempty"`
	Geolocation string `json:"Geolocation,omitempty"`
}

type DGLRecord struct {
	ApplicationNumber string  `json:"Application_Number"`
	ApplicationDate   string  `json:"Application_Date"`
	ApplicationType   string  `json:"Application_Type"`
	LoanAmount        float64 `json:"Loan_Amount"`
	Channel           string  `json:"Channel"`
	ProductType       string  `json:"Product_Type"`
	ApplicationStatus string  `json:"Application_Status"`
	MarketingCode     string  `json:"Marketing_Code"`
	ProductMarketCode string  `json:"productMarketCode"`
	LoanPurpose       string  `json:"Loan_Purpose"`
	LoanTermNumber    int     `json:"Loan_Term_Number"`
	LoanTerm          string  `json:"Loan_Term"`
	Branch            string  `json:"Branch"`

	Branch_Sub_District string  `json:"Branch_Sub_District,"`
	Branch_District     string  `json:"Branch_District,"`
	Branch_Province     string  `json:"Branch_Province,"`
	Branch_Region       string  `json:"Branch_Region,"`
	Branch_Latitude     float64 `json:"Branch_Latitude,omitempty"`
	Branch_Longitude    float64 `json:"Branch_Longitude,omitempty"`

	Applicants         []*Applicant         `json:"Applicant"`
	Loans              []*Loan              `json:"Loans"`
	Addresses          []*Address           `json:"Address"`
	Employments        []*Employment        `json:"Employment"`
	ExistingApplicants []*ExistingApplicant `json:"Existing_Applicant"`
	NCBs               []*NCB               `json:"NCB"`
	NCBAccounts        []*NCBAccount        `json:"NCB_Account"`
	NCBIDs             []*NCBID             `json:"NCB_ID"`
	NCBNames           []*NCBName           `json:"NCB_Name"`
	NCBAddresses       []*NCBAddress        `json:"NCB_Address"`
	NCBScores          []*NCBScore          `json:"NCB_Score"`
	DeviceInformations []*DeviceInformation `json:"Device_Information"`
}
