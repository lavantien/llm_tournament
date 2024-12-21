import sys
from PyQt6.QtWidgets import QApplication
from views import MainWindow
from db_manager import DBManager
import os

def main():
    app = QApplication(sys.argv)
    db_manager = DBManager(os.path.join(os.path.dirname(__file__), '..', 'db', 'benchmark.db'))
    main_window = MainWindow(db_manager)
    main_window.show()
    sys.exit(app.exec())

if __name__ == "__main__":
    main()
