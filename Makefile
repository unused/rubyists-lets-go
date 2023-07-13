PANDOC_OPTS=-t beamer -H templates/slides-header.tex

.PHONY: presentation
presentation: slides.pdf

%.pdf: %.md
	@pandoc $(PANDOC_OPTS) -o $@ $<

.PHONY: tex
tex:
	@pandoc $(PANDOC_OPTS) -o slides.tex slides.md

.PHONY: clean
clean:
	@rm -f slides.pdf

.PHONY: watch
watch:
	@echo slides.md | entr -c 'make clean && make'
