import tkinter as tk
from tkinter import messagebox, ttk
from database import Database

class ProfileManager(tk.Frame):
    def __init__(self, parent):
        super().__init__(parent)
        self.db = Database()
        self.create_widgets()

    def create_widgets(self):
        self.tree = ttk.Treeview(self, columns=('ID', 'Name', 'System Prompt', 'Dry Multiplier', 'Dry Base', 'Dry Allowed Length', 'Dry Penalty Last N', 'Repeat Penalty', 'Repeat Last N', 'Top K', 'Top P', 'Min P', 'Top A', 'XTC Threshold', 'XTC Probability', 'Temperature'), show='headings')
        self.tree.heading('ID', text='ID')
        self.tree.heading('Name', text='Name')
        self.tree.heading('System Prompt', text='System Prompt')
        self.tree.heading('Dry Multiplier', text='Dry Multiplier')
        self.tree.heading('Dry Base', text='Dry Base')
        self.tree.heading('Dry Allowed Length', text='Dry Allowed Length')
        self.tree.heading('Dry Penalty Last N', text='Dry Penalty Last N')
        self.tree.heading('Repeat Penalty', text='Repeat Penalty')
        self.tree.heading('Repeat Last N', text='Repeat Last N')
        self.tree.heading('Top K', text='Top K')
        self.tree.heading('Top P', text='Top P')
        self.tree.heading('Min P', text='Min P')
        self.tree.heading('Top A', text='Top A')
        self.tree.heading('XTC Threshold', text='XTC Threshold')
        self.tree.heading('XTC Probability', text='XTC Probability')
        self.tree.heading('Temperature', text='Temperature')
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

        system_prompt_label = tk.Label(dialog, text='System Prompt')
        system_prompt_label.pack()

        system_prompt_entry = tk.Entry(dialog)
        system_prompt_entry.pack()

        dry_multiplier_label = tk.Label(dialog, text='Dry Multiplier')
        dry_multiplier_label.pack()

        dry_multiplier_entry = tk.Entry(dialog)
        dry_multiplier_entry.pack()

        dry_base_label = tk.Label(dialog, text='Dry Base')
        dry_base_label.pack()

        dry_base_entry = tk.Entry(dialog)
        dry_base_entry.pack()

        dry_allowed_length_label = tk.Label(dialog, text='Dry Allowed Length')
        dry_allowed_length_label.pack()

        dry_allowed_length_entry = tk.Entry(dialog)
        dry_allowed_length_entry.pack()

        dry_penalty_last_n_label = tk.Label(dialog, text='Dry Penalty Last N')
        dry_penalty_last_n_label.pack()

        dry_penalty_last_n_entry = tk.Entry(dialog)
        dry_penalty_last_n_entry.pack()

        repeat_penalty_label = tk.Label(dialog, text='Repeat Penalty')
        repeat_penalty_label.pack()

        repeat_penalty_entry = tk.Entry(dialog)
        repeat_penalty_entry.pack()

        repeat_last_n_label = tk.Label(dialog, text='Repeat Last N')
        repeat_last_n_label.pack()

        repeat_last_n_entry = tk.Entry(dialog)
        repeat_last_n_entry.pack()

        top_k_label = tk.Label(dialog, text='Top K')
        top_k_label.pack()

        top_k_entry = tk.Entry(dialog)
        top_k_entry.pack()

        top_p_label = tk.Label(dialog, text='Top P')
        top_p_label.pack()

        top_p_entry = tk.Entry(dialog)
        top_p_entry.pack()

        min_p_label = tk.Label(dialog, text='Min P')
        min_p_label.pack()

        min_p_entry = tk.Entry(dialog)
        min_p_entry.pack()

        top_a_label = tk.Label(dialog, text='Top A')
        top_a_label.pack()

        top_a_entry = tk.Entry(dialog)
        top_a_entry.pack()

        xtc_threshold_label = tk.Label(dialog, text='XTC Threshold')
        xtc_threshold_label.pack()

        xtc_threshold_entry = tk.Entry(dialog)
        xtc_threshold_entry.pack()

        xtc_probability_label = tk.Label(dialog, text='XTC Probability')
        xtc_probability_label.pack()

        xtc_probability_entry = tk.Entry(dialog)
        xtc_probability_entry.pack()

        temperature_label = tk.Label(dialog, text='Temperature')
        temperature_label.pack()

        temperature_entry = tk.Entry(dialog)
        temperature_entry.pack()

        if profile:
            name_entry.insert(0, profile[1])
            system_prompt_entry.insert(0, profile[2])
            dry_multiplier_entry.insert(0, profile[3])
            dry_base_entry.insert(0, profile[4])
            dry_allowed_length_entry.insert(0, profile[5])
            dry_penalty_last_n_entry.insert(0, profile[6])
            repeat_penalty_entry.insert(0, profile[7])
            repeat_last_n_entry.insert(0, profile[8])
            top_k_entry.insert(0, profile[9])
            top_p_entry.insert(0, profile[10])
            min_p_entry.insert(0, profile[11])
            top_a_entry.insert(0, profile[12])
            xtc_threshold_entry.insert(0, profile[13])
            xtc_probability_entry.insert(0, profile[14])
            temperature_entry.insert(0, profile[15])

        def save_profile():
            name = name_entry.get()
            system_prompt = system_prompt_entry.get()
            dry_multiplier = dry_multiplier_entry.get()
            dry_base = dry_base_entry.get()
            dry_allowed_length = dry_allowed_length_entry.get()
            dry_penalty_last_n = dry_penalty_last_n_entry.get()
            repeat_penalty = repeat_penalty_entry.get()
            repeat_last_n = repeat_last_n_entry.get()
            top_k = top_k_entry.get()
            top_p = top_p_entry.get()
            min_p = min_p_entry.get()
            top_a = top_a_entry.get()
            xtc_threshold = xtc_threshold_entry.get()
            xtc_probability = xtc_probability_entry.get()
            temperature = temperature_entry.get()

            if mode == 'Add':
                self.db.execute_non_query('''
                    INSERT INTO profiles (name, system_prompt, dry_multiplier, dry_base, dry_allowed_length,
                                         dry_penalty_last_n, repeat_penalty, repeat_last_n, top_k, top_p, min_p,
                                         top_a, xtc_threshold, xtc_probability, temperature)
                    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
                ''', (name, system_prompt, dry_multiplier, dry_base, dry_allowed_length, dry_penalty_last_n,
                      repeat_penalty, repeat_last_n, top_k, top_p, min_p, top_a, xtc_threshold, xtc_probability,
                      temperature))
            elif mode == 'Edit':
                self.db.execute_non_query('''
                    UPDATE profiles
                    SET name = ?, system_prompt = ?, dry_multiplier = ?, dry_base = ?, dry_allowed_length = ?,
                        dry_penalty_last_n = ?, repeat_penalty = ?, repeat_last_n = ?, top_k = ?, top_p = ?, min_p = ?,
                        top_a = ?, xtc_threshold = ?, xtc_probability = ?, temperature = ?
                    WHERE id = ?
                ''', (name, system_prompt, dry_multiplier, dry_base, dry_allowed_length, dry_penalty_last_n,
                      repeat_penalty, repeat_last_n, top_k, top_p, min_p, top_a, xtc_threshold, xtc_probability,
                      temperature, profile[0]))
            self.load_profiles()
            dialog.destroy()

        save_button = tk.Button(dialog, text='Save', command=save_profile)
        save_button.pack()
