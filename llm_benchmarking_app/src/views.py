from PyQt6.QtWidgets import QMainWindow, QTabWidget, QVBoxLayout, QWidget, QListWidget, QPushButton, QHBoxLayout, QFormLayout, QLineEdit, QTableWidget, QValueAxis, QTableWidgetItem
from PyQt6.QtCharts import QChartView, QBarSeries, QBarSet, QBarCategoryAxis, QChart
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

        # Connect the load button to a method that loads prompts
        self.load_button.clicked.connect(self.load_prompts)

    def load_prompts(self):
        prompts = self.db_manager.fetch_all_prompts()
        self.prompt_list.clear()
        for prompt in prompts:
            self.prompt_list.addItem(f"{prompt[2]} - {prompt[3]}")

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

        # Connect the submit button to a method that saves the score
        self.submit_button.clicked.connect(self.submit_score)

    def submit_score(self):
        llm_name = self.llm_name_field.text()
        score = float(self.score_field.text())
        output = self.output_field.text()

        # For simplicity, assume llm_id and prompt_id are known
        llm_id = 1  # Replace with actual LLM ID
        prompt_id = 1  # Replace with actual Prompt ID

        self.db_manager.insert_attempt(llm_id, prompt_id, score, output)

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

        # Load leaderboard data
        self.load_leaderboard_data()

    def load_leaderboard_data(self):
        attempts = self.db_manager.fetch_all_attempts()
        llm_scores = {}

        for attempt in attempts:
            llm_id = attempt[1]
            score = attempt[3]
            if llm_id in llm_scores:
                llm_scores[llm_id]['total_score'] += score
                llm_scores[llm_id]['attempts'] += 1
            else:
                llm_scores[llm_id] = {'total_score': score, 'attempts': 1}

        self.table.setRowCount(len(llm_scores))
        row = 0
        for llm_id, data in llm_scores.items():
            llm_name = self.db_manager.fetch_llm_by_id(llm_id)[1]  # Assuming a method to fetch LLM by ID
            self.table.setItem(row, 0, QTableWidgetItem(llm_name))
            self.table.setItem(row, 1, QTableWidgetItem(str(data['total_score'])))
            self.table.setItem(row, 2, QTableWidgetItem(str(data['attempts'])))
            row += 1

class VisualizationTab(QWidget):
    def __init__(self, db_manager: DBManager):
        super().__init__()
        self.db_manager = db_manager
        self.layout = QVBoxLayout()

        self.chart = QChart()
        self.chart_view = QChartView(self.chart)
        self.layout.addWidget(self.chart_view)

        self.setLayout(self.layout)

        # Load and update the chart with data
        self.update_chart()

    def update_chart(self):
        series = QBarSeries()
        llm_scores = self.get_llm_scores()

        for llm_name, scores in llm_scores.items():
            set = QBarSet(llm_name)
            set.append(scores)
            series.append(set)

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

    def get_llm_scores(self):
        attempts = self.db_manager.fetch_all_attempts()
        llm_scores = {}

        for attempt in attempts:
            llm_id = attempt[1]
            score = attempt[3]
            llm_name = self.db_manager.fetch_llm_by_id(llm_id)[1]  # Assuming a method to fetch LLM by ID
            if llm_name in llm_scores:
                llm_scores[llm_name].append(score)
            else:
                llm_scores[llm_name] = [score]

        return llm_scores
