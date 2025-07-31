package entities

type ApplicationType struct {
	Key   string `db:"id"`
	Value string `db:"grp"`
}

type ApplicationStatus struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type ApplicantType struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type TypeOfCollateral struct {
	Key   string `db:"code"`
	Value string `db:"dsc"`
}

type SubTypeOfCollateral struct {
	Key   string `db:"clt_tp_code"`
	Key2  string `db:"clt_sbtp_code"`
	Value string `db:"dsc"`
}

type SubmitChannel struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type ProductType struct {
	Key   string `db:"pd_tp"`
	Value string `db:"pd_tp_nm"`
}

type MarketingCode struct {
	Key    *string `db:"loan_tp_id"`
	Value  string  `db:"loan_tp_nm"`
	Value2 *string `db:"loan_grp_id"`
}

type LoanPurpose struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type LoanTerm struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type Gender struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type MaritalStatus struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type Education struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type Province struct {
	Key   string `db:"country_code"`
	Key2  string `db:"state_province"`
	Value string `db:"description"`
}

type District struct {
	Key   string `db:"di"`
	Key2  string `db:"st"`
	Key3  string `db:"cn"`
	Value string `db:"thai_description"`
}

type SubDistrict struct {
	Key   string `db:"subd"`
	Key2  string `db:"di"`
	Key3  string `db:"st"`
	Key4  string `db:"cn"`
	Value string `db:"thai_description"`
}

type CovertAddress struct {
	Key    string `db:"o_province"`
	Key2   string `db:"o_district"`
	Key3   string `db:"o_sub_district"`
	Value  string `db:"n_province"`
	Value2 string `db:"n_dis"`
	Value3 string `db:"n_subdis"`
}

type SubDistrictDescription struct {
	Key   string `db:"st"`
	Key2  string `db:"di"`
	Key3  string `db:"thai_description"`
	Key4  string `db:"cn"`
	Value string `db:"subd"`
}

type OccupationSector struct {
	Key   string `db:"code"`
	Value string `db:"thai_description"`
}

type OccupationGroup struct {
	Key   string `db:"occ"`
	Value string `db:"thai_description"`
}

type Occupation struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type SubOccupation struct {
	Key   string `db:"occ"`
	Key2  string `db:"so"`
	Value string `db:"thai_description"`
}

type Country struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type Position struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type ProofOfIncome struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type incomeCode struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type SubjectReturnCode struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type AccountType struct {
	Key   string `db:"code"`
	Value string `db:"account_type"`
}

type BusinessType struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}
type BusinessType2 struct {
	Key   string `db:"isicsd"`
	Key2  string `db:"isicgc"`
	Value string `db:"description"`
}

type BusinessType3 struct {
	Key   string `db:"isicsd"`
	Key2  string `db:"isicgc"`
	Key3  string `db:"isicts"`
	Value string `db:"t_description"`
}

type OwnershipIndicator struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type LoanClass struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type LoanObjective struct {
	Key   string  `db:"code"`
	Value *string `db:"description"`
}

type Collateral struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type CreditCardType struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type IDType struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type ConsentToEnquire struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type TelephoneType struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type TypeOfBusiness struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type AddressType struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type ResidentialStatus struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type RelationshipType struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type PurchaseType struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type BuildingType struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type ReasonCode struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type KYCCCDStatus struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type KTBCustomerCode struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

type InvolvePartyGroup struct {
	Key   string `db:"involved_party_code"`
	Value string `db:"description"`
}

type SystemSourceUsernamePassword struct {
	Key    string `db:"name"`
	Value  string `db:"username"`
	Value2 string `db:"password"`
}

type SystemSourceAuthURL struct {
	Key   string `db:"name"`
	Value string `db:"path_auth"`
}

type SystemSourceActionURL struct {
	Key   string `db:"name"`
	Value string `db:"path_action"`
}

type SystemSourceURL struct {
	Key   string `db:"name"`
	Value string `db:"url"`
}
type ErrorCode struct {
	Key   string `db:"code"`
	Value string `db:"description"`
}

// key = code , value = description
type LatitudeLongitude struct {
	Key    string  `db:"sub-districts_id"`
	Key2   string  `db:"districts_id"`
	Key3   string  `db:"provnce_id"`
	Key4   string  `db:"sub_districts_th"`
	Key5   string  `db:"districts_th"`
	Key6   string  `db:"provnce_th"`
	Value  float64 `db:"latitude"`
	Value2 float64 `db:"longitude"`
}

type Branch struct {
	Key    string `db:"ccdef"`
	Key2   string `db:"brcntry"`
	Key3   string `db:"brstate"`
	Key4   string `db:"dist"`
	Value  string `db:"brcity"` //(อำเภอ name)
	Value2 string `db:"brad3"`  //(ตำบล name)
}

type Region struct {
	Key   string `db:"provincemoi_id"`
	Key2  string `db:"provincenamethai"`
	Value string `db:"regionname"`
}
