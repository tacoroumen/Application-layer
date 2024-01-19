# Gatekeeper API

This repository contains the code for the Gatekeeper API, which is responsible for managing user reservations and access to parking based on license plates.

## Gatekeeper API (gatekeeper_api/main.go)

### Configuration

The API configuration is stored in the `config/config.json` file. Make sure to provide the necessary details such as the database connection, server IP, and port.

```json
{
  "username": "your_db_username",
  "password": "your_db_password",
  "ip": "your_db_ip",
  "port": "your_db_port",
  "database": "your_database_name"
}
```

### Endpoints

#### 1. `/licenseplate`

- **Method:** `GET`
- **Parameters:** `licenseplate` (Query parameter)
- **Description:** Retrieves user information based on the provided license plate.

#### 2. `/reservering`

- **Method:** `POST`
- **Parameters:** `checkin`, `checkout`, `housenumber`, `email`, `password` (Query parameters)
- **Description:** Adds a reservation for a user.

#### 3. `/user/add`

- **Method:** `POST`
- **Parameters:** User details (Query parameters)
- **Description:** Adds a new user.

#### 4. `/user/modify`

- **Method:** `POST`
- **Parameters:** User details for modification (Query parameters)
- **Description:** Modifies user details.

#### 5. `/user/delete`

- **Method:** `POST`
- **Parameters:** `email`, `password` (Query parameters)
- **Description:** Deletes a user.

#### 6. `/user/get`

- **Method:** `GET`
- **Parameters:** `email`, `password` (Query parameters)
- **Description:** Retrieves user details based on email and password.

#### 7. `/login`

- **Method:** `GET`
- **Parameters:** `email`, `password` (Query parameters)
- **Description:** Validates user credentials for login.

### How to Run

Ensure you have Go installed on your machine. Execute the following command in the `gatekeeper_api` directory:

```bash
go run main.go
```

## ESPhome Configuration (esphome32/gatekeeper.yaml)

This YAML configuration file is for the ESP32 microcontroller, which controls the gate and communicates with the Gatekeeper API.

Ensure you provide your WiFi credentials, OTA password, and other required details in the `!secret` placeholders.

## Gatekeeper (gatekeeper/main.go)

This Go program checks the access of a license plate against the Gatekeeper API and controls the gate accordingly. It logs access details and sends commands to the ESP32.

Ensure you provide the necessary details in the `config.json` file, and execute the program with the license plate as a command-line argument:

```bash
go run main.go -plate YOUR_LICENSE_PLATE
```

## License Plate Recognition (gatekeeper/gatekeeper.py)

This Python script captures video feed from the webcam, performs license plate recognition using OpenCV and Tesseract OCR, and communicates with the Gatekeeper API.

Ensure you have OpenCV and Tesseract installed. Run the script:

```bash
python gatekeeper.py
```

Press 'q' to exit the script.

## Port Scanner

This repository contains a simple port-scanner written in Go, which allows you to scan open ports on a target IP address.

### Usage

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/tacoroumen/Application-layer
   ```

2. **Navigate to the port-scanner directory:**
   ```bash
   cd port-scanner
   ```

3. **Build the Application:**
   ```bash
   go build
   ```

4. **Run the Port Scanner:**
   ```bash
   ./port-scanner <target-IP>
   ```
   Replace `<target-IP>` with the IP address you want to scan.

### Example

```bash
./port-scanner 192.168.1.1
```

This will scan all TCP ports from 1 to 65535 on the specified IP address and report open ports.

### Customization

You can customize the port range, number of workers, and protocol by modifying the variables in the `main` function of `main.go`. Adjust the `startPort`, `endPort`, and `numWorkers` variables based on your requirements.

### Disclaimer

This port scanner is designed for educational and authorized testing purposes only. Unauthorized scanning of networks or systems may violate applicable laws and regulations. Use this tool responsibly and only on systems for which you have explicit permission. The developers and maintainers of this tool are not responsible for any misuse or unauthorized activities.

## License

This project is licensed under the [MIT License](LICENSE).