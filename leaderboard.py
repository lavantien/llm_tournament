import tkinter as tk
from tkinter import ttk
from database import Database

class Leaderboard(tk.Frame):
    def __init__(self, parent):
        super().__init__(parent)
        self.db = Database()
        self.create_widgets()

    def create_widgets(self):
        self.tree = ttk.Treeview(self, columns=('Model', 'Category1', 'Category2', 'Category3', 'Total'), show='headings')
        self.tree.heading('Model', text='Model')
        self.tree.heading('Category1', text='Category1')
        self.tree.heading('Category2', text='Category2')
        self.tree.heading('Category3', text='Category3')
        self.tree.heading('Total', text='Total')
        self.tree.pack(fill='both', expand=True)

        self.load_leaderboard()

        self.tree.bind('<ButtonRelease-1>', self.show_detail)

    def load_leaderboard(self):
        for row in self.tree.get_children():
            self.tree.delete(row)
        leaderboard = self.db.execute_query('''
            SELECT m.name,
                   SUM(CASE WHEN p.category = 'Category1' THEN s.score ELSE 0 END) AS Category1,
                   SUM(CASE WHEN p.category = 'Category2' THEN s.score ELSE 0 END) AS Category2,
                   SUM(CASE WHEN p.category = 'Category3' THEN s.score ELSE 0 END) AS Category3,
                   SUM(s.score) AS Total
            FROM models m
            LEFT JOIN scores s ON m.id = s.model_id
            LEFT JOIN prompts p ON s.prompt_id = p.id
            GROUP BY m.name
        ''')
        for row in leaderboard:
            self.tree.insert('', 'end', values=row)

    def show_detail(self, event):
        selected_item = self.tree.selection()
        if not selected_item:
            return
        model_name = self.tree.item(selected_item)['values'][0]
        self.show_detail_dialog(model_name)

    def show_detail_dialog(self, model_name):
        dialog = tk.Toplevel(self)
        dialog.title(f'Details for {model_name}')

        detail_tree = ttk.Treeview(dialog, columns=('Prompt', 'Score', 'Output'), show='headings')
        detail_tree.heading('Prompt', text='Prompt')
        detail_tree.heading('Score', text='Score')
        detail_tree.heading('Output', text='Output')
        detail_tree.pack(fill='both', expand=True)

        details = self.db.execute_query('''
            SELECT p.content, s.score, s.output
            FROM scores s
            JOIN prompts p ON s.prompt_id = p.id
            JOIN models m ON s.model_id = m.id
            WHERE m.name = ?
        ''', (model_name,))
        for detail in details:
            detail_tree.insert('', 'end', values=detail)

        def save_detail():
            # Implement saving logic here
            dialog.destroy()

        save_button = tk.Button(dialog, text='Save', command=save_detail)
        save_button.pack()
