package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ApproveStatusParse(input string) string {

	status := strings.ToLower(input)
	if status == "w" || status == "wait" {
		return "Wait"
	}

	return strings.ToUpper(input)
}

func RemoveIntroZero(input string) string {
	num, err := strconv.Atoi(input)
	if err != nil {
		return input
	}

	return strconv.Itoa(num)
}

func CleanPhoneNumber(phoneNumber string) string {
	extPatterns := []string{"ต่อ", "Ext.", "Ext"}
	for _, pattern := range extPatterns {
		if idx := strings.Index(phoneNumber, pattern); idx != -1 {
			phoneNumber = phoneNumber[:idx]
			break
		}
	}

	if strings.HasPrefix(phoneNumber, "+66") {
		phoneNumber = "0" + phoneNumber[3:]
	}

	re := regexp.MustCompile("[^0-9]+")
	phoneNumber = re.ReplaceAllString(phoneNumber, "")

	return phoneNumber
}

func IsContainsCannotUnmarshalArray(input string) string {
	if strings.Contains(input, "cannot unmarshal array") {
		return "Invalid request"
	} else {
		return input
	}
}

func AddLeadingZeroIfSingleDigit(input string) string {
	// พยายามแปลงสตริงเป็นจำนวนเต็ม
	num, err := strconv.Atoi(input)
	if err != nil {
		// ถ้าแปลงไม่ได้ (ไม่ใช่ตัวเลข) ให้รีเทิร์นค่าเดิม
		return input
	}

	// ถ้าเป็นตัวเลขหลักเดียว (0-9) ให้เติม 0 ข้างหน้า
	if num >= 0 && num <= 9 {
		return "0" + input
	}

	// ถ้าไม่ใช่ตัวเลขหลักเดียว ให้รีเทิร์นค่าเดิม
	return input
}

func FormatSubdistrictCode(input string) string {
	// พยายามแปลงสตริงเป็นจำนวนเต็ม
	_, err := strconv.Atoi(input)
	if err != nil {
		// ถ้าแปลงไม่ได้ (ไม่ใช่ตัวเลข) ให้รีเทิร์นค่าเดิม
		return input
	}
	// ถ้าความยาวของสตริงน้อยกว่า 3 ให้รีเทิร์นค่าเดิม
	if len(input) < 3 {
		return input
	}

	// ตัด 0 ออกจากทางขวาสุด 2 หลัก
	input = strings.TrimRight(input, "0")
	if len(input) == 0 {
		return "00"
	}

	// ถ้าหลังจากตัด 0 แล้วเหลือเลขหลักเดียว ให้เติม 0 ข้างหน้า
	if len(input) == 1 {
		return "0" + input
	}

	// รีเทิร์นผลลัพธ์ที่เหลือ
	return input
}

func RemoveSpecialStrings(input string) string {
	// คำเฉพาะที่ต้องการตัดออก
	specialStrings := []string{"ต.", "ตำบล", "แขวง", "อ.", "อำเภอ", "เขต", "จ.", "จังหวัด", " "}
	// แทนที่คำเฉพาะด้วย string ว่าง
	for _, s := range specialStrings {
		input = strings.ReplaceAll(input, s, "")
	}
	return input
}

func StringToFloat64(input string) (out float64) {
	if out, err := strconv.ParseFloat(input, 64); err == nil {
		return out
	}
	return 0
}

func DetectLanguage(text string) string {
	thaiPattern := `[ก-ฮ]`
	if matched, _ := regexp.MatchString(thaiPattern, text); matched {
		return "TH"
	}
	englishPattern := `[a-zA-Z]`
	if matched, _ := regexp.MatchString(englishPattern, text); matched {
		return "EN"
	}
	return ""

}
func PadNumCBS(cifNumber string) string {
	parseCifNumber, _ := strconv.Atoi(cifNumber)
	return fmt.Sprintf("%012d", parseCifNumber)
}
