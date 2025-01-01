.PHONY: updateaider pdfprompts pdfprofiles pdfbots test run
updateaider:
	python -m pip install --upgrade git+https://github.com/paul-gauthier/aider.git
pdfprompts:
	pandoc --pdf-engine=xelatex -V "mainfont:Iosevka NerdFont" -V geometry:margin=0.25in -V "monofont:Iosevka Nerd Font" --highlight-style=breezedark .\data\prompts.md -o .\data\prompts.pdf
pdfprofiles:
	pandoc --pdf-engine=xelatex -V "mainfont:Iosevka NerdFont" -V geometry:margin=0.25in -V "monofont:Iosevka Nerd Font" --highlight-style=breezedark .\data\profiles.md -o .\data\profiles.pdf
pdfbots:
	pandoc --pdf-engine=xelatex -V "mainfont:Iosevka NerdFont" -V geometry:margin=0.25in -V "monofont:Iosevka Nerd Font" --highlight-style=breezedark .\data\models.md -o .\data\bots.pdf
test:
	-go test ./... -v -race -cover > test_output.txt
	bat test_output.txt
run:
	go run . load
	go run .
