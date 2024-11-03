test_receive_otp:
	curl \
		-X POST http://0.0.0.0:3000/receive_otp \
		-H "Content-Type: application/json" \
		-d '{"otp":"123456"}'

test_register_device:
	curl \
		-X POST http://0.0.0.0:3000/register_device \
		-H "Content-Type: application/json" \
		-d '{"uuid":"device-123", "fcm_token":"abcd1234fcmToken"}'
