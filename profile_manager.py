import tkinter as tk
from tkinter import messagebox, ttk
from database import Database

class ProfileManager(tk.Frame):
    def __init__(self, parent):
        super().__init__(parent)
        self.db = Database()
        self.create_widgets()

    def create_widgets(self):
        self.tree = ttk.Treeview(self, columns=('ID', 'Name'), show='headings')
        self.tree.heading('ID', text='ID')
        self.tree.heading('Name', text='Name')
        self.tree.pack(fill='both', expand=True)

        self.load_profiles()

        self.add_button = tk.Button(self, text='Add', command=self.add_profile)
        self.add_button.pack(side='left')

        self.edit_button = tk.Button(self, text='Edit', command=self.edit_profile)
        self.edit_button.pack(side='left')

        self.delete_button = tk.Button(self, text='Delete', command=self.delete_profile)
        self.delete_button.pack(side='left')

    def load_profiles(self):
        for row in self.tree.get_children():
            self.tree.delete(row)
        profiles = self.db.execute_query('SELECT * FROM profiles')
        for profile in profiles:
            self.tree.insert('', 'end', values=profile)

    def add_profile(self):
        self.show_profile_dialog('Add')

    def edit_profile(self):
        selected_item = self.tree.selection()
        if not selected_item:
            messagebox.showwarning("Warning", "Please select a profile to edit")
            return
        self.show_profile_dialog('Edit', self.tree.item(selected_item)['values'])

    def delete_profile(self):
        selected_item = self.tree.selection()
        if not selected_item:
            messagebox.showwarning("Warning", "Please select a profile to delete")
            return
        profile_id = self.tree.item(selected_item)['values'][0]
        self.db.execute_non_query('DELETE FROM profiles WHERE id = ?', (profile_id,))
        self.load_profiles()

    def show_profile_dialog(self, mode, profile=None):
        dialog = tk.Toplevel(self)
        dialog.title(f'{mode} Profile')

        name_label = tk.Label(dialog, text='Name')
        name_label.pack()

        name_entry = tk.Entry(dialog)
        name_entry.pack()

        if profile:
            name_entry.insert(0, profile[1])

        def save_profile():
            name = name_entry.get()
            if mode == 'Add':
                self.db.execute_non_query('INSERT INTO profiles (name) VALUES (?)', (name,))
            elif mode == 'Edit':
                self.db.execute_non_query('UPDATE profiles SET name = ? WHERE id = ?', (name, profile[0]))
            self.load_profiles()
            dialog.destroy()

        save_button = tk.Button(dialog, text='Save', command=save_profile)
        save_button.pack()
