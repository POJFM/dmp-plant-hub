1. Add `XeLaTeX` config in vscode
```
    "latex-workshop.latex.recipe.default": "latexmk (xelatex)",
    "latex-workshop.latex.tools": [
        {
            "name": "xelatexmk",
            "command": "latexmk",
            "args": [
                "-xelatex",
                "-outdir=out",
                "final-report.tex"
            ]
        },
        {
            "name": "latexmk",
            "command": "latexmk",
            "args": [
                "-xelatex",
                "-synctex=1",
                "-interaction=nonstopmode",
                "-file-line-error",
                "%DOC%"
            ]
        }
    ],
    "latex-workshop.latex.recipes": [
        {
            "name": "latexmk (xelatex)",
            "tools": [
                "xelatexmk"
            ]
        }
    ],
```

2. Download `Calibri.ttf` and put it in `~/.fonts`
3. Compile with `latexmk -xelatex -outdir=final-report ./final-report/final-report.tex`