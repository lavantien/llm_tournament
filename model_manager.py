import tkinter as tk
from tkinter import messagebox, ttk
from database import Database

class ModelManager(tk.Frame):
    def __init__(self, parent):
        super().__init__(parent)
        self.db = Database()
        self.create_widgets()

    def create_widgets(self):
        self.tree = ttk.Treeview(self, columns=('ID', 'Name'), show='headings')
        self.tree.heading('ID', text='ID')
        self.tree.heading('Name', text='Name')
        self.tree.pack(fill='both', expand=True)

        self.load_models()

        self.add_button = tk.Button(self, text='Add', command=self.add_model)
        self.add_button.pack(side='left')

        self.edit_button = tk.Button(self, text='Edit', command=self.edit_model)
        self.edit_button.pack(side='left')

        self.delete_button = tk.Button(self, text='Delete', command=self.delete_model)
        self.delete_button.pack(side='left')

    def load_models(self):
        for row in self.tree.get_children():
            self.tree.delete(row)
        models = self.db.execute_query('SELECT * FROM models')
        for model in models:
            self.tree.insert('', 'end', values=model)

    def add_model(self):
        self.show_model_dialog('Add')

    def edit_model(self):
        selected_item = self.tree.selection()
        if not selected_item:
            messagebox.showwarning("Warning", "Please select a model to edit")
            return
        self.show_model_dialog('Edit', self.tree.item(selected_item)['values'])

    def delete_model(self):
        selected_item = self.tree.selection()
        if not selected_item:
            messagebox.showwarning("Warning", "Please select a model to delete")
            return
        model_id = self.tree.item(selected_item)['values'][0]
        self.db.execute_non_query('DELETE FROM models WHERE id = ?', (model_id,))
        self.load_models()

    def show_model_dialog(self, mode, model=None):
        dialog = tk.Toplevel(self)
        dialog.title(f'{mode} Model')

        name_label = tk.Label(dialog, text='Name')
        name_label.pack()

        name_entry = tk.Entry(dialog)
        name_entry.pack()

        if model:
            name_entry.insert(0, model[1])

        def save_model():
            name = name_entry.get()
            if mode == 'Add':
                self.db.execute_non_query('INSERT INTO models (name) VALUES (?)', (name,))
            elif mode == 'Edit':
                self.db.execute_non_query('UPDATE models SET name = ? WHERE id = ?', (name, model[0]))
            self.load_models()
            dialog.destroy()

        save_button = tk.Button(dialog, text='Save', command=save_model)
        save_button.pack()
