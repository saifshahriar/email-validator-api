# Email Validator API

A lightweight Go HTTP API that validates email addresses by checking their
**syntax** and **domain deliverability** (MX records).

## Features

- **Syntax validation** — checks the email against standard RFC-style formatting
  rules
- **Domain (MX) validation** — queries DNS for Mail Exchange (MX) records on the
  email's domain
- **Deliverability check** — reports whether the email passes both checks
- JSON responses with zero external dependencies beyond `godotenv`

## Requirements

- Go 1.26 or later

## Installation & Usage

```bash
git clone https://github.com/saifshahriar/email-validator-api.git
cd email-validator-api
go run .
```

The server starts on `http://localhost:5555` by default.

### Configuration

| Variable | Default | Description                |
| -------- | ------- | -------------------------- |
| `PORT`   | `5555`  | Port the server listens on |

Set the port via environment variable:

```bash
PORT=8080 go run .
```

Optional `.env` file support is available by setting
`LOAD_DOTENV_FROM_FILE = true` in `main.go`.

## Endpoints

### `GET /` — Health check

Returns `200 OK` with a status message.

```json
{
  "message": "server is running",
  "status": "ok"
}
```

### `GET /user?email=example@example.com` — Validate an email

Validates the email's syntax and domain MX records.

| Parameter | Type   | Required | Description           |
| --------- | ------ | -------- | --------------------- |
| `email`   | string | Yes      | The email to validate |

**Example request:**

```bash
curl "http://localhost:5555/user?email=test@gmail.com"
```

**Example response:**

```json
{
  "email": "test@gmail.com",
  "syntax": true,
  "domain": true,
  "deliverable": true
}
```

**Error responses:**

- `400 Bad Request` with message `email is required` if the `email` parameter is
  missing

## Response Fields

| Field         | Type   | Description                                                                                 |
| ------------- | ------ | ------------------------------------------------------------------------------------------- |
| `email`       | string | The email address that was validated                                                        |
| `syntax`      | bool   | Whether the email matches valid syntax                                                      |
| `domain`      | bool   | Whether the domain has valid MX records                                                     |
| `deliverable` | bool   | Whether the email passed both checks (`true` only if both `syntax` and `domain` are `true`) |

## Roadmap

- Batch validation with Goroutines (commented-out example in `handlers.go`)
- SMTP-level verification for stronger deliverability checks

## License

MIT
