import tkinter as tk
from tkinter import ttk
from model_manager import ModelManager
from profile_manager import ProfileManager
from prompt_manager import PromptManager
from leaderboard import Leaderboard

class App(tk.Tk):
    def __init__(self):
        super().__init__()
        self.title("LLM Benchmarking App")
        self.geometry("1280x1080")

        self.notebook = ttk.Notebook(self)
        self.notebook.pack(fill='both', expand=True)

        self.model_manager = ModelManager(self.notebook)
        self.profile_manager = ProfileManager(self.notebook)
        self.prompt_manager = PromptManager(self.notebook)
        self.leaderboard = Leaderboard(self.notebook)

        self.notebook.add(self.model_manager, text='Model Manager')
        self.notebook.add(self.profile_manager, text='Profile Manager')
        self.notebook.add(self.prompt_manager, text='Prompt Manager')
        self.notebook.add(self.leaderboard, text='Leaderboard')

if __name__ == "__main__":
    app = App()
    app.mainloop()
