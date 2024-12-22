import tkinter as tk
from tkinter import messagebox, ttk
from database import Database

class ModelManager(tk.Frame):
    def __init__(self, parent):
        super().__init__(parent)
        self.db = Database()
        self.create_widgets()

    def create_widgets(self):
        self.tree = ttk.Treeview(self, columns=('ID', 'Name', 'Path', 'GPU Layers', 'Ctx Size', 'Batch Size', 'Threads', 'Keep', 'Predict', 'Flash Attn', 'Mlock', 'Cache Type K', 'Cache Type V'), show='headings')
        self.tree.heading('ID', text='ID')
        self.tree.heading('Name', text='Name')
        self.tree.heading('Path', text='Path')
        self.tree.heading('GPU Layers', text='GPU Layers')
        self.tree.heading('Ctx Size', text='Ctx Size')
        self.tree.heading('Batch Size', text='Batch Size')
        self.tree.heading('Threads', text='Threads')
        self.tree.heading('Keep', text='Keep')
        self.tree.heading('Predict', text='Predict')
        self.tree.heading('Flash Attn', text='Flash Attn')
        self.tree.heading('Mlock', text='Mlock')
        self.tree.heading('Cache Type K', text='Cache Type K')
        self.tree.heading('Cache Type V', text='Cache Type V')
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

        path_label = tk.Label(dialog, text='Path')
        path_label.pack()

        path_entry = tk.Entry(dialog)
        path_entry.pack()

        gpu_layers_label = tk.Label(dialog, text='GPU Layers')
        gpu_layers_label.pack()

        gpu_layers_entry = tk.Entry(dialog)
        gpu_layers_entry.pack()

        ctx_size_label = tk.Label(dialog, text='Ctx Size')
        ctx_size_label.pack()

        ctx_size_entry = tk.Entry(dialog)
        ctx_size_entry.pack()

        batch_size_label = tk.Label(dialog, text='Batch Size')
        batch_size_label.pack()

        batch_size_entry = tk.Entry(dialog)
        batch_size_entry.pack()

        threads_label = tk.Label(dialog, text='Threads')
        threads_label.pack()

        threads_entry = tk.Entry(dialog)
        threads_entry.pack()

        keep_label = tk.Label(dialog, text='Keep')
        keep_label.pack()

        keep_entry = tk.Entry(dialog)
        keep_entry.pack()

        predict_label = tk.Label(dialog, text='Predict')
        predict_label.pack()

        predict_entry = tk.Entry(dialog)
        predict_entry.pack()

        flash_attn_label = tk.Label(dialog, text='Flash Attn')
        flash_attn_label.pack()

        flash_attn_entry = tk.Entry(dialog)
        flash_attn_entry.pack()

        mlock_label = tk.Label(dialog, text='Mlock')
        mlock_label.pack()

        mlock_entry = tk.Entry(dialog)
        mlock_entry.pack()

        cache_type_k_label = tk.Label(dialog, text='Cache Type K')
        cache_type_k_label.pack()

        cache_type_k_entry = tk.Entry(dialog)
        cache_type_k_entry.pack()

        cache_type_v_label = tk.Label(dialog, text='Cache Type V')
        cache_type_v_label.pack()

        cache_type_v_entry = tk.Entry(dialog)
        cache_type_v_entry.pack()

        if model:
            name_entry.insert(0, model[1])
            path_entry.insert(0, model[2])
            gpu_layers_entry.insert(0, model[3])
            ctx_size_entry.insert(0, model[4])
            batch_size_entry.insert(0, model[5])
            threads_entry.insert(0, model[6])
            keep_entry.insert(0, model[7])
            predict_entry.insert(0, model[8])
            flash_attn_entry.insert(0, model[9])
            mlock_entry.insert(0, model[10])
            cache_type_k_entry.insert(0, model[11])
            cache_type_v_entry.insert(0, model[12])

        def save_model():
            name = name_entry.get()
            path = path_entry.get()
            gpu_layers = gpu_layers_entry.get()
            ctx_size = ctx_size_entry.get()
            batch_size = batch_size_entry.get()
            threads = threads_entry.get()
            keep = keep_entry.get()
            predict = predict_entry.get()
            flash_attn = flash_attn_entry.get()
            mlock = mlock_entry.get()
            cache_type_k = cache_type_k_entry.get()
            cache_type_v = cache_type_v_entry.get()

            if mode == 'Add':
                self.db.execute_non_query('''
                    INSERT INTO models (name, path, gpu_layers, ctx_size, batch_size, threads, keep, predict, flash_attn, mlock, cache_type_k, cache_type_v)
                    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
                ''', (name, path, gpu_layers, ctx_size, batch_size, threads, keep, predict, flash_attn, mlock, cache_type_k, cache_type_v))
            elif mode == 'Edit':
                self.db.execute_non_query('''
                    UPDATE models
                    SET name = ?, path = ?, gpu_layers = ?, ctx_size = ?, batch_size = ?, threads = ?, keep = ?, predict = ?, flash_attn = ?, mlock = ?, cache_type_k = ?, cache_type_v = ?
                    WHERE id = ?
                ''', (name, path, gpu_layers, ctx_size, batch_size, threads, keep, predict, flash_attn, mlock, cache_type_k, cache_type_v, model[0]))
            self.load_models()
            dialog.destroy()

        save_button = tk.Button(dialog, text='Save', command=save_model)
        save_button.pack()
