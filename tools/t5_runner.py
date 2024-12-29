import tkinter as tk
from tkinter import ttk, scrolledtext, filedialog
from transformers import T5ForConditionalGeneration, T5Tokenizer
import torch
import os


class TranslatorApp:
    def __init__(self, root):
        self.root = root
        self.root.title("English to Vietnamese Translator (T5)")

        self.model = None
        self.tokenizer = None
        self.device = "cuda" if torch.cuda.is_available() else "cpu"

        self.setup_ui()

    def setup_ui(self):
        # Model selection frame
        model_frame = ttk.LabelFrame(
            self.root, text="Model Selection", padding="5")
        model_frame.pack(fill=tk.X, padx=5, pady=5)

        # Model path selection
        path_frame = ttk.Frame(model_frame)
        path_frame.pack(fill=tk.X, expand=True, pady=(0, 5))

        self.model_path = tk.StringVar()
        self.model_entry = ttk.Entry(path_frame, textvariable=self.model_path)
        self.model_entry.pack(side=tk.LEFT, fill=tk.X,
                              expand=True, padx=(0, 5))

        self.browse_button = ttk.Button(
            path_frame, text="Browse", command=self.browse_model)
        self.browse_button.pack(side=tk.LEFT)

        self.load_button = ttk.Button(
            path_frame, text="Load Model", command=self.load_model)
        self.load_button.pack(side=tk.LEFT, padx=(5, 0))

        # Input frame
        input_frame = ttk.LabelFrame(
            self.root, text="English Text", padding="5")
        input_frame.pack(fill=tk.BOTH, expand=True, padx=5, pady=5)

        self.input_text = scrolledtext.ScrolledText(
            input_frame, wrap=tk.WORD, height=10)
        self.input_text.pack(fill=tk.BOTH, expand=True)

        # Buttons frame
        button_frame = ttk.Frame(self.root, padding="5")
        button_frame.pack(fill=tk.X, padx=5)

        self.translate_button = ttk.Button(
            button_frame, text="Translate", command=self.translate)
        self.translate_button.pack(side=tk.LEFT, padx=5)
        # Initially disabled until model is loaded
        self.translate_button.state(['disabled'])

        self.clear_button = ttk.Button(
            button_frame, text="Clear", command=self.clear_all)
        self.clear_button.pack(side=tk.LEFT)

        # Output frame
        output_frame = ttk.LabelFrame(
            self.root, text="Vietnamese Translation", padding="5")
        output_frame.pack(fill=tk.BOTH, expand=True, padx=5, pady=5)

        self.output_text = scrolledtext.ScrolledText(
            output_frame, wrap=tk.WORD, height=10)
        self.output_text.pack(fill=tk.BOTH, expand=True)

        # Copy button
        self.copy_button = ttk.Button(
            output_frame, text="Copy to Clipboard", command=self.copy_to_clipboard)
        self.copy_button.pack(pady=5)

        # Status text
        status_frame = ttk.Frame(self.root)
        status_frame.pack(fill=tk.X, padx=5, pady=5)

        self.status_text = tk.Text(status_frame, height=2, wrap=tk.WORD)
        self.status_text.pack(fill=tk.X, expand=True)
        # Make it read-only by default
        self.status_text.configure(state='disabled')

    def set_status(self, message):
        """Update status text with new message"""
        self.status_text.configure(state='normal')
        self.status_text.delete('1.0', tk.END)
        self.status_text.insert('1.0', message)
        self.status_text.configure(state='disabled')
        self.root.update()

    def browse_model(self):
        folder = filedialog.askdirectory(
            title="Select Model Directory"
        )
        if folder:
            self.model_path.set(folder)

    def load_model(self):
        model_path = self.model_path.get()
        if not model_path or not os.path.exists(model_path):
            self.set_status("Use default: google/madlad400-3b-mt")
            model_path = 'google/madlad400-3b-mt'

        try:
            self.set_status("Loading model and tokenizer...")

            # Load tokenizer and model
            self.tokenizer = T5Tokenizer.from_pretrained(model_path)
            self.model = T5ForConditionalGeneration.from_pretrained(
                model_path,
                device_map="auto" if torch.cuda.is_available() else None
            )

            # Move model to appropriate device if not using device_map="auto"
            if not torch.cuda.is_available():
                self.model = self.model.to(self.device)

            self.translate_button.state(['!disabled'])
            self.set_status(f"Model loaded successfully!\nPath: {
                            model_path}\nDevice: {self.device}")

        except Exception as e:
            error_msg = f"Error loading model: {str(e)}\nPath: {model_path}"
            self.set_status(error_msg)
            self.translate_button.state(['disabled'])

    def translate(self):
        if not self.model or not self.tokenizer:
            self.set_status("Please load a model first")
            return

        self.set_status("Translating...")
        self.translate_button.state(['disabled'])
        self.root.update()

        try:
            input_text = self.input_text.get("1.0", tk.END).strip()
            if not input_text:
                self.set_status("Please enter some text to translate")
                return

            # Add Vietnamese target language tag
            input_text = "<2vi> " + input_text

            # Process text in chunks if it's too long
            max_length = 512

            # Tokenize and generate translation
            input_ids = self.tokenizer(
                input_text, return_tensors="pt").input_ids
            if not torch.cuda.is_available():
                input_ids = input_ids.to(self.device)

            outputs = self.model.generate(
                input_ids=input_ids,
                max_length=max_length,
                num_beams=4,
                length_penalty=0.6,
                early_stopping=True
            )

            translation = self.tokenizer.decode(
                outputs[0], skip_special_tokens=True)

            self.output_text.delete("1.0", tk.END)
            self.output_text.insert("1.0", translation)
            self.set_status("Translation completed!")

        except Exception as e:
            self.set_status(f"Error during translation: {str(e)}")
        finally:
            self.translate_button.state(['!disabled'])

    def clear_all(self):
        self.input_text.delete("1.0", tk.END)
        self.output_text.delete("1.0", tk.END)
        self.set_status("")

    def copy_to_clipboard(self):
        translation = self.output_text.get("1.0", tk.END).strip()
        self.root.clipboard_clear()
        self.root.clipboard_append(translation)
        self.set_status("Translation copied to clipboard!")


if __name__ == "__main__":
    root = tk.Tk()
    app = TranslatorApp(root)
    root.mainloop()
