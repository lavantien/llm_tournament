from database import Database

def main():
    db = Database()
    db.populate_mock_data()
    print("Mock data populated successfully.")

if __name__ == "__main__":
    main()
