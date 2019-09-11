.PHONY: init

serve:
	@docker run --rm -it --name www-gianarb \
		-p 4000:4000 \
		-v www-gianarb-cache:/usr/local/bundle \
		-e TZ='Europe/Rome' \
		-v ${PWD}:/srv/jekyll \
		jekyll/jekyll:stable \
		jekyll serve -w --future --drafts
