import tkinter as tk
from tkinter import ttk, scrolledtext
from transformers import T5ForConditionalGeneration, T5Tokenizer
import torch

class TranslatorApp:
    def __init__(self, root):
        self.root = root
        self.root.title("English to Vietnamese Translator (MadLad)")
        
        self.model = None
        self.tokenizer = None
        self.device = "cuda" if torch.cuda.is_available() else "cpu"
        
        self.setup_ui()
        self.load_model()
        
    def setup_ui(self):
        # Status text at top
        status_frame = ttk.Frame(self.root)
        status_frame.pack(fill=tk.X, padx=5, pady=5)
        
        self.status_text = tk.Text(status_frame, height=2, wrap=tk.WORD)
        self.status_text.pack(fill=tk.X, expand=True)
        self.status_text.configure(state='disabled')
        
        # Input frame
        input_frame = ttk.LabelFrame(self.root, text="English Text", padding="5")
        input_frame.pack(fill=tk.BOTH, expand=True, padx=5, pady=5)
        
        self.input_text = scrolledtext.ScrolledText(input_frame, wrap=tk.WORD, height=10)
        self.input_text.pack(fill=tk.BOTH, expand=True)
        
        # Buttons frame
        button_frame = ttk.Frame(self.root, padding="5")
        button_frame.pack(fill=tk.X, padx=5)
        
        self.translate_button = ttk.Button(button_frame, text="Translate", command=self.translate)
        self.translate_button.pack(side=tk.LEFT, padx=5)
        self.translate_button.state(['disabled'])  # Initially disabled until model is loaded
        
        self.clear_button = ttk.Button(button_frame, text="Clear", command=self.clear_all)
        self.clear_button.pack(side=tk.LEFT)
        
        # Output frame
        output_frame = ttk.LabelFrame(self.root, text="Vietnamese Translation", padding="5")
        output_frame.pack(fill=tk.BOTH, expand=True, padx=5, pady=5)
        
        self.output_text = scrolledtext.ScrolledText(output_frame, wrap=tk.WORD, height=10)
        self.output_text.pack(fill=tk.BOTH, expand=True)
        
        # Copy button
        self.copy_button = ttk.Button(output_frame, text="Copy to Clipboard", command=self.copy_to_clipboard)
        self.copy_button.pack(pady=5)
        
    def set_status(self, message):
        """Update status text with new message"""
        self.status_text.configure(state='normal')
        self.status_text.delete('1.0', tk.END)
        self.status_text.insert('1.0', message)
        self.status_text.configure(state='disabled')
        self.root.update()
    
    def load_model(self):
        try:
            self.set_status("Downloading and loading model and tokenizer...\nThis may take a while on first run.")
            
            # Load tokenizer and model from HuggingFace
            model_name = 'google/madlad400-3b-mt'
            self.tokenizer = T5Tokenizer.from_pretrained(model_name)
            self.model = T5ForConditionalGeneration.from_pretrained(
                model_name,
                device_map="auto" if torch.cuda.is_available() else None
            )
            
            # Move model to appropriate device if not using device_map="auto"
            if not torch.cuda.is_available():
                self.model = self.model.to(self.device)
            
            self.translate_button.state(['!disabled'])
            self.set_status(f"Model loaded successfully!\nDevice: {self.device}")
            
        except Exception as e:
            error_msg = f"Error loading model: {str(e)}"
            self.set_status(error_msg)
            self.translate_button.state(['disabled'])
    
    def translate(self):
        if not self.model or not self.tokenizer:
            self.set_status("Model not loaded properly")
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
            
            # Tokenize and generate translation
            input_ids = self.tokenizer(input_text, return_tensors="pt").input_ids
            if not torch.cuda.is_available():
                input_ids = input_ids.to(self.device)
                
            outputs = self.model.generate(
                input_ids=input_ids,
                max_length=512,
                num_beams=4,
                length_penalty=0.6,
                early_stopping=True
            )
            
            translation = self.tokenizer.decode(outputs[0], skip_special_tokens=True)
            
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
