.PHONY: init

serve:
	@docker run --rm -it --name www-gianarb \
		-p 4000:4000 \
		-v www-gianarb-cache:/usr/local/bundle \
		-e TZ='Europe/Rome' \
		-v ${PWD}:/srv/jekyll \
		jekyll/jekyll:stable \
		jekyll serve -w --future --drafts

build:
	npm install
	cp -r ./node_modules/jquery/dist/jquery.js ./js/jquery.js
	cp -r ./node_modules/bootstrap/dist/js/bootstrap.js ./js/bootstrap.js
	cp ./node_modules/@fortawesome/fontawesome-free/js/all.js ./js/all.js
	cp -r ./node_modules/@fortawesome/fontawesome-free/webfonts ./fonts

sass:
	sass scss/custom.scss css/style.css
