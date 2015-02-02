package models

type MResp struct {
	responseNo  int
	responseMsg string
}

type MFindPwdResp struct {
	responseNo  int
	responseMsg string
	password string
}

type MUserExistsResp struct {
	responseNo  int
	responseMsg string
	exists string
}

type MUserLoginResp struct {
	responseNo  int
	responseMsg string
	token string
	tokenExpireDatetime string
	F_phone_number string
	F_gender string
	F_grade string
	F_birthday string
	F_school string
	F_school_id int
	F_province string
	F_province_id int
	F_city string
	F_city_id int
	F_county string
	F_county_id int
	F_user_realname string
	F_crate_datetime string
	F_modify_datetime string
}

type MUserInfoResp struct {
	responseNo  int
	responseMsg string
	F_phone_number string
	F_gender string
	F_grade string
	F_birthday string
	F_school string
	F_school_id int
	F_province string
	F_province_id int
	F_city string
	F_city_id int
	F_county string
	F_county_id int
	F_user_realname string
	F_crate_datetime string
	F_modify_datetime string
}

type MAreaResp struct {
	responseNo  int
	responseMsg string
	areaList []struct{
		F_area_id	int
		F_area_name	string
	}
}

type MSchoolResp struct {
	responseNo  int
	responseMsg string
	schoolList []struct{
		F_school_id int
		F_school string
		F_school_type string
	}
}

type MGradeResp struct {
	responseNo  int
	responseMsg string
	gradeList []string
}