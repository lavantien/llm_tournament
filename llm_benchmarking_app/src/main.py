```python
import sys
from PyQt6.QtWidgets import QApplication
from views import MainWindow
from db_manager import DBManager

def main():
    app = QApplication(sys.argv)
    db_manager = DBManager('db/benchmark.db')
    main_window = MainWindow(db_manager)
    main_window.show()
    sys.exit(app.exec())

if __name__ == "__main__":
    main()
```
