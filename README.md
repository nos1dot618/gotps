> Go OTP Server for integration with the [gibotp](https://github.com/hyouteki/gibotp) Android application.
> This server registers devices and requests OTPs from specified services by sending notifications to the registered device.
---

# Getting Started
1. Ensure you have Go installed and set up on your system.
2. Ensure you have correctly added all the API keys, tokens, credentials, etc. in the `creds.bashrc.tmpl` and rename it to `creds.bashrc`.
3. Give `run_server.sh` script executable permission using `chmod +x run_server.sh`.
4. Run the server and admin portal using the command `./run_server.sh`.
---

# API documentation
Server exposes two endpoints.

## 1. Receive OTP
- **URL**: `/receive_otp`
- **Method**: `POST`
- **Description**: Receives an OTP sent from the client and validates its presence in the request. If valid, a success message is returned.

### Request Body
| Field | Type   | Description            | Required |
|-------|--------|------------------------|----------|
| otp   | String | The OTP to be verified | Yes      |

### Example
```bash
curl \
    -X POST http://0.0.0.0:3000/receive_otp \
    -H "Content-Type: application/json" \
    -d '{"otp":"123456"}'
```
```json
{
    "otp": "123456"
}
```

### Responses
| Status Code                 | Message	                  | Description                                                                                |
|-----------------------------|---------------------------|--------------------------------------------------------------------------------------------|
| `200 OK`	                  | OTP received successfully | OTP was successfully received.                                                             |
| `400 Bad Request`	          | Invalid request	          | No OTP was provided in the request.                                                        |
| `500 Internal Server Error` |	Internal processing error | Unable to process the request, possibly due to database errors or foreign key constraints. |

## 2. Register Device
- **URL**: `/register_device`
- **Method**: `POST`
- **Description**: Registers a device using the provided UUID and FCMToken. If successful, the server stores the device's FCM token using the UUID as a key.

### Request Body
| Field	    | Type	 | Description	                                     | Required |
|-----------|--------|---------------------------------------------------|----------|
| uuid	    | String | The unique identifier for the device              | Yes      |
| fcm_token	| String | The Firebase Cloud Messaging token for the device | Yes      |

### Example
```bash
curl \
    -X POST http://0.0.0.0:3000/register_device \
    -H "Content-Type: application/json" \
    -d '{"uuid":"device-123", "fcm_token":"abcd1234fcmToken"}'
```
```json
{
    "uuid": "device-123",
    "fcm_token": "abcd1234fcmToken"
}
```

### Responses
| Status Code                 | Message	                       | Description                                                                                |
|-----------------------------|--------------------------------|--------------------------------------------------------------------------------------------|
| `200 OK`	                  | Device registered successfully | Device registration was successful.                                                        |
| `400 Bad Request`	          | Invalid request	               | The request is missing required fields or has invalid data.                                |
| `500 Internal Server Error` |	Internal processing error      | Unable to process the request, possibly due to database errors during registration.        |

## 3. Register Service
- **URL**: `/register_service`
- **Method**: `POST`
- **Description**: Registers a service with details including the service name, regex pattern for OTP extraction, a sample OTP message, sample OTP, and expected OTP TTL.

### Request Body
| Field	             | Type	  | Description	                                           | Required |
|--------------------|--------|--------------------------------------------------------|----------|
| name	             | String |	The name of the service	                               | Yes      |
| regex              | String | Regular expression to match OTP in messages	           | Yes      |
| sample_otp_message | String |A sample message containing an OTP for the service      | Yes      |
| sample_otp	     | String | The sample OTP for the service                         | Yes      |
| expected_otp_ttl	 | Int    | The expected TTL (time-to-live) for the OTP in seconds | Yes      |

### Example
```bash
curl \
    -X POST http://0.0.0.0:3000/register_device \
    -H "Content-Type: application/json" \
    -d '{"name":"ExampleService","regex":"\\d{6}","sample_otp_message":"Your OTP is 123456","sample_otp":"123456","expected_otp_ttl":300}'
```
```json
{
    "uuid": "device-123",
    "fcm_token": "abcd1234fcmToken"
}
```

### Responses
| Status Code                 | Message	                        | Description                                                                                 |
|-----------------------------|---------------------------------|---------------------------------------------------------------------------------------------|
| `200 OK`	                  | Service registered successfully | The service registration was successful.                                                    |
| `400 Bad Request`	          | Invalid service registration    | The request is missing required fields or has invalid data.                                 |
| `500 Internal Server Error` |	Internal processing error       | Unable to process the request, possibly due to database errors during service registration. |
