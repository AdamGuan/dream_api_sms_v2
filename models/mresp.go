package models

type MResp struct {
	responseNo  int
	responseMsg string
}

type MRegisterResp struct {
	responseNo  int
	responseMsg string
	F_uid string
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
	F_uid string
	F_phone_number string
	F_gender string
	F_gender_id int
	F_grade string
	F_grade_id int
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
	F_user_nickname string
	F_crate_datetime string
	F_modify_datetime string
	F_class_id int
	F_class_name string
	F_avatar_url string
	F_user_email string
	F_coin int
}

type MUserInfoResp struct {
	responseNo  int
	responseMsg string
	F_uid string
	F_phone_number string
	F_gender string
	F_gender_id int
	F_grade string
	F_grade_id int
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
	F_user_nickname string
	F_crate_datetime string
	F_modify_datetime string
	F_class_id int
	F_class_name string
	F_avatar_url string
	F_user_email string
	F_coin int
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

type MClassAddResp struct {
	responseNo  int
	responseMsg string
	F_class_id int
}

type MClassListInfoResp struct {
	responseNo  int
	responseMsg string
	classList []struct{
		F_class_id int
		F_class_name string
		F_class_person_total int
	}
}

type MSchoolAreaInfoResp struct {
	responseNo  int
	responseMsg string
	areaInfoList map[string]struct{
		F_school_id int
		F_area_province_id int
		F_area_province_name string
		F_area_city_id int
		F_area_city_name string
		F_area_county_id int
		F_area_county_name string
	}
}

type MSchoolAreaInfoItemResp map[string]struct{
		F_school_id int
		F_area_province_id int
		F_area_province_name string
		F_area_city_id int
		F_area_city_name string
		F_area_county_id int
		F_area_county_name string
}

type MAreaInfoResp struct {
	responseNo  int
	responseMsg string
	areaInfoList map[string]struct{
		F_area_id int
		F_area_name string
	}
}

type MAvatarlistResp struct {
	responseNo  int
	responseMsg string
	avatarList []struct{
		F_avatar_url string
		F_avatar_id int
	}
}

type MModifyPhoneResp struct {
	responseNo  int
	responseMsg string
	token string
	tokenExpireDatetime string
}

type MModifyEmailResp struct {
	responseNo  int
	responseMsg string
	token string
	tokenExpireDatetime string
}

type MModifyCoinResp struct{
	responseNo  int
	responseMsg string
	F_newCoin int
}

type MGetCoinResp struct{
	responseNo  int
	responseMsg string
	F_coin int
}

type MUserAvatarResp struct {
	responseNo  int
	responseMsg string
	F_avatar_url string
}