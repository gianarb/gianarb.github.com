.PHONY: init tail
init:
	rm -rf _site
	@docker run --rm -it --name gianarb_blog -p 81:4000 -v ${PWD}:/srv/jekyll jekyll/jekyll jekyll serve -w --incremental --drafts --unpublished --future 
tail:
	@docker logs -f gianarb_blog
