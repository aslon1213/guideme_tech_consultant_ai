class User:
    def __init__(self) -> None:
        self.id = None
        self.first_name = ""
        self.last_name = ""
        self.email = ""
        self.passport_number = ""
        self.money = 0

    def SetUserData(self, id, first_name, second_name, email, passport_number, money):
        self.id = id
        self.first_name = first_name
        self.last_name = second_name
        self.email = email
        self.passport_number = passport_number
        self.money = money

    def SetUserDataFromJson(self, data):
        self.id = data["id"]
        self.first_name = data["first_name"]
        self.last_name = data["last_name"]
        self.email = data["email"]
        self.passport_number = data["passport_number"]
        self.money = data["money"]
