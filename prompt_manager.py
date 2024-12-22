import tkinter as tk
from tkinter import messagebox
from database import Database

class PromptManager(tk.Frame):
    def __init__(self, parent):
        super().__init__(parent)
        self.db = Database()
        self.create_widgets()

    def create_widgets(self):
        self.tree = ttk.Treeview(self, columns=('ID', 'Content', 'Solution', 'Profile', 'Category'), show='headings')
        self.tree.heading('ID', text='ID')
        self.tree.heading('Content', text='Content')
        self.tree.heading('Solution', text='Solution')
        self.tree.heading('Profile', text='Profile')
        self.tree.heading('Category', text='Category')
        self.tree.pack(fill='both', expand=True)

        self.load_prompts()

        self.add_button = tk.Button(self, text='Add', command=self.add_prompt)
        self.add_button.pack(side='left')

        self.edit_button = tk.Button(self, text='Edit', command=self.edit_prompt)
        self.edit_button.pack(side='left')

        self.delete_button = tk.Button(self, text='Delete', command=self.delete_prompt)
        self.delete_button.pack(side='left')

    def load_prompts(self):
        for row in self.tree.get_children():
            self.tree.delete(row)
        prompts = self.db.execute_query('SELECT * FROM prompts')
        for prompt in prompts:
            self.tree.insert('', 'end', values=prompt)

    def add_prompt(self):
        self.show_prompt_dialog('Add')

    def edit_prompt(self):
        selected_item = self.tree.selection()
        if not selected_item:
            messagebox.showwarning("Warning", "Please select a prompt to edit")
            return
        self.show_prompt_dialog('Edit', self.tree.item(selected_item)['values'])

    def delete_prompt(self):
        selected_item = self.tree.selection()
        if not selected_item:
            messagebox.showwarning("Warning", "Please select a prompt to delete")
            return
        prompt_id = self.tree.item(selected_item)['values'][0]
        self.db.execute_non_query('DELETE FROM prompts WHERE id = ?', (prompt_id,))
        self.load_prompts()

    def show_prompt_dialog(self, mode, prompt=None):
        dialog = tk.Toplevel(self)
        dialog.title(f'{mode} Prompt')

        content_label = tk.Label(dialog, text='Content')
        content_label.pack()

        content_entry = tk.Entry(dialog)
        content_entry.pack()

        solution_label = tk.Label(dialog, text='Solution')
        solution_label.pack()

        solution_entry = tk.Entry(dialog)
        solution_entry.pack()

        profile_label = tk.Label(dialog, text='Profile')
        profile_label.pack()

        profile_entry = tk.Entry(dialog)
        profile_entry.pack()

        category_label = tk.Label(dialog, text='Category')
        category_label.pack()

        category_entry = tk.Entry(dialog)
        category_entry.pack()

        if prompt:
            content_entry.insert(0, prompt[1])
            solution_entry.insert(0, prompt[2])
            profile_entry.insert(0, prompt[3])
            category_entry.insert(0, prompt[4])

        def save_prompt():
            content = content_entry.get()
            solution = solution_entry.get()
            profile = profile_entry.get()
            category = category_entry.get()
            if mode == 'Add':
                self.db.execute_non_query('INSERT INTO prompts (content, solution, profile, category) VALUES (?, ?, ?, ?)',
                                          (content, solution, profile, category))
            elif mode == 'Edit':
                self.db.execute_non_query('UPDATE prompts SET content = ?, solution = ?, profile = ?, category = ? WHERE id = ?',
                                          (content, solution, profile, category, prompt[0]))
            self.load_prompts()
            dialog.destroy()

        save_button = tk.Button(dialog, text='Save', command=save_prompt)
        save_button.pack()
