from PyQt6.QtWidgets import QMainWindow, QTabWidget, QVBoxLayout, QWidget, QListWidget, QPushButton, QHBoxLayout, QFormLayout, QLineEdit, QTableWidget, QTableWidgetItem, QInputDialog, QDialog, QDialogButtonBox, QListWidgetItem
from PyQt6.QtCharts import QChartView, QBarSeries, QBarSet, QBarCategoryAxis, QChart, QValueAxis
from PyQt6.QtCore import Qt
from db_manager import DBManager
import json
from PyQt6.QtWidgets import QMessageBox

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

        self.prompt_management_tab.prompt_list.itemSelectionChanged.connect(self.on_prompt_selection_changed)

    def on_prompt_selection_changed(self):
        selected_items = self.prompt_management_tab.prompt_list.selectedItems()
        if selected_items:
            prompt_id = selected_items[0].data(Qt.UserRole)
            self.score_entry_tab.set_selected_prompt(prompt_id)

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
        self.delete_button = QPushButton("Delete Prompt")
        self.add_button = QPushButton("Add Prompt")
        self.save_button = QPushButton("Save Prompts")

        self.button_layout.addWidget(self.load_button)
        self.button_layout.addWidget(self.edit_button)
        self.button_layout.addWidget(self.delete_button)
        self.button_layout.addWidget(self.add_button)
        self.button_layout.addWidget(self.save_button)

        self.layout.addLayout(self.button_layout)
        self.setLayout(self.layout)

        # Connect the load button to a method that loads prompts
        self.load_button.clicked.connect(self.load_prompts_from_json)
        self.edit_button.clicked.connect(self.edit_prompt)
        self.delete_button.clicked.connect(self.delete_prompt)
        self.add_button.clicked.connect(self.add_prompt)
        self.save_button.clicked.connect(self.save_prompts)

    def load_prompts_from_json(self):
        try:
            with open('llm_benchmarking_app/data/profiles/profile.json', 'r') as profile_file:
                profiles = json.load(profile_file)

            with open('llm_benchmarking_app/data/prompts/prompt_suite.json', 'r') as prompt_file:
                prompts = json.load(prompt_file)

            # Clear existing prompts in the database
            self.db_manager.conn.execute('DELETE FROM prompts')
            self.db_manager.conn.execute('DELETE FROM prompt_suites')

            # Insert profiles into the database
            for profile in profiles['profiles']:
                self.db_manager.insert_prompt_suite(profile['name'], profile['system_prompt'])

            # Insert prompts into the database
            for prompt in prompts['prompts']:
                self.db_manager.insert_prompt(prompt['suite_id'], prompt['content'], prompt['category'])

            # Refresh the prompt list
            self.load_prompts()

            QMessageBox.information(self, "Success", "Prompts loaded successfully!")
        except Exception as e:
            QMessageBox.critical(self, "Error", f"Failed to load prompts: {e}")

    def load_prompts(self):
        prompts = self.db_manager.fetch_all_prompts()
        self.prompt_list.clear()
        for prompt in prompts:
            item = QListWidgetItem(f"{prompt[2]} - {prompt[3]}")
            item.setData(Qt.UserRole, prompt[0])
            self.prompt_list.addItem(item)

    def edit_prompt(self):
        selected_item = self.prompt_list.currentItem()
        if selected_item:
            prompt_id = selected_item.data(Qt.UserRole)
            prompt = self.db_manager.fetch_prompt_by_id(prompt_id)
            if prompt:
                dialog = QDialog(self)
                dialog.setWindowTitle("Edit Prompt")
                layout = QVBoxLayout()

                form_layout = QFormLayout()
                content_field = QLineEdit(prompt[2])
                category_field = QLineEdit(prompt[3])

                form_layout.addRow("Content:", content_field)
                form_layout.addRow("Category:", category_field)

                layout.addLayout(form_layout)

                buttons = QDialogButtonBox(QDialogButtonBox.Ok | QDialogButtonBox.Cancel, Qt.Horizontal, dialog)
                buttons.accepted.connect(dialog.accept)
                buttons.rejected.connect(dialog.reject)
                layout.addWidget(buttons)

                dialog.setLayout(layout)

                if dialog.exec() == QDialog.Accepted:
                    new_content = content_field.text()
                    new_category = category_field.text()
                    self.db_manager.update_prompt(prompt_id, new_content, new_category)
                    self.load_prompts()

    def delete_prompt(self):
        selected_item = self.prompt_list.currentItem()
        if selected_item:
            prompt_id = selected_item.data(Qt.UserRole)
            self.db_manager.delete_prompt(prompt_id)
            self.load_prompts()

    def add_prompt(self):
        dialog = QDialog(self)
        dialog.setWindowTitle("Add Prompt")
        layout = QVBoxLayout()

        form_layout = QFormLayout()
        suite_id_field = QLineEdit()
        content_field = QLineEdit()
        category_field = QLineEdit()

        form_layout.addRow("Suite ID:", suite_id_field)
        form_layout.addRow("Content:", content_field)
        form_layout.addRow("Category:", category_field)

        layout.addLayout(form_layout)

        buttons = QDialogButtonBox(QDialogButtonBox.Ok | QDialogButtonBox.Cancel, Qt.Horizontal, dialog)
        buttons.accepted.connect(dialog.accept)
        buttons.rejected.connect(dialog.reject)
        layout.addWidget(buttons)

        dialog.setLayout(layout)

        if dialog.exec() == QDialog.Accepted:
            suite_id = int(suite_id_field.text())
            content = content_field.text()
            category = category_field.text()
            self.db_manager.insert_prompt(suite_id, content, category)
            self.load_prompts()

    def save_prompts(self):
        # Implement save prompts functionality
        pass

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

    def set_selected_prompt(self, prompt_id):
        prompt = self.db_manager.fetch_prompt_by_id(prompt_id)
        if prompt:
            self.llm_name_field.setText(prompt[2])
            self.score_field.setText(str(prompt[3]))
            self.output_field.setText(prompt[4])

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
        self.chart.setAnimationOptions(QChart.AnimationOption.SeriesAnimations)

        categories = ["Prompt 1", "Prompt 2", "Prompt 3", "Prompt 4"]
        axisX = QBarCategoryAxis()
        axisX.append(categories)
        self.chart.addAxis(axisX, Qt.AlignmentFlag.AlignBottom)
        series.attachAxis(axisX)

        axisY = QValueAxis()
        self.chart.addAxis(axisY, Qt.AlignmentFlag.AlignLeft)
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
