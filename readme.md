
ึ### return
- c.senString() --> return ข้อความกลับไปเป็น string
- c.Json() --> return json
- c.Status(fiber.StatusNotFound).SendString("Book not found")
    - custom return จะส่ง ststus และข้อความที่เราต้องการจะ custom

### Post
- c.Bodyparam() --> รับ body ที่มากาก api 