package database

import (
	"log"
)

type User struct {
	DeviceUuid string `json:"uuid"`
	DeviceFcmToken string `json:"fcm_token"`
}

type Otp struct {
	UserId int `json:"user_id"`
    Otp string `json:"otp"`
	ServiceId int `json:"service_id"`
}

type Service struct {
	Name string `json:"name"`
	Regex string `json:"regex"`
	SampleOtpMessage string `json:"sample_otp_message"`
	SampleOtp string `json:"sample_otp"`
	ExpectedOtpTtl int `json:"expected_otp_ttl"`
}

func RegisterDevice(user User) error {
	var query string = `INSERT OR REPLACE INTO users (device_uuid, device_fcm_token) ` +
		`VALUES (?, ?)`
	_, err := db.Exec(query, user.DeviceUuid, user.DeviceFcmToken)
	if err != nil {
		log.Printf("error: failed to insert user into database: %v\n", err)
		return err
	}
	return nil
}

func InsertOtp(otp Otp) error {
	var query string = `INSERT OR REPLACE INTO otps (user_id, otp, service_id, timestamp) ` +
		`VALUES (?, ?, ?, CURRENT_TIMESTAMP)`
	_, err := db.Exec(query, otp.UserId, otp.Otp, otp.ServiceId);
	if err != nil {
		log.Printf("error: failed to insert otp into database: %v\n", err)
		return err
	}
	return nil
}

func RegisterService(service Service) error {
	var query string = `INSERT OR REPLACE INTO services ` +
		`(service_name, regex, sample_otp_message, sample_otp, expected_otp_ttl) ` +
		`VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, service.Name, service.Regex, service.SampleOtpMessage,
		service.SampleOtp, service.ExpectedOtpTtl);
	if err != nil {
		log.Printf("error: failed to insert service into database: %v\n", err)
		return err
	}
	return nil
}
