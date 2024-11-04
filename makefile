test_receive_otp:
	curl \
		-X POST http://0.0.0.0:3000/receive_otp \
		-H "Content-Type: application/json" \
		-d '{"user_id": 1, "otp":"123456", "service_id": 1}'

test_register_device:
	curl \
		-X POST http://0.0.0.0:3000/register_device \
		-H "Content-Type: application/json" \
		-d '{"uuid":"device-123", "fcm_token":"abcd1234fcmToken"}'

test_register_service:
	curl \
		-X POST http://0.0.0.0:3000/register_service \
		-H "Content-Type: application/json" \
		-d '{"name":"test-service", "regex": "(\\d+)", "sample_otp_message": "Hello the otp is 12345", "sample_otp": "12345", "expected_otp_ttl": 3600}'
