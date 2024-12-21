from PyQt6.QtWidgets import QMainWindow, QTabWidget, QVBoxLayout, QWidget, QListWidget, QPushButton, QHBoxLayout, QFormLayout, QLineEdit, QTableWidget, QChartView, QBarSeries, QBarSet, QBarCategoryAxis, QValueAxis
from PyQt6.QtChart import QChart, QBarSeries, QBarSet, QBarCategoryAxis, QValueAxis
from PyQt6.QtCore import Qt
from db_manager import DBManager

class MainWindow(QMainWindow):
    def __init__(self, db_manager: DBManager):
        super().__init__()
        self.db_manager = db_manager
        self.setWindowTitle("LLM Benchmarking App")
        self.setGeometry(100, 100, 800, 600)

        self.tab_widget = QTabWidget()
        self.setCentralWidget(self.tab_widget)

        self.prompt_management_tab = PromptManagementTab(self.db_manager)
        self.score_entry_tab = ScoreEntryTab(self.db_manager)
        self.leaderboard_tab = LeaderboardTab(self.db_manager)
        self.visualization_tab = VisualizationTab(self.db_manager)

        self.tab_widget.addTab(self.prompt_management_tab, "Prompt Management")
        self.tab_widget.addTab(self.score_entry_tab, "Score Entry")
        self.tab_widget.addTab(self.leaderboard_tab, "Leaderboard")
        self.tab_widget.addTab(self.visualization_tab, "Visualization")

class PromptManagementTab(QWidget):
    def __init__(self, db_manager: DBManager):
        super().__init__()
        self.db_manager = db_manager
        self.layout = QVBoxLayout()

        self.prompt_list = QListWidget()
        self.layout.addWidget(self.prompt_list)

        self.button_layout = QHBoxLayout()
        self.load_button = QPushButton("Load Prompts")
        self.edit_button = QPushButton("Edit Prompt")
        self.save_button = QPushButton("Save Prompts")

        self.button_layout.addWidget(self.load_button)
        self.button_layout.addWidget(self.edit_button)
        self.button_layout.addWidget(self.save_button)

        self.layout.addLayout(self.button_layout)
        self.setLayout(self.layout)

class ScoreEntryTab(QWidget):
    def __init__(self, db_manager: DBManager):
        super().__init__()
        self.db_manager = db_manager
        self.layout = QVBoxLayout()

        self.form_layout = QFormLayout()
        self.llm_name_field = QLineEdit()
        self.score_field = QLineEdit()
        self.output_field = QLineEdit()

        self.form_layout.addRow("LLM Name:", self.llm_name_field)
        self.form_layout.addRow("Score:", self.score_field)
        self.form_layout.addRow("Output:", self.output_field)

        self.submit_button = QPushButton("Submit")
        self.layout.addLayout(self.form_layout)
        self.layout.addWidget(self.submit_button)

        self.setLayout(self.layout)

class LeaderboardTab(QWidget):
    def __init__(self, db_manager: DBManager):
        super().__init__()
        self.db_manager = db_manager
        self.layout = QVBoxLayout()

        self.table = QTableWidget()
        self.table.setColumnCount(3)
        self.table.setHorizontalHeaderLabels(["LLM Name", "Total Score", "Attempts"])

        self.layout.addWidget(self.table)
        self.setLayout(self.layout)

class VisualizationTab(QWidget):
    def __init__(self, db_manager: DBManager):
        super().__init__()
        self.db_manager = db_manager
        self.layout = QVBoxLayout()

        self.chart = QChart()
        self.chart_view = QChartView(self.chart)
        self.layout.addWidget(self.chart_view)

        self.setLayout(self.layout)

        self.update_chart()

    def update_chart(self):
        series = QBarSeries()
        set0 = QBarSet("LLM 1")
        set1 = QBarSet("LLM 2")

        set0.append([1, 2, 3, 4])
        set1.append([2, 3, 4, 5])

        series.append(set0)
        series.append(set1)

        self.chart.addSeries(series)
        self.chart.setTitle("LLM Scores Comparison")
        self.chart.setAnimationOptions(QChart.SeriesAnimations)

        categories = ["Prompt 1", "Prompt 2", "Prompt 3", "Prompt 4"]
        axisX = QBarCategoryAxis()
        axisX.append(categories)
        self.chart.addAxis(axisX, Qt.AlignBottom)
        series.attachAxis(axisX)

        axisY = QValueAxis()
        self.chart.addAxis(axisY, Qt.AlignLeft)
        series.attachAxis(axisY)
