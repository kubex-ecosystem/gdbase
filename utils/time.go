package utils

const (
	// GoFormat é o formato de data e hora padrão do Go.
	GoFormat = "2006-01-02 15:04:05.999999999"
	// Format é o formato de data e hora padrão.
	Format = "2006-01-02 15:04:05"
	// TimeFormat é o formato de hora padrão.
	TimeFormat = "15:04:05"
	// DateFormat é o formato de data padrão.
	DateFormat = "2006-01-02"
	// ShortDateFormat é o formato de data curta.
	ShortDateFormat = "06-01-02"
	// ShortestDateFormat é o formato de data mais curta.
	ShortestDateFormat = "06-1-2"
	// ShortTimeFormat é o formato de hora curta.
	ShortTimeFormat = "15:04"
	// LongTimeFormat é o formato de hora longa.
	LongTimeFormat = "15:04:05"
	// ShortDateTimeFormat é o formato de data e hora curtas.
	ShortDateTimeFormat = "06-01-02 15:04"
	// DateTimeFormat é o formato de data e hora padrão.
	DateTimeFormat = "2006-01-02 15:04"
	// ISO8601Date é o formato de data ISO 8601.
	ISO8601Date = "2006-01-02"
	// ISO8601Time é o formato de hora ISO 8601.
	ISO8601Time = "15:04:05"
	// ISO8601TimeMs é o formato de hora com milissegundos ISO 8601.
	ISO8601TimeMs = "15:04:05.999"
	// ISO8601DateTime é o formato de data e hora ISO 8601.
	ISO8601DateTime = "2006-01-02T15:04:05"
	// ISO8601TZ é o formato de data e hora com fuso horário ISO 8601.
	ISO8601TZ = "2006-01-02T15:04:05-0700"
	// ISO8601TZs é o formato de data e hora com fuso horário (Z) ISO 8601.
	ISO8601TZs = "2006-01-02T15:04:05Z0700"
	// ISO8601TZms é o formato de data e hora com milissegundos e fuso horário ISO 8601.
	ISO8601TZms = "2006-01-02T15:04:05.999-0700"
	// HourMinuteFormat é o formato de hora e minuto.
	HourMinuteFormat = "15:04"
	// HourFormat é o formato de hora.
	HourFormat = "15"
	// FormattedDateFormat é o formato de data formatada.
	FormattedDateFormat = "Jan 2, 2006"
	// DayDateTimeFormat é o formato de data e hora com dia da semana.
	DayDateTimeFormat = "Mon, Aug 2, 2006 3:04 PM"
	// ISO8601Format é o formato de data e hora ISO 8601.
	ISO8601Format = "2006-01-02T15:04:05-0700"
	// CookieFormat é o formato de data e hora para cookies.
	CookieFormat = "Monday, 02-Jan-2006 15:04:05 MST"
	// RFC822Format é o formato de data e hora RFC 822.
	RFC822Format = "Mon, 02 Jan 06 15:04:05 -0700"
	// RFC1036Format é o formato de data e hora RFC 1036.
	RFC1036Format = "Mon, 02 Jan 06 15:04:05 -0700"
	// RFC2822Format é o formato de data e hora RFC 2822.
	RFC2822Format = "Mon, 02 Jan 2006 15:04:05 -0700"
	// RFC3339Format é o formato de data e hora RFC 3339.
	RFC3339Format = "2006-01-02T15:04:05-07:00"
	// RSSFormat é o formato de data e hora RSS.
	RSSFormat = "Mon, 02 Jan 2006 15:04:05 -0700"
	// W3CFormat é o formato de data e hora W3C.
	W3CFormat = "2006-01-02T15:04:05-07:00"
	// UnixFormat é o formato de data e hora Unix.
	UnixFormat = "Mon Jan _2 15:04:05 MST 2006"
	// UnixDate é o formato de data e hora Unix.
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	// UnixDate2 é o formato de data e hora Unix.
	UnixDate2 = "Mon Jan 02 15:04:05 MST 2006"
	// UnixDate3 é o formato de data e hora Unix.
	UnixDate3 = "Mon Jan 02 15:04:05 -0700 2006"
	// UnixDate4 é o formato de data e hora Unix.
	UnixDate4 = "Mon Jan 02 15:04:05 -0700 MST 2006"
	// UnixDate5 é o formato de data e hora Unix.
	UnixDate5 = "Mon Jan 02 15:04:05 -0700 (MST) 2006"

	// TimezoneUTC é o fuso horário UTC.
	TimezoneUTC = "UTC"
	// TimezoneGMT é o fuso horário GMT.
	TimezoneGMT = "GMT"
	// TimezoneEST é o fuso horário EST.
	TimezoneEST = "EST"
	// TimezoneEDT é o fuso horário EDT.
	TimezoneEDT = "EDT"
	// TimezoneCST é o fuso horário CST.
	TimezoneCST = "CST"
	// TimezoneCDT é o fuso horário CDT.
	TimezoneCDT = "CDT"
	// TimezoneMST é o fuso horário MST.
	TimezoneMST = "MST"
	// TimezoneMDT é o fuso horário MDT.
	TimezoneMDT = "MDT"
	// TimezonePST é o fuso horário PST.
	TimezonePST = "PST"
	// TimezonePDT é o fuso horário PDT.
	TimezonePDT = "PDT"
	// TimezoneCET é o fuso horário CET.
	TimezoneCET = "CET"
	// TimezoneCEST é o fuso horário CEST.
	TimezoneCEST = "CEST"
	// TimezoneJST é o fuso horário JST.
	TimezoneJST = "JST"
	// TimezoneKST é o fuso horário KST.
	TimezoneKST = "KST"
	// TimezoneSGT é o fuso horário SGT.
	TimezoneSGT = "SGT"
	// TimezoneHKT é o fuso horário HKT.
	TimezoneHKT = "HKT"
	// TimezoneAEST é o fuso horário AEST.
	TimezoneAEST = "AEST"
	// TimezoneACST é o fuso horário ACST.
	TimezoneACST = "ACST"
	// TimezoneAWST é o fuso horário AWST.
	TimezoneAWST = "AWST"
	// TimezoneNZST é o fuso horário NZST.
	TimezoneNZST = "NZST"

	// TimezoneOffsetUTC é o offset do fuso horário UTC.
	TimezoneOffsetUTC = "+0000"
	// TimezoneOffsetGMT é o offset do fuso horário GMT.
	TimezoneOffsetGMT = "+0000"
	// TimezoneOffsetEST é o offset do fuso horário EST.
	TimezoneOffsetEST = "-0500"
	// TimezoneOffsetEDT é o offset do fuso horário EDT.
	TimezoneOffsetEDT = "-0400"
	// TimezoneOffsetCST é o offset do fuso horário CST.
	TimezoneOffsetCST = "-0600"
	// TimezoneOffsetCDT é o offset do fuso horário CDT.
	TimezoneOffsetCDT = "-0500"
	// TimezoneOffsetMST é o offset do fuso horário MST.
	TimezoneOffsetMST = "-0700"
	// TimezoneOffsetMDT é o offset do fuso horário MDT.
	TimezoneOffsetMDT = "-0600"
	// TimezoneOffsetPST é o offset do fuso horário PST.
	TimezoneOffsetPST = "-0800"
	// TimezoneOffsetPDT é o offset do fuso horário PDT.
	TimezoneOffsetPDT = "-0700"
	// TimezoneOffsetCET é o offset do fuso horário CET.
	TimezoneOffsetCET = "+0100"
	// TimezoneOffsetCEST é o offset do fuso horário CEST.
	TimezoneOffsetCEST = "+0200"
	// TimezoneOffsetJST é o offset do fuso horário JST.
	TimezoneOffsetJST = "+0900"
	// TimezoneOffsetKST é o offset do fuso horário KST.
	TimezoneOffsetKST = "+0900"
	// TimezoneOffsetSGT é o offset do fuso horário SGT.
	TimezoneOffsetSGT = "+0800"
	// TimezoneOffsetHKT é o offset do fuso horário HKT.
	TimezoneOffsetHKT = "+0800"
	// TimezoneOffsetAEST é o offset do fuso horário AEST.
	TimezoneOffsetAEST = "+1000"
	// TimezoneOffsetACST é o offset do fuso horário ACST.
	TimezoneOffsetACST = "+0930"
	// TimezoneOffsetAWST é o offset do fuso horário AWST.
	TimezoneOffsetAWST = "+0800"
	// TimezoneOffsetNZST é o offset do fuso horário NZST.
	TimezoneOffsetNZST = "+1200"
)

// GetTimezoneOffset retorna o offset do fuso horário.
// timezone: o fuso horário.
// Retorna uma string contendo o offset do fuso horário.
func GetTimezoneOffset(timezone string) string {
	switch timezone {
	case TimezoneUTC:
		return TimezoneOffsetUTC
	case TimezoneGMT:
		return TimezoneOffsetGMT
	case TimezoneEST:
		return TimezoneOffsetEST
	case TimezoneEDT:
		return TimezoneOffsetEDT
	case TimezoneCST:
		return TimezoneOffsetCST
	case TimezoneCDT:
		return TimezoneOffsetCDT
	case TimezoneMST:
		return TimezoneOffsetMST
	case TimezoneMDT:
		return TimezoneOffsetMDT
	case TimezonePST:
		return TimezoneOffsetPST
	case TimezonePDT:
		return TimezoneOffsetPDT
	case TimezoneCET:
		return TimezoneOffsetCET
	case TimezoneCEST:
		return TimezoneOffsetCEST
	case TimezoneJST:
		return TimezoneOffsetJST
	case TimezoneKST:
		return TimezoneOffsetKST
	case TimezoneSGT:
		return TimezoneOffsetSGT
	case TimezoneHKT:
		return TimezoneOffsetHKT
	case TimezoneAEST:
		return TimezoneOffsetAEST
	case TimezoneACST:
		return TimezoneOffsetACST
	case TimezoneAWST:
		return TimezoneOffsetAWST
	case TimezoneNZST:
		return TimezoneOffsetNZST
	default:
		return ""
	}
}

// ConvertTimezone retorna o nome completo do fuso horário.
// timezone: o fuso horário.
// Retorna uma string contendo o nome completo do fuso horário.
func ConvertTimezone(timezone string) string {
	switch timezone {
	case TimezoneUTC:
		return "Coordinated Universal Time"
	case TimezoneGMT:
		return "Greenwich Mean Time"
	case TimezoneEST:
		return "Eastern Standard Time"
	case TimezoneEDT:
		return "Eastern Daylight Time"
	case TimezoneCST:
		return "Central Standard Time"
	case TimezoneCDT:
		return "Central Daylight Time"
	case TimezoneMST:
		return "Mountain Standard Time"
	case TimezoneMDT:
		return "Mountain Daylight Time"
	case TimezonePST:
		return "Pacific Standard Time"
	case TimezonePDT:
		return "Pacific Daylight Time"
	case TimezoneCET:
		return "Central European Time"
	case TimezoneCEST:
		return "Central European Summer Time"
	case TimezoneJST:
		return "Japan Standard Time"
	case TimezoneKST:
		return "Korea Standard Time"
	case TimezoneSGT:
		return "Singapore Time"
	case TimezoneHKT:
		return "Hong Kong Time"
	case TimezoneAEST:
		return "Australian Eastern Standard Time"
	case TimezoneACST:
		return "Australian Central Standard Time"
	case TimezoneAWST:
		return "Australian Western Standard Time"
	case TimezoneNZST:
		return "New Zealand Standard Time"
	default:
		return ""
	}
}

// ConvertTimezoneOffset retorna o offset do fuso horário no formato UTC.
// timezone: o fuso horário.
// Retorna uma string contendo o offset do fuso horário no formato UTC.
func ConvertTimezoneOffset(timezone string) string {
	switch timezone {
	case TimezoneUTC:
		return "UTC+00:00"
	case TimezoneGMT:
		return "GMT+00:00"
	case TimezoneEST:
		return "UTC-05:00"
	case TimezoneEDT:
		return "UTC-04:00"
	case TimezoneCST:
		return "UTC-06:00"
	case TimezoneCDT:
		return "UTC-05:00"
	case TimezoneMST:
		return "UTC-07:00"
	case TimezoneMDT:
		return "UTC-06:00"
	case TimezonePST:
		return "UTC-08:00"
	case TimezonePDT:
		return "UTC-07:00"
	case TimezoneCET:
		return "UTC+01:00"
	case TimezoneCEST:
		return "UTC+02:00"
	case TimezoneJST:
		return "UTC+09:00"
	case TimezoneKST:
		return "UTC+09:00"
	case TimezoneSGT:
		return "UTC+08:00"
	case TimezoneHKT:
		return "UTC+08:00"
	case TimezoneAEST:
		return "UTC+10:00"
	case TimezoneACST:
		return "UTC+09:30"
	case TimezoneAWST:
		return "UTC+08:00"
	case TimezoneNZST:
		return "UTC+12:00"
	default:
		return ""
	}
}

// GetWeekdayByAnyType retorna o dia da semana a partir de qualquer tipo.
// v: o valor que representa o dia da semana.
// Retorna uma string contendo o dia da semana.
func GetWeekdayByAnyType(v interface{}) string {
	switch v.(type) {
	case string:
		return GetWeekday(v.(string))
	case int:
		return GetWeekdayByInt(v.(int))
	default:
		return ""
	}
}

// GetWeekday retorna o dia da semana a partir de uma string.
// weekday: a string que representa o dia da semana.
// Retorna uma string contendo o dia da semana em português.
func GetWeekday(weekday string) string {
	switch weekday {
	case "Sunday":
		return "Domingo"
	case "Monday":
		return "Segunda-feira"
	case "Tuesday":
		return "Terça-feira"
	case "Wednesday":
		return "Quarta-feira"
	case "Thursday":
		return "Quinta-feira"
	case "Friday":
		return "Sexta-feira"
	case "Saturday":
		return "Sábado"
	default:
		return ""
	}
}

// GetWeekdayByInt retorna o dia da semana a partir de um inteiro.
// weekday: o inteiro que representa o dia da semana (0 para Domingo, 1 para Segunda-feira, etc.).
// Retorna uma string contendo o dia da semana em português.
func GetWeekdayByInt(weekday int) string {
	switch weekday {
	case 0:
		return "Domingo"
	case 1:
		return "Segunda-feira"
	case 2:
		return "Terça-feira"
	case 3:
		return "Quarta-feira"
	case 4:
		return "Quinta-feira"
	case 5:
		return "Sexta-feira"
	case 6:
		return "Sábado"
	default:
		return ""
	}
}

// ExtractTime extrai a hora de uma string.
// s: a string que contém a hora.
// Retorna uma string contendo a hora extraída.
func ExtractTime(s string) string {
	return ExtractDateTime(s, TimeFormat)
}

// ExtractDate extrai a data de uma string.
// s: a string que contém a data.
// Retorna uma string contendo a data extraída.
func ExtractDate(s string) string {
	return ExtractDateTime(s, DateFormat)
}

// ExtractDateTime extrai a data e a hora de uma string.
// s: a string que contém a data e a hora.
// format: o formato da data e hora a ser extraído.
// Retorna uma string contendo a data e hora extraída.
func ExtractDateTime(s string, format string) string {
	if len(s) < len(format) {
		return s
	}
	return s[:len(format)]
}

// FormatTime formata a hora.
// t: a string que contém a hora.
// Retorna uma string contendo a hora formatada.
func FormatTime(t string) string {
	return FormatDateTime(t, TimeFormat)
}

// FormatDate formata a data.
// t: a string que contém a data.
// Retorna uma string contendo a data formatada.
func FormatDate(t string) string {
	return FormatDateTime(t, DateFormat)
}

// FormatDateTime formata a data e a hora.
// t: a string que contém a data e a hora.
// format: o formato da data e hora a ser formatado.
// Retorna uma string contendo a data e hora formatada.
func FormatDateTime(t string, format string) string {
	if len(t) < len(format) {
		return t
	}
	return t[:len(format)]
}
