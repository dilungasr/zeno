# **Zeno Package for Go (Golang)**

_Effortlessly Integrate Payments with Zenopay_

This package is an elegant solution for developers in the **Golang community**, designed and maintained by **Dilunga the Great**. The package simplifies interaction with the [Zenopay API](https://zenopay.net), making it easy to create payment orders and manage their statuses.

---

## **Features**

- Effortless integration with **Zenopay API**.
- Background polling to check payment statuses.
- Customizable callbacks for handling successful or failed payments.
- Open-source and community-driven by **Dilunga the Great**.

---

## **Meet the Author**

**Dilunga the Great** is a Tanzanian software engineer passionate about empowering people with tools that simplify their work. With expertise in Golang and a strong focus on simplifying tasks through code, he delivers this package to the Golang community to facilitate seamless payment workflows.

- **Contact Details**:
  - WhatsApp: [+255785453830](https://wa.me/255785453830)
  - Email: [thegreatdilunga@gmail.com](mailto:thegreatdilunga@gmail.com)
  - GitHub: [github.com/dilungasr](https://github.com/dilungasr)

> Feel free to reach out for assistance, feedback, or collaboration opportunities!

---

## **How It Works**

### **Step 1: Registration with Zenopay**

To start using this package, register and verify your account with [Zenopay](https://zenopay.net). After successful registration, you'll receive the following credentials:

- **Account ID**
- **API Key**
- **Secret Key**

### **Step 2: Install the Package**

```bash
go get github.com/dilungasr/zeno
```

---

## **Quick Start Guide**

### **1. Configure Zenopay API Credentials**

Before creating a payment, configure the package with your API credentials.

```go
package main

import "github.com/dilungasr/zeno"

func main() {
    zeno.ConfigAPI("your-account-id", "your-api-key", "your-secret-key")
}
```

### **2. Create a Payment Order**

Use the `Pay` function to create a payment order. It accepts payment details and a callback function for handling success or failure.

```go
package main

import (
    "fmt"
    "github.com/dilungasr/zeno"
)

func main() {
    // Configure Zenopay API
    zeno.ConfigAPI("your-account-id", "your-api-key", "your-secret-key")

    // Create a payment order
    orderID, err := zeno.Pay("1500", "John Doe", "0712345678", "john.doe@gmail.com", func(orderID string, success bool) {
        if success {
            fmt.Printf("Payment for Order %s was successful!\n", orderID)
        } else {
            fmt.Printf("Payment for Order %s failed.\n", orderID)
        }
    })

    if err != nil {
        fmt.Printf("Error creating payment order: %s\n", err.Error())
        return
    }

    fmt.Printf("Payment order created! Order ID: %s\n", orderID)
}
```

---

## **Detailed Function Documentation**

### **ConfigAPI**

```go
func ConfigAPI(accountID, APIKey, secreteKey string)
```

Configures the Zenopay API with your credentials.

- **Parameters**:
  - `accountID`: Your Zenopay account ID.
  - `APIKey`: Your Zenopay API key.
  - `secreteKey`: Your Zenopay secret key.

---

### **Pay**

```go
func Pay(amount, name, phone, email string, callback func(orderID string, ok bool)) (orderID string, err error)
```

Creates a payment order.

- **Parameters**:

  - `amount`: Payment amount (string).
  - `name`: Buyer's name.
  - `phone`: Buyer's phone number.
  - `email`: Buyer's email.
  - `callback`: Function called when the payment status is determined.

- **Returns**:
  - `orderID`: Unique identifier for the payment order.
  - `err`: Error, if any.

---

### **Polling in the Background**

The package automatically polls the Zenopay API to check the payment status.

- **Polling Interval**: Every 5 seconds.
- **Timeout**: 50 seconds (aligns with Zenopay's USSD session expiry).
- The `callback` function is invoked when:
  - The payment succeeds (`ok = true`).
  - The payment fails (`ok = false`).

---

### **Example Callback Function**

```go
func handlePaymentStatus(orderID string, success bool) {
    if success {
        fmt.Printf("Order %s: Payment completed successfully.\n", orderID)
        // Perform actions such as updating your database.
    } else {
        fmt.Printf("Order %s: Payment failed.\n", orderID)
        // Handle failure, notify the user, or retry.
    }
}
```

---

## **Best Practices**

- Ensure your callback function handles both success and failure scenarios.
- Use secure storage for your API credentials.
- Polling runs in the background, so you can scale your HTTP handlers without blocking.

---

## **Why 50-Second Polling Timeout?**

Zenopay automatically expires a user's USSD prompt session after 50 seconds. Polling every 5 seconds is the optimal interval to balance system performance and ensure timely updates.

---

### Support and Contact

This package is written by **Dilunga the Great** for the Golang community and is open source.

- **GitHub Repository**: [github.com/dilungasr/zeno](https://github.com/dilungasr/zeno)
- **Zenopay Official Website**: [https://zenopay.net/](https://zenopay.net/)
- **Contact Author**:
  - WhatsApp: +255785453830
  - Email: thegreatdilunga@gmail.com

Feel free to contribute, report issues, or contact me for further assistance.

## **License**

This package is **open-source**, created for the Go community by **Dilunga the Great**.  
Feel free to contribute, report issues, or suggest enhancements on [GitHub](https://github.com/dilungasr/zeno).

> "Code is poetry, and simplicity is key." — Dilunga the Great

**Happy Coding!** 🎉
