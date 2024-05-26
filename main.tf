provider "example" {
    address = "http://localhost"
    port    = "3001"
    token   = "testToken"
}

resource "test_item" "test1" {
    name    = "item_1"
    description = "test item"
    tags = [
        "hello",
        "world"
    ]
}