package repositories

import (
	"fmt"
	"integration-dgl-service/modules/dgl/entities"
	"integration-dgl-service/pkg/utils"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
)

type DglRepository interface {
	GetAndInitCache(k string, dataKey string, getFromDB func() map[string]string) string
	GetSubmitChannel(k string) string
	GetMarketingCode(k string) (name string, id string)
	GetApplicationType(k string) string
	GetProductType(k string) string
	GetLoanPurposeDlg(k string) string
	GetLoanTerm(k string) string
	GetGender(k string) string
	GetMaritalStatus(k string) string
	GetEducation(k string) string
	GetSubDistrict(k string) string
	GetDistrict(k string) string
	GetProvinceFallBack(k string) string
	GetProvince(k string) string
	GetProvinceTH(k string) string
	GetOccupation(k string) string
	GetApplicantType(k string) string
	GetLoanPurposeDLG(k string) string
}

type dglRepo struct {
	Db *sqlx.DB
	C  *cache.Cache
}

func NewDglRepository(db *sqlx.DB, c *cache.Cache) DglRepository {
	return &dglRepo{
		Db: db,
		C:  c,
	}
}

func (r *dglRepo) GetAndInitCache(k string, dataKey string, getFromDB func() map[string]string) string {
	m, found := r.C.Get(dataKey)

	times := utils.GetTimeMinsToNewDay()
	if !found {
		m = getFromDB()
		r.C.Set(dataKey, m, times)
	}

	x, found := m.(map[string]string)[k]

	if !found {
		return k
	}

	return x
}

func (r *dglRepo) GetSubmitChannel(k string) string {
	x := r.GetAndInitCache(k, "submitChannel", func() map[string]string {
		xs := []entities.SubmitChannel{}
		err := r.Db.Select(&xs, "SELECT code, description FROM submit_channel")

		if err != nil {
			fmt.Println("GetSubmitChannel Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key] = x.Value
		}

		return m
	})

	return x
}

func (r *dglRepo) GetMarketingCode(k string) (name string, id string) {
	x := r.GetAndInitCache(k, "marketingCode", func() map[string]string {
		xs := []entities.MarketingCode{}
		if r.Db != nil {
			err := r.Db.Select(&xs, "SELECT loan_tp_id, loan_tp_nm, loan_grp_id FROM market_code")

			if err != nil {
				fmt.Println("GetMarketingCode Err: ", err)
			}
		}
		m := make(map[string]string)

		for _, x := range xs {
			mapKey := ""
			if x.Key != nil {
				mapKey = *x.Key
			}

			value2 := ""
			if x.Value2 != nil {
				value2 = *x.Value2
			}
			m[mapKey] = fmt.Sprintf("%s|%s", x.Value, value2)
		}

		return m
	})

	arrStr := strings.Split(x, "|")

	if len(arrStr) > 1 {
		return arrStr[0], arrStr[1]
	}

	return k, ""
}

func (r *dglRepo) GetApplicationType(k string) string {
	x := r.GetAndInitCache(k, "applicationType", func() map[string]string {
		xs := []entities.ApplicationType{}
		err := r.Db.Select(&xs, "SELECT id, grp FROM application_type")

		if err != nil {
			fmt.Println("GetApplicationType Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key] = x.Value
		}
		return m
	})

	return x
}

func (r *dglRepo) GetProductType(k string) string {
	x := r.GetAndInitCache(k, "productType", func() map[string]string {
		xs := []entities.ProductType{}
		if r.Db != nil {
			err := r.Db.Select(&xs, "SELECT pd_tp, pd_tp_nm FROM product_type")
			if err != nil {
				fmt.Println("GetProductType Err: ", err)
			}
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key] = x.Value
		}

		return m
	})

	return x
}

func (r *dglRepo) GetLoanPurposeDlg(k string) string {
	x := r.GetAndInitCache(k, "loanPurpose", func() map[string]string {
		xs := []entities.LoanPurpose{}
		if r.Db != nil {
			err := r.Db.Select(&xs, "SELECT code, description FROM loan_app_info_loan_purpose")

			if err != nil {
				fmt.Println("GetLoanPurpose Err: ", err)
			}
		}
		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key] = x.Value
		}

		return m
	})

	return x
}

func (r *dglRepo) GetLoanTerm(k string) string {
	x := r.GetAndInitCache(k, "loanTerm", func() map[string]string {
		xs := []entities.LoanTerm{}
		err := r.Db.Select(&xs, "SELECT code, description FROM loan_app_info_term_type")

		if err != nil {
			fmt.Println("GetLoanTerm Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key] = x.Value
		}

		return m
	})

	return x
}

func (r *dglRepo) GetGender(k string) string {
	x := r.GetAndInitCache(k, "gender", func() map[string]string {
		xs := []entities.Gender{}
		if r.Db != nil {
			err := r.Db.Select(&xs, "SELECT code, description FROM gender")

			if err != nil {
				fmt.Println("GetGender Err:", err)
			}
		}
		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key] = x.Value
		}
		return m
	})

	arrStr := strings.Split(x, ",")

	if len(arrStr) > 1 {
		return arrStr[0]
	}

	return x
}

func (r *dglRepo) GetMaritalStatus(k string) string {
	x := r.GetAndInitCache(k, "maritalStatus", func() map[string]string {
		xs := []entities.MaritalStatus{}
		err := r.Db.Select(&xs, "SELECT code, description FROM marital_status")

		if err != nil {
			fmt.Println("GetMaritalStatus Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key] = x.Value
		}
		return m
	})

	return x
}

func (r *dglRepo) GetEducation(k string) string {
	x := r.GetAndInitCache(k, "education", func() map[string]string {
		xs := []entities.Education{}
		err := r.Db.Select(&xs, "SELECT code, description FROM education")

		if err != nil {
			fmt.Println("GetEducation Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key] = x.Value
		}

		return m
	})

	return x
}

func (r *dglRepo) GetSubDistrict(k string) string {
	x := r.GetAndInitCache(k, "subDistrict", func() map[string]string {
		xs := []entities.SubDistrict{}
		err := r.Db.Select(&xs, "SELECT subd, di, st, cn, thai_description FROM sub_district")

		if err != nil {
			fmt.Println("GetSubDistrict Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key+x.Key2+x.Key3+x.Key4] = x.Value
		}

		return m
	})

	return x
}

func (r *dglRepo) GetDistrict(k string) string {
	x := r.GetAndInitCache(k, "district", func() map[string]string {
		xs := []entities.District{}
		err := r.Db.Select(&xs, "SELECT di, st ,cn, thai_description FROM district")

		if err != nil {
			fmt.Println("GetDistrict Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key+x.Key2+x.Key3] = x.Value
		}

		return m
	})

	return x
}

func (r *dglRepo) GetProvinceFallBack(k string) string {
	if k == "" {
		return ""
	}

	x := r.GetProvince(k)
	if x == k {
		x = r.GetProvinceTH(k)

		if x == "TH"+k {
			return k
		}
	}

	return x
}

func (r *dglRepo) GetProvince(k string) string {
	x := r.GetAndInitCache(k, "province", func() map[string]string {
		xs := []entities.Province{}
		err := r.Db.Select(&xs, "SELECT country_code, description, state_province FROM province")

		if err != nil {
			fmt.Println("GetProvince Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key+x.Key2] = x.Value
		}

		return m
	})

	return x
}

func (r *dglRepo) GetProvinceTH(k string) string {
	country := "TH"
	x := r.GetAndInitCache(country+k, "province", func() map[string]string {
		xs := []entities.Province{}
		err := r.Db.Select(&xs, "SELECT country_code, description, state_province FROM province")

		if err != nil {
			fmt.Println("GetProvince Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key+x.Key2] = x.Value
		}

		return m
	})

	return x
}

func (r *dglRepo) GetOccupation(k string) string {
	x := r.GetAndInitCache(k, "occupation", func() map[string]string {
		xs := []entities.Occupation{}
		err := r.Db.Select(&xs, "SELECT code, description FROM occupation")

		if err != nil {
			fmt.Println("GetOccupation Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key] = x.Value
		}

		return m
	})

	return x
}

func (r *dglRepo) GetApplicantType(k string) string {
	x := r.GetAndInitCache(k, "applicantType", func() map[string]string {
		xs := []entities.ApplicantType{}
		if r.Db != nil {
			err := r.Db.Select(&xs, "SELECT code, description FROM borrower_info_role_type_code")

			if err != nil {
				fmt.Println("GetApplicantType Err: ", err)
			}
		}
		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key] = x.Value
		}

		return m
	})

	return x
}

func (r *dglRepo) GetLoanPurposeDLG(k string) string {
	x := r.GetAndInitCache(k, "loanPurpose", func() map[string]string {
		xs := []entities.LoanPurpose{}
		if r.Db != nil {
			err := r.Db.Select(&xs, "SELECT code, description FROM loan_app_info_loan_purpose")

			if err != nil {
				fmt.Println("GetLoanPurpose Err: ", err)
			}
		}
		m := make(map[string]string)

		for _, x := range xs {
			m[x.Key] = x.Value
		}

		return m
	})

	return x
}
